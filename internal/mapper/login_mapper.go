package mapper

import (
	"authservice/internal/dto"
	pb "authservice/proto/proto"
)

func ToLoginProto(dto *dto.LoginRequestDTO) *pb.LoginRequest {

	return &pb.LoginRequest{
		Email:    dto.Email,
		Password: dto.Password,
	}
}
