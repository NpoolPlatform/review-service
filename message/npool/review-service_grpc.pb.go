// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package npool

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ReviewServiceClient is the client API for ReviewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReviewServiceClient interface {
	// Method Version
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
	CreateReview(ctx context.Context, in *CreateReviewRequest, opts ...grpc.CallOption) (*CreateReviewResponse, error)
	UpdateReview(ctx context.Context, in *UpdateReviewRequest, opts ...grpc.CallOption) (*UpdateReviewResponse, error)
	GetReviewsByDomain(ctx context.Context, in *GetReviewsByDomainRequest, opts ...grpc.CallOption) (*GetReviewsByDomainResponse, error)
	CreateReviewRule(ctx context.Context, in *CreateReviewRuleRequest, opts ...grpc.CallOption) (*CreateReviewRuleResponse, error)
	UpdateReviewRule(ctx context.Context, in *UpdateReviewRuleRequest, opts ...grpc.CallOption) (*UpdateReviewRuleResponse, error)
	GetReviewRule(ctx context.Context, in *GetReviewRuleRequest, opts ...grpc.CallOption) (*GetReviewRuleResponse, error)
	GetReviewRulesByDomain(ctx context.Context, in *GetReviewRulesByDomainRequest, opts ...grpc.CallOption) (*GetReviewRulesByDomainResponse, error)
}

type reviewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReviewServiceClient(cc grpc.ClientConnInterface) ReviewServiceClient {
	return &reviewServiceClient{cc}
}

func (c *reviewServiceClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/review.service.v1.ReviewService/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) CreateReview(ctx context.Context, in *CreateReviewRequest, opts ...grpc.CallOption) (*CreateReviewResponse, error) {
	out := new(CreateReviewResponse)
	err := c.cc.Invoke(ctx, "/review.service.v1.ReviewService/CreateReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) UpdateReview(ctx context.Context, in *UpdateReviewRequest, opts ...grpc.CallOption) (*UpdateReviewResponse, error) {
	out := new(UpdateReviewResponse)
	err := c.cc.Invoke(ctx, "/review.service.v1.ReviewService/UpdateReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) GetReviewsByDomain(ctx context.Context, in *GetReviewsByDomainRequest, opts ...grpc.CallOption) (*GetReviewsByDomainResponse, error) {
	out := new(GetReviewsByDomainResponse)
	err := c.cc.Invoke(ctx, "/review.service.v1.ReviewService/GetReviewsByDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) CreateReviewRule(ctx context.Context, in *CreateReviewRuleRequest, opts ...grpc.CallOption) (*CreateReviewRuleResponse, error) {
	out := new(CreateReviewRuleResponse)
	err := c.cc.Invoke(ctx, "/review.service.v1.ReviewService/CreateReviewRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) UpdateReviewRule(ctx context.Context, in *UpdateReviewRuleRequest, opts ...grpc.CallOption) (*UpdateReviewRuleResponse, error) {
	out := new(UpdateReviewRuleResponse)
	err := c.cc.Invoke(ctx, "/review.service.v1.ReviewService/UpdateReviewRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) GetReviewRule(ctx context.Context, in *GetReviewRuleRequest, opts ...grpc.CallOption) (*GetReviewRuleResponse, error) {
	out := new(GetReviewRuleResponse)
	err := c.cc.Invoke(ctx, "/review.service.v1.ReviewService/GetReviewRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) GetReviewRulesByDomain(ctx context.Context, in *GetReviewRulesByDomainRequest, opts ...grpc.CallOption) (*GetReviewRulesByDomainResponse, error) {
	out := new(GetReviewRulesByDomainResponse)
	err := c.cc.Invoke(ctx, "/review.service.v1.ReviewService/GetReviewRulesByDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReviewServiceServer is the server API for ReviewService service.
// All implementations must embed UnimplementedReviewServiceServer
// for forward compatibility
type ReviewServiceServer interface {
	// Method Version
	Version(context.Context, *emptypb.Empty) (*VersionResponse, error)
	CreateReview(context.Context, *CreateReviewRequest) (*CreateReviewResponse, error)
	UpdateReview(context.Context, *UpdateReviewRequest) (*UpdateReviewResponse, error)
	GetReviewsByDomain(context.Context, *GetReviewsByDomainRequest) (*GetReviewsByDomainResponse, error)
	CreateReviewRule(context.Context, *CreateReviewRuleRequest) (*CreateReviewRuleResponse, error)
	UpdateReviewRule(context.Context, *UpdateReviewRuleRequest) (*UpdateReviewRuleResponse, error)
	GetReviewRule(context.Context, *GetReviewRuleRequest) (*GetReviewRuleResponse, error)
	GetReviewRulesByDomain(context.Context, *GetReviewRulesByDomainRequest) (*GetReviewRulesByDomainResponse, error)
	mustEmbedUnimplementedReviewServiceServer()
}

// UnimplementedReviewServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReviewServiceServer struct {
}

func (UnimplementedReviewServiceServer) Version(context.Context, *emptypb.Empty) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedReviewServiceServer) CreateReview(context.Context, *CreateReviewRequest) (*CreateReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReview not implemented")
}
func (UnimplementedReviewServiceServer) UpdateReview(context.Context, *UpdateReviewRequest) (*UpdateReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReview not implemented")
}
func (UnimplementedReviewServiceServer) GetReviewsByDomain(context.Context, *GetReviewsByDomainRequest) (*GetReviewsByDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReviewsByDomain not implemented")
}
func (UnimplementedReviewServiceServer) CreateReviewRule(context.Context, *CreateReviewRuleRequest) (*CreateReviewRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReviewRule not implemented")
}
func (UnimplementedReviewServiceServer) UpdateReviewRule(context.Context, *UpdateReviewRuleRequest) (*UpdateReviewRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReviewRule not implemented")
}
func (UnimplementedReviewServiceServer) GetReviewRule(context.Context, *GetReviewRuleRequest) (*GetReviewRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReviewRule not implemented")
}
func (UnimplementedReviewServiceServer) GetReviewRulesByDomain(context.Context, *GetReviewRulesByDomainRequest) (*GetReviewRulesByDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReviewRulesByDomain not implemented")
}
func (UnimplementedReviewServiceServer) mustEmbedUnimplementedReviewServiceServer() {}

// UnsafeReviewServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReviewServiceServer will
// result in compilation errors.
type UnsafeReviewServiceServer interface {
	mustEmbedUnimplementedReviewServiceServer()
}

func RegisterReviewServiceServer(s grpc.ServiceRegistrar, srv ReviewServiceServer) {
	s.RegisterService(&ReviewService_ServiceDesc, srv)
}

func _ReviewService_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.service.v1.ReviewService/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_CreateReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).CreateReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.service.v1.ReviewService/CreateReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).CreateReview(ctx, req.(*CreateReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_UpdateReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).UpdateReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.service.v1.ReviewService/UpdateReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).UpdateReview(ctx, req.(*UpdateReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_GetReviewsByDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReviewsByDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).GetReviewsByDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.service.v1.ReviewService/GetReviewsByDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).GetReviewsByDomain(ctx, req.(*GetReviewsByDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_CreateReviewRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReviewRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).CreateReviewRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.service.v1.ReviewService/CreateReviewRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).CreateReviewRule(ctx, req.(*CreateReviewRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_UpdateReviewRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReviewRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).UpdateReviewRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.service.v1.ReviewService/UpdateReviewRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).UpdateReviewRule(ctx, req.(*UpdateReviewRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_GetReviewRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReviewRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).GetReviewRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.service.v1.ReviewService/GetReviewRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).GetReviewRule(ctx, req.(*GetReviewRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_GetReviewRulesByDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReviewRulesByDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).GetReviewRulesByDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.service.v1.ReviewService/GetReviewRulesByDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).GetReviewRulesByDomain(ctx, req.(*GetReviewRulesByDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReviewService_ServiceDesc is the grpc.ServiceDesc for ReviewService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReviewService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "review.service.v1.ReviewService",
	HandlerType: (*ReviewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _ReviewService_Version_Handler,
		},
		{
			MethodName: "CreateReview",
			Handler:    _ReviewService_CreateReview_Handler,
		},
		{
			MethodName: "UpdateReview",
			Handler:    _ReviewService_UpdateReview_Handler,
		},
		{
			MethodName: "GetReviewsByDomain",
			Handler:    _ReviewService_GetReviewsByDomain_Handler,
		},
		{
			MethodName: "CreateReviewRule",
			Handler:    _ReviewService_CreateReviewRule_Handler,
		},
		{
			MethodName: "UpdateReviewRule",
			Handler:    _ReviewService_UpdateReviewRule_Handler,
		},
		{
			MethodName: "GetReviewRule",
			Handler:    _ReviewService_GetReviewRule_Handler,
		},
		{
			MethodName: "GetReviewRulesByDomain",
			Handler:    _ReviewService_GetReviewRulesByDomain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/review-service.proto",
}
