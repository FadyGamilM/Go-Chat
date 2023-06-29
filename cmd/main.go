package main

import (
	"fmt"

	"github.com/FadyGamilM/Go-Chat/db"
)

func main() {
	_, err := db.Connect()
	if err != nil {
		fmt.Println(err)
	}
}
