package service

import (
	"armiya/equipment-service/genprotos"
	"armiya/equipment-service/internal/storage"
	"context"
	"log"
	"os"
)

type (
	AuthService struct {
		genprotos.UnimplementedAuthServiceServer
		authService storage.Auth
		logger      *log.Logger
	}
)

func New(service storage.Auth) *AuthService {
	return &AuthService{
		authService: service,
		logger:      log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (s *AuthService) Register(ctx context.Context, req *genprotos.RegisterRequest) (*genprotos.RegisterResponse, error) {
	s.logger.Println("Register request")
	return s.authService.Register(ctx, req)
}

func (s *AuthService) Login(ctx context.Context, req *genprotos.LoginRequest) (*genprotos.LoginResponse, error) {
	s.logger.Println("Login request")
	return s.authService.Login(ctx, req)
}

func (s *AuthService) ShowProfile(ctx context.Context, req *genprotos.ShowProfileRequest) (*genprotos.ShowProfileResponse, error) {
	s.logger.Println("Show Profile request")
	return s.authService.ShowProfile(ctx, req)
}

func (s *AuthService) EditProfile(ctx context.Context, req *genprotos.EditProfileRequest) (*genprotos.EditProfileResponse, error) {
	s.logger.Println("Edit Profile request")
	return s.authService.EditProfile(ctx, req)
}

func (s *AuthService) EditUserType(ctx context.Context, req *genprotos.EditUserTypeRequest) (*genprotos.EditUserTypeResponse, error) {
	s.logger.Println("Edit User Type request")
	return s.authService.EditUserType(ctx, req)
}

func (s *AuthService) GetAllUsers(ctx context.Context, req *genprotos.GetAllUsersRequest) (*genprotos.GetAllUsersResponse, error) {
	s.logger.Println("Get All Users request")
	return s.authService.GetAllUsers(ctx, req)
}

func (s *AuthService) DeleteUser(ctx context.Context, req *genprotos.DeleteUserRequest) (*genprotos.AuthMessage, error) {
	s.logger.Println("Delete User request")
	return s.authService.DeleteUser(ctx, req)
}

func (s *AuthService) ResetPassword(ctx context.Context, req *genprotos.ResetPasswordRequest) (*genprotos.AuthMessage, error) {
	s.logger.Println("Reset Password request")
	return s.authService.ResetPassword(ctx, req)
}
