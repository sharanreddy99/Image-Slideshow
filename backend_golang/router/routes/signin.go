package routes

import (
	"backendgo/router/sqldb"
	"encoding/json"
	"fmt"
	"net/http"
)

//SigninHandler POST API
func SigninHandler(res http.ResponseWriter,req *http.Request,){
	
	SetHeaders(&res)
	response := make(map[string]string)

	if(sqldb.DB==nil){
		response["ModalTitle"]="Service Unavailable...";
		response["ModalBody"]="Signup Service is unavailable right now... Please try again later";
		res.WriteHeader(http.StatusServiceUnavailable)
		fmt.Println("WTF Service not available")
		panic("503 Error")
	}

	defer func(){
		if r := recover(); r!=nil{
			json.NewEncoder(res).Encode(response);
		}
	}()

	req.ParseMultipartForm(0)
	
	email := req.Form.Get("email")
	password := req.Form.Get("password")

	stmt := "select * from sureify.user where email='"+email+"' and password='"+password+"';"
	rows, _ := sqldb.DB.Query(stmt);
	userExists := rows.Next()

	if(userExists){
		token := GenerateToken(email);
		stmt = "update sureify.user set token='"+token+"' where email='"+email+"' and password='"+password+"';"
		sqldb.DB.Query(stmt);

		response["status"] = "success"
		response["email"] = email
		response["password"] = password
		response["token"] = token
		res.WriteHeader(http.StatusOK)
	}else{
		response["status"]="failure";
		response["ModalTitle"]="Invalid Credentials...";
		response["ModalBody"]="The Email or Password is incorrect. Please enter valid credentials...";    
		res.WriteHeader(http.StatusUnauthorized)
		panic("401 Error");
	}
	json.NewEncoder(res).Encode(response);
}
