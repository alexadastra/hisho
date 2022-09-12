package models

import (
	"errors"

	"github.com/alexadastra/hisho/hisho-core-service/pkg/api"
)

// Paging represents page for GET request
type Paging struct {
	IsEnabled bool
	Page      *Page
}

// Page represents a GET request page
type Page struct {
	Size   int64
	Number int64
}

// NewPaging constructs new Paging
func NewPaging(page *api.Page) (*Paging, error) {
	// if page is none or is empty, then paging is disabled
	if page == nil || (page.Limit == 0 && page.Number == 0) {
		return &Paging{
			IsEnabled: false,
			Page:      nil,
		}, nil
	}
	// if some of page fields are empty, return error
	if page.Limit == 0 || page.Number == 0 {
		return nil, errors.New("one or few empty page parameters")
	}
	return &Paging{
		IsEnabled: true,
		Page: &Page{
			Size:   int64(page.Limit),
			Number: int64(page.Number),
		},
	}, nil
}

// ToLimitOffset represents limit + offset to be selected with from the db
func (p *Page) ToLimitOffset() (int64, int64) {
	return p.Size, p.Size * p.Number
}
