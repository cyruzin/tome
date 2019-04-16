package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/cyruzin/tome"
)

// Post type is a struct for a single post.
type Post struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// Posts type is a struct for multiple posts.
type Posts []*Post

// Result type is a struct of posts with pagination.
type Result struct {
	Data *Posts `json:"data"`
	*tome.Chapter
}

func main() {
	// Creating a tome chapter with links.
	chapter := &tome.Chapter{
		// Setting base URL.
		BaseURL: "http://yourapi.com/v1/posts",
		// Enabling link results.
		Links: true,
		// Page that you captured in params inside you handler.
		NewPage: 2,
		// Total of pages, this usually comes from a SQL query total rows result.
		TotalResults: 300,
	}

	err := chapter.Paginate() // Paginating the results.
	if err != nil {
		log.Panic(err)
	}

	// Mocking results with pagination.
	res := &Result{
		Data: &Posts{
			&Post{
				Title: "What is Lorem Ipsum?",
				Body:  "Lorem Ipsum is simply dummy text of the printing and...",
			},
			&Post{
				Title: "Why do we use it?",
				Body:  "It is a long established fact that a reader will be...",
			},
		},
		Chapter: chapter,
	}

	data, err := json.MarshalIndent(res, "", " ")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(data)) // Returning JSON.
}
