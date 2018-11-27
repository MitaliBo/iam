// Code generated by protoc-gen-go. DO NOT EDIT.
// source: iam/im/config.proto

package pbim // import "openpitrix.io/iam/pkg/pb/im"

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

type Config struct {
	LogLevel             *string           `protobuf:"bytes,1,opt,name=log_level,json=logLevel" json:"log_level,omitempty"`
	Mysql                *MysqlConfig      `protobuf:"bytes,2,opt,name=mysql" json:"mysql,omitempty"`
	Extra                map[string]string `protobuf:"bytes,100,rep,name=extra" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_f6b08286cdf8e798, []int{0}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (dst *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(dst, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetLogLevel() string {
	if m != nil && m.LogLevel != nil {
		return *m.LogLevel
	}
	return ""
}

func (m *Config) GetMysql() *MysqlConfig {
	if m != nil {
		return m.Mysql
	}
	return nil
}

func (m *Config) GetExtra() map[string]string {
	if m != nil {
		return m.Extra
	}
	return nil
}

type MysqlConfig struct {
	Host                 *string  `protobuf:"bytes,1,opt,name=host,def=openpitrix-db" json:"host,omitempty"`
	Port                 *int32   `protobuf:"varint,2,opt,name=port,def=3306" json:"port,omitempty"`
	User                 *string  `protobuf:"bytes,3,opt,name=user,def=root" json:"user,omitempty"`
	Password             *string  `protobuf:"bytes,4,opt,name=password,def=password" json:"password,omitempty"`
	Database             *string  `protobuf:"bytes,5,opt,name=database,def=openpitrix" json:"database,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MysqlConfig) Reset()         { *m = MysqlConfig{} }
func (m *MysqlConfig) String() string { return proto.CompactTextString(m) }
func (*MysqlConfig) ProtoMessage()    {}
func (*MysqlConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_f6b08286cdf8e798, []int{1}
}
func (m *MysqlConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MysqlConfig.Unmarshal(m, b)
}
func (m *MysqlConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MysqlConfig.Marshal(b, m, deterministic)
}
func (dst *MysqlConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MysqlConfig.Merge(dst, src)
}
func (m *MysqlConfig) XXX_Size() int {
	return xxx_messageInfo_MysqlConfig.Size(m)
}
func (m *MysqlConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MysqlConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MysqlConfig proto.InternalMessageInfo

const Default_MysqlConfig_Host string = "openpitrix-db"
const Default_MysqlConfig_Port int32 = 3306
const Default_MysqlConfig_User string = "root"
const Default_MysqlConfig_Password string = "password"
const Default_MysqlConfig_Database string = "openpitrix"

func (m *MysqlConfig) GetHost() string {
	if m != nil && m.Host != nil {
		return *m.Host
	}
	return Default_MysqlConfig_Host
}

func (m *MysqlConfig) GetPort() int32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return Default_MysqlConfig_Port
}

func (m *MysqlConfig) GetUser() string {
	if m != nil && m.User != nil {
		return *m.User
	}
	return Default_MysqlConfig_User
}

func (m *MysqlConfig) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return Default_MysqlConfig_Password
}

func (m *MysqlConfig) GetDatabase() string {
	if m != nil && m.Database != nil {
		return *m.Database
	}
	return Default_MysqlConfig_Database
}

func init() {
	proto.RegisterType((*Config)(nil), "iam.im.Config")
	proto.RegisterMapType((map[string]string)(nil), "iam.im.Config.ExtraEntry")
	proto.RegisterType((*MysqlConfig)(nil), "iam.im.MysqlConfig")
}

func init() { proto.RegisterFile("iam/im/config.proto", fileDescriptor_config_f6b08286cdf8e798) }

var fileDescriptor_config_f6b08286cdf8e798 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0x31, 0x4f, 0xfb, 0x30,
	0x10, 0xc5, 0x95, 0x36, 0xa9, 0xda, 0xab, 0xfe, 0xd2, 0x5f, 0x2e, 0x83, 0x81, 0x25, 0x54, 0x08,
	0x95, 0x81, 0x04, 0xb5, 0x12, 0xaa, 0xca, 0x06, 0xea, 0x06, 0x8b, 0x47, 0x16, 0xe4, 0x10, 0x13,
	0xac, 0xc6, 0x3d, 0xe3, 0xb8, 0xa5, 0xfd, 0x5a, 0xac, 0x7c, 0x39, 0x64, 0x1b, 0x1a, 0xb6, 0xbb,
	0xf7, 0x7b, 0xef, 0xfc, 0x24, 0xc3, 0x48, 0x72, 0x95, 0x4b, 0x95, 0xbf, 0xe0, 0xfa, 0x55, 0x56,
	0x99, 0x36, 0x68, 0x91, 0xf4, 0x24, 0x57, 0x99, 0x54, 0xe3, 0xaf, 0x08, 0x7a, 0xf7, 0x1e, 0x90,
	0x53, 0x18, 0xd4, 0x58, 0x3d, 0xd7, 0x62, 0x2b, 0x6a, 0x1a, 0xa5, 0xd1, 0x64, 0xc0, 0xfa, 0x35,
	0x56, 0x0f, 0x6e, 0x27, 0x97, 0x90, 0xa8, 0x7d, 0xf3, 0x5e, 0xd3, 0x4e, 0x1a, 0x4d, 0x86, 0xd3,
	0x51, 0x16, 0xf2, 0xd9, 0xa3, 0x13, 0xc3, 0x01, 0x16, 0x1c, 0x24, 0x87, 0x44, 0xec, 0xac, 0xe1,
	0xb4, 0x4c, 0xbb, 0x93, 0xe1, 0xf4, 0xf8, 0xd7, 0x1a, 0x5c, 0xd9, 0xd2, 0xb1, 0xe5, 0xda, 0x9a,
	0x3d, 0x0b, 0xbe, 0x93, 0x39, 0x40, 0x2b, 0x92, 0xff, 0xd0, 0x5d, 0x89, 0xfd, 0x4f, 0x01, 0x37,
	0x92, 0x23, 0x48, 0xb6, 0xbc, 0xde, 0x08, 0xff, 0xf6, 0x80, 0x85, 0x65, 0xd1, 0x99, 0x47, 0xe3,
	0xcf, 0x08, 0x86, 0x7f, 0x1a, 0x90, 0x33, 0x88, 0xdf, 0xb0, 0xb1, 0x21, 0xbc, 0xf8, 0x87, 0x5a,
	0xac, 0xb5, 0xb4, 0x46, 0xee, 0xae, 0xca, 0x82, 0x79, 0x44, 0x28, 0xc4, 0x1a, 0x8d, 0xf5, 0xb7,
	0x92, 0x45, 0x3c, 0x9b, 0x5d, 0xdf, 0x30, 0xaf, 0x38, 0xb2, 0x69, 0x84, 0xa1, 0x5d, 0x1f, 0x8e,
	0x0d, 0xa2, 0x65, 0x5e, 0x21, 0xe7, 0xd0, 0xd7, 0xbc, 0x69, 0x3e, 0xd0, 0x94, 0x34, 0xf6, 0xf4,
	0xb0, 0xb3, 0xc3, 0x44, 0x2e, 0xa0, 0x5f, 0x72, 0xcb, 0x0b, 0xde, 0x08, 0x9a, 0x78, 0x17, 0xb4,
	0x05, 0xd8, 0x81, 0xdd, 0x8d, 0x9f, 0xd2, 0x56, 0xcf, 0x24, 0xe6, 0xee, 0x7f, 0xf4, 0xaa, 0xca,
	0x75, 0x91, 0x4b, 0x75, 0xab, 0x0b, 0xa9, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xab, 0x1f, 0x4f,
	0x5a, 0xb4, 0x01, 0x00, 0x00,
}