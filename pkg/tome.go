// Copyright (c) 2019 Cyro Dubeux. License MIT.

package tome

import "math"

// Chapter type is a struct for pagination results.
type Chapter struct {
	Data        interface{} `json:"data"`
	Offset      int         `json:"-"`
	Limit       int         `json:"-"`
	NewPage     int         `json:"-"`
	CurrentPage int         `json:"current_page"`
	LastPage    int         `json:"last_page"`
	TotalPages  int         `json:"total_pages"`
}

// Paginate handles the pagination calculation.
func Paginate(c *Chapter) *Chapter {
	setDefaults(c)
	ceilLastPage(c)
	offset, limit := doPaginate(c)

	return &Chapter{
		c.Data,
		offset,
		limit,
		c.NewPage,
		c.CurrentPage,
		c.LastPage,
		c.TotalPages,
	}
}

// Calculates the offset and the limit.
func doPaginate(c *Chapter) (int, int) {
	if c.NewPage > c.CurrentPage {
		c.CurrentPage = c.NewPage
		c.Offset = (c.CurrentPage - 1) * c.Limit
	}
	return c.Offset, c.Limit
}

// Ceils the last page and generates
// a integer number.
func ceilLastPage(c *Chapter) {
	c.LastPage = int(math.Ceil(float64(c.TotalPages) / float64(c.Limit)))
}

// Sets the defaults values for current page
// and limit if none of them were provided.
func setDefaults(c *Chapter) {
	if cp := c.CurrentPage == 0; cp {
		c.CurrentPage = 1
	}

	if l := c.Limit == 0; l {
		c.Limit = 10
	}
}
