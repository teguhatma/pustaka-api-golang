package main

import (
	docs "pustaka-api/docs"
	"pustaka-api/handlers"
	"pustaka-api/repository"
	"pustaka-api/services"
	"pustaka-api/utils"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.New()

	// Database Connection
	db := utils.DatabaseConn()

	// Migration
	// utils.DatabaseMigration(db)

	rGlobal := repository.NewRepository(db)

	// Book instance
	bookService := services.BookNewService(rGlobal)
	bookHandler := handlers.NewBookHandler(bookService)

	// User instance
	userService := services.UserNewService(rGlobal)
	userHandler := handlers.NewUserHandler(userService)

	// Documenting Swagger
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	v1 := r.Group("api/v1")
	{
		grbook := v1.Group("/books")
		{
			grbook.GET("", bookHandler.BooksHandler)
			grbook.DELETE("/:id", bookHandler.DeleteBookHandler)
			grbook.GET("/:id", bookHandler.FindByIDBookHandler)
			grbook.PUT("/:id", bookHandler.UpdateBookHandler)
			grbook.POST("", bookHandler.PostBookHandler)
		}
		gruser := v1.Group("/users")
		{
			gruser.GET("/:id", userHandler.FindUserByIdHandler)
			gruser.POST("", userHandler.CreateUserHandler)
			gruser.GET("", userHandler.FindUsersHandler)
			gruser.DELETE(":id", userHandler.DeleteUserHandler)
			gruser.PUT(":id", userHandler.UpdateUserByIdHandler)
		}
		grauth := v1.Group("/auth")
		{
			grauth.POST("/signin", userHandler.AuthLogin)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":5000")
}
