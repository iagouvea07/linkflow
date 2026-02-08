package main

import (
	"api/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func HandleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"status\": \"healthy\"}")
}

func main() {
	db, err := handlers.InitDB()

	if err != nil {
		log.Fatal(err)
	}

	if err = handlers.InitializePosition(db); err != nil {
		return
	}

	route := mux.NewRouter()

	route.HandleFunc("/", HandleHealth).Methods("GET")
	route.HandleFunc("/metrics", handlers.ExposePrometheusMetrics).Methods("GET")
	route.HandleFunc("/insert", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleInsertUrl(w, r, db) 
	}).Methods("POST")

	route.HandleFunc("/{any}", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandlerRedirect(db, r.URL.Path, w, r)
	})

	fmt.Println("API is running...")

	if err = http.ListenAndServe(":8000", route); err != nil {
		log.Fatal(err)
	}
}