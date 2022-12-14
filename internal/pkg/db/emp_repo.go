package db

import (
	"time"

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
	err := myEmpDB.db.AutoMigrate(&openapi.Emp{})
	if err != nil {
		panic("Unable to run AutoMigrate on SalDb")
	}
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
	var emp1, err1 = d.FindEmployeeById(emp.EmpId)
	if err1 == nil {
		var res1 = d.db.Model(&emp1).Updates(openapi.Emp{
			FirstName: emp.FirstName,
			LastName:  emp.LastName,
			UserName:  emp.UserName,
			Status:    emp.Status,
			Type:      emp.Type,
			MgrId:     emp.MgrId,
		})
		return emp1, res1.Error
	} else {
		var iid = d.makeTimestamp()
		emp.EmpId = iid
		emp.Id = iid
		var res2 = d.db.FirstOrCreate(&emp, emp)
		return emp, res2.Error
	}
}

func (d *EmpDb) DeleteEmp(empId int64) (openapi.Emp, error) {
	var emp1, err1 = d.FindEmployeeById(empId)
	if err1 == nil {
		err1 = d.db.Delete(&emp1, emp1.Id).Error
	}
	return emp1, err1
}

func (d *EmpDb) makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
