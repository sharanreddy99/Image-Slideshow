package routes

import (
	"backendgo/router/sqldb"
	"encoding/json"
	"net/http"
)

//FetchUserHandler POST API
func FetchUserHandler(res http.ResponseWriter,req *http.Request,){
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
	token := req.Form.Get("token");
	email := IsValidUser(token)

	if(email==""){
        response["ModalTitle"]="Not Authorized...";
        response["ModalBody"]="You are not authorized...";    
        res.WriteHeader(http.StatusUnauthorized)
		panic("401 Error");
	}
	

	stmt := "select firstname,lastname,mobile,password from sureify.user where email='"+email+"'";
	rows, _ := sqldb.DB.Query(stmt);
	
	userExists := rows.Next()

	if(userExists){
		firstname:=""
		lastname:=""
		mobile:=""
		password:=""
		rows.Scan(&firstname,&lastname,&mobile,&password)
		
		response["firstname"] = firstname;
		response["lastname"] = lastname;
		response["mobile"] = mobile;
		response["password"] = password;
		res.WriteHeader(http.StatusOK)
	}
	json.NewEncoder(res).Encode(response);
}
