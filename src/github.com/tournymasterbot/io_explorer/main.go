package main

import (
	"fmt"

	iohelper "github.com/tournymasterbot/io_helpers"
)

var testTextPath = "C:/temp/test.dat"
var testImageSourcePath = "C:/temp/bowser-source.png"
var testImageTargetPath = "C:/temp/bowser-target.png"

func main() {
	// Write a test string
	var payload = "This is a test string"
	_, textErr := iohelper.WriteAllText(testTextPath, payload)
	if textErr != nil {
		fmt.Println(textErr)
		return
	}

	// Read the source bytes
	testImageBytes, imageSourceErr := iohelper.ReadAllBytes(testImageSourcePath)
	if imageSourceErr != nil {
		fmt.Println(imageSourceErr)
		return
	}

	// Write the source bytes to the target location
	_, imageTargetErr := iohelper.WriteAllBytes(testImageTargetPath, testImageBytes)
	if imageTargetErr != nil {
		fmt.Println(imageTargetErr)
		return
	}
}
