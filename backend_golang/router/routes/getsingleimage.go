package routes

import (
	"backend_golang/constants"
	"backend_golang/router/sqldb"
	"backend_golang/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

//GetSingleImageHandler POST API
func GetSingleImageHandler(res http.ResponseWriter, req *http.Request) {
	utils.SetHeaders(&res)
	response := make(map[string]string)

	if sqldb.DB == nil {
		response["ModalTitle"] = "Service Unavailable..."
		response["ModalBody"] = "Signup Service is unavailable right now... Please try again later"
		res.WriteHeader(http.StatusServiceUnavailable)
		_ = json.NewEncoder(res).Encode(response)
		return
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
		_ = json.NewEncoder(res).Encode(response)
		return
	}

	filename := req.Form.Get("filename")

	stmt := fmt.Sprintf("select imagedata from %s.images where email = '%s' and filename='%s';", constants.MYSQL_DATABASE, email, filename)
	rows, _ := sqldb.DB.Query(stmt)
	if rows.Next() {
		imagedata := ""
		_ = rows.Scan(&imagedata)
		response["imagedata"] = imagedata
		response["fromgo"] = "true"
	}
	_ = json.NewEncoder(res).Encode(response)
}
