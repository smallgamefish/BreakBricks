// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: client.proto

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

type ClientAcceptMsg_Code int32

const (
	ClientAcceptMsg_Success ClientAcceptMsg_Code = 0
	ClientAcceptMsg_Error   ClientAcceptMsg_Code = 1
)

// Enum value maps for ClientAcceptMsg_Code.
var (
	ClientAcceptMsg_Code_name = map[int32]string{
		0: "Success",
		1: "Error",
	}
	ClientAcceptMsg_Code_value = map[string]int32{
		"Success": 0,
		"Error":   1,
	}
)

func (x ClientAcceptMsg_Code) Enum() *ClientAcceptMsg_Code {
	p := new(ClientAcceptMsg_Code)
	*p = x
	return p
}

func (x ClientAcceptMsg_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClientAcceptMsg_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_client_proto_enumTypes[0].Descriptor()
}

func (ClientAcceptMsg_Code) Type() protoreflect.EnumType {
	return &file_client_proto_enumTypes[0]
}

func (x ClientAcceptMsg_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClientAcceptMsg_Code.Descriptor instead.
func (ClientAcceptMsg_Code) EnumDescriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{1, 0}
}

//客户端发送udp的msg消息格式
type ClientSendMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Event:
	//	*ClientSendMsg_CreateRoomEvent
	Event isClientSendMsg_Event `protobuf_oneof:"Event"`
}

func (x *ClientSendMsg) Reset() {
	*x = ClientSendMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientSendMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientSendMsg) ProtoMessage() {}

func (x *ClientSendMsg) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientSendMsg.ProtoReflect.Descriptor instead.
func (*ClientSendMsg) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{0}
}

func (m *ClientSendMsg) GetEvent() isClientSendMsg_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *ClientSendMsg) GetCreateRoomEvent() *CreateRoomEvent {
	if x, ok := x.GetEvent().(*ClientSendMsg_CreateRoomEvent); ok {
		return x.CreateRoomEvent
	}
	return nil
}

type isClientSendMsg_Event interface {
	isClientSendMsg_Event()
}

type ClientSendMsg_CreateRoomEvent struct {
	CreateRoomEvent *CreateRoomEvent `protobuf:"bytes,1,opt,name=createRoomEvent,proto3,oneof"`
}

func (*ClientSendMsg_CreateRoomEvent) isClientSendMsg_Event() {}

//客户端接收udp的msg消息格式
type ClientAcceptMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code ClientAcceptMsg_Code `protobuf:"varint,1,opt,name=code,proto3,enum=break_bricks.protoc.ClientAcceptMsg_Code" json:"code,omitempty"`
	//code等于error时候的错误信息
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	// Types that are assignable to Event:
	//	*ClientAcceptMsg_CreateRoomEvent
	Event isClientAcceptMsg_Event `protobuf_oneof:"Event"`
}

func (x *ClientAcceptMsg) Reset() {
	*x = ClientAcceptMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientAcceptMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientAcceptMsg) ProtoMessage() {}

func (x *ClientAcceptMsg) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientAcceptMsg.ProtoReflect.Descriptor instead.
func (*ClientAcceptMsg) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{1}
}

func (x *ClientAcceptMsg) GetCode() ClientAcceptMsg_Code {
	if x != nil {
		return x.Code
	}
	return ClientAcceptMsg_Success
}

func (x *ClientAcceptMsg) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (m *ClientAcceptMsg) GetEvent() isClientAcceptMsg_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *ClientAcceptMsg) GetCreateRoomEvent() *CreateRoomEvent {
	if x, ok := x.GetEvent().(*ClientAcceptMsg_CreateRoomEvent); ok {
		return x.CreateRoomEvent
	}
	return nil
}

type isClientAcceptMsg_Event interface {
	isClientAcceptMsg_Event()
}

type ClientAcceptMsg_CreateRoomEvent struct {
	CreateRoomEvent *CreateRoomEvent `protobuf:"bytes,3,opt,name=createRoomEvent,proto3,oneof"`
}

func (*ClientAcceptMsg_CreateRoomEvent) isClientAcceptMsg_Event() {}

var File_client_proto protoreflect.FileDescriptor

var file_client_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x62, 0x72, 0x65, 0x61, 0x6b, 0x5f, 0x62, 0x72, 0x69, 0x63, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x1a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x6a, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73,
	0x67, 0x12, 0x50, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x62, 0x72, 0x65,
	0x61, 0x6b, 0x5f, 0x62, 0x72, 0x69, 0x63, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x48, 0x00, 0x52, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0xe1, 0x01, 0x0a,
	0x0f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x4d, 0x73, 0x67,
	0x12, 0x3d, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29,
	0x2e, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x5f, 0x62, 0x72, 0x69, 0x63, 0x6b, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x4d, 0x73, 0x67, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x50, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x6f, 0x6f, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x5f, 0x62, 0x72, 0x69, 0x63, 0x6b, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x6f, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x1e, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6d, 0x61, 0x6c, 0x6c, 0x67, 0x61, 0x6d, 0x65, 0x66, 0x69, 0x73, 0x68, 0x2f, 0x42, 0x72, 0x65,
	0x61, 0x6b, 0x42, 0x72, 0x69, 0x63, 0x6b, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_proto_rawDescOnce sync.Once
	file_client_proto_rawDescData = file_client_proto_rawDesc
)

func file_client_proto_rawDescGZIP() []byte {
	file_client_proto_rawDescOnce.Do(func() {
		file_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_proto_rawDescData)
	})
	return file_client_proto_rawDescData
}

var file_client_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_client_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_client_proto_goTypes = []interface{}{
	(ClientAcceptMsg_Code)(0), // 0: break_bricks.protoc.ClientAcceptMsg.Code
	(*ClientSendMsg)(nil),     // 1: break_bricks.protoc.ClientSendMsg
	(*ClientAcceptMsg)(nil),   // 2: break_bricks.protoc.ClientAcceptMsg
	(*CreateRoomEvent)(nil),   // 3: break_bricks.protoc.CreateRoomEvent
}
var file_client_proto_depIdxs = []int32{
	3, // 0: break_bricks.protoc.ClientSendMsg.createRoomEvent:type_name -> break_bricks.protoc.CreateRoomEvent
	0, // 1: break_bricks.protoc.ClientAcceptMsg.code:type_name -> break_bricks.protoc.ClientAcceptMsg.Code
	3, // 2: break_bricks.protoc.ClientAcceptMsg.createRoomEvent:type_name -> break_bricks.protoc.CreateRoomEvent
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_client_proto_init() }
func file_client_proto_init() {
	if File_client_proto != nil {
		return
	}
	file_event_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientSendMsg); i {
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
		file_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientAcceptMsg); i {
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
	file_client_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ClientSendMsg_CreateRoomEvent)(nil),
	}
	file_client_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ClientAcceptMsg_CreateRoomEvent)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_client_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_client_proto_goTypes,
		DependencyIndexes: file_client_proto_depIdxs,
		EnumInfos:         file_client_proto_enumTypes,
		MessageInfos:      file_client_proto_msgTypes,
	}.Build()
	File_client_proto = out.File
	file_client_proto_rawDesc = nil
	file_client_proto_goTypes = nil
	file_client_proto_depIdxs = nil
}
