// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: rw.proto

package model

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

type Cmd int32

const (
	Cmd_LOGIN Cmd = 0
	Cmd_GET   Cmd = 1
	Cmd_SET   Cmd = 2
)

// Enum value maps for Cmd.
var (
	Cmd_name = map[int32]string{
		0: "LOGIN",
		1: "GET",
		2: "SET",
	}
	Cmd_value = map[string]int32{
		"LOGIN": 0,
		"GET":   1,
		"SET":   2,
	}
)

func (x Cmd) Enum() *Cmd {
	p := new(Cmd)
	*p = x
	return p
}

func (x Cmd) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Cmd) Descriptor() protoreflect.EnumDescriptor {
	return file_rw_proto_enumTypes[0].Descriptor()
}

func (Cmd) Type() protoreflect.EnumType {
	return &file_rw_proto_enumTypes[0]
}

func (x Cmd) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Cmd.Descriptor instead.
func (Cmd) EnumDescriptor() ([]byte, []int) {
	return file_rw_proto_rawDescGZIP(), []int{0}
}

type RetCode int32

const (
	RetCode_SUCCESS       RetCode = 0
	RetCode_ERR_TOKEN     RetCode = -1
	RetCode_ERR_FREQUENCY RetCode = -2
	RetCode_ERR_REQ_PARA  RetCode = -3
)

// Enum value maps for RetCode.
var (
	RetCode_name = map[int32]string{
		0:  "SUCCESS",
		-1: "ERR_TOKEN",
		-2: "ERR_FREQUENCY",
		-3: "ERR_REQ_PARA",
	}
	RetCode_value = map[string]int32{
		"SUCCESS":       0,
		"ERR_TOKEN":     -1,
		"ERR_FREQUENCY": -2,
		"ERR_REQ_PARA":  -3,
	}
)

func (x RetCode) Enum() *RetCode {
	p := new(RetCode)
	*p = x
	return p
}

func (x RetCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RetCode) Descriptor() protoreflect.EnumDescriptor {
	return file_rw_proto_enumTypes[1].Descriptor()
}

func (RetCode) Type() protoreflect.EnumType {
	return &file_rw_proto_enumTypes[1]
}

func (x RetCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RetCode.Descriptor instead.
func (RetCode) EnumDescriptor() ([]byte, []int) {
	return file_rw_proto_rawDescGZIP(), []int{1}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seq *uint64 `protobuf:"varint,1,opt,name=seq,proto3,oneof" json:"seq,omitempty"`
	Cmd *Cmd    `protobuf:"varint,2,opt,name=cmd,proto3,enum=rw.Cmd,oneof" json:"cmd,omitempty"`
	Buf []byte  `protobuf:"bytes,3,opt,name=buf,proto3,oneof" json:"buf,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rw_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_rw_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_rw_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetSeq() uint64 {
	if x != nil && x.Seq != nil {
		return *x.Seq
	}
	return 0
}

func (x *Request) GetCmd() Cmd {
	if x != nil && x.Cmd != nil {
		return *x.Cmd
	}
	return Cmd_LOGIN
}

func (x *Request) GetBuf() []byte {
	if x != nil {
		return x.Buf
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seq *uint64 `protobuf:"varint,1,opt,name=seq,proto3,oneof" json:"seq,omitempty"`
	Cmd *Cmd    `protobuf:"varint,2,opt,name=cmd,proto3,enum=rw.Cmd,oneof" json:"cmd,omitempty"`
	Buf []byte  `protobuf:"bytes,3,opt,name=buf,proto3,oneof" json:"buf,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rw_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_rw_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_rw_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetSeq() uint64 {
	if x != nil && x.Seq != nil {
		return *x.Seq
	}
	return 0
}

func (x *Response) GetCmd() Cmd {
	if x != nil && x.Cmd != nil {
		return *x.Cmd
	}
	return Cmd_LOGIN
}

func (x *Response) GetBuf() []byte {
	if x != nil {
		return x.Buf
	}
	return nil
}

type CommonHead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Token *string `protobuf:"bytes,2,opt,name=token,proto3,oneof" json:"token,omitempty"`
}

func (x *CommonHead) Reset() {
	*x = CommonHead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rw_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonHead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonHead) ProtoMessage() {}

func (x *CommonHead) ProtoReflect() protoreflect.Message {
	mi := &file_rw_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonHead.ProtoReflect.Descriptor instead.
func (*CommonHead) Descriptor() ([]byte, []int) {
	return file_rw_proto_rawDescGZIP(), []int{2}
}

func (x *CommonHead) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *CommonHead) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Head *CommonHead `protobuf:"bytes,1,opt,name=head,proto3,oneof" json:"head,omitempty"`
	Key  *string     `protobuf:"bytes,2,opt,name=key,proto3,oneof" json:"key,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rw_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rw_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_rw_proto_rawDescGZIP(), []int{3}
}

func (x *GetRequest) GetHead() *CommonHead {
	if x != nil {
		return x.Head
	}
	return nil
}

func (x *GetRequest) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret *int32  `protobuf:"varint,1,opt,name=ret,proto3,oneof" json:"ret,omitempty"`
	Val *string `protobuf:"bytes,2,opt,name=val,proto3,oneof" json:"val,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rw_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rw_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_rw_proto_rawDescGZIP(), []int{4}
}

func (x *GetResponse) GetRet() int32 {
	if x != nil && x.Ret != nil {
		return *x.Ret
	}
	return 0
}

func (x *GetResponse) GetVal() string {
	if x != nil && x.Val != nil {
		return *x.Val
	}
	return ""
}

type SetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Head *CommonHead `protobuf:"bytes,1,opt,name=head,proto3,oneof" json:"head,omitempty"`
	Key  *string     `protobuf:"bytes,2,opt,name=key,proto3,oneof" json:"key,omitempty"`
	Val  *string     `protobuf:"bytes,3,opt,name=val,proto3,oneof" json:"val,omitempty"`
}

func (x *SetRequest) Reset() {
	*x = SetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rw_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRequest) ProtoMessage() {}

func (x *SetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rw_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRequest.ProtoReflect.Descriptor instead.
func (*SetRequest) Descriptor() ([]byte, []int) {
	return file_rw_proto_rawDescGZIP(), []int{5}
}

func (x *SetRequest) GetHead() *CommonHead {
	if x != nil {
		return x.Head
	}
	return nil
}

func (x *SetRequest) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *SetRequest) GetVal() string {
	if x != nil && x.Val != nil {
		return *x.Val
	}
	return ""
}

type SetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret *int32 `protobuf:"varint,1,opt,name=ret,proto3,oneof" json:"ret,omitempty"`
}

func (x *SetResponse) Reset() {
	*x = SetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rw_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetResponse) ProtoMessage() {}

func (x *SetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rw_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetResponse.ProtoReflect.Descriptor instead.
func (*SetResponse) Descriptor() ([]byte, []int) {
	return file_rw_proto_rawDescGZIP(), []int{6}
}

func (x *SetResponse) GetRet() int32 {
	if x != nil && x.Ret != nil {
		return *x.Ret
	}
	return 0
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Key        *string `protobuf:"bytes,2,opt,name=key,proto3,oneof" json:"key,omitempty"`
	ReadTimes  *uint32 `protobuf:"varint,3,opt,name=readTimes,proto3,oneof" json:"readTimes,omitempty"`
	WriteTimes *uint32 `protobuf:"varint,4,opt,name=writeTimes,proto3,oneof" json:"writeTimes,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rw_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rw_proto_msgTypes[7]
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
	return file_rw_proto_rawDescGZIP(), []int{7}
}

func (x *LoginRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *LoginRequest) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *LoginRequest) GetReadTimes() uint32 {
	if x != nil && x.ReadTimes != nil {
		return *x.ReadTimes
	}
	return 0
}

func (x *LoginRequest) GetWriteTimes() uint32 {
	if x != nil && x.WriteTimes != nil {
		return *x.WriteTimes
	}
	return 0
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret   *int32  `protobuf:"varint,1,opt,name=ret,proto3,oneof" json:"ret,omitempty"`
	Token *string `protobuf:"bytes,2,opt,name=token,proto3,oneof" json:"token,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rw_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rw_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_rw_proto_rawDescGZIP(), []int{8}
}

func (x *LoginResponse) GetRet() int32 {
	if x != nil && x.Ret != nil {
		return *x.Ret
	}
	return 0
}

func (x *LoginResponse) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

var File_rw_proto protoreflect.FileDescriptor

var file_rw_proto_rawDesc = []byte{
	0x0a, 0x08, 0x72, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x72, 0x77, 0x22, 0x6f,
	0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x03, 0x73, 0x65, 0x71,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x03, 0x73, 0x65, 0x71, 0x88, 0x01, 0x01,
	0x12, 0x1e, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e,
	0x72, 0x77, 0x2e, 0x43, 0x6d, 0x64, 0x48, 0x01, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x15, 0x0a, 0x03, 0x62, 0x75, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x02, 0x52,
	0x03, 0x62, 0x75, 0x66, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x73, 0x65, 0x71, 0x42,
	0x06, 0x0a, 0x04, 0x5f, 0x63, 0x6d, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x62, 0x75, 0x66, 0x22,
	0x70, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x03, 0x73,
	0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x03, 0x73, 0x65, 0x71, 0x88,
	0x01, 0x01, 0x12, 0x1e, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x07, 0x2e, 0x72, 0x77, 0x2e, 0x43, 0x6d, 0x64, 0x48, 0x01, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x62, 0x75, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x48,
	0x02, 0x52, 0x03, 0x62, 0x75, 0x66, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x73, 0x65,
	0x71, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x63, 0x6d, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x62, 0x75,
	0x66, 0x22, 0x53, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x12,
	0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x5d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x68, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x77, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x48, 0x65,
	0x61, 0x64, 0x48, 0x00, 0x52, 0x04, 0x68, 0x65, 0x61, 0x64, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x42, 0x06, 0x0a,
	0x04, 0x5f, 0x6b, 0x65, 0x79, 0x22, 0x4b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x00, 0x52, 0x03, 0x72, 0x65, 0x74, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x76,
	0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x88,
	0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x72, 0x65, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x76,
	0x61, 0x6c, 0x22, 0x7c, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x27, 0x0a, 0x04, 0x68, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x72, 0x77, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x48, 0x00,
	0x52, 0x04, 0x68, 0x65, 0x61, 0x64, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x88, 0x01, 0x01,
	0x12, 0x15, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52,
	0x03, 0x76, 0x61, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x68, 0x65, 0x61, 0x64,
	0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6b, 0x65, 0x79, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x76, 0x61, 0x6c,
	0x22, 0x2c, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x15, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x03,
	0x72, 0x65, 0x74, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x72, 0x65, 0x74, 0x22, 0xb4,
	0x01, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x88, 0x01, 0x01, 0x12,
	0x21, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x48, 0x02, 0x52, 0x09, 0x72, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x88,
	0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x03, 0x52, 0x0a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6b, 0x65, 0x79, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x72, 0x65, 0x61,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x22, 0x53, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x03, 0x72, 0x65, 0x74, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x72, 0x65, 0x74,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2a, 0x22, 0x0a, 0x03, 0x43, 0x6d,
	0x64, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03,
	0x47, 0x45, 0x54, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x45, 0x54, 0x10, 0x02, 0x2a, 0x65,
	0x0a, 0x07, 0x52, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43,
	0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x09, 0x45, 0x52, 0x52, 0x5f, 0x54, 0x4f,
	0x4b, 0x45, 0x4e, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x1a,
	0x0a, 0x0d, 0x45, 0x52, 0x52, 0x5f, 0x46, 0x52, 0x45, 0x51, 0x55, 0x45, 0x4e, 0x43, 0x59, 0x10,
	0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x19, 0x0a, 0x0c, 0x45, 0x52,
	0x52, 0x5f, 0x52, 0x45, 0x51, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x10, 0xfd, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rw_proto_rawDescOnce sync.Once
	file_rw_proto_rawDescData = file_rw_proto_rawDesc
)

func file_rw_proto_rawDescGZIP() []byte {
	file_rw_proto_rawDescOnce.Do(func() {
		file_rw_proto_rawDescData = protoimpl.X.CompressGZIP(file_rw_proto_rawDescData)
	})
	return file_rw_proto_rawDescData
}

var file_rw_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_rw_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_rw_proto_goTypes = []interface{}{
	(Cmd)(0),              // 0: rw.Cmd
	(RetCode)(0),          // 1: rw.RetCode
	(*Request)(nil),       // 2: rw.Request
	(*Response)(nil),      // 3: rw.Response
	(*CommonHead)(nil),    // 4: rw.CommonHead
	(*GetRequest)(nil),    // 5: rw.GetRequest
	(*GetResponse)(nil),   // 6: rw.GetResponse
	(*SetRequest)(nil),    // 7: rw.SetRequest
	(*SetResponse)(nil),   // 8: rw.SetResponse
	(*LoginRequest)(nil),  // 9: rw.LoginRequest
	(*LoginResponse)(nil), // 10: rw.LoginResponse
}
var file_rw_proto_depIdxs = []int32{
	0, // 0: rw.Request.cmd:type_name -> rw.Cmd
	0, // 1: rw.Response.cmd:type_name -> rw.Cmd
	4, // 2: rw.GetRequest.head:type_name -> rw.CommonHead
	4, // 3: rw.SetRequest.head:type_name -> rw.CommonHead
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_rw_proto_init() }
func file_rw_proto_init() {
	if File_rw_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rw_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_rw_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_rw_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonHead); i {
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
		file_rw_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_rw_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_rw_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRequest); i {
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
		file_rw_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetResponse); i {
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
		file_rw_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_rw_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
	file_rw_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_rw_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_rw_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_rw_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_rw_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_rw_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_rw_proto_msgTypes[6].OneofWrappers = []interface{}{}
	file_rw_proto_msgTypes[7].OneofWrappers = []interface{}{}
	file_rw_proto_msgTypes[8].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rw_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rw_proto_goTypes,
		DependencyIndexes: file_rw_proto_depIdxs,
		EnumInfos:         file_rw_proto_enumTypes,
		MessageInfos:      file_rw_proto_msgTypes,
	}.Build()
	File_rw_proto = out.File
	file_rw_proto_rawDesc = nil
	file_rw_proto_goTypes = nil
	file_rw_proto_depIdxs = nil
}