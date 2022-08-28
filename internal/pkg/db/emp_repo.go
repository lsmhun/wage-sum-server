package db

import (
	openapi "github.com/lsmhun/wage-sum-server/internal/pkg/openapi"
	"gorm.io/gorm"
)

type EmpDber interface {
	FindEmployeeById(id int64) openapi.Emp
	FindEmployeesByMgrId(id int64) []openapi.Emp
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

func (d *EmpDb) FindEmployeeById(id int64) openapi.Emp {
	var emp openapi.Emp
	d.db.First(&emp, "empId = ?", id)
	return emp
	//return openapi.Emp{}
}

func (d *EmpDb) FindEmployeesByMgrId(id int64) []openapi.Emp {
	return nil
}
