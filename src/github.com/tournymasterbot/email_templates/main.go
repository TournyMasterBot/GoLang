package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	iohelper "github.com/tournymasterbot/io_helpers"
)

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
			testTemplate, err := iohelper.ReadAllText("./templates/test.template")
			if err == nil {
				fmt.Println(testTemplate)
			} else {
				fmt.Println(err)
			}
		case "tokens":
			fmt.Println("token placeholder")
		}
		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		}
	}
}
