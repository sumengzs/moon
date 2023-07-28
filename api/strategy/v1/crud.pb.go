// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: strategy/v1/crud.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	api "prometheus-manager/api"
	strategy "prometheus-manager/api/strategy"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateCrudRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes       []string              `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	StrategyDir *strategy.StrategyDir `protobuf:"bytes,2,opt,name=strategy_dir,json=strategyDir,proto3" json:"strategy_dir,omitempty"`
}

func (x *CreateCrudRequest) Reset() {
	*x = CreateCrudRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_v1_crud_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCrudRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCrudRequest) ProtoMessage() {}

func (x *CreateCrudRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_v1_crud_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCrudRequest.ProtoReflect.Descriptor instead.
func (*CreateCrudRequest) Descriptor() ([]byte, []int) {
	return file_strategy_v1_crud_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCrudRequest) GetNodes() []string {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *CreateCrudRequest) GetStrategyDir() *strategy.StrategyDir {
	if x != nil {
		return x.StrategyDir
	}
	return nil
}

type CreateCrudReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 响应内容
	Response *api.Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *CreateCrudReply) Reset() {
	*x = CreateCrudReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_v1_crud_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCrudReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCrudReply) ProtoMessage() {}

func (x *CreateCrudReply) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_v1_crud_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCrudReply.ProtoReflect.Descriptor instead.
func (*CreateCrudReply) Descriptor() ([]byte, []int) {
	return file_strategy_v1_crud_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCrudReply) GetResponse() *api.Response {
	if x != nil {
		return x.Response
	}
	return nil
}

type UpdateCrudRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Nodes       []string              `protobuf:"bytes,2,rep,name=nodes,proto3" json:"nodes,omitempty"`
	StrategyDir *strategy.StrategyDir `protobuf:"bytes,3,opt,name=strategy_dir,json=strategyDir,proto3" json:"strategy_dir,omitempty"`
}

func (x *UpdateCrudRequest) Reset() {
	*x = UpdateCrudRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_v1_crud_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCrudRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCrudRequest) ProtoMessage() {}

func (x *UpdateCrudRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_v1_crud_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCrudRequest.ProtoReflect.Descriptor instead.
func (*UpdateCrudRequest) Descriptor() ([]byte, []int) {
	return file_strategy_v1_crud_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateCrudRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateCrudRequest) GetNodes() []string {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *UpdateCrudRequest) GetStrategyDir() *strategy.StrategyDir {
	if x != nil {
		return x.StrategyDir
	}
	return nil
}

type UpdateCrudReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 响应内容
	Response *api.Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *UpdateCrudReply) Reset() {
	*x = UpdateCrudReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_v1_crud_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCrudReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCrudReply) ProtoMessage() {}

func (x *UpdateCrudReply) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_v1_crud_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCrudReply.ProtoReflect.Descriptor instead.
func (*UpdateCrudReply) Descriptor() ([]byte, []int) {
	return file_strategy_v1_crud_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateCrudReply) GetResponse() *api.Response {
	if x != nil {
		return x.Response
	}
	return nil
}

type DeleteCrudRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Nodes []string `protobuf:"bytes,2,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *DeleteCrudRequest) Reset() {
	*x = DeleteCrudRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_v1_crud_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCrudRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCrudRequest) ProtoMessage() {}

func (x *DeleteCrudRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_v1_crud_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCrudRequest.ProtoReflect.Descriptor instead.
func (*DeleteCrudRequest) Descriptor() ([]byte, []int) {
	return file_strategy_v1_crud_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteCrudRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteCrudRequest) GetNodes() []string {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type DeleteCrudReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 响应内容
	Response *api.Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *DeleteCrudReply) Reset() {
	*x = DeleteCrudReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_v1_crud_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCrudReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCrudReply) ProtoMessage() {}

func (x *DeleteCrudReply) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_v1_crud_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCrudReply.ProtoReflect.Descriptor instead.
func (*DeleteCrudReply) Descriptor() ([]byte, []int) {
	return file_strategy_v1_crud_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteCrudReply) GetResponse() *api.Response {
	if x != nil {
		return x.Response
	}
	return nil
}

type GetCrudRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCrudRequest) Reset() {
	*x = GetCrudRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_v1_crud_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCrudRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCrudRequest) ProtoMessage() {}

func (x *GetCrudRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_v1_crud_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCrudRequest.ProtoReflect.Descriptor instead.
func (*GetCrudRequest) Descriptor() ([]byte, []int) {
	return file_strategy_v1_crud_proto_rawDescGZIP(), []int{6}
}

func (x *GetCrudRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetCrudReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 响应内容
	Response *api.Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *GetCrudReply) Reset() {
	*x = GetCrudReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_v1_crud_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCrudReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCrudReply) ProtoMessage() {}

func (x *GetCrudReply) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_v1_crud_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCrudReply.ProtoReflect.Descriptor instead.
func (*GetCrudReply) Descriptor() ([]byte, []int) {
	return file_strategy_v1_crud_proto_rawDescGZIP(), []int{7}
}

func (x *GetCrudReply) GetResponse() *api.Response {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_strategy_v1_crud_proto protoreflect.FileDescriptor

var file_strategy_v1_crud_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72,
	0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x72, 0x75, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x05, 0x6e, 0x6f,
	0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x92, 0x01,
	0x04, 0x08, 0x01, 0x10, 0x0a, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x0c,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x44, 0x69, 0x72, 0x52, 0x0b, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x44, 0x69, 0x72, 0x22, 0x3c, 0x0a, 0x0f, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x72, 0x75, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x29, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x83, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x72, 0x75, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20,
	0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x0a, 0xfa,
	0x42, 0x07, 0x92, 0x01, 0x04, 0x08, 0x01, 0x10, 0x0a, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73,
	0x12, 0x3c, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x5f, 0x64, 0x69, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x44, 0x69,
	0x72, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x44, 0x69, 0x72, 0x22, 0x3c,
	0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x75, 0x64, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x29, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0x0a, 0x11,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x72, 0x75, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x3c, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x72, 0x75, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x29, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x72, 0x75, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x39, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x72,
	0x75, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x29, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0xce, 0x03, 0x0a, 0x04, 0x43, 0x72, 0x75, 0x64, 0x12, 0x6f, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x72, 0x75, 0x64, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x72, 0x75, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x72, 0x75, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x75, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x75, 0x64, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x72, 0x75, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x75, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x1a, 0x16, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x64, 0x69, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x3a, 0x01, 0x2a, 0x12, 0x74, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x72, 0x75,
	0x64, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x72, 0x75, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x72,
	0x75, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x2a,
	0x18, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x68, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x43, 0x72, 0x75, 0x64, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x72, 0x75, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x72, 0x75, 0x64, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x42, 0x3a, 0x0a, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x25, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74,
	0x68, 0x65, 0x75, 0x73, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_strategy_v1_crud_proto_rawDescOnce sync.Once
	file_strategy_v1_crud_proto_rawDescData = file_strategy_v1_crud_proto_rawDesc
)

func file_strategy_v1_crud_proto_rawDescGZIP() []byte {
	file_strategy_v1_crud_proto_rawDescOnce.Do(func() {
		file_strategy_v1_crud_proto_rawDescData = protoimpl.X.CompressGZIP(file_strategy_v1_crud_proto_rawDescData)
	})
	return file_strategy_v1_crud_proto_rawDescData
}

var file_strategy_v1_crud_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_strategy_v1_crud_proto_goTypes = []interface{}{
	(*CreateCrudRequest)(nil),    // 0: api.strategy.v1.CreateCrudRequest
	(*CreateCrudReply)(nil),      // 1: api.strategy.v1.CreateCrudReply
	(*UpdateCrudRequest)(nil),    // 2: api.strategy.v1.UpdateCrudRequest
	(*UpdateCrudReply)(nil),      // 3: api.strategy.v1.UpdateCrudReply
	(*DeleteCrudRequest)(nil),    // 4: api.strategy.v1.DeleteCrudRequest
	(*DeleteCrudReply)(nil),      // 5: api.strategy.v1.DeleteCrudReply
	(*GetCrudRequest)(nil),       // 6: api.strategy.v1.GetCrudRequest
	(*GetCrudReply)(nil),         // 7: api.strategy.v1.GetCrudReply
	(*strategy.StrategyDir)(nil), // 8: api.strategy.StrategyDir
	(*api.Response)(nil),         // 9: api.Response
}
var file_strategy_v1_crud_proto_depIdxs = []int32{
	8,  // 0: api.strategy.v1.CreateCrudRequest.strategy_dir:type_name -> api.strategy.StrategyDir
	9,  // 1: api.strategy.v1.CreateCrudReply.response:type_name -> api.Response
	8,  // 2: api.strategy.v1.UpdateCrudRequest.strategy_dir:type_name -> api.strategy.StrategyDir
	9,  // 3: api.strategy.v1.UpdateCrudReply.response:type_name -> api.Response
	9,  // 4: api.strategy.v1.DeleteCrudReply.response:type_name -> api.Response
	9,  // 5: api.strategy.v1.GetCrudReply.response:type_name -> api.Response
	0,  // 6: api.strategy.v1.Crud.CreateCrud:input_type -> api.strategy.v1.CreateCrudRequest
	2,  // 7: api.strategy.v1.Crud.UpdateCrud:input_type -> api.strategy.v1.UpdateCrudRequest
	4,  // 8: api.strategy.v1.Crud.DeleteCrud:input_type -> api.strategy.v1.DeleteCrudRequest
	6,  // 9: api.strategy.v1.Crud.GetCrud:input_type -> api.strategy.v1.GetCrudRequest
	1,  // 10: api.strategy.v1.Crud.CreateCrud:output_type -> api.strategy.v1.CreateCrudReply
	3,  // 11: api.strategy.v1.Crud.UpdateCrud:output_type -> api.strategy.v1.UpdateCrudReply
	5,  // 12: api.strategy.v1.Crud.DeleteCrud:output_type -> api.strategy.v1.DeleteCrudReply
	7,  // 13: api.strategy.v1.Crud.GetCrud:output_type -> api.strategy.v1.GetCrudReply
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_strategy_v1_crud_proto_init() }
func file_strategy_v1_crud_proto_init() {
	if File_strategy_v1_crud_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_strategy_v1_crud_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCrudRequest); i {
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
		file_strategy_v1_crud_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCrudReply); i {
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
		file_strategy_v1_crud_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCrudRequest); i {
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
		file_strategy_v1_crud_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCrudReply); i {
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
		file_strategy_v1_crud_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCrudRequest); i {
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
		file_strategy_v1_crud_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCrudReply); i {
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
		file_strategy_v1_crud_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCrudRequest); i {
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
		file_strategy_v1_crud_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCrudReply); i {
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
			RawDescriptor: file_strategy_v1_crud_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_strategy_v1_crud_proto_goTypes,
		DependencyIndexes: file_strategy_v1_crud_proto_depIdxs,
		MessageInfos:      file_strategy_v1_crud_proto_msgTypes,
	}.Build()
	File_strategy_v1_crud_proto = out.File
	file_strategy_v1_crud_proto_rawDesc = nil
	file_strategy_v1_crud_proto_goTypes = nil
	file_strategy_v1_crud_proto_depIdxs = nil
}
