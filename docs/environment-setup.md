# Assumptions
1. You are using Windows OS
2. You are using Visual Studio Code as your IDE, and it is installed at the default path
3. Source tree is installed to the default path
4. Git is installed to the default path

# Tool Resources
1. Visual Studio Code: https://code.visualstudio.com/
2. Git: https://git-scm.com/downloads
3. GoLang: https://golang.org/
4. Source Tree: https://www.sourcetreeapp.com/
5. GoPls: https://github.com/golang/vscode-go/blob/main/docs/gopls.md
6. GoLang Installation Test: https://golang.org/doc/install#testing
7. Notepad++: https://notepad-plus-plus.org/downloads/

# Source Resources
1. Code Repository: https://github.com/TournyMasterBot/GoLang

# Web References
1. GoLang: https://golang.org/doc/

# Setup 
1. Create your GoLang source file path, for my projects I will be using `C:\Projects\GoLang\src\github.com\tournymasterbot` - if you decide to change this path, you will need to update the appropriate `go.mod` and `go imports` 
2. Download GoLang - I used the recommended defaults to install to `C:\Go\`
3. Set your GoPath variable to be the 'src' portion of the path from step 1, so example: `C:\Projects\GoLang` (To access this, open your system properties and click 'Environment Variables', change 'GOPATH')
3. Install `Go` from the extensions marketplace
4. Install `GoPls` and configure `settings.json`
* Hit `ctrl+,` and open the extension settings 
* Click on `Go`
* Scroll down to `Edit in settings.json` (it's a small hyperlink)
* `settings.json` content, this will prompt to install gopls tools, accept
```javascript
{
    "mssql.intelliSense.lowerCaseSuggestions": true,
    "git.ignoreMissingGitWarning": true,
    "terminal.integrated.shell.windows": "C:\\WINDOWS\\System32\\cmd.exe",
    "terminal.integrated.rendererType": "dom",
    "workbench.startupEditor": "newUntitledFile",
    "editor.minimap.enabled": true,
    "go.delveConfig": {
    
        "dlvLoadConfig": {
            "followPointers": true,
            "maxVariableRecurse": 1,
            "maxStringLen": 64,
            "maxArrayValues": 64,
            "maxStructFields": -1
        },
        "apiVersion": 2,
        "showGlobalVariables": true
    },
    "go.useLanguageServer": true,
    "[go]": {
        "editor.formatOnSave": true,
        "editor.codeActionsOnSave": {
            "source.organizeImports": true,
        },
        // Optional: Disable snippets, as they conflict with completion ranking.
        "editor.snippetSuggestions": "none",
    },
    "[go.mod]": {
        "editor.formatOnSave": true,
        "editor.codeActionsOnSave": {
            "source.organizeImports": true,
        },
    },
    "gopls": {
        // Add parameter placeholders when completing a function.
        "usePlaceholders": true,
 
        // If true, enable additional analyses with staticcheck.
        // Warning: This will significantly increase memory usage.
        "staticcheck": false,
    },
    "typescript.npm": ""
}
```
4. Click `View -> Terminal` in the VSCode menu
5. Follow the steps to test your installation. (Project: `src/github.com/tournymasterbot/go_installation_test/main.go`) - Install any tools it prompts to install
6. In the bottom right status bar click 'Analysis tools missing' and install the analysis tools