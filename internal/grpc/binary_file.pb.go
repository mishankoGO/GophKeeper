// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: binary_file.proto

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

// BinaryFile represents binary file instance.
type BinaryFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                            // binary file name
	File      []byte                 `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`                            // binary file
	Extension []byte                 `protobuf:"bytes,3,opt,name=extension,proto3" json:"extension,omitempty"`                  // binary file extension
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"` // time of the update or creation
	Meta      []byte                 `protobuf:"bytes,5,opt,name=meta,proto3" json:"meta,omitempty"`                            // metadata
}

func (x *BinaryFile) Reset() {
	*x = BinaryFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_binary_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinaryFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryFile) ProtoMessage() {}

func (x *BinaryFile) ProtoReflect() protoreflect.Message {
	mi := &file_binary_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryFile.ProtoReflect.Descriptor instead.
func (*BinaryFile) Descriptor() ([]byte, []int) {
	return file_binary_file_proto_rawDescGZIP(), []int{0}
}

func (x *BinaryFile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BinaryFile) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *BinaryFile) GetExtension() []byte {
	if x != nil {
		return x.Extension
	}
	return nil
}

func (x *BinaryFile) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *BinaryFile) GetMeta() []byte {
	if x != nil {
		return x.Meta
	}
	return nil
}

// Insert request.
type InsertBinaryFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User       `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	File *BinaryFile `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *InsertBinaryFileRequest) Reset() {
	*x = InsertBinaryFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_binary_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertBinaryFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertBinaryFileRequest) ProtoMessage() {}

func (x *InsertBinaryFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_binary_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertBinaryFileRequest.ProtoReflect.Descriptor instead.
func (*InsertBinaryFileRequest) Descriptor() ([]byte, []int) {
	return file_binary_file_proto_rawDescGZIP(), []int{1}
}

func (x *InsertBinaryFileRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *InsertBinaryFileRequest) GetFile() *BinaryFile {
	if x != nil {
		return x.File
	}
	return nil
}

// Get response.
type GetBinaryFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File *BinaryFile `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *GetBinaryFileResponse) Reset() {
	*x = GetBinaryFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_binary_file_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBinaryFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBinaryFileResponse) ProtoMessage() {}

func (x *GetBinaryFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_binary_file_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBinaryFileResponse.ProtoReflect.Descriptor instead.
func (*GetBinaryFileResponse) Descriptor() ([]byte, []int) {
	return file_binary_file_proto_rawDescGZIP(), []int{2}
}

func (x *GetBinaryFileResponse) GetFile() *BinaryFile {
	if x != nil {
		return x.File
	}
	return nil
}

// Update request.
type UpdateBinaryFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User       `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	File *BinaryFile `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *UpdateBinaryFileRequest) Reset() {
	*x = UpdateBinaryFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_binary_file_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBinaryFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBinaryFileRequest) ProtoMessage() {}

func (x *UpdateBinaryFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_binary_file_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBinaryFileRequest.ProtoReflect.Descriptor instead.
func (*UpdateBinaryFileRequest) Descriptor() ([]byte, []int) {
	return file_binary_file_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateBinaryFileRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UpdateBinaryFileRequest) GetFile() *BinaryFile {
	if x != nil {
		return x.File
	}
	return nil
}

// Update response.
type UpdateBinaryFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File *BinaryFile `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *UpdateBinaryFileResponse) Reset() {
	*x = UpdateBinaryFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_binary_file_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBinaryFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBinaryFileResponse) ProtoMessage() {}

func (x *UpdateBinaryFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_binary_file_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBinaryFileResponse.ProtoReflect.Descriptor instead.
func (*UpdateBinaryFileResponse) Descriptor() ([]byte, []int) {
	return file_binary_file_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateBinaryFileResponse) GetFile() *BinaryFile {
	if x != nil {
		return x.File
	}
	return nil
}

// Delete request.
type DeleteBinaryFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteBinaryFileRequest) Reset() {
	*x = DeleteBinaryFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_binary_file_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBinaryFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBinaryFileRequest) ProtoMessage() {}

func (x *DeleteBinaryFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_binary_file_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBinaryFileRequest.ProtoReflect.Descriptor instead.
func (*DeleteBinaryFileRequest) Descriptor() ([]byte, []int) {
	return file_binary_file_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteBinaryFileRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *DeleteBinaryFileRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// List request.
type ListBinaryFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *ListBinaryFileRequest) Reset() {
	*x = ListBinaryFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_binary_file_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBinaryFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBinaryFileRequest) ProtoMessage() {}

func (x *ListBinaryFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_binary_file_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBinaryFileRequest.ProtoReflect.Descriptor instead.
func (*ListBinaryFileRequest) Descriptor() ([]byte, []int) {
	return file_binary_file_proto_rawDescGZIP(), []int{6}
}

func (x *ListBinaryFileRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

// List response.
type ListBinaryFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BinaryFiles []*BinaryFile `protobuf:"bytes,1,rep,name=binary_files,json=binaryFiles,proto3" json:"binary_files,omitempty"`
}

func (x *ListBinaryFileResponse) Reset() {
	*x = ListBinaryFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_binary_file_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBinaryFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBinaryFileResponse) ProtoMessage() {}

func (x *ListBinaryFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_binary_file_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBinaryFileResponse.ProtoReflect.Descriptor instead.
func (*ListBinaryFileResponse) Descriptor() ([]byte, []int) {
	return file_binary_file_proto_rawDescGZIP(), []int{7}
}

func (x *ListBinaryFileResponse) GetBinaryFiles() []*BinaryFile {
	if x != nil {
		return x.BinaryFiles
	}
	return nil
}

var File_binary_file_proto protoreflect.FileDescriptor

var file_binary_file_proto_rawDesc = []byte{
	0x0a, 0x11, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x6c,
	0x6f, 0x67, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x01,
	0x0a, 0x0a, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x22, 0x5d, 0x0a, 0x17, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x42, 0x69, 0x6e, 0x61, 0x72,
	0x79, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x04, 0x66,
	0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65,
	0x22, 0x3c, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x66, 0x69, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x69,
	0x6e, 0x61, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x5d,
	0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x69, 0x6e,
	0x61, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x3f, 0x0a,
	0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x66, 0x69, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x69,
	0x6e, 0x61, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x4c,
	0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x36, 0x0a, 0x15,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x22, 0x4c, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x69, 0x6e, 0x61,
	0x72, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32,
	0x0a, 0x0c, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72,
	0x79, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x0b, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x46, 0x69, 0x6c,
	0x65, 0x73, 0x32, 0xc3, 0x02, 0x0a, 0x0b, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x46, 0x69, 0x6c,
	0x65, 0x73, 0x12, 0x3b, 0x0a, 0x06, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x1c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x32, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x73, 0x68, 0x61, 0x6e, 0x6b, 0x6f, 0x47,
	0x4f, 0x2f, 0x47, 0x6f, 0x70, 0x68, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_binary_file_proto_rawDescOnce sync.Once
	file_binary_file_proto_rawDescData = file_binary_file_proto_rawDesc
)

func file_binary_file_proto_rawDescGZIP() []byte {
	file_binary_file_proto_rawDescOnce.Do(func() {
		file_binary_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_binary_file_proto_rawDescData)
	})
	return file_binary_file_proto_rawDescData
}

var file_binary_file_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_binary_file_proto_goTypes = []interface{}{
	(*BinaryFile)(nil),               // 0: api.BinaryFile
	(*InsertBinaryFileRequest)(nil),  // 1: api.InsertBinaryFileRequest
	(*GetBinaryFileResponse)(nil),    // 2: api.GetBinaryFileResponse
	(*UpdateBinaryFileRequest)(nil),  // 3: api.UpdateBinaryFileRequest
	(*UpdateBinaryFileResponse)(nil), // 4: api.UpdateBinaryFileResponse
	(*DeleteBinaryFileRequest)(nil),  // 5: api.DeleteBinaryFileRequest
	(*ListBinaryFileRequest)(nil),    // 6: api.ListBinaryFileRequest
	(*ListBinaryFileResponse)(nil),   // 7: api.ListBinaryFileResponse
	(*timestamppb.Timestamp)(nil),    // 8: google.protobuf.Timestamp
	(*User)(nil),                     // 9: api.User
	(*GetRequest)(nil),               // 10: api.GetRequest
	(*InsertResponse)(nil),           // 11: api.InsertResponse
	(*DeleteResponse)(nil),           // 12: api.DeleteResponse
}
var file_binary_file_proto_depIdxs = []int32{
	8,  // 0: api.BinaryFile.updated_at:type_name -> google.protobuf.Timestamp
	9,  // 1: api.InsertBinaryFileRequest.user:type_name -> api.User
	0,  // 2: api.InsertBinaryFileRequest.file:type_name -> api.BinaryFile
	0,  // 3: api.GetBinaryFileResponse.file:type_name -> api.BinaryFile
	9,  // 4: api.UpdateBinaryFileRequest.user:type_name -> api.User
	0,  // 5: api.UpdateBinaryFileRequest.file:type_name -> api.BinaryFile
	0,  // 6: api.UpdateBinaryFileResponse.file:type_name -> api.BinaryFile
	9,  // 7: api.DeleteBinaryFileRequest.user:type_name -> api.User
	9,  // 8: api.ListBinaryFileRequest.user:type_name -> api.User
	0,  // 9: api.ListBinaryFileResponse.binary_files:type_name -> api.BinaryFile
	1,  // 10: api.BinaryFiles.Insert:input_type -> api.InsertBinaryFileRequest
	10, // 11: api.BinaryFiles.Get:input_type -> api.GetRequest
	3,  // 12: api.BinaryFiles.Update:input_type -> api.UpdateBinaryFileRequest
	5,  // 13: api.BinaryFiles.Delete:input_type -> api.DeleteBinaryFileRequest
	6,  // 14: api.BinaryFiles.List:input_type -> api.ListBinaryFileRequest
	11, // 15: api.BinaryFiles.Insert:output_type -> api.InsertResponse
	2,  // 16: api.BinaryFiles.Get:output_type -> api.GetBinaryFileResponse
	4,  // 17: api.BinaryFiles.Update:output_type -> api.UpdateBinaryFileResponse
	12, // 18: api.BinaryFiles.Delete:output_type -> api.DeleteResponse
	7,  // 19: api.BinaryFiles.List:output_type -> api.ListBinaryFileResponse
	15, // [15:20] is the sub-list for method output_type
	10, // [10:15] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_binary_file_proto_init() }
func file_binary_file_proto_init() {
	if File_binary_file_proto != nil {
		return
	}
	file_registration_proto_init()
	file_log_pass_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_binary_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinaryFile); i {
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
		file_binary_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertBinaryFileRequest); i {
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
		file_binary_file_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBinaryFileResponse); i {
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
		file_binary_file_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBinaryFileRequest); i {
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
		file_binary_file_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBinaryFileResponse); i {
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
		file_binary_file_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBinaryFileRequest); i {
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
		file_binary_file_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBinaryFileRequest); i {
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
		file_binary_file_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBinaryFileResponse); i {
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
			RawDescriptor: file_binary_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_binary_file_proto_goTypes,
		DependencyIndexes: file_binary_file_proto_depIdxs,
		MessageInfos:      file_binary_file_proto_msgTypes,
	}.Build()
	File_binary_file_proto = out.File
	file_binary_file_proto_rawDesc = nil
	file_binary_file_proto_goTypes = nil
	file_binary_file_proto_depIdxs = nil
}
