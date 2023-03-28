package handler_grpc

import (
	"context"
	"zepter/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	log "github.com/sirupsen/logrus"
	pb "zepter/common/proto/user_service"
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
	userPb := mapUserToProto(user)
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
	if handler.userService.GetByID(int(user.ID)).ID == 0 {
		err := status.Error(codes.NotFound, "User with that id does not exist.")
		log.WithFields(log.Fields{"service_name": "user-service", "method_name": "UpdateUser", "user_id": user.ID}).Warn("User with that id does not exist..")
		return nil, err
	}
	id := handler.userService.UpdateUser(user.ID, user.FirstName, user.LastName, user.Country, user.Password)
	response := &pb.UpdateUserResponse{
		Id: int64(id),
	}

	log.WithFields(log.Fields{"service_name": "user-service", "method_name": "UpdateUser", "user_id": user.ID}).Info("User successfully updated.")
	return response, nil
}

//TODO: paginated
func (handler *UserHandler) SearchUsers(ctx context.Context, request *pb.SearchUsersRequest) (*pb.SearchUsersResponse, error) {
	country := request.Country
	var users []*pb.User
	for _, user := range handler.userService.SearchUsers(country) {
		users = append(users, mapUserToProto(user))
	}
	response := &pb.SearchUsersResponse{
		Users: users,
	}
	return response, nil
}

func (handler *UserHandler) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := mapProtoToUser(request.User)
	id := handler.userService.CreateUser(user.FirstName, user.Email, user.Password, user.Username, user.LastName, user.Country)
	if id == 0 {
		err := status.Error(codes.AlreadyExists, "User with same username or email already exists.")
		log.WithFields(log.Fields{"service_name": "user-service", "method_name": "CreateUser", "username": user.Username}).Warn("User with same username or email already exists.")
		return nil, err
	}
	response := &pb.CreateUserResponse{
		Id: int64(id),
	}

	log.WithFields(log.Fields{"service_name": "user-service", "method_name": "CreateUser", "username": user.Username}).Info("User successfully created.")
	return response, nil
}

//TODO: Delete user
