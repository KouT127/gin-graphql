package repository

type UserRepository interface {
	GetUserMaxPage(limit int) int
}
