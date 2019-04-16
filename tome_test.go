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
		},
		BaseURL:      "http://yourapi.com/v1/posts",
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
		BaseURL:      "http://yourapi.com/v1/posts",
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
		Limit:        10,
		NewPage:      10,
		CurrentPage:  1,
		TotalResults: 3000,
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
		},
		BaseURL:      "http://yourapi.com/v1/posts",
		TotalResults: 3000,
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
	chapter := &Chapter{
		Data: struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		}{
			"What is Lorem Ipsum?",
			"Lorem Ipsum is simply dummy text of the printing and...",
		},
		BaseURL:      "http://yourapi.com/v1/posts",
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
		BaseURL:      "http://yourapi.com/v1/posts",
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
