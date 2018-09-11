package spider

import (
	"io/ioutil"
	"os"
	"path"
)
func saveToLocal(output, name string, data []byte) error {
	filePath := path.Join(output, name)
	return ioutil.WriteFile(filePath, data, 0644)
}

func appendUrlToLocal(output, name string, data []byte) error {
	filePath := path.Join(output, name)
	return ioutil.WriteFile(filePath, data, os.ModeAppend)
}