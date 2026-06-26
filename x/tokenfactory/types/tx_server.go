package types

import (
	context "context"

	"google.golang.org/grpc"
)

type MsgServer interface {
	CreateDenom(context.Context, *MsgCreateDenom) (*MsgCreateDenomResponse, error)
	Mint(context.Context, *MsgMint) (*MsgMintResponse, error)
	Burn(context.Context, *MsgBurn) (*MsgBurnResponse, error)
	ChangeAdmin(context.Context, *MsgChangeAdmin) (*MsgChangeAdminResponse, error)
	ForceTransfer(context.Context, *MsgForceTransfer) (*MsgForceTransferResponse, error)
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateDenom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateDenom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateDenom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heya.tokenfactory.v1.Msg/CreateDenom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateDenom(ctx, req.(*MsgCreateDenom))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Mint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Mint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heya.tokenfactory.v1.Msg/Mint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Mint(ctx, req.(*MsgMint))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Burn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBurn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Burn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heya.tokenfactory.v1.Msg/Burn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Burn(ctx, req.(*MsgBurn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ChangeAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgChangeAdmin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ChangeAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heya.tokenfactory.v1.Msg/ChangeAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ChangeAdmin(ctx, req.(*MsgChangeAdmin))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ForceTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgForceTransfer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ForceTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heya.tokenfactory.v1.Msg/ForceTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ForceTransfer(ctx, req.(*MsgForceTransfer))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "heya.tokenfactory.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDenom",
			Handler:    _Msg_CreateDenom_Handler,
		},
		{
			MethodName: "Mint",
			Handler:    _Msg_Mint_Handler,
		},
		{
			MethodName: "Burn",
			Handler:    _Msg_Burn_Handler,
		},
		{
			MethodName: "ChangeAdmin",
			Handler:    _Msg_ChangeAdmin_Handler,
		},
		{
			MethodName: "ForceTransfer",
			Handler:    _Msg_ForceTransfer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "heya/tokenfactory/v1/tx.proto",
}
