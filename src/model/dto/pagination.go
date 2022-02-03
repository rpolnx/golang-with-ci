package dto

type PaginationDTO struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
}
