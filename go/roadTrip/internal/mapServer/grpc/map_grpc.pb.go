// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package map_grpc

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

// RoadTripMapClient is the client API for RoadTripMap service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoadTripMapClient interface {
	GetTown(ctx context.Context, in *GetTownRequest, opts ...grpc.CallOption) (*Town, error)
	GetRoad(ctx context.Context, in *GetRoadRequest, opts ...grpc.CallOption) (*Road, error)
	ListStates(ctx context.Context, in *ListStatesRequest, opts ...grpc.CallOption) (*ListStatesResponse, error)
	ListTowns(ctx context.Context, in *ListTownsRequest, opts ...grpc.CallOption) (*ListTownsResponse, error)
	ListRoads(ctx context.Context, in *ListRoadsRequest, opts ...grpc.CallOption) (*ListRoadsResponse, error)
}

type roadTripMapClient struct {
	cc grpc.ClientConnInterface
}

func NewRoadTripMapClient(cc grpc.ClientConnInterface) RoadTripMapClient {
	return &roadTripMapClient{cc}
}

func (c *roadTripMapClient) GetTown(ctx context.Context, in *GetTownRequest, opts ...grpc.CallOption) (*Town, error) {
	out := new(Town)
	err := c.cc.Invoke(ctx, "/roadtrip_map.RoadTripMap/GetTown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roadTripMapClient) GetRoad(ctx context.Context, in *GetRoadRequest, opts ...grpc.CallOption) (*Road, error) {
	out := new(Road)
	err := c.cc.Invoke(ctx, "/roadtrip_map.RoadTripMap/GetRoad", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roadTripMapClient) ListStates(ctx context.Context, in *ListStatesRequest, opts ...grpc.CallOption) (*ListStatesResponse, error) {
	out := new(ListStatesResponse)
	err := c.cc.Invoke(ctx, "/roadtrip_map.RoadTripMap/ListStates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roadTripMapClient) ListTowns(ctx context.Context, in *ListTownsRequest, opts ...grpc.CallOption) (*ListTownsResponse, error) {
	out := new(ListTownsResponse)
	err := c.cc.Invoke(ctx, "/roadtrip_map.RoadTripMap/ListTowns", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roadTripMapClient) ListRoads(ctx context.Context, in *ListRoadsRequest, opts ...grpc.CallOption) (*ListRoadsResponse, error) {
	out := new(ListRoadsResponse)
	err := c.cc.Invoke(ctx, "/roadtrip_map.RoadTripMap/ListRoads", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoadTripMapServer is the server API for RoadTripMap service.
// All implementations must embed UnimplementedRoadTripMapServer
// for forward compatibility
type RoadTripMapServer interface {
	GetTown(context.Context, *GetTownRequest) (*Town, error)
	GetRoad(context.Context, *GetRoadRequest) (*Road, error)
	ListStates(context.Context, *ListStatesRequest) (*ListStatesResponse, error)
	ListTowns(context.Context, *ListTownsRequest) (*ListTownsResponse, error)
	ListRoads(context.Context, *ListRoadsRequest) (*ListRoadsResponse, error)
	mustEmbedUnimplementedRoadTripMapServer()
}

// UnimplementedRoadTripMapServer must be embedded to have forward compatible implementations.
type UnimplementedRoadTripMapServer struct {
}

func (UnimplementedRoadTripMapServer) GetTown(context.Context, *GetTownRequest) (*Town, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTown not implemented")
}
func (UnimplementedRoadTripMapServer) GetRoad(context.Context, *GetRoadRequest) (*Road, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoad not implemented")
}
func (UnimplementedRoadTripMapServer) ListStates(context.Context, *ListStatesRequest) (*ListStatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStates not implemented")
}
func (UnimplementedRoadTripMapServer) ListTowns(context.Context, *ListTownsRequest) (*ListTownsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTowns not implemented")
}
func (UnimplementedRoadTripMapServer) ListRoads(context.Context, *ListRoadsRequest) (*ListRoadsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRoads not implemented")
}
func (UnimplementedRoadTripMapServer) mustEmbedUnimplementedRoadTripMapServer() {}

// UnsafeRoadTripMapServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoadTripMapServer will
// result in compilation errors.
type UnsafeRoadTripMapServer interface {
	mustEmbedUnimplementedRoadTripMapServer()
}

func RegisterRoadTripMapServer(s grpc.ServiceRegistrar, srv RoadTripMapServer) {
	s.RegisterService(&RoadTripMap_ServiceDesc, srv)
}

func _RoadTripMap_GetTown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoadTripMapServer).GetTown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roadtrip_map.RoadTripMap/GetTown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoadTripMapServer).GetTown(ctx, req.(*GetTownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoadTripMap_GetRoad_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoadTripMapServer).GetRoad(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roadtrip_map.RoadTripMap/GetRoad",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoadTripMapServer).GetRoad(ctx, req.(*GetRoadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoadTripMap_ListStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoadTripMapServer).ListStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roadtrip_map.RoadTripMap/ListStates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoadTripMapServer).ListStates(ctx, req.(*ListStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoadTripMap_ListTowns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTownsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoadTripMapServer).ListTowns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roadtrip_map.RoadTripMap/ListTowns",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoadTripMapServer).ListTowns(ctx, req.(*ListTownsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoadTripMap_ListRoads_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRoadsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoadTripMapServer).ListRoads(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roadtrip_map.RoadTripMap/ListRoads",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoadTripMapServer).ListRoads(ctx, req.(*ListRoadsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RoadTripMap_ServiceDesc is the grpc.ServiceDesc for RoadTripMap service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoadTripMap_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "roadtrip_map.RoadTripMap",
	HandlerType: (*RoadTripMapServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTown",
			Handler:    _RoadTripMap_GetTown_Handler,
		},
		{
			MethodName: "GetRoad",
			Handler:    _RoadTripMap_GetRoad_Handler,
		},
		{
			MethodName: "ListStates",
			Handler:    _RoadTripMap_ListStates_Handler,
		},
		{
			MethodName: "ListTowns",
			Handler:    _RoadTripMap_ListTowns_Handler,
		},
		{
			MethodName: "ListRoads",
			Handler:    _RoadTripMap_ListRoads_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/mapServer/grpc/map.proto",
}
