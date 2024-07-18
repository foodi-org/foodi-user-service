// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.11.4
// source: foodi-account-service.proto

package foodi_user_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 用户类型枚举
type UserCoup int32

const (
	UserCoup_BUYER    UserCoup = 0 // 普通用户
	UserCoup_MERCHANT UserCoup = 1 // 商家用户
	UserCoup_RIDER    UserCoup = 2 // 配送骑手
)

// Enum value maps for UserCoup.
var (
	UserCoup_name = map[int32]string{
		0: "BUYER",
		1: "MERCHANT",
		2: "RIDER",
	}
	UserCoup_value = map[string]int32{
		"BUYER":    0,
		"MERCHANT": 1,
		"RIDER":    2,
	}
)

func (x UserCoup) Enum() *UserCoup {
	p := new(UserCoup)
	*p = x
	return p
}

func (x UserCoup) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserCoup) Descriptor() protoreflect.EnumDescriptor {
	return file_foodi_account_service_proto_enumTypes[0].Descriptor()
}

func (UserCoup) Type() protoreflect.EnumType {
	return &file_foodi_account_service_proto_enumTypes[0]
}

func (x UserCoup) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserCoup.Descriptor instead.
func (UserCoup) EnumDescriptor() ([]byte, []int) {
	return file_foodi_account_service_proto_rawDescGZIP(), []int{0}
}

// 用户注册与登录类型枚举
type RegisterCoup int32

const (
	RegisterCoup_Phone    RegisterCoup = 0 // 手机号一键登录注册
	RegisterCoup_Code     RegisterCoup = 1 // 手机验证码登录注册
	RegisterCoup_Password RegisterCoup = 3 // 密码登录注册
)

// Enum value maps for RegisterCoup.
var (
	RegisterCoup_name = map[int32]string{
		0: "Phone",
		1: "Code",
		3: "Password",
	}
	RegisterCoup_value = map[string]int32{
		"Phone":    0,
		"Code":     1,
		"Password": 3,
	}
)

func (x RegisterCoup) Enum() *RegisterCoup {
	p := new(RegisterCoup)
	*p = x
	return p
}

func (x RegisterCoup) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RegisterCoup) Descriptor() protoreflect.EnumDescriptor {
	return file_foodi_account_service_proto_enumTypes[1].Descriptor()
}

func (RegisterCoup) Type() protoreflect.EnumType {
	return &file_foodi_account_service_proto_enumTypes[1]
}

func (x RegisterCoup) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RegisterCoup.Descriptor instead.
func (RegisterCoup) EnumDescriptor() ([]byte, []int) {
	return file_foodi_account_service_proto_rawDescGZIP(), []int{1}
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type         UserCoup     `protobuf:"varint,1,opt,name=type,proto3,enum=foodiUserService.UserCoup" json:"type,omitempty"`
	RegisterType RegisterCoup `protobuf:"varint,2,opt,name=registerType,proto3,enum=foodiUserService.RegisterCoup" json:"registerType,omitempty"`
	Phone        int64        `protobuf:"varint,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Code         string       `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	Password     string       `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	Length       int64        `protobuf:"varint,6,opt,name=length,proto3" json:"length,omitempty"` // 密码登录时的补位长度
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foodi_account_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foodi_account_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_foodi_account_service_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetType() UserCoup {
	if x != nil {
		return x.Type
	}
	return UserCoup_BUYER
}

func (x *RegisterRequest) GetRegisterType() RegisterCoup {
	if x != nil {
		return x.RegisterType
	}
	return RegisterCoup_Phone
}

func (x *RegisterRequest) GetPhone() int64 {
	if x != nil {
		return x.Phone
	}
	return 0
}

func (x *RegisterRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterRequest) GetLength() int64 {
	if x != nil {
		return x.Length
	}
	return 0
}

type RegisterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok  bool  `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Uid int64 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *RegisterReply) Reset() {
	*x = RegisterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foodi_account_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReply) ProtoMessage() {}

func (x *RegisterReply) ProtoReflect() protoreflect.Message {
	mi := &file_foodi_account_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReply.ProtoReflect.Descriptor instead.
func (*RegisterReply) Descriptor() ([]byte, []int) {
	return file_foodi_account_service_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterReply) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *RegisterReply) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      UserCoup     `protobuf:"varint,1,opt,name=type,proto3,enum=foodiUserService.UserCoup" json:"type,omitempty"`
	LoginType RegisterCoup `protobuf:"varint,2,opt,name=loginType,proto3,enum=foodiUserService.RegisterCoup" json:"loginType,omitempty"`
	Phone     int64        `protobuf:"varint,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Code      string       `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	Password  string       `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	Length    int64        `protobuf:"varint,6,opt,name=length,proto3" json:"length,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foodi_account_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foodi_account_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_foodi_account_service_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetType() UserCoup {
	if x != nil {
		return x.Type
	}
	return UserCoup_BUYER
}

func (x *LoginRequest) GetLoginType() RegisterCoup {
	if x != nil {
		return x.LoginType
	}
	return RegisterCoup_Phone
}

func (x *LoginRequest) GetPhone() int64 {
	if x != nil {
		return x.Phone
	}
	return 0
}

func (x *LoginRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginRequest) GetLength() int64 {
	if x != nil {
		return x.Length
	}
	return 0
}

type LoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Uid   int64  `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *LoginReply) Reset() {
	*x = LoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foodi_account_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReply) ProtoMessage() {}

func (x *LoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_foodi_account_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReply.ProtoReflect.Descriptor instead.
func (*LoginReply) Descriptor() ([]byte, []int) {
	return file_foodi_account_service_proto_rawDescGZIP(), []int{3}
}

func (x *LoginReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginReply) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type PhoneExistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone int64 `protobuf:"varint,1,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *PhoneExistRequest) Reset() {
	*x = PhoneExistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foodi_account_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhoneExistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhoneExistRequest) ProtoMessage() {}

func (x *PhoneExistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foodi_account_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhoneExistRequest.ProtoReflect.Descriptor instead.
func (*PhoneExistRequest) Descriptor() ([]byte, []int) {
	return file_foodi_account_service_proto_rawDescGZIP(), []int{4}
}

func (x *PhoneExistRequest) GetPhone() int64 {
	if x != nil {
		return x.Phone
	}
	return 0
}

type PhoneExistReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exist bool `protobuf:"varint,1,opt,name=exist,proto3" json:"exist,omitempty"`
}

func (x *PhoneExistReply) Reset() {
	*x = PhoneExistReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foodi_account_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhoneExistReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhoneExistReply) ProtoMessage() {}

func (x *PhoneExistReply) ProtoReflect() protoreflect.Message {
	mi := &file_foodi_account_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhoneExistReply.ProtoReflect.Descriptor instead.
func (*PhoneExistReply) Descriptor() ([]byte, []int) {
	return file_foodi_account_service_proto_rawDescGZIP(), []int{5}
}

func (x *PhoneExistReply) GetExist() bool {
	if x != nil {
		return x.Exist
	}
	return false
}

var File_foodi_account_service_proto protoreflect.FileDescriptor

var file_foodi_account_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x66, 0x6f, 0x6f, 0x64, 0x69, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x66,
	0x6f, 0x6f, 0x64, 0x69, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22,
	0xe3, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1a, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x69, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x70, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x42, 0x0a, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x66, 0x6f, 0x6f, 0x64,
	0x69, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x70, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x31, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0xda, 0x01, 0x0a, 0x0c, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x69, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43,
	0x6f, 0x75, 0x70, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x66,
	0x6f, 0x6f, 0x64, 0x69, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x70, 0x52, 0x09, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x34, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x29, 0x0a, 0x11, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x27, 0x0a, 0x0f, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x45,
	0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x78, 0x69, 0x73, 0x74, 0x2a,
	0x2e, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x70, 0x12, 0x09, 0x0a, 0x05, 0x42,
	0x55, 0x59, 0x45, 0x52, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x45, 0x52, 0x43, 0x48, 0x41,
	0x4e, 0x54, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x49, 0x44, 0x45, 0x52, 0x10, 0x02, 0x2a,
	0x31, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x70, 0x12,
	0x09, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x6f,
	0x64, 0x65, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x10, 0x03, 0x32, 0xf6, 0x01, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x4e,
	0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x66, 0x6f, 0x6f,
	0x64, 0x69, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x66, 0x6f, 0x6f, 0x64, 0x69, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x45,
	0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1e, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x69, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x69, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x54, 0x0a, 0x0a, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x12, 0x23, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x69, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x45, 0x78, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x69,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x29, 0x5a, 0x27, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6f, 0x6f, 0x64, 0x69, 0x2d,
	0x6f, 0x72, 0x67, 0x2f, 0x66, 0x6f, 0x6f, 0x64, 0x69, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_foodi_account_service_proto_rawDescOnce sync.Once
	file_foodi_account_service_proto_rawDescData = file_foodi_account_service_proto_rawDesc
)

func file_foodi_account_service_proto_rawDescGZIP() []byte {
	file_foodi_account_service_proto_rawDescOnce.Do(func() {
		file_foodi_account_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_foodi_account_service_proto_rawDescData)
	})
	return file_foodi_account_service_proto_rawDescData
}

var file_foodi_account_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_foodi_account_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_foodi_account_service_proto_goTypes = []interface{}{
	(UserCoup)(0),             // 0: foodiUserService.UserCoup
	(RegisterCoup)(0),         // 1: foodiUserService.RegisterCoup
	(*RegisterRequest)(nil),   // 2: foodiUserService.RegisterRequest
	(*RegisterReply)(nil),     // 3: foodiUserService.RegisterReply
	(*LoginRequest)(nil),      // 4: foodiUserService.LoginRequest
	(*LoginReply)(nil),        // 5: foodiUserService.LoginReply
	(*PhoneExistRequest)(nil), // 6: foodiUserService.PhoneExistRequest
	(*PhoneExistReply)(nil),   // 7: foodiUserService.PhoneExistReply
}
var file_foodi_account_service_proto_depIdxs = []int32{
	0, // 0: foodiUserService.RegisterRequest.type:type_name -> foodiUserService.UserCoup
	1, // 1: foodiUserService.RegisterRequest.registerType:type_name -> foodiUserService.RegisterCoup
	0, // 2: foodiUserService.LoginRequest.type:type_name -> foodiUserService.UserCoup
	1, // 3: foodiUserService.LoginRequest.loginType:type_name -> foodiUserService.RegisterCoup
	2, // 4: foodiUserService.Account.Register:input_type -> foodiUserService.RegisterRequest
	4, // 5: foodiUserService.Account.Login:input_type -> foodiUserService.LoginRequest
	6, // 6: foodiUserService.Account.PhoneExist:input_type -> foodiUserService.PhoneExistRequest
	3, // 7: foodiUserService.Account.Register:output_type -> foodiUserService.RegisterReply
	5, // 8: foodiUserService.Account.Login:output_type -> foodiUserService.LoginReply
	7, // 9: foodiUserService.Account.PhoneExist:output_type -> foodiUserService.PhoneExistReply
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_foodi_account_service_proto_init() }
func file_foodi_account_service_proto_init() {
	if File_foodi_account_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_foodi_account_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_foodi_account_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReply); i {
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
		file_foodi_account_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_foodi_account_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReply); i {
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
		file_foodi_account_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhoneExistRequest); i {
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
		file_foodi_account_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhoneExistReply); i {
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
			RawDescriptor: file_foodi_account_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_foodi_account_service_proto_goTypes,
		DependencyIndexes: file_foodi_account_service_proto_depIdxs,
		EnumInfos:         file_foodi_account_service_proto_enumTypes,
		MessageInfos:      file_foodi_account_service_proto_msgTypes,
	}.Build()
	File_foodi_account_service_proto = out.File
	file_foodi_account_service_proto_rawDesc = nil
	file_foodi_account_service_proto_goTypes = nil
	file_foodi_account_service_proto_depIdxs = nil
}