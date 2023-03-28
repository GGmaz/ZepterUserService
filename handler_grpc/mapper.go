package handler_grpc

import (
	"ZepterUserService/model"
	"time"

	pb "ZepterUserService/common/proto/user_service"
)

func mapUserDtoToProto(user model.User) *pb.User {
	userPb := &pb.User{
		Id:        int64(user.ID),
		Name:      user.Name,
		Username:  user.UserName,
		Email:     user.Email,
		Gender:    string(user.Gender),
		Phone:     user.PhoneNumber,
		Date:      user.DateOfBirth.Format("02-Jan-2006"),
		Biography: user.Biography,
		IsPrivate: user.IsPrivate,
	}

	return userPb
}

func mapUserToProto(user model.User) *pb.UserWithPass {
	userPb := &pb.UserWithPass{
		Id:        int64(user.ID),
		Name:      user.Name,
		Username:  user.UserName,
		Email:     user.Email,
		Gender:    string(user.Gender),
		Phone:     user.PhoneNumber,
		Date:      user.DateOfBirth.Format("02-Jan-2006"),
		Biography: user.Biography,
		Password:  user.Password,
		IsPrivate: user.IsPrivate,
	}

	return userPb
}

func mapProtoToUser(user *pb.UserWithPass) model.User {
	date, _ := time.Parse(time.RFC3339, user.Date)
	userPb := model.User{
		ID:          uint(user.Id),
		Name:        user.Name,
		UserName:    user.Username,
		Email:       user.Email,
		Gender:      model.Gender(user.Gender),
		PhoneNumber: user.Phone,
		DateOfBirth: date,
		Biography:   user.Biography,
		Password:    user.Password,
		IsPrivate:   user.IsPrivate,
	}
	return userPb
}
