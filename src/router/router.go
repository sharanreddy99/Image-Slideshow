package router

import (
	"log"
	"net/http"
	routes "router/routes"
	"router/sqldb"

	//package for mysql

	"github.com/gorilla/mux"
)

//HandleRequests handle all the api calls for DynamicImageSlideshow
func HandleRequests() {

	sqldb.CreateConnection()
	myRouter := mux.NewRouter()

	//Routes
	myRouter.HandleFunc("/signup", routes.SignupHandler).Methods("POST")
	myRouter.HandleFunc("/signin", routes.SigninHandler).Methods("POST")
	myRouter.HandleFunc("/getallimages", routes.GetAllImagesHandler).Methods("POST")
	myRouter.HandleFunc("/getsingleimage", routes.GetSingleImageHandler).Methods("POST")
	myRouter.HandleFunc("/imageupload", routes.ImageUploadHandler).Methods("POST")
	myRouter.HandleFunc("/deletesingleimage", routes.DeleteSingleImageHandler).Methods("POST")
	myRouter.HandleFunc("/fetchuser", routes.FetchUserHandler).Methods("POST")
	myRouter.HandleFunc("/isvaliduser", routes.IsValidUserHandler).Methods("POST")
	myRouter.HandleFunc("/logout", routes.LogoutHandler).Methods("POST")
	myRouter.HandleFunc("/updateuser", routes.UpdateUserHandler).Methods("POST")
	

	http.Handle("/", myRouter)
	log.Fatal(http.ListenAndServe(":4201", nil))
}
