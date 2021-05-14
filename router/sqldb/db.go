package sqldb

import (
	"database/sql"
	"fmt"
	"os"
)

// DB is a global variable to hold db connection
var DB *sql.DB


//CreateConnection creates a database connection and creates all the tables and has a global DB pointer
func CreateConnection() bool {

	var (
		user string = os.Getenv("DB_USERNAME_LOCAL")
		pass string = os.Getenv("DB_PASSWORD_LOCAL")
		port string = os.Getenv("DB_PORT_LOCAL")
		host string = os.Getenv("DB_HOST_LOCAL")
	)

	stmt := user + ":" + pass + "@tcp("+host+":" + port + ")/"
	db, err := sql.Open("mysql", stmt)
	if err != nil {
		DB = nil
		return false
	}

	stmt = "create database if not exists sureify"
	sqlQuery, err := db.Query(stmt)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	stmt = "use sureify"
	sqlQuery, err = db.Query(stmt)
	if err != nil {
		return false;
	}

	stmt = "create table if not exists sureify.user(id int(10) auto_increment primary key,firstname varchar(50),lastname varchar(50),email varchar(50),mobile varchar(10),password varchar(50),token tinytext default NULL);"
	sqlQuery, err = db.Query(stmt)
	if err != nil {
		return false;
	}

	stmt = "create table if not exists sureify.images(fileno int(10) auto_increment primary key,filename varchar(50),imagedata mediumtext,email varchar(50));"
	sqlQuery, err = db.Query(stmt)
	if err != nil {
		return false;
	}

	defer sqlQuery.Close()
	DB = db
	return true;
}
