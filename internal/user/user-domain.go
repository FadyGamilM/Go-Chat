package user

import "context"

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(username, email, password string) *User {
	return &User{
		Username: username,
		Email:    email,
		Password: password,
	}
}

// port of data layer logic
type UserRepo interface {
	Create(context.Context, *User) (*User, error)
}

// port of usecases logic
type UserService interface {
	Create(ctx context.Context, u *CreateUserReq) (*CreateUserRes, error)
}

// dtos between the handlers and services (usecases)
type CreateUserReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type CreateUserRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (req *CreateUserReq) ToDomainEntity() *User {
	return &User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
}
