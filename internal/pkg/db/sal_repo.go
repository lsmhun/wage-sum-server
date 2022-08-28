package db

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type SalDb struct {
	db  *gorm.DB
	err error
}

type SalDber interface {
	GetSalaryByEmpId(empId int64) decimal.Decimal
}

func NewSalDb(database *gorm.DB, e error) SalDb {
	mySalDB := SalDb{
		db:  database,
		err: e,
	}
	// Migrate the schema
	mySalDB.db.AutoMigrate(&Sal{})
	return mySalDB
}

func (d *SalDb) GetSalaryByEmpId(empId int64) decimal.Decimal {
	var sal Sal
	d.db.First(&sal, "empId = ?", empId)
	if &sal == nil {
		return decimal.Decimal{}
	}
	return sal.Value
}
