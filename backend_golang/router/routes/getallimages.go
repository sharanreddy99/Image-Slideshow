package routes

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"router/sqldb"
	"strings"
)

//GetAllImagesHandler POST API
func GetAllImagesHandler(res http.ResponseWriter,req *http.Request,){
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers","Origin, X-Requested-With, Content-Type, Accept")

	response := make(map[string]string)

	err := req.ParseMultipartForm(0)
	if err!=nil {
		log.Fatal(err.Error())
	}

	
	email  := req.Form.Get("email");
	rows,err := sqldb.DB.Query("select filename from sureify.imagepaths where email = '"+email+"';")
	
	var imagesarray  = make([]string,0);
	var extensionsarray = make([]string,0);

 	for rows.Next() {
		filename := ""
		rows.Scan(&filename)
		filename = "./gallery/"+filename;
		extension := strings.ToLower(strings.Split(filename,".")[1])
		data, _ := ioutil.ReadFile(filename)
		encoded := base64.StdEncoding.EncodeToString(data)

		imagesarray = append(imagesarray,strings.Replace(encoded,"/","\\/",-1))
		extensionsarray = append(extensionsarray,extension)
	}

	if(len(imagesarray)==0){		
		response["imagesarray"] = "["+strings.Join(imagesarray,",")+"]"
		response["extensionsarray"] = "["+strings.Join(extensionsarray,",")+"]"
	}else{
		response["imagesarray"] = "[\""+strings.Join(imagesarray,`","`)+"\"]"
		response["extensionsarray"] = "[\""+strings.Join(extensionsarray,`","`)+"\"]"
	}
	json.NewEncoder(res).Encode(response);
}
