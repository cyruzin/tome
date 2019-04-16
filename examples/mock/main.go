package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/cyruzin/tome"
)

func main() {
	// Creating a tome chapter with links.
	chapter := &tome.Chapter{
		Data: struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		}{
			"What is Lorem Ipsum?",
			"Lorem Ipsum is simply dummy text of the printing and...",
		}, // Data that you want to return along with pagination settings.
		BaseURL:      "http://yourapi.com/v1/posts",
		Links:        true,
		NewPage:      2,   // Page that you captured in params.
		TotalResults: 300, // Total of pages, this usually comes from a SQL query total rows result.
	}

	err := chapter.Paginate() // Paginating the results.
	if err != nil {
		log.Panic(err)
	}

	data, err := json.MarshalIndent(chapter, "", " ")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(data)) // Returning JSON.
}
