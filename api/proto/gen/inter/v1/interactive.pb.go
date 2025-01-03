// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: inter/v1/interactive.proto

package interv1

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

// 定义请求消息
type IncrReadCntRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BizId         int64                  `protobuf:"varint,1,opt,name=biz_id,json=bizId,proto3" json:"biz_id,omitempty"` // 业务 ID
	Biz           string                 `protobuf:"bytes,2,opt,name=biz,proto3" json:"biz,omitempty"`                   // 业务类型
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IncrReadCntRequest) Reset() {
	*x = IncrReadCntRequest{}
	mi := &file_inter_v1_interactive_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IncrReadCntRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncrReadCntRequest) ProtoMessage() {}

func (x *IncrReadCntRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inter_v1_interactive_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncrReadCntRequest.ProtoReflect.Descriptor instead.
func (*IncrReadCntRequest) Descriptor() ([]byte, []int) {
	return file_inter_v1_interactive_proto_rawDescGZIP(), []int{0}
}

func (x *IncrReadCntRequest) GetBizId() int64 {
	if x != nil {
		return x.BizId
	}
	return 0
}

func (x *IncrReadCntRequest) GetBiz() string {
	if x != nil {
		return x.Biz
	}
	return ""
}

type LikeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户 ID
	BizId         int64                  `protobuf:"varint,2,opt,name=biz_id,json=bizId,proto3" json:"biz_id,omitempty"`    // 业务 ID
	Biz           string                 `protobuf:"bytes,3,opt,name=biz,proto3" json:"biz,omitempty"`                      // 业务类型
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LikeRequest) Reset() {
	*x = LikeRequest{}
	mi := &file_inter_v1_interactive_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LikeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeRequest) ProtoMessage() {}

func (x *LikeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inter_v1_interactive_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeRequest.ProtoReflect.Descriptor instead.
func (*LikeRequest) Descriptor() ([]byte, []int) {
	return file_inter_v1_interactive_proto_rawDescGZIP(), []int{1}
}

func (x *LikeRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *LikeRequest) GetBizId() int64 {
	if x != nil {
		return x.BizId
	}
	return 0
}

func (x *LikeRequest) GetBiz() string {
	if x != nil {
		return x.Biz
	}
	return ""
}

type CancelLikeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户 ID
	BizId         int64                  `protobuf:"varint,2,opt,name=biz_id,json=bizId,proto3" json:"biz_id,omitempty"`    // 业务 ID
	Biz           string                 `protobuf:"bytes,3,opt,name=biz,proto3" json:"biz,omitempty"`                      // 业务类型
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CancelLikeRequest) Reset() {
	*x = CancelLikeRequest{}
	mi := &file_inter_v1_interactive_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelLikeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelLikeRequest) ProtoMessage() {}

func (x *CancelLikeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inter_v1_interactive_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelLikeRequest.ProtoReflect.Descriptor instead.
func (*CancelLikeRequest) Descriptor() ([]byte, []int) {
	return file_inter_v1_interactive_proto_rawDescGZIP(), []int{2}
}

func (x *CancelLikeRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CancelLikeRequest) GetBizId() int64 {
	if x != nil {
		return x.BizId
	}
	return 0
}

func (x *CancelLikeRequest) GetBiz() string {
	if x != nil {
		return x.Biz
	}
	return ""
}

type CollectRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uid           int64                  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"` // 用户 ID
	Id            int64                  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`   // 业务 ID
	Cid           int64                  `protobuf:"varint,3,opt,name=cid,proto3" json:"cid,omitempty"` // 收藏夹 ID
	Biz           string                 `protobuf:"bytes,4,opt,name=biz,proto3" json:"biz,omitempty"`  // 业务类型
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CollectRequest) Reset() {
	*x = CollectRequest{}
	mi := &file_inter_v1_interactive_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CollectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectRequest) ProtoMessage() {}

func (x *CollectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inter_v1_interactive_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectRequest.ProtoReflect.Descriptor instead.
func (*CollectRequest) Descriptor() ([]byte, []int) {
	return file_inter_v1_interactive_proto_rawDescGZIP(), []int{3}
}

func (x *CollectRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *CollectRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CollectRequest) GetCid() int64 {
	if x != nil {
		return x.Cid
	}
	return 0
}

func (x *CollectRequest) GetBiz() string {
	if x != nil {
		return x.Biz
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uid           int64                  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"` // 用户 ID
	Id            int64                  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`   // 业务 ID
	Biz           string                 `protobuf:"bytes,3,opt,name=biz,proto3" json:"biz,omitempty"`  // 业务类型
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	mi := &file_inter_v1_interactive_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inter_v1_interactive_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_inter_v1_interactive_proto_rawDescGZIP(), []int{4}
}

func (x *GetRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *GetRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetRequest) GetBiz() string {
	if x != nil {
		return x.Biz
	}
	return ""
}

// 定义响应消息
type InteractiveResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ReadCnt       int32                  `protobuf:"varint,1,opt,name=read_cnt,json=readCnt,proto3" json:"read_cnt,omitempty"`          // 阅读次数
	LikeCnt       int32                  `protobuf:"varint,2,opt,name=like_cnt,json=likeCnt,proto3" json:"like_cnt,omitempty"`          // 点赞次数
	CollectCnt    int32                  `protobuf:"varint,3,opt,name=collect_cnt,json=collectCnt,proto3" json:"collect_cnt,omitempty"` // 收藏次数
	Liked         bool                   `protobuf:"varint,4,opt,name=liked,proto3" json:"liked,omitempty"`                             // 是否已点赞
	Collected     bool                   `protobuf:"varint,5,opt,name=collected,proto3" json:"collected,omitempty"`                     // 是否已收藏
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InteractiveResponse) Reset() {
	*x = InteractiveResponse{}
	mi := &file_inter_v1_interactive_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InteractiveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InteractiveResponse) ProtoMessage() {}

func (x *InteractiveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inter_v1_interactive_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InteractiveResponse.ProtoReflect.Descriptor instead.
func (*InteractiveResponse) Descriptor() ([]byte, []int) {
	return file_inter_v1_interactive_proto_rawDescGZIP(), []int{5}
}

func (x *InteractiveResponse) GetReadCnt() int32 {
	if x != nil {
		return x.ReadCnt
	}
	return 0
}

func (x *InteractiveResponse) GetLikeCnt() int32 {
	if x != nil {
		return x.LikeCnt
	}
	return 0
}

func (x *InteractiveResponse) GetCollectCnt() int32 {
	if x != nil {
		return x.CollectCnt
	}
	return 0
}

func (x *InteractiveResponse) GetLiked() bool {
	if x != nil {
		return x.Liked
	}
	return false
}

func (x *InteractiveResponse) GetCollected() bool {
	if x != nil {
		return x.Collected
	}
	return false
}

type EmptyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	mi := &file_inter_v1_interactive_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inter_v1_interactive_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_inter_v1_interactive_proto_rawDescGZIP(), []int{6}
}

var File_inter_v1_interactive_proto protoreflect.FileDescriptor

var file_inter_v1_interactive_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x3d, 0x0a, 0x12, 0x49, 0x6e, 0x63, 0x72, 0x52, 0x65,
	0x61, 0x64, 0x43, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06,
	0x62, 0x69, 0x7a, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x62, 0x69,
	0x7a, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x7a, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x62, 0x69, 0x7a, 0x22, 0x4f, 0x0a, 0x0b, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x15, 0x0a,
	0x06, 0x62, 0x69, 0x7a, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x62,
	0x69, 0x7a, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x62, 0x69, 0x7a, 0x22, 0x55, 0x0a, 0x11, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x69, 0x7a, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x62,
	0x69, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x69, 0x7a, 0x22, 0x56, 0x0a,
	0x0e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x63, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x7a, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x62, 0x69, 0x7a, 0x22, 0x40, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x7a, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x62, 0x69, 0x7a, 0x22, 0xa0, 0x01, 0x0a, 0x13, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x63, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x72, 0x65, 0x61, 0x64, 0x43, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x69,
	0x6b, 0x65, 0x5f, 0x63, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6c, 0x69,
	0x6b, 0x65, 0x43, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x5f, 0x63, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x43, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xd0, 0x02, 0x0a, 0x12,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x49, 0x6e, 0x63, 0x72, 0x52, 0x65, 0x61, 0x64, 0x43, 0x6e,
	0x74, 0x12, 0x1c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x63,
	0x72, 0x52, 0x65, 0x61, 0x64, 0x43, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x4c, 0x69, 0x6b, 0x65,
	0x12, 0x15, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6b, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x42, 0x0a, 0x0a, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x1b,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x12,
	0x18, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3a, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x9e,
	0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42,
	0x10, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x68, 0x75, 0x61, 0x6e, 0x67, 0x79, 0x75, 0x6c, 0x2f, 0x67, 0x6f, 0x2d, 0x62, 0x6c, 0x6f,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x49, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x08, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inter_v1_interactive_proto_rawDescOnce sync.Once
	file_inter_v1_interactive_proto_rawDescData = file_inter_v1_interactive_proto_rawDesc
)

func file_inter_v1_interactive_proto_rawDescGZIP() []byte {
	file_inter_v1_interactive_proto_rawDescOnce.Do(func() {
		file_inter_v1_interactive_proto_rawDescData = protoimpl.X.CompressGZIP(file_inter_v1_interactive_proto_rawDescData)
	})
	return file_inter_v1_interactive_proto_rawDescData
}

var file_inter_v1_interactive_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_inter_v1_interactive_proto_goTypes = []any{
	(*IncrReadCntRequest)(nil),  // 0: inter.v1.IncrReadCntRequest
	(*LikeRequest)(nil),         // 1: inter.v1.LikeRequest
	(*CancelLikeRequest)(nil),   // 2: inter.v1.CancelLikeRequest
	(*CollectRequest)(nil),      // 3: inter.v1.CollectRequest
	(*GetRequest)(nil),          // 4: inter.v1.GetRequest
	(*InteractiveResponse)(nil), // 5: inter.v1.InteractiveResponse
	(*EmptyResponse)(nil),       // 6: inter.v1.EmptyResponse
}
var file_inter_v1_interactive_proto_depIdxs = []int32{
	0, // 0: inter.v1.InteractiveService.IncrReadCnt:input_type -> inter.v1.IncrReadCntRequest
	1, // 1: inter.v1.InteractiveService.Like:input_type -> inter.v1.LikeRequest
	2, // 2: inter.v1.InteractiveService.CancelLike:input_type -> inter.v1.CancelLikeRequest
	3, // 3: inter.v1.InteractiveService.Collect:input_type -> inter.v1.CollectRequest
	4, // 4: inter.v1.InteractiveService.Get:input_type -> inter.v1.GetRequest
	6, // 5: inter.v1.InteractiveService.IncrReadCnt:output_type -> inter.v1.EmptyResponse
	6, // 6: inter.v1.InteractiveService.Like:output_type -> inter.v1.EmptyResponse
	6, // 7: inter.v1.InteractiveService.CancelLike:output_type -> inter.v1.EmptyResponse
	6, // 8: inter.v1.InteractiveService.Collect:output_type -> inter.v1.EmptyResponse
	5, // 9: inter.v1.InteractiveService.Get:output_type -> inter.v1.InteractiveResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_inter_v1_interactive_proto_init() }
func file_inter_v1_interactive_proto_init() {
	if File_inter_v1_interactive_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_inter_v1_interactive_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inter_v1_interactive_proto_goTypes,
		DependencyIndexes: file_inter_v1_interactive_proto_depIdxs,
		MessageInfos:      file_inter_v1_interactive_proto_msgTypes,
	}.Build()
	File_inter_v1_interactive_proto = out.File
	file_inter_v1_interactive_proto_rawDesc = nil
	file_inter_v1_interactive_proto_goTypes = nil
	file_inter_v1_interactive_proto_depIdxs = nil
}
