package handlers

import (
	"fmt"
	"net/http"
	userSchema "pustaka-api/schemas"
	"pustaka-api/utils"

	"pustaka-api/utils/auth_jwt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @Tags	Authentication
// @Summary Sign In
// @Description Sign In
// @Accept	json
// @Produce json
// @Param body body userSchema.Authentication true "Body"
// @Success 200 {object} userSchema.APIResponseToken "Success"
// @Failure 400 {object} userSchema.APIResponse400 "Bad Request"
// @Failure 404 {object} userSchema.APIResponse404 "Not Found"
// @Router /auth/signin [post]
func (h *userHandler) AuthLogin(c *gin.Context) {
	var userRequest userSchema.Authentication

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

	res, err := h.userService.FindUserByEmail(userRequest.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
			"code":  404,
			"msg":   "not found",
		})
		return
	}

	password := utils.CheckPasswordHash(userRequest.Password, res.Password)
	if !password {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"msg":   "bad request",
			"error": "invalid email and password",
		})
		return
	}

	token, _ := auth_jwt.GenerateJWT(res.Email)

	h.userService.SaveUserToken(token, res.Email)

	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"msg":   "success",
		"token": token,
	},
	)
}
