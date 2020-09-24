package iohelper

import (
	"io/ioutil"
	"os"
)

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

// WriteAllText : (Over)Writes string text to the specified file
// If the file does not exist, create the file.
func WriteAllText(filepath string, data string) (bool, error) {
	var success, err = WriteAllBytes(filepath, []byte(data))
	return success, err
}

// WriteAllBytes : (Over)Writes the byte data to the specified file.
// If the file does not exist, create the file.
func WriteAllBytes(filepath string, data []byte) (bool, error) {
	var success = false
	err := ioutil.WriteFile(filepath, []byte(data), 0700)
	if err != nil {
		success = true
	}
	return success, err
}

// Chmod : Sets file permissions to the specified value
func Chmod(filepath string, perm os.FileMode) error {
	err := os.Chmod(filepath, perm)
	return err
}
