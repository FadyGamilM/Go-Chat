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
	Login(ctx context.Context, email string) (*User, error)
}

// port of usecases logic
type UserService interface {
	Create(ctx context.Context, u *CreateUserReq) (*CreateUserRes, error)
	Login(ctx context.Context, u *LoginUserReq) (*LoginUserRes, error)
}

// dtos between the handlers and services (usecases)

// ==> for signup request
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

// ==> For login request
type LoginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginUserRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (req *LoginUserReq) ToDomainEntity() *User {
	return &User{
		Email:    req.Email,
		Password: req.Password,
	}
}
