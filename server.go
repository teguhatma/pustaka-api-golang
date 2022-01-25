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

	db := utils.DatabaseConn()

	bookRepository := repository.NewRepository(db)
	bookService := services.NewService(bookRepository)
	bookHandler := handlers.NewBookHandler(bookService)

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
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":5000")
}
