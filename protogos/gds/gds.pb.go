// Code generated by protoc-gen-go.
// source: gds.proto
// DO NOT EDIT!

/*
Package gds is a generated protocol buffer package.

It is generated from these files:
	gds.proto

It has these top-level messages:
	Building
	BuildingGds
*/
package gds

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Building struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Level            *uint32 `protobuf:"varint,2,req,name=level" json:"level,omitempty"`
	Icon             *string `protobuf:"bytes,3,req,name=icon" json:"icon,omitempty"`
	Model            *string `protobuf:"bytes,4,req,name=model" json:"model,omitempty"`
	Texture          *string `protobuf:"bytes,5,req,name=texture" json:"texture,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Building) Reset()         { *m = Building{} }
func (m *Building) String() string { return proto.CompactTextString(m) }
func (*Building) ProtoMessage()    {}

func (m *Building) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Building) GetLevel() uint32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *Building) GetIcon() string {
	if m != nil && m.Icon != nil {
		return *m.Icon
	}
	return ""
}

func (m *Building) GetModel() string {
	if m != nil && m.Model != nil {
		return *m.Model
	}
	return ""
}

func (m *Building) GetTexture() string {
	if m != nil && m.Texture != nil {
		return *m.Texture
	}
	return ""
}

type BuildingGds struct {
	Buildings        []*Building `protobuf:"bytes,1,rep,name=buildings" json:"buildings,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *BuildingGds) Reset()         { *m = BuildingGds{} }
func (m *BuildingGds) String() string { return proto.CompactTextString(m) }
func (*BuildingGds) ProtoMessage()    {}

func (m *BuildingGds) GetBuildings() []*Building {
	if m != nil {
		return m.Buildings
	}
	return nil
}
