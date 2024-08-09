package main

import (
	"log"
	"net/http"
	"todo-app-go/internal/db"
	"todo-app-go/internal/handler"

	"github.com/gorilla/mux"
)

func main() {
	db.InitDB()
	defer db.DB.Close()

	r := mux.NewRouter()
	r.HandleFunc("/api/todos", handler.ListTodosHandler).Methods("GET")
	r.HandleFunc("/api/todos", handler.CreateTodoHandler).Methods("POST")
	r.HandleFunc("/api/todos/{id:[0-9]+}", handler.UpdateTodoHandler).Methods("PUT")
	r.HandleFunc("/api/todos/{id:[0-9]+}", handler.DeleteTodoHandler).Methods("DELETE")

	http.Handle("/", r)
	log.Println("Server running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
