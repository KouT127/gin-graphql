package gateway

import (
	"database/sql"
	"fmt"
	"gin-sample/backend/domain/model"
	"gin-sample/backend/usecase/form"
	"github.com/jinzhu/gorm"
	"math"
)

type UserRepository interface {
	FindAll(p *form.Pagination) ([]*model.User, error)
	Create(frm *form.UserForm) (*model.User, error)
	GetUserMaxPage(limit int) int
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
	user := model.User{}
	u := ur.db.Model(&user).Related(&user.Tasks, "UserRefer").Order("-updated_at")
	u = p.Paging(u)
	rows, err := u.Rows()
	defer rows.Close()
	users, err := ur.getPointerList(rows)
	if err != nil {
		return users, nil
	}
	return users, nil
}

func (ur *userRepository) Create(frm *form.UserForm) (*model.User, error) {
	tx := ur.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	u := model.User{
		Name:     frm.Name,
		BirthDay: "",
		Gender:   frm.Gender,
		PhotoURL: "",
		Active:   true,
	}
	if err := tx.Create(&u).Error; err != nil {
		tx.Rollback()
		return &u, err
	}
	task := model.Task{
		UserRefer:   u.ID,
		Title:       "test",
		Description: "testd",
	}
	if err := tx.Create(&task).Error; err != nil {
		tx.Rollback()
		return &u, err
	}
	return &u, tx.Commit().Error
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

func (ur *userRepository) GetUserMaxPage(limit int) int {
	var cnt int
	ur.db.Model(&[]model.User{}).Count(&cnt)
	return int(math.Ceil(float64(cnt) / float64(limit)))
}
