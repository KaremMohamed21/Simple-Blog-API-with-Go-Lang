package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
	"strconv"
	"math/rand"
)

// Post struct (Model)
type Post struct {
	ID 							string  `json:"id"`
	Title 					string  `json:"title"`
	Description 		string  `json:"description"`		
	Author 					*Author `json:"author"`
}

// Author struct
type Author struct {
	FirstName				string  `json:"firstName"`
	LastName 				string	`json:"lastName"`
}

// Declare array of posts
var posts []Post

// Get all the posts
func getAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Send the response
	json.NewEncoder(w).Encode(posts)
}

// Get post by ID
func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _ , post := range posts {
		if params["id"] == post.ID {
			json.NewEncoder(w).Encode(post)
			return
		}
	}

	json.NewEncoder(w).Encode(&Post{})
}

// To Create new post
func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var post Post
	_ = json.NewDecoder(r.Body).Decode(&post)

	post.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	posts = append(posts, post)

	json.NewEncoder(w).Encode(post)
}

// To update post by ID
func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, post := range posts {
		if post.ID == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)

			var post Post
			_ = json.NewDecoder(r.Body).Decode(&post)
			post.ID = params["id"]
			posts = append(posts, post)
			json.NewEncoder(w).Encode(post)

			return
		}
	}

}

// To delete post by ID
func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, post := range posts {
		if post.ID == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(posts)
}


func main() {

	// Write some hard coded data
	posts = append(posts, Post{ID: "1", Description: "this is post",  Title: "Post One", Author: &Author{ FirstName: "John", LastName: "Doe"}})
	posts = append(posts, Post{ID: "2", Description: "this is post",  Title: "Post Two", Author: &Author{ FirstName: "Steve",LastName: "Smith"}})

	// setup the route
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/posts", getAllPosts).Methods("GET")
	r.HandleFunc("/api/v1/posts/{id}", getPost).Methods("GET")
	r.HandleFunc("/api/v1/posts", createPost).Methods("POST")
	r.HandleFunc("/api/v1/posts/{id}", updatePost).Methods("PUT")
	r.HandleFunc("/api/v1/posts/{id}", deletePost).Methods("DELETE")

	// Start server
	fmt.Println("Server running...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

