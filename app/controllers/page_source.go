package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/karngyan/getx/app/clients"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/karngyan/getx/app/exchanges"
	"github.com/karngyan/getx/app/models"
)

type PageSourceController struct{}

func NewPageSourceController() *PageSourceController {
	return &PageSourceController{}
}

func (sc *PageSourceController) GeneratePageSource(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	r := exchanges.GeneratePageSourceRequest{}
	fileClient := clients.NewFileClient()
	networkClient := clients.NewNetworkClient()

	json.NewDecoder(req.Body).Decode(&r)

	if _, err := r.IsValid(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s\n", err)
		return
	}

	log.Println("GeneratePageSource called for uri:", r.Uri)
	id, _ := uuid.NewUUID()

	ps := models.PageSource{
		Id:        id.String(),
		Uri:       r.Uri,
		RetyLimit: r.RetryLimit,
	}

	// fetching html bytes
	htmlBytes, err := networkClient.GetHtmlBytes(ps.Uri)
	if err != nil {
		// TODO: go routine for retries on the basis of ps.retryLimit
		// TODO: create empty html file instead and update later
		log.Println("Network Client failed to get html bytes: ", ps.Uri)
	}

	// creating file
	_, err = fileClient.CreateFile(ps.Id+".html", htmlBytes)
	if err != nil {
		// 5xx
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "File Couldn't be created%s\n", err)
		return
	}

	// file creation success
	ps.SourceUri = "files/" + ps.Id + ".html"
	psJson, err := json.Marshal(ps)
	if err != nil {
		log.Println("JSON Marshal failed: ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", ps.SourceUri)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", psJson)

}
