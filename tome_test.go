package tome

import (
	"testing"
)

const (
	baseURL           = "http://yourapi.com/v1/posts"
	baseURLError      = "BaseURL value is missing"
	newPageError      = "NewPage value is missing"
	totalResultsError = "TotalResults value is missing"
	linksFalseError   = "Links value is false, set to true"
)

func TestPaginate(t *testing.T) {
	chapter := &Chapter{
		Data: struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		}{
			"What is Lorem Ipsum?",
			"Lorem Ipsum is simply dummy text of the printing and...",
		},
		Limit:        10,
		NewPage:      10,
		CurrentPage:  1,
		TotalResults: 3000,
	}

	err := chapter.Paginate()
	if err != nil {
		t.Error(err)
	}

	if chapter.CurrentPage != 10 {
		t.Errorf("Expecting: %d, got: %d", 10, chapter.CurrentPage)
	}
}

func TestPaginateWithLinks(t *testing.T) {
	chapter := &Chapter{
		Data: struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		}{
			"What is Lorem Ipsum?",
			"Lorem Ipsum is simply dummy text of the printing and...",
		},
		BaseURL:      baseURL,
		Links:        true,
		Limit:        10,
		NewPage:      10,
		CurrentPage:  1,
		TotalResults: 3000,
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
		},
		BaseURL:      "",
		Links:        true,
		Limit:        10,
		NewPage:      10,
		CurrentPage:  1,
		TotalResults: 3000,
	}

	err := chapter.Paginate()
	if err.Error() != baseURLError {
		t.Errorf("Expecting: %s, got: %s", baseURLError, err.Error())
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
		},
		NewPage:      2,
		TotalResults: 3000,
	}

	err := chapter.Paginate()
	if err != nil {
		t.Error(err)
	}

	if chapter.CurrentPage != 2 {
		t.Errorf("Expecting: %d, got: %d", 2, chapter.CurrentPage)
	}

	if chapter.Limit != 10 {
		t.Errorf("Expecting: %d, got: %d", 10, chapter.Limit)
	}
}

func TestEmptyNewPage(t *testing.T) {
	chapter := &Chapter{
		Data: struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		}{
			"What is Lorem Ipsum?",
			"Lorem Ipsum is simply dummy text of the printing and...",
		},
		BaseURL:      baseURL,
		TotalResults: 3000,
	}

	err := chapter.Paginate()

	if err.Error() != newPageError {
		t.Errorf("Expecting: %s, got: %s", newPageError, err.Error())
	}
}

func TestEmptyTotalResults(t *testing.T) {
	chapter := &Chapter{
		Data: struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		}{
			"What is Lorem Ipsum?",
			"Lorem Ipsum is simply dummy text of the printing and...",
		},
		NewPage: 2,
		BaseURL: baseURL,
	}

	err := chapter.Paginate()

	if err.Error() != totalResultsError {
		t.Errorf("Expecting: %s, got: %s", totalResultsError, err.Error())
	}
}

func TestEmptyLinksWithBaseURL(t *testing.T) {
	chapter := &Chapter{
		Data: struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		}{
			"What is Lorem Ipsum?",
			"Lorem Ipsum is simply dummy text of the printing and...",
		},
		NewPage:      2,
		BaseURL:      baseURL,
		TotalResults: 300,
	}

	err := chapter.Paginate()

	if err.Error() != linksFalseError {
		t.Errorf("Expecting: %s, got: %s", linksFalseError, err.Error())
	}
}

func BenchmarkPaginate(b *testing.B) {
	chapter := &Chapter{
		Data: struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		}{
			"What is Lorem Ipsum?",
			"Lorem Ipsum is simply dummy text of the printing and...",
		},
		Limit:        10,
		NewPage:      10,
		CurrentPage:  1,
		TotalResults: 3000,
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := chapter.Paginate()
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkPaginateWithLinks(b *testing.B) {
	chapter := &Chapter{
		Data: struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		}{
			"What is Lorem Ipsum?",
			"Lorem Ipsum is simply dummy text of the printing and...",
		},
		BaseURL:      baseURL,
		Links:        true,
		Limit:        10,
		NewPage:      10,
		CurrentPage:  1,
		TotalResults: 3000,
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := chapter.Paginate()
		if err != nil {
			b.Error(err)
		}
	}
}
