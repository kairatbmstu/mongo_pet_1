// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: mongo_pet.proto

package grpcgen

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Link struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url    string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Status string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Link) Reset() {
	*x = Link{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mongo_pet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Link) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Link) ProtoMessage() {}

func (x *Link) ProtoReflect() protoreflect.Message {
	mi := &file_mongo_pet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Link.ProtoReflect.Descriptor instead.
func (*Link) Descriptor() ([]byte, []int) {
	return file_mongo_pet_proto_rawDescGZIP(), []int{0}
}

func (x *Link) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Link) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Link) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type CreatePageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url      string               `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Created  *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated  *timestamp.Timestamp `protobuf:"bytes,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Status   string               `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Html     string               `protobuf:"bytes,6,opt,name=html,proto3" json:"html,omitempty"`
	PageRank float64              `protobuf:"fixed64,7,opt,name=page_rank,json=pageRank,proto3" json:"page_rank,omitempty"`
	Link2    []*Link              `protobuf:"bytes,8,rep,name=link2,proto3" json:"link2,omitempty"`
}

func (x *CreatePageRequest) Reset() {
	*x = CreatePageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mongo_pet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePageRequest) ProtoMessage() {}

func (x *CreatePageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mongo_pet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePageRequest.ProtoReflect.Descriptor instead.
func (*CreatePageRequest) Descriptor() ([]byte, []int) {
	return file_mongo_pet_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePageRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreatePageRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CreatePageRequest) GetCreated() *timestamp.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *CreatePageRequest) GetUpdated() *timestamp.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

func (x *CreatePageRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreatePageRequest) GetHtml() string {
	if x != nil {
		return x.Html
	}
	return ""
}

func (x *CreatePageRequest) GetPageRank() float64 {
	if x != nil {
		return x.PageRank
	}
	return 0
}

func (x *CreatePageRequest) GetLink2() []*Link {
	if x != nil {
		return x.Link2
	}
	return nil
}

type CreatePageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url      string               `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Created  *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated  *timestamp.Timestamp `protobuf:"bytes,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Status   string               `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Html     string               `protobuf:"bytes,6,opt,name=html,proto3" json:"html,omitempty"`
	PageRank float64              `protobuf:"fixed64,7,opt,name=page_rank,json=pageRank,proto3" json:"page_rank,omitempty"`
	Link2    []*Link              `protobuf:"bytes,8,rep,name=link2,proto3" json:"link2,omitempty"`
}

func (x *CreatePageResponse) Reset() {
	*x = CreatePageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mongo_pet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePageResponse) ProtoMessage() {}

func (x *CreatePageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mongo_pet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePageResponse.ProtoReflect.Descriptor instead.
func (*CreatePageResponse) Descriptor() ([]byte, []int) {
	return file_mongo_pet_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePageResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreatePageResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CreatePageResponse) GetCreated() *timestamp.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *CreatePageResponse) GetUpdated() *timestamp.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

func (x *CreatePageResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreatePageResponse) GetHtml() string {
	if x != nil {
		return x.Html
	}
	return ""
}

func (x *CreatePageResponse) GetPageRank() float64 {
	if x != nil {
		return x.PageRank
	}
	return 0
}

func (x *CreatePageResponse) GetLink2() []*Link {
	if x != nil {
		return x.Link2
	}
	return nil
}

type UpdatePageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url      string               `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Created  *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated  *timestamp.Timestamp `protobuf:"bytes,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Status   string               `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Html     string               `protobuf:"bytes,6,opt,name=html,proto3" json:"html,omitempty"`
	PageRank float64              `protobuf:"fixed64,7,opt,name=page_rank,json=pageRank,proto3" json:"page_rank,omitempty"`
	Link2    []*Link              `protobuf:"bytes,8,rep,name=link2,proto3" json:"link2,omitempty"`
}

func (x *UpdatePageRequest) Reset() {
	*x = UpdatePageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mongo_pet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePageRequest) ProtoMessage() {}

func (x *UpdatePageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mongo_pet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePageRequest.ProtoReflect.Descriptor instead.
func (*UpdatePageRequest) Descriptor() ([]byte, []int) {
	return file_mongo_pet_proto_rawDescGZIP(), []int{3}
}

func (x *UpdatePageRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdatePageRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UpdatePageRequest) GetCreated() *timestamp.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *UpdatePageRequest) GetUpdated() *timestamp.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

func (x *UpdatePageRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *UpdatePageRequest) GetHtml() string {
	if x != nil {
		return x.Html
	}
	return ""
}

func (x *UpdatePageRequest) GetPageRank() float64 {
	if x != nil {
		return x.PageRank
	}
	return 0
}

func (x *UpdatePageRequest) GetLink2() []*Link {
	if x != nil {
		return x.Link2
	}
	return nil
}

type UpdatePageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url      string               `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Created  *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated  *timestamp.Timestamp `protobuf:"bytes,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Status   string               `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Html     string               `protobuf:"bytes,6,opt,name=html,proto3" json:"html,omitempty"`
	PageRank float64              `protobuf:"fixed64,7,opt,name=page_rank,json=pageRank,proto3" json:"page_rank,omitempty"`
	Link2    []*Link              `protobuf:"bytes,8,rep,name=link2,proto3" json:"link2,omitempty"`
}

func (x *UpdatePageResponse) Reset() {
	*x = UpdatePageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mongo_pet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePageResponse) ProtoMessage() {}

func (x *UpdatePageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mongo_pet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePageResponse.ProtoReflect.Descriptor instead.
func (*UpdatePageResponse) Descriptor() ([]byte, []int) {
	return file_mongo_pet_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatePageResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdatePageResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UpdatePageResponse) GetCreated() *timestamp.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *UpdatePageResponse) GetUpdated() *timestamp.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

func (x *UpdatePageResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *UpdatePageResponse) GetHtml() string {
	if x != nil {
		return x.Html
	}
	return ""
}

func (x *UpdatePageResponse) GetPageRank() float64 {
	if x != nil {
		return x.PageRank
	}
	return 0
}

func (x *UpdatePageResponse) GetLink2() []*Link {
	if x != nil {
		return x.Link2
	}
	return nil
}

type DeletePageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePageRequest) Reset() {
	*x = DeletePageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mongo_pet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePageRequest) ProtoMessage() {}

func (x *DeletePageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mongo_pet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePageRequest.ProtoReflect.Descriptor instead.
func (*DeletePageRequest) Descriptor() ([]byte, []int) {
	return file_mongo_pet_proto_rawDescGZIP(), []int{5}
}

func (x *DeletePageRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeletePageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePageResponse) Reset() {
	*x = DeletePageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mongo_pet_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePageResponse) ProtoMessage() {}

func (x *DeletePageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mongo_pet_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePageResponse.ProtoReflect.Descriptor instead.
func (*DeletePageResponse) Descriptor() ([]byte, []int) {
	return file_mongo_pet_proto_rawDescGZIP(), []int{6}
}

func (x *DeletePageResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FindByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindByIdRequest) Reset() {
	*x = FindByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mongo_pet_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByIdRequest) ProtoMessage() {}

func (x *FindByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mongo_pet_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByIdRequest.ProtoReflect.Descriptor instead.
func (*FindByIdRequest) Descriptor() ([]byte, []int) {
	return file_mongo_pet_proto_rawDescGZIP(), []int{7}
}

func (x *FindByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FindByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url      string               `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Created  *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated  *timestamp.Timestamp `protobuf:"bytes,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Status   string               `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Html     string               `protobuf:"bytes,6,opt,name=html,proto3" json:"html,omitempty"`
	PageRank float64              `protobuf:"fixed64,7,opt,name=page_rank,json=pageRank,proto3" json:"page_rank,omitempty"`
	Link2    []*Link              `protobuf:"bytes,8,rep,name=link2,proto3" json:"link2,omitempty"`
}

func (x *FindByIdResponse) Reset() {
	*x = FindByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mongo_pet_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByIdResponse) ProtoMessage() {}

func (x *FindByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mongo_pet_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByIdResponse.ProtoReflect.Descriptor instead.
func (*FindByIdResponse) Descriptor() ([]byte, []int) {
	return file_mongo_pet_proto_rawDescGZIP(), []int{8}
}

func (x *FindByIdResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FindByIdResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *FindByIdResponse) GetCreated() *timestamp.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *FindByIdResponse) GetUpdated() *timestamp.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

func (x *FindByIdResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *FindByIdResponse) GetHtml() string {
	if x != nil {
		return x.Html
	}
	return ""
}

func (x *FindByIdResponse) GetPageRank() float64 {
	if x != nil {
		return x.PageRank
	}
	return 0
}

func (x *FindByIdResponse) GetLink2() []*Link {
	if x != nil {
		return x.Link2
	}
	return nil
}

var File_mongo_pet_proto protoreflect.FileDescriptor

var file_mongo_pet_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x5f, 0x70, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x5f, 0x70, 0x65, 0x74, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a,
	0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x91, 0x02, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x34, 0x0a,
	0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x74, 0x6d, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x74, 0x6d, 0x6c, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x25, 0x0a, 0x05,
	0x6c, 0x69, 0x6e, 0x6b, 0x32, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f,
	0x6e, 0x67, 0x6f, 0x5f, 0x70, 0x65, 0x74, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x6c, 0x69,
	0x6e, 0x6b, 0x32, 0x22, 0x92, 0x02, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x34, 0x0a, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x74, 0x6d, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x68, 0x74, 0x6d, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x6e,
	0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x52, 0x61, 0x6e,
	0x6b, 0x12, 0x25, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x32, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x5f, 0x70, 0x65, 0x74, 0x2e, 0x4c, 0x69, 0x6e,
	0x6b, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x32, 0x22, 0x91, 0x02, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x74, 0x6d, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x74, 0x6d, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x25, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x32, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x5f, 0x70, 0x65, 0x74,
	0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x32, 0x22, 0x92, 0x02, 0x0a,
	0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x74, 0x6d,
	0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x74, 0x6d, 0x6c, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x25, 0x0a, 0x05, 0x6c, 0x69,
	0x6e, 0x6b, 0x32, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x6e, 0x67,
	0x6f, 0x5f, 0x70, 0x65, 0x74, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b,
	0x32, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x21, 0x0a, 0x0f,
	0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x90, 0x02, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x07,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x74,
	0x6d, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x74, 0x6d, 0x6c, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x25, 0x0a, 0x05, 0x6c,
	0x69, 0x6e, 0x6b, 0x32, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x6e,
	0x67, 0x6f, 0x5f, 0x70, 0x65, 0x74, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x6c, 0x69, 0x6e,
	0x6b, 0x32, 0x32, 0xb3, 0x02, 0x0a, 0x0b, 0x50, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65,
	0x12, 0x1c, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x5f, 0x70, 0x65, 0x74, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x5f, 0x70, 0x65, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a,
	0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x2e, 0x6d, 0x6f,
	0x6e, 0x67, 0x6f, 0x5f, 0x70, 0x65, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x6f, 0x6e, 0x67,
	0x6f, 0x5f, 0x70, 0x65, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x5f, 0x70,
	0x65, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x5f, 0x70, 0x65, 0x74,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x1a, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x5f, 0x70, 0x65, 0x74, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6d, 0x6f,
	0x6e, 0x67, 0x6f, 0x5f, 0x70, 0x65, 0x74, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mongo_pet_proto_rawDescOnce sync.Once
	file_mongo_pet_proto_rawDescData = file_mongo_pet_proto_rawDesc
)

func file_mongo_pet_proto_rawDescGZIP() []byte {
	file_mongo_pet_proto_rawDescOnce.Do(func() {
		file_mongo_pet_proto_rawDescData = protoimpl.X.CompressGZIP(file_mongo_pet_proto_rawDescData)
	})
	return file_mongo_pet_proto_rawDescData
}

var file_mongo_pet_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_mongo_pet_proto_goTypes = []interface{}{
	(*Link)(nil),                // 0: mongo_pet.Link
	(*CreatePageRequest)(nil),   // 1: mongo_pet.CreatePageRequest
	(*CreatePageResponse)(nil),  // 2: mongo_pet.CreatePageResponse
	(*UpdatePageRequest)(nil),   // 3: mongo_pet.UpdatePageRequest
	(*UpdatePageResponse)(nil),  // 4: mongo_pet.UpdatePageResponse
	(*DeletePageRequest)(nil),   // 5: mongo_pet.DeletePageRequest
	(*DeletePageResponse)(nil),  // 6: mongo_pet.DeletePageResponse
	(*FindByIdRequest)(nil),     // 7: mongo_pet.FindByIdRequest
	(*FindByIdResponse)(nil),    // 8: mongo_pet.FindByIdResponse
	(*timestamp.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_mongo_pet_proto_depIdxs = []int32{
	9,  // 0: mongo_pet.CreatePageRequest.created:type_name -> google.protobuf.Timestamp
	9,  // 1: mongo_pet.CreatePageRequest.updated:type_name -> google.protobuf.Timestamp
	0,  // 2: mongo_pet.CreatePageRequest.link2:type_name -> mongo_pet.Link
	9,  // 3: mongo_pet.CreatePageResponse.created:type_name -> google.protobuf.Timestamp
	9,  // 4: mongo_pet.CreatePageResponse.updated:type_name -> google.protobuf.Timestamp
	0,  // 5: mongo_pet.CreatePageResponse.link2:type_name -> mongo_pet.Link
	9,  // 6: mongo_pet.UpdatePageRequest.created:type_name -> google.protobuf.Timestamp
	9,  // 7: mongo_pet.UpdatePageRequest.updated:type_name -> google.protobuf.Timestamp
	0,  // 8: mongo_pet.UpdatePageRequest.link2:type_name -> mongo_pet.Link
	9,  // 9: mongo_pet.UpdatePageResponse.created:type_name -> google.protobuf.Timestamp
	9,  // 10: mongo_pet.UpdatePageResponse.updated:type_name -> google.protobuf.Timestamp
	0,  // 11: mongo_pet.UpdatePageResponse.link2:type_name -> mongo_pet.Link
	9,  // 12: mongo_pet.FindByIdResponse.created:type_name -> google.protobuf.Timestamp
	9,  // 13: mongo_pet.FindByIdResponse.updated:type_name -> google.protobuf.Timestamp
	0,  // 14: mongo_pet.FindByIdResponse.link2:type_name -> mongo_pet.Link
	1,  // 15: mongo_pet.PageService.CreatePage:input_type -> mongo_pet.CreatePageRequest
	3,  // 16: mongo_pet.PageService.UpdatePage:input_type -> mongo_pet.UpdatePageRequest
	5,  // 17: mongo_pet.PageService.DeletePage:input_type -> mongo_pet.DeletePageRequest
	7,  // 18: mongo_pet.PageService.FindById:input_type -> mongo_pet.FindByIdRequest
	2,  // 19: mongo_pet.PageService.CreatePage:output_type -> mongo_pet.CreatePageResponse
	4,  // 20: mongo_pet.PageService.UpdatePage:output_type -> mongo_pet.UpdatePageResponse
	6,  // 21: mongo_pet.PageService.DeletePage:output_type -> mongo_pet.DeletePageResponse
	8,  // 22: mongo_pet.PageService.FindById:output_type -> mongo_pet.FindByIdResponse
	19, // [19:23] is the sub-list for method output_type
	15, // [15:19] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_mongo_pet_proto_init() }
func file_mongo_pet_proto_init() {
	if File_mongo_pet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mongo_pet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Link); i {
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
		file_mongo_pet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePageRequest); i {
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
		file_mongo_pet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePageResponse); i {
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
		file_mongo_pet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePageRequest); i {
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
		file_mongo_pet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePageResponse); i {
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
		file_mongo_pet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePageRequest); i {
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
		file_mongo_pet_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePageResponse); i {
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
		file_mongo_pet_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByIdRequest); i {
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
		file_mongo_pet_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByIdResponse); i {
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
			RawDescriptor: file_mongo_pet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mongo_pet_proto_goTypes,
		DependencyIndexes: file_mongo_pet_proto_depIdxs,
		MessageInfos:      file_mongo_pet_proto_msgTypes,
	}.Build()
	File_mongo_pet_proto = out.File
	file_mongo_pet_proto_rawDesc = nil
	file_mongo_pet_proto_goTypes = nil
	file_mongo_pet_proto_depIdxs = nil
}
