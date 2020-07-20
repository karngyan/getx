package exchanges

import (
	"fmt"
	"github.com/karngyan/getx/server/models"
)

type GeneratePageSourceRequest struct {
	Uri        string `json:"uri"`
	RetryLimit int    `json:"retryLimit"`
}

func (r *GeneratePageSourceRequest) IsValid() (bool, error) {

	if r.Uri == "" {
		return false, fmt.Errorf("bad request: uri is missing")
	}

	if r.RetryLimit > models.RetryUpperLimit {
		return false, fmt.Errorf("bad request: retry limit too high")
	}

	if r.RetryLimit <= 0 {
		r.RetryLimit = models.DefaultRetryLimit
	}

	return true, nil
}
