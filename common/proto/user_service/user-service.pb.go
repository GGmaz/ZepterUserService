// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: user-service.proto

package user

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ForgotPasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *ForgotPasswordRequest) Reset() {
	*x = ForgotPasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForgotPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForgotPasswordRequest) ProtoMessage() {}

func (x *ForgotPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForgotPasswordRequest.ProtoReflect.Descriptor instead.
func (*ForgotPasswordRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{0}
}

func (x *ForgotPasswordRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type ForgotPasswordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ForgotPasswordResponse) Reset() {
	*x = ForgotPasswordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForgotPasswordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForgotPasswordResponse) ProtoMessage() {}

func (x *ForgotPasswordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForgotPasswordResponse.ProtoReflect.Descriptor instead.
func (*ForgotPasswordResponse) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{1}
}

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponse.ProtoReflect.Descriptor instead.
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type GetUserByUsernameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *GetUserByUsernameRequest) Reset() {
	*x = GetUserByUsernameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByUsernameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByUsernameRequest) ProtoMessage() {}

func (x *GetUserByUsernameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByUsernameRequest.ProtoReflect.Descriptor instead.
func (*GetUserByUsernameRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserByUsernameRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type GetUserByUsernameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserByUsernameResponse) Reset() {
	*x = GetUserByUsernameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByUsernameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByUsernameResponse) ProtoMessage() {}

func (x *GetUserByUsernameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByUsernameResponse.ProtoReflect.Descriptor instead.
func (*GetUserByUsernameResponse) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserByUsernameResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UpdateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateUserRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UpdateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateUserResponse) Reset() {
	*x = UpdateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserResponse) ProtoMessage() {}

func (x *UpdateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserResponse.ProtoReflect.Descriptor instead.
func (*UpdateUserResponse) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateUserResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SearchUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Country string `protobuf:"bytes,1,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *SearchUsersRequest) Reset() {
	*x = SearchUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUsersRequest) ProtoMessage() {}

func (x *SearchUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUsersRequest.ProtoReflect.Descriptor instead.
func (*SearchUsersRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{8}
}

func (x *SearchUsersRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type SearchUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *SearchUsersResponse) Reset() {
	*x = SearchUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUsersResponse) ProtoMessage() {}

func (x *SearchUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUsersResponse.ProtoReflect.Descriptor instead.
func (*SearchUsersResponse) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{9}
}

func (x *SearchUsersResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{10}
}

func (x *CreateUserRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{11}
}

func (x *CreateUserResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Username  string `protobuf:"bytes,5,opt,name=username,proto3" json:"username,omitempty"`
	Password  string `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	Country   string `protobuf:"bytes,7,opt,name=country,proto3" json:"country,omitempty"`
	CreatedAt string `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{12}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *User) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *User) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_user_service_proto protoreflect.FileDescriptor

var file_user_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x15, 0x46, 0x6f, 0x72, 0x67,
	0x6f, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x18, 0x0a,
	0x16, 0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x31, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x36, 0x0a, 0x18,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1e, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x22, 0x33, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x24, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x12,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x37, 0x0a, 0x13,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x33, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x24, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0xf8, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0xd6, 0x04, 0x0a, 0x0b,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x6f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x55, 0x0a, 0x0a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0e, 0x1a, 0x06, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x6b, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12,
	0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x12, 0x1f, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x2f, 0x7b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x7d, 0x12, 0x55, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x06, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x6e, 0x0a, 0x0e, 0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x46, 0x6f,
	0x72, 0x67, 0x6f, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x46, 0x6f, 0x72, 0x67, 0x6f,
	0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x1a, 0x19, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x7d, 0x42, 0x0d, 0x5a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_service_proto_rawDescOnce sync.Once
	file_user_service_proto_rawDescData = file_user_service_proto_rawDesc
)

func file_user_service_proto_rawDescGZIP() []byte {
	file_user_service_proto_rawDescOnce.Do(func() {
		file_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_service_proto_rawDescData)
	})
	return file_user_service_proto_rawDescData
}

var file_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_user_service_proto_goTypes = []interface{}{
	(*ForgotPasswordRequest)(nil),     // 0: user.ForgotPasswordRequest
	(*ForgotPasswordResponse)(nil),    // 1: user.ForgotPasswordResponse
	(*GetUserRequest)(nil),            // 2: user.GetUserRequest
	(*GetUserResponse)(nil),           // 3: user.GetUserResponse
	(*GetUserByUsernameRequest)(nil),  // 4: user.GetUserByUsernameRequest
	(*GetUserByUsernameResponse)(nil), // 5: user.GetUserByUsernameResponse
	(*UpdateUserRequest)(nil),         // 6: user.UpdateUserRequest
	(*UpdateUserResponse)(nil),        // 7: user.UpdateUserResponse
	(*SearchUsersRequest)(nil),        // 8: user.SearchUsersRequest
	(*SearchUsersResponse)(nil),       // 9: user.SearchUsersResponse
	(*CreateUserRequest)(nil),         // 10: user.CreateUserRequest
	(*CreateUserResponse)(nil),        // 11: user.CreateUserResponse
	(*User)(nil),                      // 12: user.User
}
var file_user_service_proto_depIdxs = []int32{
	12, // 0: user.GetUserResponse.user:type_name -> user.User
	12, // 1: user.GetUserByUsernameResponse.user:type_name -> user.User
	12, // 2: user.UpdateUserRequest.user:type_name -> user.User
	12, // 3: user.SearchUsersResponse.users:type_name -> user.User
	12, // 4: user.CreateUserRequest.user:type_name -> user.User
	2,  // 5: user.UserService.GetUser:input_type -> user.GetUserRequest
	4,  // 6: user.UserService.GetUserByUsername:input_type -> user.GetUserByUsernameRequest
	6,  // 7: user.UserService.UpdateUser:input_type -> user.UpdateUserRequest
	8,  // 8: user.UserService.SearchUsers:input_type -> user.SearchUsersRequest
	10, // 9: user.UserService.CreateUser:input_type -> user.CreateUserRequest
	0,  // 10: user.UserService.ForgotPassword:input_type -> user.ForgotPasswordRequest
	3,  // 11: user.UserService.GetUser:output_type -> user.GetUserResponse
	5,  // 12: user.UserService.GetUserByUsername:output_type -> user.GetUserByUsernameResponse
	7,  // 13: user.UserService.UpdateUser:output_type -> user.UpdateUserResponse
	9,  // 14: user.UserService.SearchUsers:output_type -> user.SearchUsersResponse
	11, // 15: user.UserService.CreateUser:output_type -> user.CreateUserResponse
	1,  // 16: user.UserService.ForgotPassword:output_type -> user.ForgotPasswordResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_user_service_proto_init() }
func file_user_service_proto_init() {
	if File_user_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForgotPasswordRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForgotPasswordResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByUsernameRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByUsernameResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchUsersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchUsersResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_service_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_service_proto_goTypes,
		DependencyIndexes: file_user_service_proto_depIdxs,
		MessageInfos:      file_user_service_proto_msgTypes,
	}.Build()
	File_user_service_proto = out.File
	file_user_service_proto_rawDesc = nil
	file_user_service_proto_goTypes = nil
	file_user_service_proto_depIdxs = nil
}
