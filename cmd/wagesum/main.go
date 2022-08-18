package main

import (
	"log"

	wagesumapp "github.com/lsmhun/wage-sum-server/internal/app/wagesum"
)

func main() {
	log.Println("WageSum application")
	wagesumapp.WageSumApp()
}
