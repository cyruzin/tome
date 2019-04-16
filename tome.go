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
	// API base URL.
	BaseURL string `json:"base_url,omitempty"`
	// The next URL link with page number.
	NextURL string `json:"next_url,omitempty"`
	// The previous URL link with page number.
	PreviousURL string `json:"prev_url,omitempty"`
	// Whether to create links or not.
	Links bool `json:"-"`
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
	// Total of results, this usually comes from a SQL query total rows result.
	TotalResults int `json:"total_results"`
}

// Paginate handles the pagination calculation.
func (c *Chapter) Paginate() error {
	c.setDefaults() // Checking if need defaults.

	err := c.ceilLastPage() // Ceiling the last page.
	if err != nil {
		return err
	}

	err = c.doPaginate() // Pagination calculation.
	if err != nil {
		return err
	}

	err = c.checkLinks() // Checking if links are necessary.
	if err != nil {
		return err
	}
	return nil
}

// Calculates the offset and the limit.
func (c *Chapter) doPaginate() error {
	if c.NewPage == 0 {
		return errors.New("NewPage value is missing")
	}

	if c.NewPage > c.CurrentPage {
		c.CurrentPage = c.NewPage
		c.Offset = (c.CurrentPage - 1) * c.Limit
	}

	return nil
}

// Ceils the last page and generates
// a integer number.
func (c *Chapter) ceilLastPage() error {
	if c.TotalResults == 0 {
		return errors.New("TotalResults value is missing")
	}

	c.LastPage = int(math.Ceil(float64(c.TotalResults) / float64(c.Limit)))
	return nil
}

// Handles links validations.
func (c *Chapter) checkLinks() error {
	if !c.Links && c.BaseURL != "" {
		return errors.New("Links value is false, set to true")
	}

	if c.Links {
		if err := c.createLinks(); err != nil {
			return err
		}
	}
	return nil
}

// Creates next and previous links using
// the given base URL.
func (c *Chapter) createLinks() error {
	if c.BaseURL == "" {
		return errors.New("BaseURL value is missing")
	}

	if c.CurrentPage < c.LastPage {
		c.NextURL = c.BaseURL + "?page=" + strconv.Itoa(c.CurrentPage+1)
	}

	if c.LastPage > c.CurrentPage {
		c.PreviousURL = c.BaseURL + "?page=" + strconv.Itoa(c.CurrentPage-1)
	}

	return nil
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
