package sqldb

import (
	"database/sql"
)

// DB is a global variable to hold db connection
var DB *sql.DB

const (
	user string = "root"
	pass string = "mysqlpassword"
	port string = "3306"
	host string = "172.17.0.2"
)

//CreateConnection creates a database connection and creates all the tables and has a global DB pointer
func CreateConnection() bool {

	stmt := user + ":" + pass + "@tcp("+host+":" + port + ")/"
	db, err := sql.Open("mysql", stmt)
	if err != nil {
		DB = nil
		return false
	}

	stmt = "create database if not exists sureify"
	sqlQuery, err := db.Query(stmt)
	if err != nil {
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
