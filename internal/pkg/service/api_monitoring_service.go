/*
 * Employee wage sum
 *
 * Wage sum - demo application with GO language
 *
 * API version: 0.3.0
 * Contact: lsmhun@github
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package service

import (
	"context"

	openapi "github.com/lsmhun/wage-sum-server/internal/pkg/openapi"
)

// MonitoringApiService is a service that implements the logic for the MonitoringApiServicer
// This service should implement the business logic for every endpoint for the MonitoringApi API.
// Include any external packages or services that will be required by this service.
type MonitoringApiService struct {
}

// NewMonitoringApiService creates a default api service
func NewMonitoringApiService() openapi.MonitoringApiServicer {
	return &MonitoringApiService{}
}

// Health - Health endpoint
func (s *MonitoringApiService) Health(ctx context.Context) (openapi.ImplResponse, error) {
	return openapi.Response(200, map[string]string{"status": "UP"}), nil
}
