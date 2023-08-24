package user

import (
	"context"
	"errors"
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
	log.Println("the hashed password is > ", u.Password)

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

func (us *userService) Login(ctx context.Context, u *LoginUserReq) (*LoginUserRes, error) {
	var err error

	// set the context timeout to pass it to the business logic and repo logic
	ctx, cancel := context.WithTimeout(ctx, us.timeout)
	defer cancel()

	// convert the login dto into user domain entity to pass it to the data layer
	user := u.ToDomainEntity()

	log.Println("the password sent in login request is : ", user.Password)
	hashIt, _ := utils.HashPassword(user.Password)
	log.Println("after we hash it : ", hashIt)

	// we need to check if this email exists or not
	registeredUser, err := us.repo.Login(ctx, user.Email)
	println("the user retrieved from database : ", registeredUser.Password)
	if err != nil {
		// check the error type and build a robust error handler mechanism later on
		return nil, err
	}

	// then we need to check the password against the hashed stored password in database
	isMatching, err := utils.CheckPassword(user.Password, registeredUser.Password)
	if err != nil {
		return nil, err
	}
	if !isMatching {
		return nil, errors.New("incorrect credentials")
	}

	// convert the retrieved domain entity into response dto
	resDto := &LoginUserRes{
		ID:       strconv.Itoa(int(registeredUser.ID)),
		Username: registeredUser.Username,
		Email:    registeredUser.Email,
	}

	// return the result
	return resDto, nil
}
