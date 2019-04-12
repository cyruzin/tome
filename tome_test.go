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
		},
		Offset:      1,
		Limit:       30,
		NewPage:     10,
		CurrentPage: 2,
		TotalPages:  100,
	}

	result := Paginate(chapter)

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
		},
		Offset:      1,
		Limit:       30,
		NewPage:     10,
		CurrentPage: 2,
		TotalPages:  100,
	}

	result := Paginate(chapter)

	b.Log(result.CurrentPage)
	b.ReportAllocs()

}
