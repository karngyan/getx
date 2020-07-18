package clients

import (
	"errors"
	"io/ioutil"
)

type FileClient struct{}

func NewFileClient() *FileClient {
	return &FileClient{}
}

// creates/updates fileName with fileContent at app/files/*filepath
// returns
//  - true, nil if file created
//  - false, error if not
func (fc *FileClient) CreateFile(fileName string, fileContent []byte) (bool, error) {
	err := ioutil.WriteFile("app/files/"+fileName, fileContent, 0644)
	if err != nil {
		return false, errors.New("file creation failed: " + fileName)
	}
	return true, nil
}
