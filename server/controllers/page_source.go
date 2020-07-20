package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/karngyan/getx/server/clients"
	"github.com/karngyan/getx/server/exchanges"
	"github.com/karngyan/getx/server/models"
	"github.com/karngyan/getx/server/utils"
	"log"
	"net/http"
)

type PageSourceController struct{}

// constructor
func NewPageSourceController() *PageSourceController {
	return &PageSourceController{}
}

// Spawns a go routine to fetch html of given url and dumps it to a file
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
		Id:         id.String(),
		Uri:        r.Uri,
		RetryLimit: r.RetryLimit,
	}

	filePath := utils.GetFilePath(ps.Id)

	// spawn a go routine to fetch the html
	go func() {
		// blocking call based on retries
		htmlBytes, err := networkClient.GetHtmlBytes(ps.Uri, ps.RetryLimit)
		if err != nil {
			log.Println("Network client failed to get html bytes: ",
				ps.Uri)
		}

		if _, err = fileClient.SaveFile(filePath, htmlBytes); err != nil {
			log.Println("Save File failed for ", ps.Uri)
		}
	}()

	// expected file path
	ps.SourceUri = utils.GetSourceUri(ps.Id)
	psJson, err := json.Marshal(ps)
	if err != nil {
		log.Println("JSON Marshal failed: ", err)
	}

	w.Header().Set("Location", ps.SourceUri)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", psJson)

}
