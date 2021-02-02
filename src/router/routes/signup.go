package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"router/sqldb"

	//package for mysql
	_ "github.com/go-sql-driver/mysql"
)

//SignupHandler POST API
func SignupHandler(res http.ResponseWriter,req *http.Request){

	SetHeaders(&res)
	response := make(map[string]string)

	if(sqldb.DB==nil){
		response["ModalTitle"]="Service Unavailable...";
		response["ModalBody"]="Signup Service is unavailable right now... Please try again later";
		res.WriteHeader(http.StatusServiceUnavailable)
		panic("503 Error")
	}

	defer func(){
		if r := recover(); r!=nil{
			json.NewEncoder(res).Encode(response);
		}
	}()

	req.ParseMultipartForm(0)

	firstname := req.Form.Get("firstname")
	lastname := req.Form.Get("lastname")
	email := req.Form.Get("email")
	mobile := req.Form.Get("mobile")
	password := req.Form.Get("password")

	stmt := "select * from sureify.user where email = '"+email+"';"
	rows,_ := sqldb.DB.Query(stmt)
	defer rows.Close()

	userExists := rows.Next()
	if(userExists){
		response["ModalTitle"] = "Email already registered..."
		response["ModalBody"] = "A User has already registered with the specified Email..Please try again with another Email..."
		res.WriteHeader(http.StatusForbidden)
		panic("403 Error")

	}else{
		stmt = "insert into sureify.user values(NULL,'"+firstname+"','"+lastname+"','"+email+"','"+mobile+"','"+password+"',NULL);";
		_, err := sqldb.DB.Query(stmt);
		if err!=nil{
			response["ModalTitle"]="Registration Failed..."
			response["ModalBody"]="Error occured during registration..Please try again later..."    
			res.WriteHeader(http.StatusForbidden)
			fmt.Println(err.Error())
			panic("403 Error")
		}else{
			response["ModalTitle"]="Registration Successful..."
			response["ModalBody"]="Registration has been completed successfully..Please Login..."                
			res.WriteHeader(http.StatusCreated)	
		}
	}
	json.NewEncoder(res).Encode(response);
}
