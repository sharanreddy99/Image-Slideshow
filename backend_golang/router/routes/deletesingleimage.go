package routes

import (
	"backend_golang/constants"
	"backend_golang/router/sqldb"
	"backend_golang/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

//DeleteSingleImageHandler POST API
func DeleteSingleImageHandler(res http.ResponseWriter, req *http.Request) {
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

	stmt := fmt.Sprintf("delete from %s.images where email = '%s' and filename='%s';", constants.MYSQL_DATABASE, email, filename)
	_, err := sqldb.DB.Query(stmt)

	if err == nil {
		response["ModalTitle"] = "Operation Successful..."
		response["ModalBody"] = "image (" + filename + ") has been deleted successfully..."
		res.WriteHeader(http.StatusOK)
	} else {
		response["ModalTitle"] = "Operation Failed..."
		response["ModalBody"] = "Image can't be deleted..."
		res.WriteHeader(http.StatusForbidden)
		_ = json.NewEncoder(res).Encode(response)
		return
	}

	_ = json.NewEncoder(res).Encode(response)
}
