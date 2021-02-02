package routes

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"router/sqldb"
	"strconv"
	"strings"
)

//ImageUploadHandler POST API
func ImageUploadHandler(res http.ResponseWriter,req *http.Request,){
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

	
	stmt := "select max(fileno) fileno from sureify.images";
	rows, _ := sqldb.DB.Query(stmt);
	
	rows.Next()
	maxfileno := 0;
	rows.Scan(&maxfileno)
	maxfileno++

	file, handler, _ := req.FormFile("imagefile")
	defer file.Close()

	filename := handler.Filename
	extension := strings.ToLower(strings.Split(filename,".")[1])

	if(extension != "jpg" && extension != "png" && extension != "jpeg") {
		response["status"] = "failure";
        response["ModalTitle"]="Invalid File Format....";
		response["ModalBody"]="Please upload files of format jpg, jpeg and png...";
		res.WriteHeader(http.StatusForbidden)
		panic("403 Error")
	}else{

		filename = "image"+strconv.Itoa(maxfileno)+"."+extension
		data, _ := ioutil.ReadAll(file)
		encoded := base64.StdEncoding.EncodeToString(data)

		stmt := "insert into sureify.images values(null,'"+filename+"','"+encoded+"','"+email+"');";
		_, err := sqldb.DB.Query(stmt);
		if err!=nil{
			
			response["ModalTitle"]="Image Uploading Failed....";
			response["ModalBody"]="Image Uploading failed.try again later...";
			res.WriteHeader(http.StatusForbidden)
			panic("403 Error")
		}
		
        response["status"]="success";
		response["filename"]= filename;
		res.WriteHeader(http.StatusOK)
	}
	json.NewEncoder(res).Encode(response);
}
