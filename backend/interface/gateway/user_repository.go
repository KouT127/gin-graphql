package gateway

import (
	"database/sql"
	"gin-sample/backend/domain/model"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	FindAll() ([]*model.User, error)
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
func (ur *userRepository) FindAll() ([]*model.User, error) {
	rows, err := ur.db.Find(&model.User{}).Rows()
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
	list := []*model.User{}
	for rows.Next() {
		mem := &model.User{}
		err := rows.Scan(&mem.ID, &mem.Name, &mem.BirthDay, &mem.CreatedAt)
		if err != nil {
			return list, err
		}
		list = append(list, mem)
	}
	return list, nil
}
