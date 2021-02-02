package routes

import (
	"router/sqldb"
)

//IsValidUser returns whether user is valid or not
func IsValidUser(email string,token string) bool{
	if(!VerifyToken(token)){
		return false;
	}

	stmt := "select * from sureify.user where email = '"+email+"' and token='"+token+"'";
	rows, _ := sqldb.DB.Query(stmt);
	userExists := rows.Next()
	return userExists;
}
