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
	Router.HandleFunc("/api/user/Login", controllers.Authenticate).Methods("POST")
	// Router.HandleFunc("/api/user/SignUp", Login).Methods("POST")
	return Router
}

func main() {
	fmt.Println("Lets Start")
	http.ListenAndServe(":8000", routes())
}
