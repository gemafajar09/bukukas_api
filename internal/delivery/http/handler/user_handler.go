package handler

import (
	"go-project/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	uc usecase.UserUsecase
}

func NewUserHandler(uc usecase.UserUsecase) *UserHandler {
	return &UserHandler{uc: uc}
}

// GetProfile godoc
// @Summary Get user profile
// @Description Ambil data user (requires JWT token)
// @Tags users
// @Security BearerAuth
// @Produce json
// @Router /api/profile [get]
func (h *UserHandler) GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User id not found"})
		return
	}
	id := userID.(uint)
	user, err := h.uc.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// GetUsers godoc
// @Summary Ambil semua user
// @Description Ambil semua pengguna dari database
// @Tags users
// @Security BearerAuth
// @Produce json
// @Router /api/get-users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.uc.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
