package databasehelper

// Assumptions:
// Execute: go get github.com/go-sql-driver/mysql

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectionData :
type ConnectionData struct {
	ServerType string
	UserName   string
	Password   string
	HostName   string
	HostPort   int
	DBName     string
}

var serverType string
var activeDatabaseConfiguration ConnectionData
var db sql.DB

// SetConnectionData : Sets the connection configuration data
func SetConnectionData(config ConnectionData) {
	activeDatabaseConfiguration = config
}

// GetDsn : Returns the active configuration DSN
func GetDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s",
		activeDatabaseConfiguration.UserName,
		activeDatabaseConfiguration.Password,
		activeDatabaseConfiguration.HostName+":"+strconv.Itoa(activeDatabaseConfiguration.HostPort),
		activeDatabaseConfiguration.DBName,
	)
}

//Initialize : Opens a sql connection, returns an error if there was an issue opening the connection
func Initialize(serverType string) (*sql.DB, error) {
	activeDatabaseConfiguration.ServerType = serverType
	var connString = GetDsn()
	db, err := sql.Open("mysql", connString)
	if err != nil {
		return db, err
	}
	return db, nil
}
