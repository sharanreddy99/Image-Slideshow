package routes

import (
	"backendgo/router/sqldb"
)

//IsValidUser returns whether user is valid or not
func IsValidUser(token string) string{
	
	email := VerifyToken(token)
	stmt := "select * from sureify.user where email = '"+email+"' and token='"+token+"'";
	rows, _ := sqldb.DB.Query(stmt);
	userExists := rows.Next()
	if userExists {
		return email
	}
		
	return ""
}
