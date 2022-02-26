package routes

import (
	"backend_golang/constants"
	"backend_golang/router/sqldb"
	"backend_golang/utils"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

//ImageUploadHandler POST API
func ImageUploadHandler(res http.ResponseWriter, req *http.Request) {
	utils.SetHeaders(&res)
	response := make(map[string]string)

	if sqldb.DB == nil {
		response["ModalTitle"] = "Service Unavailable..."
		response["ModalBody"] = "Signup Service is unavailable right now... Please try again later"
		res.WriteHeader(http.StatusServiceUnavailable)
		panic("503 Error")
	}

	defer func() {
		if r := recover(); r != nil {
			_ = json.NewEncoder(res).Encode(response)
		}
	}()

	_ = req.ParseMultipartForm(0)
	token := req.Form.Get("token")
	email := utils.IsValidUser(token)

	if email == "" {
		response["ModalTitle"] = "Not Authorized..."
		response["ModalBody"] = "You are not authorized..."
		res.WriteHeader(http.StatusUnauthorized)
		panic("401 Error")
	}

	stmt := fmt.Sprintf("select max(fileno) fileno from %s.images", constants.MYSQL_DATABASE)
	rows, _ := sqldb.DB.Query(stmt)

	rows.Next()
	maxfileno := 0
	_ = rows.Scan(&maxfileno)
	maxfileno++

	file, handler, _ := req.FormFile("imagefile")
	defer file.Close()

	filename := handler.Filename
	extension := strings.ToLower(strings.Split(filename, ".")[1])

	if extension != "jpg" && extension != "png" && extension != "jpeg" {
		response["status"] = "failure"
		response["ModalTitle"] = "Invalid File Format...."
		response["ModalBody"] = "Please upload files of format jpg, jpeg and png..."
		res.WriteHeader(http.StatusForbidden)
		panic("403 Error")
	} else {

		filename = "image" + strconv.Itoa(maxfileno) + "." + extension
		data, _ := ioutil.ReadAll(file)
		encoded := base64.StdEncoding.EncodeToString(data)

		stmt := fmt.Sprintf("insert into %s.images values(null,'%s','%s','%s');", constants.MYSQL_DATABASE, filename, encoded, email)
		_, err := sqldb.DB.Query(stmt)
		if err != nil {

			response["ModalTitle"] = "Image Uploading Failed...."
			response["ModalBody"] = "Image Uploading failed.try again later..."
			res.WriteHeader(http.StatusForbidden)
			panic("403 Error")
		}

		response["status"] = "success"
		response["filename"] = filename
		res.WriteHeader(http.StatusOK)
	}
	_ = json.NewEncoder(res).Encode(response)
}
