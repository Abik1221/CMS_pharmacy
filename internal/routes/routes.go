package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/your_project/internal/auth"
	"github.com/your_project/internal/user"
	"github.com/your_project/pkg/jwt"
)

func RegisterRoutes(router *gin.Engine) {
	// Group base API path
	api := router.Group("/api")

	// === Dependencies (init once here) ===
	userRepo := user.NewUserRepository()
	userService := user.NewUserService(userRepo)
	userHandler := user.NewHandler(userService)

	jwtService := jwt.NewJWTService()
	authHandler := auth.NewAuthHandler(userService, jwtService)

	// === Auth routes ===
	authGroup := api.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
	}

	// === User routes ===
	userGroup := api.Group("/users")
	{
		userGroup.POST("/", userHandler.CreateUser)
		userGroup.GET("/:id", userHandler.GetUserByID)
		userGroup.DELETE("/:id", userHandler.DeleteUser)
	}

	// TODO: Add /pharmacy, /cashier, /subscription routes next
}
