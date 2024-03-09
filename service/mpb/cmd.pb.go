// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.3
// source: cmd.proto

package mpb

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

type MainCmd_Cmd int32

const (
	MainCmd_Error      MainCmd_Cmd = 0
	MainCmd_TCPGateway MainCmd_Cmd = 1
	MainCmd_Team       MainCmd_Cmd = 2
	MainCmd_Chat       MainCmd_Cmd = 3
	MainCmd_Item       MainCmd_Cmd = 4
	MainCmd_Mail       MainCmd_Cmd = 5
	MainCmd_Social     MainCmd_Cmd = 6
)

// Enum value maps for MainCmd_Cmd.
var (
	MainCmd_Cmd_name = map[int32]string{
		0: "Error",
		1: "TCPGateway",
		2: "Team",
		3: "Chat",
		4: "Item",
		5: "Mail",
		6: "Social",
	}
	MainCmd_Cmd_value = map[string]int32{
		"Error":      0,
		"TCPGateway": 1,
		"Team":       2,
		"Chat":       3,
		"Item":       4,
		"Mail":       5,
		"Social":     6,
	}
)

func (x MainCmd_Cmd) Enum() *MainCmd_Cmd {
	p := new(MainCmd_Cmd)
	*p = x
	return p
}

func (x MainCmd_Cmd) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MainCmd_Cmd) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_proto_enumTypes[0].Descriptor()
}

func (MainCmd_Cmd) Type() protoreflect.EnumType {
	return &file_cmd_proto_enumTypes[0]
}

func (x MainCmd_Cmd) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MainCmd_Cmd.Descriptor instead.
func (MainCmd_Cmd) EnumDescriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{0, 0}
}

type SubCmd_Error_Cmd int32

const (
	SubCmd_Error_None SubCmd_Error_Cmd = 0
)

// Enum value maps for SubCmd_Error_Cmd.
var (
	SubCmd_Error_Cmd_name = map[int32]string{
		0: "None",
	}
	SubCmd_Error_Cmd_value = map[string]int32{
		"None": 0,
	}
)

func (x SubCmd_Error_Cmd) Enum() *SubCmd_Error_Cmd {
	p := new(SubCmd_Error_Cmd)
	*p = x
	return p
}

func (x SubCmd_Error_Cmd) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SubCmd_Error_Cmd) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_proto_enumTypes[1].Descriptor()
}

func (SubCmd_Error_Cmd) Type() protoreflect.EnumType {
	return &file_cmd_proto_enumTypes[1]
}

func (x SubCmd_Error_Cmd) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SubCmd_Error_Cmd.Descriptor instead.
func (SubCmd_Error_Cmd) EnumDescriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{1, 0}
}

type SubCmd_TCPGateway_Cmd int32

const (
	SubCmd_TCPGateway_None      SubCmd_TCPGateway_Cmd = 0
	SubCmd_TCPGateway_HeartBeat SubCmd_TCPGateway_Cmd = 1
	SubCmd_TCPGateway_LoginTCP  SubCmd_TCPGateway_Cmd = 2
)

// Enum value maps for SubCmd_TCPGateway_Cmd.
var (
	SubCmd_TCPGateway_Cmd_name = map[int32]string{
		0: "None",
		1: "HeartBeat",
		2: "LoginTCP",
	}
	SubCmd_TCPGateway_Cmd_value = map[string]int32{
		"None":      0,
		"HeartBeat": 1,
		"LoginTCP":  2,
	}
)

func (x SubCmd_TCPGateway_Cmd) Enum() *SubCmd_TCPGateway_Cmd {
	p := new(SubCmd_TCPGateway_Cmd)
	*p = x
	return p
}

func (x SubCmd_TCPGateway_Cmd) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SubCmd_TCPGateway_Cmd) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_proto_enumTypes[2].Descriptor()
}

func (SubCmd_TCPGateway_Cmd) Type() protoreflect.EnumType {
	return &file_cmd_proto_enumTypes[2]
}

func (x SubCmd_TCPGateway_Cmd) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SubCmd_TCPGateway_Cmd.Descriptor instead.
func (SubCmd_TCPGateway_Cmd) EnumDescriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{2, 0}
}

type SubCmd_Team_Cmd int32

const (
	SubCmd_Team_None SubCmd_Team_Cmd = 0
)

// Enum value maps for SubCmd_Team_Cmd.
var (
	SubCmd_Team_Cmd_name = map[int32]string{
		0: "None",
	}
	SubCmd_Team_Cmd_value = map[string]int32{
		"None": 0,
	}
)

func (x SubCmd_Team_Cmd) Enum() *SubCmd_Team_Cmd {
	p := new(SubCmd_Team_Cmd)
	*p = x
	return p
}

func (x SubCmd_Team_Cmd) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SubCmd_Team_Cmd) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_proto_enumTypes[3].Descriptor()
}

func (SubCmd_Team_Cmd) Type() protoreflect.EnumType {
	return &file_cmd_proto_enumTypes[3]
}

func (x SubCmd_Team_Cmd) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SubCmd_Team_Cmd.Descriptor instead.
func (SubCmd_Team_Cmd) EnumDescriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{3, 0}
}

type SubCmd_Chat_Cmd int32

const (
	SubCmd_Chat_None SubCmd_Chat_Cmd = 0
)

// Enum value maps for SubCmd_Chat_Cmd.
var (
	SubCmd_Chat_Cmd_name = map[int32]string{
		0: "None",
	}
	SubCmd_Chat_Cmd_value = map[string]int32{
		"None": 0,
	}
)

func (x SubCmd_Chat_Cmd) Enum() *SubCmd_Chat_Cmd {
	p := new(SubCmd_Chat_Cmd)
	*p = x
	return p
}

func (x SubCmd_Chat_Cmd) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SubCmd_Chat_Cmd) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_proto_enumTypes[4].Descriptor()
}

func (SubCmd_Chat_Cmd) Type() protoreflect.EnumType {
	return &file_cmd_proto_enumTypes[4]
}

func (x SubCmd_Chat_Cmd) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SubCmd_Chat_Cmd.Descriptor instead.
func (SubCmd_Chat_Cmd) EnumDescriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{4, 0}
}

type SubCmd_Item_Cmd int32

const (
	SubCmd_Item_None   SubCmd_Item_Cmd = 0
	SubCmd_Item_Update SubCmd_Item_Cmd = 1
)

// Enum value maps for SubCmd_Item_Cmd.
var (
	SubCmd_Item_Cmd_name = map[int32]string{
		0: "None",
		1: "Update",
	}
	SubCmd_Item_Cmd_value = map[string]int32{
		"None":   0,
		"Update": 1,
	}
)

func (x SubCmd_Item_Cmd) Enum() *SubCmd_Item_Cmd {
	p := new(SubCmd_Item_Cmd)
	*p = x
	return p
}

func (x SubCmd_Item_Cmd) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SubCmd_Item_Cmd) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_proto_enumTypes[5].Descriptor()
}

func (SubCmd_Item_Cmd) Type() protoreflect.EnumType {
	return &file_cmd_proto_enumTypes[5]
}

func (x SubCmd_Item_Cmd) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SubCmd_Item_Cmd.Descriptor instead.
func (SubCmd_Item_Cmd) EnumDescriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{5, 0}
}

type SubCmd_Mail_Cmd int32

const (
	SubCmd_Mail_None     SubCmd_Mail_Cmd = 0
	SubCmd_Mail_SendMail SubCmd_Mail_Cmd = 1
)

// Enum value maps for SubCmd_Mail_Cmd.
var (
	SubCmd_Mail_Cmd_name = map[int32]string{
		0: "None",
		1: "SendMail",
	}
	SubCmd_Mail_Cmd_value = map[string]int32{
		"None":     0,
		"SendMail": 1,
	}
)

func (x SubCmd_Mail_Cmd) Enum() *SubCmd_Mail_Cmd {
	p := new(SubCmd_Mail_Cmd)
	*p = x
	return p
}

func (x SubCmd_Mail_Cmd) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SubCmd_Mail_Cmd) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_proto_enumTypes[6].Descriptor()
}

func (SubCmd_Mail_Cmd) Type() protoreflect.EnumType {
	return &file_cmd_proto_enumTypes[6]
}

func (x SubCmd_Mail_Cmd) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SubCmd_Mail_Cmd.Descriptor instead.
func (SubCmd_Mail_Cmd) EnumDescriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{6, 0}
}

type SubCmd_Social_Cmd int32

const (
	SubCmd_Social_None         SubCmd_Social_Cmd = 0
	SubCmd_Social_ApplyFriend  SubCmd_Social_Cmd = 1
	SubCmd_Social_CreateFriend SubCmd_Social_Cmd = 2
	SubCmd_Social_RemoveFriend SubCmd_Social_Cmd = 3
)

// Enum value maps for SubCmd_Social_Cmd.
var (
	SubCmd_Social_Cmd_name = map[int32]string{
		0: "None",
		1: "ApplyFriend",
		2: "CreateFriend",
		3: "RemoveFriend",
	}
	SubCmd_Social_Cmd_value = map[string]int32{
		"None":         0,
		"ApplyFriend":  1,
		"CreateFriend": 2,
		"RemoveFriend": 3,
	}
)

func (x SubCmd_Social_Cmd) Enum() *SubCmd_Social_Cmd {
	p := new(SubCmd_Social_Cmd)
	*p = x
	return p
}

func (x SubCmd_Social_Cmd) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SubCmd_Social_Cmd) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_proto_enumTypes[7].Descriptor()
}

func (SubCmd_Social_Cmd) Type() protoreflect.EnumType {
	return &file_cmd_proto_enumTypes[7]
}

func (x SubCmd_Social_Cmd) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SubCmd_Social_Cmd.Descriptor instead.
func (SubCmd_Social_Cmd) EnumDescriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{7, 0}
}

type MainCmd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MainCmd) Reset() {
	*x = MainCmd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MainCmd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MainCmd) ProtoMessage() {}

func (x *MainCmd) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MainCmd.ProtoReflect.Descriptor instead.
func (*MainCmd) Descriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{0}
}

type SubCmd_Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubCmd_Error) Reset() {
	*x = SubCmd_Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubCmd_Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubCmd_Error) ProtoMessage() {}

func (x *SubCmd_Error) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubCmd_Error.ProtoReflect.Descriptor instead.
func (*SubCmd_Error) Descriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{1}
}

type SubCmd_TCPGateway struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubCmd_TCPGateway) Reset() {
	*x = SubCmd_TCPGateway{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubCmd_TCPGateway) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubCmd_TCPGateway) ProtoMessage() {}

func (x *SubCmd_TCPGateway) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubCmd_TCPGateway.ProtoReflect.Descriptor instead.
func (*SubCmd_TCPGateway) Descriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{2}
}

type SubCmd_Team struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubCmd_Team) Reset() {
	*x = SubCmd_Team{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubCmd_Team) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubCmd_Team) ProtoMessage() {}

func (x *SubCmd_Team) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubCmd_Team.ProtoReflect.Descriptor instead.
func (*SubCmd_Team) Descriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{3}
}

type SubCmd_Chat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubCmd_Chat) Reset() {
	*x = SubCmd_Chat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubCmd_Chat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubCmd_Chat) ProtoMessage() {}

func (x *SubCmd_Chat) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubCmd_Chat.ProtoReflect.Descriptor instead.
func (*SubCmd_Chat) Descriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{4}
}

type SubCmd_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubCmd_Item) Reset() {
	*x = SubCmd_Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubCmd_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubCmd_Item) ProtoMessage() {}

func (x *SubCmd_Item) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubCmd_Item.ProtoReflect.Descriptor instead.
func (*SubCmd_Item) Descriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{5}
}

type SubCmd_Mail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubCmd_Mail) Reset() {
	*x = SubCmd_Mail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubCmd_Mail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubCmd_Mail) ProtoMessage() {}

func (x *SubCmd_Mail) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubCmd_Mail.ProtoReflect.Descriptor instead.
func (*SubCmd_Mail) Descriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{6}
}

type SubCmd_Social struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubCmd_Social) Reset() {
	*x = SubCmd_Social{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubCmd_Social) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubCmd_Social) ProtoMessage() {}

func (x *SubCmd_Social) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubCmd_Social.ProtoReflect.Descriptor instead.
func (*SubCmd_Social) Descriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{7}
}

var File_cmd_proto protoreflect.FileDescriptor

var file_cmd_proto_rawDesc = []byte{
	0x0a, 0x09, 0x63, 0x6d, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x70, 0x62,
	0x22, 0x5f, 0x0a, 0x07, 0x4d, 0x61, 0x69, 0x6e, 0x43, 0x6d, 0x64, 0x22, 0x54, 0x0a, 0x03, 0x43,
	0x6d, 0x64, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x00, 0x12, 0x0e, 0x0a,
	0x0a, 0x54, 0x43, 0x50, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x10, 0x01, 0x12, 0x08, 0x0a,
	0x04, 0x54, 0x65, 0x61, 0x6d, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x10,
	0x03, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x4d,
	0x61, 0x69, 0x6c, 0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x10,
	0x06, 0x22, 0x1f, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x43, 0x6d, 0x64, 0x5f, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x22, 0x0f, 0x0a, 0x03, 0x43, 0x6d, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65,
	0x10, 0x00, 0x22, 0x41, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x43, 0x6d, 0x64, 0x5f, 0x54, 0x43, 0x50,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x22, 0x2c, 0x0a, 0x03, 0x43, 0x6d, 0x64, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x42, 0x65, 0x61, 0x74, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x54, 0x43, 0x50, 0x10, 0x02, 0x22, 0x1e, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x43, 0x6d, 0x64, 0x5f,
	0x54, 0x65, 0x61, 0x6d, 0x22, 0x0f, 0x0a, 0x03, 0x43, 0x6d, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x6f, 0x6e, 0x65, 0x10, 0x00, 0x22, 0x1e, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x43, 0x6d, 0x64, 0x5f,
	0x43, 0x68, 0x61, 0x74, 0x22, 0x0f, 0x0a, 0x03, 0x43, 0x6d, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x6f, 0x6e, 0x65, 0x10, 0x00, 0x22, 0x2a, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x43, 0x6d, 0x64, 0x5f,
	0x49, 0x74, 0x65, 0x6d, 0x22, 0x1b, 0x0a, 0x03, 0x43, 0x6d, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x10,
	0x01, 0x22, 0x2c, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x43, 0x6d, 0x64, 0x5f, 0x4d, 0x61, 0x69, 0x6c,
	0x22, 0x1d, 0x0a, 0x03, 0x43, 0x6d, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x10, 0x01, 0x22,
	0x55, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x43, 0x6d, 0x64, 0x5f, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c,
	0x22, 0x44, 0x0a, 0x03, 0x43, 0x6d, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10,
	0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x10, 0x03, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x6d, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_proto_rawDescOnce sync.Once
	file_cmd_proto_rawDescData = file_cmd_proto_rawDesc
)

func file_cmd_proto_rawDescGZIP() []byte {
	file_cmd_proto_rawDescOnce.Do(func() {
		file_cmd_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_proto_rawDescData)
	})
	return file_cmd_proto_rawDescData
}

var file_cmd_proto_enumTypes = make([]protoimpl.EnumInfo, 8)
var file_cmd_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_cmd_proto_goTypes = []interface{}{
	(MainCmd_Cmd)(0),           // 0: mpb.MainCmd.Cmd
	(SubCmd_Error_Cmd)(0),      // 1: mpb.SubCmd_Error.Cmd
	(SubCmd_TCPGateway_Cmd)(0), // 2: mpb.SubCmd_TCPGateway.Cmd
	(SubCmd_Team_Cmd)(0),       // 3: mpb.SubCmd_Team.Cmd
	(SubCmd_Chat_Cmd)(0),       // 4: mpb.SubCmd_Chat.Cmd
	(SubCmd_Item_Cmd)(0),       // 5: mpb.SubCmd_Item.Cmd
	(SubCmd_Mail_Cmd)(0),       // 6: mpb.SubCmd_Mail.Cmd
	(SubCmd_Social_Cmd)(0),     // 7: mpb.SubCmd_Social.Cmd
	(*MainCmd)(nil),            // 8: mpb.MainCmd
	(*SubCmd_Error)(nil),       // 9: mpb.SubCmd_Error
	(*SubCmd_TCPGateway)(nil),  // 10: mpb.SubCmd_TCPGateway
	(*SubCmd_Team)(nil),        // 11: mpb.SubCmd_Team
	(*SubCmd_Chat)(nil),        // 12: mpb.SubCmd_Chat
	(*SubCmd_Item)(nil),        // 13: mpb.SubCmd_Item
	(*SubCmd_Mail)(nil),        // 14: mpb.SubCmd_Mail
	(*SubCmd_Social)(nil),      // 15: mpb.SubCmd_Social
}
var file_cmd_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmd_proto_init() }
func file_cmd_proto_init() {
	if File_cmd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MainCmd); i {
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
		file_cmd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubCmd_Error); i {
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
		file_cmd_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubCmd_TCPGateway); i {
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
		file_cmd_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubCmd_Team); i {
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
		file_cmd_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubCmd_Chat); i {
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
		file_cmd_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubCmd_Item); i {
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
		file_cmd_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubCmd_Mail); i {
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
		file_cmd_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubCmd_Social); i {
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
			RawDescriptor: file_cmd_proto_rawDesc,
			NumEnums:      8,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_proto_goTypes,
		DependencyIndexes: file_cmd_proto_depIdxs,
		EnumInfos:         file_cmd_proto_enumTypes,
		MessageInfos:      file_cmd_proto_msgTypes,
	}.Build()
	File_cmd_proto = out.File
	file_cmd_proto_rawDesc = nil
	file_cmd_proto_goTypes = nil
	file_cmd_proto_depIdxs = nil
}
