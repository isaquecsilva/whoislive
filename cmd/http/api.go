package main

import (
	"log"

	"github.com/isaquecsilva/whoislive/src/controller"
	"github.com/isaquecsilva/whoislive/src/routes"
	"github.com/isaquecsilva/whoislive/src/service"
)

func main() {
	liveService := service.NewLiveService()
	liveController := controller.NewController(liveService)
	if err := routes.StartRoutesAndApp(":8000", "./public", liveController); err != nil {
		log.Fatal(err)
	}
}