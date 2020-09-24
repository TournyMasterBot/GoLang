package main

import (
	"fmt"

	db "github.com/tournymasterbot/database_helpers"
)

// Development test data, change to match your environment
const (
	username = "root"
	password = "L2IR94x0spPyrSQrm8oj" // This is a generated password, in your application you should store this in a secure location away from source control
	hostname = "192.168.1.2"
	hostport = 3306
	dbname   = "golang" // This database must exist
)

func main() {
	// Sets the database connection configuration data
	db.SetConnectionData(db.ConnectionData{
		UserName: username,
		Password: password,
		HostName: hostname,
		HostPort: hostport,
		DBName:   dbname,
	})

	// Uses the connection data from the set command to initialize and open the database connection
	var conn, connErr = db.Initialize("mysql")
	if connErr != nil {
		fmt.Println(connErr)
		return
	}
	defer conn.Close()

	// Creates a basic sql command for debug purposes
	var sql = "select 1 as name union all select 2 as name union all select 3 as name"
	var rows, rowErr = conn.Query(sql)
	if rowErr != nil {
		fmt.Println(rowErr)
		return
	}
	defer rows.Close()

	// Iterates over the test query
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s\n", name)
	}
}
