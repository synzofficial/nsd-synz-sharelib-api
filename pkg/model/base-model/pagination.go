package basemodel

import (
	"math"
)

type PaginationRequest struct {
	PageId   int `json:"page_id"`
	PageSize int `json:"page_size"`
}

type PaginationResponse struct {
	PageId   int  `json:"page_id"`
	PageSize int  `json:"page_size"`
	HasNext  bool `json:"has_next"`
}

type PaginationWithLastPageResponse struct {
	PageId        int `json:"page_id"`
	PageSize      int `json:"page_size"`
	LastPage      int `json:"last_page"`
	TotalElements int `json:"total_elements"`
}

func (p *PaginationWithLastPageResponse) SetLastPage() {
	p.LastPage = int(math.Ceil(float64(p.TotalElements) / float64(p.PageSize)))
}
