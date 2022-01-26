package handlers

import (
	"fmt"
	"net/http"
	userSchema "pustaka-api/schemas"
	"pustaka-api/services"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *userHandler {
	return &userHandler{userService}
}

// @Tags	Users
// @Summary Find user by ID
// @Description Find user by ID
// @Accept	json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} userSchema.APIResponseUser "Success"
// @Failure 400 {object} userSchema.APIResponse400 "Bad Request"
// @Failure 404 {object} userSchema.APIResponse404 "Not Found"
// @Router /users/{id} [get]
func (h *userHandler) FindUserByIdHandler(c *gin.Context) {
	id := c.Param("id")
	ID, _ := strconv.Atoi(id)
	result, err := h.userService.FindUserById(ID)

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
		"data": map[string]string{
			"id":         fmt.Sprint(result.ID),
			"name":       result.Name,
			"email":      result.Email,
			"created_at": fmt.Sprint(result.CreatedAt),
			"updated_at": fmt.Sprint(result.UpdatedAt),
		},
	})
}

// @Tags Users
// @Summary Create a user
// @Description Create a user
// @Accept	json
// @Produce	json
// @Param body body userSchema.UserRequest true "Body"
// @Success 200 {object} userSchema.APIResponseUser "Success"
// @Failure 400 {object} userSchema.APIResponse400 "Bad Request"
// @Failure 404 {object} userSchema.APIResponse404 "Not Found"
// @Router /users [post]
func (h *userHandler) CreateUserHandler(c *gin.Context) {
	var userRequest userSchema.UserRequest

	err := c.ShouldBindJSON(&userRequest)

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

	result, err := h.userService.CreateUser(userRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
			"code":  400,
			"msg":   "bad request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": map[string]string{
			"id":         fmt.Sprint(result.ID),
			"name":       result.Name,
			"email":      result.Email,
			"created_at": fmt.Sprint(result.CreatedAt),
			"updated_at": fmt.Sprint(result.UpdatedAt),
		},
	})
}
