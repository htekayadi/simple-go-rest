package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}


func Index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Welcome!\n")
}

func TodoIndex(writer http.ResponseWriter, request *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Study for test"},
	}

	if err := json.NewEncoder(writer).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	todoId := vars["todoId"]
	fmt.Fprintf(writer, "Todo show: %s\n", todoId)
}