package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"pustaka-api/config/auths"
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
// @Success 200 {object} bookSchema.APIResponseBooks "Success"
// @Failure 400 {object} bookSchema.APIResponse400 "Bad Request"
// @Failure 404 {object} bookSchema.APIResponse404 "Not Found"
// @Failure 401 {object} bookSchema.APIResponse401 "Unauthorized"
// @Router /books [get]
func (h *bookHandler) BooksHandler(c *gin.Context) {
	authorizer, _ := auths.Authorize(c)
	if !authorizer {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
			"code":    401,
		})
		return
	}

	result, err := h.bookService.FindBookAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
			"code":  400,
			"msg":   "bad request",
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
// @Success 200 {object} bookSchema.APIResponseBook "Success"
// @Failure 400 {object} bookSchema.APIResponse400 "Bad Request"
// @Failure 404 {object} bookSchema.APIResponse404 "Not Found"
// @Failure 401 {object} bookSchema.APIResponse401 "Unauthorized"
// @Router /books [post]
func (h *bookHandler) PostBookHandler(c *gin.Context) {
	authorizer, _ := auths.Authorize(c)
	if !authorizer {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
			"code":    401,
		})
		return
	}
	var bookRequest bookSchema.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

	errorMessages := []string{}
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
			"code":  400,
			"msg":   "bad request",
		})
		return
	}

	book, err := h.bookService.CreateBook(bookRequest)
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
// @Param id path int true "id"
// @Success 200 {object} bookSchema.APIResponse200 "Success"
// @Failure 400 {object} bookSchema.APIResponse400 "Bad Request"
// @Failure 404 {object} bookSchema.APIResponse404 "Not Found"
// @Failure 401 {object} bookSchema.APIResponse401 "Unauthorized"
// @Router /books/{id} [delete]
func (h *bookHandler) DeleteBookHandler(c *gin.Context) {
	authorizer, _ := auths.Authorize(c)
	if !authorizer {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
			"code":    401,
		})
		return
	}

	id := c.Param("id")
	ID, _ := strconv.Atoi(id)
	_, err := h.bookService.DeleteBook(ID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
			"code":  404,
			"msg":   "not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

// @Tags	Books
// @Summary Find book by ID
// @Description Find book by ID
// @Accept	json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} bookSchema.APIResponseBook "Success"
// @Failure 400 {object} bookSchema.APIResponse400 "Bad Request"
// @Failure 404 {object} bookSchema.APIResponse404 "Not Found"
// @Failure 401 {object} bookSchema.APIResponse401 "Unauthorized"
// @Router /books/{id} [get]
func (h *bookHandler) FindByIDBookHandler(c *gin.Context) {
	authorizer, _ := auths.Authorize(c)
	if !authorizer {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
			"code":    401,
		})
		return
	}

	id := c.Param("id")
	ID, _ := strconv.Atoi(id)
	result, err := h.bookService.FindBookById(ID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
			"msg":   "not found",
			"code":  404,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": result,
	})
}

// @Tags	Books
// @Summary Update the book by ID
// @Description Update the book by ID
// @Accept	json
// @Produce json
// @Param id path int true "id"
// @Param body body bookSchema.BookRequest true "Body"
// @Success 200 {object} bookSchema.APIResponseBook "Success"
// @Failure 400 {object} bookSchema.APIResponse400 "Bad Request"
// @Failure 404 {object} bookSchema.APIResponse404 "Not Found"
// @Failure 401 {object} bookSchema.APIResponse401 "Unauthorized"
// @Router /books/{id} [put]
func (h *bookHandler) UpdateBookHandler(c *gin.Context) {
	authorizer, _ := auths.Authorize(c)
	if !authorizer {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
			"code":    401,
		})
		return
	}

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

	book, err := h.bookService.UpdateBook(bookRequest, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
			"code":  404,
			"msg":   "not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
		"code": 200,
		"msg":  "success",
	})
}
