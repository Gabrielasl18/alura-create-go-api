package routes

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"API-GO-REST/controllers"
	"API-GO-REST/middleware"
)

//lidar com as requisições
func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home);
	r.HandleFunc("/api/personalidades", controllers.ExibePersonalidades).Methods("Get");
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get");
	r.HandleFunc("/api/personalidades", controllers.CriaPersonalidade).Methods("Post");
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaPersonalidade).Methods("Delete");
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaPersonalidade).Methods("Put");
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)));
}