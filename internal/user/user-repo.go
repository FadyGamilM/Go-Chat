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
