// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: card.proto

package api

import (
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

// Card represents bank card instance.
type Card struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                            // card name
	Card      []byte                 `protobuf:"bytes,2,opt,name=card,proto3" json:"card,omitempty"`                            // bank card
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"` // update or creation time
	Meta      []byte                 `protobuf:"bytes,4,opt,name=meta,proto3" json:"meta,omitempty"`                            // metadata
}

func (x *Card) Reset() {
	*x = Card{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Card) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Card) ProtoMessage() {}

func (x *Card) ProtoReflect() protoreflect.Message {
	mi := &file_card_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Card.ProtoReflect.Descriptor instead.
func (*Card) Descriptor() ([]byte, []int) {
	return file_card_proto_rawDescGZIP(), []int{0}
}

func (x *Card) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Card) GetCard() []byte {
	if x != nil {
		return x.Card
	}
	return nil
}

func (x *Card) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Card) GetMeta() []byte {
	if x != nil {
		return x.Meta
	}
	return nil
}

// Insert request.
type InsertCardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Card *Card `protobuf:"bytes,2,opt,name=card,proto3" json:"card,omitempty"`
}

func (x *InsertCardRequest) Reset() {
	*x = InsertCardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertCardRequest) ProtoMessage() {}

func (x *InsertCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_card_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertCardRequest.ProtoReflect.Descriptor instead.
func (*InsertCardRequest) Descriptor() ([]byte, []int) {
	return file_card_proto_rawDescGZIP(), []int{1}
}

func (x *InsertCardRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *InsertCardRequest) GetCard() *Card {
	if x != nil {
		return x.Card
	}
	return nil
}

// Get response.
type GetCardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Card  *Card  `protobuf:"bytes,1,opt,name=card,proto3" json:"card,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetCardResponse) Reset() {
	*x = GetCardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCardResponse) ProtoMessage() {}

func (x *GetCardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_card_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCardResponse.ProtoReflect.Descriptor instead.
func (*GetCardResponse) Descriptor() ([]byte, []int) {
	return file_card_proto_rawDescGZIP(), []int{2}
}

func (x *GetCardResponse) GetCard() *Card {
	if x != nil {
		return x.Card
	}
	return nil
}

func (x *GetCardResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// Update request.
type UpdateCardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Card *Card  `protobuf:"bytes,3,opt,name=card,proto3" json:"card,omitempty"`
}

func (x *UpdateCardRequest) Reset() {
	*x = UpdateCardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCardRequest) ProtoMessage() {}

func (x *UpdateCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_card_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCardRequest.ProtoReflect.Descriptor instead.
func (*UpdateCardRequest) Descriptor() ([]byte, []int) {
	return file_card_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateCardRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UpdateCardRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateCardRequest) GetCard() *Card {
	if x != nil {
		return x.Card
	}
	return nil
}

// Update response.
type UpdateCardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Card  *Card  `protobuf:"bytes,1,opt,name=card,proto3" json:"card,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *UpdateCardResponse) Reset() {
	*x = UpdateCardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCardResponse) ProtoMessage() {}

func (x *UpdateCardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_card_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCardResponse.ProtoReflect.Descriptor instead.
func (*UpdateCardResponse) Descriptor() ([]byte, []int) {
	return file_card_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCardResponse) GetCard() *Card {
	if x != nil {
		return x.Card
	}
	return nil
}

func (x *UpdateCardResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// Delete request.
type DeleteCardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteCardRequest) Reset() {
	*x = DeleteCardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCardRequest) ProtoMessage() {}

func (x *DeleteCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_card_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCardRequest.ProtoReflect.Descriptor instead.
func (*DeleteCardRequest) Descriptor() ([]byte, []int) {
	return file_card_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteCardRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *DeleteCardRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// List request.
type ListCardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *ListCardRequest) Reset() {
	*x = ListCardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCardRequest) ProtoMessage() {}

func (x *ListCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_card_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCardRequest.ProtoReflect.Descriptor instead.
func (*ListCardRequest) Descriptor() ([]byte, []int) {
	return file_card_proto_rawDescGZIP(), []int{6}
}

func (x *ListCardRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

// List response.
type ListCardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cards []*Card `protobuf:"bytes,1,rep,name=cards,proto3" json:"cards,omitempty"`
}

func (x *ListCardResponse) Reset() {
	*x = ListCardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCardResponse) ProtoMessage() {}

func (x *ListCardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_card_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCardResponse.ProtoReflect.Descriptor instead.
func (*ListCardResponse) Descriptor() ([]byte, []int) {
	return file_card_proto_rawDescGZIP(), []int{7}
}

func (x *ListCardResponse) GetCards() []*Card {
	if x != nil {
		return x.Cards
	}
	return nil
}

var File_card_proto protoreflect.FileDescriptor

var file_card_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70,
	0x69, 0x1a, 0x12, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x6c, 0x6f, 0x67, 0x5f, 0x70, 0x61, 0x73, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x63, 0x61, 0x72, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x51, 0x0a, 0x11, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x43,
	0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x04, 0x63, 0x61, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x61,
	0x72, 0x64, 0x52, 0x04, 0x63, 0x61, 0x72, 0x64, 0x22, 0x46, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x63,
	0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x43, 0x61, 0x72, 0x64, 0x52, 0x04, 0x63, 0x61, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x22, 0x65, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x61, 0x72,
	0x64, 0x52, 0x04, 0x63, 0x61, 0x72, 0x64, 0x22, 0x49, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a,
	0x04, 0x63, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x04, 0x63, 0x61, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x46, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x30, 0x0a, 0x0f, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x33, 0x0a, 0x10,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x05, 0x63, 0x61, 0x72, 0x64,
	0x73, 0x32, 0x93, 0x02, 0x0a, 0x05, 0x43, 0x61, 0x72, 0x64, 0x73, 0x12, 0x35, 0x0a, 0x06, 0x49,
	0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6e, 0x73, 0x65,
	0x72, 0x74, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x39, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x73, 0x68, 0x61, 0x6e, 0x6b, 0x6f, 0x47, 0x4f,
	0x2f, 0x47, 0x6f, 0x70, 0x68, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_card_proto_rawDescOnce sync.Once
	file_card_proto_rawDescData = file_card_proto_rawDesc
)

func file_card_proto_rawDescGZIP() []byte {
	file_card_proto_rawDescOnce.Do(func() {
		file_card_proto_rawDescData = protoimpl.X.CompressGZIP(file_card_proto_rawDescData)
	})
	return file_card_proto_rawDescData
}

var file_card_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_card_proto_goTypes = []interface{}{
	(*Card)(nil),                  // 0: api.Card
	(*InsertCardRequest)(nil),     // 1: api.InsertCardRequest
	(*GetCardResponse)(nil),       // 2: api.GetCardResponse
	(*UpdateCardRequest)(nil),     // 3: api.UpdateCardRequest
	(*UpdateCardResponse)(nil),    // 4: api.UpdateCardResponse
	(*DeleteCardRequest)(nil),     // 5: api.DeleteCardRequest
	(*ListCardRequest)(nil),       // 6: api.ListCardRequest
	(*ListCardResponse)(nil),      // 7: api.ListCardResponse
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*User)(nil),                  // 9: api.User
	(*GetRequest)(nil),            // 10: api.GetRequest
	(*InsertResponse)(nil),        // 11: api.InsertResponse
	(*DeleteResponse)(nil),        // 12: api.DeleteResponse
}
var file_card_proto_depIdxs = []int32{
	8,  // 0: api.Card.updated_at:type_name -> google.protobuf.Timestamp
	9,  // 1: api.InsertCardRequest.user:type_name -> api.User
	0,  // 2: api.InsertCardRequest.card:type_name -> api.Card
	0,  // 3: api.GetCardResponse.card:type_name -> api.Card
	9,  // 4: api.UpdateCardRequest.user:type_name -> api.User
	0,  // 5: api.UpdateCardRequest.card:type_name -> api.Card
	0,  // 6: api.UpdateCardResponse.card:type_name -> api.Card
	9,  // 7: api.DeleteCardRequest.user:type_name -> api.User
	9,  // 8: api.ListCardRequest.user:type_name -> api.User
	0,  // 9: api.ListCardResponse.cards:type_name -> api.Card
	1,  // 10: api.Cards.Insert:input_type -> api.InsertCardRequest
	10, // 11: api.Cards.Get:input_type -> api.GetRequest
	3,  // 12: api.Cards.Update:input_type -> api.UpdateCardRequest
	5,  // 13: api.Cards.Delete:input_type -> api.DeleteCardRequest
	6,  // 14: api.Cards.List:input_type -> api.ListCardRequest
	11, // 15: api.Cards.Insert:output_type -> api.InsertResponse
	2,  // 16: api.Cards.Get:output_type -> api.GetCardResponse
	4,  // 17: api.Cards.Update:output_type -> api.UpdateCardResponse
	12, // 18: api.Cards.Delete:output_type -> api.DeleteResponse
	7,  // 19: api.Cards.List:output_type -> api.ListCardResponse
	15, // [15:20] is the sub-list for method output_type
	10, // [10:15] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_card_proto_init() }
func file_card_proto_init() {
	if File_card_proto != nil {
		return
	}
	file_registration_proto_init()
	file_log_pass_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_card_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Card); i {
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
		file_card_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertCardRequest); i {
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
		file_card_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCardResponse); i {
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
		file_card_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCardRequest); i {
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
		file_card_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCardResponse); i {
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
		file_card_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCardRequest); i {
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
		file_card_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCardRequest); i {
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
		file_card_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCardResponse); i {
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
			RawDescriptor: file_card_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_card_proto_goTypes,
		DependencyIndexes: file_card_proto_depIdxs,
		MessageInfos:      file_card_proto_msgTypes,
	}.Build()
	File_card_proto = out.File
	file_card_proto_rawDesc = nil
	file_card_proto_goTypes = nil
	file_card_proto_depIdxs = nil
}
