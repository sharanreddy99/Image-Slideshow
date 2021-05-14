package routes

import (
	"backendgo/router/sqldb"
	"encoding/json"
	"net/http"
)

//UpdateUserHandler POST API
func UpdateUserHandler(res http.ResponseWriter,req *http.Request,){
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
	
	firstname := req.Form.Get("firstname");
    lastname := req.Form.Get("lastname");
    mobile := req.Form.Get("mobile");
    password := req.Form.Get("password");
    newemail := req.Form.Get("email");

	stmt := "select * from sureify.user where email='"+newemail+"'";
	rows, _ := sqldb.DB.Query(stmt);
	
	userExists := rows.Next()

	if(userExists){
		
		res.WriteHeader(http.StatusForbidden)
        response["ModalTitle"]="Email already registered...";
        response["ModalBody"]="A User has already registered with the specified Email..Please try again with another Email...";
        panic("403 Error");
	}

	stmt = "update sureify.user set firstname = '"+firstname+"',lastname='"+lastname+"',email='"+newemail+"',password='"+password+"',mobile='"+mobile+"' where email = '"+email+"';";
	_,err := sqldb.DB.Query(stmt);
	

	if(err==nil){

      stmt = "update sureify.images set email='"+newemail+"' where email='"+email+"'";
	  sqldb.DB.Query(stmt);
	
	  res.WriteHeader(http.StatusOK)
      response["ModalTitle"]="User Updated...";
      response["ModalBody"]="User details have been updated successfully.Please login again...";
        
    }else{
      res.WriteHeader(http.StatusForbidden)
      response["ModalTitle"]="User Updation Failed...";
      response["ModalBody"]="Updating user has failed..Please try again...";
      panic("403 Error");
    }
	json.NewEncoder(res).Encode(response);
}
