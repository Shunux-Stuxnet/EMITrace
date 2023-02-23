package main

import (
	"log"
	"os"

	"github.com/Shunux-Stuxnet/EMITrace/controllers"
	"github.com/Shunux-Stuxnet/EMITrace/database"
	"github.com/Shunux-Stuxnet/EMITrace/middleware"
	"github.com/Shunux-Stuxnet/EMITrace/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	app := controllers.NewApplication(database.EMIdata(database.Client, "EMI"), database.Userdata(database.Client, "Users"))
	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())
	router.GET("/listcart", controllers.GetItemFromCart())
	router.GET("/addemi", app.AddEMI())
	router.GET("/removemi", app.RemovEMI())

	log.Fatal(router.Run(":" + port))
}
