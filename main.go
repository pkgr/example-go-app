package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func CityHandler(res http.ResponseWriter, req *http.Request) {
	data, _ := json.Marshal("{'cities':'San Francisco, Amsterdam, Berlin, New York'}")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write(data)
}

func main() {
	r := mux.NewRouter()
	port := os.Getenv("PORT")
	r.HandleFunc("/cities.json", CityHandler)
	http.Handle("/", r)
	log.Println("Listening on port "+port+"...")
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
