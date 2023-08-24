package main

import (
	"fmt"

	"github.com/FadyGamilM/Go-Chat/db"
	"github.com/FadyGamilM/Go-Chat/internal/user"
	"github.com/FadyGamilM/Go-Chat/router"
)

func main() {

	connStr := setupConnString()

	dbConnPool, err := db.Connect(connStr)
	if err != nil {
		fmt.Println(err)
	}

	// inject anything implements the DBTX (inject *sql.Tx or *sql.DB)
	userRepo := user.NewUserRepo(dbConnPool.GetDbPool())
	// inject userRepo into userService
	userSrv := user.NewUserService(userRepo)
	// inject userSrv into userHandler
	userHandler := user.NewUserHandler(userSrv)
	// map the handler methods to routes
	router.InitRouter(userHandler)
	router.Start("127.0.0.1:5050")
}

func setupConnString() string {

	// Replace these values with your PostgreSQL container settings
	host := "127.0.0.1"    // Use the container IP if needed
	port := 1234           // Default PostgreSQL port
	username := "go_chat"  // PostgreSQL username
	password := "go_chat"  // PostgreSQL password
	dbname := "go_chat_db" // Database name

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)
	return connStr
}
