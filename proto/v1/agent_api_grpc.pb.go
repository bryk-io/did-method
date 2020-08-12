// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protov1

import (
	context "context"

	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AgentAPIClient is the client API for AgentAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentAPIClient interface {
	// Reachability test.
	Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PingResponse, error)
	// Process an incoming request ticket.
	Process(ctx context.Context, in *ProcessRequest, opts ...grpc.CallOption) (*ProcessResponse, error)
	// Return the current state of a DID subject.
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type agentAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentAPIClient(cc grpc.ClientConnInterface) AgentAPIClient {
	return &agentAPIClient{cc}
}

func (c *agentAPIClient) Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/bryk.did.proto.v1.AgentAPI/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentAPIClient) Process(ctx context.Context, in *ProcessRequest, opts ...grpc.CallOption) (*ProcessResponse, error) {
	out := new(ProcessResponse)
	err := c.cc.Invoke(ctx, "/bryk.did.proto.v1.AgentAPI/Process", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentAPIClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/bryk.did.proto.v1.AgentAPI/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentAPIServer is the server API for AgentAPI service.
// All implementations must embed UnimplementedAgentAPIServer
// for forward compatibility
type AgentAPIServer interface {
	// Reachability test.
	Ping(context.Context, *empty.Empty) (*PingResponse, error)
	// Process an incoming request ticket.
	Process(context.Context, *ProcessRequest) (*ProcessResponse, error)
	// Return the current state of a DID subject.
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
	mustEmbedUnimplementedAgentAPIServer()
}

// UnimplementedAgentAPIServer must be embedded to have forward compatible implementations.
type UnimplementedAgentAPIServer struct {
}

func (*UnimplementedAgentAPIServer) Ping(context.Context, *empty.Empty) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedAgentAPIServer) Process(context.Context, *ProcessRequest) (*ProcessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Process not implemented")
}
func (*UnimplementedAgentAPIServer) Query(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (*UnimplementedAgentAPIServer) mustEmbedUnimplementedAgentAPIServer() {}

func RegisterAgentAPIServer(s *grpc.Server, srv AgentAPIServer) {
	s.RegisterService(&_AgentAPI_serviceDesc, srv)
}

func _AgentAPI_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentAPIServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bryk.did.proto.v1.AgentAPI/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentAPIServer).Ping(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentAPI_Process_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentAPIServer).Process(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bryk.did.proto.v1.AgentAPI/Process",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentAPIServer).Process(ctx, req.(*ProcessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentAPI_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentAPIServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bryk.did.proto.v1.AgentAPI/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentAPIServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AgentAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bryk.did.proto.v1.AgentAPI",
	HandlerType: (*AgentAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _AgentAPI_Ping_Handler,
		},
		{
			MethodName: "Process",
			Handler:    _AgentAPI_Process_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _AgentAPI_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/agent_api.proto",
}