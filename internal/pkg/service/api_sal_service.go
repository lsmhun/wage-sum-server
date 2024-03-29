/*
 * Employee wage sum
 *
 * Wage sum - demo application with GO language
 *
 * API version: 0.2.0
 * Contact: lsmhun@github
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package service

import (
	"context"
	"errors"
	"net/http"

	db "github.com/lsmhun/wage-sum-server/internal/pkg/db"
	openapi "github.com/lsmhun/wage-sum-server/internal/pkg/openapi"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"

	empSalService "github.com/lsmhun/wage-sum-server/internal/pkg/emp_sal_service"
)

// SalApiService is a service that implements the logic for the SalApiServicer
// This service should implement the business logic for every endpoint for the SalApi API.
// Include any external packages or services that will be required by this service.
type SalApiService struct {
	salDb         db.SalDber
	empSalService empSalService.EmpSalService
}

// NewSalApiService creates a default api service
func NewSalApiService(s db.SalDb, es empSalService.EmpSalService) openapi.SalApiServicer {
	return &SalApiService{
		salDb:         &s,
		empSalService: es,
	}
}

// GetSalByEmpId - Find sal by ID
func (s *SalApiService) GetSalByEmpId(ctx context.Context, empId int64) (openapi.ImplResponse, error) {
	salValue := s.salDb.GetSalaryByEmpId(empId)
	return openapi.Response(200, salValue), nil
}

// GetWageSumByMgrId - Find sum sal by manager ID
func (s *SalApiService) GetWageSumByMgrId(ctx context.Context, empId int64) (openapi.ImplResponse, error) {
	wageSumValue := s.empSalService.GetSumSalariesByMgrId(empId)
	return openapi.Response(200, wageSumValue), nil
}

// UpdateSalWithForm - Updates a sal in the store with form data
func (s *SalApiService) UpdateSalWithForm(ctx context.Context, empId int64, value string) (openapi.ImplResponse, error) {
	salary, errSal := decimal.NewFromString(value)
	if errSal != nil {
		return openapi.Response(http.StatusInternalServerError, nil), errSal
	}
	salById, err := s.salDb.CreateOrUpdateSalary(empId, salary)
	if err == nil {
		return openapi.Response(200, salById), nil
	} else {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return openapi.Response(404, nil), nil
		}
		return openapi.Response(http.StatusInternalServerError, nil), err
	}
}

// DeleteSal - Deletes a sal
func (s *SalApiService) DeleteSal(ctx context.Context, empId int64) (openapi.ImplResponse, error) {
	_, err := s.salDb.DeleteByEmpId(empId)
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), err
	}
	return openapi.Response(200, empId), nil
}
