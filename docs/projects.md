# Overview
A demo project is one that you would navigate to with the command line and execute using `go run main.go`. These projects are to try and show the business user usage.

A helper project is one that you would look at the code for to look at specific implementation examples. These helpers try to show the underlying data handling.

# Demo Projects
## api_explorer
Shows an extremely basic http GET and http POST request

## ddl_query
Shows a basic set of sql DDL examples, and some simple insert/query examples

## email_templates
Shows basic usage of the Go templating language and potential uses. Worth noting: Officially this is html templating

## go_installation_test
Runs a simple test to see if Go has been properly installed

## io_explorer
A simple file test to read and write files. Also shows basic file permission editing

## simple_query
Executes an extremely basic SQL query to test connecting to a database and running a select query

## web_server
Hosts a simple web server that supports basic slug usage

# Helper Projects
## database_helpers
Connects to a database, opens connections and allows querying the database

## io_helpers
Simple file system interactions to read files, also supports gzip compression and inflation

## web_request_helpers
Simple HTTP client to support GET and POST data

## web_server_helpers
Simple HTTP server that listens for traffic and handles routing using slugs