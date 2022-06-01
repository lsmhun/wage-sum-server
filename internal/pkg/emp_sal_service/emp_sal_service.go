package empsalservice

import (
	"fmt"
	"sync"

	"github.com/lsmhun/wage-sum-server/internal/pkg/db"
	"github.com/lsmhun/wage-sum-server/internal/pkg/openapi"

	"github.com/shopspring/decimal"
)

// BEGIN: interface for hiding of direct DB

type dbRepository interface {
	GetSalaryByEmpId(empId int64) decimal.Decimal
	FindEmployeesByMgrId(mgrId int64) []openapi.Emp
}

type salDB struct{}

func (s salDB) GetSalaryByEmpId(empId int64) decimal.Decimal {
	return db.GetSalaryByEmpId(empId)
}

func (s salDB) FindEmployeesByMgrId(mgrId int64) []openapi.Emp {
	return db.FindEmployeesByMgrId(mgrId)
}

// END: direct DB hiding

var salaryDB dbRepository

func init() {
	salaryDB = salDB{}
}

func GetSalaryByEmpId(empId int64) decimal.Decimal {
	return salaryDB.GetSalaryByEmpId(empId)
}

func GetSumSalariesByMgrId(mgrId int64) decimal.Decimal {
	var wg sync.WaitGroup

	//var salChannel chan decimal.Decimal

	//var workChannel chan bool
	//var resChannel chan decimal.Decimal

	salChannel := make(chan decimal.Decimal)
	// unbuffered channel
	resChannel := make(chan decimal.Decimal, 1)

	fmt.Printf("root wg.add(1)\n")
	wg.Add(1)
	go getSumSalariesByMgrIdRec(mgrId, salChannel, resChannel, wg)

	/*go func() {
		for i := range salChannel {
			fmt.Println(i)
		}
	}()*/
	wg.Wait()
	//wg.Add(1)
	sumSalaryForEmployee(salChannel, resChannel, wg)
	//res := <-resChannel
	//fmt.Printf("result %d", res)
	return decimal.NewFromInt(11)
}

// Producer
func addSalaryForEmployee(empId int64, salChannel chan<- decimal.Decimal, wg sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("addSalaryForEmployee Writing into salChannel %d \n", empId)
	salChannel <- GetSalaryByEmpId(empId)
}

// Consumer
func sumSalaryForEmployee(salChannel <-chan decimal.Decimal, res chan decimal.Decimal, wg sync.WaitGroup) {
	//defer wg.Done()
	fmt.Printf(" salChannel size: %d \n", cap(salChannel))
	for sal := range salChannel {
		current := <-res
		fmt.Printf("Working with res channel %d \n", current.BigInt().Uint64())
		current = current.Add(sal)
		res <- current
	}
}

func getSumSalariesByMgrIdRec(mgrId int64, salChannel chan decimal.Decimal, res chan decimal.Decimal, wg sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("getSumSalariesByMgrIdRec Writing into mgrId %d \n", mgrId)
	var employees []openapi.Emp = salaryDB.FindEmployeesByMgrId(mgrId)
	for _, emp := range employees {
		if emp.Status == "ACTIVE" {
			fmt.Printf("mgr wg.add(1)\n")
			wg.Add(1) // one for produce, one for consume
			go addSalaryForEmployee(emp.EmpId, salChannel, wg)
			if emp.Type == "MANAGER" {
				wg.Add(1)
				fmt.Printf("sal wg.add(1)\n")
				go getSumSalariesByMgrIdRec(emp.EmpId, salChannel, res, wg)
			}
		}
	}
}
