package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Content struct
type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"description"`
	Content string `json:"content"`
}

var Articles []Article // Database of contents

// HomePage endpoint => http://127.0.0.1:5656/
func home_page(resp_write http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp_write, "Welcome!!!")
	fmt.Println("Endpoint hit: /")
}

// AllArticles endpoint => http://127.0.0.1:5656/all_articles
func all_articles(resp_write http.ResponseWriter, req *http.Request) {
	fmt.Println("Endpoint hit: /all_articles")
	json.NewEncoder(resp_write).Encode(Articles)
}

// Main route function
func handle_request() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/articles", all_articles)
	log.Fatal(http.ListenAndServe(":5656", nil))
}

func main() {
	// Contents
	Articles = []Article{
		Article{Title: "Title 1", Desc: "Desc 1", Content: "Content 1"},
		Article{Title: "Title 2", Desc: "Desc 2", Content: "Content 2"},
	}
	handle_request()
}
