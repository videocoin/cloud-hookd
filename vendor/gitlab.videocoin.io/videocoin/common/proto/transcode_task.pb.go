// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transcode_task.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/golang/protobuf/ptypes/timestamp"

import time "time"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TranscodeStatus int32

const (
	TranscodeStatusNone        TranscodeStatus = 0
	TranscodeStatusPending     TranscodeStatus = 1
	TranscodeStatusTranscoding TranscodeStatus = 2
	TranscodeStatusCanceled    TranscodeStatus = 3
	TranscodeStatusFailed      TranscodeStatus = 4
	TranscodeStatusCompleted   TranscodeStatus = 5
)

var TranscodeStatus_name = map[int32]string{
	0: "none",
	1: "pending",
	2: "transcoding",
	3: "canceld",
	4: "failed",
	5: "completed",
}
var TranscodeStatus_value = map[string]int32{
	"none":        0,
	"pending":     1,
	"transcoding": 2,
	"canceld":     3,
	"failed":      4,
	"completed":   5,
}

func (x TranscodeStatus) String() string {
	return proto.EnumName(TranscodeStatus_name, int32(x))
}
func (TranscodeStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_transcode_task_2c89ecba2ce6acaf, []int{0}
}

type SimpleTranscodeTask struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	InputUrl             string     `protobuf:"bytes,2,opt,name=input_url,json=inputUrl,proto3" json:"input_url,omitempty"`
	OutputUrl            string     `protobuf:"bytes,3,opt,name=output_url,json=outputUrl,proto3" json:"output_url,omitempty"`
	Status               string     `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	UserId               int32      `protobuf:"varint,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CreatedAt            *time.Time `protobuf:"bytes,6,opt,name=created_at,json=createdAt,stdtime" json:"created_at,omitempty"`
	UpdatedAt            *time.Time `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,stdtime" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SimpleTranscodeTask) Reset()         { *m = SimpleTranscodeTask{} }
func (m *SimpleTranscodeTask) String() string { return proto.CompactTextString(m) }
func (*SimpleTranscodeTask) ProtoMessage()    {}
func (*SimpleTranscodeTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_transcode_task_2c89ecba2ce6acaf, []int{0}
}
func (m *SimpleTranscodeTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleTranscodeTask.Unmarshal(m, b)
}
func (m *SimpleTranscodeTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleTranscodeTask.Marshal(b, m, deterministic)
}
func (dst *SimpleTranscodeTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleTranscodeTask.Merge(dst, src)
}
func (m *SimpleTranscodeTask) XXX_Size() int {
	return xxx_messageInfo_SimpleTranscodeTask.Size(m)
}
func (m *SimpleTranscodeTask) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleTranscodeTask.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleTranscodeTask proto.InternalMessageInfo

func (m *SimpleTranscodeTask) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SimpleTranscodeTask) GetInputUrl() string {
	if m != nil {
		return m.InputUrl
	}
	return ""
}

func (m *SimpleTranscodeTask) GetOutputUrl() string {
	if m != nil {
		return m.OutputUrl
	}
	return ""
}

func (m *SimpleTranscodeTask) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *SimpleTranscodeTask) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *SimpleTranscodeTask) GetCreatedAt() *time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *SimpleTranscodeTask) GetUpdatedAt() *time.Time {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*SimpleTranscodeTask)(nil), "proto.SimpleTranscodeTask")
	proto.RegisterEnum("proto.TranscodeStatus", TranscodeStatus_name, TranscodeStatus_value)
}

func init() {
	proto.RegisterFile("transcode_task.proto", fileDescriptor_transcode_task_2c89ecba2ce6acaf)
}

var fileDescriptor_transcode_task_2c89ecba2ce6acaf = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0x93, 0xac, 0x4d, 0xc9, 0x37, 0x09, 0x22, 0x0f, 0xd6, 0xcc, 0x83, 0xcc, 0x20, 0x21,
	0x2a, 0x10, 0xa9, 0x04, 0x0f, 0x80, 0x36, 0x24, 0x04, 0x17, 0x84, 0xba, 0x72, 0xe1, 0x52, 0xb9,
	0xb1, 0x17, 0xac, 0xa5, 0x76, 0x94, 0xd8, 0xef, 0x80, 0x72, 0xe2, 0xc0, 0x81, 0x4b, 0x4e, 0xc0,
	0x4b, 0xf0, 0x04, 0x1c, 0x79, 0x03, 0x50, 0x79, 0x11, 0x14, 0xa7, 0x61, 0x52, 0x4e, 0x9c, 0x9c,
	0xff, 0xf7, 0xfb, 0xf9, 0x4b, 0x3e, 0x3b, 0x70, 0x53, 0x97, 0x54, 0x56, 0xa9, 0x62, 0x7c, 0xa5,
	0x69, 0x75, 0x99, 0x14, 0xa5, 0xd2, 0x0a, 0x8d, 0xed, 0x82, 0x4f, 0x32, 0xa5, 0xb2, 0x9c, 0xcf,
	0x6d, 0x5a, 0x9b, 0x8b, 0xb9, 0x16, 0x1b, 0x5e, 0x69, 0xba, 0x29, 0x3a, 0x0f, 0x3f, 0xce, 0x84,
	0x7e, 0x6f, 0xd6, 0x49, 0xaa, 0x36, 0xf3, 0x4c, 0x65, 0xea, 0xca, 0x6c, 0x93, 0x0d, 0xf6, 0xa9,
	0xd3, 0xef, 0x7d, 0xf2, 0xe0, 0xe0, 0x5c, 0x6c, 0x8a, 0x9c, 0x2f, 0xfb, 0xb7, 0x2e, 0x69, 0x75,
	0x89, 0xae, 0x83, 0x27, 0x58, 0xe4, 0x12, 0x77, 0x16, 0x2c, 0x3c, 0xc1, 0xd0, 0x31, 0x04, 0x42,
	0x16, 0x46, 0xaf, 0x4c, 0x99, 0x47, 0x9e, 0x2d, 0x5f, 0xb3, 0x85, 0xb7, 0x65, 0x8e, 0xee, 0x00,
	0x28, 0xa3, 0x7b, 0xba, 0x67, 0x69, 0xd0, 0x55, 0x5a, 0x7c, 0x08, 0x7e, 0xa5, 0xa9, 0x36, 0x55,
	0x34, 0xb2, 0x68, 0x97, 0xd0, 0x14, 0x26, 0xa6, 0xe2, 0xe5, 0x4a, 0xb0, 0x68, 0x4c, 0xdc, 0xd9,
	0x78, 0xe1, 0xb7, 0xf1, 0x15, 0x43, 0xcf, 0x00, 0xd2, 0x92, 0x53, 0xcd, 0xd9, 0x8a, 0xea, 0xc8,
	0x27, 0xee, 0x6c, 0xff, 0x09, 0x4e, 0xba, 0xc9, 0x93, 0x7e, 0x9e, 0x64, 0xd9, 0x4f, 0x7e, 0x36,
	0xfa, 0xf8, 0xeb, 0xc4, 0x5d, 0x04, 0xbb, 0x3d, 0xa7, 0xba, 0x6d, 0x60, 0x0a, 0xd6, 0x37, 0x98,
	0xfc, 0x6f, 0x83, 0xdd, 0x9e, 0x53, 0xfd, 0xf0, 0x9b, 0x07, 0x37, 0xfe, 0x1d, 0xc8, 0x79, 0xf7,
	0xb9, 0x77, 0x61, 0x24, 0x95, 0xe4, 0xa1, 0x83, 0xa7, 0x75, 0x43, 0x0e, 0x06, 0xf8, 0xb5, 0x92,
	0x1c, 0x3d, 0x80, 0x49, 0xc1, 0x25, 0x13, 0x32, 0x0b, 0x5d, 0x8c, 0xeb, 0x86, 0x1c, 0x0e, 0xac,
	0x37, 0x1d, 0x45, 0x73, 0xd8, 0xef, 0x6f, 0xb9, 0x95, 0x3d, 0x1c, 0xd7, 0x0d, 0xc1, 0x03, 0x79,
	0x79, 0x65, 0xa0, 0x19, 0x4c, 0x52, 0x2a, 0x53, 0x9e, 0xb3, 0x70, 0x0f, 0x1f, 0xd7, 0x0d, 0x99,
	0x0e, 0xe4, 0xe7, 0x96, 0x72, 0x86, 0xee, 0x83, 0x7f, 0x41, 0x45, 0xce, 0x59, 0x38, 0xc2, 0x47,
	0x75, 0x43, 0x6e, 0x0d, 0xc4, 0x17, 0x16, 0xa2, 0x47, 0x10, 0xa4, 0xaa, 0xbd, 0x77, 0xcd, 0x59,
	0x38, 0xc6, 0xb7, 0xeb, 0x86, 0x44, 0xc3, 0x96, 0x3d, 0xc7, 0xd3, 0x0f, 0x5f, 0x62, 0xe7, 0xfb,
	0xd7, 0x78, 0x78, 0x26, 0x67, 0x47, 0x3f, 0xb6, 0xb1, 0xf3, 0x73, 0x1b, 0x3b, 0xbf, 0xb7, 0xb1,
	0xf3, 0xf9, 0x4f, 0xec, 0xbc, 0x74, 0xdf, 0x75, 0x7f, 0xea, 0xda, 0xb7, 0xcb, 0xd3, 0xbf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x3b, 0x78, 0xda, 0x7d, 0xcf, 0x02, 0x00, 0x00,
}
