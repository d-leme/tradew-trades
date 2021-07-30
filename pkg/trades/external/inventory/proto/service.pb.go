// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: pkg/trades/external/inventory/proto/service.proto

package proto

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_trades_external_inventory_proto_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_trades_external_inventory_proto_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_pkg_trades_external_inventory_proto_service_proto_rawDescGZIP(), []int{0}
}

type ItemToLock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Quantity int64  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *ItemToLock) Reset() {
	*x = ItemToLock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_trades_external_inventory_proto_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemToLock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemToLock) ProtoMessage() {}

func (x *ItemToLock) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_trades_external_inventory_proto_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemToLock.ProtoReflect.Descriptor instead.
func (*ItemToLock) Descriptor() ([]byte, []int) {
	return file_pkg_trades_external_inventory_proto_service_proto_rawDescGZIP(), []int{1}
}

func (x *ItemToLock) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ItemToLock) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type LockItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LockedBy           string        `protobuf:"bytes,1,opt,name=lockedBy,proto3" json:"lockedBy,omitempty"`
	OwnerID            string        `protobuf:"bytes,2,opt,name=ownerID,proto3" json:"ownerID,omitempty"`
	WantedItemsOwnerID string        `protobuf:"bytes,3,opt,name=wantedItemsOwnerID,proto3" json:"wantedItemsOwnerID,omitempty"`
	OfferedItems       []*ItemToLock `protobuf:"bytes,4,rep,name=offeredItems,proto3" json:"offeredItems,omitempty"`
	WantedItems        []*ItemToLock `protobuf:"bytes,5,rep,name=wantedItems,proto3" json:"wantedItems,omitempty"`
}

func (x *LockItemsRequest) Reset() {
	*x = LockItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_trades_external_inventory_proto_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockItemsRequest) ProtoMessage() {}

func (x *LockItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_trades_external_inventory_proto_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockItemsRequest.ProtoReflect.Descriptor instead.
func (*LockItemsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_trades_external_inventory_proto_service_proto_rawDescGZIP(), []int{2}
}

func (x *LockItemsRequest) GetLockedBy() string {
	if x != nil {
		return x.LockedBy
	}
	return ""
}

func (x *LockItemsRequest) GetOwnerID() string {
	if x != nil {
		return x.OwnerID
	}
	return ""
}

func (x *LockItemsRequest) GetWantedItemsOwnerID() string {
	if x != nil {
		return x.WantedItemsOwnerID
	}
	return ""
}

func (x *LockItemsRequest) GetOfferedItems() []*ItemToLock {
	if x != nil {
		return x.OfferedItems
	}
	return nil
}

func (x *LockItemsRequest) GetWantedItems() []*ItemToLock {
	if x != nil {
		return x.WantedItems
	}
	return nil
}

type ItemToTrade struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Quantity int64  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *ItemToTrade) Reset() {
	*x = ItemToTrade{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_trades_external_inventory_proto_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemToTrade) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemToTrade) ProtoMessage() {}

func (x *ItemToTrade) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_trades_external_inventory_proto_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemToTrade.ProtoReflect.Descriptor instead.
func (*ItemToTrade) Descriptor() ([]byte, []int) {
	return file_pkg_trades_external_inventory_proto_service_proto_rawDescGZIP(), []int{3}
}

func (x *ItemToTrade) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ItemToTrade) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type TradeItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TradeID            string         `protobuf:"bytes,1,opt,name=tradeID,proto3" json:"tradeID,omitempty"`
	OwnerID            string         `protobuf:"bytes,2,opt,name=ownerID,proto3" json:"ownerID,omitempty"`
	WantedItemsOwnerID string         `protobuf:"bytes,3,opt,name=wantedItemsOwnerID,proto3" json:"wantedItemsOwnerID,omitempty"`
	OfferedItems       []*ItemToTrade `protobuf:"bytes,4,rep,name=offeredItems,proto3" json:"offeredItems,omitempty"`
	WantedItems        []*ItemToTrade `protobuf:"bytes,5,rep,name=wantedItems,proto3" json:"wantedItems,omitempty"`
}

func (x *TradeItemsRequest) Reset() {
	*x = TradeItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_trades_external_inventory_proto_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TradeItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TradeItemsRequest) ProtoMessage() {}

func (x *TradeItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_trades_external_inventory_proto_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TradeItemsRequest.ProtoReflect.Descriptor instead.
func (*TradeItemsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_trades_external_inventory_proto_service_proto_rawDescGZIP(), []int{4}
}

func (x *TradeItemsRequest) GetTradeID() string {
	if x != nil {
		return x.TradeID
	}
	return ""
}

func (x *TradeItemsRequest) GetOwnerID() string {
	if x != nil {
		return x.OwnerID
	}
	return ""
}

func (x *TradeItemsRequest) GetWantedItemsOwnerID() string {
	if x != nil {
		return x.WantedItemsOwnerID
	}
	return ""
}

func (x *TradeItemsRequest) GetOfferedItems() []*ItemToTrade {
	if x != nil {
		return x.OfferedItems
	}
	return nil
}

func (x *TradeItemsRequest) GetWantedItems() []*ItemToTrade {
	if x != nil {
		return x.WantedItems
	}
	return nil
}

var File_pkg_trades_external_inventory_proto_service_proto protoreflect.FileDescriptor

var file_pkg_trades_external_inventory_proto_service_proto_rawDesc = []byte{
	0x0a, 0x31, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x73, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x07,
	0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x38, 0x0a, 0x0a, 0x49, 0x74, 0x65, 0x6d, 0x54,
	0x6f, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x22, 0xec, 0x01, 0x0a, 0x10, 0x4c, 0x6f, 0x63, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64,
	0x42, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64,
	0x42, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x12, 0x2e, 0x0a, 0x12,
	0x77, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x77, 0x61, 0x6e, 0x74, 0x65, 0x64,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x12, 0x39, 0x0a, 0x0c,
	0x6f, 0x66, 0x66, 0x65, 0x72, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x54, 0x6f, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x0c, 0x6f, 0x66, 0x66, 0x65, 0x72,
	0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x37, 0x0a, 0x0b, 0x77, 0x61, 0x6e, 0x74, 0x65,
	0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x6f, 0x4c,
	0x6f, 0x63, 0x6b, 0x52, 0x0b, 0x77, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x22, 0x39, 0x0a, 0x0b, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x6f, 0x54, 0x72, 0x61, 0x64, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0xed, 0x01, 0x0a, 0x11,
	0x54, 0x72, 0x61, 0x64, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x64, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x64, 0x65, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x49, 0x44, 0x12, 0x2e, 0x0a, 0x12, 0x77, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x77, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x49, 0x44, 0x12, 0x3a, 0x0a, 0x0c, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x65, 0x64,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x6f, 0x54, 0x72,
	0x61, 0x64, 0x65, 0x52, 0x0c, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x12, 0x38, 0x0a, 0x0b, 0x77, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x6f, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x0b,
	0x77, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x32, 0x90, 0x01, 0x0a, 0x10,
	0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3c, 0x0a, 0x09, 0x4c, 0x6f, 0x63, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1b, 0x2e,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3e,
	0x0a, 0x0a, 0x54, 0x72, 0x61, 0x64, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1c, 0x2e, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x25,
	0x5a, 0x23, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x73, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_trades_external_inventory_proto_service_proto_rawDescOnce sync.Once
	file_pkg_trades_external_inventory_proto_service_proto_rawDescData = file_pkg_trades_external_inventory_proto_service_proto_rawDesc
)

func file_pkg_trades_external_inventory_proto_service_proto_rawDescGZIP() []byte {
	file_pkg_trades_external_inventory_proto_service_proto_rawDescOnce.Do(func() {
		file_pkg_trades_external_inventory_proto_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_trades_external_inventory_proto_service_proto_rawDescData)
	})
	return file_pkg_trades_external_inventory_proto_service_proto_rawDescData
}

var file_pkg_trades_external_inventory_proto_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_trades_external_inventory_proto_service_proto_goTypes = []interface{}{
	(*Empty)(nil),             // 0: inventory.Empty
	(*ItemToLock)(nil),        // 1: inventory.ItemToLock
	(*LockItemsRequest)(nil),  // 2: inventory.LockItemsRequest
	(*ItemToTrade)(nil),       // 3: inventory.ItemToTrade
	(*TradeItemsRequest)(nil), // 4: inventory.TradeItemsRequest
}
var file_pkg_trades_external_inventory_proto_service_proto_depIdxs = []int32{
	1, // 0: inventory.LockItemsRequest.offeredItems:type_name -> inventory.ItemToLock
	1, // 1: inventory.LockItemsRequest.wantedItems:type_name -> inventory.ItemToLock
	3, // 2: inventory.TradeItemsRequest.offeredItems:type_name -> inventory.ItemToTrade
	3, // 3: inventory.TradeItemsRequest.wantedItems:type_name -> inventory.ItemToTrade
	2, // 4: inventory.InventoryService.LockItems:input_type -> inventory.LockItemsRequest
	4, // 5: inventory.InventoryService.TradeItems:input_type -> inventory.TradeItemsRequest
	0, // 6: inventory.InventoryService.LockItems:output_type -> inventory.Empty
	0, // 7: inventory.InventoryService.TradeItems:output_type -> inventory.Empty
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_trades_external_inventory_proto_service_proto_init() }
func file_pkg_trades_external_inventory_proto_service_proto_init() {
	if File_pkg_trades_external_inventory_proto_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_trades_external_inventory_proto_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_pkg_trades_external_inventory_proto_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemToLock); i {
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
		file_pkg_trades_external_inventory_proto_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockItemsRequest); i {
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
		file_pkg_trades_external_inventory_proto_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemToTrade); i {
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
		file_pkg_trades_external_inventory_proto_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TradeItemsRequest); i {
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
			RawDescriptor: file_pkg_trades_external_inventory_proto_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_trades_external_inventory_proto_service_proto_goTypes,
		DependencyIndexes: file_pkg_trades_external_inventory_proto_service_proto_depIdxs,
		MessageInfos:      file_pkg_trades_external_inventory_proto_service_proto_msgTypes,
	}.Build()
	File_pkg_trades_external_inventory_proto_service_proto = out.File
	file_pkg_trades_external_inventory_proto_service_proto_rawDesc = nil
	file_pkg_trades_external_inventory_proto_service_proto_goTypes = nil
	file_pkg_trades_external_inventory_proto_service_proto_depIdxs = nil
}
