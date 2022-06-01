package db

import (
	"github.com/lsmhun/wage-sum-server/internal/pkg/openapi"
)

func FindEmployeeById(id int64) openapi.Emp {
	return openapi.Emp{}
}

func FindEmployeesByMgrId(id int64) []openapi.Emp {
	return nil
}
