package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"router/sqldb"
	"strconv"
	"strings"
)

//ImageUploadHandler POST API
func ImageUploadHandler(res http.ResponseWriter,req *http.Request,){
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers","Origin, X-Requested-With, Content-Type, Accept")

	response := make(map[string]string)

	err := req.ParseMultipartForm(0)
	if err!=nil {
		log.Fatal(err.Error())
	}
	
	stmt := "select max(fileno) fileno from sureify.imagepaths";
	rows, err := sqldb.DB.Query(stmt);
	if err!=nil{
		panic(err.Error())
	}
	
	rows.Next()
	maxfileno := 0;
	rows.Scan(&maxfileno)
	maxfileno++

	file, handler, err := req.FormFile("imagefile")
	defer file.Close()

	email := req.Form.Get("email")
	filename := handler.Filename
	extension := strings.ToLower(strings.Split(filename,".")[1])

	if(extension != "jpg" && extension != "png" && extension != "jpeg") {
		response["status"] = "failure";
        response["ModalTitle"]="Invalid File Format....";
        response["ModalBody"]="Please upload files of format jpg, jpeg and png...";
	}else{

		filename = "image"+strconv.Itoa(maxfileno)+"."+extension
		tempFile, err := os.Create("./gallery/"+filename)
		if err != nil {
			fmt.Println(err)
		}
		
		defer tempFile.Close()
		_, err = io.Copy(tempFile, file)
		if err != nil {
				fmt.Fprintln(res, err)
		}

		stmt := "insert into sureify.imagepaths values(null,'"+filename+"','"+email+"');";
		_, err = sqldb.DB.Query(stmt);
		if err!=nil{
			panic(err.Error())
		}
		
        response["status"]="success";
        response["filename"]= filename;
	}
	json.NewEncoder(res).Encode(response);
}
