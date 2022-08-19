package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	myConfig "github.com/lsmhun/wage-sum-server/internal/pkg/configuration"
)

type DatabaseConfigurator interface {
	DatabaseSetup() (gorm.DB, error)
}

func PostgresDatabaseSetup() (*gorm.DB, error) {
	host := myConfig.GetConfigValue("wagesum.db.host")
	port := myConfig.GetConfigValue("wagesum.db.port")
	name := myConfig.GetConfigValue("wagesum.db.name")
	user := myConfig.GetConfigValue("wagesum.db.username")
	password := myConfig.GetConfigValue("wagesum.db.password")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		name,
		password,
	)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

/*func SqliteDatabaseSetup() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}*/
