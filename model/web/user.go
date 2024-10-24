package web

type UserCreateRequest struct {
	Username string `validate:"required,min=1,max=100" json:"username"`
	Password string `validate:"required,min=6,max=100" json:"password"` 
}

type UserUpdateRequest struct {
	Id       int    `validate:"required" json:"id"`
	Username string `validate:"required,min=1,max=100" json:"username"`
	Password string `validate:"omitempty,min=6,max=100" json:"password"` 
}
type UserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}
