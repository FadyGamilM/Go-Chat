package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("error while trying to hash user password")
	}
	return string(hashed), nil
}

// this function recieves the `providedPass` which is provided by user at login time and the `hashedPass` which is the stored hashed password in database to compare them
func CheckPassword(providedPass, hashedPass string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(providedPass))
	if err != nil {
		log.Printf("error while trying to compare the provided password against the hashed password from database %v \n", err)
		return false, err
	}
	return true, nil
}
