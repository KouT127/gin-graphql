package response

type UserResponse struct {
	ID     uint
	Name   string
	Gender string
}

type UsersResponse struct {
	MaxPage int
	Users   []*UserResponse
}
