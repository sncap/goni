// Code generated by protoc-gen-go.
// source: metric.proto
// DO NOT EDIT!

/*
Package metric is a generated protocol buffer package.

It is generated from these files:
	metric.proto

It has these top-level messages:
	ApplicationMetric
	SystemMetric
	Metric
*/
package metric

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ApplicationMetric struct {
	Error       []*ApplicationMetric_Error     `protobuf:"bytes,1,rep,name=error" json:"error,omitempty"`
	Transaction *ApplicationMetric_Transaction `protobuf:"bytes,2,opt,name=transaction" json:"transaction,omitempty"`
	User        []*ApplicationMetric_User      `protobuf:"bytes,3,rep,name=user" json:"user,omitempty"`
	Realtime    []*ApplicationMetric_Realtime  `protobuf:"bytes,4,rep,name=realtime" json:"realtime,omitempty"`
}

func (m *ApplicationMetric) Reset()                    { *m = ApplicationMetric{} }
func (m *ApplicationMetric) String() string            { return proto.CompactTextString(m) }
func (*ApplicationMetric) ProtoMessage()               {}
func (*ApplicationMetric) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ApplicationMetric) GetError() []*ApplicationMetric_Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ApplicationMetric) GetTransaction() *ApplicationMetric_Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *ApplicationMetric) GetUser() []*ApplicationMetric_User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *ApplicationMetric) GetRealtime() []*ApplicationMetric_Realtime {
	if m != nil {
		return m.Realtime
	}
	return nil
}

type ApplicationMetric_Breadcrumb struct {
	Tag  []string `protobuf:"bytes,1,rep,name=tag" json:"tag,omitempty"`
	TagT []int64  `protobuf:"varint,2,rep,name=tagT" json:"tagT,omitempty"`
}

func (m *ApplicationMetric_Breadcrumb) Reset()                    { *m = ApplicationMetric_Breadcrumb{} }
func (m *ApplicationMetric_Breadcrumb) String() string            { return proto.CompactTextString(m) }
func (*ApplicationMetric_Breadcrumb) ProtoMessage()               {}
func (*ApplicationMetric_Breadcrumb) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type ApplicationMetric_Error struct {
	Tag       string `protobuf:"bytes,1,opt,name=tag" json:"tag,omitempty"`
	Log       string `protobuf:"bytes,2,opt,name=log" json:"log,omitempty"`
	Timestamp string `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *ApplicationMetric_Error) Reset()                    { *m = ApplicationMetric_Error{} }
func (m *ApplicationMetric_Error) String() string            { return proto.CompactTextString(m) }
func (*ApplicationMetric_Error) ProtoMessage()               {}
func (*ApplicationMetric_Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

type ApplicationMetric_Transaction struct {
	Detail []*ApplicationMetric_TransactionDetail `protobuf:"bytes,1,rep,name=detail" json:"detail,omitempty"`
}

func (m *ApplicationMetric_Transaction) Reset()         { *m = ApplicationMetric_Transaction{} }
func (m *ApplicationMetric_Transaction) String() string { return proto.CompactTextString(m) }
func (*ApplicationMetric_Transaction) ProtoMessage()    {}
func (*ApplicationMetric_Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 2}
}

func (m *ApplicationMetric_Transaction) GetDetail() []*ApplicationMetric_TransactionDetail {
	if m != nil {
		return m.Detail
	}
	return nil
}

type ApplicationMetric_TransactionDetail struct {
	Path     string                               `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Data     []*ApplicationMetric_TransactionData `protobuf:"bytes,2,rep,name=data" json:"data,omitempty"`
	Method   string                               `protobuf:"bytes,3,opt,name=method" json:"method,omitempty"`
	Realpath string                               `protobuf:"bytes,4,opt,name=realpath" json:"realpath,omitempty"`
}

func (m *ApplicationMetric_TransactionDetail) Reset()         { *m = ApplicationMetric_TransactionDetail{} }
func (m *ApplicationMetric_TransactionDetail) String() string { return proto.CompactTextString(m) }
func (*ApplicationMetric_TransactionDetail) ProtoMessage()    {}
func (*ApplicationMetric_TransactionDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 3}
}

func (m *ApplicationMetric_TransactionDetail) GetData() []*ApplicationMetric_TransactionData {
	if m != nil {
		return m.Data
	}
	return nil
}

type ApplicationMetric_TransactionData struct {
	Status     *ApplicationMetric_TransactionStatus `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Browser    string                               `protobuf:"bytes,2,opt,name=browser" json:"browser,omitempty"`
	Breadcrumb *ApplicationMetric_Breadcrumb        `protobuf:"bytes,3,opt,name=breadcrumb" json:"breadcrumb,omitempty"`
}

func (m *ApplicationMetric_TransactionData) Reset()         { *m = ApplicationMetric_TransactionData{} }
func (m *ApplicationMetric_TransactionData) String() string { return proto.CompactTextString(m) }
func (*ApplicationMetric_TransactionData) ProtoMessage()    {}
func (*ApplicationMetric_TransactionData) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 4}
}

func (m *ApplicationMetric_TransactionData) GetStatus() *ApplicationMetric_TransactionStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ApplicationMetric_TransactionData) GetBreadcrumb() *ApplicationMetric_Breadcrumb {
	if m != nil {
		return m.Breadcrumb
	}
	return nil
}

type ApplicationMetric_TransactionStatus struct {
	Status    string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Duration  int64  `protobuf:"varint,2,opt,name=duration" json:"duration,omitempty"`
	Panic     bool   `protobuf:"varint,3,opt,name=panic" json:"panic,omitempty"`
	Timestamp string `protobuf:"bytes,4,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *ApplicationMetric_TransactionStatus) Reset()         { *m = ApplicationMetric_TransactionStatus{} }
func (m *ApplicationMetric_TransactionStatus) String() string { return proto.CompactTextString(m) }
func (*ApplicationMetric_TransactionStatus) ProtoMessage()    {}
func (*ApplicationMetric_TransactionStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 5}
}

type ApplicationMetric_User struct {
	Ip string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
}

func (m *ApplicationMetric_User) Reset()                    { *m = ApplicationMetric_User{} }
func (m *ApplicationMetric_User) String() string            { return proto.CompactTextString(m) }
func (*ApplicationMetric_User) ProtoMessage()               {}
func (*ApplicationMetric_User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 6} }

type ApplicationMetric_Realtime struct {
	Timegroup int64 `protobuf:"varint,1,opt,name=timegroup" json:"timegroup,omitempty"`
	Count     int64 `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
}

func (m *ApplicationMetric_Realtime) Reset()                    { *m = ApplicationMetric_Realtime{} }
func (m *ApplicationMetric_Realtime) String() string            { return proto.CompactTextString(m) }
func (*ApplicationMetric_Realtime) ProtoMessage()               {}
func (*ApplicationMetric_Realtime) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 7} }

type SystemMetric struct {
	Expvar   map[string]string      `protobuf:"bytes,1,rep,name=expvar" json:"expvar,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Resource *SystemMetric_Resource `protobuf:"bytes,2,opt,name=resource" json:"resource,omitempty"`
	Runtime  *SystemMetric_Runtime  `protobuf:"bytes,3,opt,name=runtime" json:"runtime,omitempty"`
}

func (m *SystemMetric) Reset()                    { *m = SystemMetric{} }
func (m *SystemMetric) String() string            { return proto.CompactTextString(m) }
func (*SystemMetric) ProtoMessage()               {}
func (*SystemMetric) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SystemMetric) GetExpvar() map[string]string {
	if m != nil {
		return m.Expvar
	}
	return nil
}

func (m *SystemMetric) GetResource() *SystemMetric_Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *SystemMetric) GetRuntime() *SystemMetric_Runtime {
	if m != nil {
		return m.Runtime
	}
	return nil
}

type SystemMetric_Resource struct {
	Cpu float64 `protobuf:"fixed64,1,opt,name=cpu" json:"cpu,omitempty"`
}

func (m *SystemMetric_Resource) Reset()                    { *m = SystemMetric_Resource{} }
func (m *SystemMetric_Resource) String() string            { return proto.CompactTextString(m) }
func (*SystemMetric_Resource) ProtoMessage()               {}
func (*SystemMetric_Resource) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type SystemMetric_Runtime struct {
	Cgo       int64 `protobuf:"varint,1,opt,name=cgo" json:"cgo,omitempty"`
	Goroutine int32 `protobuf:"varint,2,opt,name=goroutine" json:"goroutine,omitempty"`
}

func (m *SystemMetric_Runtime) Reset()                    { *m = SystemMetric_Runtime{} }
func (m *SystemMetric_Runtime) String() string            { return proto.CompactTextString(m) }
func (*SystemMetric_Runtime) ProtoMessage()               {}
func (*SystemMetric_Runtime) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 1} }

type Metric struct {
	Apikey      string             `protobuf:"bytes,1,opt,name=apikey" json:"apikey,omitempty"`
	Instance    string             `protobuf:"bytes,2,opt,name=instance" json:"instance,omitempty"`
	Timestamp   string             `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
	Application *ApplicationMetric `protobuf:"bytes,4,opt,name=application" json:"application,omitempty"`
	System      *SystemMetric      `protobuf:"bytes,5,opt,name=system" json:"system,omitempty"`
}

func (m *Metric) Reset()                    { *m = Metric{} }
func (m *Metric) String() string            { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()               {}
func (*Metric) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Metric) GetApplication() *ApplicationMetric {
	if m != nil {
		return m.Application
	}
	return nil
}

func (m *Metric) GetSystem() *SystemMetric {
	if m != nil {
		return m.System
	}
	return nil
}

func init() {
	proto.RegisterType((*ApplicationMetric)(nil), "ApplicationMetric")
	proto.RegisterType((*ApplicationMetric_Breadcrumb)(nil), "ApplicationMetric.Breadcrumb")
	proto.RegisterType((*ApplicationMetric_Error)(nil), "ApplicationMetric.Error")
	proto.RegisterType((*ApplicationMetric_Transaction)(nil), "ApplicationMetric.Transaction")
	proto.RegisterType((*ApplicationMetric_TransactionDetail)(nil), "ApplicationMetric.TransactionDetail")
	proto.RegisterType((*ApplicationMetric_TransactionData)(nil), "ApplicationMetric.TransactionData")
	proto.RegisterType((*ApplicationMetric_TransactionStatus)(nil), "ApplicationMetric.TransactionStatus")
	proto.RegisterType((*ApplicationMetric_User)(nil), "ApplicationMetric.User")
	proto.RegisterType((*ApplicationMetric_Realtime)(nil), "ApplicationMetric.Realtime")
	proto.RegisterType((*SystemMetric)(nil), "SystemMetric")
	proto.RegisterType((*SystemMetric_Resource)(nil), "SystemMetric.Resource")
	proto.RegisterType((*SystemMetric_Runtime)(nil), "SystemMetric.Runtime")
	proto.RegisterType((*Metric)(nil), "Metric")
}

func init() { proto.RegisterFile("metric.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 649 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x54, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x63, 0x27, 0x69, 0xc7, 0xe5, 0xd2, 0x15, 0x14, 0x63, 0x0a, 0xaa, 0x22, 0x90, 0x2a,
	0x21, 0x19, 0x61, 0x10, 0x37, 0x01, 0x02, 0x44, 0x1e, 0x10, 0xe2, 0x65, 0x5b, 0x3e, 0x60, 0xe3,
	0x58, 0xa9, 0x45, 0xe2, 0xb5, 0xd6, 0xeb, 0x42, 0xc4, 0x67, 0xf0, 0x13, 0xfc, 0x01, 0x12, 0xff,
	0xc0, 0x3f, 0xb1, 0x3b, 0xbb, 0xbe, 0xe4, 0xa2, 0xf6, 0x29, 0x33, 0x3b, 0xe7, 0x78, 0xe6, 0xcc,
	0x9e, 0x0d, 0xec, 0x2d, 0x52, 0x29, 0xb2, 0x24, 0x2a, 0x04, 0x97, 0x7c, 0xf4, 0x6f, 0x08, 0xfb,
	0xef, 0x8b, 0x62, 0x9e, 0x25, 0x4c, 0x66, 0x3c, 0xff, 0x82, 0x35, 0x12, 0x41, 0x3f, 0x15, 0x82,
	0x8b, 0xc0, 0x39, 0x72, 0x8f, 0xfd, 0x38, 0x88, 0x36, 0x20, 0xd1, 0x58, 0xd7, 0xa9, 0x81, 0x91,
	0x77, 0xe0, 0x4b, 0xc1, 0xf2, 0x92, 0x25, 0x1a, 0x11, 0xf4, 0x8e, 0x1c, 0xc5, 0xba, 0xb7, 0x85,
	0x75, 0xda, 0xa2, 0x68, 0x97, 0x42, 0x1e, 0x82, 0x57, 0x95, 0xa9, 0x08, 0x5c, 0x6c, 0x78, 0x6b,
	0x0b, 0xf5, 0xab, 0x2a, 0x53, 0x04, 0x91, 0xe7, 0xb0, 0x23, 0x52, 0x36, 0x97, 0xd9, 0x22, 0x0d,
	0x3c, 0x24, 0xdc, 0xd9, 0x42, 0xa0, 0x16, 0x42, 0x1b, 0x70, 0x18, 0x03, 0x7c, 0x50, 0xc9, 0x34,
	0x11, 0xd5, 0x62, 0x42, 0xae, 0x83, 0x2b, 0xd9, 0x0c, 0x35, 0xee, 0x52, 0x1d, 0x12, 0x02, 0x9e,
	0xfa, 0x39, 0x55, 0x02, 0xdc, 0x63, 0x97, 0x62, 0x1c, 0x7e, 0x82, 0x3e, 0x6a, 0x6d, 0xe1, 0x4e,
	0x0d, 0x57, 0x27, 0x73, 0x3e, 0x43, 0xb9, 0xea, 0x44, 0x85, 0xe4, 0x10, 0x76, 0x75, 0xa3, 0x52,
	0xb2, 0x45, 0xa1, 0xb4, 0xe8, 0xf3, 0xf6, 0x20, 0xfc, 0x0c, 0x7e, 0x67, 0x01, 0xe4, 0x35, 0x0c,
	0xa6, 0xa9, 0x64, 0xd9, 0xdc, 0xae, 0xf9, 0xfe, 0xc5, 0x0b, 0xfb, 0x88, 0x58, 0x6a, 0x39, 0xe1,
	0x2f, 0x07, 0xf6, 0x37, 0xaa, 0x5a, 0x41, 0xc1, 0xe4, 0x99, 0x9d, 0x12, 0x63, 0xf2, 0x0c, 0xbc,
	0x29, 0x93, 0x0c, 0x55, 0xf9, 0xf1, 0xe8, 0x92, 0x2e, 0x0a, 0x49, 0x11, 0x4f, 0x0e, 0x60, 0xa0,
	0xbc, 0x72, 0xc6, 0xa7, 0x56, 0x89, 0xcd, 0x48, 0x68, 0xd6, 0x8f, 0x7d, 0x3c, 0xac, 0x34, 0x79,
	0xf8, 0xdb, 0x81, 0x6b, 0x6b, 0x5f, 0xd3, 0x3a, 0x95, 0x7e, 0x59, 0x95, 0x38, 0xd5, 0xa5, 0x3a,
	0x4f, 0x10, 0x4b, 0x2d, 0x87, 0x04, 0x30, 0x9c, 0x08, 0xfe, 0x5d, 0x9b, 0xc3, 0x2c, 0xba, 0x4e,
	0xc9, 0x1b, 0x80, 0x49, 0x73, 0x9b, 0x38, 0xa3, 0x1f, 0xdf, 0xdd, 0xf2, 0xed, 0xf6, 0xca, 0x69,
	0x87, 0x10, 0xfe, 0x5c, 0xd9, 0x9f, 0xe9, 0xaa, 0x35, 0x77, 0x66, 0xdd, 0x6d, 0xa6, 0x50, 0x9a,
	0xa7, 0x95, 0x60, 0x8d, 0xbd, 0x5d, 0xda, 0xe4, 0xe4, 0x06, 0xf4, 0x0b, 0x96, 0x67, 0x09, 0x8e,
	0xb0, 0x43, 0x4d, 0xb2, 0x6a, 0x05, 0x6f, 0xdd, 0x0a, 0x07, 0xe0, 0x69, 0x43, 0x93, 0xab, 0xd0,
	0xcb, 0x0a, 0xdb, 0x4b, 0x45, 0xe1, 0x5b, 0xd8, 0xa9, 0x7d, 0x5b, 0x7f, 0x61, 0x26, 0x78, 0x65,
	0x20, 0x2e, 0x6d, 0x0f, 0x74, 0xd7, 0x84, 0x57, 0xb9, 0xb4, 0xe3, 0x98, 0x64, 0xf4, 0xb7, 0x07,
	0x7b, 0x27, 0xcb, 0x52, 0xa6, 0x0b, 0xfb, 0x94, 0x1f, 0xc3, 0x20, 0xfd, 0x51, 0x9c, 0xb3, 0xfa,
	0x2d, 0xdf, 0x8e, 0xba, 0xe5, 0x68, 0x8c, 0xb5, 0x71, 0x2e, 0xc5, 0x92, 0x5a, 0x20, 0x89, 0xf5,
	0xfd, 0x96, 0xbc, 0x12, 0x49, 0x6a, 0x9f, 0xf2, 0xc1, 0x2a, 0x89, 0xda, 0x2a, 0x6d, 0x70, 0xe4,
	0x11, 0x0c, 0x85, 0xea, 0xaf, 0x5f, 0xa4, 0xb9, 0x88, 0x9b, 0x6b, 0x14, 0x53, 0xa4, 0x35, 0x2a,
	0x3c, 0xd4, 0x42, 0x2d, 0x59, 0xbd, 0xa3, 0xa4, 0xa8, 0x50, 0xa2, 0x43, 0x75, 0x18, 0xbe, 0x84,
	0xa1, 0x65, 0x60, 0x71, 0xc6, 0xad, 0x7e, 0x1d, 0xea, 0xbd, 0xcc, 0xb8, 0xda, 0x81, 0xcc, 0x72,
	0x33, 0x60, 0x9f, 0xb6, 0x07, 0x8a, 0xea, 0x77, 0x44, 0x69, 0xfa, 0xb7, 0x74, 0x59, 0xbf, 0x5a,
	0x15, 0xea, 0xc5, 0x9d, 0xb3, 0x79, 0x95, 0x5a, 0x3b, 0x99, 0xe4, 0x55, 0xef, 0x85, 0x33, 0xfa,
	0xe3, 0xc0, 0xc0, 0xae, 0x4d, 0xf9, 0x80, 0x15, 0x59, 0xcb, 0xb4, 0x99, 0xf6, 0x41, 0x96, 0xab,
	0x2b, 0xcc, 0x93, 0x9a, 0xdf, 0xe4, 0x17, 0x3f, 0x7e, 0xf2, 0x14, 0x7c, 0xd6, 0x5a, 0x13, 0x1d,
	0xe1, 0xc7, 0x64, 0xd3, 0xae, 0xb4, 0x0b, 0x23, 0x0f, 0x94, 0x1f, 0x71, 0x8f, 0x41, 0x1f, 0x09,
	0x57, 0x56, 0xd6, 0x4a, 0x6d, 0x71, 0x32, 0xc0, 0x7f, 0xf3, 0x27, 0xff, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x65, 0xfa, 0xfd, 0x72, 0xdd, 0x05, 0x00, 0x00,
}
