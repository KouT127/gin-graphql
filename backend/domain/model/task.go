package model

import (
	"time"
)

type Task struct {
	ID          int
	UserRefer   int
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
