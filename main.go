package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

//json file name
var jsonFile = "articles.json"

//single article struct
type Article struct {
	ID      int      `json:"ID"`
	Title   string   `json:"Title"`
	Content string   `json:"Content"`
	Tags    []string `json:"Tags"`
}

//article list struct
type Articles []Article

//read json and return articles slice
func readJsonFile() Articles {
	plan, _ := ioutil.ReadFile(jsonFile)
	var articles Articles
	err := json.Unmarshal(plan, &articles)
	if err != nil {
		fmt.Print(err)
	}
	return articles
}

//return all articles as json format
func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := readJsonFile()

	fmt.Println("Endpoint All Articles")
	err := json.NewEncoder(w).Encode(articles)

	if err != nil {
		fmt.Print(err)
	}
}

//creat new article by post request
func createArticle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.Form.Get("title")
	content := r.Form.Get("content")

	newArticle := Article{
		ID:      3,
		Title:   title,
		Content: content,
		Tags:    []string{"tag1", "tag2"},
	}

	oldArticles := readJsonFile()
	newArticles := append(oldArticles, newArticle)

	articleJson, _ := json.Marshal(newArticles)
	err := ioutil.WriteFile(jsonFile, articleJson, 0644)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Fprintf(w, "New Article Created")
	}
}

//home page route
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint")
}

//http request handler developed using mux package
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/article/create", createArticle).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

//program starts from here
func main() {
	handleRequests()
}
