package tome

import (
	"testing"
)

func TestPaginate(t *testing.T) {
	chapter := &Chapter{
		Data: struct {
			Nome  string `json:"nome"`
			Email string `json:"email"`
		}{
			"Cyro",
			"xorycx@gmail.com",
		}, // Data that you want to return along with pagination settings.
		BaseURL:     "http://yourapi.com/v1/posts", // End-point base URL.
		Limit:       10,                            // Limit per page.
		NewPage:     10,                            // Page that you captured in params.
		CurrentPage: 1,                             // Inicial Page.
		TotalPages:  3000,                          // Total of pages, this usually comes from a SQL query total rows result.
	}

	result, err := chapter.Paginate()
	if err != nil {
		t.Error(err)
	}

	if result.CurrentPage != 10 {
		t.Errorf("Expecting: %d, got: %d", 10, result.CurrentPage)
	}
}

func BenchmarkPaginate(b *testing.B) {
	chapter := &Chapter{
		Data: struct {
			Nome  string `json:"nome"`
			Email string `json:"email"`
		}{
			"Cyro",
			"xorycx@gmail.com",
		}, // Data that you want to return along with pagination settings.
		BaseURL:     "http://yourapi.com/v1/posts", // End-point base URL.
		Limit:       10,                            // Limit per page.
		NewPage:     10,                            // Page that you captured in params.
		CurrentPage: 1,                             // Inicial Page.
		TotalPages:  3000,                          // Total of pages, this usually comes from a SQL query total rows result.
	}

	result, err := chapter.Paginate()
	if err != nil {
		b.Error(err)
	}

	b.Log(result.CurrentPage)
	b.ReportAllocs()

}
