package empsalservice

import (
	"sync"

	slog "github.com/go-eden/slf4go"
	db "github.com/lsmhun/wage-sum-server/internal/pkg/db"
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
	slog.Debugf("GetSumSalariesByMgrId mgrId= %d", mgrId)
	var wg sync.WaitGroup

	salChannel := make(chan decimal.Decimal)
	// unbuffered channel
	resChannel := make(chan decimal.Decimal, 1)
	// init with zero sum
	resChannel <- decimal.NewFromInt(0)

	// pass a pointer to the WaitGroup, otherwise it won't work
	getSumSalariesByMgrIdRec(mgrId, salChannel, &wg)
	// we have to wait until all salaries are collected
	go func() {
		wg.Wait()
		close(salChannel)
	}()
	// summarizing the salary channel, it can be solved with a mutex as well
	for sallll := range salChannel {
		current := <-resChannel
		current = current.Add(sallll)
		resChannel <- current
	}

	res := <-resChannel
	return res
}

// Producer
func addSalaryForEmployee(empId int64, salChannel chan decimal.Decimal, wg *sync.WaitGroup) {
	defer wg.Done()
	//add salary for employee writing into salChannel
	salChannel <- GetSalaryByEmpId(empId)
}

func getSumSalariesByMgrIdRec(mgrId int64, salChannel chan decimal.Decimal, wg *sync.WaitGroup) {
	var employees []openapi.Emp = salaryDB.FindEmployeesByMgrId(mgrId)
	for _, emp := range employees {
		if emp.Status == "ACTIVE" {
			wg.Add(1) // one for produce, one for consume
			go addSalaryForEmployee(emp.EmpId, salChannel, wg)
			if emp.Type == "MANAGER" {
				// this won't be a goroutin
				getSumSalariesByMgrIdRec(emp.EmpId, salChannel, wg)
			}
		}
	}
}
