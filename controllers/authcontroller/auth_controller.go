package authcontroller

import (
	"fmt"
	models "moodly/Models"
	"moodly/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service *services.AuthService
}

func NewAuthController(services *services.AuthService) *AuthController {
	return &AuthController{service: services}
}
func (ac *AuthController) HandleRegister(c *gin.Context) {

	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := ac.service.CreateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created",
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}

func (ac *AuthController) HandleLogin(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("bind error:", err)

		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	token, err := ac.service.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"token":   token,
	})
}
