package dtos

type FindAllDto struct {
	OrderBy     string
	Sort        string
	Search      string
	Limit       uint
	Offset      uint
	ShowDeleted bool
	OnlyDeleted bool
}

type PaginatedResponse[T any] struct {
	Data   []T   `json:"data"`
	Total  int64 `json:"total"`
	Limit  uint  `json:"limit"`
	Offset uint  `json:"offset"`
	Pages  uint  `json:"pages"`
}

type Response struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
}
