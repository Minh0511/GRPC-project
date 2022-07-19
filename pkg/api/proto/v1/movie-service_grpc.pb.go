// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.11.4
// source: api/proto/v1/movie-service.proto

package v1

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

// MoviesServiceClient is the client API for MoviesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MoviesServiceClient interface {
	CreateMovies(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	GetAllMovies(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error)
	GetMovieByGenre(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	UpdateMovies(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	DeleteMovies(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type moviesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMoviesServiceClient(cc grpc.ClientConnInterface) MoviesServiceClient {
	return &moviesServiceClient{cc}
}

func (c *moviesServiceClient) CreateMovies(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/v1.MoviesService/CreateMovies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesServiceClient) GetAllMovies(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error) {
	out := new(ReadAllResponse)
	err := c.cc.Invoke(ctx, "/v1.MoviesService/GetAllMovies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesServiceClient) GetMovieByGenre(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/v1.MoviesService/GetMovieByGenre", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesServiceClient) UpdateMovies(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/v1.MoviesService/UpdateMovies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesServiceClient) DeleteMovies(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/v1.MoviesService/DeleteMovies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MoviesServiceServer is the server API for MoviesService service.
// All implementations must embed UnimplementedMoviesServiceServer
// for forward compatibility
type MoviesServiceServer interface {
	CreateMovies(context.Context, *CreateRequest) (*CreateResponse, error)
	GetAllMovies(context.Context, *ReadAllRequest) (*ReadAllResponse, error)
	GetMovieByGenre(context.Context, *ReadRequest) (*ReadResponse, error)
	UpdateMovies(context.Context, *UpdateRequest) (*UpdateResponse, error)
	DeleteMovies(context.Context, *DeleteRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedMoviesServiceServer()
}

// UnimplementedMoviesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMoviesServiceServer struct {
}

func (UnimplementedMoviesServiceServer) CreateMovies(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMovies not implemented")
}
func (UnimplementedMoviesServiceServer) GetAllMovies(context.Context, *ReadAllRequest) (*ReadAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllMovies not implemented")
}
func (UnimplementedMoviesServiceServer) GetMovieByGenre(context.Context, *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovieByGenre not implemented")
}
func (UnimplementedMoviesServiceServer) UpdateMovies(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMovies not implemented")
}
func (UnimplementedMoviesServiceServer) DeleteMovies(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMovies not implemented")
}
func (UnimplementedMoviesServiceServer) mustEmbedUnimplementedMoviesServiceServer() {}

// UnsafeMoviesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MoviesServiceServer will
// result in compilation errors.
type UnsafeMoviesServiceServer interface {
	mustEmbedUnimplementedMoviesServiceServer()
}

func RegisterMoviesServiceServer(s grpc.ServiceRegistrar, srv MoviesServiceServer) {
	s.RegisterService(&MoviesService_ServiceDesc, srv)
}

func _MoviesService_CreateMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesServiceServer).CreateMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.MoviesService/CreateMovies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesServiceServer).CreateMovies(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MoviesService_GetAllMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesServiceServer).GetAllMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.MoviesService/GetAllMovies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesServiceServer).GetAllMovies(ctx, req.(*ReadAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MoviesService_GetMovieByGenre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesServiceServer).GetMovieByGenre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.MoviesService/GetMovieByGenre",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesServiceServer).GetMovieByGenre(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MoviesService_UpdateMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesServiceServer).UpdateMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.MoviesService/UpdateMovies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesServiceServer).UpdateMovies(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MoviesService_DeleteMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesServiceServer).DeleteMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.MoviesService/DeleteMovies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesServiceServer).DeleteMovies(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MoviesService_ServiceDesc is the grpc.ServiceDesc for MoviesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MoviesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.MoviesService",
	HandlerType: (*MoviesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMovies",
			Handler:    _MoviesService_CreateMovies_Handler,
		},
		{
			MethodName: "GetAllMovies",
			Handler:    _MoviesService_GetAllMovies_Handler,
		},
		{
			MethodName: "GetMovieByGenre",
			Handler:    _MoviesService_GetMovieByGenre_Handler,
		},
		{
			MethodName: "UpdateMovies",
			Handler:    _MoviesService_UpdateMovies_Handler,
		},
		{
			MethodName: "DeleteMovies",
			Handler:    _MoviesService_DeleteMovies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/v1/movie-service.proto",
}
