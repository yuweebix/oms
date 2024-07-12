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

type AcceptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId uint64 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	UserId  uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *AcceptRequest) Reset() {
	*x = AcceptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_returns_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptRequest) ProtoMessage() {}

func (x *AcceptRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AcceptRequest.ProtoReflect.Descriptor instead.
func (*AcceptRequest) Descriptor() ([]byte, []int) {
	return file_returns_proto_rawDescGZIP(), []int{0}
}

func (x *AcceptRequest) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *AcceptRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type AcceptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AcceptResponse) Reset() {
	*x = AcceptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_returns_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptResponse) ProtoMessage() {}

func (x *AcceptResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AcceptResponse.ProtoReflect.Descriptor instead.
func (*AcceptResponse) Descriptor() ([]byte, []int) {
	return file_returns_proto_rawDescGZIP(), []int{1}
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset uint64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_returns_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_returns_proto_rawDescGZIP(), []int{2}
}

func (x *ListRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListRequest) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*ListResponse_Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_returns_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_returns_proto_rawDescGZIP(), []int{3}
}

func (x *ListResponse) GetOrders() []*ListResponse_Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

type ListResponse_Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId   uint64                 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	UserId    uint64                 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Expiry    *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=expiry,proto3" json:"expiry,omitempty"`
	ReturnBy  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=return_by,json=returnBy,proto3" json:"return_by,omitempty"`
	Status    string                 `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Hash      string                 `protobuf:"bytes,6,opt,name=hash,proto3" json:"hash,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Cost      float32                `protobuf:"fixed32,8,opt,name=cost,proto3" json:"cost,omitempty"`
	Weight    float32                `protobuf:"fixed32,9,opt,name=weight,proto3" json:"weight,omitempty"`
	Packaging string                 `protobuf:"bytes,10,opt,name=packaging,proto3" json:"packaging,omitempty"`
}

func (x *ListResponse_Order) Reset() {
	*x = ListResponse_Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_returns_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse_Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse_Order) ProtoMessage() {}

func (x *ListResponse_Order) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListResponse_Order.ProtoReflect.Descriptor instead.
func (*ListResponse_Order) Descriptor() ([]byte, []int) {
	return file_returns_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ListResponse_Order) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *ListResponse_Order) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ListResponse_Order) GetExpiry() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiry
	}
	return nil
}

func (x *ListResponse_Order) GetReturnBy() *timestamppb.Timestamp {
	if x != nil {
		return x.ReturnBy
	}
	return nil
}

func (x *ListResponse_Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ListResponse_Order) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *ListResponse_Order) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ListResponse_Order) GetCost() float32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *ListResponse_Order) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *ListResponse_Order) GetPackaging() string {
	if x != nil {
		return x.Packaging
	}
	return ""
}

var File_returns_proto protoreflect.FileDescriptor

var file_returns_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x5b, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x25, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x0a, 0xe0, 0x41, 0x02, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0a, 0xe0, 0x41, 0x02, 0xfa, 0x42,
	0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x10, 0x0a,
	0x0e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x4d, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x32, 0x02, 0x28, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1f, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x32, 0x02, 0x28, 0x00, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x9e,
	0x03, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x32, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x1a, 0xd9, 0x02, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x32, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x37, 0x0a, 0x09, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f,
	0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x42, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x32,
	0x91, 0x06, 0x0a, 0x07, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x12, 0xb6, 0x03, 0x0a, 0x06,
	0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x12, 0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e,
	0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xfc, 0x02, 0x92, 0x41, 0xde, 0x02, 0x12, 0x31, 0xd0, 0x9f,
	0xd1, 0x80, 0xd0, 0xb8, 0xd0, 0xbd, 0xd1, 0x8f, 0xd1, 0x82, 0xd1, 0x8c, 0x20, 0xd0, 0xb2, 0xd0,
	0xbe, 0xd0, 0xb7, 0xd0, 0xb2, 0xd1, 0x80, 0xd0, 0xb0, 0xd1, 0x82, 0x20, 0xd0, 0xbe, 0xd1, 0x82,
	0x20, 0xd0, 0xba, 0xd0, 0xbb, 0xd0, 0xb8, 0xd0, 0xb5, 0xd0, 0xbd, 0xd1, 0x82, 0xd0, 0xb0, 0x1a,
	0xa8, 0x02, 0xd0, 0x98, 0xd1, 0x81, 0xd0, 0xbf, 0xd0, 0xbe, 0xd0, 0xbb, 0xd1, 0x8c, 0xd0, 0xb7,
	0xd1, 0x83, 0xd0, 0xb5, 0xd1, 0x82, 0xd1, 0x81, 0xd1, 0x8f, 0x20, 0xd0, 0xb4, 0xd0, 0xbb, 0xd1,
	0x8f, 0x20, 0xd0, 0xbf, 0xd1, 0x80, 0xd0, 0xb8, 0xd0, 0xbd, 0xd1, 0x8f, 0xd1, 0x82, 0xd0, 0xb8,
	0xd1, 0x8f, 0x20, 0xd0, 0xb2, 0xd0, 0xbe, 0xd0, 0xb7, 0xd0, 0xb2, 0xd1, 0x80, 0xd0, 0xb0, 0xd1,
	0x82, 0xd0, 0xb0, 0x20, 0xd0, 0xb7, 0xd0, 0xb0, 0xd0, 0xba, 0xd0, 0xb0, 0xd0, 0xb7, 0xd0, 0xb0,
	0x20, 0xd0, 0xbe, 0xd1, 0x82, 0x20, 0xd0, 0xba, 0xd0, 0xbb, 0xd0, 0xb8, 0xd0, 0xb5, 0xd0, 0xbd,
	0xd1, 0x82, 0xd0, 0xb0, 0x2e, 0x20, 0xd0, 0x92, 0xd0, 0xbe, 0xd0, 0xb7, 0xd0, 0xb2, 0xd1, 0x80,
	0xd0, 0xb0, 0xd1, 0x82, 0x20, 0xd0, 0xbc, 0xd0, 0xbe, 0xd0, 0xb6, 0xd0, 0xb5, 0xd1, 0x82, 0x20,
	0xd0, 0xb1, 0xd1, 0x8b, 0xd1, 0x82, 0xd1, 0x8c, 0x20, 0xd0, 0xbf, 0xd1, 0x80, 0xd0, 0xb8, 0xd0,
	0xbd, 0xd1, 0x8f, 0xd1, 0x82, 0x20, 0xd0, 0xb2, 0x20, 0xd1, 0x82, 0xd0, 0xb5, 0xd1, 0x87, 0xd0,
	0xb5, 0xd0, 0xbd, 0xd0, 0xb8, 0xd0, 0xb5, 0x20, 0xd0, 0xb4, 0xd0, 0xb2, 0xd1, 0x83, 0xd1, 0x85,
	0x20, 0xd0, 0xb4, 0xd0, 0xbd, 0xd0, 0xb5, 0xd0, 0xb9, 0x20, 0xd1, 0x81, 0x20, 0xd0, 0xbc, 0xd0,
	0xbe, 0xd0, 0xbc, 0xd0, 0xb5, 0xd0, 0xbd, 0xd1, 0x82, 0xd0, 0xb0, 0x20, 0xd0, 0xb2, 0xd1, 0x8b,
	0xd0, 0xb4, 0xd0, 0xb0, 0xd1, 0x87, 0xd0, 0xb8, 0x20, 0xd0, 0xb7, 0xd0, 0xb0, 0xd0, 0xba, 0xd0,
	0xb0, 0xd0, 0xb7, 0xd0, 0xb0, 0x2e, 0x20, 0xd0, 0x97, 0xd0, 0xb0, 0xd0, 0xba, 0xd0, 0xb0, 0xd0,
	0xb7, 0x20, 0xd0, 0xb4, 0xd0, 0xbe, 0xd0, 0xbb, 0xd0, 0xb6, 0xd0, 0xb5, 0xd0, 0xbd, 0x20, 0xd0,
	0xb1, 0xd1, 0x8b, 0xd1, 0x82, 0xd1, 0x8c, 0x20, 0xd0, 0xb2, 0xd1, 0x8b, 0xd0, 0xb4, 0xd0, 0xb0,
	0xd0, 0xbd, 0x20, 0xd0, 0xb8, 0xd0, 0xb7, 0x20, 0xd1, 0x8d, 0xd1, 0x82, 0xd0, 0xbe, 0xd0, 0xb3,
	0xd0, 0xbe, 0x20, 0xd0, 0x9f, 0xd0, 0x92, 0xd0, 0x97, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14,
	0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x2f, 0x61, 0x63,
	0x63, 0x65, 0x70, 0x74, 0x12, 0xcc, 0x02, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x13, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x98, 0x02, 0x92, 0x41, 0xfc, 0x01, 0x12,
	0x30, 0xd0, 0x9f, 0xd0, 0xbe, 0xd0, 0xbb, 0xd1, 0x83, 0xd1, 0x87, 0xd0, 0xb8, 0xd1, 0x82, 0xd1,
	0x8c, 0x20, 0xd1, 0x81, 0xd0, 0xbf, 0xd0, 0xb8, 0xd1, 0x81, 0xd0, 0xbe, 0xd0, 0xba, 0x20, 0xd0,
	0xb2, 0xd0, 0xbe, 0xd0, 0xb7, 0xd0, 0xb2, 0xd1, 0x80, 0xd0, 0xb0, 0xd1, 0x82, 0xd0, 0xbe, 0xd0,
	0xb2, 0x1a, 0xc7, 0x01, 0xd0, 0x98, 0xd1, 0x81, 0xd0, 0xbf, 0xd0, 0xbe, 0xd0, 0xbb, 0xd1, 0x8c,
	0xd0, 0xb7, 0xd1, 0x83, 0xd0, 0xb5, 0xd1, 0x82, 0xd1, 0x81, 0xd1, 0x8f, 0x20, 0xd0, 0xb4, 0xd0,
	0xbb, 0xd1, 0x8f, 0x20, 0xd0, 0xbf, 0xd0, 0xbe, 0xd0, 0xbb, 0xd1, 0x83, 0xd1, 0x87, 0xd0, 0xb5,
	0xd0, 0xbd, 0xd0, 0xb8, 0xd1, 0x8f, 0x20, 0xd1, 0x81, 0xd0, 0xbf, 0xd0, 0xb8, 0xd1, 0x81, 0xd0,
	0xba, 0xd0, 0xb0, 0x20, 0xd0, 0xb2, 0xd0, 0xbe, 0xd0, 0xb7, 0xd0, 0xb2, 0xd1, 0x80, 0xd0, 0xb0,
	0xd1, 0x82, 0xd0, 0xbe, 0xd0, 0xb2, 0x2e, 0x20, 0xd0, 0x9a, 0xd0, 0xbe, 0xd0, 0xbc, 0xd0, 0xb0,
	0xd0, 0xbd, 0xd0, 0xb4, 0xd0, 0xb0, 0x20, 0xd0, 0xb2, 0xd0, 0xbe, 0xd0, 0xb7, 0xd0, 0xb2, 0xd1,
	0x80, 0xd0, 0xb0, 0xd1, 0x89, 0xd0, 0xb0, 0xd0, 0xb5, 0xd1, 0x82, 0x20, 0xd1, 0x81, 0xd0, 0xbf,
	0xd0, 0xb8, 0xd1, 0x81, 0xd0, 0xbe, 0xd0, 0xba, 0x20, 0xd0, 0xb2, 0xd0, 0xbe, 0xd0, 0xb7, 0xd0,
	0xb2, 0xd1, 0x80, 0xd0, 0xb0, 0xd1, 0x82, 0xd0, 0xbe, 0xd0, 0xb2, 0x20, 0xd1, 0x81, 0x20, 0xd0,
	0xb2, 0xd0, 0xbe, 0xd0, 0xb7, 0xd0, 0xbc, 0xd0, 0xbe, 0xd0, 0xb6, 0xd0, 0xbd, 0xd0, 0xbe, 0xd1,
	0x81, 0xd1, 0x82, 0xd1, 0x8c, 0xd1, 0x8e, 0x20, 0xd0, 0xbf, 0xd0, 0xb0, 0xd0, 0xb3, 0xd0, 0xb8,
	0xd0, 0xbd, 0xd0, 0xb0, 0xd1, 0x86, 0xd0, 0xb8, 0xd0, 0xb8, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x12, 0x12, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6f, 0x7a,
	0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x79, 0x75, 0x77, 0x65, 0x65, 0x62, 0x69, 0x78, 0x2f,
	0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2d, 0x31, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x3b, 0x72, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_returns_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_returns_proto_goTypes = []any{
	(*AcceptRequest)(nil),         // 0: orders.AcceptRequest
	(*AcceptResponse)(nil),        // 1: orders.AcceptResponse
	(*ListRequest)(nil),           // 2: orders.ListRequest
	(*ListResponse)(nil),          // 3: orders.ListResponse
	(*ListResponse_Order)(nil),    // 4: orders.ListResponse.Order
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_returns_proto_depIdxs = []int32{
	4, // 0: orders.ListResponse.orders:type_name -> orders.ListResponse.Order
	5, // 1: orders.ListResponse.Order.expiry:type_name -> google.protobuf.Timestamp
	5, // 2: orders.ListResponse.Order.return_by:type_name -> google.protobuf.Timestamp
	5, // 3: orders.ListResponse.Order.created_at:type_name -> google.protobuf.Timestamp
	0, // 4: orders.Returns.Accept:input_type -> orders.AcceptRequest
	2, // 5: orders.Returns.List:input_type -> orders.ListRequest
	1, // 6: orders.Returns.Accept:output_type -> orders.AcceptResponse
	3, // 7: orders.Returns.List:output_type -> orders.ListResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_returns_proto_init() }
func file_returns_proto_init() {
	if File_returns_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_returns_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AcceptRequest); i {
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
			switch v := v.(*AcceptResponse); i {
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
			switch v := v.(*ListRequest); i {
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
			switch v := v.(*ListResponse); i {
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
			switch v := v.(*ListResponse_Order); i {
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
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_returns_proto_goTypes,
		DependencyIndexes: file_returns_proto_depIdxs,
		MessageInfos:      file_returns_proto_msgTypes,
	}.Build()
	File_returns_proto = out.File
	file_returns_proto_rawDesc = nil
	file_returns_proto_goTypes = nil
	file_returns_proto_depIdxs = nil
}
