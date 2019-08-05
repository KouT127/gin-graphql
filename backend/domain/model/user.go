package model

import "time"

type User struct {
	ID        int
	Name      string
	BirthDay  string
	Gender    string
	PhotoURL  string
	Active    bool
	Tasks     []Task
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
