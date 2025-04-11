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
