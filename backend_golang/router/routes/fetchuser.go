package routes

import (
	"backend_golang/constants"
	"backend_golang/router/sqldb"
	"backend_golang/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

//FetchUserHandler POST API
func FetchUserHandler(res http.ResponseWriter, req *http.Request) {
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

	stmt := fmt.Sprintf("select firstname,lastname,mobile,password from %s.user where email='%s';", constants.MYSQL_DATABASE, email)
	rows, _ := sqldb.DB.Query(stmt)

	userExists := rows.Next()

	if userExists {
		firstname := ""
		lastname := ""
		mobile := ""
		password := ""
		_ = rows.Scan(&firstname, &lastname, &mobile, &password)

		response["firstname"] = firstname
		response["lastname"] = lastname
		response["mobile"] = mobile
		response["password"] = password
		res.WriteHeader(http.StatusOK)
	}
	_ = json.NewEncoder(res).Encode(response)
}
