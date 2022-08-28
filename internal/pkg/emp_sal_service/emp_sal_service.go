package empsalservice

import (
	"sync"

	slog "github.com/go-eden/slf4go"
	db "github.com/lsmhun/wage-sum-server/internal/pkg/db"
	"github.com/shopspring/decimal"
)

type EmpSalService struct {
	empDb    db.EmpDber
	salaryDb db.SalDber
}

type EmpSalServicer interface {
}

func NewEmpSalService(edb db.EmpDber, sdb db.SalDber) EmpSalService {
	myEmpSalService := EmpSalService{
		empDb:    edb,
		salaryDb: sdb,
	}
	return myEmpSalService
}

func (es *EmpSalService) GetSalaryByEmpId(empId int64) decimal.Decimal {
	return es.salaryDb.GetSalaryByEmpId(empId)
}

func (es *EmpSalService) GetSumSalariesByMgrId(mgrId int64) decimal.Decimal {
	slog.Debugf("GetSumSalariesByMgrId mgrId= %d", mgrId)
	var wg sync.WaitGroup

	salChannel := make(chan decimal.Decimal)
	// unbuffered channel
	resChannel := make(chan decimal.Decimal, 1)
	// init with zero sum
	resChannel <- decimal.NewFromInt(0)

	// pass a pointer to the WaitGroup, otherwise it won't work
	es.getSumSalariesByMgrIdRec(mgrId, salChannel, &wg)
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
func (es *EmpSalService) addSalaryForEmployee(empId int64, salChannel chan decimal.Decimal, wg *sync.WaitGroup) {
	defer wg.Done()
	//add salary for employee writing into salChannel
	salChannel <- es.GetSalaryByEmpId(empId)
}

func (es *EmpSalService) getSumSalariesByMgrIdRec(mgrId int64, salChannel chan decimal.Decimal, wg *sync.WaitGroup) {
	var employees, err1 = es.empDb.FindEmployeesByMgrId(mgrId)
	if err1 == nil {
		for _, emp := range employees {
			if emp.Status == "ACTIVE" {
				wg.Add(1) // one for produce, one for consume
				go es.addSalaryForEmployee(emp.EmpId, salChannel, wg)
				if emp.Type == "MANAGER" {
					// this won't be a goroutin
					es.getSumSalariesByMgrIdRec(emp.EmpId, salChannel, wg)
				}
			}
		}
	}
}
