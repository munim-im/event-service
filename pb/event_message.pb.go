// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: proto/event_message.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EventInputParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email         string           `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Environment   string           `protobuf:"bytes,2,opt,name=environment,proto3" json:"environment,omitempty"`
	Component     string           `protobuf:"bytes,3,opt,name=component,proto3" json:"component,omitempty"`
	MessageString string           `protobuf:"bytes,4,opt,name=message_string,json=messageString,proto3" json:"message_string,omitempty"`
	Data          *structpb.Struct `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *EventInputParams) Reset() {
	*x = EventInputParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventInputParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventInputParams) ProtoMessage() {}

func (x *EventInputParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventInputParams.ProtoReflect.Descriptor instead.
func (*EventInputParams) Descriptor() ([]byte, []int) {
	return file_proto_event_message_proto_rawDescGZIP(), []int{0}
}

func (x *EventInputParams) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *EventInputParams) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *EventInputParams) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *EventInputParams) GetMessageString() string {
	if x != nil {
		return x.MessageString
	}
	return ""
}

func (x *EventInputParams) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

type EventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email         string           `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Environment   string           `protobuf:"bytes,3,opt,name=environment,proto3" json:"environment,omitempty"`
	Component     string           `protobuf:"bytes,4,opt,name=component,proto3" json:"component,omitempty"`
	MessageString string           `protobuf:"bytes,5,opt,name=message_string,json=messageString,proto3" json:"message_string,omitempty"`
	Data          *structpb.Struct `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *EventResponse) Reset() {
	*x = EventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventResponse) ProtoMessage() {}

func (x *EventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventResponse.ProtoReflect.Descriptor instead.
func (*EventResponse) Descriptor() ([]byte, []int) {
	return file_proto_event_message_proto_rawDescGZIP(), []int{1}
}

func (x *EventResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EventResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *EventResponse) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *EventResponse) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *EventResponse) GetMessageString() string {
	if x != nil {
		return x.MessageString
	}
	return ""
}

func (x *EventResponse) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

type EventFilterInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params map[string]string `protobuf:"bytes,1,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *EventFilterInput) Reset() {
	*x = EventFilterInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventFilterInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventFilterInput) ProtoMessage() {}

func (x *EventFilterInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventFilterInput.ProtoReflect.Descriptor instead.
func (*EventFilterInput) Descriptor() ([]byte, []int) {
	return file_proto_event_message_proto_rawDescGZIP(), []int{2}
}

func (x *EventFilterInput) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

type EventFilterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*EventResponse `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *EventFilterResponse) Reset() {
	*x = EventFilterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventFilterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventFilterResponse) ProtoMessage() {}

func (x *EventFilterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventFilterResponse.ProtoReflect.Descriptor instead.
func (*EventFilterResponse) Descriptor() ([]byte, []int) {
	return file_proto_event_message_proto_rawDescGZIP(), []int{3}
}

func (x *EventFilterResponse) GetEvents() []*EventResponse {
	if x != nil {
		return x.Events
	}
	return nil
}

var File_proto_event_message_proto protoreflect.FileDescriptor

var file_proto_event_message_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x01, 0x0a, 0x10, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xc9, 0x01, 0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12,
	0x25, 0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x92, 0x01, 0x0a, 0x10, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x43, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x39, 0x0a,
	0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4b, 0x0a, 0x13, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x34, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_event_message_proto_rawDescOnce sync.Once
	file_proto_event_message_proto_rawDescData = file_proto_event_message_proto_rawDesc
)

func file_proto_event_message_proto_rawDescGZIP() []byte {
	file_proto_event_message_proto_rawDescOnce.Do(func() {
		file_proto_event_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_event_message_proto_rawDescData)
	})
	return file_proto_event_message_proto_rawDescData
}

var file_proto_event_message_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_event_message_proto_goTypes = []interface{}{
	(*EventInputParams)(nil),    // 0: event_service.EventInputParams
	(*EventResponse)(nil),       // 1: event_service.EventResponse
	(*EventFilterInput)(nil),    // 2: event_service.EventFilterInput
	(*EventFilterResponse)(nil), // 3: event_service.EventFilterResponse
	nil,                         // 4: event_service.EventFilterInput.ParamsEntry
	(*structpb.Struct)(nil),     // 5: google.protobuf.Struct
}
var file_proto_event_message_proto_depIdxs = []int32{
	5, // 0: event_service.EventInputParams.data:type_name -> google.protobuf.Struct
	5, // 1: event_service.EventResponse.data:type_name -> google.protobuf.Struct
	4, // 2: event_service.EventFilterInput.params:type_name -> event_service.EventFilterInput.ParamsEntry
	1, // 3: event_service.EventFilterResponse.events:type_name -> event_service.EventResponse
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_event_message_proto_init() }
func file_proto_event_message_proto_init() {
	if File_proto_event_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_event_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventInputParams); i {
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
		file_proto_event_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventResponse); i {
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
		file_proto_event_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventFilterInput); i {
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
		file_proto_event_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventFilterResponse); i {
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
			RawDescriptor: file_proto_event_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_event_message_proto_goTypes,
		DependencyIndexes: file_proto_event_message_proto_depIdxs,
		MessageInfos:      file_proto_event_message_proto_msgTypes,
	}.Build()
	File_proto_event_message_proto = out.File
	file_proto_event_message_proto_rawDesc = nil
	file_proto_event_message_proto_goTypes = nil
	file_proto_event_message_proto_depIdxs = nil
}
