package routes

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//GetSingleImageHandler POST API
func GetSingleImageHandler(res http.ResponseWriter,req *http.Request,){
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers","Origin, X-Requested-With, Content-Type, Accept")

	response := make(map[string]string)

	err := req.ParseMultipartForm(0)
	if err!=nil {
		log.Fatal(err.Error())
	}
	
	filename := "./gallery/"+req.Form.Get("filename");
	data, err := ioutil.ReadFile(filename)
	encoded := base64.StdEncoding.EncodeToString(data)
	response["imagedata"] = encoded
	response["fromgo"] = "true"
	json.NewEncoder(res).Encode(response);
	
}
