// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: foodi-user-service.proto

package foodi_user_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foodi_user_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foodi_user_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_foodi_user_service_proto_rawDescGZIP(), []int{0}
}

func (x *UserRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type UserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	NikeName string `protobuf:"bytes,2,opt,name=NikeName,proto3" json:"NikeName,omitempty"`
	Image    string `protobuf:"bytes,3,opt,name=Image,proto3" json:"Image,omitempty"`
	LV       int32  `protobuf:"varint,4,opt,name=LV,proto3" json:"LV,omitempty"`
}

func (x *UserReply) Reset() {
	*x = UserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foodi_user_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserReply) ProtoMessage() {}

func (x *UserReply) ProtoReflect() protoreflect.Message {
	mi := &file_foodi_user_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserReply.ProtoReflect.Descriptor instead.
func (*UserReply) Descriptor() ([]byte, []int) {
	return file_foodi_user_service_proto_rawDescGZIP(), []int{1}
}

func (x *UserReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserReply) GetNikeName() string {
	if x != nil {
		return x.NikeName
	}
	return ""
}

func (x *UserReply) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *UserReply) GetLV() int32 {
	if x != nil {
		return x.LV
	}
	return 0
}

// 更新用户头像
type UserImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid   int64  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Image string `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *UserImageRequest) Reset() {
	*x = UserImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foodi_user_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserImageRequest) ProtoMessage() {}

func (x *UserImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foodi_user_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserImageRequest.ProtoReflect.Descriptor instead.
func (*UserImageRequest) Descriptor() ([]byte, []int) {
	return file_foodi_user_service_proto_rawDescGZIP(), []int{2}
}

func (x *UserImageRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *UserImageRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

// 获取用户详情
type UserDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
}

func (x *UserDetailRequest) Reset() {
	*x = UserDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foodi_user_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDetailRequest) ProtoMessage() {}

func (x *UserDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foodi_user_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDetailRequest.ProtoReflect.Descriptor instead.
func (*UserDetailRequest) Descriptor() ([]byte, []int) {
	return file_foodi_user_service_proto_rawDescGZIP(), []int{3}
}

func (x *UserDetailRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type UserDetailReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	NikeName string `protobuf:"bytes,2,opt,name=NikeName,proto3" json:"NikeName,omitempty"`
	Image    string `protobuf:"bytes,3,opt,name=Image,proto3" json:"Image,omitempty"`
	Age      int32  `protobuf:"varint,4,opt,name=Age,proto3" json:"Age,omitempty"`
	Gender   string `protobuf:"bytes,5,opt,name=Gender,proto3" json:"Gender,omitempty"`
	Birthday string `protobuf:"bytes,6,opt,name=Birthday,proto3" json:"Birthday,omitempty"`
	Region   string `protobuf:"bytes,7,opt,name=Region,proto3" json:"Region,omitempty"`
	LV       int32  `protobuf:"varint,8,opt,name=LV,proto3" json:"LV,omitempty"`
	VIP      int32  `protobuf:"varint,9,opt,name=VIP,proto3" json:"VIP,omitempty"`
}

func (x *UserDetailReply) Reset() {
	*x = UserDetailReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foodi_user_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDetailReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDetailReply) ProtoMessage() {}

func (x *UserDetailReply) ProtoReflect() protoreflect.Message {
	mi := &file_foodi_user_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDetailReply.ProtoReflect.Descriptor instead.
func (*UserDetailReply) Descriptor() ([]byte, []int) {
	return file_foodi_user_service_proto_rawDescGZIP(), []int{4}
}

func (x *UserDetailReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserDetailReply) GetNikeName() string {
	if x != nil {
		return x.NikeName
	}
	return ""
}

func (x *UserDetailReply) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *UserDetailReply) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UserDetailReply) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *UserDetailReply) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *UserDetailReply) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *UserDetailReply) GetLV() int32 {
	if x != nil {
		return x.LV
	}
	return 0
}

func (x *UserDetailReply) GetVIP() int32 {
	if x != nil {
		return x.VIP
	}
	return 0
}

type UserOKReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *UserOKReply) Reset() {
	*x = UserOKReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foodi_user_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserOKReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserOKReply) ProtoMessage() {}

func (x *UserOKReply) ProtoReflect() protoreflect.Message {
	mi := &file_foodi_user_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserOKReply.ProtoReflect.Descriptor instead.
func (*UserOKReply) Descriptor() ([]byte, []int) {
	return file_foodi_user_service_proto_rawDescGZIP(), []int{5}
}

func (x *UserOKReply) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type UpdateUserInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid     int64  `protobuf:"varint,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	NikeName string `protobuf:"bytes,2,opt,name=nikeName,proto3" json:"nikeName,omitempty"`
	CardID   string `protobuf:"bytes,3,opt,name=cardID,proto3" json:"cardID,omitempty"`
	Gender   string `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
	Age      int64  `protobuf:"varint,5,opt,name=age,proto3" json:"age,omitempty"`
	Birthday string `protobuf:"bytes,6,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Region   string `protobuf:"bytes,7,opt,name=region,proto3" json:"region,omitempty"`
	Name     string `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateUserInfoRequest) Reset() {
	*x = UpdateUserInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foodi_user_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserInfoRequest) ProtoMessage() {}

func (x *UpdateUserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foodi_user_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserInfoRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserInfoRequest) Descriptor() ([]byte, []int) {
	return file_foodi_user_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateUserInfoRequest) GetUuid() int64 {
	if x != nil {
		return x.Uuid
	}
	return 0
}

func (x *UpdateUserInfoRequest) GetNikeName() string {
	if x != nil {
		return x.NikeName
	}
	return ""
}

func (x *UpdateUserInfoRequest) GetCardID() string {
	if x != nil {
		return x.CardID
	}
	return ""
}

func (x *UpdateUserInfoRequest) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *UpdateUserInfoRequest) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UpdateUserInfoRequest) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *UpdateUserInfoRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *UpdateUserInfoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_foodi_user_service_proto protoreflect.FileDescriptor

var file_foodi_user_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x66, 0x6f, 0x6f, 0x64, 0x69, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x66, 0x6f, 0x6f, 0x64,
	0x69, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x55, 0x69, 0x64, 0x22, 0x61, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69, 0x6b,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x69, 0x6b,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x4c,
	0x56, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x4c, 0x56, 0x22, 0x3a, 0x0a, 0x10, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x25, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x55, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x55, 0x69, 0x64, 0x22, 0xd7,
	0x01, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69, 0x6b, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x69, 0x6b, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x67, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x41, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x47, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x4c, 0x56, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x4c, 0x56, 0x12, 0x10, 0x0a, 0x03, 0x56, 0x49, 0x50, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x56, 0x49, 0x50, 0x22, 0x1d, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72,
	0x4f, 0x4b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0xd1, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x6b, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x6b, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x64, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x61, 0x72, 0x64, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xce, 0x02, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x08, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1d, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x69, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x69, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x54, 0x0a, 0x0a,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x2e, 0x66, 0x6f, 0x6f,
	0x64, 0x69, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x69, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x4e, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x22, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x69, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x69, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x4b, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x58, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x27, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x69, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x66, 0x6f, 0x6f, 0x64, 0x69, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x4b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x29, 0x5a, 0x27,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6f, 0x6f, 0x64, 0x69,
	0x2d, 0x6f, 0x72, 0x67, 0x2f, 0x66, 0x6f, 0x6f, 0x64, 0x69, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_foodi_user_service_proto_rawDescOnce sync.Once
	file_foodi_user_service_proto_rawDescData = file_foodi_user_service_proto_rawDesc
)

func file_foodi_user_service_proto_rawDescGZIP() []byte {
	file_foodi_user_service_proto_rawDescOnce.Do(func() {
		file_foodi_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_foodi_user_service_proto_rawDescData)
	})
	return file_foodi_user_service_proto_rawDescData
}

var file_foodi_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_foodi_user_service_proto_goTypes = []interface{}{
	(*UserRequest)(nil),           // 0: foodiUserService.UserRequest
	(*UserReply)(nil),             // 1: foodiUserService.UserReply
	(*UserImageRequest)(nil),      // 2: foodiUserService.UserImageRequest
	(*UserDetailRequest)(nil),     // 3: foodiUserService.UserDetailRequest
	(*UserDetailReply)(nil),       // 4: foodiUserService.UserDetailReply
	(*UserOKReply)(nil),           // 5: foodiUserService.UserOKReply
	(*UpdateUserInfoRequest)(nil), // 6: foodiUserService.UpdateUserInfoRequest
}
var file_foodi_user_service_proto_depIdxs = []int32{
	0, // 0: foodiUserService.User.BaseInfo:input_type -> foodiUserService.UserRequest
	3, // 1: foodiUserService.User.DetailInfo:input_type -> foodiUserService.UserDetailRequest
	2, // 2: foodiUserService.User.UserImage:input_type -> foodiUserService.UserImageRequest
	6, // 3: foodiUserService.User.UpdateUserInfo:input_type -> foodiUserService.UpdateUserInfoRequest
	1, // 4: foodiUserService.User.BaseInfo:output_type -> foodiUserService.UserReply
	4, // 5: foodiUserService.User.DetailInfo:output_type -> foodiUserService.UserDetailReply
	5, // 6: foodiUserService.User.UserImage:output_type -> foodiUserService.UserOKReply
	5, // 7: foodiUserService.User.UpdateUserInfo:output_type -> foodiUserService.UserOKReply
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_foodi_user_service_proto_init() }
func file_foodi_user_service_proto_init() {
	if File_foodi_user_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_foodi_user_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRequest); i {
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
		file_foodi_user_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserReply); i {
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
		file_foodi_user_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserImageRequest); i {
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
		file_foodi_user_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDetailRequest); i {
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
		file_foodi_user_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDetailReply); i {
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
		file_foodi_user_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserOKReply); i {
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
		file_foodi_user_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserInfoRequest); i {
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
			RawDescriptor: file_foodi_user_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_foodi_user_service_proto_goTypes,
		DependencyIndexes: file_foodi_user_service_proto_depIdxs,
		MessageInfos:      file_foodi_user_service_proto_msgTypes,
	}.Build()
	File_foodi_user_service_proto = out.File
	file_foodi_user_service_proto_rawDesc = nil
	file_foodi_user_service_proto_goTypes = nil
	file_foodi_user_service_proto_depIdxs = nil
}
