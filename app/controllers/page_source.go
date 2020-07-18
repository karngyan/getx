package controllers

import (
	"encoding/json"
	"fmt"
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
	json.NewDecoder(req.Body).Decode(&r)

	if _, err := r.IsValid(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s\n", err)
		return
	}

	log.Println("GeneratePageSource called for uri:", r.Uri)

	// TODO: Create HTML File add sourceUri for the same
	// TODO: go fetchFile and do shite (async) worker queue may be

	id, _ := uuid.NewUUID()
	ps := models.PageSource{
		Id:        id.String(),
		Uri:       r.Uri,
		RetyLimit: r.RetryLimit,
	}

	log.Println(ps)

	psJson, err := json.Marshal(ps)
	if err != nil {
		log.Println("JSON Marshal failed: ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", "sourceUriPlaceholder")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", psJson)

}
