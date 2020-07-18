package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/karngyan/getx/app/controllers"
)

const (
	ServerPort = ":7771"
)

func StartApp() {
	pageSourceController := controllers.NewPageSourceController()

	router := httprouter.New()
	router.POST("/pagesource", pageSourceController.GeneratePageSource)
	router.ServeFiles("/files/*filepath", http.Dir("app/files"))

	fmt.Println("Server is running on port:", ServerPort)
	log.Fatal(http.ListenAndServe(ServerPort, router))
}
