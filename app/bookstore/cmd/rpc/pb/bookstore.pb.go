// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: bookstore.proto

package pb

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

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price int64  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	Desc  string `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Book) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Book) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Book) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type CreateBookReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price int64  `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	Desc  string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *CreateBookReq) Reset() {
	*x = CreateBookReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookReq) ProtoMessage() {}

func (x *CreateBookReq) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookReq.ProtoReflect.Descriptor instead.
func (*CreateBookReq) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBookReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateBookReq) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateBookReq) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type CreateBookResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateBookResp) Reset() {
	*x = CreateBookResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookResp) ProtoMessage() {}

func (x *CreateBookResp) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookResp.ProtoReflect.Descriptor instead.
func (*CreateBookResp) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{2}
}

func (x *CreateBookResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateBookReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price int64  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	Desc  string `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *UpdateBookReq) Reset() {
	*x = UpdateBookReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBookReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookReq) ProtoMessage() {}

func (x *UpdateBookReq) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookReq.ProtoReflect.Descriptor instead.
func (*UpdateBookReq) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateBookReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateBookReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateBookReq) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdateBookReq) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type UpdateBookResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateBookResp) Reset() {
	*x = UpdateBookResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBookResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookResp) ProtoMessage() {}

func (x *UpdateBookResp) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookResp.ProtoReflect.Descriptor instead.
func (*UpdateBookResp) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateBookResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteBookReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteBookReq) Reset() {
	*x = DeleteBookReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBookReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookReq) ProtoMessage() {}

func (x *DeleteBookReq) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookReq.ProtoReflect.Descriptor instead.
func (*DeleteBookReq) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteBookReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteBookResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteBookResp) Reset() {
	*x = DeleteBookResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBookResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookResp) ProtoMessage() {}

func (x *DeleteBookResp) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookResp.ProtoReflect.Descriptor instead.
func (*DeleteBookResp) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteBookResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetBookReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBookReq) Reset() {
	*x = GetBookReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookReq) ProtoMessage() {}

func (x *GetBookReq) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookReq.ProtoReflect.Descriptor instead.
func (*GetBookReq) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{7}
}

func (x *GetBookReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetBookResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Book *Book `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *GetBookResp) Reset() {
	*x = GetBookResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookResp) ProtoMessage() {}

func (x *GetBookResp) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookResp.ProtoReflect.Descriptor instead.
func (*GetBookResp) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{8}
}

func (x *GetBookResp) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

type GetBookListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Price    *int64  `protobuf:"varint,2,opt,name=price,proto3,oneof" json:"price,omitempty"`
	OrderBy  *string `protobuf:"bytes,3,opt,name=order_by,json=orderBy,proto3,oneof" json:"order_by,omitempty"`
	Page     *int64  `protobuf:"varint,4,opt,name=page,proto3,oneof" json:"page,omitempty"`
	PageSize *int64  `protobuf:"varint,5,opt,name=page_size,json=pageSize,proto3,oneof" json:"page_size,omitempty"`
}

func (x *GetBookListReq) Reset() {
	*x = GetBookListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookListReq) ProtoMessage() {}

func (x *GetBookListReq) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookListReq.ProtoReflect.Descriptor instead.
func (*GetBookListReq) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{9}
}

func (x *GetBookListReq) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *GetBookListReq) GetPrice() int64 {
	if x != nil && x.Price != nil {
		return *x.Price
	}
	return 0
}

func (x *GetBookListReq) GetOrderBy() string {
	if x != nil && x.OrderBy != nil {
		return *x.OrderBy
	}
	return ""
}

func (x *GetBookListReq) GetPage() int64 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *GetBookListReq) GetPageSize() int64 {
	if x != nil && x.PageSize != nil {
		return *x.PageSize
	}
	return 0
}

type GetBookListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*Book `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *GetBookListResp) Reset() {
	*x = GetBookListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookListResp) ProtoMessage() {}

func (x *GetBookListResp) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookListResp.ProtoReflect.Descriptor instead.
func (*GetBookListResp) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{10}
}

func (x *GetBookListResp) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

var File_bookstore_proto protoreflect.FileDescriptor

var file_bookstore_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x54, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x22, 0x4d, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x22, 0x20, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5d, 0x0a, 0x0d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x22, 0x20, 0x0a, 0x0e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1f, 0x0a,
	0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x20,
	0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x1c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2b,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a,
	0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62,
	0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0xd6, 0x01, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x17,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x1e, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x88,
	0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x48, 0x04,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x22, 0x31, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x32, 0x98, 0x02, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x6f, 0x6f, 0x6b, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e,
	0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f,
	0x6b, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42,
	0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_bookstore_proto_rawDescOnce sync.Once
	file_bookstore_proto_rawDescData = file_bookstore_proto_rawDesc
)

func file_bookstore_proto_rawDescGZIP() []byte {
	file_bookstore_proto_rawDescOnce.Do(func() {
		file_bookstore_proto_rawDescData = protoimpl.X.CompressGZIP(file_bookstore_proto_rawDescData)
	})
	return file_bookstore_proto_rawDescData
}

var file_bookstore_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_bookstore_proto_goTypes = []interface{}{
	(*Book)(nil),            // 0: pb.Book
	(*CreateBookReq)(nil),   // 1: pb.CreateBookReq
	(*CreateBookResp)(nil),  // 2: pb.CreateBookResp
	(*UpdateBookReq)(nil),   // 3: pb.UpdateBookReq
	(*UpdateBookResp)(nil),  // 4: pb.UpdateBookResp
	(*DeleteBookReq)(nil),   // 5: pb.DeleteBookReq
	(*DeleteBookResp)(nil),  // 6: pb.DeleteBookResp
	(*GetBookReq)(nil),      // 7: pb.GetBookReq
	(*GetBookResp)(nil),     // 8: pb.GetBookResp
	(*GetBookListReq)(nil),  // 9: pb.GetBookListReq
	(*GetBookListResp)(nil), // 10: pb.GetBookListResp
}
var file_bookstore_proto_depIdxs = []int32{
	0,  // 0: pb.GetBookResp.book:type_name -> pb.Book
	0,  // 1: pb.GetBookListResp.books:type_name -> pb.Book
	1,  // 2: pb.bookstore.CreateBook:input_type -> pb.CreateBookReq
	3,  // 3: pb.bookstore.UpdateBook:input_type -> pb.UpdateBookReq
	5,  // 4: pb.bookstore.DeleteBook:input_type -> pb.DeleteBookReq
	7,  // 5: pb.bookstore.GetBook:input_type -> pb.GetBookReq
	9,  // 6: pb.bookstore.GetBookList:input_type -> pb.GetBookListReq
	2,  // 7: pb.bookstore.CreateBook:output_type -> pb.CreateBookResp
	4,  // 8: pb.bookstore.UpdateBook:output_type -> pb.UpdateBookResp
	6,  // 9: pb.bookstore.DeleteBook:output_type -> pb.DeleteBookResp
	8,  // 10: pb.bookstore.GetBook:output_type -> pb.GetBookResp
	10, // 11: pb.bookstore.GetBookList:output_type -> pb.GetBookListResp
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_bookstore_proto_init() }
func file_bookstore_proto_init() {
	if File_bookstore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bookstore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
		file_bookstore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBookReq); i {
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
		file_bookstore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBookResp); i {
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
		file_bookstore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBookReq); i {
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
		file_bookstore_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBookResp); i {
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
		file_bookstore_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBookReq); i {
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
		file_bookstore_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBookResp); i {
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
		file_bookstore_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookReq); i {
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
		file_bookstore_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookResp); i {
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
		file_bookstore_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookListReq); i {
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
		file_bookstore_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookListResp); i {
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
	file_bookstore_proto_msgTypes[9].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bookstore_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bookstore_proto_goTypes,
		DependencyIndexes: file_bookstore_proto_depIdxs,
		MessageInfos:      file_bookstore_proto_msgTypes,
	}.Build()
	File_bookstore_proto = out.File
	file_bookstore_proto_rawDesc = nil
	file_bookstore_proto_goTypes = nil
	file_bookstore_proto_depIdxs = nil
}
