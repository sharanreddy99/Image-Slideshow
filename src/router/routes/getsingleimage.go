package routes

import (
	"encoding/json"
	"net/http"
	"router/sqldb"
)

//GetSingleImageHandler POST API
func GetSingleImageHandler(res http.ResponseWriter,req *http.Request,){
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
	email  := req.Form.Get("email");
	token := req.Form.Get("token");


	if(!IsValidUser(email,token)){
        response["ModalTitle"]="Not Authorized...";
        response["ModalBody"]="You are not authorized...";    
        res.WriteHeader(http.StatusUnauthorized)
		panic("401 Error");
	}


	filename := req.Form.Get("filename")

	stmt := "select imagedata from sureify.images where email = '"+email+"' and filename='"+filename+"';";
	rows,_ := sqldb.DB.Query(stmt)
	if(rows.Next()){
		imagedata := ""
		rows.Scan(&imagedata)
		response["imagedata"] = imagedata;
		response["fromgo"] = "true"
	}
	json.NewEncoder(res).Encode(response)
}
