package main

import (
	"database/sql"
	"fmt"
	"strconv"

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

	dropTables(conn)
	createTables(conn)
	insertRows(conn)
	queryRows(conn)
}

// dropTables : Drops the demo table if it already exists for later recreation and testing
func dropTables(db *sql.DB) {
	var sql = "drop table if exists " + dbname + ".DemoTable;"
	var rows, rowErr = db.Query(sql)
	if rowErr != nil {
		fmt.Println(rowErr)
		return
	}
	defer rows.Close()
}

// createTables : Creates the demo table for later data insert and query
func createTables(db *sql.DB) {
	var sql = `
		create table if not exists ` + dbname + `.DemoTable
		(
			id int not null auto_increment,
			first_name varchar(50) not null,
			last_name varchar(50) not null,
			primary key (id)
		) default character set utf8mb4 collate utf8mb4_general_ci;
	`
	var rows, rowErr = db.Query(sql)
	if rowErr != nil {
		fmt.Println(rowErr)
		return
	}
	defer rows.Close()
}

// insertRows : Inserts test rows into the demo table that can later be queried
func insertRows(db *sql.DB) {
	for i := 0; i < 50; i++ {
		var sql = `
			insert into ` + dbname + `.DemoTable (first_name, last_name)
			values (?, ?);
		`
		var rows, rowErr = db.Query(sql, "FName"+strconv.Itoa(i), "LName"+strconv.Itoa(i))
		if rowErr != nil {
			fmt.Println(rowErr)
			return
		}
		defer rows.Close()
	}
}

// queryRows : Runs a basic select query against the DemoTable and outputs the result
func queryRows(db *sql.DB) {
	// Set the command to execute
	var sql = `
		select id, first_name, last_name
		from ` + dbname + `.DemoTable;
	`

	// Get row results
	var rows, rowErr = db.Query(sql)
	if rowErr != nil {
		fmt.Println(rowErr)
		return
	}
	defer rows.Close()

	// Iterate over the rowset and output the demo table columns
	for rows.Next() {
		var id int
		var fname string
		var lname string
		if err := rows.Scan(&id, &fname, &lname); err != nil {
			fmt.Println(err)
		}
		fmt.Printf("id: %d, fname: %s, lname: %s\n", id, fname, lname)
	}
}
