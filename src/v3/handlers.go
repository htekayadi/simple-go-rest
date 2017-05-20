package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

func Index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Welcome!\n")
}

func TodoIndex(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Todo Index!\n")
}

func TodoShow(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	todoId := vars["todoId"]
	fmt.Fprintf(writer, "Todo show: %s\n", todoId)
}