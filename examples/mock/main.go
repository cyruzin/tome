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
			Nome  string `json:"nome"`
			Email string `json:"email"`
		}{
			"Cyro",
			"xorycx@gmail.com",
		}, // Data that you want to return along with pagination settings.
		BaseURL:     "https://feelthemovies.com.br/v1/recommendations",
		Offset:      0,    // Inicial offset.
		Limit:       10,   // Limit per page.
		NewPage:     200,  // Page that you captured in params.
		CurrentPage: 1,    // Inicial Page.
		TotalPages:  3000, // Total of pages, this usually comes from a SQL query total rows result.
	}

	pg, err := tome.Paginate(chapter) // Paginating the results.
	if err != nil {
		log.Println(err)
	}

	data, err := json.MarshalIndent(pg, "", " ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(data)) // Returning JSON.
}
