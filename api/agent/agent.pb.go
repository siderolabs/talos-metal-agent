// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v4.24.4
// source: agent/agent.proto

package agentpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HelloRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	mi := &file_agent_agent_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{0}
}

type HelloResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	mi := &file_agent_agent_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{1}
}

type GetPowerManagementRequest struct {
	state         protoimpl.MessageState          `protogen:"open.v1"`
	Ipmi          *GetPowerManagementRequest_IPMI `protobuf:"bytes,1,opt,name=ipmi,proto3" json:"ipmi,omitempty"`
	Api           *GetPowerManagementRequest_API  `protobuf:"bytes,2,opt,name=api,proto3" json:"api,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPowerManagementRequest) Reset() {
	*x = GetPowerManagementRequest{}
	mi := &file_agent_agent_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPowerManagementRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPowerManagementRequest) ProtoMessage() {}

func (x *GetPowerManagementRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPowerManagementRequest.ProtoReflect.Descriptor instead.
func (*GetPowerManagementRequest) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{2}
}

func (x *GetPowerManagementRequest) GetIpmi() *GetPowerManagementRequest_IPMI {
	if x != nil {
		return x.Ipmi
	}
	return nil
}

func (x *GetPowerManagementRequest) GetApi() *GetPowerManagementRequest_API {
	if x != nil {
		return x.Api
	}
	return nil
}

type GetPowerManagementResponse struct {
	state         protoimpl.MessageState           `protogen:"open.v1"`
	Ipmi          *GetPowerManagementResponse_IPMI `protobuf:"bytes,1,opt,name=ipmi,proto3" json:"ipmi,omitempty"`
	Api           *GetPowerManagementResponse_API  `protobuf:"bytes,2,opt,name=api,proto3" json:"api,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPowerManagementResponse) Reset() {
	*x = GetPowerManagementResponse{}
	mi := &file_agent_agent_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPowerManagementResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPowerManagementResponse) ProtoMessage() {}

func (x *GetPowerManagementResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPowerManagementResponse.ProtoReflect.Descriptor instead.
func (*GetPowerManagementResponse) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{3}
}

func (x *GetPowerManagementResponse) GetIpmi() *GetPowerManagementResponse_IPMI {
	if x != nil {
		return x.Ipmi
	}
	return nil
}

func (x *GetPowerManagementResponse) GetApi() *GetPowerManagementResponse_API {
	if x != nil {
		return x.Api
	}
	return nil
}

type SetPowerManagementRequest struct {
	state         protoimpl.MessageState          `protogen:"open.v1"`
	Ipmi          *SetPowerManagementRequest_IPMI `protobuf:"bytes,1,opt,name=ipmi,proto3" json:"ipmi,omitempty"`
	Api           *SetPowerManagementRequest_API  `protobuf:"bytes,2,opt,name=api,proto3" json:"api,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetPowerManagementRequest) Reset() {
	*x = SetPowerManagementRequest{}
	mi := &file_agent_agent_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetPowerManagementRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPowerManagementRequest) ProtoMessage() {}

func (x *SetPowerManagementRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPowerManagementRequest.ProtoReflect.Descriptor instead.
func (*SetPowerManagementRequest) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{4}
}

func (x *SetPowerManagementRequest) GetIpmi() *SetPowerManagementRequest_IPMI {
	if x != nil {
		return x.Ipmi
	}
	return nil
}

func (x *SetPowerManagementRequest) GetApi() *SetPowerManagementRequest_API {
	if x != nil {
		return x.Api
	}
	return nil
}

type SetPowerManagementResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetPowerManagementResponse) Reset() {
	*x = SetPowerManagementResponse{}
	mi := &file_agent_agent_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetPowerManagementResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPowerManagementResponse) ProtoMessage() {}

func (x *SetPowerManagementResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPowerManagementResponse.ProtoReflect.Descriptor instead.
func (*SetPowerManagementResponse) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{5}
}

type RebootRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RebootRequest) Reset() {
	*x = RebootRequest{}
	mi := &file_agent_agent_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RebootRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RebootRequest) ProtoMessage() {}

func (x *RebootRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RebootRequest.ProtoReflect.Descriptor instead.
func (*RebootRequest) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{6}
}

type RebootResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RebootResponse) Reset() {
	*x = RebootResponse{}
	mi := &file_agent_agent_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RebootResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RebootResponse) ProtoMessage() {}

func (x *RebootResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RebootResponse.ProtoReflect.Descriptor instead.
func (*RebootResponse) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{7}
}

type WipeDisksRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Zeroes        bool                   `protobuf:"varint,1,opt,name=zeroes,proto3" json:"zeroes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WipeDisksRequest) Reset() {
	*x = WipeDisksRequest{}
	mi := &file_agent_agent_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WipeDisksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WipeDisksRequest) ProtoMessage() {}

func (x *WipeDisksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WipeDisksRequest.ProtoReflect.Descriptor instead.
func (*WipeDisksRequest) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{8}
}

func (x *WipeDisksRequest) GetZeroes() bool {
	if x != nil {
		return x.Zeroes
	}
	return false
}

type WipeDisksResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WipeDisksResponse) Reset() {
	*x = WipeDisksResponse{}
	mi := &file_agent_agent_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WipeDisksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WipeDisksResponse) ProtoMessage() {}

func (x *WipeDisksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WipeDisksResponse.ProtoReflect.Descriptor instead.
func (*WipeDisksResponse) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{9}
}

type GetPowerManagementRequest_IPMI struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CheckUsername string                 `protobuf:"bytes,1,opt,name=check_username,json=checkUsername,proto3" json:"check_username,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPowerManagementRequest_IPMI) Reset() {
	*x = GetPowerManagementRequest_IPMI{}
	mi := &file_agent_agent_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPowerManagementRequest_IPMI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPowerManagementRequest_IPMI) ProtoMessage() {}

func (x *GetPowerManagementRequest_IPMI) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPowerManagementRequest_IPMI.ProtoReflect.Descriptor instead.
func (*GetPowerManagementRequest_IPMI) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{2, 0}
}

func (x *GetPowerManagementRequest_IPMI) GetCheckUsername() string {
	if x != nil {
		return x.CheckUsername
	}
	return ""
}

type GetPowerManagementRequest_API struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPowerManagementRequest_API) Reset() {
	*x = GetPowerManagementRequest_API{}
	mi := &file_agent_agent_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPowerManagementRequest_API) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPowerManagementRequest_API) ProtoMessage() {}

func (x *GetPowerManagementRequest_API) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPowerManagementRequest_API.ProtoReflect.Descriptor instead.
func (*GetPowerManagementRequest_API) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{2, 1}
}

type GetPowerManagementResponse_IPMI struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       string                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Port          uint32                 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	UserExists    bool                   `protobuf:"varint,3,opt,name=user_exists,json=userExists,proto3" json:"user_exists,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPowerManagementResponse_IPMI) Reset() {
	*x = GetPowerManagementResponse_IPMI{}
	mi := &file_agent_agent_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPowerManagementResponse_IPMI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPowerManagementResponse_IPMI) ProtoMessage() {}

func (x *GetPowerManagementResponse_IPMI) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPowerManagementResponse_IPMI.ProtoReflect.Descriptor instead.
func (*GetPowerManagementResponse_IPMI) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{3, 0}
}

func (x *GetPowerManagementResponse_IPMI) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GetPowerManagementResponse_IPMI) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *GetPowerManagementResponse_IPMI) GetUserExists() bool {
	if x != nil {
		return x.UserExists
	}
	return false
}

type GetPowerManagementResponse_API struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPowerManagementResponse_API) Reset() {
	*x = GetPowerManagementResponse_API{}
	mi := &file_agent_agent_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPowerManagementResponse_API) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPowerManagementResponse_API) ProtoMessage() {}

func (x *GetPowerManagementResponse_API) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPowerManagementResponse_API.ProtoReflect.Descriptor instead.
func (*GetPowerManagementResponse_API) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{3, 1}
}

type SetPowerManagementRequest_IPMI struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetPowerManagementRequest_IPMI) Reset() {
	*x = SetPowerManagementRequest_IPMI{}
	mi := &file_agent_agent_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetPowerManagementRequest_IPMI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPowerManagementRequest_IPMI) ProtoMessage() {}

func (x *SetPowerManagementRequest_IPMI) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPowerManagementRequest_IPMI.ProtoReflect.Descriptor instead.
func (*SetPowerManagementRequest_IPMI) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{4, 0}
}

func (x *SetPowerManagementRequest_IPMI) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SetPowerManagementRequest_IPMI) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SetPowerManagementRequest_API struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetPowerManagementRequest_API) Reset() {
	*x = SetPowerManagementRequest_API{}
	mi := &file_agent_agent_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetPowerManagementRequest_API) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPowerManagementRequest_API) ProtoMessage() {}

func (x *SetPowerManagementRequest_API) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPowerManagementRequest_API.ProtoReflect.Descriptor instead.
func (*SetPowerManagementRequest_API) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{4, 1}
}

var File_agent_agent_proto protoreflect.FileDescriptor

var file_agent_agent_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0e, 0x0a,
	0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a,
	0x0d, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xce,
	0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x04,
	0x69, 0x70, 0x6d, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x77, 0x65, 0x72,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x49, 0x50, 0x4d, 0x49, 0x52, 0x04, 0x69, 0x70, 0x6d, 0x69, 0x12, 0x3b, 0x0a, 0x03,
	0x61, 0x70, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x41, 0x50, 0x49, 0x52, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x2d, 0x0a, 0x04, 0x49, 0x50, 0x4d,
	0x49, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x05, 0x0a, 0x03, 0x41, 0x50, 0x49, 0x22,
	0xf9, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f,
	0x0a, 0x04, 0x69, 0x70, 0x6d, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x77,
	0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49, 0x50, 0x4d, 0x49, 0x52, 0x04, 0x69, 0x70, 0x6d, 0x69, 0x12,
	0x3c, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x77,
	0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x55, 0x0a,
	0x04, 0x49, 0x50, 0x4d, 0x49, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x69, 0x73,
	0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x73, 0x1a, 0x05, 0x0a, 0x03, 0x41, 0x50, 0x49, 0x22, 0xdf, 0x01, 0x0a, 0x19,
	0x53, 0x65, 0x74, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x04, 0x69, 0x70, 0x6d,
	0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x49,
	0x50, 0x4d, 0x49, 0x52, 0x04, 0x69, 0x70, 0x6d, 0x69, 0x12, 0x3b, 0x0a, 0x03, 0x61, 0x70, 0x69,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x50,
	0x49, 0x52, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x3e, 0x0a, 0x04, 0x49, 0x50, 0x4d, 0x49, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x1a, 0x05, 0x0a, 0x03, 0x41, 0x50, 0x49, 0x22, 0x1c, 0x0a,
	0x1a, 0x53, 0x65, 0x74, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x52,
	0x65, 0x62, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e,
	0x52, 0x65, 0x62, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a,
	0x0a, 0x10, 0x57, 0x69, 0x70, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x7a, 0x65, 0x72, 0x6f, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x7a, 0x65, 0x72, 0x6f, 0x65, 0x73, 0x22, 0x13, 0x0a, 0x11, 0x57, 0x69,
	0x70, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xa1, 0x03, 0x0a, 0x0c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3c, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x18, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x77, 0x65,
	0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65,
	0x74, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x52, 0x65, 0x62, 0x6f,
	0x6f, 0x74, 0x12, 0x19, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x52, 0x65, 0x62, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x62, 0x6f, 0x6f,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x09, 0x57, 0x69, 0x70,
	0x65, 0x44, 0x69, 0x73, 0x6b, 0x73, 0x12, 0x1c, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x57, 0x69, 0x70, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x57, 0x69, 0x70, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x72, 0x6f, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x74, 0x61, 0x6c,
	0x6f, 0x73, 0x2d, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_agent_agent_proto_rawDescOnce sync.Once
	file_agent_agent_proto_rawDescData = file_agent_agent_proto_rawDesc
)

func file_agent_agent_proto_rawDescGZIP() []byte {
	file_agent_agent_proto_rawDescOnce.Do(func() {
		file_agent_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_agent_agent_proto_rawDescData)
	})
	return file_agent_agent_proto_rawDescData
}

var file_agent_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_agent_agent_proto_goTypes = []any{
	(*HelloRequest)(nil),                    // 0: management.HelloRequest
	(*HelloResponse)(nil),                   // 1: management.HelloResponse
	(*GetPowerManagementRequest)(nil),       // 2: management.GetPowerManagementRequest
	(*GetPowerManagementResponse)(nil),      // 3: management.GetPowerManagementResponse
	(*SetPowerManagementRequest)(nil),       // 4: management.SetPowerManagementRequest
	(*SetPowerManagementResponse)(nil),      // 5: management.SetPowerManagementResponse
	(*RebootRequest)(nil),                   // 6: management.RebootRequest
	(*RebootResponse)(nil),                  // 7: management.RebootResponse
	(*WipeDisksRequest)(nil),                // 8: management.WipeDisksRequest
	(*WipeDisksResponse)(nil),               // 9: management.WipeDisksResponse
	(*GetPowerManagementRequest_IPMI)(nil),  // 10: management.GetPowerManagementRequest.IPMI
	(*GetPowerManagementRequest_API)(nil),   // 11: management.GetPowerManagementRequest.API
	(*GetPowerManagementResponse_IPMI)(nil), // 12: management.GetPowerManagementResponse.IPMI
	(*GetPowerManagementResponse_API)(nil),  // 13: management.GetPowerManagementResponse.API
	(*SetPowerManagementRequest_IPMI)(nil),  // 14: management.SetPowerManagementRequest.IPMI
	(*SetPowerManagementRequest_API)(nil),   // 15: management.SetPowerManagementRequest.API
}
var file_agent_agent_proto_depIdxs = []int32{
	10, // 0: management.GetPowerManagementRequest.ipmi:type_name -> management.GetPowerManagementRequest.IPMI
	11, // 1: management.GetPowerManagementRequest.api:type_name -> management.GetPowerManagementRequest.API
	12, // 2: management.GetPowerManagementResponse.ipmi:type_name -> management.GetPowerManagementResponse.IPMI
	13, // 3: management.GetPowerManagementResponse.api:type_name -> management.GetPowerManagementResponse.API
	14, // 4: management.SetPowerManagementRequest.ipmi:type_name -> management.SetPowerManagementRequest.IPMI
	15, // 5: management.SetPowerManagementRequest.api:type_name -> management.SetPowerManagementRequest.API
	0,  // 6: management.AgentService.Hello:input_type -> management.HelloRequest
	2,  // 7: management.AgentService.GetPowerManagement:input_type -> management.GetPowerManagementRequest
	4,  // 8: management.AgentService.SetPowerManagement:input_type -> management.SetPowerManagementRequest
	6,  // 9: management.AgentService.Reboot:input_type -> management.RebootRequest
	8,  // 10: management.AgentService.WipeDisks:input_type -> management.WipeDisksRequest
	1,  // 11: management.AgentService.Hello:output_type -> management.HelloResponse
	3,  // 12: management.AgentService.GetPowerManagement:output_type -> management.GetPowerManagementResponse
	5,  // 13: management.AgentService.SetPowerManagement:output_type -> management.SetPowerManagementResponse
	7,  // 14: management.AgentService.Reboot:output_type -> management.RebootResponse
	9,  // 15: management.AgentService.WipeDisks:output_type -> management.WipeDisksResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_agent_agent_proto_init() }
func file_agent_agent_proto_init() {
	if File_agent_agent_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_agent_agent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_agent_agent_proto_goTypes,
		DependencyIndexes: file_agent_agent_proto_depIdxs,
		MessageInfos:      file_agent_agent_proto_msgTypes,
	}.Build()
	File_agent_agent_proto = out.File
	file_agent_agent_proto_rawDesc = nil
	file_agent_agent_proto_goTypes = nil
	file_agent_agent_proto_depIdxs = nil
}
