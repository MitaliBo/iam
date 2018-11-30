// Code generated by protoc-gen-go. DO NOT EDIT.
// source: iam/am/am.proto

package pbam // import "openpitrix.io/iam/pkg/pb/am"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Action struct {
	RoleName             []string `protobuf:"bytes,1,rep,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
	Xid                  []string `protobuf:"bytes,2,rep,name=xid,proto3" json:"xid,omitempty"`
	Verb                 string   `protobuf:"bytes,3,opt,name=verb,proto3" json:"verb,omitempty"`
	Path                 string   `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Action) Reset()         { *m = Action{} }
func (m *Action) String() string { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()    {}
func (*Action) Descriptor() ([]byte, []int) {
	return fileDescriptor_am_ca50cc6bb79f02db, []int{0}
}
func (m *Action) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Action.Unmarshal(m, b)
}
func (m *Action) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Action.Marshal(b, m, deterministic)
}
func (dst *Action) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Action.Merge(dst, src)
}
func (m *Action) XXX_Size() int {
	return xxx_messageInfo_Action.Size(m)
}
func (m *Action) XXX_DiscardUnknown() {
	xxx_messageInfo_Action.DiscardUnknown(m)
}

var xxx_messageInfo_Action proto.InternalMessageInfo

func (m *Action) GetRoleName() []string {
	if m != nil {
		return m.RoleName
	}
	return nil
}

func (m *Action) GetXid() []string {
	if m != nil {
		return m.Xid
	}
	return nil
}

func (m *Action) GetVerb() string {
	if m != nil {
		return m.Verb
	}
	return ""
}

func (m *Action) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type Rule struct {
	PathPattern          string   `protobuf:"bytes,1,opt,name=path_pattern,json=pathPattern,proto3" json:"path_pattern,omitempty"`
	VerbPattern          []string `protobuf:"bytes,2,rep,name=verb_pattern,json=verbPattern,proto3" json:"verb_pattern,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rule) Reset()         { *m = Rule{} }
func (m *Rule) String() string { return proto.CompactTextString(m) }
func (*Rule) ProtoMessage()    {}
func (*Rule) Descriptor() ([]byte, []int) {
	return fileDescriptor_am_ca50cc6bb79f02db, []int{1}
}
func (m *Rule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rule.Unmarshal(m, b)
}
func (m *Rule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rule.Marshal(b, m, deterministic)
}
func (dst *Rule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rule.Merge(dst, src)
}
func (m *Rule) XXX_Size() int {
	return xxx_messageInfo_Rule.Size(m)
}
func (m *Rule) XXX_DiscardUnknown() {
	xxx_messageInfo_Rule.DiscardUnknown(m)
}

var xxx_messageInfo_Rule proto.InternalMessageInfo

func (m *Rule) GetPathPattern() string {
	if m != nil {
		return m.PathPattern
	}
	return ""
}

func (m *Rule) GetVerbPattern() []string {
	if m != nil {
		return m.VerbPattern
	}
	return nil
}

type Role struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Rule                 []*Rule  `protobuf:"bytes,2,rep,name=rule,proto3" json:"rule,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_am_ca50cc6bb79f02db, []int{2}
}
func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (dst *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(dst, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Role) GetRule() []*Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

type RoleList struct {
	Value                []*Role  `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleList) Reset()         { *m = RoleList{} }
func (m *RoleList) String() string { return proto.CompactTextString(m) }
func (*RoleList) ProtoMessage()    {}
func (*RoleList) Descriptor() ([]byte, []int) {
	return fileDescriptor_am_ca50cc6bb79f02db, []int{3}
}
func (m *RoleList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleList.Unmarshal(m, b)
}
func (m *RoleList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleList.Marshal(b, m, deterministic)
}
func (dst *RoleList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleList.Merge(dst, src)
}
func (m *RoleList) XXX_Size() int {
	return xxx_messageInfo_RoleList.Size(m)
}
func (m *RoleList) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleList.DiscardUnknown(m)
}

var xxx_messageInfo_RoleList proto.InternalMessageInfo

func (m *RoleList) GetValue() []*Role {
	if m != nil {
		return m.Value
	}
	return nil
}

type RoleBinding struct {
	RoleName             string   `protobuf:"bytes,1,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
	Xid                  string   `protobuf:"bytes,2,opt,name=xid,proto3" json:"xid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleBinding) Reset()         { *m = RoleBinding{} }
func (m *RoleBinding) String() string { return proto.CompactTextString(m) }
func (*RoleBinding) ProtoMessage()    {}
func (*RoleBinding) Descriptor() ([]byte, []int) {
	return fileDescriptor_am_ca50cc6bb79f02db, []int{4}
}
func (m *RoleBinding) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleBinding.Unmarshal(m, b)
}
func (m *RoleBinding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleBinding.Marshal(b, m, deterministic)
}
func (dst *RoleBinding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleBinding.Merge(dst, src)
}
func (m *RoleBinding) XXX_Size() int {
	return xxx_messageInfo_RoleBinding.Size(m)
}
func (m *RoleBinding) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleBinding.DiscardUnknown(m)
}

var xxx_messageInfo_RoleBinding proto.InternalMessageInfo

func (m *RoleBinding) GetRoleName() string {
	if m != nil {
		return m.RoleName
	}
	return ""
}

func (m *RoleBinding) GetXid() string {
	if m != nil {
		return m.Xid
	}
	return ""
}

type RoleBindingList struct {
	Value                []*RoleBinding `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RoleBindingList) Reset()         { *m = RoleBindingList{} }
func (m *RoleBindingList) String() string { return proto.CompactTextString(m) }
func (*RoleBindingList) ProtoMessage()    {}
func (*RoleBindingList) Descriptor() ([]byte, []int) {
	return fileDescriptor_am_ca50cc6bb79f02db, []int{5}
}
func (m *RoleBindingList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleBindingList.Unmarshal(m, b)
}
func (m *RoleBindingList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleBindingList.Marshal(b, m, deterministic)
}
func (dst *RoleBindingList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleBindingList.Merge(dst, src)
}
func (m *RoleBindingList) XXX_Size() int {
	return xxx_messageInfo_RoleBindingList.Size(m)
}
func (m *RoleBindingList) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleBindingList.DiscardUnknown(m)
}

var xxx_messageInfo_RoleBindingList proto.InternalMessageInfo

func (m *RoleBindingList) GetValue() []*RoleBinding {
	if m != nil {
		return m.Value
	}
	return nil
}

type Bool struct {
	Value                bool     `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bool) Reset()         { *m = Bool{} }
func (m *Bool) String() string { return proto.CompactTextString(m) }
func (*Bool) ProtoMessage()    {}
func (*Bool) Descriptor() ([]byte, []int) {
	return fileDescriptor_am_ca50cc6bb79f02db, []int{6}
}
func (m *Bool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bool.Unmarshal(m, b)
}
func (m *Bool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bool.Marshal(b, m, deterministic)
}
func (dst *Bool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bool.Merge(dst, src)
}
func (m *Bool) XXX_Size() int {
	return xxx_messageInfo_Bool.Size(m)
}
func (m *Bool) XXX_DiscardUnknown() {
	xxx_messageInfo_Bool.DiscardUnknown(m)
}

var xxx_messageInfo_Bool proto.InternalMessageInfo

func (m *Bool) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

type String struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *String) Reset()         { *m = String{} }
func (m *String) String() string { return proto.CompactTextString(m) }
func (*String) ProtoMessage()    {}
func (*String) Descriptor() ([]byte, []int) {
	return fileDescriptor_am_ca50cc6bb79f02db, []int{7}
}
func (m *String) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_String.Unmarshal(m, b)
}
func (m *String) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_String.Marshal(b, m, deterministic)
}
func (dst *String) XXX_Merge(src proto.Message) {
	xxx_messageInfo_String.Merge(dst, src)
}
func (m *String) XXX_Size() int {
	return xxx_messageInfo_String.Size(m)
}
func (m *String) XXX_DiscardUnknown() {
	xxx_messageInfo_String.DiscardUnknown(m)
}

var xxx_messageInfo_String proto.InternalMessageInfo

func (m *String) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Xid struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Xid) Reset()         { *m = Xid{} }
func (m *Xid) String() string { return proto.CompactTextString(m) }
func (*Xid) ProtoMessage()    {}
func (*Xid) Descriptor() ([]byte, []int) {
	return fileDescriptor_am_ca50cc6bb79f02db, []int{8}
}
func (m *Xid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Xid.Unmarshal(m, b)
}
func (m *Xid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Xid.Marshal(b, m, deterministic)
}
func (dst *Xid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Xid.Merge(dst, src)
}
func (m *Xid) XXX_Size() int {
	return xxx_messageInfo_Xid.Size(m)
}
func (m *Xid) XXX_DiscardUnknown() {
	xxx_messageInfo_Xid.DiscardUnknown(m)
}

var xxx_messageInfo_Xid proto.InternalMessageInfo

func (m *Xid) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type XidList struct {
	Value                []string `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XidList) Reset()         { *m = XidList{} }
func (m *XidList) String() string { return proto.CompactTextString(m) }
func (*XidList) ProtoMessage()    {}
func (*XidList) Descriptor() ([]byte, []int) {
	return fileDescriptor_am_ca50cc6bb79f02db, []int{9}
}
func (m *XidList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XidList.Unmarshal(m, b)
}
func (m *XidList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XidList.Marshal(b, m, deterministic)
}
func (dst *XidList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XidList.Merge(dst, src)
}
func (m *XidList) XXX_Size() int {
	return xxx_messageInfo_XidList.Size(m)
}
func (m *XidList) XXX_DiscardUnknown() {
	xxx_messageInfo_XidList.DiscardUnknown(m)
}

var xxx_messageInfo_XidList proto.InternalMessageInfo

func (m *XidList) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

type RoleNameRegexp struct {
	// Use RE2 syntax
	// See https://github.com/google/re2/wiki/Syntax
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleNameRegexp) Reset()         { *m = RoleNameRegexp{} }
func (m *RoleNameRegexp) String() string { return proto.CompactTextString(m) }
func (*RoleNameRegexp) ProtoMessage()    {}
func (*RoleNameRegexp) Descriptor() ([]byte, []int) {
	return fileDescriptor_am_ca50cc6bb79f02db, []int{10}
}
func (m *RoleNameRegexp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleNameRegexp.Unmarshal(m, b)
}
func (m *RoleNameRegexp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleNameRegexp.Marshal(b, m, deterministic)
}
func (dst *RoleNameRegexp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleNameRegexp.Merge(dst, src)
}
func (m *RoleNameRegexp) XXX_Size() int {
	return xxx_messageInfo_RoleNameRegexp.Size(m)
}
func (m *RoleNameRegexp) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleNameRegexp.DiscardUnknown(m)
}

var xxx_messageInfo_RoleNameRegexp proto.InternalMessageInfo

func (m *RoleNameRegexp) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Action)(nil), "iam.am.Action")
	proto.RegisterType((*Rule)(nil), "iam.am.Rule")
	proto.RegisterType((*Role)(nil), "iam.am.Role")
	proto.RegisterType((*RoleList)(nil), "iam.am.RoleList")
	proto.RegisterType((*RoleBinding)(nil), "iam.am.RoleBinding")
	proto.RegisterType((*RoleBindingList)(nil), "iam.am.RoleBindingList")
	proto.RegisterType((*Bool)(nil), "iam.am.Bool")
	proto.RegisterType((*String)(nil), "iam.am.String")
	proto.RegisterType((*Xid)(nil), "iam.am.Xid")
	proto.RegisterType((*XidList)(nil), "iam.am.XidList")
	proto.RegisterType((*RoleNameRegexp)(nil), "iam.am.RoleNameRegexp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccessManagerClient is the client API for AccessManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccessManagerClient interface {
	CreateRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Role, error)
	ModifyRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Role, error)
	DeleteRoleByName(ctx context.Context, in *String, opts ...grpc.CallOption) (*Bool, error)
	GetRoleByRoleName(ctx context.Context, in *String, opts ...grpc.CallOption) (*Role, error)
	GetRoleByXidList(ctx context.Context, in *XidList, opts ...grpc.CallOption) (*RoleList, error)
	ListRoles(ctx context.Context, in *RoleNameRegexp, opts ...grpc.CallOption) (*RoleList, error)
	CreateRoleBinding(ctx context.Context, in *RoleBindingList, opts ...grpc.CallOption) (*Bool, error)
	DeleteRoleBinding(ctx context.Context, in *XidList, opts ...grpc.CallOption) (*Bool, error)
	GetRoleBindingByRoleName(ctx context.Context, in *String, opts ...grpc.CallOption) (*RoleBindingList, error)
	GetRoleBindingByXidList(ctx context.Context, in *XidList, opts ...grpc.CallOption) (*RoleBindingList, error)
	ListRoleBindings(ctx context.Context, in *RoleNameRegexp, opts ...grpc.CallOption) (*RoleBindingList, error)
	// check any of xid or role_name can do the action
	CanDo(ctx context.Context, in *Action, opts ...grpc.CallOption) (*Bool, error)
}

type accessManagerClient struct {
	cc *grpc.ClientConn
}

func NewAccessManagerClient(cc *grpc.ClientConn) AccessManagerClient {
	return &accessManagerClient{cc}
}

func (c *accessManagerClient) CreateRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, "/iam.am.AccessManager/CreateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessManagerClient) ModifyRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, "/iam.am.AccessManager/ModifyRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessManagerClient) DeleteRoleByName(ctx context.Context, in *String, opts ...grpc.CallOption) (*Bool, error) {
	out := new(Bool)
	err := c.cc.Invoke(ctx, "/iam.am.AccessManager/DeleteRoleByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessManagerClient) GetRoleByRoleName(ctx context.Context, in *String, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, "/iam.am.AccessManager/GetRoleByRoleName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessManagerClient) GetRoleByXidList(ctx context.Context, in *XidList, opts ...grpc.CallOption) (*RoleList, error) {
	out := new(RoleList)
	err := c.cc.Invoke(ctx, "/iam.am.AccessManager/GetRoleByXidList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessManagerClient) ListRoles(ctx context.Context, in *RoleNameRegexp, opts ...grpc.CallOption) (*RoleList, error) {
	out := new(RoleList)
	err := c.cc.Invoke(ctx, "/iam.am.AccessManager/ListRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessManagerClient) CreateRoleBinding(ctx context.Context, in *RoleBindingList, opts ...grpc.CallOption) (*Bool, error) {
	out := new(Bool)
	err := c.cc.Invoke(ctx, "/iam.am.AccessManager/CreateRoleBinding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessManagerClient) DeleteRoleBinding(ctx context.Context, in *XidList, opts ...grpc.CallOption) (*Bool, error) {
	out := new(Bool)
	err := c.cc.Invoke(ctx, "/iam.am.AccessManager/DeleteRoleBinding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessManagerClient) GetRoleBindingByRoleName(ctx context.Context, in *String, opts ...grpc.CallOption) (*RoleBindingList, error) {
	out := new(RoleBindingList)
	err := c.cc.Invoke(ctx, "/iam.am.AccessManager/GetRoleBindingByRoleName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessManagerClient) GetRoleBindingByXidList(ctx context.Context, in *XidList, opts ...grpc.CallOption) (*RoleBindingList, error) {
	out := new(RoleBindingList)
	err := c.cc.Invoke(ctx, "/iam.am.AccessManager/GetRoleBindingByXidList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessManagerClient) ListRoleBindings(ctx context.Context, in *RoleNameRegexp, opts ...grpc.CallOption) (*RoleBindingList, error) {
	out := new(RoleBindingList)
	err := c.cc.Invoke(ctx, "/iam.am.AccessManager/ListRoleBindings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessManagerClient) CanDo(ctx context.Context, in *Action, opts ...grpc.CallOption) (*Bool, error) {
	out := new(Bool)
	err := c.cc.Invoke(ctx, "/iam.am.AccessManager/CanDo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccessManagerServer is the server API for AccessManager service.
type AccessManagerServer interface {
	CreateRole(context.Context, *Role) (*Role, error)
	ModifyRole(context.Context, *Role) (*Role, error)
	DeleteRoleByName(context.Context, *String) (*Bool, error)
	GetRoleByRoleName(context.Context, *String) (*Role, error)
	GetRoleByXidList(context.Context, *XidList) (*RoleList, error)
	ListRoles(context.Context, *RoleNameRegexp) (*RoleList, error)
	CreateRoleBinding(context.Context, *RoleBindingList) (*Bool, error)
	DeleteRoleBinding(context.Context, *XidList) (*Bool, error)
	GetRoleBindingByRoleName(context.Context, *String) (*RoleBindingList, error)
	GetRoleBindingByXidList(context.Context, *XidList) (*RoleBindingList, error)
	ListRoleBindings(context.Context, *RoleNameRegexp) (*RoleBindingList, error)
	// check any of xid or role_name can do the action
	CanDo(context.Context, *Action) (*Bool, error)
}

func RegisterAccessManagerServer(s *grpc.Server, srv AccessManagerServer) {
	s.RegisterService(&_AccessManager_serviceDesc, srv)
}

func _AccessManager_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Role)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessManagerServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iam.am.AccessManager/CreateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessManagerServer).CreateRole(ctx, req.(*Role))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessManager_ModifyRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Role)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessManagerServer).ModifyRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iam.am.AccessManager/ModifyRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessManagerServer).ModifyRole(ctx, req.(*Role))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessManager_DeleteRoleByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessManagerServer).DeleteRoleByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iam.am.AccessManager/DeleteRoleByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessManagerServer).DeleteRoleByName(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessManager_GetRoleByRoleName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessManagerServer).GetRoleByRoleName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iam.am.AccessManager/GetRoleByRoleName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessManagerServer).GetRoleByRoleName(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessManager_GetRoleByXidList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(XidList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessManagerServer).GetRoleByXidList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iam.am.AccessManager/GetRoleByXidList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessManagerServer).GetRoleByXidList(ctx, req.(*XidList))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessManager_ListRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleNameRegexp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessManagerServer).ListRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iam.am.AccessManager/ListRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessManagerServer).ListRoles(ctx, req.(*RoleNameRegexp))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessManager_CreateRoleBinding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleBindingList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessManagerServer).CreateRoleBinding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iam.am.AccessManager/CreateRoleBinding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessManagerServer).CreateRoleBinding(ctx, req.(*RoleBindingList))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessManager_DeleteRoleBinding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(XidList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessManagerServer).DeleteRoleBinding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iam.am.AccessManager/DeleteRoleBinding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessManagerServer).DeleteRoleBinding(ctx, req.(*XidList))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessManager_GetRoleBindingByRoleName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessManagerServer).GetRoleBindingByRoleName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iam.am.AccessManager/GetRoleBindingByRoleName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessManagerServer).GetRoleBindingByRoleName(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessManager_GetRoleBindingByXidList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(XidList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessManagerServer).GetRoleBindingByXidList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iam.am.AccessManager/GetRoleBindingByXidList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessManagerServer).GetRoleBindingByXidList(ctx, req.(*XidList))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessManager_ListRoleBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleNameRegexp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessManagerServer).ListRoleBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iam.am.AccessManager/ListRoleBindings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessManagerServer).ListRoleBindings(ctx, req.(*RoleNameRegexp))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessManager_CanDo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Action)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessManagerServer).CanDo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iam.am.AccessManager/CanDo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessManagerServer).CanDo(ctx, req.(*Action))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccessManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "iam.am.AccessManager",
	HandlerType: (*AccessManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRole",
			Handler:    _AccessManager_CreateRole_Handler,
		},
		{
			MethodName: "ModifyRole",
			Handler:    _AccessManager_ModifyRole_Handler,
		},
		{
			MethodName: "DeleteRoleByName",
			Handler:    _AccessManager_DeleteRoleByName_Handler,
		},
		{
			MethodName: "GetRoleByRoleName",
			Handler:    _AccessManager_GetRoleByRoleName_Handler,
		},
		{
			MethodName: "GetRoleByXidList",
			Handler:    _AccessManager_GetRoleByXidList_Handler,
		},
		{
			MethodName: "ListRoles",
			Handler:    _AccessManager_ListRoles_Handler,
		},
		{
			MethodName: "CreateRoleBinding",
			Handler:    _AccessManager_CreateRoleBinding_Handler,
		},
		{
			MethodName: "DeleteRoleBinding",
			Handler:    _AccessManager_DeleteRoleBinding_Handler,
		},
		{
			MethodName: "GetRoleBindingByRoleName",
			Handler:    _AccessManager_GetRoleBindingByRoleName_Handler,
		},
		{
			MethodName: "GetRoleBindingByXidList",
			Handler:    _AccessManager_GetRoleBindingByXidList_Handler,
		},
		{
			MethodName: "ListRoleBindings",
			Handler:    _AccessManager_ListRoleBindings_Handler,
		},
		{
			MethodName: "CanDo",
			Handler:    _AccessManager_CanDo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iam/am/am.proto",
}

func init() { proto.RegisterFile("iam/am/am.proto", fileDescriptor_am_ca50cc6bb79f02db) }

var fileDescriptor_am_ca50cc6bb79f02db = []byte{
	// 534 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x51, 0x6f, 0xd3, 0x3c,
	0x14, 0x55, 0xbe, 0xa6, 0xfd, 0x9a, 0xdb, 0xb1, 0xa6, 0x66, 0x62, 0xd1, 0x86, 0xa0, 0x44, 0x02,
	0x95, 0x97, 0x14, 0x8a, 0xf6, 0x02, 0x7d, 0x59, 0x3b, 0x89, 0x97, 0x0d, 0xa1, 0xf0, 0x32, 0xf1,
	0x52, 0xb9, 0x8d, 0x29, 0x16, 0x49, 0x1c, 0xa5, 0xee, 0xd4, 0xfd, 0x17, 0x7e, 0x2c, 0xba, 0x76,
	0xdc, 0x24, 0xa3, 0x19, 0xd2, 0xa4, 0xd9, 0xe7, 0x9e, 0x73, 0xcf, 0xbd, 0xa7, 0x91, 0xa1, 0xcf,
	0x69, 0x32, 0x56, 0x7f, 0x41, 0x96, 0x0b, 0x29, 0x48, 0x87, 0xd3, 0x24, 0xa0, 0x89, 0xbf, 0x80,
	0xce, 0xe5, 0x4a, 0x72, 0x91, 0x92, 0x73, 0x70, 0x72, 0x11, 0xb3, 0x45, 0x4a, 0x13, 0xe6, 0x59,
	0xc3, 0xd6, 0xc8, 0x09, 0xbb, 0x08, 0x7c, 0xa1, 0x09, 0x23, 0x2e, 0xb4, 0x76, 0x3c, 0xf2, 0xfe,
	0x53, 0x30, 0x1e, 0x09, 0x01, 0xfb, 0x8e, 0xe5, 0x4b, 0xaf, 0x35, 0xb4, 0x46, 0x4e, 0xa8, 0xce,
	0x88, 0x65, 0x54, 0xfe, 0xf4, 0x6c, 0x8d, 0xe1, 0xd9, 0xbf, 0x06, 0x3b, 0xdc, 0xc6, 0x8c, 0xbc,
	0x82, 0x23, 0xbc, 0x2f, 0x32, 0x2a, 0x25, 0xcb, 0x53, 0xcf, 0x52, 0x9c, 0x1e, 0x62, 0x5f, 0x35,
	0x84, 0x14, 0x6c, 0xb3, 0xa7, 0x68, 0xb7, 0x1e, 0x62, 0x05, 0xc5, 0x9f, 0x82, 0x1d, 0x8a, 0x98,
	0xa1, 0x53, 0x31, 0xa7, 0x72, 0xc2, 0x33, 0x19, 0x82, 0x9d, 0x6f, 0x63, 0xa6, 0x64, 0xbd, 0xc9,
	0x51, 0xa0, 0x37, 0x0c, 0xd0, 0x3d, 0x54, 0x15, 0x3f, 0x80, 0x2e, 0xaa, 0xaf, 0xf9, 0x46, 0x12,
	0x1f, 0xda, 0x77, 0x34, 0xde, 0xea, 0x55, 0xab, 0x74, 0x11, 0xb3, 0x50, 0x97, 0xfc, 0x29, 0xf4,
	0xf0, 0x3a, 0xe3, 0x69, 0xc4, 0xd3, 0xf5, 0xc3, 0x84, 0xac, 0xc3, 0x09, 0x59, 0x45, 0x42, 0xfe,
	0x14, 0xfa, 0x15, 0xb5, 0x32, 0x7d, 0x5b, 0x37, 0x7d, 0x5a, 0x35, 0x2d, 0x78, 0xc6, 0xfb, 0x39,
	0xd8, 0x33, 0x21, 0x62, 0x72, 0x52, 0x4a, 0xac, 0x51, 0xd7, 0x54, 0x5f, 0x40, 0xe7, 0x9b, 0xcc,
	0x71, 0xa8, 0x5a, 0xdd, 0x31, 0xf5, 0x73, 0x68, 0xdd, 0xf2, 0xa8, 0xa1, 0xf8, 0x12, 0xfe, 0xbf,
	0xe5, 0x91, 0x1a, 0xe8, 0xa4, 0x3a, 0xd0, 0x9e, 0xf0, 0x06, 0x8e, 0xc3, 0x62, 0xaf, 0x90, 0xad,
	0xd9, 0x2e, 0x3b, 0xdc, 0x68, 0xf2, 0xbb, 0x0d, 0x4f, 0x2e, 0x57, 0x2b, 0xb6, 0xd9, 0xdc, 0xd0,
	0x94, 0xae, 0x59, 0x4e, 0x46, 0x00, 0xf3, 0x9c, 0x51, 0xc9, 0xd4, 0xaf, 0x54, 0x0b, 0xf5, 0xac,
	0x76, 0x43, 0xe6, 0x8d, 0x88, 0xf8, 0x8f, 0xfb, 0x7f, 0x32, 0xdf, 0x81, 0x7b, 0xc5, 0x62, 0xa6,
	0x7b, 0xce, 0xee, 0x55, 0xda, 0xc7, 0x86, 0xa1, 0x53, 0x28, 0x15, 0x2a, 0xb3, 0xf7, 0x30, 0xf8,
	0xcc, 0xa4, 0xa6, 0x9b, 0x45, 0x9a, 0x25, 0xca, 0xe4, 0x02, 0xdc, 0xbd, 0xc4, 0x84, 0xd3, 0x37,
	0x8c, 0x02, 0x38, 0x73, 0xab, 0x12, 0x45, 0xb9, 0x00, 0x07, 0xff, 0xe3, 0x7d, 0x43, 0x9e, 0x55,
	0xcb, 0x65, 0x78, 0x07, 0x64, 0x1f, 0x61, 0x50, 0xc6, 0x64, 0x3e, 0xaf, 0xd3, 0x03, 0x5f, 0x83,
	0xb2, 0xad, 0x2f, 0x37, 0x81, 0x41, 0x25, 0x8e, 0x42, 0xfb, 0xd7, 0xa8, 0x75, 0xcd, 0x1c, 0x3c,
	0xb3, 0x9d, 0x16, 0x3c, 0x92, 0x4b, 0xd3, 0x18, 0x64, 0x0e, 0xa7, 0x0f, 0x9b, 0x34, 0x26, 0xf5,
	0x48, 0x13, 0xd7, 0x04, 0x56, 0xc0, 0xcd, 0xb9, 0x35, 0x36, 0x79, 0x0d, 0xed, 0x39, 0x4d, 0xaf,
	0x44, 0x39, 0xbb, 0x7e, 0xc3, 0xea, 0x5b, 0xcf, 0xfc, 0xef, 0x43, 0x91, 0xb1, 0x34, 0xe3, 0x32,
	0xe7, 0xbb, 0x80, 0x8b, 0x31, 0x3e, 0x82, 0xd9, 0xaf, 0xf5, 0x38, 0x5b, 0x8e, 0x69, 0xf2, 0x29,
	0x5b, 0xd2, 0x64, 0xd9, 0x51, 0xcf, 0xe1, 0x87, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc1, 0xed,
	0x42, 0x87, 0x21, 0x05, 0x00, 0x00,
}
