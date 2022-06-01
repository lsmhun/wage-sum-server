package wagesumapp

import (
	"fmt"
	"log"
	//"net/http"
	//"github.com/lsmhun/wage-sum-server/internal/pkg/swagger"
)

func WageSumApp() {
	fmt.Println("Wagesum app started")
	startServer()
}

func startServer() {
	log.Printf("Server started")

	//router := swagger.NewRouter()

	//log.Fatal(http.ListenAndServe(":8080", router))
}
