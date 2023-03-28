package handler_grpc

import (
	"ZepterUserService/service"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "ZepterUserService/common/proto/user_service"
	log "github.com/sirupsen/logrus"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	userService *service.UserService
}

func New() (*UserHandler, error) {

	userService, err := service.New()
	if err != nil {
		log.WithFields(log.Fields{"service_name": "user-service", "method_name": "NewUserHandler"}).Error("Error creating user service.")
		return nil, err
	}
	log.WithFields(log.Fields{"service_name": "user-service", "method_name": "NewUserHandler"}).Info("Successfully created user handler.")
	return &UserHandler{
		userService: userService,
	}, nil
}

func (handler *UserHandler) GetUser(ctx context.Context, request *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	id := request.Id
	user := handler.userService.GetByID(int(id))
	if user.ID == 0 {
		err := status.Error(codes.NotFound, "User with that id does not exist.")
		log.WithFields(log.Fields{"service_name": "user-service", "method_name": "GetUser", "user_id": id}).Warn("User with that id does not exist.")
		return nil, err
	}
	userPb := mapUserDtoToProto(user)
	response := &pb.GetUserResponse{
		User: userPb,
	}

	log.WithFields(log.Fields{"service_name": "user-service", "method_name": "GetUser", "user_id": id}).Info("User successfully retrieved.")
	return response, nil
}

func (handler *UserHandler) GetUserByUsername(ctx context.Context, request *pb.GetUserByUsernameRequest) (*pb.GetUserByUsernameResponse, error) {
	username := request.Username
	user := handler.userService.GetByUsername(username)
	if user.ID == 0 {
		err := status.Error(codes.NotFound, "User with that username does not exist.")
		log.WithFields(log.Fields{"service_name": "user-service", "method_name": "GetUserByUsername", "username": username}).Warn("User with that username does not exist.")
		return nil, err
	}
	userPb := mapUserToProto(user)
	response := &pb.GetUserByUsernameResponse{
		User: userPb,
	}

	log.WithFields(log.Fields{"service_name": "user-service", "method_name": "GetUserByUsername", "username": username}).Info("User successfully retrieved.")
	return response, nil
}

func (handler *UserHandler) UpdateUser(ctx context.Context, request *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	user := mapProtoToUser(request.User)
	user.ID = GetUserID(ctx)
	if handler.userService.GetByID(int(user.ID)).ID == 0 {
		err := status.Error(codes.NotFound, "User with that id does not exist.")
		log.WithFields(log.Fields{"service_name": "user-service", "method_name": "UpdateUser", "user_id": user.ID}).Warn("User with that id does not exist..")
		return nil, err
	}
	id := handler.userService.UpdateUser(user.ID, user.Name, user.Email, user.Password, user.UserName, user.Gender, user.PhoneNumber, user.DateOfBirth, user.Biography, user.IsPrivate)
	response := &pb.UpdateUserResponse{
		Id: int64(id),
	}

	log.WithFields(log.Fields{"service_name": "user-service", "method_name": "UpdateUser", "user_id": user.ID}).Info("User successfully updated.")
	return response, nil
}

func (handler *UserHandler) SearchUsers(ctx context.Context, request *pb.SearchUsersRequest) (*pb.SearchUsersResponse, error) {
	username := request.Username
	loggedUserId := request.LoggedUserId
	var users []*pb.User
	for _, user := range handler.userService.SearchUsers(username, uint(loggedUserId)) {
		users = append(users, mapUserDtoToProto(user))
	}
	response := &pb.SearchUsersResponse{
		Users: users,
	}
	return response, nil
}

func (handler *UserHandler) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := mapProtoToUser(request.User)
	id := handler.userService.CreateUser(user.Name, user.Email, user.Password, user.UserName, user.Gender, user.PhoneNumber, user.DateOfBirth, user.Biography)
	if id == 0 {
		err := status.Error(codes.AlreadyExists, "User with same username or email already exists.")
		log.WithFields(log.Fields{"service_name": "user-service", "method_name": "CreateUser", "username": user.UserName}).Warn("User with same username or email already exists.")
		return nil, err
	}
	response := &pb.CreateUserResponse{}

	log.WithFields(log.Fields{"service_name": "user-service", "method_name": "CreateUser", "username": user.UserName}).Info("User successfully created.")
	return response, nil
}

func (handler *UserHandler) ForgotPassword(ctx context.Context, request *pb.ForgotPasswordRequest) (*pb.ForgotPasswordResponse, error) {
	i := handler.userService.ForgotPassword(request.Username)
	if i == 0 {
		err := status.Error(codes.InvalidArgument, "User with that username does not exist.")
		log.WithFields(log.Fields{"service_name": "user-service", "method_name": "ForgotPassword", "username": request.Username}).Warn("User with that username does not exist.")
		return nil, err
	}
	response := &pb.ForgotPasswordResponse{}
	log.WithFields(log.Fields{"service_name": "user-service", "method_name": "ForgotPassword", "username": request.Username}).Info("Temporary password successfully created.")
	return response, nil
}
