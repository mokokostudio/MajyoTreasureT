// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.3
// source: grpc_account.proto

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

type ResLoginTest struct {
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

func (x *ResLoginTest) Reset() {
	*x = ResLoginTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResLoginTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResLoginTest) ProtoMessage() {}

func (x *ResLoginTest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResLoginTest.ProtoReflect.Descriptor instead.
func (*ResLoginTest) Descriptor() ([]byte, []int) {
	return file_grpc_account_proto_rawDescGZIP(), []int{0}
}

func (x *ResLoginTest) GetAccount() *AccountInfo {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *ResLoginTest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ResLoginTest) GetEnergy() uint32 {
	if x != nil {
		return x.Energy
	}
	return 0
}

func (x *ResLoginTest) GetEnergyUpdateAt() int64 {
	if x != nil {
		return x.EnergyUpdateAt
	}
	return 0
}

func (x *ResLoginTest) GetBossDefeatHistory() *BossDefeatHistory {
	if x != nil {
		return x.BossDefeatHistory
	}
	return nil
}

func (x *ResLoginTest) GetBuffCards() []*BuffCard {
	if x != nil {
		return x.BuffCards
	}
	return nil
}

func (x *ResLoginTest) GetBuffCardStatus() EGame_BuffCardStatus {
	if x != nil {
		return x.BuffCardStatus
	}
	return EGame_BuffCardStatus_None
}

type ReqGenerateLoginToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TgId         uint64 `protobuf:"varint,1,opt,name=tg_id,json=tgId,proto3" json:"tg_id,omitempty"`
	FirstName    string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName     string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	LanguageCode string `protobuf:"bytes,4,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
}

func (x *ReqGenerateLoginToken) Reset() {
	*x = ReqGenerateLoginToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqGenerateLoginToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqGenerateLoginToken) ProtoMessage() {}

func (x *ReqGenerateLoginToken) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqGenerateLoginToken.ProtoReflect.Descriptor instead.
func (*ReqGenerateLoginToken) Descriptor() ([]byte, []int) {
	return file_grpc_account_proto_rawDescGZIP(), []int{1}
}

func (x *ReqGenerateLoginToken) GetTgId() uint64 {
	if x != nil {
		return x.TgId
	}
	return 0
}

func (x *ReqGenerateLoginToken) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *ReqGenerateLoginToken) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *ReqGenerateLoginToken) GetLanguageCode() string {
	if x != nil {
		return x.LanguageCode
	}
	return ""
}

type ResGenerateLoginToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ResGenerateLoginToken) Reset() {
	*x = ResGenerateLoginToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResGenerateLoginToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResGenerateLoginToken) ProtoMessage() {}

func (x *ResGenerateLoginToken) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResGenerateLoginToken.ProtoReflect.Descriptor instead.
func (*ResGenerateLoginToken) Descriptor() ([]byte, []int) {
	return file_grpc_account_proto_rawDescGZIP(), []int{2}
}

func (x *ResGenerateLoginToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ReqLoginByToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ReqLoginByToken) Reset() {
	*x = ReqLoginByToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqLoginByToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqLoginByToken) ProtoMessage() {}

func (x *ReqLoginByToken) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqLoginByToken.ProtoReflect.Descriptor instead.
func (*ReqLoginByToken) Descriptor() ([]byte, []int) {
	return file_grpc_account_proto_rawDescGZIP(), []int{3}
}

func (x *ReqLoginByToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ResLoginByToken struct {
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

func (x *ResLoginByToken) Reset() {
	*x = ResLoginByToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResLoginByToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResLoginByToken) ProtoMessage() {}

func (x *ResLoginByToken) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResLoginByToken.ProtoReflect.Descriptor instead.
func (*ResLoginByToken) Descriptor() ([]byte, []int) {
	return file_grpc_account_proto_rawDescGZIP(), []int{4}
}

func (x *ResLoginByToken) GetAccount() *AccountInfo {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *ResLoginByToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ResLoginByToken) GetEnergy() uint32 {
	if x != nil {
		return x.Energy
	}
	return 0
}

func (x *ResLoginByToken) GetEnergyUpdateAt() int64 {
	if x != nil {
		return x.EnergyUpdateAt
	}
	return 0
}

func (x *ResLoginByToken) GetBossDefeatHistory() *BossDefeatHistory {
	if x != nil {
		return x.BossDefeatHistory
	}
	return nil
}

func (x *ResLoginByToken) GetBuffCards() []*BuffCard {
	if x != nil {
		return x.BuffCards
	}
	return nil
}

func (x *ResLoginByToken) GetBuffCardStatus() EGame_BuffCardStatus {
	if x != nil {
		return x.BuffCardStatus
	}
	return EGame_BuffCardStatus_None
}

type ReqGetAccountByTGUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TgId         uint64 `protobuf:"varint,1,opt,name=tg_id,json=tgId,proto3" json:"tg_id,omitempty"`
	FirstName    string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName     string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	LanguageCode string `protobuf:"bytes,4,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
}

func (x *ReqGetAccountByTGUser) Reset() {
	*x = ReqGetAccountByTGUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_account_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqGetAccountByTGUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqGetAccountByTGUser) ProtoMessage() {}

func (x *ReqGetAccountByTGUser) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_account_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqGetAccountByTGUser.ProtoReflect.Descriptor instead.
func (*ReqGetAccountByTGUser) Descriptor() ([]byte, []int) {
	return file_grpc_account_proto_rawDescGZIP(), []int{5}
}

func (x *ReqGetAccountByTGUser) GetTgId() uint64 {
	if x != nil {
		return x.TgId
	}
	return 0
}

func (x *ReqGetAccountByTGUser) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *ReqGetAccountByTGUser) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *ReqGetAccountByTGUser) GetLanguageCode() string {
	if x != nil {
		return x.LanguageCode
	}
	return ""
}

type ResGetAccountByTGUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account        *AccountInfo `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Energy         uint32       `protobuf:"varint,2,opt,name=energy,proto3" json:"energy,omitempty"`
	EnergyUpdateAt int64        `protobuf:"varint,3,opt,name=energy_update_at,json=energyUpdateAt,proto3" json:"energy_update_at,omitempty"`
}

func (x *ResGetAccountByTGUser) Reset() {
	*x = ResGetAccountByTGUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_account_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResGetAccountByTGUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResGetAccountByTGUser) ProtoMessage() {}

func (x *ResGetAccountByTGUser) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_account_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResGetAccountByTGUser.ProtoReflect.Descriptor instead.
func (*ResGetAccountByTGUser) Descriptor() ([]byte, []int) {
	return file_grpc_account_proto_rawDescGZIP(), []int{6}
}

func (x *ResGetAccountByTGUser) GetAccount() *AccountInfo {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *ResGetAccountByTGUser) GetEnergy() uint32 {
	if x != nil {
		return x.Energy
	}
	return 0
}

func (x *ResGetAccountByTGUser) GetEnergyUpdateAt() int64 {
	if x != nil {
		return x.EnergyUpdateAt
	}
	return 0
}

type ResGetAccountByUserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *AccountInfo `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *ResGetAccountByUserId) Reset() {
	*x = ResGetAccountByUserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_account_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResGetAccountByUserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResGetAccountByUserId) ProtoMessage() {}

func (x *ResGetAccountByUserId) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_account_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResGetAccountByUserId.ProtoReflect.Descriptor instead.
func (*ResGetAccountByUserId) Descriptor() ([]byte, []int) {
	return file_grpc_account_proto_rawDescGZIP(), []int{7}
}

func (x *ResGetAccountByUserId) GetAccount() *AccountInfo {
	if x != nil {
		return x.Account
	}
	return nil
}

type ReqSetAccountTGLan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TgId         uint64 `protobuf:"varint,1,opt,name=tg_id,json=tgId,proto3" json:"tg_id,omitempty"`
	LanguageCode string `protobuf:"bytes,2,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
}

func (x *ReqSetAccountTGLan) Reset() {
	*x = ReqSetAccountTGLan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_account_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqSetAccountTGLan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqSetAccountTGLan) ProtoMessage() {}

func (x *ReqSetAccountTGLan) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_account_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqSetAccountTGLan.ProtoReflect.Descriptor instead.
func (*ReqSetAccountTGLan) Descriptor() ([]byte, []int) {
	return file_grpc_account_proto_rawDescGZIP(), []int{8}
}

func (x *ReqSetAccountTGLan) GetTgId() uint64 {
	if x != nil {
		return x.TgId
	}
	return 0
}

func (x *ReqSetAccountTGLan) GetLanguageCode() string {
	if x != nil {
		return x.LanguageCode
	}
	return ""
}

type ResBatchGetAccountsByUserIds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accounts map[uint64]*AccountInfo `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ResBatchGetAccountsByUserIds) Reset() {
	*x = ResBatchGetAccountsByUserIds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_account_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResBatchGetAccountsByUserIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResBatchGetAccountsByUserIds) ProtoMessage() {}

func (x *ResBatchGetAccountsByUserIds) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_account_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResBatchGetAccountsByUserIds.ProtoReflect.Descriptor instead.
func (*ResBatchGetAccountsByUserIds) Descriptor() ([]byte, []int) {
	return file_grpc_account_proto_rawDescGZIP(), []int{9}
}

func (x *ResBatchGetAccountsByUserIds) GetAccounts() map[uint64]*AccountInfo {
	if x != nil {
		return x.Accounts
	}
	return nil
}

var File_grpc_account_proto protoreflect.FileDescriptor

var file_grpc_account_proto_rawDesc = []byte{
	0x0a, 0x12, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x70, 0x62, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x02, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x70, 0x62, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e,
	0x65, 0x72, 0x67, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x65, 0x6e, 0x65, 0x72,
	0x67, 0x79, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x5f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x65, 0x6e,
	0x65, 0x72, 0x67, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x46, 0x0a, 0x13,
	0x62, 0x6f, 0x73, 0x73, 0x5f, 0x64, 0x65, 0x66, 0x65, 0x61, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x70, 0x62, 0x2e,
	0x42, 0x6f, 0x73, 0x73, 0x44, 0x65, 0x66, 0x65, 0x61, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x11, 0x62, 0x6f, 0x73, 0x73, 0x44, 0x65, 0x66, 0x65, 0x61, 0x74, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x12, 0x2c, 0x0a, 0x0a, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x63, 0x61, 0x72,
	0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x42,
	0x75, 0x66, 0x66, 0x43, 0x61, 0x72, 0x64, 0x52, 0x09, 0x62, 0x75, 0x66, 0x66, 0x43, 0x61, 0x72,
	0x64, 0x73, 0x12, 0x43, 0x0a, 0x10, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6d,
	0x70, 0x62, 0x2e, 0x45, 0x47, 0x61, 0x6d, 0x65, 0x2e, 0x42, 0x75, 0x66, 0x66, 0x43, 0x61, 0x72,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0e, 0x62, 0x75, 0x66, 0x66, 0x43, 0x61, 0x72,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x8d, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x71, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x13, 0x0a, 0x05, 0x74, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x04, 0x74, 0x67, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x2d, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x27, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x42, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0xd0, 0x02, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x2a, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x12, 0x28, 0x0a,
	0x10, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x46, 0x0a, 0x13, 0x62, 0x6f, 0x73, 0x73, 0x5f,
	0x64, 0x65, 0x66, 0x65, 0x61, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x42, 0x6f, 0x73, 0x73, 0x44,
	0x65, 0x66, 0x65, 0x61, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x11, 0x62, 0x6f,
	0x73, 0x73, 0x44, 0x65, 0x66, 0x65, 0x61, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12,
	0x2c, 0x0a, 0x0a, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x42, 0x75, 0x66, 0x66, 0x43, 0x61,
	0x72, 0x64, 0x52, 0x09, 0x62, 0x75, 0x66, 0x66, 0x43, 0x61, 0x72, 0x64, 0x73, 0x12, 0x43, 0x0a,
	0x10, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x45, 0x47,
	0x61, 0x6d, 0x65, 0x2e, 0x42, 0x75, 0x66, 0x66, 0x43, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x0e, 0x62, 0x75, 0x66, 0x66, 0x43, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x8d, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x71, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x54, 0x47, 0x55, 0x73, 0x65, 0x72, 0x12, 0x13, 0x0a, 0x05,
	0x74, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x67, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x22, 0x85, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x54, 0x47, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x07,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x6d, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x65, 0x72,
	0x67, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79,
	0x12, 0x28, 0x0a, 0x10, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x65, 0x6e, 0x65, 0x72,
	0x67, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x22, 0x43, 0x0a, 0x15, 0x52, 0x65,
	0x73, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x4e, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x53, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x54, 0x47, 0x4c, 0x61, 0x6e, 0x12, 0x13, 0x0a, 0x05, 0x74, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x67, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x22,
	0xba, 0x01, 0x0a, 0x1c, 0x52, 0x65, 0x73, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73,
	0x12, 0x4b, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x1a, 0x4d, 0x0a,
	0x0d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xe3, 0x03, 0x0a,
	0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2e, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x2e, 0x6d,
	0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x11, 0x2e, 0x6d,
	0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x12,
	0x4c, 0x0a, 0x12, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x1a, 0x1a, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x3a, 0x0a,
	0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x2e,
	0x6d, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x1a, 0x14, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x42, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x4c, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x54, 0x47, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x1a, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x42, 0x79, 0x54, 0x47, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x1a, 0x2e, 0x6d, 0x70,
	0x62, 0x2e, 0x52, 0x65, 0x73, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42,
	0x79, 0x54, 0x47, 0x55, 0x73, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x2e,
	0x6d, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x1a, 0x2e,
	0x6d, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x4f, 0x0a, 0x19, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x0f, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x1a, 0x21, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x73, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x36, 0x0a, 0x0f, 0x53, 0x65,
	0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x47, 0x4c, 0x61, 0x6e, 0x12, 0x17, 0x2e,
	0x6d, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x53, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x54, 0x47, 0x4c, 0x61, 0x6e, 0x1a, 0x0a, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x6d, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_grpc_account_proto_rawDescOnce sync.Once
	file_grpc_account_proto_rawDescData = file_grpc_account_proto_rawDesc
)

func file_grpc_account_proto_rawDescGZIP() []byte {
	file_grpc_account_proto_rawDescOnce.Do(func() {
		file_grpc_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_account_proto_rawDescData)
	})
	return file_grpc_account_proto_rawDescData
}

var file_grpc_account_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_grpc_account_proto_goTypes = []interface{}{
	(*ResLoginTest)(nil),                 // 0: mpb.ResLoginTest
	(*ReqGenerateLoginToken)(nil),        // 1: mpb.ReqGenerateLoginToken
	(*ResGenerateLoginToken)(nil),        // 2: mpb.ResGenerateLoginToken
	(*ReqLoginByToken)(nil),              // 3: mpb.ReqLoginByToken
	(*ResLoginByToken)(nil),              // 4: mpb.ResLoginByToken
	(*ReqGetAccountByTGUser)(nil),        // 5: mpb.ReqGetAccountByTGUser
	(*ResGetAccountByTGUser)(nil),        // 6: mpb.ResGetAccountByTGUser
	(*ResGetAccountByUserId)(nil),        // 7: mpb.ResGetAccountByUserId
	(*ReqSetAccountTGLan)(nil),           // 8: mpb.ReqSetAccountTGLan
	(*ResBatchGetAccountsByUserIds)(nil), // 9: mpb.ResBatchGetAccountsByUserIds
	nil,                                  // 10: mpb.ResBatchGetAccountsByUserIds.AccountsEntry
	(*AccountInfo)(nil),                  // 11: mpb.AccountInfo
	(*BossDefeatHistory)(nil),            // 12: mpb.BossDefeatHistory
	(*BuffCard)(nil),                     // 13: mpb.BuffCard
	(EGame_BuffCardStatus)(0),            // 14: mpb.EGame.BuffCardStatus
	(*ReqUserId)(nil),                    // 15: mpb.ReqUserId
	(*ReqUserIds)(nil),                   // 16: mpb.ReqUserIds
	(*Empty)(nil),                        // 17: mpb.Empty
}
var file_grpc_account_proto_depIdxs = []int32{
	11, // 0: mpb.ResLoginTest.account:type_name -> mpb.AccountInfo
	12, // 1: mpb.ResLoginTest.boss_defeat_history:type_name -> mpb.BossDefeatHistory
	13, // 2: mpb.ResLoginTest.buff_cards:type_name -> mpb.BuffCard
	14, // 3: mpb.ResLoginTest.buff_card_status:type_name -> mpb.EGame.BuffCardStatus
	11, // 4: mpb.ResLoginByToken.account:type_name -> mpb.AccountInfo
	12, // 5: mpb.ResLoginByToken.boss_defeat_history:type_name -> mpb.BossDefeatHistory
	13, // 6: mpb.ResLoginByToken.buff_cards:type_name -> mpb.BuffCard
	14, // 7: mpb.ResLoginByToken.buff_card_status:type_name -> mpb.EGame.BuffCardStatus
	11, // 8: mpb.ResGetAccountByTGUser.account:type_name -> mpb.AccountInfo
	11, // 9: mpb.ResGetAccountByUserId.account:type_name -> mpb.AccountInfo
	10, // 10: mpb.ResBatchGetAccountsByUserIds.accounts:type_name -> mpb.ResBatchGetAccountsByUserIds.AccountsEntry
	11, // 11: mpb.ResBatchGetAccountsByUserIds.AccountsEntry.value:type_name -> mpb.AccountInfo
	15, // 12: mpb.AccountService.LoginTest:input_type -> mpb.ReqUserId
	1,  // 13: mpb.AccountService.GenerateLoginToken:input_type -> mpb.ReqGenerateLoginToken
	3,  // 14: mpb.AccountService.LoginByToken:input_type -> mpb.ReqLoginByToken
	5,  // 15: mpb.AccountService.GetAccountByTGUser:input_type -> mpb.ReqGetAccountByTGUser
	15, // 16: mpb.AccountService.GetAccountByUserId:input_type -> mpb.ReqUserId
	16, // 17: mpb.AccountService.BatchGetAccountsByUserIds:input_type -> mpb.ReqUserIds
	8,  // 18: mpb.AccountService.SetAccountTGLan:input_type -> mpb.ReqSetAccountTGLan
	0,  // 19: mpb.AccountService.LoginTest:output_type -> mpb.ResLoginTest
	2,  // 20: mpb.AccountService.GenerateLoginToken:output_type -> mpb.ResGenerateLoginToken
	4,  // 21: mpb.AccountService.LoginByToken:output_type -> mpb.ResLoginByToken
	6,  // 22: mpb.AccountService.GetAccountByTGUser:output_type -> mpb.ResGetAccountByTGUser
	7,  // 23: mpb.AccountService.GetAccountByUserId:output_type -> mpb.ResGetAccountByUserId
	9,  // 24: mpb.AccountService.BatchGetAccountsByUserIds:output_type -> mpb.ResBatchGetAccountsByUserIds
	17, // 25: mpb.AccountService.SetAccountTGLan:output_type -> mpb.Empty
	19, // [19:26] is the sub-list for method output_type
	12, // [12:19] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_grpc_account_proto_init() }
func file_grpc_account_proto_init() {
	if File_grpc_account_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_grpc_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResLoginTest); i {
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
		file_grpc_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqGenerateLoginToken); i {
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
		file_grpc_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResGenerateLoginToken); i {
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
		file_grpc_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqLoginByToken); i {
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
		file_grpc_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResLoginByToken); i {
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
		file_grpc_account_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqGetAccountByTGUser); i {
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
		file_grpc_account_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResGetAccountByTGUser); i {
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
		file_grpc_account_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResGetAccountByUserId); i {
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
		file_grpc_account_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqSetAccountTGLan); i {
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
		file_grpc_account_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResBatchGetAccountsByUserIds); i {
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
			RawDescriptor: file_grpc_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_account_proto_goTypes,
		DependencyIndexes: file_grpc_account_proto_depIdxs,
		MessageInfos:      file_grpc_account_proto_msgTypes,
	}.Build()
	File_grpc_account_proto = out.File
	file_grpc_account_proto_rawDesc = nil
	file_grpc_account_proto_goTypes = nil
	file_grpc_account_proto_depIdxs = nil
}
