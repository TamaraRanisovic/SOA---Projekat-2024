// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: users/userservice.proto

package users

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *TokenRequest) Reset() {
	*x = TokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_userservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenRequest) ProtoMessage() {}

func (x *TokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_userservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenRequest.ProtoReflect.Descriptor instead.
func (*TokenRequest) Descriptor() ([]byte, []int) {
	return file_users_userservice_proto_rawDescGZIP(), []int{0}
}

func (x *TokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UserIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserIdResponse) Reset() {
	*x = UserIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_userservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIdResponse) ProtoMessage() {}

func (x *UserIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_userservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIdResponse.ProtoReflect.Descriptor instead.
func (*UserIdResponse) Descriptor() ([]byte, []int) {
	return file_users_userservice_proto_rawDescGZIP(), []int{1}
}

func (x *UserIdResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AccountIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AccountIdRequest) Reset() {
	*x = AccountIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_userservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountIdRequest) ProtoMessage() {}

func (x *AccountIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_userservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountIdRequest.ProtoReflect.Descriptor instead.
func (*AccountIdRequest) Descriptor() ([]byte, []int) {
	return file_users_userservice_proto_rawDescGZIP(), []int{2}
}

func (x *AccountIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UsernameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *UsernameRequest) Reset() {
	*x = UsernameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_userservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsernameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsernameRequest) ProtoMessage() {}

func (x *UsernameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_userservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsernameRequest.ProtoReflect.Descriptor instead.
func (*UsernameRequest) Descriptor() ([]byte, []int) {
	return file_users_userservice_proto_rawDescGZIP(), []int{3}
}

func (x *UsernameRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type Credentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Credentials) Reset() {
	*x = Credentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_userservice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Credentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Credentials) ProtoMessage() {}

func (x *Credentials) ProtoReflect() protoreflect.Message {
	mi := &file_users_userservice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Credentials.ProtoReflect.Descriptor instead.
func (*Credentials) Descriptor() ([]byte, []int) {
	return file_users_userservice_proto_rawDescGZIP(), []int{4}
}

func (x *Credentials) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Credentials) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AccountDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username  string      `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email     string      `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Role      int32       `protobuf:"varint,4,opt,name=role,proto3" json:"role,omitempty"`
	Isblocked bool        `protobuf:"varint,5,opt,name=isblocked,proto3" json:"isblocked,omitempty"`
	User      *UserDetail `protobuf:"bytes,6,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *AccountDetailResponse) Reset() {
	*x = AccountDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_userservice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountDetailResponse) ProtoMessage() {}

func (x *AccountDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_userservice_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountDetailResponse.ProtoReflect.Descriptor instead.
func (*AccountDetailResponse) Descriptor() ([]byte, []int) {
	return file_users_userservice_proto_rawDescGZIP(), []int{5}
}

func (x *AccountDetailResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AccountDetailResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AccountDetailResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AccountDetailResponse) GetRole() int32 {
	if x != nil {
		return x.Role
	}
	return 0
}

func (x *AccountDetailResponse) GetIsblocked() bool {
	if x != nil {
		return x.Isblocked
	}
	return false
}

func (x *AccountDetailResponse) GetUser() *UserDetail {
	if x != nil {
		return x.User
	}
	return nil
}

type UserDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Surname   string `protobuf:"bytes,3,opt,name=surname,proto3" json:"surname,omitempty"`
	Picture   string `protobuf:"bytes,4,opt,name=picture,proto3" json:"picture,omitempty"`
	Biography string `protobuf:"bytes,5,opt,name=biography,proto3" json:"biography,omitempty"`
	Moto      string `protobuf:"bytes,6,opt,name=moto,proto3" json:"moto,omitempty"`
}

func (x *UserDetail) Reset() {
	*x = UserDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_userservice_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDetail) ProtoMessage() {}

func (x *UserDetail) ProtoReflect() protoreflect.Message {
	mi := &file_users_userservice_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDetail.ProtoReflect.Descriptor instead.
func (*UserDetail) Descriptor() ([]byte, []int) {
	return file_users_userservice_proto_rawDescGZIP(), []int{6}
}

func (x *UserDetail) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserDetail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserDetail) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *UserDetail) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *UserDetail) GetBiography() string {
	if x != nil {
		return x.Biography
	}
	return ""
}

func (x *UserDetail) GetMoto() string {
	if x != nil {
		return x.Moto
	}
	return ""
}

type AccountsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accounts []*AccountDetailResponse `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *AccountsResponse) Reset() {
	*x = AccountsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_userservice_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountsResponse) ProtoMessage() {}

func (x *AccountsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_userservice_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountsResponse.ProtoReflect.Descriptor instead.
func (*AccountsResponse) Descriptor() ([]byte, []int) {
	return file_users_userservice_proto_rawDescGZIP(), []int{7}
}

func (x *AccountsResponse) GetAccounts() []*AccountDetailResponse {
	if x != nil {
		return x.Accounts
	}
	return nil
}

type UserIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserIdRequest) Reset() {
	*x = UserIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_userservice_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIdRequest) ProtoMessage() {}

func (x *UserIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_userservice_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIdRequest.ProtoReflect.Descriptor instead.
func (*UserIdRequest) Descriptor() ([]byte, []int) {
	return file_users_userservice_proto_rawDescGZIP(), []int{8}
}

func (x *UserIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_users_userservice_proto protoreflect.FileDescriptor

var file_users_userservice_proto_rawDesc = []byte{
	0x0a, 0x17, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x20, 0x0a, 0x0e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x22, 0x0a, 0x10,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x2d, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x45, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xac, 0x01, 0x0a, 0x15, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x96, 0x01, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x62, 0x69, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x62, 0x69, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f,
	0x74, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x6f, 0x74, 0x6f, 0x22, 0x46,
	0x0a, 0x10, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0x1f, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xd4, 0x05, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x47, 0x75, 0x69, 0x64, 0x65, 0x12, 0x0d, 0x2e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x67, 0x75,
	0x69, 0x64, 0x65, 0x12, 0x4f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0d, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a,
	0x22, 0x12, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x54, 0x0a, 0x0c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x11, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x66, 0x0a, 0x18, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x41, 0x6e, 0x64, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x0c, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x1a, 0x16, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x79,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x56, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x62,
	0x79, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4d, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x67, 0x65,
	0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x56, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x11, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11,
	0x2f, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x12, 0x5b, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x16, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0d,
	0x5a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_users_userservice_proto_rawDescOnce sync.Once
	file_users_userservice_proto_rawDescData = file_users_userservice_proto_rawDesc
)

func file_users_userservice_proto_rawDescGZIP() []byte {
	file_users_userservice_proto_rawDescOnce.Do(func() {
		file_users_userservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_users_userservice_proto_rawDescData)
	})
	return file_users_userservice_proto_rawDescData
}

var file_users_userservice_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_users_userservice_proto_goTypes = []interface{}{
	(*TokenRequest)(nil),          // 0: TokenRequest
	(*UserIdResponse)(nil),        // 1: UserIdResponse
	(*AccountIdRequest)(nil),      // 2: AccountIdRequest
	(*UsernameRequest)(nil),       // 3: UsernameRequest
	(*Credentials)(nil),           // 4: Credentials
	(*AccountDetailResponse)(nil), // 5: AccountDetailResponse
	(*UserDetail)(nil),            // 6: UserDetail
	(*AccountsResponse)(nil),      // 7: AccountsResponse
	(*UserIdRequest)(nil),         // 8: UserIdRequest
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_users_userservice_proto_depIdxs = []int32{
	6,  // 0: AccountDetailResponse.user:type_name -> UserDetail
	5,  // 1: AccountsResponse.accounts:type_name -> AccountDetailResponse
	0,  // 2: UserService.AuthenticateGuide:input_type -> TokenRequest
	0,  // 3: UserService.GetUserByToken:input_type -> TokenRequest
	2,  // 4: UserService.BlockAccount:input_type -> AccountIdRequest
	4,  // 5: UserService.GetByUsernameAndPassword:input_type -> Credentials
	3,  // 6: UserService.GetByUsername:input_type -> UsernameRequest
	8,  // 7: UserService.GetAccount:input_type -> UserIdRequest
	9,  // 8: UserService.GetAllAccounts:input_type -> google.protobuf.Empty
	5,  // 9: UserService.CreateAccount:input_type -> AccountDetailResponse
	9,  // 10: UserService.AuthenticateGuide:output_type -> google.protobuf.Empty
	1,  // 11: UserService.GetUserByToken:output_type -> UserIdResponse
	9,  // 12: UserService.BlockAccount:output_type -> google.protobuf.Empty
	5,  // 13: UserService.GetByUsernameAndPassword:output_type -> AccountDetailResponse
	5,  // 14: UserService.GetByUsername:output_type -> AccountDetailResponse
	5,  // 15: UserService.GetAccount:output_type -> AccountDetailResponse
	7,  // 16: UserService.GetAllAccounts:output_type -> AccountsResponse
	9,  // 17: UserService.CreateAccount:output_type -> google.protobuf.Empty
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_users_userservice_proto_init() }
func file_users_userservice_proto_init() {
	if File_users_userservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_users_userservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenRequest); i {
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
		file_users_userservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserIdResponse); i {
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
		file_users_userservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountIdRequest); i {
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
		file_users_userservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsernameRequest); i {
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
		file_users_userservice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Credentials); i {
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
		file_users_userservice_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountDetailResponse); i {
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
		file_users_userservice_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDetail); i {
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
		file_users_userservice_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountsResponse); i {
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
		file_users_userservice_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserIdRequest); i {
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
			RawDescriptor: file_users_userservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_users_userservice_proto_goTypes,
		DependencyIndexes: file_users_userservice_proto_depIdxs,
		MessageInfos:      file_users_userservice_proto_msgTypes,
	}.Build()
	File_users_userservice_proto = out.File
	file_users_userservice_proto_rawDesc = nil
	file_users_userservice_proto_goTypes = nil
	file_users_userservice_proto_depIdxs = nil
}