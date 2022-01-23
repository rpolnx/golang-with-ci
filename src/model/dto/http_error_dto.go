package dto

import (
	"time"
)

type ErrorDTO struct {
	Timestamp time.Time
	Status    int
	Error     string
	Message   string
	Path      string
}
