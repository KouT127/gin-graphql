package graph

import (
	"github.com/KouT127/gin-sample/backend/domain/model"
	"strconv"
)

type UserConnection struct {
	TotalCount int
	PageInfo   *PageInfo
	Edges      []*UserEdge
}
type UserEdge struct {
	Cursor string
	Node   *User
}

type User struct {
	ID       string
	Name     string
	BirthDay string
	Gender   string
	PhotoURL string
	Active   bool
}

func NewUser(user *model.User) *User {
	u := &User{
		ID:       strconv.Itoa(int(user.ID)),
		Name:     user.Name,
		BirthDay: user.BirthDay,
		Gender:   user.Gender,
		PhotoURL: user.PhotoURL,
		Active:   user.Active,
	}
	return u
}
