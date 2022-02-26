package routes

import (
	"backend_golang/constants"
	"backend_golang/router/sqldb"
	"backend_golang/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

//LogoutHandler POST API
func LogoutHandler(res http.ResponseWriter, req *http.Request) {
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
	email := req.Form.Get("email")

	stmt := fmt.Sprintf("update %s.user set token=null where email='%s';", constants.MYSQL_DATABASE, email)
	_, _ = sqldb.DB.Exec(stmt)
	_ = json.NewEncoder(res).Encode("Successfully logged out")
}
