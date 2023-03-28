package handler_grpc

import (
	"time"
	"zepter/model"

	pb "zepter/common/proto/user_service"
)

func mapUserToProto(user model.User) *pb.User {
	userPb := &pb.User{
		Id:        int64(user.ID),
		FirstName: user.FirstName,
		Username:  user.Username,
		Email:     user.Email,
		LastName:  user.LastName,
		Password:  user.Password,
		Country:   user.Country,
		CreatedAt: user.CreatedAt.Format("02-Jan-2006 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("02-Jan-2006 15:04:05"),
	}

	return userPb
}

func mapProtoToUser(user *pb.User) model.User {
	createdDate, _ := time.Parse(time.RFC3339, user.CreatedAt)
	updatedDate, _ := time.Parse(time.RFC3339, user.UpdatedAt)
	userPb := model.User{
		ID:        uint(user.Id),
		FirstName: user.FirstName,
		Username:  user.Username,
		Email:     user.Email,
		LastName:  user.LastName,
		Password:  user.Password,
		Country:   user.Country,
		CreatedAt: createdDate,
		UpdatedAt: updatedDate,
	}
	return userPb
}
