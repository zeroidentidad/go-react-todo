package router

import (
	"net/http"

	"../middleware"
	"github.com/gorilla/mux"
)

// Router se exporta y se utiliza en main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/task", middleware.GetAllTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task", middleware.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/task/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undoTask/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/deleteAllTask", middleware.DeleteAllTask).Methods("DELETE", "OPTIONS")

	prefix := "/"
	router.PathPrefix("/").Handler(http.StripPrefix(prefix, http.FileServer(http.Dir("../client/build/"))))

	return router
}
