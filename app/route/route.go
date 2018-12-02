package route

import (
	"github.com/MockArch/hehehe/app/controllers"
	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	routes := mux.NewRouter().StrictSlash(false)

	r := routes.PathPrefix("/api/v1/").Subrouter()

	routes.HandleFunc("/hello", controllers.Hello).Methods("GET")
	r.HandleFunc("/hello", controllers.Hello_api)
	r.HandleFunc("/reg", controllers.Post_method).Methods("POST")
	return routes
}
