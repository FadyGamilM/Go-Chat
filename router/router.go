package router

import (
	"github.com/FadyGamilM/Go-Chat/internal/user"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

// receives all handlers of all our domains
func InitRouter(uh *user.UserHandler) {
	r = gin.Default()
	authRoutes := r.Group("/api/auth")
	authRoutes.POST("/signup", uh.HandleSignup)
	authRoutes.POST("/login", uh.HandleLogin)

}

func Start(addr string) error {
	return r.Run(addr)
}
