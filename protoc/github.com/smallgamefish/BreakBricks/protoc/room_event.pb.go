// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: room_event.proto

package protoc

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

//创建房间事件
type CreateRoomEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId string `protobuf:"bytes,1,opt,name=roomId,proto3" json:"roomId,omitempty"` //创建房间的id
}

func (x *CreateRoomEvent) Reset() {
	*x = CreateRoomEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoomEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomEvent) ProtoMessage() {}

func (x *CreateRoomEvent) ProtoReflect() protoreflect.Message {
	mi := &file_room_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomEvent.ProtoReflect.Descriptor instead.
func (*CreateRoomEvent) Descriptor() ([]byte, []int) {
	return file_room_event_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRoomEvent) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

//加入房间事件
type JoinRoomEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId string `protobuf:"bytes,1,opt,name=roomId,proto3" json:"roomId,omitempty"` //加入房间的id
}

func (x *JoinRoomEvent) Reset() {
	*x = JoinRoomEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRoomEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRoomEvent) ProtoMessage() {}

func (x *JoinRoomEvent) ProtoReflect() protoreflect.Message {
	mi := &file_room_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRoomEvent.ProtoReflect.Descriptor instead.
func (*JoinRoomEvent) Descriptor() ([]byte, []int) {
	return file_room_event_proto_rawDescGZIP(), []int{1}
}

func (x *JoinRoomEvent) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

//离开房间事件
type LeaveRoomEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId string `protobuf:"bytes,1,opt,name=roomId,proto3" json:"roomId,omitempty"` //离开房间的id
}

func (x *LeaveRoomEvent) Reset() {
	*x = LeaveRoomEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveRoomEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveRoomEvent) ProtoMessage() {}

func (x *LeaveRoomEvent) ProtoReflect() protoreflect.Message {
	mi := &file_room_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveRoomEvent.ProtoReflect.Descriptor instead.
func (*LeaveRoomEvent) Descriptor() ([]byte, []int) {
	return file_room_event_proto_rawDescGZIP(), []int{2}
}

func (x *LeaveRoomEvent) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

//刷新房间的玩家事件
type RefreshRoomPlayerEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Players []*Player `protobuf:"bytes,1,rep,name=Players,proto3" json:"Players,omitempty"`
}

func (x *RefreshRoomPlayerEvent) Reset() {
	*x = RefreshRoomPlayerEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshRoomPlayerEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshRoomPlayerEvent) ProtoMessage() {}

func (x *RefreshRoomPlayerEvent) ProtoReflect() protoreflect.Message {
	mi := &file_room_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshRoomPlayerEvent.ProtoReflect.Descriptor instead.
func (*RefreshRoomPlayerEvent) Descriptor() ([]byte, []int) {
	return file_room_event_proto_rawDescGZIP(), []int{3}
}

func (x *RefreshRoomPlayerEvent) GetPlayers() []*Player {
	if x != nil {
		return x.Players
	}
	return nil
}

//准备事件
type ReadyEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId string `protobuf:"bytes,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	Ready  bool   `protobuf:"varint,2,opt,name=ready,proto3" json:"ready,omitempty"` //true准备，false取消准备
}

func (x *ReadyEvent) Reset() {
	*x = ReadyEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadyEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadyEvent) ProtoMessage() {}

func (x *ReadyEvent) ProtoReflect() protoreflect.Message {
	mi := &file_room_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadyEvent.ProtoReflect.Descriptor instead.
func (*ReadyEvent) Descriptor() ([]byte, []int) {
	return file_room_event_proto_rawDescGZIP(), []int{4}
}

func (x *ReadyEvent) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *ReadyEvent) GetReady() bool {
	if x != nil {
		return x.Ready
	}
	return false
}

//开始游戏事件
type StartGameEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` //展示没有什么用
}

func (x *StartGameEvent) Reset() {
	*x = StartGameEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_event_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartGameEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartGameEvent) ProtoMessage() {}

func (x *StartGameEvent) ProtoReflect() protoreflect.Message {
	mi := &file_room_event_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartGameEvent.ProtoReflect.Descriptor instead.
func (*StartGameEvent) Descriptor() ([]byte, []int) {
	return file_room_event_proto_rawDescGZIP(), []int{5}
}

func (x *StartGameEvent) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

//ping，心跳事件
type PingEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId string `protobuf:"bytes,1,opt,name=roomId,proto3" json:"roomId,omitempty"` //房间id
}

func (x *PingEvent) Reset() {
	*x = PingEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_event_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingEvent) ProtoMessage() {}

func (x *PingEvent) ProtoReflect() protoreflect.Message {
	mi := &file_room_event_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingEvent.ProtoReflect.Descriptor instead.
func (*PingEvent) Descriptor() ([]byte, []int) {
	return file_room_event_proto_rawDescGZIP(), []int{6}
}

func (x *PingEvent) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

//针数据
type FrameDataEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId string `protobuf:"bytes,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	//游戏的数据帧数据
	FrameData []byte `protobuf:"bytes,2,opt,name=frameData,proto3" json:"frameData,omitempty"`
}

func (x *FrameDataEvent) Reset() {
	*x = FrameDataEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_event_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrameDataEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrameDataEvent) ProtoMessage() {}

func (x *FrameDataEvent) ProtoReflect() protoreflect.Message {
	mi := &file_room_event_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrameDataEvent.ProtoReflect.Descriptor instead.
func (*FrameDataEvent) Descriptor() ([]byte, []int) {
	return file_room_event_proto_rawDescGZIP(), []int{7}
}

func (x *FrameDataEvent) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *FrameDataEvent) GetFrameData() []byte {
	if x != nil {
		return x.FrameData
	}
	return nil
}

var File_room_event_proto protoreflect.FileDescriptor

var file_room_event_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x5f, 0x62, 0x72, 0x69, 0x63, 0x6b, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x1a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x6f, 0x6f, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64,
	0x22, 0x27, 0x0a, 0x0d, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x0e, 0x4c, 0x65, 0x61,
	0x76, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f,
	0x6d, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x16, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x6f,
	0x6f, 0x6d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x35, 0x0a,
	0x07, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x5f, 0x62, 0x72, 0x69, 0x63, 0x6b, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x73, 0x22, 0x3a, 0x0a, 0x0a, 0x52, 0x65, 0x61, 0x64, 0x79, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65,
	0x61, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79,
	0x22, 0x24, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x23, 0x0a, 0x09, 0x50, 0x69, 0x6e, 0x67, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x22, 0x46, 0x0a, 0x0e, 0x46,
	0x72, 0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x67, 0x61, 0x6d, 0x65, 0x66, 0x69, 0x73, 0x68, 0x2f,
	0x42, 0x72, 0x65, 0x61, 0x6b, 0x42, 0x72, 0x69, 0x63, 0x6b, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_room_event_proto_rawDescOnce sync.Once
	file_room_event_proto_rawDescData = file_room_event_proto_rawDesc
)

func file_room_event_proto_rawDescGZIP() []byte {
	file_room_event_proto_rawDescOnce.Do(func() {
		file_room_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_room_event_proto_rawDescData)
	})
	return file_room_event_proto_rawDescData
}

var file_room_event_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_room_event_proto_goTypes = []interface{}{
	(*CreateRoomEvent)(nil),        // 0: break_bricks.protoc.CreateRoomEvent
	(*JoinRoomEvent)(nil),          // 1: break_bricks.protoc.JoinRoomEvent
	(*LeaveRoomEvent)(nil),         // 2: break_bricks.protoc.LeaveRoomEvent
	(*RefreshRoomPlayerEvent)(nil), // 3: break_bricks.protoc.RefreshRoomPlayerEvent
	(*ReadyEvent)(nil),             // 4: break_bricks.protoc.ReadyEvent
	(*StartGameEvent)(nil),         // 5: break_bricks.protoc.StartGameEvent
	(*PingEvent)(nil),              // 6: break_bricks.protoc.PingEvent
	(*FrameDataEvent)(nil),         // 7: break_bricks.protoc.FrameDataEvent
	(*Player)(nil),                 // 8: break_bricks.protoc.Player
}
var file_room_event_proto_depIdxs = []int32{
	8, // 0: break_bricks.protoc.RefreshRoomPlayerEvent.Players:type_name -> break_bricks.protoc.Player
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_room_event_proto_init() }
func file_room_event_proto_init() {
	if File_room_event_proto != nil {
		return
	}
	file_player_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_room_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoomEvent); i {
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
		file_room_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRoomEvent); i {
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
		file_room_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveRoomEvent); i {
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
		file_room_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshRoomPlayerEvent); i {
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
		file_room_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadyEvent); i {
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
		file_room_event_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartGameEvent); i {
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
		file_room_event_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingEvent); i {
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
		file_room_event_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrameDataEvent); i {
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
			RawDescriptor: file_room_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_room_event_proto_goTypes,
		DependencyIndexes: file_room_event_proto_depIdxs,
		MessageInfos:      file_room_event_proto_msgTypes,
	}.Build()
	File_room_event_proto = out.File
	file_room_event_proto_rawDesc = nil
	file_room_event_proto_goTypes = nil
	file_room_event_proto_depIdxs = nil
}
