package utility

import (
	"io/ioutil"
)

//ReadFile is a wrapper and reads a file for a given path
func ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}
