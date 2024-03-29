// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.3
// source: grpc_market.proto

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

type ReqGetGoodsOrdersOnSell struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId  uint32 `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	PageNum uint32 `protobuf:"varint,2,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
}

func (x *ReqGetGoodsOrdersOnSell) Reset() {
	*x = ReqGetGoodsOrdersOnSell{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_market_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqGetGoodsOrdersOnSell) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqGetGoodsOrdersOnSell) ProtoMessage() {}

func (x *ReqGetGoodsOrdersOnSell) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_market_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqGetGoodsOrdersOnSell.ProtoReflect.Descriptor instead.
func (*ReqGetGoodsOrdersOnSell) Descriptor() ([]byte, []int) {
	return file_grpc_market_proto_rawDescGZIP(), []int{0}
}

func (x *ReqGetGoodsOrdersOnSell) GetItemId() uint32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *ReqGetGoodsOrdersOnSell) GetPageNum() uint32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

type ResGetGoodsOrdersOnSell struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId   uint32        `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	PageNum  uint32        `protobuf:"varint,2,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	Orders   []*GoodsOrder `protobuf:"bytes,3,rep,name=orders,proto3" json:"orders,omitempty"`
	OrderCnt uint32        `protobuf:"varint,4,opt,name=order_cnt,json=orderCnt,proto3" json:"order_cnt,omitempty"`
}

func (x *ResGetGoodsOrdersOnSell) Reset() {
	*x = ResGetGoodsOrdersOnSell{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_market_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResGetGoodsOrdersOnSell) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResGetGoodsOrdersOnSell) ProtoMessage() {}

func (x *ResGetGoodsOrdersOnSell) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_market_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResGetGoodsOrdersOnSell.ProtoReflect.Descriptor instead.
func (*ResGetGoodsOrdersOnSell) Descriptor() ([]byte, []int) {
	return file_grpc_market_proto_rawDescGZIP(), []int{1}
}

func (x *ResGetGoodsOrdersOnSell) GetItemId() uint32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *ResGetGoodsOrdersOnSell) GetPageNum() uint32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *ResGetGoodsOrdersOnSell) GetOrders() []*GoodsOrder {
	if x != nil {
		return x.Orders
	}
	return nil
}

func (x *ResGetGoodsOrdersOnSell) GetOrderCnt() uint32 {
	if x != nil {
		return x.OrderCnt
	}
	return 0
}

type ResGetMyGoodsOrdersOnSell struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*GoodsOrder `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *ResGetMyGoodsOrdersOnSell) Reset() {
	*x = ResGetMyGoodsOrdersOnSell{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_market_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResGetMyGoodsOrdersOnSell) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResGetMyGoodsOrdersOnSell) ProtoMessage() {}

func (x *ResGetMyGoodsOrdersOnSell) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_market_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResGetMyGoodsOrdersOnSell.ProtoReflect.Descriptor instead.
func (*ResGetMyGoodsOrdersOnSell) Descriptor() ([]byte, []int) {
	return file_grpc_market_proto_rawDescGZIP(), []int{2}
}

func (x *ResGetMyGoodsOrdersOnSell) GetOrders() []*GoodsOrder {
	if x != nil {
		return x.Orders
	}
	return nil
}

type ReqPublishGoodsOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ItemId uint32 `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Num    uint64 `protobuf:"varint,3,opt,name=num,proto3" json:"num,omitempty"`
	Price  uint64 `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *ReqPublishGoodsOrder) Reset() {
	*x = ReqPublishGoodsOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_market_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqPublishGoodsOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqPublishGoodsOrder) ProtoMessage() {}

func (x *ReqPublishGoodsOrder) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_market_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqPublishGoodsOrder.ProtoReflect.Descriptor instead.
func (*ReqPublishGoodsOrder) Descriptor() ([]byte, []int) {
	return file_grpc_market_proto_rawDescGZIP(), []int{3}
}

func (x *ReqPublishGoodsOrder) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ReqPublishGoodsOrder) GetItemId() uint32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *ReqPublishGoodsOrder) GetNum() uint64 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *ReqPublishGoodsOrder) GetPrice() uint64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type ReqPurchaseGoodsOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrderUuid uint64 `protobuf:"varint,2,opt,name=order_uuid,json=orderUuid,proto3" json:"order_uuid,omitempty"`
}

func (x *ReqPurchaseGoodsOrder) Reset() {
	*x = ReqPurchaseGoodsOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_market_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqPurchaseGoodsOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqPurchaseGoodsOrder) ProtoMessage() {}

func (x *ReqPurchaseGoodsOrder) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_market_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqPurchaseGoodsOrder.ProtoReflect.Descriptor instead.
func (*ReqPurchaseGoodsOrder) Descriptor() ([]byte, []int) {
	return file_grpc_market_proto_rawDescGZIP(), []int{4}
}

func (x *ReqPurchaseGoodsOrder) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ReqPurchaseGoodsOrder) GetOrderUuid() uint64 {
	if x != nil {
		return x.OrderUuid
	}
	return 0
}

type ResPurchaseGoodsOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Goods    *Item  `protobuf:"bytes,1,opt,name=goods,proto3" json:"goods,omitempty"`
	ManaCost uint64 `protobuf:"varint,2,opt,name=mana_cost,json=manaCost,proto3" json:"mana_cost,omitempty"`
	ManaLeft uint64 `protobuf:"varint,3,opt,name=mana_left,json=manaLeft,proto3" json:"mana_left,omitempty"`
}

func (x *ResPurchaseGoodsOrder) Reset() {
	*x = ResPurchaseGoodsOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_market_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResPurchaseGoodsOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResPurchaseGoodsOrder) ProtoMessage() {}

func (x *ResPurchaseGoodsOrder) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_market_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResPurchaseGoodsOrder.ProtoReflect.Descriptor instead.
func (*ResPurchaseGoodsOrder) Descriptor() ([]byte, []int) {
	return file_grpc_market_proto_rawDescGZIP(), []int{5}
}

func (x *ResPurchaseGoodsOrder) GetGoods() *Item {
	if x != nil {
		return x.Goods
	}
	return nil
}

func (x *ResPurchaseGoodsOrder) GetManaCost() uint64 {
	if x != nil {
		return x.ManaCost
	}
	return 0
}

func (x *ResPurchaseGoodsOrder) GetManaLeft() uint64 {
	if x != nil {
		return x.ManaLeft
	}
	return 0
}

type ReqTakeOffGoodsOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrderUuid uint64 `protobuf:"varint,2,opt,name=order_uuid,json=orderUuid,proto3" json:"order_uuid,omitempty"`
}

func (x *ReqTakeOffGoodsOrder) Reset() {
	*x = ReqTakeOffGoodsOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_market_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqTakeOffGoodsOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqTakeOffGoodsOrder) ProtoMessage() {}

func (x *ReqTakeOffGoodsOrder) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_market_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqTakeOffGoodsOrder.ProtoReflect.Descriptor instead.
func (*ReqTakeOffGoodsOrder) Descriptor() ([]byte, []int) {
	return file_grpc_market_proto_rawDescGZIP(), []int{6}
}

func (x *ReqTakeOffGoodsOrder) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ReqTakeOffGoodsOrder) GetOrderUuid() uint64 {
	if x != nil {
		return x.OrderUuid
	}
	return 0
}

type ResTakeOffGoodsOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddItem    *Item `protobuf:"bytes,1,opt,name=add_item,json=addItem,proto3" json:"add_item,omitempty"`
	UpdateItem *Item `protobuf:"bytes,2,opt,name=update_item,json=updateItem,proto3" json:"update_item,omitempty"`
}

func (x *ResTakeOffGoodsOrder) Reset() {
	*x = ResTakeOffGoodsOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_market_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResTakeOffGoodsOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResTakeOffGoodsOrder) ProtoMessage() {}

func (x *ResTakeOffGoodsOrder) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_market_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResTakeOffGoodsOrder.ProtoReflect.Descriptor instead.
func (*ResTakeOffGoodsOrder) Descriptor() ([]byte, []int) {
	return file_grpc_market_proto_rawDescGZIP(), []int{7}
}

func (x *ResTakeOffGoodsOrder) GetAddItem() *Item {
	if x != nil {
		return x.AddItem
	}
	return nil
}

func (x *ResTakeOffGoodsOrder) GetUpdateItem() *Item {
	if x != nil {
		return x.UpdateItem
	}
	return nil
}

var File_grpc_market_proto protoreflect.FileDescriptor

var file_grpc_market_proto_rawDesc = []byte{
	0x0a, 0x11, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x70, 0x62, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x17, 0x52, 0x65, 0x71, 0x47, 0x65, 0x74,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x4f, 0x6e, 0x53, 0x65, 0x6c,
	0x6c, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x70, 0x61,
	0x67, 0x65, 0x4e, 0x75, 0x6d, 0x22, 0x93, 0x01, 0x0a, 0x17, 0x52, 0x65, 0x73, 0x47, 0x65, 0x74,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x4f, 0x6e, 0x53, 0x65, 0x6c,
	0x6c, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x70, 0x61,
	0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x27, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x63, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6e, 0x74, 0x22, 0x44, 0x0a, 0x19, 0x52,
	0x65, 0x73, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x4f, 0x6e, 0x53, 0x65, 0x6c, 0x6c, 0x12, 0x27, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x47,
	0x6f, 0x6f, 0x64, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x22, 0x70, 0x0a, 0x14, 0x52, 0x65, 0x71, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x47,
	0x6f, 0x6f, 0x64, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e,
	0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x22, 0x4f, 0x0a, 0x15, 0x52, 0x65, 0x71, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61,
	0x73, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x55, 0x75, 0x69, 0x64, 0x22, 0x72, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x50, 0x75, 0x72, 0x63, 0x68,
	0x61, 0x73, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1f, 0x0a,
	0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x6d,
	0x70, 0x62, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x6d, 0x61, 0x6e, 0x61, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x6d, 0x61, 0x6e, 0x61, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d,
	0x61, 0x6e, 0x61, 0x5f, 0x6c, 0x65, 0x66, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x6d, 0x61, 0x6e, 0x61, 0x4c, 0x65, 0x66, 0x74, 0x22, 0x4e, 0x0a, 0x14, 0x52, 0x65, 0x71, 0x54,
	0x61, 0x6b, 0x65, 0x4f, 0x66, 0x66, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x22, 0x68, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x54,
	0x61, 0x6b, 0x65, 0x4f, 0x66, 0x66, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x24, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x2a, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x6d, 0x70,
	0x62, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x32, 0xb9, 0x03, 0x0a, 0x0d, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x4f, 0x6e, 0x53, 0x65, 0x6c, 0x6c, 0x12, 0x1c, 0x2e, 0x6d,
	0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x4f, 0x6e, 0x53, 0x65, 0x6c, 0x6c, 0x1a, 0x1c, 0x2e, 0x6d, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x73, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x4f, 0x6e, 0x53, 0x65, 0x6c, 0x6c, 0x12, 0x48, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4d,
	0x79, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x4f, 0x6e, 0x53, 0x65,
	0x6c, 0x6c, 0x12, 0x0e, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x1a, 0x1e, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x47, 0x65, 0x74, 0x4d,
	0x79, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x4f, 0x6e, 0x53, 0x65,
	0x6c, 0x6c, 0x12, 0x3a, 0x0a, 0x11, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x47, 0x6f, 0x6f,
	0x64, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x71, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x1a, 0x0a, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4c,
	0x0a, 0x12, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x50, 0x75,
	0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x1a, 0x1a, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61,
	0x73, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x49, 0x0a, 0x11,
	0x54, 0x61, 0x6b, 0x65, 0x4f, 0x66, 0x66, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x12, 0x19, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x54, 0x61, 0x6b, 0x65, 0x4f,
	0x66, 0x66, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x19, 0x2e, 0x6d,
	0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x54, 0x61, 0x6b, 0x65, 0x4f, 0x66, 0x66, 0x47, 0x6f, 0x6f,
	0x64, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x17, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x46, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x72, 0x61,
	0x64, 0x65, 0x12, 0x0e, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x1a, 0x0a, 0x2e, 0x6d, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x07,
	0x5a, 0x05, 0x2e, 0x2f, 0x6d, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_market_proto_rawDescOnce sync.Once
	file_grpc_market_proto_rawDescData = file_grpc_market_proto_rawDesc
)

func file_grpc_market_proto_rawDescGZIP() []byte {
	file_grpc_market_proto_rawDescOnce.Do(func() {
		file_grpc_market_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_market_proto_rawDescData)
	})
	return file_grpc_market_proto_rawDescData
}

var file_grpc_market_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_grpc_market_proto_goTypes = []interface{}{
	(*ReqGetGoodsOrdersOnSell)(nil),   // 0: mpb.ReqGetGoodsOrdersOnSell
	(*ResGetGoodsOrdersOnSell)(nil),   // 1: mpb.ResGetGoodsOrdersOnSell
	(*ResGetMyGoodsOrdersOnSell)(nil), // 2: mpb.ResGetMyGoodsOrdersOnSell
	(*ReqPublishGoodsOrder)(nil),      // 3: mpb.ReqPublishGoodsOrder
	(*ReqPurchaseGoodsOrder)(nil),     // 4: mpb.ReqPurchaseGoodsOrder
	(*ResPurchaseGoodsOrder)(nil),     // 5: mpb.ResPurchaseGoodsOrder
	(*ReqTakeOffGoodsOrder)(nil),      // 6: mpb.ReqTakeOffGoodsOrder
	(*ResTakeOffGoodsOrder)(nil),      // 7: mpb.ResTakeOffGoodsOrder
	(*GoodsOrder)(nil),                // 8: mpb.GoodsOrder
	(*Item)(nil),                      // 9: mpb.Item
	(*ReqUserId)(nil),                 // 10: mpb.ReqUserId
	(*Empty)(nil),                     // 11: mpb.Empty
}
var file_grpc_market_proto_depIdxs = []int32{
	8,  // 0: mpb.ResGetGoodsOrdersOnSell.orders:type_name -> mpb.GoodsOrder
	8,  // 1: mpb.ResGetMyGoodsOrdersOnSell.orders:type_name -> mpb.GoodsOrder
	9,  // 2: mpb.ResPurchaseGoodsOrder.goods:type_name -> mpb.Item
	9,  // 3: mpb.ResTakeOffGoodsOrder.add_item:type_name -> mpb.Item
	9,  // 4: mpb.ResTakeOffGoodsOrder.update_item:type_name -> mpb.Item
	0,  // 5: mpb.MarketService.GetGoodsOrdersOnSell:input_type -> mpb.ReqGetGoodsOrdersOnSell
	10, // 6: mpb.MarketService.GetMyGoodsOrdersOnSell:input_type -> mpb.ReqUserId
	3,  // 7: mpb.MarketService.PublishGoodsOrder:input_type -> mpb.ReqPublishGoodsOrder
	4,  // 8: mpb.MarketService.PurchaseGoodsOrder:input_type -> mpb.ReqPurchaseGoodsOrder
	6,  // 9: mpb.MarketService.TakeOffGoodsOrder:input_type -> mpb.ReqTakeOffGoodsOrder
	10, // 10: mpb.MarketService.AdminFreezeAccountTrade:input_type -> mpb.ReqUserId
	1,  // 11: mpb.MarketService.GetGoodsOrdersOnSell:output_type -> mpb.ResGetGoodsOrdersOnSell
	2,  // 12: mpb.MarketService.GetMyGoodsOrdersOnSell:output_type -> mpb.ResGetMyGoodsOrdersOnSell
	11, // 13: mpb.MarketService.PublishGoodsOrder:output_type -> mpb.Empty
	5,  // 14: mpb.MarketService.PurchaseGoodsOrder:output_type -> mpb.ResPurchaseGoodsOrder
	7,  // 15: mpb.MarketService.TakeOffGoodsOrder:output_type -> mpb.ResTakeOffGoodsOrder
	11, // 16: mpb.MarketService.AdminFreezeAccountTrade:output_type -> mpb.Empty
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_grpc_market_proto_init() }
func file_grpc_market_proto_init() {
	if File_grpc_market_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_grpc_market_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqGetGoodsOrdersOnSell); i {
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
		file_grpc_market_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResGetGoodsOrdersOnSell); i {
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
		file_grpc_market_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResGetMyGoodsOrdersOnSell); i {
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
		file_grpc_market_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqPublishGoodsOrder); i {
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
		file_grpc_market_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqPurchaseGoodsOrder); i {
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
		file_grpc_market_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResPurchaseGoodsOrder); i {
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
		file_grpc_market_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqTakeOffGoodsOrder); i {
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
		file_grpc_market_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResTakeOffGoodsOrder); i {
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
			RawDescriptor: file_grpc_market_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_market_proto_goTypes,
		DependencyIndexes: file_grpc_market_proto_depIdxs,
		MessageInfos:      file_grpc_market_proto_msgTypes,
	}.Build()
	File_grpc_market_proto = out.File
	file_grpc_market_proto_rawDesc = nil
	file_grpc_market_proto_goTypes = nil
	file_grpc_market_proto_depIdxs = nil
}
