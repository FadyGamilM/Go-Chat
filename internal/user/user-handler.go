package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	srv *userService
}

func NewUserHandler(s *userService) *UserHandler {
	return &UserHandler{
		srv: s,
	}
}

func (uh *UserHandler) HandleSignup(c *gin.Context) {
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return the response to the user
	c.JSON(http.StatusCreated, gin.H{"data": uRes})
}

func (uh *UserHandler) HandleLogin(c *gin.Context) {
	var err error
	// define the user req variable to map the gin request body into it
	var loginReq LoginUserReq

	// bind the request
	err = c.ShouldBindJSON(&loginReq)
	if err != nil {
		log.Printf("error while trying to bind request data into login user dto : %v \n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// call the service layer to login the user
	loginRes, err := uh.srv.Login(c.Request.Context(), &loginReq)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": loginRes})
}
