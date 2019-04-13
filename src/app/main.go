package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"app/random"
	"app/redis"

	"github.com/gorilla/mux"
)

type url struct {
	URL string `json:"url"`
}

func main() {
	r := mux.NewRouter()
	// /submit url to be short by check body
	r.HandleFunc("/submit", submitHandler).Methods("POST")
	r.HandleFunc("/{url:[0-9a-zA-Z]{8}}", getHandler).Methods("GET")

	fmt.Println("server running on both 3000 and 3030")
	go http.ListenAndServe(":3030", nil)
	http.ListenAndServe(":3000", r)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var v url
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		http.Error(w, fmt.Sprintf("failed to decode request body into json: %+v", err), http.StatusBadRequest)
		return
	}
	if v.URL == "" {
		http.Error(w, "empty 'url' value", http.StatusBadRequest)
		return
	}
	client := redis.Client
	setDone := false
	for !setDone {
		genKey := random.GenerateRandom(8)
		set, err := client.SetNX(genKey, v.URL, 168*time.Hour).Result()
		if err != nil {
			panic(err)
		} else if set {
			fmt.Println("generated", genKey, v.URL)
			fmt.Fprintf(w, "Your generated short url is %s and it will redirect to %s\n", genKey, v.URL)
			setDone = true
		} else {
			fmt.Println("generated url already exist! And need to generate new one")
			setDone = false
		}
	}

}
func getHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	randomVar := vars["url"]
	fmt.Println("get ", randomVar)
	client := redis.Client
	url, err := client.Get(randomVar).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key", url)
		http.Redirect(w, r, "http://"+url, 301)
	}
}
