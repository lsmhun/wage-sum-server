package empsalservice

import (
	"fmt"

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
var salChannel chan decimal.Decimal
var workChannel chan bool
var resChannel chan decimal.Decimal

func init() {
	salaryDB = salDB{}
	salChannel = make(chan decimal.Decimal)
	workChannel = make(chan bool)
	// unbuffered channel
	resChannel = make(chan decimal.Decimal, 1)

	go sumSalaryForEmployee(salChannel, resChannel, workChannel)
}

// Producer
func addSalaryForEmployee(empId int64, salChannel chan<- decimal.Decimal, workChannel chan<- bool) {
	fmt.Printf("addSalaryForEmployee Writing into salChannel %d \n", empId)
	salChannel <- GetSalaryByEmpId(empId)
	fmt.Printf("addSalaryForEmployee Writing into workChannel %d \n", empId)
	workChannel <- true
}

// Consumer
func sumSalaryForEmployee(salChannel <-chan decimal.Decimal, res chan decimal.Decimal, workChannel <-chan bool) {
	fmt.Printf(" salChannel size: %d \n", cap(salChannel))
	for sal := range salChannel {
		fmt.Printf("Working with res channel %d \n", sal.IntPart())
		current := <-res
		current = current.Add(sal)
		res <- current
		fmt.Printf("Reading from workChannel \n")
		<-workChannel
	}
}

func GetSumSalariesByMgrId(mgrId int64) decimal.Decimal {
	var employees []openapi.Emp = salaryDB.FindEmployeesByMgrId(mgrId)

	for _, emp := range employees {
		if emp.Status == "ACTIVE" {
			go addSalaryForEmployee(emp.EmpId, salChannel, workChannel)
			if emp.Type == "MANAGER" {
				go GetSumSalariesByMgrId(emp.EmpId)
			}
		}
	}
	<-workChannel
	res := <-resChannel
	fmt.Printf("result %d", res)
	return res
}

func GetSalaryByEmpId(empId int64) decimal.Decimal {
	return salaryDB.GetSalaryByEmpId(empId)
}
