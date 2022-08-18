package db

import (
	openapi "github.com/lsmhun/wage-sum-server/internal/pkg/openapi"
)

type EmpDb struct {
}

func FindEmployeeById(id int64) openapi.Emp {
	return openapi.Emp{}
}

func FindEmployeesByMgrId(id int64) []openapi.Emp {
	return nil
}
