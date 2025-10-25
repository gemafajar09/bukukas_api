package handler

import (
	"net/http"

	"go-project/internal/delivery/http/response"
	"go-project/internal/delivery/http/validator"
	"go-project/internal/domain"
	"go-project/internal/usecase"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type authHandler struct {
	uc usecase.UserUsecase
}

func NewAuthHandler(uc usecase.UserUsecase) *authHandler {
	return &authHandler{uc: uc}
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Register godoc
// @Summary Daftarkan pengguna baru
// @Description Buat akun pengguna baru
// @Tags auth
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "Register request"
// @Router /auth/register [post]
func (h *authHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Failed to register user",
			Errors:  validator.PesanError(err),
		})
		return
	}

	user := domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	createdUser, err := h.uc.Register(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Failed to register user",
			Errors:  validator.PesanError(err),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":    createdUser.ID,
		"name":  createdUser.Name,
		"email": createdUser.Email,
	})
}

// Login godoc
// @Summary Login user
// @Description Otentikasi pengguna dan dapatkan token JWT
// @Tags auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login request"
// @Router /auth/login [post]
func (h *authHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Failed to login user",
			Errors:  validator.PesanError(err),
		})
		return
	}

	token, err := h.uc.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse{
			Success: false,
			Message: "Invalid email or password",
			Errors:  validator.PesanError(err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
