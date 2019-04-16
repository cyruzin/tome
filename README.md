# Tome

[![Build Status](https://travis-ci.org/cyruzin/tome.svg?branch=master)](https://travis-ci.org/cyruzin/tome) [![Coverage Status](https://coveralls.io/repos/github/cyruzin/tome/badge.svg?branch=master)](https://coveralls.io/github/cyruzin/tome?branch=master) [![GoDoc](https://godoc.org/github.com/cyruzin/tome?status.svg)](https://godoc.org/github.com/cyruzin/tome) [![Go Report Card](https://goreportcard.com/badge/github.com/cyruzin/tome)](https://goreportcard.com/report/github.com/cyruzin/tome) [![GitHub license](https://img.shields.io/github/license/Naereen/StrapDown.js.svg)](https://github.com/Naereen/StrapDown.js/blob/master/LICENSE)

Package tome was designed to paginate simple RESTful APIs.

## Installation

```sh
go get -u github.com/cyruzin/tome
```
## Usage

To get started, import the `tome` package and initiate the pagination:

```go
import "github.com/cyruzin/tome"

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
	Links: 		  true, // Create links.
	Limit:        10,  // Limit per page.
	NewPage:      2,   // Page that you captured in params.
	CurrentPage:  1,   // Inicial Page.
	TotalResults: 300, // Total of pages, this usually comes from a SQL query total rows result.
}

err := chapter.Paginate() // Paginating the results.
 if err != nil {
	log.Println(err)
 }
    
w.WriteHeader(http.StatusOK)  // Setting status 200 (Inside your handler).
json.NewEncoder(w).Encode(chapter) // Returning JSON.
```

Output: 

```json
{
 "data": {
  "title": "What is Lorem Ipsum?",
  "body": "Lorem Ipsum is simply dummy text of the printing and..."
 },
 "base_url": "http://yourapi.com/v1/posts",
 "next_url": "http://yourapi.com/v1/posts?page=3",
 "prev_url": "http://yourapi.com/v1/posts?page=1",
 "per_page": 10,
 "current_page": 2,
 "last_page": 30,
 "total_results": 300
}
```

## Performance

Without links:

| Iterations | ns/op | B/op | allocs/op |
|------------|-------|------|-----------|
| 300000000  | 5.20  | 0    | 0         |

With links:

| Iterations | ns/op | B/op | allocs/op |
|------------|-------|------|-----------|
| 1000000    | 120   | 96   | 2         |