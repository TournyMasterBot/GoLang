package main

import (
	"fmt"

	iohelper "github.com/tournymasterbot/io_helpers"
)

// Uncompressed text written to a file
// Then compress the file
// Then inflate the file
// And verify the results
var basepath1 string = "C:/temp/test1_base.txt"
var gzwritepath1 string = "C:/temp/test1.txt.gz"
var inflatewritepath1 string = "C:/temp/inflate_test1.txt"

// Uncompressed text written to a compressed byte stream
// Then compressed stream written to a file
// Then compressed file is inflated
// And verify the results
var gzwritepath2 string = "C:/temp/test2.txt.gz"
var inflatewritepath2 string = "C:/temp/inflate_test2.txt"

func main() {
	// Test string that starts it all
	var testString = "This is a test string"

	// ===== TEST 1 =====
	// Write the base file
	success, writeErr := iohelper.WriteAllText(basepath1, testString)
	if writeErr != nil {
		fmt.Println(writeErr)
	}

	// Compress the base file
	if success == true {
		fmt.Println("Compressing base file")
		compressErr := iohelper.CompressFile(basepath1, gzwritepath1)
		if compressErr != nil {
			fmt.Println(compressErr)
		}
	} else {
		fmt.Println("Failed to write test1 base file")
		return
	}

	// Create the inflated file
	inflateErr := iohelper.InflateFile(gzwritepath1, inflatewritepath1)
	if inflateErr != nil {
		fmt.Println(inflateErr)
		return
	}

	// Read the inflated file text
	inflateTest1Result, inflateReadResultErr := iohelper.ReadAllText(inflatewritepath1)
	if inflateReadResultErr != nil {
		fmt.Println(inflateReadResultErr)
		return
	}

	// Validate the result
	if testString != inflateTest1Result {
		fmt.Println("Test1 Failed")
	} else {
		fmt.Println("Test1 Success")
	}
	// ===== END TEST 1 =====

	// ===== TEST 2 =====
	// Create the compressed bytes
	compressedTextBytes, compressedTextErr := iohelper.CompressText(testString)
	if compressedTextErr != nil {
		fmt.Println(compressedTextErr)
		return
	}

	// Write the compressed bytes to a file
	compressedTextSuccess, compressedTextWriteErr := iohelper.WriteAllBytes(gzwritepath2, compressedTextBytes)
	if compressedTextWriteErr != nil {
		fmt.Println(compressedTextWriteErr)
		return
	}

	if compressedTextSuccess == true {
		gzwritefileErr := iohelper.InflateFile(gzwritepath2, inflatewritepath2)
		if gzwritefileErr != nil {
			fmt.Println(gzwritefileErr)
			return
		}
	} else {
		fmt.Println("Failed to write test2 gzip file")
		return
	}

	// Re-inflate the compressed text and check if we got the same text we're expecting
	inflatedText, inflatedTextErr := iohelper.InflateText(compressedTextBytes)
	if inflatedTextErr != nil {
		fmt.Println(inflatedTextErr)
		return
	}

	if inflatedText == testString {
		fmt.Println("Test2 Success")
	} else {
		fmt.Println("Test2 Fail")
	}

	// ===== END TEST 2 =====
}
