// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.8
// source: characters.proto

package characterspb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateCharacterMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	UserId   uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BreedId  uint64 `protobuf:"varint,3,opt,name=breed_id,json=breedId,proto3" json:"breed_id,omitempty"`
	GenderId uint64 `protobuf:"varint,4,opt,name=gender_id,json=genderId,proto3" json:"gender_id,omitempty"`
	RealmId  uint64 `protobuf:"varint,5,opt,name=realm_id,json=realmId,proto3" json:"realm_id,omitempty"`
}

func (x *CreateCharacterMessage) Reset() {
	*x = CreateCharacterMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_characters_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCharacterMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCharacterMessage) ProtoMessage() {}

func (x *CreateCharacterMessage) ProtoReflect() protoreflect.Message {
	mi := &file_characters_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCharacterMessage.ProtoReflect.Descriptor instead.
func (*CreateCharacterMessage) Descriptor() ([]byte, []int) {
	return file_characters_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCharacterMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCharacterMessage) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateCharacterMessage) GetBreedId() uint64 {
	if x != nil {
		return x.BreedId
	}
	return 0
}

func (x *CreateCharacterMessage) GetGenderId() uint64 {
	if x != nil {
		return x.GenderId
	}
	return 0
}

func (x *CreateCharacterMessage) GetRealmId() uint64 {
	if x != nil {
		return x.RealmId
	}
	return 0
}

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_characters_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_characters_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_characters_proto_rawDescGZIP(), []int{1}
}

var File_characters_proto protoreflect.FileDescriptor

var file_characters_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x73, 0x72, 0x6f, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x73, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x72,
	0x65, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x62, 0x72,
	0x65, 0x65, 0x64, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x49, 0x64, 0x22, 0x0f, 0x0a,
	0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x81,
	0x01, 0x0a, 0x10, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x6d, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x28, 0x2e,
	0x73, 0x72, 0x6f, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x70, 0x62,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1f, 0x2e, 0x73, 0x72, 0x6f, 0x2e, 0x63, 0x68,
	0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12,
	0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x3a,
	0x01, 0x2a, 0x42, 0x12, 0x5a, 0x10, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63,
	0x74, 0x65, 0x72, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_characters_proto_rawDescOnce sync.Once
	file_characters_proto_rawDescData = file_characters_proto_rawDesc
)

func file_characters_proto_rawDescGZIP() []byte {
	file_characters_proto_rawDescOnce.Do(func() {
		file_characters_proto_rawDescData = protoimpl.X.CompressGZIP(file_characters_proto_rawDescData)
	})
	return file_characters_proto_rawDescData
}

var file_characters_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_characters_proto_goTypes = []interface{}{
	(*CreateCharacterMessage)(nil), // 0: sro.characterspb.CreateCharacterMessage
	(*EmptyResponse)(nil),          // 1: sro.characterspb.EmptyResponse
}
var file_characters_proto_depIdxs = []int32{
	0, // 0: sro.characterspb.CharacterService.Create:input_type -> sro.characterspb.CreateCharacterMessage
	1, // 1: sro.characterspb.CharacterService.Create:output_type -> sro.characterspb.EmptyResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_characters_proto_init() }
func file_characters_proto_init() {
	if File_characters_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_characters_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCharacterMessage); i {
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
		file_characters_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResponse); i {
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
			RawDescriptor: file_characters_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_characters_proto_goTypes,
		DependencyIndexes: file_characters_proto_depIdxs,
		MessageInfos:      file_characters_proto_msgTypes,
	}.Build()
	File_characters_proto = out.File
	file_characters_proto_rawDesc = nil
	file_characters_proto_goTypes = nil
	file_characters_proto_depIdxs = nil
}