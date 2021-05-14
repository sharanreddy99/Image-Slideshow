package routes

import (
	"backendgo/router/sqldb"
	"encoding/json"
	"net/http"
	"strings"
)

//GetAllImagesHandler POST API
func GetAllImagesHandler(res http.ResponseWriter,req *http.Request,){
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

	rows,_ := sqldb.DB.Query("select filename,imagedata from sureify.images where email = '"+email+"';")
	
	var imagesarray  = make([]string,0);
	var extensionsarray = make([]string,0);
	var filenamesarray = make([]string,0);

 	for rows.Next() {
		filename := ""
		imagedata := ""
		
		rows.Scan(&filename,&imagedata)
		
		extension := strings.ToLower(strings.Split(filename,".")[1])

		imagesarray = append(imagesarray,strings.Replace(imagedata,"/","\\/",-1))
		extensionsarray = append(extensionsarray,extension)
		filenamesarray = append(filenamesarray,filename)
	}

	if(len(imagesarray)==0){		
		response["imagesarray"] = "["+strings.Join(imagesarray,",")+"]"
		response["extensionsarray"] = "["+strings.Join(extensionsarray,",")+"]"
		response["filenamesarray"] = "["+strings.Join(extensionsarray,",")+"]"
		
	}else{
		response["imagesarray"] = "[\""+strings.Join(imagesarray,`","`)+"\"]"
		response["extensionsarray"] = "[\""+strings.Join(extensionsarray,`","`)+"\"]"
		response["filenamesarray"] = "[\""+strings.Join(filenamesarray,`","`)+"\"]"
	}
	json.NewEncoder(res).Encode(response);
}
