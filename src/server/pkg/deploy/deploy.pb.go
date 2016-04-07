// Code generated by protoc-gen-go.
// source: server/pkg/deploy/deploy.proto
// DO NOT EDIT!

/*
Package deploy is a generated protocol buffer package.

It is generated from these files:
	server/pkg/deploy/deploy.proto

It has these top-level messages:
	KubeEndpoint
	Cluster
	ClusterInfo
	ClusterInfos
	CreateClusterRequest
	UpdateClusterRequest
	InspectClusterRequest
	ListClusterRequest
	DeleteClusterRequest
*/
package deploy

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gengo/grpc-gateway/third_party/googleapis/google/api"
import google_protobuf1 "go.pedge.io/pb/go/google/protobuf"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type KubeEndpoint struct {
	Host string `protobuf:"bytes,1,opt,name=host" json:"host,omitempty"`
}

func (m *KubeEndpoint) Reset()                    { *m = KubeEndpoint{} }
func (m *KubeEndpoint) String() string            { return proto.CompactTextString(m) }
func (*KubeEndpoint) ProtoMessage()               {}
func (*KubeEndpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Cluster struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Cluster) Reset()                    { *m = Cluster{} }
func (m *Cluster) String() string            { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()               {}
func (*Cluster) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ClusterInfo struct {
	Cluster *Cluster `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Shards  uint64   `protobuf:"varint,2,opt,name=shards" json:"shards,omitempty"`
}

func (m *ClusterInfo) Reset()                    { *m = ClusterInfo{} }
func (m *ClusterInfo) String() string            { return proto.CompactTextString(m) }
func (*ClusterInfo) ProtoMessage()               {}
func (*ClusterInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ClusterInfo) GetCluster() *Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

type ClusterInfos struct {
	ClusterInfos []*ClusterInfo `protobuf:"bytes,1,rep,name=cluster_infos,json=clusterInfos" json:"cluster_infos,omitempty"`
}

func (m *ClusterInfos) Reset()                    { *m = ClusterInfos{} }
func (m *ClusterInfos) String() string            { return proto.CompactTextString(m) }
func (*ClusterInfos) ProtoMessage()               {}
func (*ClusterInfos) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ClusterInfos) GetClusterInfos() []*ClusterInfo {
	if m != nil {
		return m.ClusterInfos
	}
	return nil
}

type CreateClusterRequest struct {
	Cluster *Cluster `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Shards  uint64   `protobuf:"varint,2,opt,name=shards" json:"shards,omitempty"`
}

func (m *CreateClusterRequest) Reset()                    { *m = CreateClusterRequest{} }
func (m *CreateClusterRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateClusterRequest) ProtoMessage()               {}
func (*CreateClusterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CreateClusterRequest) GetCluster() *Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

type UpdateClusterRequest struct {
	Cluster *Cluster `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Nodes   uint64   `protobuf:"varint,2,opt,name=nodes" json:"nodes,omitempty"`
}

func (m *UpdateClusterRequest) Reset()                    { *m = UpdateClusterRequest{} }
func (m *UpdateClusterRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateClusterRequest) ProtoMessage()               {}
func (*UpdateClusterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateClusterRequest) GetCluster() *Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

type InspectClusterRequest struct {
	Cluster *Cluster `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
}

func (m *InspectClusterRequest) Reset()                    { *m = InspectClusterRequest{} }
func (m *InspectClusterRequest) String() string            { return proto.CompactTextString(m) }
func (*InspectClusterRequest) ProtoMessage()               {}
func (*InspectClusterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *InspectClusterRequest) GetCluster() *Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

type ListClusterRequest struct {
}

func (m *ListClusterRequest) Reset()                    { *m = ListClusterRequest{} }
func (m *ListClusterRequest) String() string            { return proto.CompactTextString(m) }
func (*ListClusterRequest) ProtoMessage()               {}
func (*ListClusterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type DeleteClusterRequest struct {
	Cluster *Cluster `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
}

func (m *DeleteClusterRequest) Reset()                    { *m = DeleteClusterRequest{} }
func (m *DeleteClusterRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteClusterRequest) ProtoMessage()               {}
func (*DeleteClusterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *DeleteClusterRequest) GetCluster() *Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

func init() {
	proto.RegisterType((*KubeEndpoint)(nil), "deploy.KubeEndpoint")
	proto.RegisterType((*Cluster)(nil), "deploy.Cluster")
	proto.RegisterType((*ClusterInfo)(nil), "deploy.ClusterInfo")
	proto.RegisterType((*ClusterInfos)(nil), "deploy.ClusterInfos")
	proto.RegisterType((*CreateClusterRequest)(nil), "deploy.CreateClusterRequest")
	proto.RegisterType((*UpdateClusterRequest)(nil), "deploy.UpdateClusterRequest")
	proto.RegisterType((*InspectClusterRequest)(nil), "deploy.InspectClusterRequest")
	proto.RegisterType((*ListClusterRequest)(nil), "deploy.ListClusterRequest")
	proto.RegisterType((*DeleteClusterRequest)(nil), "deploy.DeleteClusterRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.

// Client API for API service

type APIClient interface {
	CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	UpdateCluster(ctx context.Context, in *UpdateClusterRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	InspectCluster(ctx context.Context, in *InspectClusterRequest, opts ...grpc.CallOption) (*ClusterInfo, error)
	ListCluster(ctx context.Context, in *ListClusterRequest, opts ...grpc.CallOption) (*ClusterInfos, error)
	DeleteCluster(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/deploy.API/CreateCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) UpdateCluster(ctx context.Context, in *UpdateClusterRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/deploy.API/UpdateCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) InspectCluster(ctx context.Context, in *InspectClusterRequest, opts ...grpc.CallOption) (*ClusterInfo, error) {
	out := new(ClusterInfo)
	err := grpc.Invoke(ctx, "/deploy.API/InspectCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListCluster(ctx context.Context, in *ListClusterRequest, opts ...grpc.CallOption) (*ClusterInfos, error) {
	out := new(ClusterInfos)
	err := grpc.Invoke(ctx, "/deploy.API/ListCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) DeleteCluster(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/deploy.API/DeleteCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for API service

type APIServer interface {
	CreateCluster(context.Context, *CreateClusterRequest) (*google_protobuf1.Empty, error)
	UpdateCluster(context.Context, *UpdateClusterRequest) (*google_protobuf1.Empty, error)
	InspectCluster(context.Context, *InspectClusterRequest) (*ClusterInfo, error)
	ListCluster(context.Context, *ListClusterRequest) (*ClusterInfos, error)
	DeleteCluster(context.Context, *DeleteClusterRequest) (*google_protobuf1.Empty, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_CreateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CreateClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).CreateCluster(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_UpdateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(UpdateClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).UpdateCluster(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_InspectCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(InspectClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).InspectCluster(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_ListCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ListClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).ListCluster(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_DeleteCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DeleteClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).DeleteCluster(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "deploy.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCluster",
			Handler:    _API_CreateCluster_Handler,
		},
		{
			MethodName: "UpdateCluster",
			Handler:    _API_UpdateCluster_Handler,
		},
		{
			MethodName: "InspectCluster",
			Handler:    _API_InspectCluster_Handler,
		},
		{
			MethodName: "ListCluster",
			Handler:    _API_ListCluster_Handler,
		},
		{
			MethodName: "DeleteCluster",
			Handler:    _API_DeleteCluster_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x92, 0x4f, 0x6f, 0xda, 0x40,
	0x10, 0xc5, 0xcb, 0x9f, 0x82, 0x3a, 0x40, 0x2b, 0x6d, 0x5d, 0x84, 0x5c, 0xa8, 0xaa, 0x3d, 0xb5,
	0x17, 0x5b, 0xa2, 0x97, 0x5e, 0x09, 0x21, 0x09, 0x4a, 0x0e, 0xc8, 0x52, 0x14, 0xe5, 0x14, 0x19,
	0x3c, 0x80, 0x15, 0xb3, 0xbb, 0xf1, 0xae, 0x23, 0xe5, 0x33, 0xe7, 0x4b, 0xc4, 0x7f, 0xd6, 0x01,
	0x13, 0x27, 0x87, 0x90, 0x93, 0x77, 0xf6, 0xbd, 0xfd, 0x79, 0xe6, 0x69, 0xe0, 0x97, 0xc4, 0xf0,
	0x1e, 0x43, 0x5b, 0xdc, 0xae, 0x6c, 0x0f, 0x45, 0xc0, 0x1f, 0xf4, 0xc7, 0x12, 0x21, 0x57, 0x9c,
	0x34, 0xb2, 0xca, 0xec, 0xaf, 0x38, 0x5f, 0x05, 0x68, 0xbb, 0xc2, 0xb7, 0x5d, 0xc6, 0xb8, 0x72,
	0x95, 0xcf, 0x99, 0xcc, 0x5c, 0xe6, 0x4f, 0xad, 0xa6, 0xd5, 0x3c, 0x5a, 0xda, 0xb8, 0x11, 0x4a,
	0x23, 0x28, 0x85, 0xf6, 0x79, 0x34, 0xc7, 0x09, 0xf3, 0x04, 0xf7, 0x99, 0x22, 0x04, 0xea, 0x6b,
	0x2e, 0x55, 0xaf, 0xf2, 0xbb, 0xf2, 0xe7, 0x8b, 0x93, 0x9e, 0xe9, 0x00, 0x9a, 0xe3, 0x20, 0x92,
	0x0a, 0xc3, 0x44, 0x66, 0xee, 0x06, 0x73, 0x39, 0x39, 0xd3, 0x19, 0xb4, 0xb4, 0x3c, 0x65, 0x4b,
	0x4e, 0xfe, 0x42, 0x73, 0x91, 0x95, 0xa9, 0xab, 0x35, 0xfc, 0x66, 0xe9, 0xa6, 0xb5, 0xcb, 0xc9,
	0x75, 0xd2, 0x85, 0x86, 0x5c, 0xbb, 0xa1, 0x27, 0x7b, 0xd5, 0xd8, 0x59, 0x77, 0x74, 0x45, 0xcf,
	0xa0, 0xbd, 0x43, 0x94, 0xe4, 0x3f, 0x74, 0xf4, 0x93, 0x1b, 0x3f, 0xb9, 0x88, 0xc1, 0xb5, 0x18,
	0xfc, 0x7d, 0x0f, 0x9c, 0x98, 0x9d, 0xf6, 0x62, 0xe7, 0x25, 0xbd, 0x06, 0x63, 0x1c, 0xa2, 0xab,
	0x30, 0xff, 0x37, 0xde, 0x45, 0x28, 0xd5, 0x47, 0x34, 0x79, 0x05, 0xc6, 0xa5, 0xf0, 0x0e, 0x42,
	0x1b, 0xf0, 0x99, 0x71, 0x0f, 0x73, 0x72, 0x56, 0xd0, 0x23, 0xf8, 0x31, 0x65, 0x52, 0xe0, 0x42,
	0xbd, 0x9b, 0x4c, 0x0d, 0x20, 0x17, 0xbe, 0xdc, 0x03, 0xd0, 0x11, 0x18, 0xc7, 0x18, 0xe0, 0x01,
	0x2d, 0x0f, 0x1f, 0xab, 0x50, 0x1b, 0xcd, 0xa6, 0xe4, 0x14, 0x3a, 0x85, 0x60, 0x49, 0xff, 0xf9,
	0x49, 0x49, 0xde, 0x66, 0xd7, 0xca, 0x96, 0xd0, 0xca, 0x97, 0xd0, 0x9a, 0x24, 0x4b, 0x48, 0x3f,
	0x25, 0xa0, 0x42, 0x8c, 0x5b, 0x50, 0x59, 0xba, 0x6f, 0x80, 0x4e, 0xe0, 0x6b, 0x31, 0x36, 0x32,
	0xc8, 0x49, 0xa5, 0x71, 0x9a, 0x65, 0xeb, 0x13, 0x73, 0x46, 0xd0, 0xda, 0x89, 0x8e, 0x98, 0xb9,
	0xeb, 0x65, 0x9e, 0xa6, 0x51, 0x42, 0x90, 0xd9, 0x4c, 0x85, 0x9c, 0xb7, 0x33, 0x95, 0xc5, 0xff,
	0xfa, 0x4c, 0xf3, 0x46, 0x7a, 0xf3, 0xef, 0x29, 0x00, 0x00, 0xff, 0xff, 0x16, 0x2d, 0x2a, 0xb8,
	0x09, 0x04, 0x00, 0x00,
}
