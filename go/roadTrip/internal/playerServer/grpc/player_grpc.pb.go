// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package player_grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RoadTripPlayerClient is the client API for RoadTripPlayer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoadTripPlayerClient interface {
	CreateCharacter(ctx context.Context, in *CreateCharacterRequest, opts ...grpc.CallOption) (*Character, error)
	GetCharacter(ctx context.Context, in *GetCharacterRequest, opts ...grpc.CallOption) (*Character, error)
	UpdateCharacter(ctx context.Context, in *UpdateCharacterRequest, opts ...grpc.CallOption) (*Character, error)
	UpdateCar(ctx context.Context, in *UpdateCarRequest, opts ...grpc.CallOption) (*Car, error)
	GetTown(ctx context.Context, in *GetTownRequest, opts ...grpc.CallOption) (*Town, error)
}

type roadTripPlayerClient struct {
	cc grpc.ClientConnInterface
}

func NewRoadTripPlayerClient(cc grpc.ClientConnInterface) RoadTripPlayerClient {
	return &roadTripPlayerClient{cc}
}

func (c *roadTripPlayerClient) CreateCharacter(ctx context.Context, in *CreateCharacterRequest, opts ...grpc.CallOption) (*Character, error) {
	out := new(Character)
	err := c.cc.Invoke(ctx, "/roadtrip.RoadTripPlayer/CreateCharacter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roadTripPlayerClient) GetCharacter(ctx context.Context, in *GetCharacterRequest, opts ...grpc.CallOption) (*Character, error) {
	out := new(Character)
	err := c.cc.Invoke(ctx, "/roadtrip.RoadTripPlayer/GetCharacter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roadTripPlayerClient) UpdateCharacter(ctx context.Context, in *UpdateCharacterRequest, opts ...grpc.CallOption) (*Character, error) {
	out := new(Character)
	err := c.cc.Invoke(ctx, "/roadtrip.RoadTripPlayer/UpdateCharacter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roadTripPlayerClient) UpdateCar(ctx context.Context, in *UpdateCarRequest, opts ...grpc.CallOption) (*Car, error) {
	out := new(Car)
	err := c.cc.Invoke(ctx, "/roadtrip.RoadTripPlayer/UpdateCar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roadTripPlayerClient) GetTown(ctx context.Context, in *GetTownRequest, opts ...grpc.CallOption) (*Town, error) {
	out := new(Town)
	err := c.cc.Invoke(ctx, "/roadtrip.RoadTripPlayer/GetTown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoadTripPlayerServer is the server API for RoadTripPlayer service.
// All implementations must embed UnimplementedRoadTripPlayerServer
// for forward compatibility
type RoadTripPlayerServer interface {
	CreateCharacter(context.Context, *CreateCharacterRequest) (*Character, error)
	GetCharacter(context.Context, *GetCharacterRequest) (*Character, error)
	UpdateCharacter(context.Context, *UpdateCharacterRequest) (*Character, error)
	UpdateCar(context.Context, *UpdateCarRequest) (*Car, error)
	GetTown(context.Context, *GetTownRequest) (*Town, error)
	mustEmbedUnimplementedRoadTripPlayerServer()
}

// UnimplementedRoadTripPlayerServer must be embedded to have forward compatible implementations.
type UnimplementedRoadTripPlayerServer struct {
}

func (UnimplementedRoadTripPlayerServer) CreateCharacter(context.Context, *CreateCharacterRequest) (*Character, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCharacter not implemented")
}
func (UnimplementedRoadTripPlayerServer) GetCharacter(context.Context, *GetCharacterRequest) (*Character, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCharacter not implemented")
}
func (UnimplementedRoadTripPlayerServer) UpdateCharacter(context.Context, *UpdateCharacterRequest) (*Character, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCharacter not implemented")
}
func (UnimplementedRoadTripPlayerServer) UpdateCar(context.Context, *UpdateCarRequest) (*Car, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCar not implemented")
}
func (UnimplementedRoadTripPlayerServer) GetTown(context.Context, *GetTownRequest) (*Town, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTown not implemented")
}
func (UnimplementedRoadTripPlayerServer) mustEmbedUnimplementedRoadTripPlayerServer() {}

// UnsafeRoadTripPlayerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoadTripPlayerServer will
// result in compilation errors.
type UnsafeRoadTripPlayerServer interface {
	mustEmbedUnimplementedRoadTripPlayerServer()
}

func RegisterRoadTripPlayerServer(s grpc.ServiceRegistrar, srv RoadTripPlayerServer) {
	s.RegisterService(&RoadTripPlayer_ServiceDesc, srv)
}

func _RoadTripPlayer_CreateCharacter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCharacterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoadTripPlayerServer).CreateCharacter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roadtrip.RoadTripPlayer/CreateCharacter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoadTripPlayerServer).CreateCharacter(ctx, req.(*CreateCharacterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoadTripPlayer_GetCharacter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCharacterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoadTripPlayerServer).GetCharacter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roadtrip.RoadTripPlayer/GetCharacter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoadTripPlayerServer).GetCharacter(ctx, req.(*GetCharacterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoadTripPlayer_UpdateCharacter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCharacterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoadTripPlayerServer).UpdateCharacter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roadtrip.RoadTripPlayer/UpdateCharacter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoadTripPlayerServer).UpdateCharacter(ctx, req.(*UpdateCharacterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoadTripPlayer_UpdateCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoadTripPlayerServer).UpdateCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roadtrip.RoadTripPlayer/UpdateCar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoadTripPlayerServer).UpdateCar(ctx, req.(*UpdateCarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoadTripPlayer_GetTown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoadTripPlayerServer).GetTown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roadtrip.RoadTripPlayer/GetTown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoadTripPlayerServer).GetTown(ctx, req.(*GetTownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RoadTripPlayer_ServiceDesc is the grpc.ServiceDesc for RoadTripPlayer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoadTripPlayer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "roadtrip.RoadTripPlayer",
	HandlerType: (*RoadTripPlayerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCharacter",
			Handler:    _RoadTripPlayer_CreateCharacter_Handler,
		},
		{
			MethodName: "GetCharacter",
			Handler:    _RoadTripPlayer_GetCharacter_Handler,
		},
		{
			MethodName: "UpdateCharacter",
			Handler:    _RoadTripPlayer_UpdateCharacter_Handler,
		},
		{
			MethodName: "UpdateCar",
			Handler:    _RoadTripPlayer_UpdateCar_Handler,
		},
		{
			MethodName: "GetTown",
			Handler:    _RoadTripPlayer_GetTown_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/playerServer/grpc/player.proto",
}