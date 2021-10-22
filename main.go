package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Article struct {
	ID	string	`json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Author string `json:"author"`
	CreateDate time.Time
	LastUpdated time.Time
}

var articles []Article

func handleRequests() {
	router := mux.NewRouter()

	router.HandleFunc("/api/articles", getArticles).Methods("GET")
	router.HandleFunc("/api/articles/{id}", getArticle).Methods("GET")
	router.HandleFunc("/api/articles", createArticle).Methods("POST")
	router.HandleFunc("/api/articles/{id}", updateArticle).Methods("PUT")
	router.HandleFunc("/api/articles/{id}", deleteArticle).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newArticle Article
	json.NewDecoder(r.Body).Decode(&newArticle)
	newArticle.ID = strconv.Itoa(len(articles)+1)
	newArticle.CreateDate = time.Now()
	newArticle.LastUpdated = time.Now() 

	articles = append(articles, newArticle)

	json.NewEncoder(w).Encode(articles)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range articles {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Article{})
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var articleToUpdate Article
	json.NewDecoder(r.Body).Decode(&articleToUpdate)

	for index, item := range articles {
		if item.ID == params["id"] {
			if articleToUpdate.Title != "" {
				articles[index].Title = articleToUpdate.Title
			}

			if articleToUpdate.Description != "" {
				articles[index].Description = articleToUpdate.Description
			}

			if articleToUpdate.Author != "" {
				articles[index].Author = articleToUpdate.Author
			}
			
			item.LastUpdated = time.Now()
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(articleToUpdate)

}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range articles {
		if item.ID == params["id"] {
			articles = append(articles[:index], articles[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(articles)
}

func main(){
	fmt.Println("Hello world")

	var x int
	x = 1

	y := 1

	fmt.Println("x: ", x)
	fmt.Println("y: ", y)

	var arr [5]int
	fmt.Println("arr", arr)
	
	arr[0] = 1
	arr[1] = 2
	
	fmt.Println("arr", arr)
	
	arr = [5]int{5,4,3,2,1}
	fmt.Println("arr", arr)
	
	slice := []int{1,2,3,4,5}
	fmt.Println("slice", slice)
	
	slice = append(slice, 6)
	fmt.Println("slice", slice)
	
	testmap := make(map[string]string)
	fmt.Println("map", testmap)
	testmap["key1"] = "value1"
	testmap["key2"] = "value2"
	fmt.Println("map", testmap)
	
	delete(testmap, "key2")
	fmt.Println("map", testmap)

	for i := 0; i < 10; i++ {
		fmt.Println("i: ", i)
	}

	for index, value := range arr {
		fmt.Print("index: ", index)
		fmt.Println("value: ", value)
	}
	articles = append(articles, Article{ID: "1", Title: "Article 1", Description: "Article 1 Description", Author:  "Bobby", CreateDate: time.Date(2021, time.October, 10, 0, 0, 0, 0, time.UTC), LastUpdated: time.Date(2021, time.October, 10, 0, 0, 0, 0, time.UTC)})
	articles = append(articles, Article{ID: "2", Title: "Article 2", Description: "Article 2 Description", Author:  "Robert", CreateDate: time.Date(2021, time.October, 11, 13, 10, 0, 0, time.UTC), LastUpdated: time.Date(2021, time.October, 11, 13, 10, 0, 0, time.UTC)})
	handleRequests()
}