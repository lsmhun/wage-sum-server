package db

import (
	"github.com/shopspring/decimal"
)

func GetSalaryByEmpId(empId int64) decimal.Decimal {
	return decimal.NewFromFloat(2.0)
}
