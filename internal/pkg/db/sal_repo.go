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
	DeleteByEmpId(empId int64) (Sal, error)
	CreateOrUpdateSalary(empId int64, value decimal.Decimal) (Sal, error)
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
	d.db.First(&sal, "emp_id = ?", empId)
	if d.err != nil {
		return decimal.Decimal{}
	}
	return sal.Value
}

func (d *SalDb) DeleteByEmpId(empId int64) (Sal, error) {
	var sal, err1 = d.findSalaryByEmpId(empId)
	if err1 != nil {
		return sal, err1
	}
	err1 = d.db.Delete(&sal, 1).Error
	return sal, err1
}

func (d *SalDb) CreateOrUpdateSalary(empId int64, value decimal.Decimal) (Sal, error) {
	var sal, err1 = d.findSalaryByEmpId(empId)
	if err1 != nil {
		sal = Sal{
			EmpId: empId,
			Value: value,
		}
		result := d.db.FirstOrCreate(&sal, sal)
		return sal, result.Error
	}
	err1 = d.db.Model(&sal).Update("value", value).Error
	return sal, err1
}

func (d *SalDb) findSalaryByEmpId(empId int64) (Sal, error) {
	var sal Sal
	err := d.db.First(&sal, "emp_id = ?", empId).Error
	return sal, err
}
