package routes

import (
	"github.com/Shunux-Stuxnet/EMITrace/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRouts *gin.Engine) {
	incomingRouts.POST("/users/signup", controllers.SignUp())
	incomingRouts.POST("/users/login", controllers.Login())
	incomingRouts.POST("/users/addemi", controllers.AddEMI())
	incomingRouts.GET("/users/searchEMI", controllers.SearchEMI())
}
