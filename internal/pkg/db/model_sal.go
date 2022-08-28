package db

import (
	"github.com/shopspring/decimal"
)

type Sal struct {
	Id int64 `json:"id,omitempty" gorm:"primaryKey;autoIncrement:true;unique"`

	EmpId int64 `json:"empId,omitempty" gorm:"unique"`

	Value decimal.Decimal `json:"value,omitempty"`
}
