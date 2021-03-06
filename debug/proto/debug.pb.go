// Code generated by protoc-gen-go. DO NOT EDIT.
// source: debug.proto

package debug

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SpanType int32

const (
	SpanType_INBOUND  SpanType = 0
	SpanType_OUTBOUND SpanType = 1
)

var SpanType_name = map[int32]string{
	0: "INBOUND",
	1: "OUTBOUND",
}

var SpanType_value = map[string]int32{
	"INBOUND":  0,
	"OUTBOUND": 1,
}

func (x SpanType) String() string {
	return proto.EnumName(SpanType_name, int32(x))
}

func (SpanType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8d9d361be58531fb, []int{0}
}

type HealthRequest struct {
	// optional service name
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthRequest) Reset()         { *m = HealthRequest{} }
func (m *HealthRequest) String() string { return proto.CompactTextString(m) }
func (*HealthRequest) ProtoMessage()    {}
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d9d361be58531fb, []int{0}
}

func (m *HealthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthRequest.Unmarshal(m, b)
}
func (m *HealthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthRequest.Marshal(b, m, deterministic)
}
func (m *HealthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthRequest.Merge(m, src)
}
func (m *HealthRequest) XXX_Size() int {
	return xxx_messageInfo_HealthRequest.Size(m)
}
func (m *HealthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthRequest proto.InternalMessageInfo

func (m *HealthRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

type HealthResponse struct {
	// default: ok
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthResponse) Reset()         { *m = HealthResponse{} }
func (m *HealthResponse) String() string { return proto.CompactTextString(m) }
func (*HealthResponse) ProtoMessage()    {}
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d9d361be58531fb, []int{1}
}

func (m *HealthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthResponse.Unmarshal(m, b)
}
func (m *HealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthResponse.Marshal(b, m, deterministic)
}
func (m *HealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthResponse.Merge(m, src)
}
func (m *HealthResponse) XXX_Size() int {
	return xxx_messageInfo_HealthResponse.Size(m)
}
func (m *HealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthResponse proto.InternalMessageInfo

func (m *HealthResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type StatsRequest struct {
	// optional service name
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatsRequest) Reset()         { *m = StatsRequest{} }
func (m *StatsRequest) String() string { return proto.CompactTextString(m) }
func (*StatsRequest) ProtoMessage()    {}
func (*StatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d9d361be58531fb, []int{2}
}

func (m *StatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsRequest.Unmarshal(m, b)
}
func (m *StatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsRequest.Marshal(b, m, deterministic)
}
func (m *StatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsRequest.Merge(m, src)
}
func (m *StatsRequest) XXX_Size() int {
	return xxx_messageInfo_StatsRequest.Size(m)
}
func (m *StatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatsRequest proto.InternalMessageInfo

func (m *StatsRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

type StatsResponse struct {
	// timestamp of recording
	Timestamp uint64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// unix timestamp
	Started uint64 `protobuf:"varint,2,opt,name=started,proto3" json:"started,omitempty"`
	// in seconds
	Uptime uint64 `protobuf:"varint,3,opt,name=uptime,proto3" json:"uptime,omitempty"`
	// in bytes
	Memory uint64 `protobuf:"varint,4,opt,name=memory,proto3" json:"memory,omitempty"`
	// num threads
	Threads uint64 `protobuf:"varint,5,opt,name=threads,proto3" json:"threads,omitempty"`
	// total gc in nanoseconds
	Gc uint64 `protobuf:"varint,6,opt,name=gc,proto3" json:"gc,omitempty"`
	// total number of requests
	Requests uint64 `protobuf:"varint,7,opt,name=requests,proto3" json:"requests,omitempty"`
	// total number of errors
	Errors               uint64   `protobuf:"varint,8,opt,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatsResponse) Reset()         { *m = StatsResponse{} }
func (m *StatsResponse) String() string { return proto.CompactTextString(m) }
func (*StatsResponse) ProtoMessage()    {}
func (*StatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d9d361be58531fb, []int{3}
}

func (m *StatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsResponse.Unmarshal(m, b)
}
func (m *StatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsResponse.Marshal(b, m, deterministic)
}
func (m *StatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsResponse.Merge(m, src)
}
func (m *StatsResponse) XXX_Size() int {
	return xxx_messageInfo_StatsResponse.Size(m)
}
func (m *StatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatsResponse proto.InternalMessageInfo

func (m *StatsResponse) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *StatsResponse) GetStarted() uint64 {
	if m != nil {
		return m.Started
	}
	return 0
}

func (m *StatsResponse) GetUptime() uint64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *StatsResponse) GetMemory() uint64 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *StatsResponse) GetThreads() uint64 {
	if m != nil {
		return m.Threads
	}
	return 0
}

func (m *StatsResponse) GetGc() uint64 {
	if m != nil {
		return m.Gc
	}
	return 0
}

func (m *StatsResponse) GetRequests() uint64 {
	if m != nil {
		return m.Requests
	}
	return 0
}

func (m *StatsResponse) GetErrors() uint64 {
	if m != nil {
		return m.Errors
	}
	return 0
}

// LogRequest requests service logs
type LogRequest struct {
	// service to request logs for
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// stream records continuously
	Stream bool `protobuf:"varint,2,opt,name=stream,proto3" json:"stream,omitempty"`
	// count of records to request
	Count int64 `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	// relative time in seconds
	// before the current time
	// from which to show logs
	Since                int64    `protobuf:"varint,4,opt,name=since,proto3" json:"since,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogRequest) Reset()         { *m = LogRequest{} }
func (m *LogRequest) String() string { return proto.CompactTextString(m) }
func (*LogRequest) ProtoMessage()    {}
func (*LogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d9d361be58531fb, []int{4}
}

func (m *LogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogRequest.Unmarshal(m, b)
}
func (m *LogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogRequest.Marshal(b, m, deterministic)
}
func (m *LogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogRequest.Merge(m, src)
}
func (m *LogRequest) XXX_Size() int {
	return xxx_messageInfo_LogRequest.Size(m)
}
func (m *LogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogRequest proto.InternalMessageInfo

func (m *LogRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *LogRequest) GetStream() bool {
	if m != nil {
		return m.Stream
	}
	return false
}

func (m *LogRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *LogRequest) GetSince() int64 {
	if m != nil {
		return m.Since
	}
	return 0
}

// Record is service log record
type Record struct {
	// timestamp of log record
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// record metadata
	Metadata map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// message
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d9d361be58531fb, []int{5}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Record) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Record) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type TraceRequest struct {
	// trace id to retrieve
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TraceRequest) Reset()         { *m = TraceRequest{} }
func (m *TraceRequest) String() string { return proto.CompactTextString(m) }
func (*TraceRequest) ProtoMessage()    {}
func (*TraceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d9d361be58531fb, []int{6}
}

func (m *TraceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TraceRequest.Unmarshal(m, b)
}
func (m *TraceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TraceRequest.Marshal(b, m, deterministic)
}
func (m *TraceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceRequest.Merge(m, src)
}
func (m *TraceRequest) XXX_Size() int {
	return xxx_messageInfo_TraceRequest.Size(m)
}
func (m *TraceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TraceRequest proto.InternalMessageInfo

func (m *TraceRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type TraceResponse struct {
	Spans                []*Span  `protobuf:"bytes,1,rep,name=spans,proto3" json:"spans,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TraceResponse) Reset()         { *m = TraceResponse{} }
func (m *TraceResponse) String() string { return proto.CompactTextString(m) }
func (*TraceResponse) ProtoMessage()    {}
func (*TraceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d9d361be58531fb, []int{7}
}

func (m *TraceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TraceResponse.Unmarshal(m, b)
}
func (m *TraceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TraceResponse.Marshal(b, m, deterministic)
}
func (m *TraceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceResponse.Merge(m, src)
}
func (m *TraceResponse) XXX_Size() int {
	return xxx_messageInfo_TraceResponse.Size(m)
}
func (m *TraceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TraceResponse proto.InternalMessageInfo

func (m *TraceResponse) GetSpans() []*Span {
	if m != nil {
		return m.Spans
	}
	return nil
}

type Span struct {
	// the trace id
	Trace string `protobuf:"bytes,1,opt,name=trace,proto3" json:"trace,omitempty"`
	// id of the span
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// parent span
	Parent string `protobuf:"bytes,3,opt,name=parent,proto3" json:"parent,omitempty"`
	// name of the resource
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// time of start in nanoseconds
	Started uint64 `protobuf:"varint,5,opt,name=started,proto3" json:"started,omitempty"`
	// duration of the execution in nanoseconds
	Duration uint64 `protobuf:"varint,6,opt,name=duration,proto3" json:"duration,omitempty"`
	// associated metadata
	Metadata             map[string]string `protobuf:"bytes,7,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Type                 SpanType          `protobuf:"varint,8,opt,name=type,proto3,enum=SpanType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Span) Reset()         { *m = Span{} }
func (m *Span) String() string { return proto.CompactTextString(m) }
func (*Span) ProtoMessage()    {}
func (*Span) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d9d361be58531fb, []int{8}
}

func (m *Span) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Span.Unmarshal(m, b)
}
func (m *Span) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Span.Marshal(b, m, deterministic)
}
func (m *Span) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Span.Merge(m, src)
}
func (m *Span) XXX_Size() int {
	return xxx_messageInfo_Span.Size(m)
}
func (m *Span) XXX_DiscardUnknown() {
	xxx_messageInfo_Span.DiscardUnknown(m)
}

var xxx_messageInfo_Span proto.InternalMessageInfo

func (m *Span) GetTrace() string {
	if m != nil {
		return m.Trace
	}
	return ""
}

func (m *Span) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Span) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *Span) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Span) GetStarted() uint64 {
	if m != nil {
		return m.Started
	}
	return 0
}

func (m *Span) GetDuration() uint64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Span) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Span) GetType() SpanType {
	if m != nil {
		return m.Type
	}
	return SpanType_INBOUND
}

func init() {
	proto.RegisterEnum("SpanType", SpanType_name, SpanType_value)
	proto.RegisterType((*HealthRequest)(nil), "HealthRequest")
	proto.RegisterType((*HealthResponse)(nil), "HealthResponse")
	proto.RegisterType((*StatsRequest)(nil), "StatsRequest")
	proto.RegisterType((*StatsResponse)(nil), "StatsResponse")
	proto.RegisterType((*LogRequest)(nil), "LogRequest")
	proto.RegisterType((*Record)(nil), "Record")
	proto.RegisterMapType((map[string]string)(nil), "Record.MetadataEntry")
	proto.RegisterType((*TraceRequest)(nil), "TraceRequest")
	proto.RegisterType((*TraceResponse)(nil), "TraceResponse")
	proto.RegisterType((*Span)(nil), "Span")
	proto.RegisterMapType((map[string]string)(nil), "Span.MetadataEntry")
}

func init() { proto.RegisterFile("debug.proto", fileDescriptor_8d9d361be58531fb) }

var fileDescriptor_8d9d361be58531fb = []byte{
	// 586 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0xed, 0x38, 0xb1, 0x27, 0x8d, 0xa9, 0x96, 0x8b, 0x2c, 0x73, 0xab, 0x2c, 0x21, 0x85,
	0x8b, 0x0c, 0x94, 0x17, 0x04, 0x6f, 0xa8, 0x48, 0x20, 0x95, 0x56, 0xda, 0xb6, 0x1f, 0xb0, 0xb5,
	0x47, 0xae, 0xa1, 0xbe, 0xb0, 0xbb, 0xae, 0x94, 0x6f, 0xe1, 0x0b, 0x78, 0xe3, 0x67, 0xf8, 0x1f,
	0xb4, 0x17, 0xa7, 0xb1, 0x10, 0xea, 0x03, 0x6f, 0x3e, 0x67, 0x67, 0x4f, 0x66, 0x4e, 0xce, 0x0e,
	0x2c, 0x0a, 0x3c, 0xef, 0xcb, 0xac, 0xe3, 0xad, 0x6c, 0xd3, 0xa7, 0xb0, 0xfc, 0x84, 0xec, 0x52,
	0x5e, 0x50, 0xfc, 0xde, 0xa3, 0x90, 0x24, 0x86, 0xb9, 0x40, 0x7e, 0x55, 0xe5, 0x18, 0x3b, 0x7b,
	0xce, 0x2a, 0xa4, 0x03, 0x4c, 0x57, 0x10, 0x0d, 0xa5, 0xa2, 0x6b, 0x1b, 0x81, 0xe4, 0x1e, 0xcc,
	0x84, 0x64, 0xb2, 0x17, 0xb6, 0xd4, 0xa2, 0x74, 0x05, 0x3b, 0x27, 0x92, 0x49, 0x71, 0xb3, 0xe6,
	0x6f, 0x07, 0x96, 0xb6, 0xd4, 0x6a, 0x3e, 0x80, 0x50, 0x56, 0x35, 0x0a, 0xc9, 0xea, 0x4e, 0x57,
	0x4f, 0xe9, 0x35, 0xa1, 0x95, 0x24, 0xe3, 0x12, 0x8b, 0xd8, 0xd5, 0x67, 0x03, 0x54, 0xbd, 0xf4,
	0x9d, 0x2a, 0x8c, 0x3d, 0x7d, 0x60, 0x91, 0xe2, 0x6b, 0xac, 0x5b, 0xbe, 0x8e, 0xa7, 0x86, 0x37,
	0x48, 0x29, 0xc9, 0x0b, 0x8e, 0xac, 0x10, 0xb1, 0x6f, 0x94, 0x2c, 0x24, 0x11, 0xb8, 0x65, 0x1e,
	0xcf, 0x34, 0xe9, 0x96, 0x39, 0x49, 0x20, 0xe0, 0x66, 0x10, 0x11, 0xcf, 0x35, 0xbb, 0xc1, 0x4a,
	0x1d, 0x39, 0x6f, 0xb9, 0x88, 0x03, 0xa3, 0x6e, 0x50, 0xfa, 0x15, 0xe0, 0xb0, 0x2d, 0x6f, 0x9c,
	0xdf, 0x38, 0xc8, 0x91, 0xd5, 0x7a, 0x9c, 0x80, 0x5a, 0x44, 0xee, 0x80, 0x9f, 0xb7, 0x7d, 0x23,
	0xf5, 0x30, 0x1e, 0x35, 0x40, 0xb1, 0xa2, 0x6a, 0x72, 0xd4, 0xa3, 0x78, 0xd4, 0x80, 0xf4, 0x97,
	0x03, 0x33, 0x8a, 0x79, 0xcb, 0x8b, 0xbf, 0xcd, 0xf3, 0xb6, 0xcd, 0x7b, 0x0d, 0x41, 0x8d, 0x92,
	0x15, 0x4c, 0xb2, 0xd8, 0xdd, 0xf3, 0x56, 0x8b, 0xfd, 0xbb, 0x99, 0xb9, 0x98, 0x7d, 0xb1, 0xfc,
	0xc7, 0x46, 0xf2, 0x35, 0xdd, 0x94, 0xa9, 0xce, 0x6b, 0x14, 0x82, 0x95, 0xc6, 0xd6, 0x90, 0x0e,
	0x30, 0x79, 0x0f, 0xcb, 0xd1, 0x25, 0xb2, 0x0b, 0xde, 0x37, 0x5c, 0xdb, 0x01, 0xd5, 0xa7, 0x6a,
	0xf7, 0x8a, 0x5d, 0xf6, 0xa8, 0x67, 0x0b, 0xa9, 0x01, 0xef, 0xdc, 0xb7, 0x4e, 0xfa, 0x08, 0x76,
	0x4e, 0x39, 0xcb, 0x71, 0x30, 0x28, 0x02, 0xb7, 0x2a, 0xec, 0x55, 0xb7, 0x2a, 0xd2, 0x17, 0xb0,
	0xb4, 0xe7, 0x36, 0x15, 0xf7, 0xc1, 0x17, 0x1d, 0x6b, 0x54, 0xd0, 0x54, 0xdf, 0x7e, 0x76, 0xd2,
	0xb1, 0x86, 0x1a, 0x2e, 0xfd, 0xe1, 0xc2, 0x54, 0x61, 0xf5, 0x83, 0x52, 0x5d, 0xb3, 0x4a, 0x06,
	0x58, 0x71, 0x77, 0x10, 0x57, 0x9e, 0x77, 0x8c, 0xa3, 0x35, 0x37, 0xa4, 0x16, 0x11, 0x02, 0xd3,
	0x86, 0xd5, 0xc6, 0xdc, 0x90, 0xea, 0xef, 0xed, 0xbc, 0xf9, 0xe3, 0xbc, 0x25, 0x10, 0x14, 0x3d,
	0x67, 0xb2, 0x6a, 0x1b, 0x9b, 0x95, 0x0d, 0x26, 0x2f, 0xb7, 0x8c, 0x9e, 0xeb, 0x86, 0x6f, 0xeb,
	0x86, 0xff, 0x69, 0xf3, 0x43, 0x98, 0xca, 0x75, 0x87, 0x3a, 0x44, 0xd1, 0x7e, 0xa8, 0x8b, 0x4f,
	0xd7, 0x1d, 0x52, 0x4d, 0xff, 0x97, 0xd7, 0xcf, 0x9e, 0x40, 0x30, 0xc8, 0x91, 0x05, 0xcc, 0x3f,
	0x1f, 0x7d, 0x38, 0x3e, 0x3b, 0x3a, 0xd8, 0x9d, 0x90, 0x1d, 0x08, 0x8e, 0xcf, 0x4e, 0x0d, 0x72,
	0xf6, 0x7f, 0x3a, 0xe0, 0x1f, 0xa8, 0xc5, 0x40, 0x1e, 0x83, 0x77, 0xd8, 0x96, 0x64, 0x91, 0x5d,
	0x27, 0x38, 0x99, 0xdb, 0xa0, 0xa4, 0x93, 0x57, 0x0e, 0x79, 0x0e, 0x33, 0xb3, 0x08, 0x48, 0x94,
	0x8d, 0x96, 0x47, 0x72, 0x2b, 0x1b, 0x6f, 0x88, 0x74, 0x42, 0x56, 0xe0, 0xeb, 0x07, 0x4e, 0x96,
	0xd9, 0xf6, 0x4e, 0x48, 0xa2, 0x6c, 0xf4, 0xee, 0x4d, 0xa5, 0xfe, 0xd3, 0xc9, 0x32, 0xdb, 0x0e,
	0x47, 0x12, 0x65, 0xa3, 0x2c, 0xa4, 0x93, 0xf3, 0x99, 0xde, 0x5d, 0x6f, 0xfe, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x08, 0x25, 0x21, 0x57, 0xca, 0x04, 0x00, 0x00,
}
