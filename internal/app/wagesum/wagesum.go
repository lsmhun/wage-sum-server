package wagesumapp

import (
	"log"
	"net/http"

	myConfig "github.com/lsmhun/wage-sum-server/internal/pkg/configuration"
	db "github.com/lsmhun/wage-sum-server/internal/pkg/db"
	openapi "github.com/lsmhun/wage-sum-server/internal/pkg/openapi"
	service "github.com/lsmhun/wage-sum-server/internal/pkg/service"

	"gorm.io/gorm"
)

func WageSumApp() {
	startServer()
}

func startServer() {

	database, errr := initializeDatabaseConnection()
	empDbRepo := initEmpDb(database, errr)

	listeningHttpPort := myConfig.GetConfigValue("wagesum.http.service.port")
	log.Printf("WageSum HTTP Server is starting on port :%s ...", listeningHttpPort)

	// registering new controllers
	router := openapi.NewRouter(
		empApiController(empDbRepo),
	)
	log.Fatal(http.ListenAndServe(":"+listeningHttpPort, router))

}

func initializeDatabaseConnection() (*gorm.DB, error) {
	db, err := db.PostgresDatabaseSetup()
	if err != nil {
		panic("Unable to connect to database")
	}
	return db, err
}

func empApiController(empDB db.EmpDb) openapi.Router {
	empApiService := service.NewEmpApiService(empDB)
	empApiController := openapi.NewEmpApiController(empApiService)
	return empApiController
}

func initEmpDb(database *gorm.DB, err error) db.EmpDb {
	return db.NewEmpDb(database, err)
}
