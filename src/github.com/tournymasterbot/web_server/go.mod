module github.com/tournymasterbot/web_server

replace github.com/tournymasterbot/web_server_helpers => ../web_server_helpers // 'replace' should be changed to a 'require (...)' block

go 1.15

require (
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/tournymasterbot/web_server_helpers v0.0.0-00010101000000-000000000000
)
