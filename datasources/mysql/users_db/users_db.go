package users_db

import (
	"database/sql"
	"fmt"
	"github.com/adershrp/bookstore_users-api/logger"
	"os"
	/**
	  Just loading the driver.
	*/
	_ "github.com/go-sql-driver/mysql"
)

const (
	/**
	  help to get values from environment variables
	*/
	mysqlUsersUsername = "mysql_users_username"
	mysqlUsersPassword = "mysql_users_password"
	mysqlUsersHost     = "mysql_users_host"
	mysqlUsersSchema   = "mysql_users_schema"
)

var (
	Client   *sql.DB
	username = os.Getenv(mysqlUsersUsername)
	password = os.Getenv(mysqlUsersPassword)
	host     = os.Getenv(mysqlUsersHost)
	schema   = os.Getenv(mysqlUsersSchema)
)

func init() {
	/**
	  constructing the dataSourceName
	*/
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		logger.Error("Error while connecting to Database.", err)
		panic(err)
	}
	/**
	  pinging the database to verify the connection
	*/
	err = Client.Ping()
	if err != nil {
		logger.Error("Error while connecting to Database.", err)
		panic(err)
	}
	logger.Info("Database successfully connected.")
}
