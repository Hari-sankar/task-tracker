package controllers

import (
	"net/http"
	"task-tracker/api/schemas"
	"task-tracker/api/usecases"
	"task-tracker/api/utils"

	"github.com/gin-gonic/gin"
)

// AuthController handles authentication-related HTTP requests
type AuthController struct {
	authUseCase usecases.AuthUseCase
}

// / NewAuthController creates a new AuthController instance
func NewAuthController(authUseCase usecases.AuthUseCase) *AuthController {
	return &AuthController{authUseCase: authUseCase}
}

// SignIn godoc
// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body schemas.SignInRequest true "Sign in credentials"
// @Success 200 {object} schemas.SignInResponse
// @Failure 401 {object} utils.ErrorStruct "Invalid credentials"
// @Failure 500 {object} utils.ErrorStruct "Internal server error"
// @Router /auth/signin [post]
func (ctrl *AuthController) SignIn(c *gin.Context) {
	ctx := c.Request.Context()

	// Parse request body
	var req schemas.SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(schemas.ErrInvalidInput)
		return
	}
	// Call the SignIn method of the authUseCase
	response, err := ctrl.authUseCase.SignIn(ctx, req.Username, req.Password)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// SignUp godoc
// @Summary Register a new user
// @Description Register a new user with username, password and type
// @Tags auth
// @Accept json
// @Produce json
// @Param request body schemas.SignUpRequest true "Sign up request"
// @Success 200 {object} schemas.SignUpResponse
// @Failure 400 {object} utils.ErrorStruct
// @Failure 409 {object} utils.ErrorStruct "Username already exists"
// @Failure 503 {object} utils.ErrorStruct "Database connection error"
// @Router /auth/signup [post]
// funtion to handle the user signup
func (ctrl *AuthController) SignUp(c *gin.Context) {
	ctx := c.Request.Context()

	// Parse request body
	var err error
	var response *schemas.SignUpResponse
	var req schemas.SignUpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(utils.NewErrorStruct(400, err.Error()))
		return
	}

	// Call the SignUp method of the authUseCase
	if response, err = ctrl.authUseCase.SignUp(ctx, &req); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, response)
}
