package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"router/sqldb"

	//package for mysql
	_ "github.com/go-sql-driver/mysql"
)

//SignupHandler POST API
func SignupHandler(res http.ResponseWriter,req *http.Request){

	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers","Origin, X-Requested-With, Content-Type, Accept")
	
	response := make(map[string]string)

	err := req.ParseMultipartForm(0)
	if err!=nil {
		log.Fatal(err.Error())
	}

	firstname := req.Form.Get("firstname")
	lastname := req.Form.Get("lastname")
	email := req.Form.Get("email")
	mobile := req.Form.Get("mobile")
	password := req.Form.Get("password")

	stmt := "select * from sureify.user where email = '"+email+"';"
	rows, err := sqldb.DB.Query(stmt)
	if err!=nil{
		panic(err.Error())
	}
	defer rows.Close()

	userExists := rows.Next()
	if(userExists){
		response["ModalTitle"] = "Email already registered..."
		response["ModalBody"] = "A User has already registered with the specified Email..Please try again with another Email..."
		res.WriteHeader(http.StatusOK)
	}else{
		stmt = "insert into sureify.user values(NULL,'"+firstname+"','"+lastname+"','"+email+"','"+mobile+"','"+password+"');";
		_, err := sqldb.DB.Query(stmt);
		if err!=nil{
			response["ModalTitle"]="Registration Failed..."
			response["ModalBody"]="Error occured during registration..Please try again later..."    
			res.WriteHeader(http.StatusOK)
		}else{
			response["ModalTitle"]="Registration Successful..."
			response["ModalBody"]="Registration has been completed successfully..Please Login..."                
			res.WriteHeader(http.StatusCreated)	
		}
	}
	json.NewEncoder(res).Encode(response);
}
