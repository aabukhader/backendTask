package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/aabukhader/backEnd/controllers"
)

func routes() *mux.Router {
	Router := mux.NewRouter()
	Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1 style='text-align:center'>Backend task </h1>"))
	})
	Router.HandleFunc("/api/search/{query}", controllers.Search).Methods("GET")
	Router.HandleFunc("/api/user/Login", controllers.Authenticate).Methods("POST")
	Router.HandleFunc("/api/user/signUp", controllers.Registration).Methods("POST")
	Router.HandleFunc("/api/post/create", controllers.Post).Methods("POST")
	Router.HandleFunc("/api/post/getAll", controllers.GetPosts).Methods("GET")
	return Router
}

func main() {
	fmt.Println("Server Started...")
	http.ListenAndServe(":8000", routes())
}
