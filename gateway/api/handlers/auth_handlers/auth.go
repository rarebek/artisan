package authhandlers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/armiya-gateway/api/handlers/auth"
	genprotos "github.com/ruziba3vich/armiya-gateway/genprotos"
)

type AuthHandlers struct {
	client genprotos.AuthServiceClient
	logger *log.Logger
}

func NewAuthHandlers(client genprotos.AuthServiceClient, logger *log.Logger) *AuthHandlers {
	return &AuthHandlers{
		client: client,
		logger: logger,
	}
}

// Register godoc
// @Summary Register user
// @Description This endpoint for registering user.
// @Accept json
// @Produce json
// @Param request body genprotos.RegisterRequest true "User details to register"
// @Success 200 {object} genprotos.RegisterResponse
// @Failure 400 {object} genprotos.Message
// @Failure 500 {object} genprotos.Message
// @Router /auth/register [post]
func (a *AuthHandlers) Register(ctx *gin.Context) {
	var req genprotos.RegisterRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := a.client.Register(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// Login godoc
// @Summary Login user
// @Description This endpoint for logging in user.
// @Accept json
// @Produce json
// @Param request body genprotos.LoginRequest true "User login details"
// @Success 200 {object} genprotos.LoginResponse
// @Failure 400 {object} genprotos.Message
// @Failure 500 {object} genprotos.Message
// @Router /auth/login [post]
func (a *AuthHandlers) Login(ctx *gin.Context) {
	var req genprotos.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := a.client.Login(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	manager := auth.NewTokenManager("secret_key_nodirbek")

	token, _ := manager.GenerateToken(req.Email)

	ctx.IndentedJSON(200, token)
}

// ShowProfile godoc
// @Summary Show user profile
// @Description This endpoint for showing user profile.
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} genprotos.ShowProfileResponse
// @Failure 400 {object} genprotos.Message
// @Failure 500 {object} genprotos.Message
// @Router /auth/profile/{id} [get]
func (a *AuthHandlers) ShowProfile(ctx *gin.Context) {
	var req genprotos.ShowProfileRequest
	req.Id = ctx.Param("id")

	resp, err := a.client.ShowProfile(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// EditProfile godoc
// @Summary Edit user profile
// @Description This endpoint for editing user profile.
// @Accept json
// @Produce json
// @Param request body genprotos.EditProfileRequest true "User profile details to edit"
// @Success 200 {object} genprotos.EditProfileResponse
// @Failure 400 {object} genprotos.Message
// @Failure 500 {object} genprotos.Message
// @Router /auth/profile/edit [put]
func (a *AuthHandlers) EditProfile(ctx *gin.Context) {
	var req genprotos.EditProfileRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := a.client.EditProfile(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// EditUserType godoc
// @Summary Edit user type
// @Description This endpoint for editing user type.
// @Accept json
// @Produce json
// @Param request body genprotos.EditUserTypeRequest true "User type details to edit"
// @Success 200 {object} genprotos.EditUserTypeResponse
// @Failure 400 {object} genprotos.Message
// @Failure 500 {object} genprotos.Message
// @Router /auth/usertype/edit [put]
func (a *AuthHandlers) EditUserType(ctx *gin.Context) {
	var req genprotos.EditUserTypeRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := a.client.EditUserType(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// GetAllUsers godoc
// @Summary Get all users
// @Description This endpoint for getting all users with pagination.
// @Accept json
// @Produce json
// @Param page query uint64 true "Page number"
// @Param limit query uint64 true "Number of users per page"
// @Success 200 {object} genprotos.GetAllUsersResponse
// @Failure 400 {object} genprotos.Message
// @Failure 500 {object} genprotos.Message
// @Router /auth/users [get]
func (a *AuthHandlers) GetAllUsers(ctx *gin.Context) {
	var req genprotos.GetAllUsersRequest

	page, err := strconv.ParseUint(ctx.Query("page"), 10, 64)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "Invalid page number"})
		return
	}
	limit, err := strconv.ParseUint(ctx.Query("limit"), 10, 64)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "Invalid limit"})
		return
	}

	req.Page = page
	req.Limit = limit

	resp, err := a.client.GetAllUsers(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// DeleteUser godoc
// @Summary Delete user
// @Description This endpoint for deleting user.
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} genprotos.AuthMessage
// @Failure 400 {object} genprotos.Message
// @Failure 500 {object} genprotos.Message
// @Router /auth/delete/{id} [delete]
func (a *AuthHandlers) DeleteUser(ctx *gin.Context) {
	var req genprotos.DeleteUserRequest
	req.Id = ctx.Param("id")

	resp, err := a.client.DeleteUser(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// ResetPassword godoc
// @Summary Reset password
// @Description This endpoint for resetting user password.
// @Accept json
// @Produce json
// @Param request body genprotos.ResetPasswordRequest true "User email to reset password"
// @Success 200 {object} genprotos.AuthMessage
// @Failure 400 {object} genprotos.Message
// @Failure 500 {object} genprotos.Message
// @Router /auth/reset [post]
func (a *AuthHandlers) ResetPassword(ctx *gin.Context) {
	var req genprotos.ResetPasswordRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := a.client.ResetPassword(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}
