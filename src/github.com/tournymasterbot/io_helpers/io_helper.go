package iohelper

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
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
	if err == nil {
		success = true
	}
	return success, err
}

// Chmod : Sets file permissions to the specified value
func Chmod(filepath string, perm os.FileMode) error {
	err := os.Chmod(filepath, perm)
	return err
}

// CompressFile : Compresses (Zips) a file using gzip
func CompressFile(source, target string) error {
	reader, err := os.Open(source)
	if err != nil {
		return err
	}

	CreateFolderIfNotExists(filepath.Dir(target))

	filename := filepath.Base(source)
	writer, err := os.Create(target)
	if err != nil {
		return err
	}
	defer writer.Close()

	archiver := gzip.NewWriter(writer)
	archiver.Name = filename
	defer archiver.Close()

	_, err = io.Copy(archiver, reader)
	return err
}

// InflateFile : Inflates (Unzips) a gzip file
func InflateFile(source, target string) error {
	reader, err := os.Open(source)
	if err != nil {
		return err
	}
	defer reader.Close()

	archive, err := gzip.NewReader(reader)
	if err != nil {
		return err
	}
	defer archive.Close()

	CreateFolderIfNotExists(filepath.Dir(target))
	writer, err := os.Create(target)
	if err != nil {
		return err
	}
	defer writer.Close()

	_, err = io.Copy(writer, archive)
	return err
}

// CompressText : Returns a compressed byte array for the specified string
func CompressText(data string) ([]byte, error) {
	return CompressBytes([]byte(data))
}

// CompressBytes : Compresses bytes and returns the byte array result
func CompressBytes(inflatedBytes []byte) ([]byte, error) {
	var buffer bytes.Buffer
	gz := gzip.NewWriter(&buffer)
	if _, err := gz.Write(inflatedBytes); err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// CreateFolderIfNotExists : Creates the specified folder if it doesn't already exist
func CreateFolderIfNotExists(directoryPath string) error {
	_, err := os.Stat(directoryPath)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(directoryPath, 0755)
		return errDir
	}
	return err
}

// InflateText : Inflates bytes and returns a string result
func InflateText(compressedBytes []byte) (string, error) {
	var result string
	inflatedBytes, err := InflateBytes(compressedBytes)
	if err == nil {
		result = string(inflatedBytes[:])
	}
	return result, err
}

// InflateBytes : Inflates bytes and returns the result
func InflateBytes(compressedBytes []byte) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	reader := bytes.NewReader(compressedBytes)
	gzipreader, gzipErr := gzip.NewReader(reader)
	writer := bufio.NewWriter(buffer)
	io.Copy(writer, gzipreader)
	return buffer.Bytes(), gzipErr
}
