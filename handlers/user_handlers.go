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
	res, err := h.userService.FindUserById(ID)

	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
			"msg":   "not found",
			"code":  404,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": res,
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

	res, err := h.userService.CreateUser(userRequest)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
			"code":  400,
			"msg":   "bad request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": map[string]interface{}{
			"ID":        res.ID,
			"Name":      res.Name,
			"Email":     res.Email,
			"CreatedAt": res.CreatedAt,
			"UpdatedAt": res.UpdatedAt,
		},
	})
}

// @Tags Users
// @Summary Get all user
// @Description Get all user
// @Accept	json
// @Produce	json
// @Success 200 {object} userSchema.APIResponseUsers "Success"
// @Failure 400 {object} userSchema.APIResponse400 "Bad Request"
// @Failure 404 {object} userSchema.APIResponse404 "Not Found"
// @Router /users [get]
func (h *userHandler) FindUsersHandler(c *gin.Context) {
	res, err := h.userService.FindUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"code":  400,
			"msg":   "bad request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
		"code": 200,
		"msg":  "success",
	})
}

// @Tags	Users
// @Summary Delete user by ID
// @Description Delete user by ID
// @Accept	json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} userSchema.APIResponseDelete200 "Success"
// @Failure 400 {object} userSchema.APIResponse400 "Bad Request"
// @Failure 404 {object} userSchema.APIResponse404 "Not Found"
// @Router /users/{id} [delete]
func (h *userHandler) DeleteUserHandler(c *gin.Context) {
	id := c.Param("id")
	ID, _ := strconv.Atoi(id)

	result, err := h.userService.DeleteUser(ID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":  404,
			"msg":   "not found",
			"error": err.Error(),
		})
		return
	}

	fmt.Println(result)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"ID":   result.ID,
	})
}

// @Tags Users
// @Summary Update user by Id
// @Description Update user by Id
// @Accept	json
// @Produce	json
// @Param id path int true "id"
// @Param body body userSchema.UserRequest true "Body"
// @Success 200 {object} userSchema.APIResponseUser "Success"
// @Failure 400 {object} userSchema.APIResponse400 "Bad Request"
// @Failure 404 {object} userSchema.APIResponse404 "Not Found"
// @Router /users/{id} [put]
func (h *userHandler) UpdateUserByIdHandler(c *gin.Context) {
	var userRequest userSchema.UserRequest

	id := c.Param("id")
	ID, _ := strconv.Atoi(id)

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

	res, err := h.userService.UpdateUser(userRequest, ID)
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
		"data": map[string]interface{}{
			"ID":        res.ID,
			"Name":      res.Name,
			"Email":     res.Email,
			"CreatedAt": res.CreatedAt,
			"UpdatedAt": res.UpdatedAt,
		},
	},
	)
}
