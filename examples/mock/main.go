package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/cyruzin/tome"
)

func main() {
	// Creating a tome chapter.
	chapter := &tome.Chapter{
		Data: struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		}{
			"What is Lorem Ipsum?",
			"Lorem Ipsum is simply dummy text of the printing and...",
		}, // Data that you want to return along with pagination settings.
		BaseURL:     "http://yourapi.com/v1/posts",
		Offset:      0,    // Inicial offset.
		Limit:       10,   // Limit per page.
		NewPage:     2,    // Page that you captured in params.
		CurrentPage: 1,    // Inicial Page.
		TotalPages:  3000, // Total of pages, this usually comes from a SQL query total rows result.
	}

	pg, err := chapter.Paginate() // Paginating the results.
	if err != nil {
		log.Println(err)
	}

	data, err := json.MarshalIndent(pg, "", " ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(data)) // Returning JSON.
}
