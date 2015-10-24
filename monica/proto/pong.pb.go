// Code generated by protoc-gen-go.
// source: pong.proto
// DO NOT EDIT!

/*
Package pong is a generated protocol buffer package.

It is generated from these files:
	pong.proto

It has these top-level messages:
	Pong
*/
package pong

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Pong struct {
	Header           *common.Header `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	Timestamp        *int64         `protobuf:"varint,2,req,name=timestamp" json:"timestamp,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}

func (m *Pong) GetHeader() *common.Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pong) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}