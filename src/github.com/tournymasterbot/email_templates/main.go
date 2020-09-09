package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
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
	// Sanity checks to make sure we only get one arg which should be a template name
	if os.Args == nil || len(os.Args) <= 1 || len(os.Args) > 2 {
		fmt.Println("You must pass one argument: The template name without the file extension.")
		return
	}

	// Process template name from args
	var templateName = os.Args[1]
	templateName = strings.TrimSpace(templateName)
	if len(templateName) == 0 {
		fmt.Println("The template name may not be blank")
		return
	}

	// Read the template file
	templateText, err := iohelper.ReadAllText("./templates/" + templateName + ".template")
	// Output the template file, or error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Base Template:")
	fmt.Println(templateText)

	// Load tokens
	rawTokenData, err := iohelper.ReadAllBytes("./templates/" + templateName + ".tokens")
	if err == nil {
		var data = TokenData{}
		// Process tokens
		if err := json.Unmarshal(rawTokenData, &data); err == nil {
			t := template.New("everything")
			template.Must(t.New("/").Parse(templateText))
			fmt.Println("Tokens Found:")
			// Add all tokens for the template
			for i := 0; i < len(data.Tokens); i++ {
				template.Must(t.New(data.Tokens[i].Key).Parse(data.Tokens[i].Value))
				fmt.Println("Key: ", data.Tokens[i].Key, ", Value: ", data.Tokens[i].Value)
			}
			fmt.Println("***")
			// Output the template to the command line
			if err := t.ExecuteTemplate(os.Stdout, "/", templateText); err != nil {
				log.Fatal("Failed to execute:", err)
			}
		}
	}
}
