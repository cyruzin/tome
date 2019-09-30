<p align="center"><img src="./img/logo.png"></p>

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

// GetPosts gets the latest 10 posts with pagination.
func GetPosts(w http.ResponseWriter, r *http.Request) {
	// Creating a tome chapter with links.
	chapter := &tome.Chapter{
		// Setting base URL.
		BaseURL: "http://yourapi.com/v1/posts",
		// Enabling link results.
		Links: true,
		// Page that you captured in params inside you handler.
		NewPage: 2,
		// Total of pages, this usually comes from a SQL query total rows result.
		TotalResults: model.GetPostsTotalResults(),
	}

	// Paginating the results.
	if err := chapter.Paginate(); err != nil { 
		log.Panic(err)
	}

	// Here you pass the offset and limit.
	database, err := model.GetPosts(chapter.Offset, chapter.Limit)
	if err != nil {
		log.Panic(err)
	}

	// Mocking results with pagination.
	res := &Result{Data: database, Chapter: chapter}
    
	w.WriteHeader(http.StatusOK)  // Setting status 200.
	json.NewEncoder(w).Encode(res) // Returning JSON.
}
```

Output: 

```json
{
 "data": [
  {
   "title": "What is Lorem Ipsum?",
   "body": "Lorem Ipsum is simply dummy text of the printing and..."
  },
  {
   "title": "Why do we use it?",
   "body": "It is a long established fact that a reader will be..."
  }
 ],
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
| 200000000  | 7.80  | 0    | 0         |

With links:

| Iterations | ns/op | B/op | allocs/op |
|------------|-------|------|-----------|
| 10000000   | 133   | 96   | 2         |