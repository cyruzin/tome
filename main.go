package main

import (
	"encoding/json"
	"fmt"
	"log"

	tome "github.com/cyruzin/tome/pkg"
)

func main() {
	chapter := &tome.Chapter{
		Data: struct {
			Nome  string `json:"nome"`
			Email string `json:"email"`
		}{
			"Cyro",
			"xorycx@gmail.com",
		},
		Offset:      0,
		Limit:       10,
		NewPage:     1,
		CurrentPage: 1,
		TotalPages:  60,
	}

	pg := tome.Paginate(chapter)

	data, err := json.MarshalIndent(pg, "", " ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(data))
}
