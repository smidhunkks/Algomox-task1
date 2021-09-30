package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func openjson() {
	type user struct {
		id       string
		name     string
		keywords []string
	}
	var users []user
	data, err := ioutil.ReadFile("./license.json")
	if err != nil {
		fmt.Print(err)
	} else {
		err = json.Unmarshal(data, &users)
		fmt.Println(users)
	}
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	openjson()
	fmt.Println("Endpoint Hit: homePage")
}
func savePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the SaveFile!")
	fmt.Println("Endpoint Hit: saveFile")
}
func handleRequests() {
	http.HandleFunc("/readFile", homePage)
	http.HandleFunc("/saveFile", savePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
