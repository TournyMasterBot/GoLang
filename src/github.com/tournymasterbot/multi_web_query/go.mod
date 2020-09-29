module github.com/tournymasterbot/multi_web_query

replace github.com/tournymasterbot/web_request_helpers => ../web_request_helpers // 'replace' should be changed to a 'require (...)' block

go 1.15

require github.com/tournymasterbot/web_request_helpers v0.0.0-00010101000000-000000000000
