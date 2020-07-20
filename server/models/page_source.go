package models

const (
	DefaultRetryLimit = 4
	RetryUpperLimit   = 10
)

type PageSource struct {
	Id         string `json:"id"`
	Uri        string `json:"uri"`
	SourceUri  string `json:"sourceUri"`
	RetryLimit int    `json:"retryLimit"`
}
