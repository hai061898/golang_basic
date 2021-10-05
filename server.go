package main

import (
	"api1/config"
    "api1/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection() 
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
)
func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/login", authController.Register)
	}
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello",
	// 	})
	// })

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}