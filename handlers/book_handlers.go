package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	bookSchema "pustaka-api/schemas"
	"pustaka-api/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService services.BookService
}

func NewBookHandler(bookService services.BookService) *bookHandler {
	return &bookHandler{bookService}
}

// @Tags Books
// @Summary Get all books
// @Description Get all books
// @Accept	json
// @Produce	json
// @Success 200 {object} bookSchema.APIResponseID "Success"
// @Failure 400 {object} bookSchema.APIResponse400 "We need ID!!"
// @Failure 404 {object} bookSchema.APIResponse404 "Can not find ID"
// @Router /books [get]
func (h *bookHandler) BooksHandler(c *gin.Context) {
	result, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":  err,
			"code": 400,
			"msg":  "bad request",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
		"code": 200,
		"msg":  "success",
	})
}

// @Tags Books
// @Summary Post a book
// @Description Post a book
// @Accept	json
// @Produce	json
// @Param body body bookSchema.BookRequest true "Body"
// @Success 200 {object} bookSchema.BookRequest "Success"
// @Failure 400 {object} bookSchema.APIResponse400 "We need ID!!"
// @Failure 404 {object} bookSchema.APIResponse404 "Can not find ID"
// @Router /books [post]
func (h *bookHandler) PostBookHandler(c *gin.Context) {
	var bookRequest bookSchema.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

	errorMessages := []string{}
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
			"code":   400,
			"msg":    "bad request",
		})
		return
	}

	book, err := h.bookService.Create(bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
			"code":  400,
			"msg":   "bad request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
		"code": 200,
		"msg":  "success",
	})
}

// @Tags	Books
// @Summary Delete book by ID
// @Description Delete book by ID
// @Accept	json
// @Produce json
// @Param ID path int true "Id"
// @Success 200 {object} bookSchema.BookRequest "Success"
// @Failure 400 {object} bookSchema.APIResponse400 "We need ID!!"
// @Failure 404 {object} bookSchema.APIResponse404 "Can not find ID"
// @Router /books/{id} [delete]
func (h *bookHandler) DeleteBookHandler(c *gin.Context) {
	id := c.Param("id")
	ID, _ := strconv.Atoi(id)
	_, err := h.bookService.Delete(ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "success",
	})
}

// @Tags	Books
// @Summary Find book by ID
// @Description Find book by ID
// @Accept	json
// @Produce json
// @Param ID path int true "Id"
// @Success 200 {object} bookSchema.BookRequest "Success"
// @Failure 400 {object} bookSchema.APIResponse400 "We need ID!!"
// @Failure 404 {object} bookSchema.APIResponse404 "Can not find ID"
// @Router /books/{id} [get]
func (h *bookHandler) FindByIDBookHandler(c *gin.Context) {
	id := c.Param("id")
	ID, _ := strconv.Atoi(id)
	result, err := h.bookService.FindById(ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "success",
		"data":   result,
	})
}

// @Tags	Books
// @Summary Update the book by ID
// @Description Update the book by ID
// @Accept	json
// @Produce json
// @Param ID path int true "Id"
// @Param body body bookSchema.BookRequest true "Body"
// @Success 200 {object} bookSchema.BookRequest "Success"
// @Failure 400 {object} bookSchema.APIResponse400 "We need ID!!"
// @Failure 404 {object} bookSchema.APIResponse404 "Can not find ID"
// @Router /books/{id} [put]
func (h *bookHandler) UpdateBookHandler(c *gin.Context) {
	id := c.Param("id")
	ID, _ := strconv.Atoi(id)

	var bookRequest bookSchema.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

	errorMessages := []string{}
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
			"code":   400,
			"msg":    "bad request",
		})
		return
	}

	book, err := h.bookService.Update(bookRequest, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
			"code":  400,
			"msg":   "bad request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
		"code": 200,
		"msg":  "success",
	})
}
