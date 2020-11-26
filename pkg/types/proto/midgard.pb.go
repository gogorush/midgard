// Copyright 2020 The golang.design Initiative authors.
// All rights reserved. Use of this source code is governed by
// a GNU GPL-3.0 license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: midgard.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ClipboardDataType int32

const (
	ClipboardDataType_PlainText ClipboardDataType = 0
	ClipboardDataType_ImagePNG  ClipboardDataType = 1
)

// Enum value maps for ClipboardDataType.
var (
	ClipboardDataType_name = map[int32]string{
		0: "PlainText",
		1: "ImagePNG",
	}
	ClipboardDataType_value = map[string]int32{
		"PlainText": 0,
		"ImagePNG":  1,
	}
)

func (x ClipboardDataType) Enum() *ClipboardDataType {
	p := new(ClipboardDataType)
	*p = x
	return p
}

func (x ClipboardDataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClipboardDataType) Descriptor() protoreflect.EnumDescriptor {
	return file_midgard_proto_enumTypes[0].Descriptor()
}

func (ClipboardDataType) Type() protoreflect.EnumType {
	return &file_midgard_proto_enumTypes[0]
}

func (x ClipboardDataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClipboardDataType.Descriptor instead.
func (ClipboardDataType) EnumDescriptor() ([]byte, []int) {
	return file_midgard_proto_rawDescGZIP(), []int{0}
}

type SourceType int32

const (
	SourceType_UniversalClipboard SourceType = 0
	SourceType_Attachment         SourceType = 1
)

// Enum value maps for SourceType.
var (
	SourceType_name = map[int32]string{
		0: "UniversalClipboard",
		1: "Attachment",
	}
	SourceType_value = map[string]int32{
		"UniversalClipboard": 0,
		"Attachment":         1,
	}
)

func (x SourceType) Enum() *SourceType {
	p := new(SourceType)
	*p = x
	return p
}

func (x SourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_midgard_proto_enumTypes[1].Descriptor()
}

func (SourceType) Type() protoreflect.EnumType {
	return &file_midgard_proto_enumTypes[1]
}

func (x SourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SourceType.Descriptor instead.
func (SourceType) EnumDescriptor() ([]byte, []int) {
	return file_midgard_proto_rawDescGZIP(), []int{1}
}

type PingInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingInput) Reset() {
	*x = PingInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_midgard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingInput) ProtoMessage() {}

func (x *PingInput) ProtoReflect() protoreflect.Message {
	mi := &file_midgard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingInput.ProtoReflect.Descriptor instead.
func (*PingInput) Descriptor() ([]byte, []int) {
	return file_midgard_proto_rawDescGZIP(), []int{0}
}

type PingOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *PingOutput) Reset() {
	*x = PingOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_midgard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingOutput) ProtoMessage() {}

func (x *PingOutput) ProtoReflect() protoreflect.Message {
	mi := &file_midgard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingOutput.ProtoReflect.Descriptor instead.
func (*PingOutput) Descriptor() ([]byte, []int) {
	return file_midgard_proto_rawDescGZIP(), []int{1}
}

func (x *PingOutput) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetFromUniversalClipboardInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFromUniversalClipboardInput) Reset() {
	*x = GetFromUniversalClipboardInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_midgard_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFromUniversalClipboardInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFromUniversalClipboardInput) ProtoMessage() {}

func (x *GetFromUniversalClipboardInput) ProtoReflect() protoreflect.Message {
	mi := &file_midgard_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFromUniversalClipboardInput.ProtoReflect.Descriptor instead.
func (*GetFromUniversalClipboardInput) Descriptor() ([]byte, []int) {
	return file_midgard_proto_rawDescGZIP(), []int{2}
}

type GetFromUniversalClipboardOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type ClipboardDataType `protobuf:"varint,1,opt,name=Type,proto3,enum=proto.ClipboardDataType" json:"Type,omitempty"`
	Data string            `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *GetFromUniversalClipboardOutput) Reset() {
	*x = GetFromUniversalClipboardOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_midgard_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFromUniversalClipboardOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFromUniversalClipboardOutput) ProtoMessage() {}

func (x *GetFromUniversalClipboardOutput) ProtoReflect() protoreflect.Message {
	mi := &file_midgard_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFromUniversalClipboardOutput.ProtoReflect.Descriptor instead.
func (*GetFromUniversalClipboardOutput) Descriptor() ([]byte, []int) {
	return file_midgard_proto_rawDescGZIP(), []int{3}
}

func (x *GetFromUniversalClipboardOutput) GetType() ClipboardDataType {
	if x != nil {
		return x.Type
	}
	return ClipboardDataType_PlainText
}

func (x *GetFromUniversalClipboardOutput) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type PutToUniversalClipboardInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type ClipboardDataType `protobuf:"varint,1,opt,name=Type,proto3,enum=proto.ClipboardDataType" json:"Type,omitempty"`
	Data string            `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *PutToUniversalClipboardInput) Reset() {
	*x = PutToUniversalClipboardInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_midgard_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutToUniversalClipboardInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutToUniversalClipboardInput) ProtoMessage() {}

func (x *PutToUniversalClipboardInput) ProtoReflect() protoreflect.Message {
	mi := &file_midgard_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutToUniversalClipboardInput.ProtoReflect.Descriptor instead.
func (*PutToUniversalClipboardInput) Descriptor() ([]byte, []int) {
	return file_midgard_proto_rawDescGZIP(), []int{4}
}

func (x *PutToUniversalClipboardInput) GetType() ClipboardDataType {
	if x != nil {
		return x.Type
	}
	return ClipboardDataType_PlainText
}

func (x *PutToUniversalClipboardInput) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type PutToUniversalClipboardOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *PutToUniversalClipboardOutput) Reset() {
	*x = PutToUniversalClipboardOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_midgard_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutToUniversalClipboardOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutToUniversalClipboardOutput) ProtoMessage() {}

func (x *PutToUniversalClipboardOutput) ProtoReflect() protoreflect.Message {
	mi := &file_midgard_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutToUniversalClipboardOutput.ProtoReflect.Descriptor instead.
func (*PutToUniversalClipboardOutput) Descriptor() ([]byte, []int) {
	return file_midgard_proto_rawDescGZIP(), []int{5}
}

func (x *PutToUniversalClipboardOutput) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type AllocateURLInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source SourceType `protobuf:"varint,1,opt,name=Source,proto3,enum=proto.SourceType" json:"Source,omitempty"`
	URI    string     `protobuf:"bytes,2,opt,name=URI,proto3" json:"URI,omitempty"`
	Data   string     `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *AllocateURLInput) Reset() {
	*x = AllocateURLInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_midgard_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllocateURLInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocateURLInput) ProtoMessage() {}

func (x *AllocateURLInput) ProtoReflect() protoreflect.Message {
	mi := &file_midgard_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocateURLInput.ProtoReflect.Descriptor instead.
func (*AllocateURLInput) Descriptor() ([]byte, []int) {
	return file_midgard_proto_rawDescGZIP(), []int{6}
}

func (x *AllocateURLInput) GetSource() SourceType {
	if x != nil {
		return x.Source
	}
	return SourceType_UniversalClipboard
}

func (x *AllocateURLInput) GetURI() string {
	if x != nil {
		return x.URI
	}
	return ""
}

func (x *AllocateURLInput) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type AllocateURLOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL     string `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *AllocateURLOutput) Reset() {
	*x = AllocateURLOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_midgard_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllocateURLOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocateURLOutput) ProtoMessage() {}

func (x *AllocateURLOutput) ProtoReflect() protoreflect.Message {
	mi := &file_midgard_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocateURLOutput.ProtoReflect.Descriptor instead.
func (*AllocateURLOutput) Descriptor() ([]byte, []int) {
	return file_midgard_proto_rawDescGZIP(), []int{7}
}

func (x *AllocateURLOutput) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

func (x *AllocateURLOutput) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_midgard_proto protoreflect.FileDescriptor

var file_midgard_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x69, 0x64, 0x67, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0b, 0x0a, 0x09, 0x50, 0x69, 0x6e, 0x67, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x22, 0x26, 0x0a, 0x0a, 0x50, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x20, 0x0a, 0x1e, 0x47,
	0x65, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x43,
	0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x63, 0x0a,
	0x1f, 0x47, 0x65, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x6c, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x2c, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x22, 0x60, 0x0a, 0x1c, 0x50, 0x75, 0x74, 0x54, 0x6f, 0x55, 0x6e, 0x69, 0x76, 0x65,
	0x72, 0x73, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x39, 0x0a, 0x1d, 0x50, 0x75, 0x74, 0x54, 0x6f, 0x55, 0x6e, 0x69,
	0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x63, 0x0a, 0x10, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x55, 0x52, 0x4c, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x29, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x55, 0x52, 0x49, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x49,
	0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x3f, 0x0a, 0x11, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65,
	0x55, 0x52, 0x4c, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x12, 0x18, 0x0a, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x30, 0x0a, 0x11, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x6c,
	0x61, 0x69, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x50, 0x4e, 0x47, 0x10, 0x01, 0x2a, 0x34, 0x0a, 0x0a, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x6c, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x10, 0x00, 0x12, 0x0e, 0x0a,
	0x0a, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x01, 0x32, 0xd2, 0x02,
	0x0a, 0x07, 0x4d, 0x69, 0x64, 0x67, 0x61, 0x72, 0x64, 0x12, 0x2d, 0x0a, 0x04, 0x50, 0x69, 0x6e,
	0x67, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x69, 0x6e, 0x67,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x12, 0x6c, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x46,
	0x72, 0x6f, 0x6d, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x70,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x46, 0x72, 0x6f, 0x6d, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x43, 0x6c,
	0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x26, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x55, 0x6e, 0x69, 0x76,
	0x65, 0x72, 0x73, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x12, 0x66, 0x0a, 0x17, 0x50, 0x75, 0x74, 0x54, 0x6f, 0x55,
	0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x12, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x74, 0x54, 0x6f, 0x55,
	0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x75, 0x74, 0x54, 0x6f, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x43, 0x6c, 0x69,
	0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x12, 0x42,
	0x0a, 0x0b, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x55, 0x52, 0x4c, 0x12, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x55, 0x52,
	0x4c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x55, 0x52, 0x4c, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_midgard_proto_rawDescOnce sync.Once
	file_midgard_proto_rawDescData = file_midgard_proto_rawDesc
)

func file_midgard_proto_rawDescGZIP() []byte {
	file_midgard_proto_rawDescOnce.Do(func() {
		file_midgard_proto_rawDescData = protoimpl.X.CompressGZIP(file_midgard_proto_rawDescData)
	})
	return file_midgard_proto_rawDescData
}

var file_midgard_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_midgard_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_midgard_proto_goTypes = []interface{}{
	(ClipboardDataType)(0),                  // 0: proto.ClipboardDataType
	(SourceType)(0),                         // 1: proto.SourceType
	(*PingInput)(nil),                       // 2: proto.PingInput
	(*PingOutput)(nil),                      // 3: proto.PingOutput
	(*GetFromUniversalClipboardInput)(nil),  // 4: proto.GetFromUniversalClipboardInput
	(*GetFromUniversalClipboardOutput)(nil), // 5: proto.GetFromUniversalClipboardOutput
	(*PutToUniversalClipboardInput)(nil),    // 6: proto.PutToUniversalClipboardInput
	(*PutToUniversalClipboardOutput)(nil),   // 7: proto.PutToUniversalClipboardOutput
	(*AllocateURLInput)(nil),                // 8: proto.AllocateURLInput
	(*AllocateURLOutput)(nil),               // 9: proto.AllocateURLOutput
}
var file_midgard_proto_depIdxs = []int32{
	0, // 0: proto.GetFromUniversalClipboardOutput.Type:type_name -> proto.ClipboardDataType
	0, // 1: proto.PutToUniversalClipboardInput.Type:type_name -> proto.ClipboardDataType
	1, // 2: proto.AllocateURLInput.Source:type_name -> proto.SourceType
	2, // 3: proto.Midgard.Ping:input_type -> proto.PingInput
	4, // 4: proto.Midgard.GetFromUniversalClipboard:input_type -> proto.GetFromUniversalClipboardInput
	6, // 5: proto.Midgard.PutToUniversalClipboard:input_type -> proto.PutToUniversalClipboardInput
	8, // 6: proto.Midgard.AllocateURL:input_type -> proto.AllocateURLInput
	3, // 7: proto.Midgard.Ping:output_type -> proto.PingOutput
	5, // 8: proto.Midgard.GetFromUniversalClipboard:output_type -> proto.GetFromUniversalClipboardOutput
	7, // 9: proto.Midgard.PutToUniversalClipboard:output_type -> proto.PutToUniversalClipboardOutput
	9, // 10: proto.Midgard.AllocateURL:output_type -> proto.AllocateURLOutput
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_midgard_proto_init() }
func file_midgard_proto_init() {
	if File_midgard_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_midgard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingInput); i {
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
		file_midgard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingOutput); i {
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
		file_midgard_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFromUniversalClipboardInput); i {
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
		file_midgard_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFromUniversalClipboardOutput); i {
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
		file_midgard_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutToUniversalClipboardInput); i {
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
		file_midgard_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutToUniversalClipboardOutput); i {
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
		file_midgard_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllocateURLInput); i {
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
		file_midgard_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllocateURLOutput); i {
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
			RawDescriptor: file_midgard_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_midgard_proto_goTypes,
		DependencyIndexes: file_midgard_proto_depIdxs,
		EnumInfos:         file_midgard_proto_enumTypes,
		MessageInfos:      file_midgard_proto_msgTypes,
	}.Build()
	File_midgard_proto = out.File
	file_midgard_proto_rawDesc = nil
	file_midgard_proto_goTypes = nil
	file_midgard_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MidgardClient is the client API for Midgard service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MidgardClient interface {
	Ping(ctx context.Context, in *PingInput, opts ...grpc.CallOption) (*PingOutput, error)
	GetFromUniversalClipboard(ctx context.Context, in *GetFromUniversalClipboardInput, opts ...grpc.CallOption) (*GetFromUniversalClipboardOutput, error)
	PutToUniversalClipboard(ctx context.Context, in *PutToUniversalClipboardInput, opts ...grpc.CallOption) (*PutToUniversalClipboardOutput, error)
	AllocateURL(ctx context.Context, in *AllocateURLInput, opts ...grpc.CallOption) (*AllocateURLOutput, error)
}

type midgardClient struct {
	cc grpc.ClientConnInterface
}

func NewMidgardClient(cc grpc.ClientConnInterface) MidgardClient {
	return &midgardClient{cc}
}

func (c *midgardClient) Ping(ctx context.Context, in *PingInput, opts ...grpc.CallOption) (*PingOutput, error) {
	out := new(PingOutput)
	err := c.cc.Invoke(ctx, "/proto.Midgard/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *midgardClient) GetFromUniversalClipboard(ctx context.Context, in *GetFromUniversalClipboardInput, opts ...grpc.CallOption) (*GetFromUniversalClipboardOutput, error) {
	out := new(GetFromUniversalClipboardOutput)
	err := c.cc.Invoke(ctx, "/proto.Midgard/GetFromUniversalClipboard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *midgardClient) PutToUniversalClipboard(ctx context.Context, in *PutToUniversalClipboardInput, opts ...grpc.CallOption) (*PutToUniversalClipboardOutput, error) {
	out := new(PutToUniversalClipboardOutput)
	err := c.cc.Invoke(ctx, "/proto.Midgard/PutToUniversalClipboard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *midgardClient) AllocateURL(ctx context.Context, in *AllocateURLInput, opts ...grpc.CallOption) (*AllocateURLOutput, error) {
	out := new(AllocateURLOutput)
	err := c.cc.Invoke(ctx, "/proto.Midgard/AllocateURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MidgardServer is the server API for Midgard service.
type MidgardServer interface {
	Ping(context.Context, *PingInput) (*PingOutput, error)
	GetFromUniversalClipboard(context.Context, *GetFromUniversalClipboardInput) (*GetFromUniversalClipboardOutput, error)
	PutToUniversalClipboard(context.Context, *PutToUniversalClipboardInput) (*PutToUniversalClipboardOutput, error)
	AllocateURL(context.Context, *AllocateURLInput) (*AllocateURLOutput, error)
}

// UnimplementedMidgardServer can be embedded to have forward compatible implementations.
type UnimplementedMidgardServer struct {
}

func (*UnimplementedMidgardServer) Ping(context.Context, *PingInput) (*PingOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedMidgardServer) GetFromUniversalClipboard(context.Context, *GetFromUniversalClipboardInput) (*GetFromUniversalClipboardOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFromUniversalClipboard not implemented")
}
func (*UnimplementedMidgardServer) PutToUniversalClipboard(context.Context, *PutToUniversalClipboardInput) (*PutToUniversalClipboardOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutToUniversalClipboard not implemented")
}
func (*UnimplementedMidgardServer) AllocateURL(context.Context, *AllocateURLInput) (*AllocateURLOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllocateURL not implemented")
}

func RegisterMidgardServer(s *grpc.Server, srv MidgardServer) {
	s.RegisterService(&_Midgard_serviceDesc, srv)
}

func _Midgard_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MidgardServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Midgard/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MidgardServer).Ping(ctx, req.(*PingInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Midgard_GetFromUniversalClipboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFromUniversalClipboardInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MidgardServer).GetFromUniversalClipboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Midgard/GetFromUniversalClipboard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MidgardServer).GetFromUniversalClipboard(ctx, req.(*GetFromUniversalClipboardInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Midgard_PutToUniversalClipboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutToUniversalClipboardInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MidgardServer).PutToUniversalClipboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Midgard/PutToUniversalClipboard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MidgardServer).PutToUniversalClipboard(ctx, req.(*PutToUniversalClipboardInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Midgard_AllocateURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocateURLInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MidgardServer).AllocateURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Midgard/AllocateURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MidgardServer).AllocateURL(ctx, req.(*AllocateURLInput))
	}
	return interceptor(ctx, in, info, handler)
}

var _Midgard_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Midgard",
	HandlerType: (*MidgardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Midgard_Ping_Handler,
		},
		{
			MethodName: "GetFromUniversalClipboard",
			Handler:    _Midgard_GetFromUniversalClipboard_Handler,
		},
		{
			MethodName: "PutToUniversalClipboard",
			Handler:    _Midgard_PutToUniversalClipboard_Handler,
		},
		{
			MethodName: "AllocateURL",
			Handler:    _Midgard_AllocateURL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "midgard.proto",
}