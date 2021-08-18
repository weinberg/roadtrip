// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: internal/grpc/player.proto

package grpc

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
		mi := &file_internal_grpc_player_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_player_proto_msgTypes[0]
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
	return file_internal_grpc_player_proto_rawDescGZIP(), []int{0}
}

type Character struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // output only
	CharacterName string `protobuf:"bytes,2,opt,name=character_name,json=characterName,proto3" json:"character_name,omitempty"`
	Car           *Car   `protobuf:"bytes,3,opt,name=car,proto3" json:"car,omitempty"` // singleton
}

func (x *Character) Reset() {
	*x = Character{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_player_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Character) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Character) ProtoMessage() {}

func (x *Character) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_player_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Character.ProtoReflect.Descriptor instead.
func (*Character) Descriptor() ([]byte, []int) {
	return file_internal_grpc_player_proto_rawDescGZIP(), []int{1}
}

func (x *Character) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Character) GetCharacterName() string {
	if x != nil {
		return x.CharacterName
	}
	return ""
}

func (x *Character) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

type GetCharacterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCharacterRequest) Reset() {
	*x = GetCharacterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_player_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCharacterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCharacterRequest) ProtoMessage() {}

func (x *GetCharacterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_player_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCharacterRequest.ProtoReflect.Descriptor instead.
func (*GetCharacterRequest) Descriptor() ([]byte, []int) {
	return file_internal_grpc_player_proto_rawDescGZIP(), []int{2}
}

func (x *GetCharacterRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateCharacterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CaptchaId     string `protobuf:"bytes,1,opt,name=captcha_id,json=captchaId,proto3" json:"captcha_id,omitempty"`
	CharacterName string `protobuf:"bytes,2,opt,name=character_name,json=characterName,proto3" json:"character_name,omitempty"`
}

func (x *CreateCharacterRequest) Reset() {
	*x = CreateCharacterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_player_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCharacterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCharacterRequest) ProtoMessage() {}

func (x *CreateCharacterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_player_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCharacterRequest.ProtoReflect.Descriptor instead.
func (*CreateCharacterRequest) Descriptor() ([]byte, []int) {
	return file_internal_grpc_player_proto_rawDescGZIP(), []int{3}
}

func (x *CreateCharacterRequest) GetCaptchaId() string {
	if x != nil {
		return x.CaptchaId
	}
	return ""
}

func (x *CreateCharacterRequest) GetCharacterName() string {
	if x != nil {
		return x.CharacterName
	}
	return ""
}

type UpdateCharacterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Character *Character `protobuf:"bytes,1,opt,name=character,proto3" json:"character,omitempty"`
}

func (x *UpdateCharacterRequest) Reset() {
	*x = UpdateCharacterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_player_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCharacterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCharacterRequest) ProtoMessage() {}

func (x *UpdateCharacterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_player_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCharacterRequest.ProtoReflect.Descriptor instead.
func (*UpdateCharacterRequest) Descriptor() ([]byte, []int) {
	return file_internal_grpc_player_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCharacterRequest) GetCharacter() *Character {
	if x != nil {
		return x.Character
	}
	return nil
}

type UpdateCarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Car *Car `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"`
}

func (x *UpdateCarRequest) Reset() {
	*x = UpdateCarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_player_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCarRequest) ProtoMessage() {}

func (x *UpdateCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_player_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCarRequest.ProtoReflect.Descriptor instead.
func (*UpdateCarRequest) Descriptor() ([]byte, []int) {
	return file_internal_grpc_player_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateCarRequest) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

type Car struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // output only
	Plate   string `protobuf:"bytes,2,opt,name=plate,proto3" json:"plate,omitempty"`
	CarName string `protobuf:"bytes,3,opt,name=car_name,json=carName,proto3" json:"car_name,omitempty"`
}

func (x *Car) Reset() {
	*x = Car{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_player_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Car) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Car) ProtoMessage() {}

func (x *Car) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_player_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Car.ProtoReflect.Descriptor instead.
func (*Car) Descriptor() ([]byte, []int) {
	return file_internal_grpc_player_proto_rawDescGZIP(), []int{6}
}

func (x *Car) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Car) GetPlate() string {
	if x != nil {
		return x.Plate
	}
	return ""
}

func (x *Car) GetCarName() string {
	if x != nil {
		return x.CarName
	}
	return ""
}

var File_internal_grpc_player_proto protoreflect.FileDescriptor

var file_internal_grpc_player_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x72, 0x6f,
	0x61, 0x64, 0x74, 0x72, 0x69, 0x70, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x63, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e,
	0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x03, 0x63, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x72, 0x6f, 0x61, 0x64, 0x74, 0x72, 0x69, 0x70, 0x2e, 0x43, 0x61, 0x72, 0x52,
	0x03, 0x63, 0x61, 0x72, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5e, 0x0a, 0x16, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x70, 0x74, 0x63,
	0x68, 0x61, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x68,
	0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4b, 0x0a, 0x16, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x6f, 0x61, 0x64, 0x74,
	0x72, 0x69, 0x70, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x09, 0x63,
	0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x22, 0x33, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x03,
	0x63, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x6f, 0x61, 0x64,
	0x74, 0x72, 0x69, 0x70, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x03, 0x63, 0x61, 0x72, 0x22, 0x46, 0x0a,
	0x03, 0x43, 0x61, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x61,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0xa8, 0x02, 0x0a, 0x0e, 0x52, 0x6f, 0x61, 0x64, 0x54, 0x72,
	0x69, 0x70, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x72, 0x6f,
	0x61, 0x64, 0x74, 0x72, 0x69, 0x70, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x72, 0x6f, 0x61, 0x64, 0x74, 0x72, 0x69, 0x70, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74,
	0x65, 0x72, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x72, 0x6f, 0x61, 0x64, 0x74, 0x72, 0x69, 0x70, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x72, 0x6f, 0x61, 0x64, 0x74, 0x72, 0x69, 0x70, 0x2e, 0x43,
	0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0f, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x20, 0x2e,
	0x72, 0x6f, 0x61, 0x64, 0x74, 0x72, 0x69, 0x70, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x72, 0x6f, 0x61, 0x64, 0x74, 0x72, 0x69, 0x70, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x72, 0x12, 0x1a, 0x2e, 0x72, 0x6f, 0x61, 0x64, 0x74, 0x72, 0x69, 0x70, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0d, 0x2e, 0x72, 0x6f, 0x61, 0x64, 0x74, 0x72, 0x69, 0x70, 0x2e, 0x43, 0x61, 0x72, 0x22, 0x00,
	0x42, 0x1b, 0x5a, 0x19, 0x69, 0x6e, 0x73, 0x6f, 0x66, 0x61, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x72, 0x6f, 0x61, 0x64, 0x54, 0x72, 0x69, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_grpc_player_proto_rawDescOnce sync.Once
	file_internal_grpc_player_proto_rawDescData = file_internal_grpc_player_proto_rawDesc
)

func file_internal_grpc_player_proto_rawDescGZIP() []byte {
	file_internal_grpc_player_proto_rawDescOnce.Do(func() {
		file_internal_grpc_player_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_grpc_player_proto_rawDescData)
	})
	return file_internal_grpc_player_proto_rawDescData
}

var file_internal_grpc_player_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_internal_grpc_player_proto_goTypes = []interface{}{
	(*Empty)(nil),                  // 0: roadtrip.Empty
	(*Character)(nil),              // 1: roadtrip.Character
	(*GetCharacterRequest)(nil),    // 2: roadtrip.GetCharacterRequest
	(*CreateCharacterRequest)(nil), // 3: roadtrip.CreateCharacterRequest
	(*UpdateCharacterRequest)(nil), // 4: roadtrip.UpdateCharacterRequest
	(*UpdateCarRequest)(nil),       // 5: roadtrip.UpdateCarRequest
	(*Car)(nil),                    // 6: roadtrip.Car
}
var file_internal_grpc_player_proto_depIdxs = []int32{
	6, // 0: roadtrip.Character.car:type_name -> roadtrip.Car
	1, // 1: roadtrip.UpdateCharacterRequest.character:type_name -> roadtrip.Character
	6, // 2: roadtrip.UpdateCarRequest.car:type_name -> roadtrip.Car
	3, // 3: roadtrip.RoadTripPlayer.CreateCharacter:input_type -> roadtrip.CreateCharacterRequest
	2, // 4: roadtrip.RoadTripPlayer.GetCharacter:input_type -> roadtrip.GetCharacterRequest
	4, // 5: roadtrip.RoadTripPlayer.UpdateCharacter:input_type -> roadtrip.UpdateCharacterRequest
	5, // 6: roadtrip.RoadTripPlayer.UpdateCar:input_type -> roadtrip.UpdateCarRequest
	1, // 7: roadtrip.RoadTripPlayer.CreateCharacter:output_type -> roadtrip.Character
	1, // 8: roadtrip.RoadTripPlayer.GetCharacter:output_type -> roadtrip.Character
	1, // 9: roadtrip.RoadTripPlayer.UpdateCharacter:output_type -> roadtrip.Character
	6, // 10: roadtrip.RoadTripPlayer.UpdateCar:output_type -> roadtrip.Car
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_internal_grpc_player_proto_init() }
func file_internal_grpc_player_proto_init() {
	if File_internal_grpc_player_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_grpc_player_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_internal_grpc_player_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Character); i {
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
		file_internal_grpc_player_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCharacterRequest); i {
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
		file_internal_grpc_player_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCharacterRequest); i {
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
		file_internal_grpc_player_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCharacterRequest); i {
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
		file_internal_grpc_player_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCarRequest); i {
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
		file_internal_grpc_player_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Car); i {
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
			RawDescriptor: file_internal_grpc_player_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_grpc_player_proto_goTypes,
		DependencyIndexes: file_internal_grpc_player_proto_depIdxs,
		MessageInfos:      file_internal_grpc_player_proto_msgTypes,
	}.Build()
	File_internal_grpc_player_proto = out.File
	file_internal_grpc_player_proto_rawDesc = nil
	file_internal_grpc_player_proto_goTypes = nil
	file_internal_grpc_player_proto_depIdxs = nil
}
