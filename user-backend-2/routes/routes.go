package routes

import (
	"go-api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)	

func HandlerRequest() {
	router := gin.Default()

	router.Use(cors.Default())
	router.GET("/users", controllers.GetAllUsers )	
	router.GET("/user/:id", controllers.GetUserById )	
	router.POST("/user", controllers.CreateNewUser )	
	router.PUT("/user/:id", controllers.UpdateUserData )
	router.DELETE("/user/:id", controllers.DeletUserById )
	router.Run(":8002")
}