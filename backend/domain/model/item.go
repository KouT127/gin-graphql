package model

import "time"

type Item struct {
	ID          int
	Name        string
	Description string
	Price       float32
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

type Cart struct {
	ID        int
	UserRefer int
	ItemRefer int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
