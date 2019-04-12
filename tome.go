// Copyright (c) 2019 Cyro Dubeux. License MIT.

// Package tome was designed to paginate simple RESTful APIs.
package tome

import "math"

// Chapter type is a struct for pagination results.
type Chapter struct {
	// Data that you want to return along with pagination settings.
	Data interface{} `json:"data"`
	// The inicial offset position.
	Offset int `json:"-"`
	// Limit per page.
	Limit int `json:"per_page"`
	// The page number captured on the request params.
	NewPage int `json:"-"`
	// Current page of the tome.
	CurrentPage int `json:"current_page"`
	// The last page of the tome.
	LastPage int `json:"last_page"`
	// Total of pages, this usually comes from a SQL query total rows result.
	TotalPages int `json:"total"`
}

// Paginate handles the pagination calculation.
func Paginate(c *Chapter) *Chapter {
	setDefaults(c)                 // Checking if need defaults
	ceilLastPage(c)                // Ceiling the last page.
	offset, limit := doPaginate(c) // Pagination calculation.

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
