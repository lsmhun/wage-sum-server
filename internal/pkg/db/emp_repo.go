package db

import (
	openapi "github.com/lsmhun/wage-sum-server/internal/pkg/openapi"
	"gorm.io/gorm"
)

type EmpDber interface {
	FindEmployeeById(empId int64) (openapi.Emp, error)
	FindEmployeesByMgrId(mgrId int64) ([]openapi.Emp, error)
	FindEmployeesByType(empType string) ([]openapi.Emp, error)
	CreateOrUpdateEmp(emp openapi.Emp) (openapi.Emp, error)
	DeleteEmp(empId int64) (openapi.Emp, error)
}

type EmpDb struct {
	db  *gorm.DB
	err error
}

func NewEmpDb(database *gorm.DB, e error) EmpDb {
	myEmpDB := EmpDb{
		db:  database,
		err: e,
	}
	// Migrate the schema
	myEmpDB.db.AutoMigrate(&openapi.Emp{})
	return myEmpDB
}

func (d *EmpDb) FindEmployeeById(id int64) (openapi.Emp, error) {
	var emp openapi.Emp
	var err1 = d.db.First(&emp, "emp_id = ?", id).Error
	return emp, err1
	//return openapi.Emp{}
}

func (d *EmpDb) FindEmployeesByMgrId(mgrId int64) ([]openapi.Emp, error) {
	var emps []openapi.Emp
	var err1 = d.db.Where("mgr_id = ?", mgrId).Find(&emps).Error
	return emps, err1
}

func (d *EmpDb) FindEmployeesByType(empType string) ([]openapi.Emp, error) {
	var emps []openapi.Emp
	var err1 = d.db.Where("type = ?", empType).Find(&emps).Error
	return emps, err1
}

func (d *EmpDb) CreateOrUpdateEmp(emp openapi.Emp) (openapi.Emp, error) {
	var emp1 openapi.Emp
	var err1 error
	if emp.EmpId != 0 {
		emp1, err1 = d.FindEmployeeById(emp.EmpId)
		if err1 != nil {
			return emp1, err1
		}
		d.db.Model(&emp1).Updates(emp)
	} else {
		err1 = d.db.Create(&emp).Error
	}
	return emp1, err1
}

func (d *EmpDb) DeleteEmp(empId int64) (openapi.Emp, error) {
	var emp1, err1 = d.FindEmployeeById(empId)
	if err1 != nil {
		d.db.Delete(&emp1)
	}
	return emp1, d.err
}
