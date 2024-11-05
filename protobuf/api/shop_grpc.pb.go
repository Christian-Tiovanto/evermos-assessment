// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: protobuf/api/shop.proto

package commerce_api

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ShopService_CreateShop_FullMethodName      = "/mughieams.evermosassessment.v1.ShopService/CreateShop"
	ShopService_CreateWarehouse_FullMethodName = "/mughieams.evermosassessment.v1.ShopService/CreateWarehouse"
	ShopService_GetWarehouses_FullMethodName   = "/mughieams.evermosassessment.v1.ShopService/GetWarehouses"
	ShopService_AddProduct_FullMethodName      = "/mughieams.evermosassessment.v1.ShopService/AddProduct"
	ShopService_GetProducts_FullMethodName     = "/mughieams.evermosassessment.v1.ShopService/GetProducts"
)

// ShopServiceClient is the client API for ShopService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShopServiceClient interface {
	CreateShop(ctx context.Context, in *CreateShopRequest, opts ...grpc.CallOption) (*MessageResponse, error)
	CreateWarehouse(ctx context.Context, in *CreateWarehouseRequest, opts ...grpc.CallOption) (*MessageResponse, error)
	GetWarehouses(ctx context.Context, in *ShopIDRequest, opts ...grpc.CallOption) (*GetWarehousesResponse, error)
	AddProduct(ctx context.Context, in *AddProductRequest, opts ...grpc.CallOption) (*MessageResponse, error)
	GetProducts(ctx context.Context, in *ShopIDRequest, opts ...grpc.CallOption) (*GetProductsResponse, error)
}

type shopServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShopServiceClient(cc grpc.ClientConnInterface) ShopServiceClient {
	return &shopServiceClient{cc}
}

func (c *shopServiceClient) CreateShop(ctx context.Context, in *CreateShopRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, ShopService_CreateShop_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) CreateWarehouse(ctx context.Context, in *CreateWarehouseRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, ShopService_CreateWarehouse_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) GetWarehouses(ctx context.Context, in *ShopIDRequest, opts ...grpc.CallOption) (*GetWarehousesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetWarehousesResponse)
	err := c.cc.Invoke(ctx, ShopService_GetWarehouses_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) AddProduct(ctx context.Context, in *AddProductRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, ShopService_AddProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) GetProducts(ctx context.Context, in *ShopIDRequest, opts ...grpc.CallOption) (*GetProductsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProductsResponse)
	err := c.cc.Invoke(ctx, ShopService_GetProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopServiceServer is the server API for ShopService service.
// All implementations must embed UnimplementedShopServiceServer
// for forward compatibility.
type ShopServiceServer interface {
	CreateShop(context.Context, *CreateShopRequest) (*MessageResponse, error)
	CreateWarehouse(context.Context, *CreateWarehouseRequest) (*MessageResponse, error)
	GetWarehouses(context.Context, *ShopIDRequest) (*GetWarehousesResponse, error)
	AddProduct(context.Context, *AddProductRequest) (*MessageResponse, error)
	GetProducts(context.Context, *ShopIDRequest) (*GetProductsResponse, error)
	mustEmbedUnimplementedShopServiceServer()
}

// UnimplementedShopServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedShopServiceServer struct{}

func (UnimplementedShopServiceServer) CreateShop(context.Context, *CreateShopRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShop not implemented")
}
func (UnimplementedShopServiceServer) CreateWarehouse(context.Context, *CreateWarehouseRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWarehouse not implemented")
}
func (UnimplementedShopServiceServer) GetWarehouses(context.Context, *ShopIDRequest) (*GetWarehousesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWarehouses not implemented")
}
func (UnimplementedShopServiceServer) AddProduct(context.Context, *AddProductRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProduct not implemented")
}
func (UnimplementedShopServiceServer) GetProducts(context.Context, *ShopIDRequest) (*GetProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProducts not implemented")
}
func (UnimplementedShopServiceServer) mustEmbedUnimplementedShopServiceServer() {}
func (UnimplementedShopServiceServer) testEmbeddedByValue()                     {}

// UnsafeShopServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShopServiceServer will
// result in compilation errors.
type UnsafeShopServiceServer interface {
	mustEmbedUnimplementedShopServiceServer()
}

func RegisterShopServiceServer(s grpc.ServiceRegistrar, srv ShopServiceServer) {
	// If the following call pancis, it indicates UnimplementedShopServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ShopService_ServiceDesc, srv)
}

func _ShopService_CreateShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).CreateShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_CreateShop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).CreateShop(ctx, req.(*CreateShopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_CreateWarehouse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWarehouseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).CreateWarehouse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_CreateWarehouse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).CreateWarehouse(ctx, req.(*CreateWarehouseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_GetWarehouses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).GetWarehouses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_GetWarehouses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).GetWarehouses(ctx, req.(*ShopIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_AddProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).AddProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_AddProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).AddProduct(ctx, req.(*AddProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_GetProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).GetProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_GetProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).GetProducts(ctx, req.(*ShopIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShopService_ServiceDesc is the grpc.ServiceDesc for ShopService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShopService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mughieams.evermosassessment.v1.ShopService",
	HandlerType: (*ShopServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateShop",
			Handler:    _ShopService_CreateShop_Handler,
		},
		{
			MethodName: "CreateWarehouse",
			Handler:    _ShopService_CreateWarehouse_Handler,
		},
		{
			MethodName: "GetWarehouses",
			Handler:    _ShopService_GetWarehouses_Handler,
		},
		{
			MethodName: "AddProduct",
			Handler:    _ShopService_AddProduct_Handler,
		},
		{
			MethodName: "GetProducts",
			Handler:    _ShopService_GetProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/api/shop.proto",
}
