/*
 * Employee wage sum
 *
 * Wage sum - demo application with GO language
 *
 * API version: 0.3.0
 * Contact: lsmhun@github
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"
	"strings"
)

// MonitoringApiController binds http requests to an api service and writes the service results to the http response
type MonitoringApiController struct {
	service      MonitoringApiServicer
	errorHandler ErrorHandler
}

// MonitoringApiOption for how the controller is set up.
type MonitoringApiOption func(*MonitoringApiController)

// WithMonitoringApiErrorHandler inject ErrorHandler into controller
func WithMonitoringApiErrorHandler(h ErrorHandler) MonitoringApiOption {
	return func(c *MonitoringApiController) {
		c.errorHandler = h
	}
}

// NewMonitoringApiController creates a default api controller
func NewMonitoringApiController(s MonitoringApiServicer, opts ...MonitoringApiOption) Router {
	controller := &MonitoringApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the MonitoringApiController
func (c *MonitoringApiController) Routes() Routes {
	return Routes{
		{
			"Health",
			strings.ToUpper("Get"),
			"/api/health",
			c.Health,
		},
	}
}

// Health - Health endpoint
func (c *MonitoringApiController) Health(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.Health(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	err = EncodeJSONResponse(result.Body, &result.Code, w)
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}

}
