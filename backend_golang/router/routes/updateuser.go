package routes

import (
	"backend_golang/constants"
	"backend_golang/router/sqldb"
	"backend_golang/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

//UpdateUserHandler POST API
func UpdateUserHandler(res http.ResponseWriter, req *http.Request) {
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

	firstname := req.Form.Get("firstname")
	lastname := req.Form.Get("lastname")
	mobile := req.Form.Get("mobile")
	password := req.Form.Get("password")
	newemail := req.Form.Get("email")

	stmt := fmt.Sprintf("select * from %s.user where email='%s';", constants.MYSQL_DATABASE, newemail)
	rows, _ := sqldb.DB.Query(stmt)

	userExists := rows.Next()

	if userExists {

		res.WriteHeader(http.StatusForbidden)
		response["ModalTitle"] = "Email already registered..."
		response["ModalBody"] = "A User has already registered with the specified Email..Please try again with another Email..."
		panic("403 Error")
	}

	stmt = fmt.Sprintf("update %s.user set firstname = '%s',lastname='%s',email='%s',password='%s',mobile='%s' where email = '%s';", constants.MYSQL_DATABASE, firstname, lastname, newemail, password, mobile, email)
	_, err := sqldb.DB.Query(stmt)

	if err == nil {

		stmt = fmt.Sprintf("update %s.images set email='%s' where email='%s", constants.MYSQL_DATABASE, newemail, email)
		_, _ = sqldb.DB.Exec(stmt)

		res.WriteHeader(http.StatusOK)
		response["ModalTitle"] = "User Updated..."
		response["ModalBody"] = "User details have been updated successfully.Please login again..."

	} else {
		res.WriteHeader(http.StatusForbidden)
		response["ModalTitle"] = "User Updation Failed..."
		response["ModalBody"] = "Updating user has failed..Please try again..."
		panic("403 Error")
	}
	_ = json.NewEncoder(res).Encode(response)
}
