package common

// Define a struct and its methods for paging
type Paging struct {
	// use tag "form" because page/limit are passed as query string ?page=1&limit=3
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"-"` // not pass -> set "-"
}

func (p *Paging) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 {
		p.Limit = 1
	}
}
