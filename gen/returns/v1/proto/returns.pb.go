// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: returns.proto

package returns

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PackagingType int32

const (
	PackagingType_unknown_packaging PackagingType = 0
	PackagingType_bag               PackagingType = 1
	PackagingType_wrap              PackagingType = 2
	PackagingType_box               PackagingType = 3
)

// Enum value maps for PackagingType.
var (
	PackagingType_name = map[int32]string{
		0: "unknown_packaging",
		1: "bag",
		2: "wrap",
		3: "box",
	}
	PackagingType_value = map[string]int32{
		"unknown_packaging": 0,
		"bag":               1,
		"wrap":              2,
		"box":               3,
	}
)

func (x PackagingType) Enum() *PackagingType {
	p := new(PackagingType)
	*p = x
	return p
}

func (x PackagingType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PackagingType) Descriptor() protoreflect.EnumDescriptor {
	return file_returns_proto_enumTypes[0].Descriptor()
}

func (PackagingType) Type() protoreflect.EnumType {
	return &file_returns_proto_enumTypes[0]
}

func (x PackagingType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PackagingType.Descriptor instead.
func (PackagingType) EnumDescriptor() ([]byte, []int) {
	return file_returns_proto_rawDescGZIP(), []int{0}
}

type Status int32

const (
	Status_unknown_status Status = 0
	Status_pending        Status = 1
	Status_accepted       Status = 2
	Status_delivered      Status = 3
	Status_returned       Status = 4
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "unknown_status",
		1: "pending",
		2: "accepted",
		3: "delivered",
		4: "returned",
	}
	Status_value = map[string]int32{
		"unknown_status": 0,
		"pending":        1,
		"accepted":       2,
		"delivered":      3,
		"returned":       4,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_returns_proto_enumTypes[1].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_returns_proto_enumTypes[1]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_returns_proto_rawDescGZIP(), []int{1}
}

type AcceptReturnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId uint64 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	UserId  uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *AcceptReturnRequest) Reset() {
	*x = AcceptReturnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_returns_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptReturnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptReturnRequest) ProtoMessage() {}

func (x *AcceptReturnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_returns_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptReturnRequest.ProtoReflect.Descriptor instead.
func (*AcceptReturnRequest) Descriptor() ([]byte, []int) {
	return file_returns_proto_rawDescGZIP(), []int{0}
}

func (x *AcceptReturnRequest) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *AcceptReturnRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type AcceptReturnResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AcceptReturnResponse) Reset() {
	*x = AcceptReturnResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_returns_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptReturnResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptReturnResponse) ProtoMessage() {}

func (x *AcceptReturnResponse) ProtoReflect() protoreflect.Message {
	mi := &file_returns_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptReturnResponse.ProtoReflect.Descriptor instead.
func (*AcceptReturnResponse) Descriptor() ([]byte, []int) {
	return file_returns_proto_rawDescGZIP(), []int{1}
}

type ListReturnsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset uint64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListReturnsRequest) Reset() {
	*x = ListReturnsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_returns_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReturnsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReturnsRequest) ProtoMessage() {}

func (x *ListReturnsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_returns_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReturnsRequest.ProtoReflect.Descriptor instead.
func (*ListReturnsRequest) Descriptor() ([]byte, []int) {
	return file_returns_proto_rawDescGZIP(), []int{2}
}

func (x *ListReturnsRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListReturnsRequest) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListReturnsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*ListReturnsResponse_Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *ListReturnsResponse) Reset() {
	*x = ListReturnsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_returns_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReturnsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReturnsResponse) ProtoMessage() {}

func (x *ListReturnsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_returns_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReturnsResponse.ProtoReflect.Descriptor instead.
func (*ListReturnsResponse) Descriptor() ([]byte, []int) {
	return file_returns_proto_rawDescGZIP(), []int{3}
}

func (x *ListReturnsResponse) GetOrders() []*ListReturnsResponse_Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

type ListReturnsResponse_Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId   uint64                 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	UserId    uint64                 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Expiry    *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=expiry,proto3" json:"expiry,omitempty"`
	ReturnBy  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=return_by,json=returnBy,proto3" json:"return_by,omitempty"`
	Status    Status                 `protobuf:"varint,5,opt,name=status,proto3,enum=returns.Status" json:"status,omitempty"`
	Hash      string                 `protobuf:"bytes,6,opt,name=hash,proto3" json:"hash,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Cost      float64                `protobuf:"fixed64,8,opt,name=cost,proto3" json:"cost,omitempty"`
	Weight    float64                `protobuf:"fixed64,9,opt,name=weight,proto3" json:"weight,omitempty"`
	Packaging PackagingType          `protobuf:"varint,10,opt,name=packaging,proto3,enum=returns.PackagingType" json:"packaging,omitempty"`
}

func (x *ListReturnsResponse_Order) Reset() {
	*x = ListReturnsResponse_Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_returns_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReturnsResponse_Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReturnsResponse_Order) ProtoMessage() {}

func (x *ListReturnsResponse_Order) ProtoReflect() protoreflect.Message {
	mi := &file_returns_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReturnsResponse_Order.ProtoReflect.Descriptor instead.
func (*ListReturnsResponse_Order) Descriptor() ([]byte, []int) {
	return file_returns_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ListReturnsResponse_Order) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *ListReturnsResponse_Order) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ListReturnsResponse_Order) GetExpiry() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiry
	}
	return nil
}

func (x *ListReturnsResponse_Order) GetReturnBy() *timestamppb.Timestamp {
	if x != nil {
		return x.ReturnBy
	}
	return nil
}

func (x *ListReturnsResponse_Order) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_unknown_status
}

func (x *ListReturnsResponse_Order) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *ListReturnsResponse_Order) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ListReturnsResponse_Order) GetCost() float64 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *ListReturnsResponse_Order) GetWeight() float64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *ListReturnsResponse_Order) GetPackaging() PackagingType {
	if x != nil {
		return x.Packaging
	}
	return PackagingType_unknown_packaging
}

var File_returns_proto protoreflect.FileDescriptor

var file_returns_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x61, 0x0a, 0x13, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x52, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0a, 0xe0, 0x41, 0x02, 0xfa,
	0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x23, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x0a, 0xe0, 0x41, 0x02, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x54, 0x0a, 0x12,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x28, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x1f, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x28, 0x00, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x22, 0xd6, 0x03, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x82, 0x03, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x37, 0x0a, 0x09, 0x72, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x42,
	0x79, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0f, 0x2e, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x34, 0x0a, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x69,
	0x6e, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x73, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2a, 0x48, 0x0a, 0x0d, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11,
	0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x62, 0x61, 0x67, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04,
	0x77, 0x72, 0x61, 0x70, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x62, 0x6f, 0x78, 0x10, 0x03, 0x22,
	0x04, 0x08, 0x04, 0x10, 0x64, 0x2a, 0x5a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x12, 0x0a, 0x0e, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x01,
	0x12, 0x0c, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x10, 0x02, 0x12, 0x0d,
	0x0a, 0x09, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x10, 0x03, 0x12, 0x0c, 0x0a,
	0x08, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x10, 0x04, 0x22, 0x04, 0x08, 0x05, 0x10,
	0x64, 0x32, 0xbc, 0x06, 0x0a, 0x07, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x12, 0xca, 0x03,
	0x0a, 0x0c, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x12, 0x1c,
	0x2e, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x52,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x72,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x52, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xfc, 0x02, 0x92, 0x41,
	0xde, 0x02, 0x12, 0x31, 0xd0, 0x9f, 0xd1, 0x80, 0xd0, 0xb8, 0xd0, 0xbd, 0xd1, 0x8f, 0xd1, 0x82,
	0xd1, 0x8c, 0x20, 0xd0, 0xb2, 0xd0, 0xbe, 0xd0, 0xb7, 0xd0, 0xb2, 0xd1, 0x80, 0xd0, 0xb0, 0xd1,
	0x82, 0x20, 0xd0, 0xbe, 0xd1, 0x82, 0x20, 0xd0, 0xba, 0xd0, 0xbb, 0xd0, 0xb8, 0xd0, 0xb5, 0xd0,
	0xbd, 0xd1, 0x82, 0xd0, 0xb0, 0x1a, 0xa8, 0x02, 0xd0, 0x98, 0xd1, 0x81, 0xd0, 0xbf, 0xd0, 0xbe,
	0xd0, 0xbb, 0xd1, 0x8c, 0xd0, 0xb7, 0xd1, 0x83, 0xd0, 0xb5, 0xd1, 0x82, 0xd1, 0x81, 0xd1, 0x8f,
	0x20, 0xd0, 0xb4, 0xd0, 0xbb, 0xd1, 0x8f, 0x20, 0xd0, 0xbf, 0xd1, 0x80, 0xd0, 0xb8, 0xd0, 0xbd,
	0xd1, 0x8f, 0xd1, 0x82, 0xd0, 0xb8, 0xd1, 0x8f, 0x20, 0xd0, 0xb2, 0xd0, 0xbe, 0xd0, 0xb7, 0xd0,
	0xb2, 0xd1, 0x80, 0xd0, 0xb0, 0xd1, 0x82, 0xd0, 0xb0, 0x20, 0xd0, 0xb7, 0xd0, 0xb0, 0xd0, 0xba,
	0xd0, 0xb0, 0xd0, 0xb7, 0xd0, 0xb0, 0x20, 0xd0, 0xbe, 0xd1, 0x82, 0x20, 0xd0, 0xba, 0xd0, 0xbb,
	0xd0, 0xb8, 0xd0, 0xb5, 0xd0, 0xbd, 0xd1, 0x82, 0xd0, 0xb0, 0x2e, 0x20, 0xd0, 0x92, 0xd0, 0xbe,
	0xd0, 0xb7, 0xd0, 0xb2, 0xd1, 0x80, 0xd0, 0xb0, 0xd1, 0x82, 0x20, 0xd0, 0xbc, 0xd0, 0xbe, 0xd0,
	0xb6, 0xd0, 0xb5, 0xd1, 0x82, 0x20, 0xd0, 0xb1, 0xd1, 0x8b, 0xd1, 0x82, 0xd1, 0x8c, 0x20, 0xd0,
	0xbf, 0xd1, 0x80, 0xd0, 0xb8, 0xd0, 0xbd, 0xd1, 0x8f, 0xd1, 0x82, 0x20, 0xd0, 0xb2, 0x20, 0xd1,
	0x82, 0xd0, 0xb5, 0xd1, 0x87, 0xd0, 0xb5, 0xd0, 0xbd, 0xd0, 0xb8, 0xd0, 0xb5, 0x20, 0xd0, 0xb4,
	0xd0, 0xb2, 0xd1, 0x83, 0xd1, 0x85, 0x20, 0xd0, 0xb4, 0xd0, 0xbd, 0xd0, 0xb5, 0xd0, 0xb9, 0x20,
	0xd1, 0x81, 0x20, 0xd0, 0xbc, 0xd0, 0xbe, 0xd0, 0xbc, 0xd0, 0xb5, 0xd0, 0xbd, 0xd1, 0x82, 0xd0,
	0xb0, 0x20, 0xd0, 0xb2, 0xd1, 0x8b, 0xd0, 0xb4, 0xd0, 0xb0, 0xd1, 0x87, 0xd0, 0xb8, 0x20, 0xd0,
	0xb7, 0xd0, 0xb0, 0xd0, 0xba, 0xd0, 0xb0, 0xd0, 0xb7, 0xd0, 0xb0, 0x2e, 0x20, 0xd0, 0x97, 0xd0,
	0xb0, 0xd0, 0xba, 0xd0, 0xb0, 0xd0, 0xb7, 0x20, 0xd0, 0xb4, 0xd0, 0xbe, 0xd0, 0xbb, 0xd0, 0xb6,
	0xd0, 0xb5, 0xd0, 0xbd, 0x20, 0xd0, 0xb1, 0xd1, 0x8b, 0xd1, 0x82, 0xd1, 0x8c, 0x20, 0xd0, 0xb2,
	0xd1, 0x8b, 0xd0, 0xb4, 0xd0, 0xb0, 0xd0, 0xbd, 0x20, 0xd0, 0xb8, 0xd0, 0xb7, 0x20, 0xd1, 0x8d,
	0xd1, 0x82, 0xd0, 0xbe, 0xd0, 0xb3, 0xd0, 0xbe, 0x20, 0xd0, 0x9f, 0xd0, 0x92, 0xd0, 0x97, 0x2e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x12, 0xe3, 0x02, 0x0a, 0x0b, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x12, 0x1b, 0x2e, 0x72, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x98, 0x02, 0x92, 0x41, 0xfc, 0x01, 0x12, 0x30, 0xd0, 0x9f,
	0xd0, 0xbe, 0xd0, 0xbb, 0xd1, 0x83, 0xd1, 0x87, 0xd0, 0xb8, 0xd1, 0x82, 0xd1, 0x8c, 0x20, 0xd1,
	0x81, 0xd0, 0xbf, 0xd0, 0xb8, 0xd1, 0x81, 0xd0, 0xbe, 0xd0, 0xba, 0x20, 0xd0, 0xb2, 0xd0, 0xbe,
	0xd0, 0xb7, 0xd0, 0xb2, 0xd1, 0x80, 0xd0, 0xb0, 0xd1, 0x82, 0xd0, 0xbe, 0xd0, 0xb2, 0x1a, 0xc7,
	0x01, 0xd0, 0x98, 0xd1, 0x81, 0xd0, 0xbf, 0xd0, 0xbe, 0xd0, 0xbb, 0xd1, 0x8c, 0xd0, 0xb7, 0xd1,
	0x83, 0xd0, 0xb5, 0xd1, 0x82, 0xd1, 0x81, 0xd1, 0x8f, 0x20, 0xd0, 0xb4, 0xd0, 0xbb, 0xd1, 0x8f,
	0x20, 0xd0, 0xbf, 0xd0, 0xbe, 0xd0, 0xbb, 0xd1, 0x83, 0xd1, 0x87, 0xd0, 0xb5, 0xd0, 0xbd, 0xd0,
	0xb8, 0xd1, 0x8f, 0x20, 0xd1, 0x81, 0xd0, 0xbf, 0xd0, 0xb8, 0xd1, 0x81, 0xd0, 0xba, 0xd0, 0xb0,
	0x20, 0xd0, 0xb2, 0xd0, 0xbe, 0xd0, 0xb7, 0xd0, 0xb2, 0xd1, 0x80, 0xd0, 0xb0, 0xd1, 0x82, 0xd0,
	0xbe, 0xd0, 0xb2, 0x2e, 0x20, 0xd0, 0x9a, 0xd0, 0xbe, 0xd0, 0xbc, 0xd0, 0xb0, 0xd0, 0xbd, 0xd0,
	0xb4, 0xd0, 0xb0, 0x20, 0xd0, 0xb2, 0xd0, 0xbe, 0xd0, 0xb7, 0xd0, 0xb2, 0xd1, 0x80, 0xd0, 0xb0,
	0xd1, 0x89, 0xd0, 0xb0, 0xd0, 0xb5, 0xd1, 0x82, 0x20, 0xd1, 0x81, 0xd0, 0xbf, 0xd0, 0xb8, 0xd1,
	0x81, 0xd0, 0xbe, 0xd0, 0xba, 0x20, 0xd0, 0xb2, 0xd0, 0xbe, 0xd0, 0xb7, 0xd0, 0xb2, 0xd1, 0x80,
	0xd0, 0xb0, 0xd1, 0x82, 0xd0, 0xbe, 0xd0, 0xb2, 0x20, 0xd1, 0x81, 0x20, 0xd0, 0xb2, 0xd0, 0xbe,
	0xd0, 0xb7, 0xd0, 0xbc, 0xd0, 0xbe, 0xd0, 0xb6, 0xd0, 0xbd, 0xd0, 0xbe, 0xd1, 0x81, 0xd1, 0x82,
	0xd1, 0x8c, 0xd1, 0x8e, 0x20, 0xd0, 0xbf, 0xd0, 0xb0, 0xd0, 0xb3, 0xd0, 0xb8, 0xd0, 0xbd, 0xd0,
	0xb0, 0xd1, 0x86, 0xd0, 0xb8, 0xd0, 0xb8, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74,
	0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e,
	0x64, 0x65, 0x76, 0x2f, 0x79, 0x75, 0x77, 0x65, 0x65, 0x62, 0x69, 0x78, 0x2f, 0x68, 0x6f, 0x6d,
	0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2d, 0x31, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x3b, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_returns_proto_rawDescOnce sync.Once
	file_returns_proto_rawDescData = file_returns_proto_rawDesc
)

func file_returns_proto_rawDescGZIP() []byte {
	file_returns_proto_rawDescOnce.Do(func() {
		file_returns_proto_rawDescData = protoimpl.X.CompressGZIP(file_returns_proto_rawDescData)
	})
	return file_returns_proto_rawDescData
}

var file_returns_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_returns_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_returns_proto_goTypes = []any{
	(PackagingType)(0),                // 0: returns.PackagingType
	(Status)(0),                       // 1: returns.Status
	(*AcceptReturnRequest)(nil),       // 2: returns.AcceptReturnRequest
	(*AcceptReturnResponse)(nil),      // 3: returns.AcceptReturnResponse
	(*ListReturnsRequest)(nil),        // 4: returns.ListReturnsRequest
	(*ListReturnsResponse)(nil),       // 5: returns.ListReturnsResponse
	(*ListReturnsResponse_Order)(nil), // 6: returns.ListReturnsResponse.Order
	(*timestamppb.Timestamp)(nil),     // 7: google.protobuf.Timestamp
}
var file_returns_proto_depIdxs = []int32{
	6, // 0: returns.ListReturnsResponse.orders:type_name -> returns.ListReturnsResponse.Order
	7, // 1: returns.ListReturnsResponse.Order.expiry:type_name -> google.protobuf.Timestamp
	7, // 2: returns.ListReturnsResponse.Order.return_by:type_name -> google.protobuf.Timestamp
	1, // 3: returns.ListReturnsResponse.Order.status:type_name -> returns.Status
	7, // 4: returns.ListReturnsResponse.Order.created_at:type_name -> google.protobuf.Timestamp
	0, // 5: returns.ListReturnsResponse.Order.packaging:type_name -> returns.PackagingType
	2, // 6: returns.Returns.AcceptReturn:input_type -> returns.AcceptReturnRequest
	4, // 7: returns.Returns.ListReturns:input_type -> returns.ListReturnsRequest
	3, // 8: returns.Returns.AcceptReturn:output_type -> returns.AcceptReturnResponse
	5, // 9: returns.Returns.ListReturns:output_type -> returns.ListReturnsResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_returns_proto_init() }
func file_returns_proto_init() {
	if File_returns_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_returns_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AcceptReturnRequest); i {
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
		file_returns_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AcceptReturnResponse); i {
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
		file_returns_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ListReturnsRequest); i {
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
		file_returns_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ListReturnsResponse); i {
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
		file_returns_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ListReturnsResponse_Order); i {
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
			RawDescriptor: file_returns_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_returns_proto_goTypes,
		DependencyIndexes: file_returns_proto_depIdxs,
		EnumInfos:         file_returns_proto_enumTypes,
		MessageInfos:      file_returns_proto_msgTypes,
	}.Build()
	File_returns_proto = out.File
	file_returns_proto_rawDesc = nil
	file_returns_proto_goTypes = nil
	file_returns_proto_depIdxs = nil
}
