// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metrics.proto

package metrics

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

type Batch struct {
	Events               []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Batch) Reset()         { *m = Batch{} }
func (m *Batch) String() string { return proto.CompactTextString(m) }
func (*Batch) ProtoMessage()    {}
func (*Batch) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{0}
}

func (m *Batch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Batch.Unmarshal(m, b)
}
func (m *Batch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Batch.Marshal(b, m, deterministic)
}
func (m *Batch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Batch.Merge(m, src)
}
func (m *Batch) XXX_Size() int {
	return xxx_messageInfo_Batch.Size(m)
}
func (m *Batch) XXX_DiscardUnknown() {
	xxx_messageInfo_Batch.DiscardUnknown(m)
}

var xxx_messageInfo_Batch proto.InternalMessageInfo

func (m *Batch) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type Event struct {
	Signature  string `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	TargetId   string `protobuf:"bytes,2,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	IsProvider bool   `protobuf:"varint,3,opt,name=is_provider,json=isProvider,proto3" json:"is_provider,omitempty"`
	// Types that are valid to be assigned to Metric:
	//	*Event_VersionPayload
	//	*Event_SessionEventPayload
	//	*Event_SessionStatisticsPayload
	//	*Event_ProposalPayload
	//	*Event_SessionTokensPayload
	Metric               isEvent_Metric `protobuf_oneof:"metric"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{1}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *Event) GetTargetId() string {
	if m != nil {
		return m.TargetId
	}
	return ""
}

func (m *Event) GetIsProvider() bool {
	if m != nil {
		return m.IsProvider
	}
	return false
}

type isEvent_Metric interface {
	isEvent_Metric()
}

type Event_VersionPayload struct {
	VersionPayload *VersionPayload `protobuf:"bytes,4,opt,name=version_payload,json=versionPayload,proto3,oneof"`
}

type Event_SessionEventPayload struct {
	SessionEventPayload *SessionEventPayload `protobuf:"bytes,5,opt,name=session_event_payload,json=sessionEventPayload,proto3,oneof"`
}

type Event_SessionStatisticsPayload struct {
	SessionStatisticsPayload *SessionStatisticsPayload `protobuf:"bytes,6,opt,name=session_statistics_payload,json=sessionStatisticsPayload,proto3,oneof"`
}

type Event_ProposalPayload struct {
	ProposalPayload *ProposalPayload `protobuf:"bytes,7,opt,name=proposal_payload,json=proposalPayload,proto3,oneof"`
}

type Event_SessionTokensPayload struct {
	SessionTokensPayload *SessionTokensPayload `protobuf:"bytes,8,opt,name=session_tokens_payload,json=sessionTokensPayload,proto3,oneof"`
}

func (*Event_VersionPayload) isEvent_Metric() {}

func (*Event_SessionEventPayload) isEvent_Metric() {}

func (*Event_SessionStatisticsPayload) isEvent_Metric() {}

func (*Event_ProposalPayload) isEvent_Metric() {}

func (*Event_SessionTokensPayload) isEvent_Metric() {}

func (m *Event) GetMetric() isEvent_Metric {
	if m != nil {
		return m.Metric
	}
	return nil
}

func (m *Event) GetVersionPayload() *VersionPayload {
	if x, ok := m.GetMetric().(*Event_VersionPayload); ok {
		return x.VersionPayload
	}
	return nil
}

func (m *Event) GetSessionEventPayload() *SessionEventPayload {
	if x, ok := m.GetMetric().(*Event_SessionEventPayload); ok {
		return x.SessionEventPayload
	}
	return nil
}

func (m *Event) GetSessionStatisticsPayload() *SessionStatisticsPayload {
	if x, ok := m.GetMetric().(*Event_SessionStatisticsPayload); ok {
		return x.SessionStatisticsPayload
	}
	return nil
}

func (m *Event) GetProposalPayload() *ProposalPayload {
	if x, ok := m.GetMetric().(*Event_ProposalPayload); ok {
		return x.ProposalPayload
	}
	return nil
}

func (m *Event) GetSessionTokensPayload() *SessionTokensPayload {
	if x, ok := m.GetMetric().(*Event_SessionTokensPayload); ok {
		return x.SessionTokensPayload
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Event) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Event_VersionPayload)(nil),
		(*Event_SessionEventPayload)(nil),
		(*Event_SessionStatisticsPayload)(nil),
		(*Event_ProposalPayload)(nil),
		(*Event_SessionTokensPayload)(nil),
	}
}

type VersionPayload struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Os                   string   `protobuf:"bytes,2,opt,name=os,proto3" json:"os,omitempty"`
	Arch                 string   `protobuf:"bytes,3,opt,name=arch,proto3" json:"arch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionPayload) Reset()         { *m = VersionPayload{} }
func (m *VersionPayload) String() string { return proto.CompactTextString(m) }
func (*VersionPayload) ProtoMessage()    {}
func (*VersionPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{2}
}

func (m *VersionPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionPayload.Unmarshal(m, b)
}
func (m *VersionPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionPayload.Marshal(b, m, deterministic)
}
func (m *VersionPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionPayload.Merge(m, src)
}
func (m *VersionPayload) XXX_Size() int {
	return xxx_messageInfo_VersionPayload.Size(m)
}
func (m *VersionPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionPayload.DiscardUnknown(m)
}

var xxx_messageInfo_VersionPayload proto.InternalMessageInfo

func (m *VersionPayload) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *VersionPayload) GetOs() string {
	if m != nil {
		return m.Os
	}
	return ""
}

func (m *VersionPayload) GetArch() string {
	if m != nil {
		return m.Arch
	}
	return ""
}

type SessionEventPayload struct {
	Session              *SessionPayload `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	Event                string          `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SessionEventPayload) Reset()         { *m = SessionEventPayload{} }
func (m *SessionEventPayload) String() string { return proto.CompactTextString(m) }
func (*SessionEventPayload) ProtoMessage()    {}
func (*SessionEventPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{3}
}

func (m *SessionEventPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionEventPayload.Unmarshal(m, b)
}
func (m *SessionEventPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionEventPayload.Marshal(b, m, deterministic)
}
func (m *SessionEventPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionEventPayload.Merge(m, src)
}
func (m *SessionEventPayload) XXX_Size() int {
	return xxx_messageInfo_SessionEventPayload.Size(m)
}
func (m *SessionEventPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionEventPayload.DiscardUnknown(m)
}

var xxx_messageInfo_SessionEventPayload proto.InternalMessageInfo

func (m *SessionEventPayload) GetSession() *SessionPayload {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *SessionEventPayload) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

type SessionStatisticsPayload struct {
	Session              *SessionPayload `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	BytesSent            uint64          `protobuf:"varint,2,opt,name=bytes_sent,json=bytesSent,proto3" json:"bytes_sent,omitempty"`
	BytesReceived        uint64          `protobuf:"varint,3,opt,name=bytes_received,json=bytesReceived,proto3" json:"bytes_received,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SessionStatisticsPayload) Reset()         { *m = SessionStatisticsPayload{} }
func (m *SessionStatisticsPayload) String() string { return proto.CompactTextString(m) }
func (*SessionStatisticsPayload) ProtoMessage()    {}
func (*SessionStatisticsPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{4}
}

func (m *SessionStatisticsPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionStatisticsPayload.Unmarshal(m, b)
}
func (m *SessionStatisticsPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionStatisticsPayload.Marshal(b, m, deterministic)
}
func (m *SessionStatisticsPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionStatisticsPayload.Merge(m, src)
}
func (m *SessionStatisticsPayload) XXX_Size() int {
	return xxx_messageInfo_SessionStatisticsPayload.Size(m)
}
func (m *SessionStatisticsPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionStatisticsPayload.DiscardUnknown(m)
}

var xxx_messageInfo_SessionStatisticsPayload proto.InternalMessageInfo

func (m *SessionStatisticsPayload) GetSession() *SessionPayload {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *SessionStatisticsPayload) GetBytesSent() uint64 {
	if m != nil {
		return m.BytesSent
	}
	return 0
}

func (m *SessionStatisticsPayload) GetBytesReceived() uint64 {
	if m != nil {
		return m.BytesReceived
	}
	return 0
}

type SessionTokensPayload struct {
	Session              *SessionPayload `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	Tokens               uint64          `protobuf:"varint,2,opt,name=tokens,proto3" json:"tokens,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SessionTokensPayload) Reset()         { *m = SessionTokensPayload{} }
func (m *SessionTokensPayload) String() string { return proto.CompactTextString(m) }
func (*SessionTokensPayload) ProtoMessage()    {}
func (*SessionTokensPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{5}
}

func (m *SessionTokensPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionTokensPayload.Unmarshal(m, b)
}
func (m *SessionTokensPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionTokensPayload.Marshal(b, m, deterministic)
}
func (m *SessionTokensPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionTokensPayload.Merge(m, src)
}
func (m *SessionTokensPayload) XXX_Size() int {
	return xxx_messageInfo_SessionTokensPayload.Size(m)
}
func (m *SessionTokensPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionTokensPayload.DiscardUnknown(m)
}

var xxx_messageInfo_SessionTokensPayload proto.InternalMessageInfo

func (m *SessionTokensPayload) GetSession() *SessionPayload {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *SessionTokensPayload) GetTokens() uint64 {
	if m != nil {
		return m.Tokens
	}
	return 0
}

type SessionPayload struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ServiceType          string   `protobuf:"bytes,2,opt,name=service_type,json=serviceType,proto3" json:"service_type,omitempty"`
	ProviderContry       string   `protobuf:"bytes,3,opt,name=provider_contry,json=providerContry,proto3" json:"provider_contry,omitempty"`
	ConsumerContry       string   `protobuf:"bytes,4,opt,name=consumer_contry,json=consumerContry,proto3" json:"consumer_contry,omitempty"`
	AccountantAddress    string   `protobuf:"bytes,5,opt,name=accountant_address,json=accountantAddress,proto3" json:"accountant_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionPayload) Reset()         { *m = SessionPayload{} }
func (m *SessionPayload) String() string { return proto.CompactTextString(m) }
func (*SessionPayload) ProtoMessage()    {}
func (*SessionPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{6}
}

func (m *SessionPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionPayload.Unmarshal(m, b)
}
func (m *SessionPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionPayload.Marshal(b, m, deterministic)
}
func (m *SessionPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionPayload.Merge(m, src)
}
func (m *SessionPayload) XXX_Size() int {
	return xxx_messageInfo_SessionPayload.Size(m)
}
func (m *SessionPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionPayload.DiscardUnknown(m)
}

var xxx_messageInfo_SessionPayload proto.InternalMessageInfo

func (m *SessionPayload) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SessionPayload) GetServiceType() string {
	if m != nil {
		return m.ServiceType
	}
	return ""
}

func (m *SessionPayload) GetProviderContry() string {
	if m != nil {
		return m.ProviderContry
	}
	return ""
}

func (m *SessionPayload) GetConsumerContry() string {
	if m != nil {
		return m.ConsumerContry
	}
	return ""
}

func (m *SessionPayload) GetAccountantAddress() string {
	if m != nil {
		return m.AccountantAddress
	}
	return ""
}

type ProposalPayload struct {
	ProviderId           string          `protobuf:"bytes,1,opt,name=provider_id,json=providerId,proto3" json:"provider_id,omitempty"`
	ServiceType          string          `protobuf:"bytes,2,opt,name=service_type,json=serviceType,proto3" json:"service_type,omitempty"`
	Version              *VersionPayload `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	NodeType             string          `protobuf:"bytes,4,opt,name=node_type,json=nodeType,proto3" json:"node_type,omitempty"`
	Country              string          `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ProposalPayload) Reset()         { *m = ProposalPayload{} }
func (m *ProposalPayload) String() string { return proto.CompactTextString(m) }
func (*ProposalPayload) ProtoMessage()    {}
func (*ProposalPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{7}
}

func (m *ProposalPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProposalPayload.Unmarshal(m, b)
}
func (m *ProposalPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProposalPayload.Marshal(b, m, deterministic)
}
func (m *ProposalPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalPayload.Merge(m, src)
}
func (m *ProposalPayload) XXX_Size() int {
	return xxx_messageInfo_ProposalPayload.Size(m)
}
func (m *ProposalPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalPayload.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalPayload proto.InternalMessageInfo

func (m *ProposalPayload) GetProviderId() string {
	if m != nil {
		return m.ProviderId
	}
	return ""
}

func (m *ProposalPayload) GetServiceType() string {
	if m != nil {
		return m.ServiceType
	}
	return ""
}

func (m *ProposalPayload) GetVersion() *VersionPayload {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *ProposalPayload) GetNodeType() string {
	if m != nil {
		return m.NodeType
	}
	return ""
}

func (m *ProposalPayload) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func init() {
	proto.RegisterType((*Batch)(nil), "metrics.Batch")
	proto.RegisterType((*Event)(nil), "metrics.Event")
	proto.RegisterType((*VersionPayload)(nil), "metrics.VersionPayload")
	proto.RegisterType((*SessionEventPayload)(nil), "metrics.SessionEventPayload")
	proto.RegisterType((*SessionStatisticsPayload)(nil), "metrics.SessionStatisticsPayload")
	proto.RegisterType((*SessionTokensPayload)(nil), "metrics.SessionTokensPayload")
	proto.RegisterType((*SessionPayload)(nil), "metrics.SessionPayload")
	proto.RegisterType((*ProposalPayload)(nil), "metrics.ProposalPayload")
}

func init() {
	proto.RegisterFile("metrics.proto", fileDescriptor_6039342a2ba47b72)
}

var fileDescriptor_6039342a2ba47b72 = []byte{
	// 586 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x26, 0x6d, 0xfa, 0x93, 0x53, 0x96, 0x82, 0x57, 0x46, 0xc4, 0x36, 0xd1, 0x45, 0x02, 0x7a,
	0xc3, 0xd0, 0xc6, 0x13, 0x50, 0x34, 0x89, 0xdd, 0xa0, 0x2a, 0x1d, 0x5c, 0x12, 0x79, 0x89, 0xb5,
	0x59, 0x6c, 0x71, 0xe4, 0xe3, 0x46, 0xca, 0x0b, 0xf0, 0x02, 0x3c, 0x10, 0x37, 0x3c, 0x18, 0x8a,
	0x63, 0x27, 0xb4, 0x74, 0x68, 0xda, 0x5d, 0xfc, 0x9d, 0xcf, 0x9f, 0xbf, 0x93, 0xef, 0xd8, 0xb0,
	0x73, 0xcb, 0x94, 0xe4, 0x09, 0x1e, 0xe7, 0x52, 0x28, 0x41, 0x06, 0x66, 0x19, 0xbe, 0x83, 0xde,
	0x9c, 0xaa, 0xe4, 0x9a, 0xbc, 0x86, 0x3e, 0x2b, 0x58, 0xa6, 0x30, 0x70, 0xa6, 0xdd, 0xd9, 0xe8,
	0xd4, 0x3f, 0xb6, 0x3b, 0xce, 0x2a, 0x38, 0x32, 0xd5, 0xf0, 0x87, 0x0b, 0x3d, 0x8d, 0x90, 0x03,
	0xf0, 0x90, 0x5f, 0x65, 0x54, 0xad, 0x24, 0x0b, 0x9c, 0xa9, 0x33, 0xf3, 0xa2, 0x16, 0x20, 0xfb,
	0xe0, 0x29, 0x2a, 0xaf, 0x98, 0x8a, 0x79, 0x1a, 0x74, 0x74, 0x75, 0x58, 0x03, 0xe7, 0x29, 0x79,
	0x09, 0x23, 0x8e, 0x71, 0x2e, 0x45, 0xc1, 0x53, 0x26, 0x83, 0xee, 0xd4, 0x99, 0x0d, 0x23, 0xe0,
	0xb8, 0x30, 0x08, 0x99, 0xc3, 0xb8, 0x60, 0x12, 0xb9, 0xc8, 0xe2, 0x9c, 0x96, 0x37, 0x82, 0xa6,
	0x81, 0x3b, 0x75, 0x66, 0xa3, 0xd3, 0xe7, 0x8d, 0xad, 0xaf, 0x75, 0x7d, 0x51, 0x97, 0x3f, 0x3d,
	0x8a, 0xfc, 0x62, 0x0d, 0x21, 0x11, 0x3c, 0x43, 0x86, 0x5a, 0x43, 0x7b, 0x6f, 0x94, 0x7a, 0x5a,
	0xe9, 0xa0, 0x51, 0x5a, 0xd6, 0x2c, 0xdd, 0x55, 0x2b, 0xb7, 0x8b, 0xff, 0xc2, 0x84, 0xc2, 0x0b,
	0xab, 0x89, 0x8a, 0x2a, 0x8e, 0x8a, 0x27, 0xd8, 0x08, 0xf7, 0xb5, 0xf0, 0xd1, 0xa6, 0xf0, 0xb2,
	0x61, 0xb6, 0xea, 0x01, 0xde, 0x51, 0x23, 0x67, 0xf0, 0x24, 0x97, 0x22, 0x17, 0x48, 0x6f, 0x1a,
	0xe1, 0x81, 0x16, 0x0e, 0x1a, 0xe1, 0x85, 0x21, 0xb4, 0x7a, 0xe3, 0x7c, 0x1d, 0x22, 0x5f, 0x60,
	0xcf, 0x3a, 0x55, 0xe2, 0x3b, 0xcb, 0x5a, 0x97, 0x43, 0x2d, 0x76, 0xb8, 0xe9, 0xf2, 0x42, 0xb3,
	0x5a, 0xc5, 0x09, 0x6e, 0xc1, 0xe7, 0x43, 0xe8, 0xd7, 0xfb, 0xc2, 0xcf, 0xe0, 0xaf, 0x47, 0x40,
	0x02, 0x18, 0x98, 0x08, 0xcc, 0x38, 0xd8, 0x25, 0xf1, 0xa1, 0x23, 0xd0, 0x4c, 0x41, 0x47, 0x20,
	0x21, 0xe0, 0x52, 0x99, 0x5c, 0xeb, 0xe0, 0xbd, 0x48, 0x7f, 0x87, 0xdf, 0x60, 0x77, 0x4b, 0x10,
	0xe4, 0x04, 0x06, 0xc6, 0x88, 0x16, 0xfd, 0x7b, 0x02, 0x0c, 0xdd, 0x30, 0x23, 0xcb, 0x23, 0x13,
	0xe8, 0xe9, 0xc0, 0xcd, 0x81, 0xf5, 0x22, 0xfc, 0xe9, 0x40, 0x70, 0x57, 0x20, 0x0f, 0x39, 0xe5,
	0x10, 0xe0, 0xb2, 0x54, 0x0c, 0x63, 0xb4, 0x47, 0xb9, 0x91, 0xa7, 0x91, 0x65, 0x75, 0x3b, 0x5e,
	0x81, 0x5f, 0x97, 0x25, 0x4b, 0x18, 0x2f, 0x58, 0xaa, 0x9b, 0x75, 0xa3, 0x1d, 0x8d, 0x46, 0x06,
	0x0c, 0x29, 0x4c, 0xb6, 0xfd, 0xff, 0x87, 0x18, 0xda, 0x83, 0x7e, 0x9d, 0xb4, 0x31, 0x63, 0x56,
	0xe1, 0x6f, 0x07, 0xfc, 0xf5, 0x3d, 0x55, 0x1e, 0x3c, 0x35, 0x21, 0x75, 0x78, 0x4a, 0x8e, 0xe0,
	0x31, 0x32, 0x59, 0xf0, 0x84, 0xc5, 0xaa, 0xcc, 0x99, 0xf9, 0x71, 0x23, 0x83, 0x5d, 0x94, 0x39,
	0x23, 0x6f, 0x60, 0x6c, 0xef, 0x6b, 0x9c, 0x88, 0x4c, 0xc9, 0xd2, 0xa4, 0xe7, 0x5b, 0xf8, 0xa3,
	0x46, 0x2b, 0x62, 0x22, 0x32, 0x5c, 0xdd, 0xb6, 0x44, 0xb7, 0x26, 0x5a, 0xd8, 0x10, 0xdf, 0x02,
	0xa1, 0x49, 0x22, 0x56, 0x99, 0xa2, 0x99, 0x8a, 0x69, 0x9a, 0x4a, 0x86, 0xa8, 0x2f, 0xa7, 0x17,
	0x3d, 0x6d, 0x2b, 0x1f, 0xea, 0x42, 0xf8, 0xcb, 0x81, 0xf1, 0xc6, 0xdc, 0x57, 0xef, 0x48, 0x63,
	0xaa, 0x69, 0x08, 0x2c, 0x74, 0x7e, 0xaf, 0xc6, 0x4e, 0xda, 0xa9, 0xed, 0xfe, 0xf7, 0x89, 0x69,
	0xc7, 0x79, 0x1f, 0xbc, 0x4c, 0xa4, 0x46, 0xb2, 0x6e, 0x6e, 0x58, 0x01, 0x5a, 0x2f, 0x80, 0x81,
	0xb6, 0x2e, 0x4b, 0xd3, 0x8b, 0x5d, 0x5e, 0xf6, 0xf5, 0xdb, 0xfb, 0xfe, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x7b, 0xce, 0xb0, 0x42, 0x8c, 0x05, 0x00, 0x00,
}
