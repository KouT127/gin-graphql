package repository

import (
	"github.com/KouT127/gin-sample/backend/application/form"
	"github.com/KouT127/gin-sample/backend/domain/model/user"
)

type UserRepository interface {
	FindAll(p *form.Pagination) ([]*user.User, error)
	Create(frm *form.UserForm) (*user.User, error)
	GetUserMaxPage(limit int) int
}