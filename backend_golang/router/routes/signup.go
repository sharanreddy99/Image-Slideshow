package routes

import (
	"backend_golang/constants"
	"backend_golang/router/sqldb"
	"backend_golang/utils"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

//SignupHandler POST API
func SignupHandler(res http.ResponseWriter, req *http.Request) {

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

	firstname := req.Form.Get("firstname")
	lastname := req.Form.Get("lastname")
	email := req.Form.Get("email")
	mobile := req.Form.Get("mobile")
	password := req.Form.Get("password")

	stmt := fmt.Sprintf("select * from %s.user where email = '%s';", constants.MYSQL_DATABASE, email)
	rows, _ := sqldb.DB.Query(stmt)
	defer rows.Close()

	userExists := rows.Next()
	if userExists {
		response["ModalTitle"] = "Email already registered..."
		response["ModalBody"] = "A User has already registered with the specified Email..Please try again with another Email..."
		res.WriteHeader(http.StatusForbidden)
		panic("403 Error")

	} else {
		stmt = fmt.Sprintf("insert into %s.user values(NULL,'%s','%s','%s','%s','%s',NULL);", constants.MYSQL_DATABASE, firstname, lastname, email, mobile, password)
		_, err := sqldb.DB.Query(stmt)
		if err != nil {
			response["ModalTitle"] = "Registration Failed..."
			response["ModalBody"] = "Error occured during registration..Please try again later..."
			res.WriteHeader(http.StatusForbidden)
			fmt.Println(err.Error())
			panic("403 Error")
		} else {
			response["ModalTitle"] = "Registration Successful..."
			response["ModalBody"] = "Registration has been completed successfully..Please Login..."
			res.WriteHeader(http.StatusCreated)
		}
	}
	_ = json.NewEncoder(res).Encode(response)
}
