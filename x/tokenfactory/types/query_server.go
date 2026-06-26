package types

import (
	context "context"

	"github.com/cosmos/cosmos-sdk/client"
	"google.golang.org/grpc"
)

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(c client.Context) QueryServer {
	return &queryClient{cc: c.GRPCClient}
}

func (c *queryClient) DenomAdmin(ctx context.Context, req *QueryDenomAdminRequest) (*QueryDenomAdminResponse, error) {
	out := new(QueryDenomAdminResponse)
	err := c.cc.Invoke(ctx, "/heya.tokenfactory.v1.Query/DenomAdmin", req, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/heya.tokenfactory.v1.Query/Params", req, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type QueryServer interface {
	DenomAdmin(context.Context, *QueryDenomAdminRequest) (*QueryDenomAdminResponse, error)
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_DenomAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDenomAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DenomAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heya.tokenfactory.v1.Query/DenomAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DenomAdmin(ctx, req.(*QueryDenomAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heya.tokenfactory.v1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "heya.tokenfactory.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DenomAdmin",
			Handler:    _Query_DenomAdmin_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "heya/tokenfactory/v1/query.proto",
}
