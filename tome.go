// Copyright (c) 2019 Cyro Dubeux. License MIT.

// Package tome was designed to paginate simple RESTful APIs.
package tome

import (
	"errors"
	"math"
	"strconv"
)

// Chapter handles pagination results.
type Chapter struct {
	// The base URL for the endpoint.
	// It is only necessary when using links.
	// Will be omitted from JSON when links are set to false.
	BaseURL string `json:"base_url,omitempty"`
	// The next URL string.
	// Will be omitted from JSON when links are set to false.
	NextURL string `json:"next_url,omitempty"`
	// The previous URL string.
	// Will be omitted from JSON when links are set to false.
	PreviousURL string `json:"prev_url,omitempty"`
	// Whether to create links or not.
	// Pagination without links is faster.
	Links bool `json:"-"`
	// The inicial offset position.
	Offset int `json:"-"`
	// The limit per page.
	// If none is provided, the limit will be setted to 10.
	Limit int `json:"per_page"`
	// The new page number captured on the request params.
	// Will be omitted from JSON, since there is no need for it.
	NewPage int `json:"-"`
	// The current page of the tome.
	// If none is provided, the current page will be setted to 1.
	CurrentPage int `json:"current_page"`
	// The last page of the tome.
	LastPage int `json:"last_page"`
	// The total results, this usually comes from
	// a database query.
	TotalResults int `json:"total_results"`
}

// Paginate handles the pagination calculation.
func (c *Chapter) Paginate() error {
	c.setDefaults()

	if err := c.ceilLastPage(); err != nil {
		return err
	}

	if err := c.doPaginate(); err != nil {
		return err
	}

	if err := c.checkLinks(); err != nil {
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
