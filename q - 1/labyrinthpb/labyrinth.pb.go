// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: labyrinth.proto

package __

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
		mi := &file_labyrinth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_labyrinth_proto_msgTypes[0]
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
	return file_labyrinth_proto_rawDescGZIP(), []int{0}
}

type LabyrinthInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Width     int32    `protobuf:"varint,1,opt,name=Width,proto3" json:"Width,omitempty"`
	Height    int32    `protobuf:"varint,2,opt,name=Height,proto3" json:"Height,omitempty"`
	Labyrinth []string `protobuf:"bytes,3,rep,name=Labyrinth,proto3" json:"Labyrinth,omitempty"`
}

func (x *LabyrinthInfo) Reset() {
	*x = LabyrinthInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_labyrinth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabyrinthInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabyrinthInfo) ProtoMessage() {}

func (x *LabyrinthInfo) ProtoReflect() protoreflect.Message {
	mi := &file_labyrinth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabyrinthInfo.ProtoReflect.Descriptor instead.
func (*LabyrinthInfo) Descriptor() ([]byte, []int) {
	return file_labyrinth_proto_rawDescGZIP(), []int{1}
}

func (x *LabyrinthInfo) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *LabyrinthInfo) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *LabyrinthInfo) GetLabyrinth() []string {
	if x != nil {
		return x.Labyrinth
	}
	return nil
}

type PlayerStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Score           int32 `protobuf:"varint,1,opt,name=Score,proto3" json:"Score,omitempty"`
	Health          int32 `protobuf:"varint,2,opt,name=Health,proto3" json:"Health,omitempty"`
	X               int32 `protobuf:"varint,3,opt,name=X,proto3" json:"X,omitempty"`
	Y               int32 `protobuf:"varint,4,opt,name=Y,proto3" json:"Y,omitempty"`
	RemainingSpells int32 `protobuf:"varint,5,opt,name=RemainingSpells,proto3" json:"RemainingSpells,omitempty"`
}

func (x *PlayerStatus) Reset() {
	*x = PlayerStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_labyrinth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerStatus) ProtoMessage() {}

func (x *PlayerStatus) ProtoReflect() protoreflect.Message {
	mi := &file_labyrinth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerStatus.ProtoReflect.Descriptor instead.
func (*PlayerStatus) Descriptor() ([]byte, []int) {
	return file_labyrinth_proto_rawDescGZIP(), []int{2}
}

func (x *PlayerStatus) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *PlayerStatus) GetHealth() int32 {
	if x != nil {
		return x.Health
	}
	return 0
}

func (x *PlayerStatus) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *PlayerStatus) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *PlayerStatus) GetRemainingSpells() int32 {
	if x != nil {
		return x.RemainingSpells
	}
	return 0
}

type MoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Direction string `protobuf:"bytes,1,opt,name=Direction,proto3" json:"Direction,omitempty"` // Can be "up", "down", "left", "right"
}

func (x *MoveRequest) Reset() {
	*x = MoveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_labyrinth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveRequest) ProtoMessage() {}

func (x *MoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_labyrinth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveRequest.ProtoReflect.Descriptor instead.
func (*MoveRequest) Descriptor() ([]byte, []int) {
	return file_labyrinth_proto_rawDescGZIP(), []int{3}
}

func (x *MoveRequest) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

type MoveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"` // Can be "success", "failure", "victory", "death"
}

func (x *MoveResponse) Reset() {
	*x = MoveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_labyrinth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveResponse) ProtoMessage() {}

func (x *MoveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_labyrinth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveResponse.ProtoReflect.Descriptor instead.
func (*MoveResponse) Descriptor() ([]byte, []int) {
	return file_labyrinth_proto_rawDescGZIP(), []int{4}
}

func (x *MoveResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type RevelioRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X        int32  `protobuf:"varint,1,opt,name=X,proto3" json:"X,omitempty"`
	Y        int32  `protobuf:"varint,2,opt,name=Y,proto3" json:"Y,omitempty"`
	TileType string `protobuf:"bytes,3,opt,name=TileType,proto3" json:"TileType,omitempty"` // Can be "empty" or "coin" or "wall"
}

func (x *RevelioRequest) Reset() {
	*x = RevelioRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_labyrinth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevelioRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevelioRequest) ProtoMessage() {}

func (x *RevelioRequest) ProtoReflect() protoreflect.Message {
	mi := &file_labyrinth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevelioRequest.ProtoReflect.Descriptor instead.
func (*RevelioRequest) Descriptor() ([]byte, []int) {
	return file_labyrinth_proto_rawDescGZIP(), []int{5}
}

func (x *RevelioRequest) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *RevelioRequest) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *RevelioRequest) GetTileType() string {
	if x != nil {
		return x.TileType
	}
	return ""
}

type TilePosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=X,proto3" json:"X,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=Y,proto3" json:"Y,omitempty"`
}

func (x *TilePosition) Reset() {
	*x = TilePosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_labyrinth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TilePosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TilePosition) ProtoMessage() {}

func (x *TilePosition) ProtoReflect() protoreflect.Message {
	mi := &file_labyrinth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TilePosition.ProtoReflect.Descriptor instead.
func (*TilePosition) Descriptor() ([]byte, []int) {
	return file_labyrinth_proto_rawDescGZIP(), []int{6}
}

func (x *TilePosition) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *TilePosition) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

var File_labyrinth_proto protoreflect.FileDescriptor

var file_labyrinth_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6c, 0x61, 0x62, 0x79, 0x72, 0x69, 0x6e, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x6c, 0x61, 0x62, 0x79, 0x72, 0x69, 0x6e, 0x74, 0x68, 0x22, 0x07, 0x0a, 0x05,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x5b, 0x0a, 0x0d, 0x4c, 0x61, 0x62, 0x79, 0x72, 0x69, 0x6e,
	0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06,
	0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x48, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x61, 0x62, 0x79, 0x72, 0x69, 0x6e, 0x74,
	0x68, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x4c, 0x61, 0x62, 0x79, 0x72, 0x69, 0x6e,
	0x74, 0x68, 0x22, 0x82, 0x01, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x12, 0x0c, 0x0a, 0x01, 0x58, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x58, 0x12,
	0x0c, 0x0a, 0x01, 0x59, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x59, 0x12, 0x28, 0x0a,
	0x0f, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x65, 0x6c, 0x6c, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x53, 0x70, 0x65, 0x6c, 0x6c, 0x73, 0x22, 0x2b, 0x0a, 0x0b, 0x4d, 0x6f, 0x76, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x26, 0x0a, 0x0c, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x48, 0x0a, 0x0e,
	0x52, 0x65, 0x76, 0x65, 0x6c, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c,
	0x0a, 0x01, 0x58, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x58, 0x12, 0x0c, 0x0a, 0x01,
	0x59, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x59, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x69,
	0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x69,
	0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x2a, 0x0a, 0x0c, 0x54, 0x69, 0x6c, 0x65, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x58, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x01, 0x58, 0x12, 0x0c, 0x0a, 0x01, 0x59, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x01, 0x59, 0x32, 0xc4, 0x02, 0x0a, 0x09, 0x4c, 0x61, 0x62, 0x79, 0x72, 0x69, 0x6e, 0x74, 0x68,
	0x12, 0x3e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x62, 0x79, 0x72, 0x69, 0x6e, 0x74, 0x68,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x2e, 0x6c, 0x61, 0x62, 0x79, 0x72, 0x69, 0x6e, 0x74, 0x68,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x6c, 0x61, 0x62, 0x79, 0x72, 0x69, 0x6e,
	0x74, 0x68, 0x2e, 0x4c, 0x61, 0x62, 0x79, 0x72, 0x69, 0x6e, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x3c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x10, 0x2e, 0x6c, 0x61, 0x62, 0x79, 0x72, 0x69, 0x6e, 0x74, 0x68, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x6c, 0x61, 0x62, 0x79, 0x72, 0x69, 0x6e, 0x74,
	0x68, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3f,
	0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x6f, 0x76, 0x65, 0x12, 0x16,
	0x2e, 0x6c, 0x61, 0x62, 0x79, 0x72, 0x69, 0x6e, 0x74, 0x68, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6c, 0x61, 0x62, 0x79, 0x72, 0x69, 0x6e,
	0x74, 0x68, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3f, 0x0a, 0x07, 0x52, 0x65, 0x76, 0x65, 0x6c, 0x69, 0x6f, 0x12, 0x19, 0x2e, 0x6c, 0x61, 0x62,
	0x79, 0x72, 0x69, 0x6e, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x76, 0x65, 0x6c, 0x69, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6c, 0x61, 0x62, 0x79, 0x72, 0x69, 0x6e, 0x74,
	0x68, 0x2e, 0x54, 0x69, 0x6c, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x30, 0x01,
	0x12, 0x37, 0x0a, 0x08, 0x42, 0x6f, 0x6d, 0x62, 0x61, 0x72, 0x64, 0x61, 0x12, 0x17, 0x2e, 0x6c,
	0x61, 0x62, 0x79, 0x72, 0x69, 0x6e, 0x74, 0x68, 0x2e, 0x54, 0x69, 0x6c, 0x65, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x10, 0x2e, 0x6c, 0x61, 0x62, 0x79, 0x72, 0x69, 0x6e, 0x74,
	0x68, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x28, 0x01, 0x42, 0x03, 0x5a, 0x01, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_labyrinth_proto_rawDescOnce sync.Once
	file_labyrinth_proto_rawDescData = file_labyrinth_proto_rawDesc
)

func file_labyrinth_proto_rawDescGZIP() []byte {
	file_labyrinth_proto_rawDescOnce.Do(func() {
		file_labyrinth_proto_rawDescData = protoimpl.X.CompressGZIP(file_labyrinth_proto_rawDescData)
	})
	return file_labyrinth_proto_rawDescData
}

var file_labyrinth_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_labyrinth_proto_goTypes = []any{
	(*Empty)(nil),          // 0: labyrinth.Empty
	(*LabyrinthInfo)(nil),  // 1: labyrinth.LabyrinthInfo
	(*PlayerStatus)(nil),   // 2: labyrinth.PlayerStatus
	(*MoveRequest)(nil),    // 3: labyrinth.MoveRequest
	(*MoveResponse)(nil),   // 4: labyrinth.MoveResponse
	(*RevelioRequest)(nil), // 5: labyrinth.RevelioRequest
	(*TilePosition)(nil),   // 6: labyrinth.TilePosition
}
var file_labyrinth_proto_depIdxs = []int32{
	0, // 0: labyrinth.Labyrinth.GetLabyrinthInfo:input_type -> labyrinth.Empty
	0, // 1: labyrinth.Labyrinth.GetPlayerStatus:input_type -> labyrinth.Empty
	3, // 2: labyrinth.Labyrinth.RegisterMove:input_type -> labyrinth.MoveRequest
	5, // 3: labyrinth.Labyrinth.Revelio:input_type -> labyrinth.RevelioRequest
	6, // 4: labyrinth.Labyrinth.Bombarda:input_type -> labyrinth.TilePosition
	1, // 5: labyrinth.Labyrinth.GetLabyrinthInfo:output_type -> labyrinth.LabyrinthInfo
	2, // 6: labyrinth.Labyrinth.GetPlayerStatus:output_type -> labyrinth.PlayerStatus
	4, // 7: labyrinth.Labyrinth.RegisterMove:output_type -> labyrinth.MoveResponse
	6, // 8: labyrinth.Labyrinth.Revelio:output_type -> labyrinth.TilePosition
	0, // 9: labyrinth.Labyrinth.Bombarda:output_type -> labyrinth.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_labyrinth_proto_init() }
func file_labyrinth_proto_init() {
	if File_labyrinth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_labyrinth_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_labyrinth_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*LabyrinthInfo); i {
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
		file_labyrinth_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*PlayerStatus); i {
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
		file_labyrinth_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*MoveRequest); i {
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
		file_labyrinth_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*MoveResponse); i {
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
		file_labyrinth_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*RevelioRequest); i {
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
		file_labyrinth_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*TilePosition); i {
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
			RawDescriptor: file_labyrinth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_labyrinth_proto_goTypes,
		DependencyIndexes: file_labyrinth_proto_depIdxs,
		MessageInfos:      file_labyrinth_proto_msgTypes,
	}.Build()
	File_labyrinth_proto = out.File
	file_labyrinth_proto_rawDesc = nil
	file_labyrinth_proto_goTypes = nil
	file_labyrinth_proto_depIdxs = nil
}