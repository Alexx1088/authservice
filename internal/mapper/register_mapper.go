package mapper

import (
	"authservice/internal/dto"
	pb "authservice/proto/proto"
)

func ToRegisterProto(dto *dto.RegisterRequestDTO) *pb.RegisterRequest {
	return &pb.RegisterRequest{
		Name:     dto.Name,
		Surname:  dto.Surname,
		Email:    dto.Email,
		Password: dto.Password,
	}
}
