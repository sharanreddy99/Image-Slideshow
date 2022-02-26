package routes

import (
	"backend_golang/constants"
	"backend_golang/router/sqldb"
	"backend_golang/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

//SigninHandler POST API
func SigninHandler(res http.ResponseWriter, req *http.Request) {

	utils.SetHeaders(&res)
	response := make(map[string]string)

	if sqldb.DB == nil {
		response["ModalTitle"] = "Service Unavailable..."
		response["ModalBody"] = "Signup Service is unavailable right now... Please try again later"
		res.WriteHeader(http.StatusServiceUnavailable)
		fmt.Println("WTF Service not available")
		_ = json.NewEncoder(res).Encode(response)
		return
	}

	defer func() {
		if r := recover(); r != nil {
			_ = json.NewEncoder(res).Encode(response)
		}
	}()

	_ = req.ParseMultipartForm(0)

	email := req.Form.Get("email")
	password := req.Form.Get("password")

	stmt := fmt.Sprintf("select * from %s.user where email = '%s' and password = '%s';", constants.MYSQL_DATABASE, email, password)
	rows, _ := sqldb.DB.Query(stmt)
	userExists := rows.Next()

	if userExists {
		token := utils.GenerateToken(email)
		stmt = fmt.Sprintf("update %s.user set token='%s' where email='%s' and password='%s';", constants.MYSQL_DATABASE, token, email, password)
		_, _ = sqldb.DB.Exec(stmt)

		response["status"] = "success"
		response["email"] = email
		response["password"] = password
		response["token"] = token
		res.WriteHeader(http.StatusOK)
	} else {
		response["status"] = "failure"
		response["ModalTitle"] = "Invalid Credentials..."
		response["ModalBody"] = "The Email or Password is incorrect. Please enter valid credentials..."
		res.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(res).Encode(response)
		return
	}
	_ = json.NewEncoder(res).Encode(response)
}
