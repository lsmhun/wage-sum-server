package empsalservice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/shopspring/decimal"

	db "github.com/lsmhun/wage-sum-server/internal/pkg/db"
	openapi "github.com/lsmhun/wage-sum-server/internal/pkg/openapi"
)

// This helps in assigning mock at the runtime instead of compile time
var getSalaryByEmpIdMock func(empId int64) decimal.Decimal
var findEmployeesByMgrIdMock func(mgrId int64) []openapi.Emp

type salaryDbMock struct {
	db.SalDb
}

func (s *salaryDbMock) GetSalaryByEmpId(empId int64) decimal.Decimal {
	return getSalaryByEmpIdMock(empId)
}

type empDbMock struct {
	db.EmpDb
}

func (e *empDbMock) FindEmployeesByMgrId(mgrId int64) []openapi.Emp {
	return findEmployeesByMgrIdMock(mgrId)
}

func TestGetSalaryByEmpId(t *testing.T) {
	// define inputs
	var empId int64 = 1

	// define expected result
	var expected decimal.Decimal = decimal.NewFromFloat(3.0)

	//mocking
	getSalaryByEmpIdMock = func(empId int64) decimal.Decimal {
		return decimal.NewFromFloat(3.0)
	}

	origSalDb := db.SalDb{}
	tSalDb := &salaryDbMock{origSalDb}
	origEmpDb := db.EmpDb{}
	tEmpDb := &empDbMock{origEmpDb}
	ess := NewEmpSalService(tEmpDb, tSalDb)

	// perform test
	actual := ess.GetSalaryByEmpId(empId)

	// assert that the actual result is equal to expected
	assert.Equal(t, expected, actual)
}

func TestGetSumSalariesByMgrId(t *testing.T) {

	//   Boss - id=1
	//     Manager1 - id=2
	//       Clerk1 - id=3
	//     Manager2 - id=4

	// define expected result
	var expected decimal.Decimal = decimal.NewFromFloat(3.0)

	//mocking
	getSalaryByEmpIdMock = func(empId int64) decimal.Decimal {
		return decimal.NewFromInt(empId)
	}

	findEmployeesByMgrIdMock = func(mgrId int64) []openapi.Emp {
		//empBoss := openapi.Emp{EmpId: 1, Status: "ACTIVE", Type: "MANAGER"}
		empManager1 := openapi.Emp{EmpId: 2, MgrId: 1, Status: "ACTIVE", Type: "MANAGER"}
		empClerk1 := openapi.Emp{EmpId: 3, MgrId: 2, Status: "ACTIVE", Type: "EMPLOYEE"}
		empManager2 := openapi.Emp{EmpId: 4, MgrId: 1, Status: "ACTIVE", Type: "MANAGER"}
		//employees := [4]openapi.Emp{empBoss, empManager1, empClerk1, empManager2}
		switch mgrId {
		case 1:
			return []openapi.Emp{empManager1, empManager2}
		case 2:
			return []openapi.Emp{empClerk1}
		default:
			return []openapi.Emp{}
		}
	}

	origSalDb := db.SalDb{}
	tSalDb := &salaryDbMock{origSalDb}
	origEmpDb := db.EmpDb{}
	tEmpDb := &empDbMock{origEmpDb}
	ess := NewEmpSalService(tEmpDb, tSalDb)

	// perform test
	actual := ess.GetSumSalariesByMgrId(1)
	fmt.Printf("actual=%d", actual)
	expected = decimal.NewFromInt(2 + 3 + 4)
	assert.Equal(t, expected, actual)

	actual = ess.GetSumSalariesByMgrId(2)
	expected = decimal.NewFromInt(3)
	assert.Equal(t, expected, actual)

	actual = ess.GetSumSalariesByMgrId(4)
	expected = decimal.NewFromInt(0)
	assert.Equal(t, expected, actual)

}
