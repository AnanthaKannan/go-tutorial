package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string
	Desc    string
	Content string
}

type Articles []Article

func allArticle(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Desc: "some desc", Content: "some content is here"},
		Article{Title: "Test Title-2", Desc: "some desc-2", Content: "some content is here-2"},
	}
	fmt.Println("End point hit: all article end point")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage Endpoint Hit")
}

func postArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "postArticle Hit")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all-article", allArticle).Methods("GET")
	myRouter.HandleFunc("/articles", postArticle).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequests()
}
