package user

import (
	"context"
	"strconv"
	"time"
)

type userService struct {
	repo    UserRepo
	timeout time.Duration
}

func NewUserService(ur UserRepo) *userService {
	return &userService{
		repo:    ur,
		timeout: time.Duration(5) * time.Second,
	}
}

func (us *userService) Create(ctx context.Context, u *CreateUserReq) (*CreateUserRes, error) {
	// set the timeout of the passed context and defer the canceling in case of any business logic or data layer method exceed the timeout
	ctx, cancel := context.WithTimeout(ctx, us.timeout)
	defer cancel()

	// convert the req dto into domain entity to pass to the data layer
	domainUser := u.ToDomainEntity()

	// call the data layer method and set the id of the created user (if created successfully !)
	createdUser, err := us.repo.Create(ctx, domainUser)
	if err != nil {
		return nil, err
	}
	domainUser.ID = createdUser.ID

	// convert the domain entity into response dto to return it to the handler
	userRespDto := &CreateUserRes{
		ID:       strconv.Itoa(int(domainUser.ID)),
		Username: domainUser.Username,
		Email:    domainUser.Email,
	}

	// return the response to the handler component
	return userRespDto, nil
}
