package mapper

import (
	"authservice/internal/dto"
	pb "authservice/proto/proto"
)

func ToAuthResponseDTO(protoResp *pb.AuthResponse) *dto.AuthResponseDTO {
	return &dto.AuthResponseDTO{
		Token:   protoResp.Token,
		UserID:  protoResp.User.UserId,
		Name:    protoResp.User.Name,
		Surname: protoResp.User.Surname,
		Email:   protoResp.User.Email,
	}
}
