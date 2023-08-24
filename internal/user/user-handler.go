package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	srv *userService
}

func NewUserHandler(s *userService) *userHandler {
	return &userHandler{
		srv: s,
	}
}

func (uh *userHandler) HandleCreateUser(c *gin.Context) {
	// define a user request variaable to map the gin request data into it
	var uReq CreateUserReq
	var err error

	// bind the json request into our req dto type
	err = c.ShouldBindJSON(&uReq)
	if err != nil {
		log.Printf("error while trying to bind request data into create user dto :%v \n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uRes, err := uh.srv.Create(c.Request.Context(), &uReq)
	if err != nil {
		log.Printf("error while trying to create user  :%v \n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return the response to the user
	c.JSON(http.StatusCreated, gin.H{"data": uRes})
}
