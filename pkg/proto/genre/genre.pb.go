// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: pkg/proto/genre.proto

package genre

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

type Genre struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   *int64  `protobuf:"varint,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Name *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
}

func (x *Genre) Reset() {
	*x = Genre{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_genre_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Genre) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Genre) ProtoMessage() {}

func (x *Genre) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_genre_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Genre.ProtoReflect.Descriptor instead.
func (*Genre) Descriptor() ([]byte, []int) {
	return file_pkg_proto_genre_proto_rawDescGZIP(), []int{0}
}

func (x *Genre) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Genre) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type GetGenreByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetGenreByIdRequest) Reset() {
	*x = GetGenreByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_genre_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGenreByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGenreByIdRequest) ProtoMessage() {}

func (x *GetGenreByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_genre_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGenreByIdRequest.ProtoReflect.Descriptor instead.
func (*GetGenreByIdRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_genre_proto_rawDescGZIP(), []int{1}
}

func (x *GetGenreByIdRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetGenreByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Genre *Genre `protobuf:"bytes,2,opt,name=genre,proto3" json:"genre,omitempty"`
}

func (x *GetGenreByIdResponse) Reset() {
	*x = GetGenreByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_genre_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGenreByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGenreByIdResponse) ProtoMessage() {}

func (x *GetGenreByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_genre_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGenreByIdResponse.ProtoReflect.Descriptor instead.
func (*GetGenreByIdResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_genre_proto_rawDescGZIP(), []int{2}
}

func (x *GetGenreByIdResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *GetGenreByIdResponse) GetGenre() *Genre {
	if x != nil {
		return x.Genre
	}
	return nil
}

type GetAllGenresRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllGenresRequest) Reset() {
	*x = GetAllGenresRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_genre_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllGenresRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllGenresRequest) ProtoMessage() {}

func (x *GetAllGenresRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_genre_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllGenresRequest.ProtoReflect.Descriptor instead.
func (*GetAllGenresRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_genre_proto_rawDescGZIP(), []int{3}
}

type GetAllGenresResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Genres []*Genre `protobuf:"bytes,2,rep,name=genres,proto3" json:"genres,omitempty"`
}

func (x *GetAllGenresResponse) Reset() {
	*x = GetAllGenresResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_genre_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllGenresResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllGenresResponse) ProtoMessage() {}

func (x *GetAllGenresResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_genre_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllGenresResponse.ProtoReflect.Descriptor instead.
func (*GetAllGenresResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_genre_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllGenresResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *GetAllGenresResponse) GetGenres() []*Genre {
	if x != nil {
		return x.Genres
	}
	return nil
}

type AddGenderToMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MovieId int64 `protobuf:"varint,1,opt,name=movie_id,json=movieId,proto3" json:"movie_id,omitempty"`
	GenreId int64 `protobuf:"varint,2,opt,name=genre_id,json=genreId,proto3" json:"genre_id,omitempty"`
}

func (x *AddGenderToMovieRequest) Reset() {
	*x = AddGenderToMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_genre_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddGenderToMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddGenderToMovieRequest) ProtoMessage() {}

func (x *AddGenderToMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_genre_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddGenderToMovieRequest.ProtoReflect.Descriptor instead.
func (*AddGenderToMovieRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_genre_proto_rawDescGZIP(), []int{5}
}

func (x *AddGenderToMovieRequest) GetMovieId() int64 {
	if x != nil {
		return x.MovieId
	}
	return 0
}

func (x *AddGenderToMovieRequest) GetGenreId() int64 {
	if x != nil {
		return x.GenreId
	}
	return 0
}

type AddGenderToMovieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *AddGenderToMovieResponse) Reset() {
	*x = AddGenderToMovieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_genre_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddGenderToMovieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddGenderToMovieResponse) ProtoMessage() {}

func (x *AddGenderToMovieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_genre_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddGenderToMovieResponse.ProtoReflect.Descriptor instead.
func (*AddGenderToMovieResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_genre_proto_rawDescGZIP(), []int{6}
}

func (x *AddGenderToMovieResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AddGenderToMovieResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type RemoveGenderFromMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MovieId int64 `protobuf:"varint,1,opt,name=movie_id,json=movieId,proto3" json:"movie_id,omitempty"`
	GenreId int64 `protobuf:"varint,2,opt,name=genre_id,json=genreId,proto3" json:"genre_id,omitempty"`
}

func (x *RemoveGenderFromMovieRequest) Reset() {
	*x = RemoveGenderFromMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_genre_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveGenderFromMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveGenderFromMovieRequest) ProtoMessage() {}

func (x *RemoveGenderFromMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_genre_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveGenderFromMovieRequest.ProtoReflect.Descriptor instead.
func (*RemoveGenderFromMovieRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_genre_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveGenderFromMovieRequest) GetMovieId() int64 {
	if x != nil {
		return x.MovieId
	}
	return 0
}

func (x *RemoveGenderFromMovieRequest) GetGenreId() int64 {
	if x != nil {
		return x.GenreId
	}
	return 0
}

type RemoveGenderFromMovieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *RemoveGenderFromMovieResponse) Reset() {
	*x = RemoveGenderFromMovieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_genre_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveGenderFromMovieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveGenderFromMovieResponse) ProtoMessage() {}

func (x *RemoveGenderFromMovieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_genre_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveGenderFromMovieResponse.ProtoReflect.Descriptor instead.
func (*RemoveGenderFromMovieResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_genre_proto_rawDescGZIP(), []int{8}
}

func (x *RemoveGenderFromMovieResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *RemoveGenderFromMovieResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_pkg_proto_genre_proto protoreflect.FileDescriptor

var file_pkg_proto_genre_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x72,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x22, 0x45,
	0x0a, 0x05, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x47, 0x65, 0x6e, 0x72,
	0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x50, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x22, 0x0a, 0x05, 0x67, 0x65,
	0x6e, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x65, 0x6e, 0x72,
	0x65, 0x2e, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x52, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x22, 0x15,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x52, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x47,
	0x65, 0x6e, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x24, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x6e, 0x72,
	0x65, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x22, 0x4f, 0x0a, 0x17, 0x41, 0x64, 0x64,
	0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x6f, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x18, 0x41, 0x64,
	0x64, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x6f, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x54, 0x0a, 0x1c, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x47, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x1d, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xe1, 0x02, 0x0a, 0x0c, 0x47, 0x65,
	0x6e, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x2e, 0x67, 0x65, 0x6e,
	0x72, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x47,
	0x65, 0x6e, 0x72, 0x65, 0x73, 0x12, 0x1a, 0x2e, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x47, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x55, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x6f, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x12, 0x1e, 0x2e, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x2e, 0x41, 0x64, 0x64,
	0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x6f, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x2e, 0x41, 0x64, 0x64,
	0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x6f, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x12, 0x23, 0x2e, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x47,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x13, 0x5a,
	0x11, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e,
	0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_genre_proto_rawDescOnce sync.Once
	file_pkg_proto_genre_proto_rawDescData = file_pkg_proto_genre_proto_rawDesc
)

func file_pkg_proto_genre_proto_rawDescGZIP() []byte {
	file_pkg_proto_genre_proto_rawDescOnce.Do(func() {
		file_pkg_proto_genre_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_genre_proto_rawDescData)
	})
	return file_pkg_proto_genre_proto_rawDescData
}

var file_pkg_proto_genre_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pkg_proto_genre_proto_goTypes = []interface{}{
	(*Genre)(nil),                         // 0: genre.Genre
	(*GetGenreByIdRequest)(nil),           // 1: genre.GetGenreByIdRequest
	(*GetGenreByIdResponse)(nil),          // 2: genre.GetGenreByIdResponse
	(*GetAllGenresRequest)(nil),           // 3: genre.GetAllGenresRequest
	(*GetAllGenresResponse)(nil),          // 4: genre.GetAllGenresResponse
	(*AddGenderToMovieRequest)(nil),       // 5: genre.AddGenderToMovieRequest
	(*AddGenderToMovieResponse)(nil),      // 6: genre.AddGenderToMovieResponse
	(*RemoveGenderFromMovieRequest)(nil),  // 7: genre.RemoveGenderFromMovieRequest
	(*RemoveGenderFromMovieResponse)(nil), // 8: genre.RemoveGenderFromMovieResponse
}
var file_pkg_proto_genre_proto_depIdxs = []int32{
	0, // 0: genre.GetGenreByIdResponse.genre:type_name -> genre.Genre
	0, // 1: genre.GetAllGenresResponse.genres:type_name -> genre.Genre
	1, // 2: genre.GenreService.GetGenreById:input_type -> genre.GetGenreByIdRequest
	3, // 3: genre.GenreService.GetAllGenres:input_type -> genre.GetAllGenresRequest
	5, // 4: genre.GenreService.AddGenderToMovie:input_type -> genre.AddGenderToMovieRequest
	7, // 5: genre.GenreService.RemoveGenderFromMovie:input_type -> genre.RemoveGenderFromMovieRequest
	2, // 6: genre.GenreService.GetGenreById:output_type -> genre.GetGenreByIdResponse
	4, // 7: genre.GenreService.GetAllGenres:output_type -> genre.GetAllGenresResponse
	6, // 8: genre.GenreService.AddGenderToMovie:output_type -> genre.AddGenderToMovieResponse
	8, // 9: genre.GenreService.RemoveGenderFromMovie:output_type -> genre.RemoveGenderFromMovieResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_proto_genre_proto_init() }
func file_pkg_proto_genre_proto_init() {
	if File_pkg_proto_genre_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_genre_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Genre); i {
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
		file_pkg_proto_genre_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGenreByIdRequest); i {
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
		file_pkg_proto_genre_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGenreByIdResponse); i {
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
		file_pkg_proto_genre_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllGenresRequest); i {
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
		file_pkg_proto_genre_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllGenresResponse); i {
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
		file_pkg_proto_genre_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddGenderToMovieRequest); i {
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
		file_pkg_proto_genre_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddGenderToMovieResponse); i {
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
		file_pkg_proto_genre_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveGenderFromMovieRequest); i {
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
		file_pkg_proto_genre_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveGenderFromMovieResponse); i {
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
	file_pkg_proto_genre_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_proto_genre_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_genre_proto_goTypes,
		DependencyIndexes: file_pkg_proto_genre_proto_depIdxs,
		MessageInfos:      file_pkg_proto_genre_proto_msgTypes,
	}.Build()
	File_pkg_proto_genre_proto = out.File
	file_pkg_proto_genre_proto_rawDesc = nil
	file_pkg_proto_genre_proto_goTypes = nil
	file_pkg_proto_genre_proto_depIdxs = nil
}