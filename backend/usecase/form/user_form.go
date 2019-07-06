package form

type UserForm struct {
	Name   string `json:"name" binding:"required,max=24"`
	Gender string `json:"gender" binding:"required,max=5"`
}
