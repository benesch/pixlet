//Generated by gRPC Go plugin
//If you make any local changes, they will be lost
//source: canvas

package fb

import (
	context "context"
	flatbuffers "github.com/google/flatbuffers/go"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Client API for Canvas service
type CanvasClient interface {
	Perform(ctx context.Context, in *flatbuffers.Builder,
		opts ...grpc.CallOption) (*Image, error)
}

type canvasClient struct {
	cc grpc.ClientConnInterface
}

func NewCanvasClient(cc grpc.ClientConnInterface) CanvasClient {
	return &canvasClient{cc}
}

func (c *canvasClient) Perform(ctx context.Context, in *flatbuffers.Builder,
	opts ...grpc.CallOption) (*Image, error) {
	out := new(Image)
	err := c.cc.Invoke(ctx, "/fb.Canvas/Perform", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Canvas service
type CanvasServer interface {
	Perform(context.Context, *CanvasOperationList) (*flatbuffers.Builder, error)
	mustEmbedUnimplementedCanvasServer()
}

type UnimplementedCanvasServer struct {
}

func (UnimplementedCanvasServer) Perform(context.Context, *CanvasOperationList) (*flatbuffers.Builder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Perform not implemented")
}

func (UnimplementedCanvasServer) mustEmbedUnimplementedCanvasServer() {}

type UnsafeCanvasServer interface {
	mustEmbedUnimplementedCanvasServer()
}

func RegisterCanvasServer(s grpc.ServiceRegistrar, srv CanvasServer) {
	s.RegisterService(&_Canvas_serviceDesc, srv)
}

func _Canvas_Perform_Handler(srv interface{}, ctx context.Context,
	dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CanvasOperationList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CanvasServer).Perform(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fb.Canvas/Perform",
	}

	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CanvasServer).Perform(ctx, req.(*CanvasOperationList))
	}
	return interceptor(ctx, in, info, handler)
}
var _Canvas_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fb.Canvas",
	HandlerType: (*CanvasServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Perform",
			Handler:    _Canvas_Perform_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
	},
}
