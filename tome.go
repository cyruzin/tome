// Copyright (c) 2019 Cyro Dubeux. License MIT.

// Package tome was designed to paginate simple RESTful APIs.
package tome

import (
	"errors"
	"math"
	"strconv"
)

// Chapter type is a struct for pagination results.
type Chapter struct {
	// Data that you want to return along with pagination settings.
	Data interface{} `json:"data"`
	// API base URL.
	BaseURL string `json:"base_url"`
	// The next URL link with page number.
	NextURL string `json:"next_url"`
	// The previous URL link with page number.
	PreviousURL string `json:"prev_url"`
	// The inicial offset position.
	offset int
	// Limit per page.
	Limit int `json:"per_page"`
	// The page number captured on the request params.
	NewPage int `json:"-"`
	// Current page of the tome.
	CurrentPage int `json:"current_page"`
	// The last page of the tome.
	LastPage int `json:"last_page"`
	// Total of results, this usually comes from a SQL query total rows result.
	TotalResults int `json:"total_results"`
}

// Paginate handles the pagination calculation.
func (c *Chapter) Paginate() error {
	if c.BaseURL == "" {
		return errors.New("Base URL is missing")
	}

	c.setDefaults()  // Checking if need defaults
	c.ceilLastPage() // Ceiling the last page.
	c.doPaginate()   // Pagination calculation.
	c.createLinks()  // Creating links.

	return nil
}

// Calculates the offset and the limit.
func (c *Chapter) doPaginate() {
	if c.NewPage > c.CurrentPage {
		c.CurrentPage = c.NewPage
		c.offset = (c.CurrentPage - 1) * c.Limit
	}
}

// Ceils the last page and generates
// a integer number.
func (c *Chapter) ceilLastPage() {
	c.LastPage = int(math.Ceil(float64(c.TotalResults) / float64(c.Limit)))
}

// Creates next and previous links using
// the given base URL.
func (c *Chapter) createLinks() {
	if c.CurrentPage < c.LastPage {
		c.NextURL = c.BaseURL + "?page=" + strconv.Itoa(c.CurrentPage+1)
	}

	if c.LastPage > c.CurrentPage {
		c.PreviousURL = c.BaseURL + "?page=" + strconv.Itoa(c.CurrentPage-1)
	}
}

// Sets the defaults values for current page
// and limit if none of them were provided.
func (c *Chapter) setDefaults() {
	if c.CurrentPage == 0 {
		c.CurrentPage = 1
	}

	if c.Limit == 0 {
		c.Limit = 10
	}
}
