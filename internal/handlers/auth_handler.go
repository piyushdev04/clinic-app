package handlers

import (
	"clinic-app/internal/auth"
	"clinic-app/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}

// @Security BearerAuth
// LoginHandler godoc
// @Summary Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body LoginRequest true "User credentials"
// @Success 200 {object} LoginResponse
// @Router /login [post]
func LoginHandler(userService *services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		user, err := userService.GetByUsername(req.Username)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		token, err := auth.GenerateJWT(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}
		c.JSON(http.StatusOK, LoginResponse{Token: token, Role: user.Role})
	}
}
