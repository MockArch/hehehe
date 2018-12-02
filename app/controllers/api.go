package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type myData struct {
	name string
}

func Hello_api(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Hostname, r.URL.Path, r.RequestURI)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Hello mock"})
}

func Post_method(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, postmethod")
	decoder := json.NewDecoder(r.Body)
	var data myData
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data.name)

}
