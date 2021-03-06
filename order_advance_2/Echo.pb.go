// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: Echo.proto

package order

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type EchoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EchoReq) Reset() {
	*x = EchoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Echo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoReq) ProtoMessage() {}

func (x *EchoReq) ProtoReflect() protoreflect.Message {
	mi := &file_Echo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoReq.ProtoReflect.Descriptor instead.
func (*EchoReq) Descriptor() ([]byte, []int) {
	return file_Echo_proto_rawDescGZIP(), []int{0}
}

func (x *EchoReq) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type EchoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Addr    string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (x *EchoResp) Reset() {
	*x = EchoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Echo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoResp) ProtoMessage() {}

func (x *EchoResp) ProtoReflect() protoreflect.Message {
	mi := &file_Echo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoResp.ProtoReflect.Descriptor instead.
func (*EchoResp) Descriptor() ([]byte, []int) {
	return file_Echo_proto_rawDescGZIP(), []int{1}
}

func (x *EchoResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *EchoResp) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

var File_Echo_proto protoreflect.FileDescriptor

var file_Echo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x45, 0x63, 0x68, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x07, 0x65, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x38, 0x0a, 0x08, 0x65, 0x63, 0x68, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64,
	0x64, 0x72, 0x32, 0x3b, 0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x2c, 0x0a, 0x09, 0x75, 0x6e, 0x61, 0x72, 0x79, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x0e,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x0f,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Echo_proto_rawDescOnce sync.Once
	file_Echo_proto_rawDescData = file_Echo_proto_rawDesc
)

func file_Echo_proto_rawDescGZIP() []byte {
	file_Echo_proto_rawDescOnce.Do(func() {
		file_Echo_proto_rawDescData = protoimpl.X.CompressGZIP(file_Echo_proto_rawDescData)
	})
	return file_Echo_proto_rawDescData
}

var file_Echo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Echo_proto_goTypes = []interface{}{
	(*EchoReq)(nil),  // 0: order.echoReq
	(*EchoResp)(nil), // 1: order.echoResp
}
var file_Echo_proto_depIdxs = []int32{
	0, // 0: order.EchoService.unaryEcho:input_type -> order.echoReq
	1, // 1: order.EchoService.unaryEcho:output_type -> order.echoResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Echo_proto_init() }
func file_Echo_proto_init() {
	if File_Echo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Echo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Echo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Echo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Echo_proto_goTypes,
		DependencyIndexes: file_Echo_proto_depIdxs,
		MessageInfos:      file_Echo_proto_msgTypes,
	}.Build()
	File_Echo_proto = out.File
	file_Echo_proto_rawDesc = nil
	file_Echo_proto_goTypes = nil
	file_Echo_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EchoServiceClient is the client API for EchoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoServiceClient interface {
	UnaryEcho(ctx context.Context, in *EchoReq, opts ...grpc.CallOption) (*EchoResp, error)
}

type echoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEchoServiceClient(cc grpc.ClientConnInterface) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) UnaryEcho(ctx context.Context, in *EchoReq, opts ...grpc.CallOption) (*EchoResp, error) {
	out := new(EchoResp)
	err := c.cc.Invoke(ctx, "/order.EchoService/unaryEcho", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoServiceServer is the server API for EchoService service.
type EchoServiceServer interface {
	UnaryEcho(context.Context, *EchoReq) (*EchoResp, error)
}

// UnimplementedEchoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEchoServiceServer struct {
}

func (*UnimplementedEchoServiceServer) UnaryEcho(context.Context, *EchoReq) (*EchoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryEcho not implemented")
}

func RegisterEchoServiceServer(s *grpc.Server, srv EchoServiceServer) {
	s.RegisterService(&_EchoService_serviceDesc, srv)
}

func _EchoService_UnaryEcho_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).UnaryEcho(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.EchoService/UnaryEcho",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).UnaryEcho(ctx, req.(*EchoReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _EchoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "order.EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "unaryEcho",
			Handler:    _EchoService_UnaryEcho_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Echo.proto",
}
