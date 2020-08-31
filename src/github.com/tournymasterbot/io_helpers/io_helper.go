package iohelper

import "io/ioutil"

//ReadAllBytes Opens requested file and returns entire file result
func ReadAllBytes(filepath string) ([]byte, error) {
	result, err := ioutil.ReadFile(filepath)
	if err != nil {
		result = nil
	}

	return result, err
}

//ReadAllText Opens requested file and returns entire file as a string
func ReadAllText(filepath string) (string, error) {
	var result string

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		result = ""
	} else {
		result = string(data)
	}

	return result, err
}
