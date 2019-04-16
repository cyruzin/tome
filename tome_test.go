package tome

import (
	"testing"
)

func TestPaginate(t *testing.T) {
	chapter := &Chapter{
		Data: struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		}{
			"What is Lorem Ipsum?",
			"Lorem Ipsum is simply dummy text of the printing and...",
		}, // Data that you want to return along with pagination settings.
		BaseURL:      "http://yourapi.com/v1/posts", // End-point base URL.
		Limit:        10,                            // Limit per page.
		NewPage:      10,                            // Page that you captured in params.
		CurrentPage:  1,                             // Inicial Page.
		TotalResults: 3000,                          // Total of pages, this usually comes from a SQL query total rows result.
	}

	err := chapter.Paginate()
	if err != nil {
		t.Error(err)
	}

	if chapter.CurrentPage != 10 {
		t.Errorf("Expecting: %d, got: %d", 10, chapter.CurrentPage)
	}
}

func TestEmptyBaseURL(t *testing.T) {
	chapter := &Chapter{
		Data: struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		}{
			"What is Lorem Ipsum?",
			"Lorem Ipsum is simply dummy text of the printing and...",
		}, // Data that you want to return along with pagination settings.
		BaseURL:      "",   // End-point base URL.
		Limit:        10,   // Limit per page.
		NewPage:      10,   // Page that you captured in params.
		CurrentPage:  1,    // Inicial Page.
		TotalResults: 3000, // Total of pages, this usually comes from a SQL query total rows result.
	}

	err := chapter.Paginate()
	if err.Error() != "Base URL is missing" {
		t.Errorf("Expecting: %s, got: %s", "Base URL is missing", err.Error())
	}
}

func TestDefaultValues(t *testing.T) {
	chapter := &Chapter{
		Data: struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		}{
			"What is Lorem Ipsum?",
			"Lorem Ipsum is simply dummy text of the printing and...",
		}, // Data that you want to return along with pagination settings.
		BaseURL:      "http://yourapi.com/v1/posts", // End-point base URL.
		TotalResults: 3000,                          // Total of pages, this usually comes from a SQL query total rows result.
	}

	err := chapter.Paginate()
	if err != nil {
		t.Error(err)
	}

	if chapter.CurrentPage != 1 {
		t.Errorf("Expecting: %d, got: %d", 1, chapter.CurrentPage)
	}

	if chapter.Limit != 10 {
		t.Errorf("Expecting: %d, got: %d", 10, chapter.Limit)
	}
}

func BenchmarkPaginate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		chapter := &Chapter{
			Data: struct {
				Title string `json:"title"`
				Body  string `json:"body"`
			}{
				"What is Lorem Ipsum?",
				"Lorem Ipsum is simply dummy text of the printing and...",
			}, // Data that you want to return along with pagination settings.
			BaseURL:      "http://yourapi.com/v1/posts", // End-point base URL.
			Limit:        10,                            // Limit per page.
			NewPage:      10,                            // Page that you captured in params.
			CurrentPage:  1,                             // Inicial Page.
			TotalResults: 3000,                          // Total of pages, this usually comes from a SQL query total rows result.
		}
		err := chapter.Paginate()
		if err != nil {
			b.Error(err)
		}
		b.ReportAllocs()
	}
}
