package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
)

type Article struct {
	Id string `json:"Id"`
	Title string `json:"Title"`
	Desc string `json:"Desc"`
	Content string `json:"Content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
    // return the string response containing the request body 
    reqBody, _ := ioutil.ReadAll(r.Body)
    
    var article Article
    json.Unmarshal(reqBody, &article)

    // update our global Articles array to include
    // our new Article
    Articles = append(Articles, article)
    json.NewEncoder(w).Encode(Articles)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)

	// we will need to extract the `id` of the article we
    // wish to delete
    id := vars["id"]

    // we then need to loop through all our articles
    for index, article := range Articles {
    	// if our id path parameter matches one of our articles
    	if article.Id == id {
    		// updates our Articles array to remove the article
    		Articles = append(Articles[:index], Articles[index+1:]...)
    		break
    	}
    }
}

func handleRequest() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	// replace http.HandleFunc with myRouter.HandleFunc
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/all", returnAllArticles)
    myRouter.HandleFunc("/article/{id}", returnSingleArticle)
    myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
    myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")

    // finally, instead of passing in nil, we want
    // to pass in our newly created router as the second argument
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article {
		Article{Id:"1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{Id:"2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequest()
}