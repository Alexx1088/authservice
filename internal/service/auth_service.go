package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/Alexx1088/authservice/internal/dto"
	"github.com/Alexx1088/authservice/internal/model"
	"github.com/Alexx1088/authservice/internal/repository"
	"github.com/Alexx1088/authservice/pkg/jwtutil"
	pb "github.com/Alexx1088/authservice/proto"
	userpb "github.com/Alexx1088/userservice/proto"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServiceServer struct {
	pb.UnimplementedAuthServiceServer
	Repo        repository.AuthRepository
	UserService userpb.UserServiceClient
}

func NewAuthServiceServer(repo repository.AuthRepository, userService userpb.UserServiceClient) *AuthServiceServer {
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

	// 4. Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// 5. Save the user to the database
	userID := uuid.New().String()
	user := &model.User{
		ID:             userID,
		Email:          req.GetEmail(),
		HashedPassword: string(hashedPassword),
	}

	if err := s.Repo.Create(ctx, user); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}

	// 6. Create user in UserService (if necessary)
	_, err = s.UserService.CreateUser(ctx, &userpb.CreateUserRequest{
		UserId:  userID,
		Name:    req.GetName(),
		Surname: req.GetSurname(),
		Email:   req.GetEmail(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create user in UserService: %w", err)
	}

	// 7. Generate a token
	token, err := jwtutil.GenerateToken(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	// 8. Map DTO back to Protobuf AuthResponse
	response := &pb.AuthResponse{
		Token: token,
		User: &pb.User{
			UserId:  user.ID,
			Name:    req.GetName(),
			Surname: req.GetSurname(),
			Email:   user.Email,
		},
	}
	return response, nil
}
