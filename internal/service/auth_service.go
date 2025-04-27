package service

import (
	pb "authservice/proto/proto"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceServer struct {
	pb.UnimplementedAuthServiceServer
}

func NewAuthServiceServer() *AuthServiceServer {
	return &AuthServiceServer{}
}

// Register handles user registration
func (s *AuthServiceServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	// 1. Validate incoming data
	if req.GetEmail() == "" || req.GetPassword() == "" {
		return nil, errors.New("email or password are required")
	}
	// 2. Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	// 3. Save the user to the database
	userID := uuid.New().String()

	response := &pb.AuthResponse{
		Token: token,
		User: &pb.User{
			UserId:  createdUser.ID,
			Name:    createdUser.Name,
			Surname: createdUser.Surname,
			Email:   createdUser.Email,
		},
	}
	return response, nil
}
