package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/Alexx1088/authservice/internal/dto"
	pb "github.com/Alexx1088/authservice/proto"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// AuthServiceServer implements AuthServiceServer
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

	// 2. Convert RegisterRequest to DTO for internal processing
	registerDTO := dto.RegisterRequestDTO{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Name:     req.GetName(),
		Surname:  req.GetSurname(),
	}

	// 3. Validate DTO
	validate := validator.New()
	err := validate.Struct(registerDTO)
	if err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	// 3. Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	// 4. Save the user to the database
	userID := uuid.New().String()

	// 5. Generate a token (here we mock it)

	// 6. Map DTO back to Protobuf AuthResponse

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
