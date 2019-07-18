package repository

import (
	"github.com/KouT127/gin-sample/backend/application/form"
	"github.com/KouT127/gin-sample/backend/domain/model"
)

type UserRepository interface {
	FindAll(p *form.Pagination) ([]*model.User, error)
	Create(frm *form.UserForm) (*model.User, error)
	GetUserMaxPage(limit int) int
}