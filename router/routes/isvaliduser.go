package routes

import (
	"backendgo/router/sqldb"
	"encoding/json"
	"net/http"
)

//IsValidUserHandler POST API
func IsValidUserHandler(res http.ResponseWriter,req *http.Request,){
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

    res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(response);
}
