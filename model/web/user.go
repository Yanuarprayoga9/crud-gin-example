package web

type UserCreateRequest struct {
	Username string `validate:"required,min=1,max=100" json:"username"`
	Password string `validate:"required,min=6,max=100" json:"password"` // Password should have a reasonable length, e.g., min=6.
}

type UserUpdateRequest struct {
	Id       int    `validate:"required" json:"id"`
	Username string `validate:"required,min=1,max=100" json:"username"`
	Password string `validate:"omitempty,min=6,max=100" json:"password"` // Password is optional during updates
}
type UserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}
