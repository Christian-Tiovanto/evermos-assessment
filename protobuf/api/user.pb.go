// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: protobuf/api/user.proto

package commerce_api

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegisterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Phone         string                 `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Password      string                 `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	mi := &file_protobuf_api_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_api_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_api_user_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Phone         string                 `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_protobuf_api_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_api_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_api_user_proto_rawDescGZIP(), []int{1}
}

func (x *LoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_protobuf_api_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_api_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_api_user_proto_rawDescGZIP(), []int{2}
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Phone         string                 `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_protobuf_api_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_api_user_proto_msgTypes[3]
	if x != nil {
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
	return file_protobuf_api_user_proto_rawDescGZIP(), []int{3}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type GetUsersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Users         []*User                `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUsersResponse) Reset() {
	*x = GetUsersResponse{}
	mi := &file_protobuf_api_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersResponse) ProtoMessage() {}

func (x *GetUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_api_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersResponse.ProtoReflect.Descriptor instead.
func (*GetUsersResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_api_user_proto_rawDescGZIP(), []int{4}
}

func (x *GetUsersResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type UpdateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone         string                 `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	mi := &file_protobuf_api_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_api_user_proto_msgTypes[5]
	if x != nil {
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
	return file_protobuf_api_user_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateUserRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UpdateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateUserRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

var File_protobuf_api_user_proto protoreflect.FileDescriptor

var file_protobuf_api_user_proto_rawDesc = string([]byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x6d, 0x75, 0x67, 0x68, 0x69,
	0x65, 0x61, 0x6d, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x72, 0x6d, 0x6f, 0x73, 0x61, 0x73, 0x73, 0x65,
	0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x6d, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x5b, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1f, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x25, 0x0a, 0x0d,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x42, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x4e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6d, 0x75, 0x67,
	0x68, 0x69, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x72, 0x6d, 0x6f, 0x73, 0x61, 0x73,
	0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x6b, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x32, 0x8c, 0x04, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x82, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x2f, 0x2e, 0x6d, 0x75, 0x67, 0x68, 0x69, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x65, 0x76,
	0x65, 0x72, 0x6d, 0x6f, 0x73, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x6d, 0x75, 0x67, 0x68, 0x69, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x65,
	0x76, 0x65, 0x72, 0x6d, 0x6f, 0x73, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x80, 0x01, 0x0a, 0x05, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x2c, 0x2e, 0x6d, 0x75, 0x67, 0x68, 0x69, 0x65, 0x61, 0x6d, 0x73, 0x2e,
	0x65, 0x76, 0x65, 0x72, 0x6d, 0x6f, 0x73, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2d, 0x2e, 0x6d, 0x75, 0x67, 0x68, 0x69, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x65, 0x76,
	0x65, 0x72, 0x6d, 0x6f, 0x73, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x67, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x30, 0x2e, 0x6d, 0x75, 0x67, 0x68, 0x69, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x65, 0x76, 0x65,
	0x72, 0x6d, 0x6f, 0x73, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x8b, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x31, 0x2e, 0x6d, 0x75, 0x67, 0x68, 0x69, 0x65, 0x61, 0x6d, 0x73,
	0x2e, 0x65, 0x76, 0x65, 0x72, 0x6d, 0x6f, 0x73, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x6d, 0x75, 0x67, 0x68, 0x69, 0x65,
	0x61, 0x6d, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x72, 0x6d, 0x6f, 0x73, 0x61, 0x73, 0x73, 0x65, 0x73,
	0x73, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13,
	0x3a, 0x01, 0x2a, 0x32, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6d, 0x75, 0x67, 0x68, 0x69, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x72,
	0x6d, 0x6f, 0x73, 0x2d, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x72, 0x63, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_protobuf_api_user_proto_rawDescOnce sync.Once
	file_protobuf_api_user_proto_rawDescData []byte
)

func file_protobuf_api_user_proto_rawDescGZIP() []byte {
	file_protobuf_api_user_proto_rawDescOnce.Do(func() {
		file_protobuf_api_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_protobuf_api_user_proto_rawDesc), len(file_protobuf_api_user_proto_rawDesc)))
	})
	return file_protobuf_api_user_proto_rawDescData
}

var file_protobuf_api_user_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protobuf_api_user_proto_goTypes = []any{
	(*RegisterRequest)(nil),   // 0: mughieams.evermosassessment.v1.RegisterRequest
	(*LoginRequest)(nil),      // 1: mughieams.evermosassessment.v1.LoginRequest
	(*LoginResponse)(nil),     // 2: mughieams.evermosassessment.v1.LoginResponse
	(*User)(nil),              // 3: mughieams.evermosassessment.v1.User
	(*GetUsersResponse)(nil),  // 4: mughieams.evermosassessment.v1.GetUsersResponse
	(*UpdateUserRequest)(nil), // 5: mughieams.evermosassessment.v1.UpdateUserRequest
	(*emptypb.Empty)(nil),     // 6: google.protobuf.Empty
	(*MessageResponse)(nil),   // 7: mughieams.evermosassessment.v1.MessageResponse
}
var file_protobuf_api_user_proto_depIdxs = []int32{
	3, // 0: mughieams.evermosassessment.v1.GetUsersResponse.users:type_name -> mughieams.evermosassessment.v1.User
	0, // 1: mughieams.evermosassessment.v1.UserService.Register:input_type -> mughieams.evermosassessment.v1.RegisterRequest
	1, // 2: mughieams.evermosassessment.v1.UserService.Login:input_type -> mughieams.evermosassessment.v1.LoginRequest
	6, // 3: mughieams.evermosassessment.v1.UserService.GetUsers:input_type -> google.protobuf.Empty
	5, // 4: mughieams.evermosassessment.v1.UserService.UpdateUser:input_type -> mughieams.evermosassessment.v1.UpdateUserRequest
	7, // 5: mughieams.evermosassessment.v1.UserService.Register:output_type -> mughieams.evermosassessment.v1.MessageResponse
	2, // 6: mughieams.evermosassessment.v1.UserService.Login:output_type -> mughieams.evermosassessment.v1.LoginResponse
	4, // 7: mughieams.evermosassessment.v1.UserService.GetUsers:output_type -> mughieams.evermosassessment.v1.GetUsersResponse
	7, // 8: mughieams.evermosassessment.v1.UserService.UpdateUser:output_type -> mughieams.evermosassessment.v1.MessageResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protobuf_api_user_proto_init() }
func file_protobuf_api_user_proto_init() {
	if File_protobuf_api_user_proto != nil {
		return
	}
	file_protobuf_api_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_protobuf_api_user_proto_rawDesc), len(file_protobuf_api_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_api_user_proto_goTypes,
		DependencyIndexes: file_protobuf_api_user_proto_depIdxs,
		MessageInfos:      file_protobuf_api_user_proto_msgTypes,
	}.Build()
	File_protobuf_api_user_proto = out.File
	file_protobuf_api_user_proto_goTypes = nil
	file_protobuf_api_user_proto_depIdxs = nil
}
