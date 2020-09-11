package main

import (
	webserverhelper "github.com/tournymasterbot/web_server_helpers"
)

func main() {
	var serverConfig = webserverhelper.ServerConfig{
		IP:   "0.0.0.0",
		Port: 8080,
	}
	webserverhelper.Start(serverConfig)
}
