package server

import (
	"fmt"
	"github.com/rs/cors"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/karngyan/getx/server/controllers"
)

const (
	Port = ":8080"
)

func StartApp() {
	pageSourceController := controllers.NewPageSourceController()

	router := httprouter.New()
	router.POST("/pagesource", pageSourceController.GeneratePageSource)
	router.ServeFiles("/files/*filepath", http.Dir("server/files"))

	handler := cors.Default().Handler(router)
	fmt.Println("Server is running on port", Port)
	log.Fatal(http.ListenAndServe(Port, handler))
}
