// Code generated by protoc-gen-go. DO NOT EDIT.
// source: seer.proto

/*
Package seer is a generated protocol buffer package.

It is generated from these files:
	seer.proto

It has these top-level messages:
	Stream
	Event
	Interval
	Forecast
	CreateStreamRequest
	GetStreamRequest
	DeleteStreamRequest
	ListStreamsRequest
	ListStreamsResponse
	UpdateStreamRequest
	GetForecastRequest
*/
package seer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/empty"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Domain int32

const (
	Domain_CONTINUOUS          Domain = 0
	Domain_CONTINUOUS_RIGHT    Domain = 1
	Domain_CONTINUOUS_INTERVAL Domain = 2
	Domain_DISCRETE_RIGHT      Domain = 3
	Domain_DISCRETE_INTERVAL   Domain = 4
)

var Domain_name = map[int32]string{
	0: "CONTINUOUS",
	1: "CONTINUOUS_RIGHT",
	2: "CONTINUOUS_INTERVAL",
	3: "DISCRETE_RIGHT",
	4: "DISCRETE_INTERVAL",
}
var Domain_value = map[string]int32{
	"CONTINUOUS":          0,
	"CONTINUOUS_RIGHT":    1,
	"CONTINUOUS_INTERVAL": 2,
	"DISCRETE_RIGHT":      3,
	"DISCRETE_INTERVAL":   4,
}

func (x Domain) String() string {
	return proto.EnumName(Domain_name, int32(x))
}
func (Domain) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// A data stream
type Stream struct {
	Name          string                      `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Period        float64                     `protobuf:"fixed64,2,opt,name=period" json:"period,omitempty"`
	LastEventTime *google_protobuf1.Timestamp `protobuf:"bytes,3,opt,name=last_event_time,json=lastEventTime" json:"last_event_time,omitempty"`
	Domain        Domain                      `protobuf:"varint,4,opt,name=domain,enum=seer.Domain" json:"domain,omitempty"`
	Min           float64                     `protobuf:"fixed64,5,opt,name=min" json:"min,omitempty"`
	Max           float64                     `protobuf:"fixed64,6,opt,name=max" json:"max,omitempty"`
}

func (m *Stream) Reset()                    { *m = Stream{} }
func (m *Stream) String() string            { return proto.CompactTextString(m) }
func (*Stream) ProtoMessage()               {}
func (*Stream) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Stream) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Stream) GetPeriod() float64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *Stream) GetLastEventTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.LastEventTime
	}
	return nil
}

func (m *Stream) GetDomain() Domain {
	if m != nil {
		return m.Domain
	}
	return Domain_CONTINUOUS
}

func (m *Stream) GetMin() float64 {
	if m != nil {
		return m.Min
	}
	return 0
}

func (m *Stream) GetMax() float64 {
	if m != nil {
		return m.Max
	}
	return 0
}

// A set of ordered events (values and times) in a stream
type Event struct {
	Times  []*google_protobuf1.Timestamp `protobuf:"bytes,1,rep,name=times" json:"times,omitempty"`
	Values []float64                     `protobuf:"fixed64,2,rep,packed,name=values" json:"values,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Event) GetTimes() []*google_protobuf1.Timestamp {
	if m != nil {
		return m.Times
	}
	return nil
}

func (m *Event) GetValues() []float64 {
	if m != nil {
		return m.Values
	}
	return nil
}

// A confidence interval
type Interval struct {
	Probability float64   `protobuf:"fixed64,1,opt,name=probability" json:"probability,omitempty"`
	LowerBound  []float64 `protobuf:"fixed64,2,rep,packed,name=lower_bound,json=lowerBound" json:"lower_bound,omitempty"`
	UpperBound  []float64 `protobuf:"fixed64,3,rep,packed,name=upper_bound,json=upperBound" json:"upper_bound,omitempty"`
}

func (m *Interval) Reset()                    { *m = Interval{} }
func (m *Interval) String() string            { return proto.CompactTextString(m) }
func (*Interval) ProtoMessage()               {}
func (*Interval) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Interval) GetProbability() float64 {
	if m != nil {
		return m.Probability
	}
	return 0
}

func (m *Interval) GetLowerBound() []float64 {
	if m != nil {
		return m.LowerBound
	}
	return nil
}

func (m *Interval) GetUpperBound() []float64 {
	if m != nil {
		return m.UpperBound
	}
	return nil
}

// A forecast, with point predictions and confidence intervals
type Forecast struct {
	Times     []*google_protobuf1.Timestamp `protobuf:"bytes,1,rep,name=times" json:"times,omitempty"`
	Values    []float64                     `protobuf:"fixed64,2,rep,packed,name=values" json:"values,omitempty"`
	Intervals []*Interval                   `protobuf:"bytes,3,rep,name=intervals" json:"intervals,omitempty"`
}

func (m *Forecast) Reset()                    { *m = Forecast{} }
func (m *Forecast) String() string            { return proto.CompactTextString(m) }
func (*Forecast) ProtoMessage()               {}
func (*Forecast) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Forecast) GetTimes() []*google_protobuf1.Timestamp {
	if m != nil {
		return m.Times
	}
	return nil
}

func (m *Forecast) GetValues() []float64 {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *Forecast) GetIntervals() []*Interval {
	if m != nil {
		return m.Intervals
	}
	return nil
}

// The request message containing the stream to be created
type CreateStreamRequest struct {
	Stream *Stream `protobuf:"bytes,1,opt,name=stream" json:"stream,omitempty"`
}

func (m *CreateStreamRequest) Reset()                    { *m = CreateStreamRequest{} }
func (m *CreateStreamRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateStreamRequest) ProtoMessage()               {}
func (*CreateStreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CreateStreamRequest) GetStream() *Stream {
	if m != nil {
		return m.Stream
	}
	return nil
}

// The request message containing the name of the requested stream
type GetStreamRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetStreamRequest) Reset()                    { *m = GetStreamRequest{} }
func (m *GetStreamRequest) String() string            { return proto.CompactTextString(m) }
func (*GetStreamRequest) ProtoMessage()               {}
func (*GetStreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetStreamRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The request message containing the name of the stream to be deleted
type DeleteStreamRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteStreamRequest) Reset()                    { *m = DeleteStreamRequest{} }
func (m *DeleteStreamRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteStreamRequest) ProtoMessage()               {}
func (*DeleteStreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeleteStreamRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The request message containing the paging data for the stream to list
type ListStreamsRequest struct {
	PageSize   int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	PageNumber int32 `protobuf:"varint,2,opt,name=page_number,json=pageNumber" json:"page_number,omitempty"`
}

func (m *ListStreamsRequest) Reset()                    { *m = ListStreamsRequest{} }
func (m *ListStreamsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListStreamsRequest) ProtoMessage()               {}
func (*ListStreamsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ListStreamsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListStreamsRequest) GetPageNumber() int32 {
	if m != nil {
		return m.PageNumber
	}
	return 0
}

// The response message containing a list of streams
type ListStreamsResponse struct {
	Streams []*Stream `protobuf:"bytes,1,rep,name=streams" json:"streams,omitempty"`
}

func (m *ListStreamsResponse) Reset()                    { *m = ListStreamsResponse{} }
func (m *ListStreamsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListStreamsResponse) ProtoMessage()               {}
func (*ListStreamsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ListStreamsResponse) GetStreams() []*Stream {
	if m != nil {
		return m.Streams
	}
	return nil
}

// The request message containing events to apply to the stream
type UpdateStreamRequest struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Event *Event `protobuf:"bytes,2,opt,name=event" json:"event,omitempty"`
}

func (m *UpdateStreamRequest) Reset()                    { *m = UpdateStreamRequest{} }
func (m *UpdateStreamRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateStreamRequest) ProtoMessage()               {}
func (*UpdateStreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *UpdateStreamRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateStreamRequest) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

// The request message containing the forecast length
type GetForecastRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	N    int32  `protobuf:"varint,2,opt,name=n" json:"n,omitempty"`
}

func (m *GetForecastRequest) Reset()                    { *m = GetForecastRequest{} }
func (m *GetForecastRequest) String() string            { return proto.CompactTextString(m) }
func (*GetForecastRequest) ProtoMessage()               {}
func (*GetForecastRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *GetForecastRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetForecastRequest) GetN() int32 {
	if m != nil {
		return m.N
	}
	return 0
}

func init() {
	proto.RegisterType((*Stream)(nil), "seer.Stream")
	proto.RegisterType((*Event)(nil), "seer.Event")
	proto.RegisterType((*Interval)(nil), "seer.Interval")
	proto.RegisterType((*Forecast)(nil), "seer.Forecast")
	proto.RegisterType((*CreateStreamRequest)(nil), "seer.CreateStreamRequest")
	proto.RegisterType((*GetStreamRequest)(nil), "seer.GetStreamRequest")
	proto.RegisterType((*DeleteStreamRequest)(nil), "seer.DeleteStreamRequest")
	proto.RegisterType((*ListStreamsRequest)(nil), "seer.ListStreamsRequest")
	proto.RegisterType((*ListStreamsResponse)(nil), "seer.ListStreamsResponse")
	proto.RegisterType((*UpdateStreamRequest)(nil), "seer.UpdateStreamRequest")
	proto.RegisterType((*GetForecastRequest)(nil), "seer.GetForecastRequest")
	proto.RegisterEnum("seer.Domain", Domain_name, Domain_value)
}

func init() { proto.RegisterFile("seer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 670 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6f, 0xd3, 0x4e,
	0x10, 0x8d, 0xf3, 0xf5, 0x6b, 0xc6, 0xf9, 0x85, 0x30, 0x81, 0xe2, 0xba, 0x87, 0x1a, 0x0b, 0x55,
	0x01, 0xa1, 0x14, 0xa5, 0x12, 0x52, 0x85, 0x38, 0xd0, 0x24, 0x94, 0x48, 0x55, 0x2a, 0x36, 0x29,
	0xd7, 0x68, 0x43, 0x86, 0xca, 0x92, 0xbf, 0xb0, 0x9d, 0xd2, 0xf6, 0xc8, 0x9f, 0xc6, 0x3f, 0xc5,
	0x15, 0xed, 0xae, 0x9d, 0x3a, 0x1f, 0x02, 0x0e, 0xdc, 0xbc, 0x6f, 0xde, 0xcc, 0xec, 0x9b, 0x79,
	0x6b, 0x80, 0x98, 0x28, 0xea, 0x84, 0x51, 0x90, 0x04, 0x58, 0x16, 0xdf, 0xe6, 0xfe, 0x55, 0x10,
	0x5c, 0xb9, 0x74, 0x24, 0xb1, 0xd9, 0xe2, 0xcb, 0x11, 0x79, 0x61, 0x72, 0xab, 0x28, 0xe6, 0xc1,
	0x7a, 0x30, 0x71, 0x3c, 0x8a, 0x13, 0xee, 0x85, 0x8a, 0x60, 0xff, 0xd0, 0xa0, 0x3a, 0x4e, 0x22,
	0xe2, 0x1e, 0x22, 0x94, 0x7d, 0xee, 0x91, 0xa1, 0x59, 0x5a, 0xbb, 0xc6, 0xe4, 0x37, 0xee, 0x42,
	0x35, 0xa4, 0xc8, 0x09, 0xe6, 0x46, 0xd1, 0xd2, 0xda, 0x1a, 0x4b, 0x4f, 0x78, 0x0a, 0x0f, 0x5c,
	0x1e, 0x27, 0x53, 0xba, 0x26, 0x3f, 0x99, 0x8a, 0xa2, 0x46, 0xc9, 0xd2, 0xda, 0x7a, 0xd7, 0xec,
	0xa8, 0x8e, 0x9d, 0xac, 0x63, 0x67, 0x92, 0x75, 0x64, 0xff, 0x8b, 0x94, 0x81, 0xc8, 0x10, 0x18,
	0x3e, 0x83, 0xea, 0x3c, 0xf0, 0xb8, 0xe3, 0x1b, 0x65, 0x4b, 0x6b, 0x37, 0xba, 0xf5, 0x8e, 0xd4,
	0xd6, 0x97, 0x18, 0x4b, 0x63, 0xd8, 0x84, 0x92, 0xe7, 0xf8, 0x46, 0x45, 0xb6, 0x17, 0x9f, 0x12,
	0xe1, 0x37, 0x46, 0x35, 0x45, 0xf8, 0x8d, 0xfd, 0x11, 0x2a, 0xb2, 0x2c, 0xbe, 0x82, 0x8a, 0x14,
	0x68, 0x68, 0x56, 0xe9, 0x0f, 0x97, 0x51, 0x44, 0x21, 0xf0, 0x9a, 0xbb, 0x0b, 0x8a, 0x8d, 0xa2,
	0x55, 0x12, 0x02, 0xd5, 0xc9, 0xf6, 0x61, 0x67, 0xe8, 0x27, 0x14, 0x5d, 0x73, 0x17, 0x2d, 0xd0,
	0xc3, 0x28, 0x98, 0xf1, 0x99, 0xe3, 0x3a, 0xc9, 0xad, 0x9c, 0x8f, 0xc6, 0xf2, 0x10, 0x1e, 0x80,
	0xee, 0x06, 0xdf, 0x28, 0x9a, 0xce, 0x82, 0x85, 0x3f, 0x4f, 0x4b, 0x81, 0x84, 0x4e, 0x05, 0x22,
	0x08, 0x8b, 0x30, 0x5c, 0x12, 0x4a, 0x8a, 0x20, 0x21, 0x49, 0xb0, 0xbf, 0x6b, 0xb0, 0xf3, 0x3e,
	0x88, 0xe8, 0x33, 0x8f, 0xff, 0xa1, 0x0c, 0x7c, 0x09, 0x35, 0x27, 0x95, 0x11, 0xcb, 0xae, 0x7a,
	0xb7, 0xa1, 0xc6, 0x9c, 0xa9, 0x63, 0xf7, 0x04, 0xfb, 0x0d, 0xb4, 0x7a, 0x11, 0xf1, 0x84, 0x94,
	0x23, 0x18, 0x7d, 0x5d, 0x50, 0x9c, 0x88, 0x45, 0xc5, 0x12, 0x90, 0xd2, 0xf5, 0x6c, 0x51, 0x29,
	0x29, 0x8d, 0xd9, 0x87, 0xd0, 0x3c, 0xa3, 0x64, 0x35, 0x73, 0x8b, 0xa5, 0xec, 0xe7, 0xd0, 0xea,
	0x93, 0x4b, 0xeb, 0x4d, 0xb6, 0x51, 0x19, 0xe0, 0xb9, 0x13, 0xa7, 0x35, 0xe3, 0x8c, 0xb9, 0x0f,
	0xb5, 0x90, 0x5f, 0xd1, 0x34, 0x76, 0xee, 0x14, 0xbd, 0xc2, 0x76, 0x04, 0x30, 0x76, 0xee, 0x48,
	0x0c, 0x5a, 0x06, 0xfd, 0x85, 0x37, 0xa3, 0x48, 0xba, 0xb6, 0xc2, 0x40, 0x40, 0x23, 0x89, 0xd8,
	0x6f, 0xa1, 0xb5, 0x52, 0x33, 0x0e, 0x03, 0x3f, 0x26, 0x3c, 0x84, 0xff, 0x94, 0x8e, 0x6c, 0xe8,
	0xab, 0x22, 0xb3, 0xa0, 0x7d, 0x0e, 0xad, 0xcb, 0x70, 0xce, 0xff, 0xe2, 0xf6, 0xf8, 0x14, 0x2a,
	0xf2, 0x79, 0xc8, 0x4b, 0xe8, 0x5d, 0x5d, 0x15, 0x94, 0x46, 0x65, 0x2a, 0x62, 0xbf, 0x06, 0x3c,
	0xa3, 0x24, 0xdb, 0xfb, 0xef, 0x8a, 0xd5, 0x41, 0xf3, 0x53, 0x35, 0x9a, 0xff, 0x22, 0x82, 0xaa,
	0x7a, 0x26, 0xd8, 0x00, 0xe8, 0x5d, 0x8c, 0x26, 0xc3, 0xd1, 0xe5, 0xc5, 0xe5, 0xb8, 0x59, 0xc0,
	0x47, 0xd0, 0xbc, 0x3f, 0x4f, 0xd9, 0xf0, 0xec, 0xc3, 0xa4, 0xa9, 0xe1, 0x13, 0x68, 0xe5, 0xd0,
	0xe1, 0x68, 0x32, 0x60, 0x9f, 0xde, 0x9d, 0x37, 0x8b, 0x88, 0xd0, 0xe8, 0x0f, 0xc7, 0x3d, 0x36,
	0x98, 0x0c, 0x52, 0x72, 0x09, 0x1f, 0xc3, 0xc3, 0x25, 0xb6, 0xa4, 0x96, 0xbb, 0x3f, 0x8b, 0x50,
	0x1e, 0x13, 0x45, 0x78, 0x02, 0xf5, 0xbc, 0x4b, 0x70, 0x4f, 0x09, 0xdb, 0xe2, 0x1c, 0x73, 0x65,
	0x88, 0x76, 0x01, 0x8f, 0xa1, 0xb6, 0xf4, 0x08, 0xee, 0xaa, 0xe0, 0xba, 0x69, 0x36, 0x92, 0x4e,
	0xa0, 0x9e, 0x1f, 0x79, 0xd6, 0x6f, 0xcb, 0x1a, 0x36, 0x52, 0x7b, 0x50, 0xcf, 0x7b, 0x2d, 0x4b,
	0xdd, 0xe2, 0x3f, 0x73, 0x77, 0xe3, 0x91, 0x0d, 0xc4, 0x7f, 0xd4, 0x2e, 0x60, 0x1f, 0xf4, 0x9c,
	0x63, 0xd0, 0x50, 0x35, 0x36, 0x8d, 0x69, 0xee, 0x6d, 0x89, 0x28, 0x7b, 0x49, 0x15, 0x7a, 0x6e,
	0xd5, 0x59, 0x95, 0xcd, 0xed, 0x9b, 0xe9, 0xfb, 0xcc, 0x60, 0xbb, 0x30, 0xab, 0xca, 0x2b, 0x1d,
	0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x31, 0xcb, 0x5c, 0x49, 0xfc, 0x05, 0x00, 0x00,
}
