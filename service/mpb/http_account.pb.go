// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.3
// source: http_account.proto

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

// account
type CResTelegramLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *AccountInfo `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *CResTelegramLogin) Reset() {
	*x = CResTelegramLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CResTelegramLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CResTelegramLogin) ProtoMessage() {}

func (x *CResTelegramLogin) ProtoReflect() protoreflect.Message {
	mi := &file_http_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CResTelegramLogin.ProtoReflect.Descriptor instead.
func (*CResTelegramLogin) Descriptor() ([]byte, []int) {
	return file_http_account_proto_rawDescGZIP(), []int{0}
}

func (x *CResTelegramLogin) GetAccount() *AccountInfo {
	if x != nil {
		return x.Account
	}
	return nil
}

type CReqLoginTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CReqLoginTest) Reset() {
	*x = CReqLoginTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CReqLoginTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CReqLoginTest) ProtoMessage() {}

func (x *CReqLoginTest) ProtoReflect() protoreflect.Message {
	mi := &file_http_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CReqLoginTest.ProtoReflect.Descriptor instead.
func (*CReqLoginTest) Descriptor() ([]byte, []int) {
	return file_http_account_proto_rawDescGZIP(), []int{1}
}

func (x *CReqLoginTest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CResLoginTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account           *AccountInfo         `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Token             string               `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Energy            uint32               `protobuf:"varint,3,opt,name=energy,proto3" json:"energy,omitempty"`
	EnergyUpdateAt    int64                `protobuf:"varint,4,opt,name=energy_update_at,json=energyUpdateAt,proto3" json:"energy_update_at,omitempty"`
	BossDefeatHistory *BossDefeatHistory   `protobuf:"bytes,5,opt,name=boss_defeat_history,json=bossDefeatHistory,proto3" json:"boss_defeat_history,omitempty"`
	BuffCards         []*BuffCard          `protobuf:"bytes,6,rep,name=buff_cards,json=buffCards,proto3" json:"buff_cards,omitempty"`
	BuffCardStatus    EGame_BuffCardStatus `protobuf:"varint,7,opt,name=buff_card_status,json=buffCardStatus,proto3,enum=mpb.EGame_BuffCardStatus" json:"buff_card_status,omitempty"`
}

func (x *CResLoginTest) Reset() {
	*x = CResLoginTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CResLoginTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CResLoginTest) ProtoMessage() {}

func (x *CResLoginTest) ProtoReflect() protoreflect.Message {
	mi := &file_http_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CResLoginTest.ProtoReflect.Descriptor instead.
func (*CResLoginTest) Descriptor() ([]byte, []int) {
	return file_http_account_proto_rawDescGZIP(), []int{2}
}

func (x *CResLoginTest) GetAccount() *AccountInfo {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *CResLoginTest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CResLoginTest) GetEnergy() uint32 {
	if x != nil {
		return x.Energy
	}
	return 0
}

func (x *CResLoginTest) GetEnergyUpdateAt() int64 {
	if x != nil {
		return x.EnergyUpdateAt
	}
	return 0
}

func (x *CResLoginTest) GetBossDefeatHistory() *BossDefeatHistory {
	if x != nil {
		return x.BossDefeatHistory
	}
	return nil
}

func (x *CResLoginTest) GetBuffCards() []*BuffCard {
	if x != nil {
		return x.BuffCards
	}
	return nil
}

func (x *CResLoginTest) GetBuffCardStatus() EGame_BuffCardStatus {
	if x != nil {
		return x.BuffCardStatus
	}
	return EGame_BuffCardStatus_None
}

type CReqLoginByToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CReqLoginByToken) Reset() {
	*x = CReqLoginByToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CReqLoginByToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CReqLoginByToken) ProtoMessage() {}

func (x *CReqLoginByToken) ProtoReflect() protoreflect.Message {
	mi := &file_http_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CReqLoginByToken.ProtoReflect.Descriptor instead.
func (*CReqLoginByToken) Descriptor() ([]byte, []int) {
	return file_http_account_proto_rawDescGZIP(), []int{3}
}

func (x *CReqLoginByToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type CResLoginByToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account           *AccountInfo         `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Token             string               `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Energy            uint32               `protobuf:"varint,3,opt,name=energy,proto3" json:"energy,omitempty"`
	EnergyUpdateAt    int64                `protobuf:"varint,4,opt,name=energy_update_at,json=energyUpdateAt,proto3" json:"energy_update_at,omitempty"`
	BossDefeatHistory *BossDefeatHistory   `protobuf:"bytes,5,opt,name=boss_defeat_history,json=bossDefeatHistory,proto3" json:"boss_defeat_history,omitempty"`
	BuffCards         []*BuffCard          `protobuf:"bytes,6,rep,name=buff_cards,json=buffCards,proto3" json:"buff_cards,omitempty"`
	BuffCardStatus    EGame_BuffCardStatus `protobuf:"varint,7,opt,name=buff_card_status,json=buffCardStatus,proto3,enum=mpb.EGame_BuffCardStatus" json:"buff_card_status,omitempty"`
}

func (x *CResLoginByToken) Reset() {
	*x = CResLoginByToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CResLoginByToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CResLoginByToken) ProtoMessage() {}

func (x *CResLoginByToken) ProtoReflect() protoreflect.Message {
	mi := &file_http_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CResLoginByToken.ProtoReflect.Descriptor instead.
func (*CResLoginByToken) Descriptor() ([]byte, []int) {
	return file_http_account_proto_rawDescGZIP(), []int{4}
}

func (x *CResLoginByToken) GetAccount() *AccountInfo {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *CResLoginByToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CResLoginByToken) GetEnergy() uint32 {
	if x != nil {
		return x.Energy
	}
	return 0
}

func (x *CResLoginByToken) GetEnergyUpdateAt() int64 {
	if x != nil {
		return x.EnergyUpdateAt
	}
	return 0
}

func (x *CResLoginByToken) GetBossDefeatHistory() *BossDefeatHistory {
	if x != nil {
		return x.BossDefeatHistory
	}
	return nil
}

func (x *CResLoginByToken) GetBuffCards() []*BuffCard {
	if x != nil {
		return x.BuffCards
	}
	return nil
}

func (x *CResLoginByToken) GetBuffCardStatus() EGame_BuffCardStatus {
	if x != nil {
		return x.BuffCardStatus
	}
	return EGame_BuffCardStatus_None
}

var File_http_account_proto protoreflect.FileDescriptor

var file_http_account_proto_rawDesc = []byte{
	0x0a, 0x12, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x70, 0x62, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x11, 0x43, 0x52, 0x65, 0x73, 0x54,
	0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x2a, 0x0a, 0x07,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x6d, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x28, 0x0a, 0x0d, 0x43, 0x52, 0x65, 0x71,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0xce, 0x02, 0x0a, 0x0d, 0x43, 0x52, 0x65, 0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x54, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x12, 0x28,
	0x0a, 0x10, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x46, 0x0a, 0x13, 0x62, 0x6f, 0x73, 0x73,
	0x5f, 0x64, 0x65, 0x66, 0x65, 0x61, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x42, 0x6f, 0x73, 0x73,
	0x44, 0x65, 0x66, 0x65, 0x61, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x11, 0x62,
	0x6f, 0x73, 0x73, 0x44, 0x65, 0x66, 0x65, 0x61, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x2c, 0x0a, 0x0a, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x42, 0x75, 0x66, 0x66, 0x43,
	0x61, 0x72, 0x64, 0x52, 0x09, 0x62, 0x75, 0x66, 0x66, 0x43, 0x61, 0x72, 0x64, 0x73, 0x12, 0x43,
	0x0a, 0x10, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x45,
	0x47, 0x61, 0x6d, 0x65, 0x2e, 0x42, 0x75, 0x66, 0x66, 0x43, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x0e, 0x62, 0x75, 0x66, 0x66, 0x43, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x28, 0x0a, 0x10, 0x43, 0x52, 0x65, 0x71, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x42, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xd1, 0x02,
	0x0a, 0x10, 0x43, 0x52, 0x65, 0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x2a, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x12, 0x28, 0x0a, 0x10,
	0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x46, 0x0a, 0x13, 0x62, 0x6f, 0x73, 0x73, 0x5f, 0x64,
	0x65, 0x66, 0x65, 0x61, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x42, 0x6f, 0x73, 0x73, 0x44, 0x65,
	0x66, 0x65, 0x61, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x11, 0x62, 0x6f, 0x73,
	0x73, 0x44, 0x65, 0x66, 0x65, 0x61, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x2c,
	0x0a, 0x0a, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x42, 0x75, 0x66, 0x66, 0x43, 0x61, 0x72,
	0x64, 0x52, 0x09, 0x62, 0x75, 0x66, 0x66, 0x43, 0x61, 0x72, 0x64, 0x73, 0x12, 0x43, 0x0a, 0x10,
	0x62, 0x75, 0x66, 0x66, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x45, 0x47, 0x61,
	0x6d, 0x65, 0x2e, 0x42, 0x75, 0x66, 0x66, 0x43, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x0e, 0x62, 0x75, 0x66, 0x66, 0x43, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x6d, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_http_account_proto_rawDescOnce sync.Once
	file_http_account_proto_rawDescData = file_http_account_proto_rawDesc
)

func file_http_account_proto_rawDescGZIP() []byte {
	file_http_account_proto_rawDescOnce.Do(func() {
		file_http_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_http_account_proto_rawDescData)
	})
	return file_http_account_proto_rawDescData
}

var file_http_account_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_http_account_proto_goTypes = []interface{}{
	(*CResTelegramLogin)(nil), // 0: mpb.CResTelegramLogin
	(*CReqLoginTest)(nil),     // 1: mpb.CReqLoginTest
	(*CResLoginTest)(nil),     // 2: mpb.CResLoginTest
	(*CReqLoginByToken)(nil),  // 3: mpb.CReqLoginByToken
	(*CResLoginByToken)(nil),  // 4: mpb.CResLoginByToken
	(*AccountInfo)(nil),       // 5: mpb.AccountInfo
	(*BossDefeatHistory)(nil), // 6: mpb.BossDefeatHistory
	(*BuffCard)(nil),          // 7: mpb.BuffCard
	(EGame_BuffCardStatus)(0), // 8: mpb.EGame.BuffCardStatus
}
var file_http_account_proto_depIdxs = []int32{
	5, // 0: mpb.CResTelegramLogin.account:type_name -> mpb.AccountInfo
	5, // 1: mpb.CResLoginTest.account:type_name -> mpb.AccountInfo
	6, // 2: mpb.CResLoginTest.boss_defeat_history:type_name -> mpb.BossDefeatHistory
	7, // 3: mpb.CResLoginTest.buff_cards:type_name -> mpb.BuffCard
	8, // 4: mpb.CResLoginTest.buff_card_status:type_name -> mpb.EGame.BuffCardStatus
	5, // 5: mpb.CResLoginByToken.account:type_name -> mpb.AccountInfo
	6, // 6: mpb.CResLoginByToken.boss_defeat_history:type_name -> mpb.BossDefeatHistory
	7, // 7: mpb.CResLoginByToken.buff_cards:type_name -> mpb.BuffCard
	8, // 8: mpb.CResLoginByToken.buff_card_status:type_name -> mpb.EGame.BuffCardStatus
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_http_account_proto_init() }
func file_http_account_proto_init() {
	if File_http_account_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_http_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CResTelegramLogin); i {
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
		file_http_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CReqLoginTest); i {
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
		file_http_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CResLoginTest); i {
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
		file_http_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CReqLoginByToken); i {
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
		file_http_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CResLoginByToken); i {
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
			RawDescriptor: file_http_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_http_account_proto_goTypes,
		DependencyIndexes: file_http_account_proto_depIdxs,
		MessageInfos:      file_http_account_proto_msgTypes,
	}.Build()
	File_http_account_proto = out.File
	file_http_account_proto_rawDesc = nil
	file_http_account_proto_goTypes = nil
	file_http_account_proto_depIdxs = nil
}
