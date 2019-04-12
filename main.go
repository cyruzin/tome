package main

import (
	"encoding/json"
	"fmt"
	"log"

	tome "github.com/cyruzin/tome/pkg"
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
		Offset:      0,  // Inicial offset.
		Limit:       10, // Limit per page.
		NewPage:     1,  // Page that you captured in params.
		CurrentPage: 1,  // Inicial Page.
		TotalPages:  60, // Total of pages, this usually comes from SQL query total rows result.
	}

	pg := tome.Paginate(chapter) // Paginating the results.

	data, err := json.MarshalIndent(pg, "", " ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(data)) // Returning JSON.
}
