package routes

import (
	"hacktiv8-chapter-two-project-one/controller"
	"hacktiv8-chapter-two-project-one/repository"
	"hacktiv8-chapter-two-project-one/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupBookRoute(router *gin.Engine, db *gorm.DB) {
	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository)
	bookController := controller.NewBookController(bookService)

	router.POST("/books", bookController.CreateBook)
	router.GET("/books", bookController.GetAllBook)
	router.GET("/books/:book_id", bookController.GetBookByID)
	router.PUT("/books/:book_id", bookController.UpdateBook)
	router.DELETE("/books/:book_id", bookController.DeleteBook)
}
