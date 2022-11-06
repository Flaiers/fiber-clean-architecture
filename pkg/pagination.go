package pkg

import (
	"gorm.io/gorm"
)

const (
	DefaultPage = 1
	DefaultSize = 50
)

type Pagination struct {
	Items any `json:"items"`
	Page  int `json:"page"`
	Size  int `json:"size"`
	Total int `json:"total"`
}

func (p *Pagination) GetPage() int {
	if p.Page <= 0 {
		p.Page = DefaultPage
	}

	return p.Page
}

func (p *Pagination) GetSize() int {
	if p.Size <= 0 {
		p.Size = DefaultSize
	}

	return p.Size
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetSize()
}

func (p *Pagination) Query(db *gorm.DB) *gorm.DB {
	return db.Offset(p.GetOffset()).Limit(p.GetSize())
}

func (p *Pagination) Paginate(items any, total int) *Pagination {
	p.Items = items
	p.Total = total

	return p
}
