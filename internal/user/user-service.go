package user

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/FadyGamilM/Go-Chat/internal/utils"
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
	var err error
	// set the timeout of the passed context and defer the canceling in case of any business logic or data layer method exceed the timeout
	ctx, cancel := context.WithTimeout(ctx, us.timeout)
	defer cancel()

	// hash the password of the req dto data
	u.Password, err = utils.HashPassword(u.Password)
	if err != nil {
		log.Printf("error while trying to hash the user password : %v \n", err)
	}

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
