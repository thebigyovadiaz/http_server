package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprintf(writer, "Welcome to my Website!")
		if err != nil {
			fmt.Println("Server Error: ", err.Error())
			return
		}

		fmt.Println("GET HELLO PATH ")
		return
	}).Methods("GET")

	// Path Prefix and Subrouters
	bookRouter := router.PathPrefix("/books").Subrouter()

	bookRouter.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		_, err := fmt.Fprintf(writer, "List All BOOKS")
		if err != nil {
			fmt.Println("Server Error: ", err.Error())
			return
		}

		fmt.Println("GET ALL BOOKS")
		return
	}).Methods("GET")

	bookRouter.HandleFunc("/{title}", func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		title := vars["title"]

		_, err := fmt.Fprintf(writer, "You've requested the book: %s\n", title)
		if err != nil {
			fmt.Println("Server Error: ", err.Error())
			return
		}

		fmt.Println("GET BOOK TITLE")
		return
	}).Methods("GET")

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	err := http.ListenAndServe(":2323", router)
	if err != nil {
		fmt.Println("Server Error: ", err.Error())
		return
	}
}
