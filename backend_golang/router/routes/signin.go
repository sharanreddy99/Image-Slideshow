package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"router/sqldb"
)

//SigninHandler POST API
func SigninHandler(res http.ResponseWriter,req *http.Request,){
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers","Origin, X-Requested-With, Content-Type, Accept")

	response := make(map[string]string)

	err := req.ParseMultipartForm(0)
	if err!=nil {
		log.Fatal(err.Error())
	}
	
	
	email := req.Form.Get("email")
	password := req.Form.Get("password")

	stmt := "select * from sureify.user where email='"+email+"' and password='"+password+"';"
	rows, err := sqldb.DB.Query(stmt);
	if err!=nil{
		panic(err.Error())
	}
	
	userExists := rows.Next()

	if(userExists){
		response["status"] = "success"
		response["email"] = email
		response["password"] = password
		res.WriteHeader(http.StatusCreated)
	}else{
		response["status"] = "failure"
		response["ModalTitle"] = "Invalid Credentials..."
		response["ModalBody"]="The Email or Password is incorrect. Please enter valid credentials..."
		res.WriteHeader(http.StatusOK)
	}
	json.NewEncoder(res).Encode(response);
}
