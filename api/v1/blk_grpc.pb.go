// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// BlkClient is the client API for Blk service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlkClient interface {
	GetBalance(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*Balance, error)
}

type blkClient struct {
	cc grpc.ClientConnInterface
}

func NewBlkClient(cc grpc.ClientConnInterface) BlkClient {
	return &blkClient{cc}
}

func (c *blkClient) GetBalance(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*Balance, error) {
	out := new(Balance)
	err := c.cc.Invoke(ctx, "/api.v1.Blk/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlkServer is the server API for Blk service.
// All implementations must embed UnimplementedBlkServer
// for forward compatibility
type BlkServer interface {
	GetBalance(context.Context, *AccountRequest) (*Balance, error)
	mustEmbedUnimplementedBlkServer()
}

// UnimplementedBlkServer must be embedded to have forward compatible implementations.
type UnimplementedBlkServer struct {
}

func (UnimplementedBlkServer) GetBalance(context.Context, *AccountRequest) (*Balance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedBlkServer) mustEmbedUnimplementedBlkServer() {}

// UnsafeBlkServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlkServer will
// result in compilation errors.
type UnsafeBlkServer interface {
	mustEmbedUnimplementedBlkServer()
}

func RegisterBlkServer(s *grpc.Server, srv BlkServer) {
	s.RegisterService(&_Blk_serviceDesc, srv)
}

func _Blk_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlkServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.Blk/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlkServer).GetBalance(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Blk_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.Blk",
	HandlerType: (*BlkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBalance",
			Handler:    _Blk_GetBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/blk.proto",
}

// BlkWalletClient is the client API for BlkWallet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlkWalletClient interface {
	// rpc BuyToken(TokenTxRequest)returns (Wallets){}
	// rpc Transfer(TokenTxRequest)returns (Wallets){}
	GetWallets(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*Wallets, error)
}

type blkWalletClient struct {
	cc grpc.ClientConnInterface
}

func NewBlkWalletClient(cc grpc.ClientConnInterface) BlkWalletClient {
	return &blkWalletClient{cc}
}

func (c *blkWalletClient) GetWallets(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*Wallets, error) {
	out := new(Wallets)
	err := c.cc.Invoke(ctx, "/api.v1.BlkWallet/GetWallets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlkWalletServer is the server API for BlkWallet service.
// All implementations must embed UnimplementedBlkWalletServer
// for forward compatibility
type BlkWalletServer interface {
	// rpc BuyToken(TokenTxRequest)returns (Wallets){}
	// rpc Transfer(TokenTxRequest)returns (Wallets){}
	GetWallets(context.Context, *AccountRequest) (*Wallets, error)
	mustEmbedUnimplementedBlkWalletServer()
}

// UnimplementedBlkWalletServer must be embedded to have forward compatible implementations.
type UnimplementedBlkWalletServer struct {
}

func (UnimplementedBlkWalletServer) GetWallets(context.Context, *AccountRequest) (*Wallets, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWallets not implemented")
}
func (UnimplementedBlkWalletServer) mustEmbedUnimplementedBlkWalletServer() {}

// UnsafeBlkWalletServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlkWalletServer will
// result in compilation errors.
type UnsafeBlkWalletServer interface {
	mustEmbedUnimplementedBlkWalletServer()
}

func RegisterBlkWalletServer(s *grpc.Server, srv BlkWalletServer) {
	s.RegisterService(&_BlkWallet_serviceDesc, srv)
}

func _BlkWallet_GetWallets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlkWalletServer).GetWallets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.BlkWallet/GetWallets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlkWalletServer).GetWallets(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlkWallet_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.BlkWallet",
	HandlerType: (*BlkWalletServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWallets",
			Handler:    _BlkWallet_GetWallets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/blk.proto",
}
