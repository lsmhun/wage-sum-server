package db

import (
	"github.com/shopspring/decimal"
)

type Sal struct {
	Id int64 `json:"id,omitempty"`

	EmpId int64 `json:"empId,omitempty"`

	Value decimal.Decimal `json:"value,omitempty"`
}
