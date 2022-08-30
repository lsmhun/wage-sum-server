package main

import (
	"log"

	"github.com/caarlos0/env/v6"
	wagesumapp "github.com/lsmhun/wage-sum-server/internal/app/wagesum"
	config "github.com/lsmhun/wage-sum-server/internal/pkg/configuration"
)

var (
	version string
	build   string
)

func main() {
	log.Printf("WageSum application version %s build %s", version, build)

	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Printf("%+v\n", err)
	}

	wagesumapp.WageSumApp(cfg)
}
