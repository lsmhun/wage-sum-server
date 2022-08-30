package wagesumapp

import (
	"errors"
	"log"
	"net/http"

	config "github.com/lsmhun/wage-sum-server/internal/pkg/configuration"
	db "github.com/lsmhun/wage-sum-server/internal/pkg/db"
	empSalService "github.com/lsmhun/wage-sum-server/internal/pkg/emp_sal_service"
	openapi "github.com/lsmhun/wage-sum-server/internal/pkg/openapi"
	service "github.com/lsmhun/wage-sum-server/internal/pkg/service"
	"github.com/shopspring/decimal"

	"gorm.io/gorm"
)

func WageSumApp(conf config.Config) {
	startServer(conf)
}

func startServer(conf config.Config) {

	database, errr := initializeDatabaseConnection(conf)
	empDbRepo := initEmpDb(database, errr)
	salDbRepo := initSalDb(database, errr)
	empSalService := initEmpSalService(empDbRepo, salDbRepo)

	initDbWithDemoData(empDbRepo, salDbRepo)

	listeningHttpPort := conf.HttpServerPort
	log.Printf("WageSum HTTP Server is starting on port :%s ...", listeningHttpPort)

	// registering new controllers
	router := openapi.NewRouter(
		empApiController(empDbRepo),
		salApiController(salDbRepo, empSalService),
	)
	log.Fatal(http.ListenAndServe(":"+listeningHttpPort, router))

}

func initializeDatabaseConnection(conf config.Config) (*gorm.DB, error) {
	db, err := db.PostgresDatabaseSetup(conf)
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

func salApiController(salDB db.SalDb, empSalService empSalService.EmpSalService) openapi.Router {
	salApiService := service.NewSalApiService(salDB, empSalService)
	salApiController := openapi.NewSalApiController(salApiService)
	return salApiController
}

func initSalDb(database *gorm.DB, err error) db.SalDb {
	return db.NewSalDb(database, err)
}

func initEmpSalService(empDB db.EmpDb, salDB db.SalDb) empSalService.EmpSalService {
	empSalService := empSalService.NewEmpSalService(&empDB, &salDB)
	return empSalService
}

func initDbWithDemoData(empDB db.EmpDb, salDB db.SalDb) {
	// This is just for demo purposes
	_, errr := empDB.FindEmployeeById(1)
	if errr != nil {
		if errors.Is(errr, gorm.ErrRecordNotFound) {
			empBoss := openapi.Emp{EmpId: 1, Status: "ACTIVE", Type: "MANAGER"}
			empManager1 := openapi.Emp{EmpId: 2, MgrId: 1, Status: "ACTIVE", Type: "MANAGER"}
			empClerk1 := openapi.Emp{EmpId: 3, MgrId: 2, Status: "ACTIVE", Type: "EMPLOYEE"}
			empManager2 := openapi.Emp{EmpId: 4, MgrId: 1, Status: "ACTIVE", Type: "MANAGER"}
			employees := [4]openapi.Emp{empBoss, empManager1, empClerk1, empManager2}
			for _, emp := range employees {
				_, pErr := empDB.CreateOrUpdateEmp(emp)
				if pErr == nil {
					var salValue int64 = 240_000 / emp.EmpId
					salDB.CreateOrUpdateSalary(emp.EmpId, decimal.NewFromInt(salValue))
				}
			}
		} else {
			log.Println("Demo data has been already registered")
		}
	}
}
