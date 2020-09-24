module github.com/tournymasterbot/ddl_query

replace github.com/tournymasterbot/database_helpers => ../database_helpers // 'replace' should be changed to a 'require (...)' block

go 1.15

require github.com/tournymasterbot/database_helpers v0.0.0-00010101000000-000000000000
