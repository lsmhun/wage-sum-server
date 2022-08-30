package db

import (
	"fmt"

	config "github.com/lsmhun/wage-sum-server/internal/pkg/configuration"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfigurator interface {
	DatabaseSetup() (gorm.DB, error)
}

func PostgresDatabaseSetup(conf config.Config) (*gorm.DB, error) {
	host := conf.DbHost
	port := conf.DbPort
	name := conf.DbName
	user := conf.DbUsername
	password := conf.DbPassword

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
