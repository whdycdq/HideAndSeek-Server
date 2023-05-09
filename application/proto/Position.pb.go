// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.0
// source: Position.proto

package hasproto

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

type SyncPlayerToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *SyncPlayerToken) Reset() {
	*x = SyncPlayerToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Position_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncPlayerToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncPlayerToken) ProtoMessage() {}

func (x *SyncPlayerToken) ProtoReflect() protoreflect.Message {
	mi := &file_Position_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncPlayerToken.ProtoReflect.Descriptor instead.
func (*SyncPlayerToken) Descriptor() ([]byte, []int) {
	return file_Position_proto_rawDescGZIP(), []int{0}
}

func (x *SyncPlayerToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type SyncPlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token         *SyncPlayerToken `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	Pid           int32            `protobuf:"varint,2,opt,name=Pid,proto3" json:"Pid,omitempty"`
	Name          string           `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	CharacterType int32            `protobuf:"varint,4,opt,name=CharacterType,proto3" json:"CharacterType,omitempty"`
}

func (x *SyncPlayerInfo) Reset() {
	*x = SyncPlayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Position_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncPlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncPlayerInfo) ProtoMessage() {}

func (x *SyncPlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_Position_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncPlayerInfo.ProtoReflect.Descriptor instead.
func (*SyncPlayerInfo) Descriptor() ([]byte, []int) {
	return file_Position_proto_rawDescGZIP(), []int{1}
}

func (x *SyncPlayerInfo) GetToken() *SyncPlayerToken {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *SyncPlayerInfo) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *SyncPlayerInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SyncPlayerInfo) GetCharacterType() int32 {
	if x != nil {
		return x.CharacterType
	}
	return 0
}

type PlayerChat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type int32   `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Pid  []int32 `protobuf:"varint,2,rep,packed,name=pid,proto3" json:"pid,omitempty"`
	Chat string  `protobuf:"bytes,3,opt,name=Chat,proto3" json:"Chat,omitempty"`
}

func (x *PlayerChat) Reset() {
	*x = PlayerChat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Position_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerChat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerChat) ProtoMessage() {}

func (x *PlayerChat) ProtoReflect() protoreflect.Message {
	mi := &file_Position_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerChat.ProtoReflect.Descriptor instead.
func (*PlayerChat) Descriptor() ([]byte, []int) {
	return file_Position_proto_rawDescGZIP(), []int{2}
}

func (x *PlayerChat) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *PlayerChat) GetPid() []int32 {
	if x != nil {
		return x.Pid
	}
	return nil
}

func (x *PlayerChat) GetChat() string {
	if x != nil {
		return x.Chat
	}
	return ""
}

type RayIntersect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Origin    *Location `protobuf:"bytes,1,opt,name=Origin,proto3" json:"Origin,omitempty"`
	Direction *Location `protobuf:"bytes,2,opt,name=Direction,proto3" json:"Direction,omitempty"`
	Actor     string    `protobuf:"bytes,3,opt,name=Actor,proto3" json:"Actor,omitempty"`
}

func (x *RayIntersect) Reset() {
	*x = RayIntersect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Position_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RayIntersect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RayIntersect) ProtoMessage() {}

func (x *RayIntersect) ProtoReflect() protoreflect.Message {
	mi := &file_Position_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RayIntersect.ProtoReflect.Descriptor instead.
func (*RayIntersect) Descriptor() ([]byte, []int) {
	return file_Position_proto_rawDescGZIP(), []int{3}
}

func (x *RayIntersect) GetOrigin() *Location {
	if x != nil {
		return x.Origin
	}
	return nil
}

func (x *RayIntersect) GetDirection() *Location {
	if x != nil {
		return x.Direction
	}
	return nil
}

func (x *RayIntersect) GetActor() string {
	if x != nil {
		return x.Actor
	}
	return ""
}

// 玩家位置
type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float32 `protobuf:"fixed32,1,opt,name=X,proto3" json:"X,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=Y,proto3" json:"Y,omitempty"`
	Z float32 `protobuf:"fixed32,3,opt,name=Z,proto3" json:"Z,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Position_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_Position_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_Position_proto_rawDescGZIP(), []int{4}
}

func (x *Location) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Location) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Location) GetZ() float32 {
	if x != nil {
		return x.Z
	}
	return 0
}

type Rotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Roll  float32 `protobuf:"fixed32,1,opt,name=Roll,proto3" json:"Roll,omitempty"`
	Pitch float32 `protobuf:"fixed32,2,opt,name=Pitch,proto3" json:"Pitch,omitempty"`
	Yaw   float32 `protobuf:"fixed32,3,opt,name=Yaw,proto3" json:"Yaw,omitempty"`
}

func (x *Rotation) Reset() {
	*x = Rotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Position_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rotation) ProtoMessage() {}

func (x *Rotation) ProtoReflect() protoreflect.Message {
	mi := &file_Position_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rotation.ProtoReflect.Descriptor instead.
func (*Rotation) Descriptor() ([]byte, []int) {
	return file_Position_proto_rawDescGZIP(), []int{5}
}

func (x *Rotation) GetRoll() float32 {
	if x != nil {
		return x.Roll
	}
	return 0
}

func (x *Rotation) GetPitch() float32 {
	if x != nil {
		return x.Pitch
	}
	return 0
}

func (x *Rotation) GetYaw() float32 {
	if x != nil {
		return x.Yaw
	}
	return 0
}

type Scale struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float32 `protobuf:"fixed32,1,opt,name=X,proto3" json:"X,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=Y,proto3" json:"Y,omitempty"`
	Z float32 `protobuf:"fixed32,3,opt,name=Z,proto3" json:"Z,omitempty"`
}

func (x *Scale) Reset() {
	*x = Scale{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Position_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scale) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scale) ProtoMessage() {}

func (x *Scale) ProtoReflect() protoreflect.Message {
	mi := &file_Position_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scale.ProtoReflect.Descriptor instead.
func (*Scale) Descriptor() ([]byte, []int) {
	return file_Position_proto_rawDescGZIP(), []int{6}
}

func (x *Scale) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Scale) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Scale) GetZ() float32 {
	if x != nil {
		return x.Z
	}
	return 0
}

type ControllerRotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Roll  float32 `protobuf:"fixed32,1,opt,name=Roll,proto3" json:"Roll,omitempty"`
	Pitch float32 `protobuf:"fixed32,2,opt,name=Pitch,proto3" json:"Pitch,omitempty"`
	Yaw   float32 `protobuf:"fixed32,3,opt,name=Yaw,proto3" json:"Yaw,omitempty"`
}

func (x *ControllerRotation) Reset() {
	*x = ControllerRotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Position_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControllerRotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControllerRotation) ProtoMessage() {}

func (x *ControllerRotation) ProtoReflect() protoreflect.Message {
	mi := &file_Position_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControllerRotation.ProtoReflect.Descriptor instead.
func (*ControllerRotation) Descriptor() ([]byte, []int) {
	return file_Position_proto_rawDescGZIP(), []int{7}
}

func (x *ControllerRotation) GetRoll() float32 {
	if x != nil {
		return x.Roll
	}
	return 0
}

func (x *ControllerRotation) GetPitch() float32 {
	if x != nil {
		return x.Pitch
	}
	return 0
}

func (x *ControllerRotation) GetYaw() float32 {
	if x != nil {
		return x.Yaw
	}
	return 0
}

type Transform struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	L *Location `protobuf:"bytes,1,opt,name=L,proto3" json:"L,omitempty"`
	R *Rotation `protobuf:"bytes,2,opt,name=R,proto3" json:"R,omitempty"`
	S *Scale    `protobuf:"bytes,3,opt,name=S,proto3" json:"S,omitempty"`
}

func (x *Transform) Reset() {
	*x = Transform{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Position_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transform) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transform) ProtoMessage() {}

func (x *Transform) ProtoReflect() protoreflect.Message {
	mi := &file_Position_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transform.ProtoReflect.Descriptor instead.
func (*Transform) Descriptor() ([]byte, []int) {
	return file_Position_proto_rawDescGZIP(), []int{8}
}

func (x *Transform) GetL() *Location {
	if x != nil {
		return x.L
	}
	return nil
}

func (x *Transform) GetR() *Rotation {
	if x != nil {
		return x.R
	}
	return nil
}

func (x *Transform) GetS() *Scale {
	if x != nil {
		return x.S
	}
	return nil
}

type Axis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MoveForward float32 `protobuf:"fixed32,1,opt,name=MoveForward,proto3" json:"MoveForward,omitempty"`
	MoveRight   float32 `protobuf:"fixed32,2,opt,name=MoveRight,proto3" json:"MoveRight,omitempty"`
}

func (x *Axis) Reset() {
	*x = Axis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Position_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Axis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Axis) ProtoMessage() {}

func (x *Axis) ProtoReflect() protoreflect.Message {
	mi := &file_Position_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Axis.ProtoReflect.Descriptor instead.
func (*Axis) Descriptor() ([]byte, []int) {
	return file_Position_proto_rawDescGZIP(), []int{9}
}

func (x *Axis) GetMoveForward() float32 {
	if x != nil {
		return x.MoveForward
	}
	return 0
}

func (x *Axis) GetMoveRight() float32 {
	if x != nil {
		return x.MoveRight
	}
	return 0
}

type PlayerMovementInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InputAxis *Axis               `protobuf:"bytes,1,opt,name=InputAxis,proto3" json:"InputAxis,omitempty"`
	Trans     *Transform          `protobuf:"bytes,2,opt,name=Trans,proto3" json:"Trans,omitempty"`
	CRotation *ControllerRotation `protobuf:"bytes,3,opt,name=CRotation,proto3" json:"CRotation,omitempty"`
}

func (x *PlayerMovementInfo) Reset() {
	*x = PlayerMovementInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Position_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerMovementInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerMovementInfo) ProtoMessage() {}

func (x *PlayerMovementInfo) ProtoReflect() protoreflect.Message {
	mi := &file_Position_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerMovementInfo.ProtoReflect.Descriptor instead.
func (*PlayerMovementInfo) Descriptor() ([]byte, []int) {
	return file_Position_proto_rawDescGZIP(), []int{10}
}

func (x *PlayerMovementInfo) GetInputAxis() *Axis {
	if x != nil {
		return x.InputAxis
	}
	return nil
}

func (x *PlayerMovementInfo) GetTrans() *Transform {
	if x != nil {
		return x.Trans
	}
	return nil
}

func (x *PlayerMovementInfo) GetCRotation() *ControllerRotation {
	if x != nil {
		return x.CRotation
	}
	return nil
}

// 玩家广播数据
type BroadCast struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      *SyncPlayerInfo     `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	Content    int32               `protobuf:"varint,2,opt,name=Content,proto3" json:"Content,omitempty"`
	Info       *PlayerMovementInfo `protobuf:"bytes,3,opt,name=Info,proto3" json:"Info,omitempty"`
	ActionData int32               `protobuf:"varint,4,opt,name=ActionData,proto3" json:"ActionData,omitempty"`
}

func (x *BroadCast) Reset() {
	*x = BroadCast{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Position_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadCast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadCast) ProtoMessage() {}

func (x *BroadCast) ProtoReflect() protoreflect.Message {
	mi := &file_Position_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadCast.ProtoReflect.Descriptor instead.
func (*BroadCast) Descriptor() ([]byte, []int) {
	return file_Position_proto_rawDescGZIP(), []int{11}
}

func (x *BroadCast) GetToken() *SyncPlayerInfo {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *BroadCast) GetContent() int32 {
	if x != nil {
		return x.Content
	}
	return 0
}

func (x *BroadCast) GetInfo() *PlayerMovementInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *BroadCast) GetActionData() int32 {
	if x != nil {
		return x.ActionData
	}
	return 0
}

var File_Position_proto protoreflect.FileDescriptor

var file_Position_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x27, 0x0a, 0x0f, 0x53, 0x79, 0x6e, 0x63,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x8b, 0x01, 0x0a, 0x0e, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2d, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x53, 0x79, 0x6e,
	0x63, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x05, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x50, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x68, 0x61,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x46, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03,
	0x70, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x43, 0x68, 0x61, 0x74, 0x22, 0x7e, 0x0a, 0x0c, 0x52, 0x61, 0x79, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x73, 0x65, 0x63, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x4f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x4f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x12, 0x2e, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x34, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x58, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01,
	0x58, 0x12, 0x0c, 0x0a, 0x01, 0x59, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x59, 0x12,
	0x0c, 0x0a, 0x01, 0x5a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x5a, 0x22, 0x46, 0x0a,
	0x08, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x6f, 0x6c,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x52, 0x6f, 0x6c, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x50, 0x69, 0x74, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x50, 0x69,
	0x74, 0x63, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x59, 0x61, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x03, 0x59, 0x61, 0x77, 0x22, 0x31, 0x0a, 0x05, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x0c,
	0x0a, 0x01, 0x58, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x58, 0x12, 0x0c, 0x0a, 0x01,
	0x59, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x59, 0x12, 0x0c, 0x0a, 0x01, 0x5a, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x5a, 0x22, 0x50, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x52, 0x6f,
	0x6c, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x69, 0x74, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x50, 0x69, 0x74, 0x63, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x59, 0x61, 0x77, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x59, 0x61, 0x77, 0x22, 0x68, 0x0a, 0x09, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x1e, 0x0a, 0x01, 0x4c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x01, 0x4c, 0x12, 0x1e, 0x0a, 0x01, 0x52, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x01, 0x52, 0x12, 0x1b, 0x0a, 0x01, 0x53, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x53, 0x63, 0x61, 0x6c,
	0x65, 0x52, 0x01, 0x53, 0x22, 0x46, 0x0a, 0x04, 0x41, 0x78, 0x69, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x4d, 0x6f, 0x76, 0x65, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0b, 0x4d, 0x6f, 0x76, 0x65, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x09, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x69, 0x67, 0x68, 0x74, 0x22, 0xa3, 0x01, 0x0a,
	0x12, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4d, 0x6f, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x2a, 0x0a, 0x09, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x41, 0x78, 0x69, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e,
	0x41, 0x78, 0x69, 0x73, 0x52, 0x09, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x41, 0x78, 0x69, 0x73, 0x12,
	0x27, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x52, 0x05, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x52, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x43, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0xa3, 0x01, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74,
	0x12, 0x2c, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4d, 0x6f, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x3b, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Position_proto_rawDescOnce sync.Once
	file_Position_proto_rawDescData = file_Position_proto_rawDesc
)

func file_Position_proto_rawDescGZIP() []byte {
	file_Position_proto_rawDescOnce.Do(func() {
		file_Position_proto_rawDescData = protoimpl.X.CompressGZIP(file_Position_proto_rawDescData)
	})
	return file_Position_proto_rawDescData
}

var file_Position_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_Position_proto_goTypes = []interface{}{
	(*SyncPlayerToken)(nil),    // 0: Player.SyncPlayerToken
	(*SyncPlayerInfo)(nil),     // 1: Player.SyncPlayerInfo
	(*PlayerChat)(nil),         // 2: Player.PlayerChat
	(*RayIntersect)(nil),       // 3: Player.RayIntersect
	(*Location)(nil),           // 4: Player.Location
	(*Rotation)(nil),           // 5: Player.Rotation
	(*Scale)(nil),              // 6: Player.Scale
	(*ControllerRotation)(nil), // 7: Player.ControllerRotation
	(*Transform)(nil),          // 8: Player.Transform
	(*Axis)(nil),               // 9: Player.Axis
	(*PlayerMovementInfo)(nil), // 10: Player.PlayerMovementInfo
	(*BroadCast)(nil),          // 11: Player.BroadCast
}
var file_Position_proto_depIdxs = []int32{
	0,  // 0: Player.SyncPlayerInfo.Token:type_name -> Player.SyncPlayerToken
	4,  // 1: Player.RayIntersect.Origin:type_name -> Player.Location
	4,  // 2: Player.RayIntersect.Direction:type_name -> Player.Location
	4,  // 3: Player.Transform.L:type_name -> Player.Location
	5,  // 4: Player.Transform.R:type_name -> Player.Rotation
	6,  // 5: Player.Transform.S:type_name -> Player.Scale
	9,  // 6: Player.PlayerMovementInfo.InputAxis:type_name -> Player.Axis
	8,  // 7: Player.PlayerMovementInfo.Trans:type_name -> Player.Transform
	7,  // 8: Player.PlayerMovementInfo.CRotation:type_name -> Player.ControllerRotation
	1,  // 9: Player.BroadCast.Token:type_name -> Player.SyncPlayerInfo
	10, // 10: Player.BroadCast.Info:type_name -> Player.PlayerMovementInfo
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_Position_proto_init() }
func file_Position_proto_init() {
	if File_Position_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Position_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncPlayerToken); i {
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
		file_Position_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncPlayerInfo); i {
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
		file_Position_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerChat); i {
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
		file_Position_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RayIntersect); i {
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
		file_Position_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_Position_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rotation); i {
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
		file_Position_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scale); i {
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
		file_Position_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControllerRotation); i {
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
		file_Position_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transform); i {
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
		file_Position_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Axis); i {
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
		file_Position_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerMovementInfo); i {
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
		file_Position_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadCast); i {
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
			RawDescriptor: file_Position_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Position_proto_goTypes,
		DependencyIndexes: file_Position_proto_depIdxs,
		MessageInfos:      file_Position_proto_msgTypes,
	}.Build()
	File_Position_proto = out.File
	file_Position_proto_rawDesc = nil
	file_Position_proto_goTypes = nil
	file_Position_proto_depIdxs = nil
}