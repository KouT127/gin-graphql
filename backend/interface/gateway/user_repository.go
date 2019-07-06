package gateway

import (
	"database/sql"
	"fmt"
	"gin-sample/backend/domain/model"
	"gin-sample/backend/usecase/form"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	FindAll(p *form.Pagination) ([]*model.User, error)
	Create(user *model.User) (*model.User, error)
	getPointerList(rows *sql.Rows) ([]*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}
func (ur *userRepository) FindAll(p *form.Pagination) ([]*model.User, error) {
	u := ur.db.Model(&model.User{})
	u = p.Paging(u)
	rows, err := u.Rows()
	defer rows.Close()
	users, err := ur.getPointerList(rows)
	if err != nil {
		return users, nil
	}
	return users, nil
}

func (ur *userRepository) Create(user *model.User) (*model.User, error) {
	if err := ur.db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (ur *userRepository) getPointerList(rows *sql.Rows) ([]*model.User, error) {
	var list []*model.User
	for rows.Next() {
		mem := &model.User{}
		err := ur.db.ScanRows(rows, &mem)
		if err != nil {
			fmt.Print(err.Error())
			return list, err
		}
		list = append(list, mem)
	}
	return list, nil
}
