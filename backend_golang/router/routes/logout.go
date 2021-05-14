package routes

import (
	"backendgo/router/sqldb"
	"encoding/json"
	"net/http"
)

//LogoutHandler POST API
func LogoutHandler(res http.ResponseWriter,req *http.Request,){
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
	email := req.Form.Get("email")


    stmt := "update sureify.user set token=null where email='"+email+"'";
    sqldb.DB.Query(stmt);
	json.NewEncoder(res).Encode("Successfully logged out");
}
