package sqldb

import (
	"database/sql"
	"log"
)

// DB is a global variable to hold db connection
var DB *sql.DB

const (
	user = "root"
	pass = ""
	port = "3306"
)

//CreateConnection creates a database connection and creates all the tables and has a global DB pointer
func CreateConnection(){
	stmt := user+":"+pass+"@tcp(127.0.0.1:"+port+")/"
	db, err := sql.Open("mysql",stmt)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt = "create database if not exists sureify"
	sqlQuery, err := db.Query(stmt);
	if err!=nil{
		panic(err.Error())
	}

	stmt = "use sureify"
	sqlQuery, err = db.Query(stmt);
	if err!=nil{
		panic(err.Error())
	}

	stmt = "create table if not exists sureify.user(id int(10) auto_increment primary key,firstname varchar(50),lastname varchar(50),email varchar(50),mobile varchar(10),password varchar(50));"
	sqlQuery, err = db.Query(stmt)
	if err!=nil{
		panic(err.Error())
	}

	stmt = "create table if not exists sureify.imagepaths(fileno int(10) auto_increment primary key,filename varchar(50),email varchar(50));"
	sqlQuery, err = db.Query(stmt)
	if err!=nil{
		panic(err.Error())
	}
	defer sqlQuery.Close()

	DB = db
}
