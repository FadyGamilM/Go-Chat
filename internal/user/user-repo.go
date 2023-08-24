package user

import (
	"context"
	"log"

	"github.com/FadyGamilM/Go-Chat/db"
)

type pg_repository struct {
	db db.DBTX
}

func NewUserRepo(db db.DBTX) UserRepo {
	return &pg_repository{
		db: db,
	}
}

func (pgRepo *pg_repository) Create(ctx context.Context, u *User) (*User, error) {
	create_user_query := `
		INSERT INTO users
		(username, email, password)
		VALUES
		($1, $2, $3)
		RETURNING id
	`
	insertedUserID := new(int64)
	err := pgRepo.db.QueryRowContext(ctx, create_user_query, u.Username, u.Email, u.Password).Scan(&insertedUserID)
	if err != nil {
		log.Printf("error while trying to insert a new user %v \n", err)
		return nil, err
	}

	log.Printf("the inserted user id is : %v \n", *insertedUserID)
	// set the id into the received user domain entity
	u.ID = *insertedUserID

	return u, nil
}

func (pgRepo *pg_repository) Login(ctx context.Context, email string) (*User, error) {
	login_user_query := `
		SELECT id, username, password, email 
		FROM users 
		WHERE email = $1
	`
	var err error
	foundUser := new(User)

	// fetch the user with provided email to check if there is user with this email in the system or not
	err = pgRepo.db.QueryRowContext(ctx, login_user_query, email).Scan(&foundUser.ID, &foundUser.Username, &foundUser.Password, &foundUser.Email)
	if err != nil {
		log.Printf("error while trying to fetch user by email : %v \n", err)
		return nil, err
	}

	return foundUser, nil
}
