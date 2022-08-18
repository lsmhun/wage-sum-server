package wagesumapp

import (
	"log"
	"net/http"

	myConfig "github.com/lsmhun/wage-sum-server/internal/pkg/configuration"
	oa "github.com/lsmhun/wage-sum-server/internal/pkg/openapi"
)

func WageSumApp() {
	startServer()
}

func startServer() {
	listeningHttpPort := myConfig.GetConfigValue("wagesum.http.service.port")
	log.Printf("WageSum HTTP Server is starting on port :%s ...", listeningHttpPort)

	// registering new controllers
	router := oa.NewRouter(
		empApiController(),
	)
	log.Fatal(http.ListenAndServe(":"+listeningHttpPort, router))

}

func empApiController() oa.Router {
	empApiService := oa.NewEmpApiService()
	empApiController := oa.NewEmpApiController(empApiService)
	return empApiController
}
