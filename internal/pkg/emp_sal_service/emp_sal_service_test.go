package empsalservice

import (
	"testing"

	"github.com/lsmhun/wage-sum-server/internal/pkg/openapi"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

// This helps in assigning mock at the runtime instead of compile time
var getSalaryByEmpIdMock func(empId int64) decimal.Decimal
var findEmployeesByMgrIdMock func(mgrId int64) []openapi.Emp

type salaryDbMock struct{}

func (s salaryDbMock) GetSalaryByEmpId(empId int64) decimal.Decimal {
	return getSalaryByEmpIdMock(empId)
}

func (s salaryDbMock) FindEmployeesByMgrId(mgrId int64) []openapi.Emp {
	return findEmployeesByMgrIdMock(mgrId)
}

func TestGetSalaryByEmpId(t *testing.T) {
	// define inputs
	var empId int64 = 1

	// define expected result
	var expected decimal.Decimal = decimal.NewFromFloat(3.0)

	//mocking
	salaryDB = salaryDbMock{}
	getSalaryByEmpIdMock = func(empId int64) decimal.Decimal {
		return decimal.NewFromFloat(3.0)
	}

	// perform test
	actual := GetSalaryByEmpId(empId)

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
	salaryDB = salaryDbMock{}
	getSalaryByEmpIdMock = func(empId int64) decimal.Decimal {
		return decimal.NewFromInt(empId)
	}

	findEmployeesByMgrIdMock = func(mgrId int64) []openapi.Emp {
		empBoss := openapi.Emp{EmpId: 1, Status: "ACTIVE", Type: "MANAGER"}
		empManager1 := openapi.Emp{EmpId: 2, MgrId: 1, Status: "ACTIVE", Type: "MANAGER"}
		empClerk1 := openapi.Emp{EmpId: 3, MgrId: 2, Status: "ACTIVE", Type: "EMPLOYEE"}
		empManager2 := openapi.Emp{EmpId: 4, MgrId: 1, Status: "ACTIVE", Type: "MANAGER"}
		employees := [4]openapi.Emp{empBoss, empManager1, empClerk1, empManager2}
		switch mgrId {
		case 1:
			return []openapi.Emp{empManager1, empManager2}
		case 2:
			return employees[2:2]
		default:
			return []openapi.Emp{}
		}
	}

	// perform test
	actual := GetSumSalariesByMgrId(1)
	expected = decimal.NewFromInt(1 + 2 + 3 + 4)
	assert.Equal(t, expected, actual)
	/*
		actual = GetSumSalariesByMgrId(2)
		expected = decimal.NewFromInt(2 + 3)
		assert.Equal(t, expected, actual)

		actual = GetSumSalariesByMgrId(4)
		expected = decimal.NewFromInt(4)
		assert.Equal(t, expected, actual)
	*/
}
