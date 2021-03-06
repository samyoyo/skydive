// Code generated by protoc-gen-go.
// source: flow/flow.proto
// DO NOT EDIT!

/*
Package flow is a generated protocol buffer package.

It is generated from these files:
	flow/flow.proto

It has these top-level messages:
	FlowEndpointStatistics
	FlowEndpointsStatistics
	FlowStatistics
	Flow
*/
package flow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type FlowEndpointLayer int32

const (
	FlowEndpointLayer_LINK      FlowEndpointLayer = 0
	FlowEndpointLayer_NETWORK   FlowEndpointLayer = 1
	FlowEndpointLayer_TRANSPORT FlowEndpointLayer = 2
)

var FlowEndpointLayer_name = map[int32]string{
	0: "LINK",
	1: "NETWORK",
	2: "TRANSPORT",
}
var FlowEndpointLayer_value = map[string]int32{
	"LINK":      0,
	"NETWORK":   1,
	"TRANSPORT": 2,
}

func (x FlowEndpointLayer) String() string {
	return proto.EnumName(FlowEndpointLayer_name, int32(x))
}
func (FlowEndpointLayer) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type FlowEndpointType int32

const (
	FlowEndpointType_ETHERNET FlowEndpointType = 0
	FlowEndpointType_IPV4     FlowEndpointType = 1
	FlowEndpointType_TCPPORT  FlowEndpointType = 2
	FlowEndpointType_UDPPORT  FlowEndpointType = 3
	FlowEndpointType_SCTPPORT FlowEndpointType = 4
)

var FlowEndpointType_name = map[int32]string{
	0: "ETHERNET",
	1: "IPV4",
	2: "TCPPORT",
	3: "UDPPORT",
	4: "SCTPPORT",
}
var FlowEndpointType_value = map[string]int32{
	"ETHERNET": 0,
	"IPV4":     1,
	"TCPPORT":  2,
	"UDPPORT":  3,
	"SCTPPORT": 4,
}

func (x FlowEndpointType) String() string {
	return proto.EnumName(FlowEndpointType_name, int32(x))
}
func (FlowEndpointType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type FlowEndpointStatistics struct {
	Value   string `protobuf:"bytes,2,opt,name=Value" json:"Value,omitempty"`
	Packets uint64 `protobuf:"varint,5,opt,name=Packets" json:"Packets,omitempty"`
	Bytes   uint64 `protobuf:"varint,6,opt,name=Bytes" json:"Bytes,omitempty"`
}

func (m *FlowEndpointStatistics) Reset()                    { *m = FlowEndpointStatistics{} }
func (m *FlowEndpointStatistics) String() string            { return proto.CompactTextString(m) }
func (*FlowEndpointStatistics) ProtoMessage()               {}
func (*FlowEndpointStatistics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type FlowEndpointsStatistics struct {
	Type FlowEndpointType        `protobuf:"varint,1,opt,name=Type,enum=flow.FlowEndpointType" json:"Type,omitempty"`
	Hash []byte                  `protobuf:"bytes,2,opt,name=Hash,proto3" json:"Hash,omitempty"`
	AB   *FlowEndpointStatistics `protobuf:"bytes,3,opt,name=AB" json:"AB,omitempty"`
	BA   *FlowEndpointStatistics `protobuf:"bytes,4,opt,name=BA" json:"BA,omitempty"`
}

func (m *FlowEndpointsStatistics) Reset()                    { *m = FlowEndpointsStatistics{} }
func (m *FlowEndpointsStatistics) String() string            { return proto.CompactTextString(m) }
func (*FlowEndpointsStatistics) ProtoMessage()               {}
func (*FlowEndpointsStatistics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FlowEndpointsStatistics) GetAB() *FlowEndpointStatistics {
	if m != nil {
		return m.AB
	}
	return nil
}

func (m *FlowEndpointsStatistics) GetBA() *FlowEndpointStatistics {
	if m != nil {
		return m.BA
	}
	return nil
}

type FlowStatistics struct {
	Start     int64                      `protobuf:"varint,1,opt,name=Start" json:"Start,omitempty"`
	Last      int64                      `protobuf:"varint,2,opt,name=Last" json:"Last,omitempty"`
	Endpoints []*FlowEndpointsStatistics `protobuf:"bytes,3,rep,name=Endpoints" json:"Endpoints,omitempty"`
}

func (m *FlowStatistics) Reset()                    { *m = FlowStatistics{} }
func (m *FlowStatistics) String() string            { return proto.CompactTextString(m) }
func (*FlowStatistics) ProtoMessage()               {}
func (*FlowStatistics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FlowStatistics) GetEndpoints() []*FlowEndpointsStatistics {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

type Flow struct {
	UUID       string `protobuf:"bytes,1,opt,name=UUID" json:"UUID,omitempty"`
	LayersPath string `protobuf:"bytes,2,opt,name=LayersPath" json:"LayersPath,omitempty"`
	// Data Flow info
	Statistics *FlowStatistics `protobuf:"bytes,3,opt,name=Statistics" json:"Statistics,omitempty"`
	FlowUUID   string          `protobuf:"bytes,5,opt,name=FlowUUID" json:"FlowUUID,omitempty"`
	// Topology info
	ProbeGraphPath string `protobuf:"bytes,11,opt,name=ProbeGraphPath" json:"ProbeGraphPath,omitempty"`
	IfSrcGraphPath string `protobuf:"bytes,14,opt,name=IfSrcGraphPath" json:"IfSrcGraphPath,omitempty"`
	IfDstGraphPath string `protobuf:"bytes,19,opt,name=IfDstGraphPath" json:"IfDstGraphPath,omitempty"`
}

func (m *Flow) Reset()                    { *m = Flow{} }
func (m *Flow) String() string            { return proto.CompactTextString(m) }
func (*Flow) ProtoMessage()               {}
func (*Flow) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Flow) GetStatistics() *FlowStatistics {
	if m != nil {
		return m.Statistics
	}
	return nil
}

func init() {
	proto.RegisterType((*FlowEndpointStatistics)(nil), "flow.FlowEndpointStatistics")
	proto.RegisterType((*FlowEndpointsStatistics)(nil), "flow.FlowEndpointsStatistics")
	proto.RegisterType((*FlowStatistics)(nil), "flow.FlowStatistics")
	proto.RegisterType((*Flow)(nil), "flow.Flow")
	proto.RegisterEnum("flow.FlowEndpointLayer", FlowEndpointLayer_name, FlowEndpointLayer_value)
	proto.RegisterEnum("flow.FlowEndpointType", FlowEndpointType_name, FlowEndpointType_value)
}

var fileDescriptor0 = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x25, 0xf6, 0xba, 0x89, 0x27, 0xc5, 0x84, 0xa5, 0x2a, 0x16, 0x02, 0x84, 0x72, 0x40, 0x28,
	0x42, 0x45, 0x2a, 0xbd, 0x20, 0x4e, 0x49, 0x13, 0x68, 0xd4, 0x2a, 0xb5, 0xd6, 0x4e, 0xb9, 0x21,
	0xad, 0x83, 0x23, 0x5b, 0x58, 0xb5, 0xe5, 0xdd, 0x52, 0xe5, 0xce, 0x2f, 0xf1, 0x7f, 0xcc, 0xac,
	0xd3, 0xda, 0xa1, 0x17, 0x2e, 0x9b, 0x79, 0x6f, 0xdf, 0x9b, 0x37, 0x3b, 0x49, 0xe0, 0xc9, 0x3a,
	0x2f, 0x6e, 0x3f, 0xd0, 0x71, 0x54, 0x56, 0x85, 0x2e, 0x38, 0xa3, 0x7a, 0xf8, 0x1d, 0x0e, 0xbf,
	0xe0, 0xe7, 0xec, 0xfa, 0x47, 0x59, 0x64, 0xd7, 0x3a, 0xd4, 0x52, 0x67, 0x4a, 0x67, 0x2b, 0xc5,
	0x0f, 0xc0, 0xb9, 0x92, 0xf9, 0x4d, 0xe2, 0x5b, 0x6f, 0x3a, 0xef, 0x5c, 0xe1, 0xfc, 0x22, 0xc0,
	0x7d, 0xe8, 0x06, 0x72, 0xf5, 0x33, 0xd1, 0xca, 0x77, 0x90, 0x67, 0xa2, 0x5b, 0xd6, 0x90, 0xf4,
	0x93, 0x8d, 0x4e, 0x94, 0xbf, 0x67, 0x78, 0x27, 0x26, 0x30, 0xfc, 0xd3, 0x81, 0xe7, 0xed, 0x00,
	0xd5, 0x4a, 0x18, 0x01, 0x8b, 0x36, 0x65, 0xe2, 0x77, 0xd0, 0xe0, 0x1d, 0x1f, 0x1e, 0x99, 0xe1,
	0xda, 0x62, 0xba, 0x15, 0x4c, 0xe3, 0xc9, 0x39, 0xb0, 0x33, 0xa9, 0x52, 0x33, 0xcc, 0xbe, 0x60,
	0x29, 0xd6, 0xfc, 0x3d, 0x58, 0xe3, 0x89, 0x6f, 0x23, 0xd3, 0x3f, 0x7e, 0xf9, 0xd0, 0xdd, 0x24,
	0x09, 0x4b, 0x4e, 0x48, 0x3d, 0x19, 0xfb, 0xec, 0x7f, 0xd4, 0xf1, 0x78, 0x78, 0x0b, 0x1e, 0xdd,
	0xee, 0xee, 0x03, 0x51, 0xa5, 0xcd, 0xb8, 0xb6, 0x70, 0x14, 0x01, 0x9a, 0xeb, 0x42, 0x2a, 0x6d,
	0xe6, 0xb2, 0x05, 0xcb, 0xb1, 0xe6, 0x9f, 0xc1, 0xbd, 0x7f, 0x2e, 0x8e, 0x67, 0x63, 0xe0, 0xab,
	0x87, 0x81, 0xad, 0x4d, 0x08, 0x37, 0xb9, 0x23, 0x87, 0xbf, 0x2d, 0x60, 0x24, 0xa3, 0xce, 0xcb,
	0xe5, 0x7c, 0x6a, 0xe2, 0x5c, 0xc1, 0x6e, 0xb0, 0xe6, 0xaf, 0x01, 0x2e, 0xe4, 0x26, 0xa9, 0x54,
	0x20, 0x75, 0xba, 0xfd, 0x62, 0x20, 0xbf, 0x67, 0xf8, 0x09, 0x40, 0xd3, 0x75, 0xbb, 0x99, 0x83,
	0x26, 0xba, 0x95, 0x08, 0xaa, 0x79, 0xd9, 0x0b, 0xe8, 0xd1, 0xad, 0x49, 0x73, 0x4c, 0xcf, 0xde,
	0x7a, 0x8b, 0xf9, 0x5b, 0xf0, 0x82, 0xaa, 0x88, 0x93, 0xaf, 0x95, 0x2c, 0x53, 0x93, 0xda, 0x37,
	0x0a, 0xaf, 0xdc, 0x61, 0x49, 0x37, 0x5f, 0x87, 0xd5, 0xaa, 0xd1, 0x79, 0xb5, 0x2e, 0xdb, 0x61,
	0x6b, 0xdd, 0x54, 0xe9, 0x46, 0xf7, 0xec, 0x4e, 0xd7, 0x66, 0x47, 0x9f, 0xe0, 0x69, 0x7b, 0x59,
	0xe6, 0xd5, 0xbc, 0x87, 0xcb, 0x9e, 0x2f, 0xce, 0x07, 0x8f, 0x78, 0x1f, 0xba, 0x8b, 0x59, 0xf4,
	0xed, 0x52, 0x9c, 0x0f, 0x3a, 0xfc, 0x31, 0xb8, 0x91, 0x18, 0x2f, 0xc2, 0xe0, 0x52, 0x44, 0x03,
	0x6b, 0x24, 0x60, 0xf0, 0xef, 0x8f, 0x88, 0xef, 0x43, 0x6f, 0x16, 0x9d, 0xcd, 0x04, 0x9a, 0xd0,
	0x8d, 0x7d, 0xe6, 0xc1, 0xd5, 0x09, 0x5a, 0xb1, 0x4f, 0x74, 0x1a, 0xd4, 0x46, 0x02, 0xcb, 0x69,
	0x0d, 0x6c, 0x72, 0x84, 0xa7, 0x51, 0x8d, 0x58, 0xbc, 0x67, 0xfe, 0x33, 0x1f, 0xff, 0x06, 0x00,
	0x00, 0xff, 0xff, 0xec, 0xad, 0x94, 0xb2, 0x46, 0x03, 0x00, 0x00,
}
