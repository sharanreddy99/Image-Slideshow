package sqldb

import (
	"backend_golang/constants"
	"database/sql"
	"fmt"
	"log"
)

// DB is a global variable to hold db connection
var DB *sql.DB

//CreateConnection creates a database connection and creates all the tables and has a global DB pointer
func CreateConnection() bool {

	stmt := constants.MYSQL_USERNAME + ":" + constants.MYSQL_PASSWORD + "@tcp(" + constants.MYSQL_HOST + ":" + constants.MYSQL_PORT + ")/"
	db, err := sql.Open("mysql", stmt)
	if err != nil {
		log.Println(err.Error())
		DB = nil
		return false
	}

	stmt = "create database if not exists " + constants.MYSQL_DATABASE
	_, err = db.Exec(stmt)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	stmt = "use " + constants.MYSQL_DATABASE
	if _, err = db.Exec(stmt); err != nil {
		log.Println(err.Error())
		return false
	}

	stmt = fmt.Sprintf("create table if not exists %v.user(id int(10) auto_increment primary key,firstname varchar(50),lastname varchar(50),email varchar(50),mobile varchar(10),password varchar(50),token tinytext default NULL);", constants.MYSQL_DATABASE)
	if _, err = db.Query(stmt); err != nil {
		log.Println(err.Error())
		return false
	}

	stmt = fmt.Sprintf("create table if not exists %v.images(fileno int(10) auto_increment primary key,filename varchar(50),imagedata mediumtext,email varchar(50));", constants.MYSQL_DATABASE)
	if _, err = db.Query(stmt); err != nil {
		log.Println(err.Error())
		return false
	}

	DB = db
	return true
}
