package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	iohelper "github.com/tournymasterbot/io_helpers"
)

// TokenData : Raw token configuration details
type TokenData struct {
	Tokens Tokens `json:"tokens"`
}

// Tokens : Token details, primarily: Token substitution name and substitution value
type Tokens []struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")
	// Simple readline loop to get user input and do something with it.
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		switch text {
		case "help":
			fmt.Println("Available Commands: 'templates', 'tokens'")
		case "templates":
			// Read the template file
			testTemplate, err := iohelper.ReadAllText("./templates/test.template")
			// Output the template file, or error
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(testTemplate)
		case "tokens":
			rawTokenData, err := iohelper.ReadAllBytes("./templates/test.tokens")
			if err != nil {
				fmt.Println(err)
				return
			}

			var data = TokenData{}
			if err := json.Unmarshal(rawTokenData, &data); err != nil {
				fmt.Println(err)
				return
			}
			for i := 0; i < len(data.Tokens); i++ {
				fmt.Println("Key: ", data.Tokens[i].Key, ", Value: ", data.Tokens[i].Value)
				fmt.Println("---")
			}
		}
		// Debug
		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		}
	}
}
