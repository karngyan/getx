package clients

import (
	"errors"
	"io/ioutil"
)

type FileClient struct{}

func NewFileClient() *FileClient {
	return &FileClient{}
}

// creates/updates filePath with fileContent at server/files/*filepath
// returns
//  - true, nil if file created
//  - false, error if not
func (fc *FileClient) SaveFile(filePath string, fileContent []byte) (bool, error) {
	err := ioutil.WriteFile(filePath, fileContent, 0644)
	if err != nil {
		return false, errors.New("file creation failed: " + filePath)
	}
	return true, nil
}
