package constants

import "os"

var (
	MYSQL_USERNAME string = os.Getenv("IMAGE_VIEWER_MYSQL_USER")
	MYSQL_PASSWORD string = os.Getenv("IMAGE_VIEWER_MYSQL_PASSWORD")
	MYSQL_PORT     string = os.Getenv("IMAGE_VIEWER_MYSQL_PORT")
	MYSQL_HOST     string = os.Getenv("IMAGE_VIEWER_MYSQL_HOST")
	MYSQL_DATABASE string = os.Getenv("IMAGE_VIEWER_MYSQL_DATABASE")
	GOLANG_PORT    string = os.Getenv("IMAGE_VIEWER_GOLANG_PORT")
)
