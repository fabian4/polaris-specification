// Code generated by protoc-gen-go. DO NOT EDIT.
// source: routing.proto

package traffic_manage // import "github.com/polarismesh/specification/source/go/api/v1/traffic_manage"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import model "github.com/polarismesh/specification/source/go/api/v1/model"
import anypb "google.golang.org/protobuf/types/known/anypb"
import wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RoutingPolicy int32

const (
	// Route by rule rule => RuleRoutingConfig
	RoutingPolicy_RulePolicy RoutingPolicy = 0
	// Route by destination metadata ==> MetadataRoutingConfig
	RoutingPolicy_MetadataPolicy RoutingPolicy = 1
)

var RoutingPolicy_name = map[int32]string{
	0: "RulePolicy",
	1: "MetadataPolicy",
}
var RoutingPolicy_value = map[string]int32{
	"RulePolicy":     0,
	"MetadataPolicy": 1,
}

func (x RoutingPolicy) String() string {
	return proto.EnumName(RoutingPolicy_name, int32(x))
}
func (RoutingPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_routing_2204fcc6be83eb17, []int{0}
}

type MetadataFailover_FailoverRange int32

const (
	// ALL return all instances
	MetadataFailover_ALL MetadataFailover_FailoverRange = 0
	// OTHERS retuen without thie labels instances
	MetadataFailover_OTHERS MetadataFailover_FailoverRange = 1
	// OTHER_KEYS return other instances which match keys
	MetadataFailover_OTHER_KEYS MetadataFailover_FailoverRange = 2
)

var MetadataFailover_FailoverRange_name = map[int32]string{
	0: "ALL",
	1: "OTHERS",
	2: "OTHER_KEYS",
}
var MetadataFailover_FailoverRange_value = map[string]int32{
	"ALL":        0,
	"OTHERS":     1,
	"OTHER_KEYS": 2,
}

func (x MetadataFailover_FailoverRange) String() string {
	return proto.EnumName(MetadataFailover_FailoverRange_name, int32(x))
}
func (MetadataFailover_FailoverRange) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_routing_2204fcc6be83eb17, []int{5, 0}
}

// label type for gateway request
type SourceMatch_Type int32

const (
	// custom arguments
	SourceMatch_CUSTOM SourceMatch_Type = 0
	// method, match the http post/get/put/delete or grpc method
	SourceMatch_METHOD SourceMatch_Type = 1
	// header, match the http header, dubbo attachment, grpc header
	SourceMatch_HEADER SourceMatch_Type = 2
	// query, match the http query, dubbo argument
	SourceMatch_QUERY SourceMatch_Type = 3
	// caller host ip
	SourceMatch_CALLER_IP SourceMatch_Type = 4
	// path, math the http url
	SourceMatch_PATH SourceMatch_Type = 5
	// cookie match http cookie
	SourceMatch_COOKIE SourceMatch_Type = 6
)

var SourceMatch_Type_name = map[int32]string{
	0: "CUSTOM",
	1: "METHOD",
	2: "HEADER",
	3: "QUERY",
	4: "CALLER_IP",
	5: "PATH",
	6: "COOKIE",
}
var SourceMatch_Type_value = map[string]int32{
	"CUSTOM":    0,
	"METHOD":    1,
	"HEADER":    2,
	"QUERY":     3,
	"CALLER_IP": 4,
	"PATH":      5,
	"COOKIE":    6,
}

func (x SourceMatch_Type) String() string {
	return proto.EnumName(SourceMatch_Type_name, int32(x))
}
func (SourceMatch_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_routing_2204fcc6be83eb17, []int{10, 0}
}

// deprecated: only for compatible to the old version server
type Routing struct {
	// 规则所属服务以及命名空间
	Service   *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Namespace *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// 每个服务可以配置多条入站或者出站规则
	// 对于每个请求，从上到下依次匹配，若命中则终止
	Inbounds             []*Route                `protobuf:"bytes,3,rep,name=inbounds,proto3" json:"inbounds,omitempty"`
	Outbounds            []*Route                `protobuf:"bytes,4,rep,name=outbounds,proto3" json:"outbounds,omitempty"`
	Ctime                *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=ctime,proto3" json:"ctime,omitempty"`
	Mtime                *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=mtime,proto3" json:"mtime,omitempty"`
	Revision             *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=revision,proto3" json:"revision,omitempty"`
	ServiceToken         *wrapperspb.StringValue `protobuf:"bytes,8,opt,name=service_token,proto3" json:"service_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Routing) Reset()         { *m = Routing{} }
func (m *Routing) String() string { return proto.CompactTextString(m) }
func (*Routing) ProtoMessage()    {}
func (*Routing) Descriptor() ([]byte, []int) {
	return fileDescriptor_routing_2204fcc6be83eb17, []int{0}
}
func (m *Routing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Routing.Unmarshal(m, b)
}
func (m *Routing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Routing.Marshal(b, m, deterministic)
}
func (dst *Routing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Routing.Merge(dst, src)
}
func (m *Routing) XXX_Size() int {
	return xxx_messageInfo_Routing.Size(m)
}
func (m *Routing) XXX_DiscardUnknown() {
	xxx_messageInfo_Routing.DiscardUnknown(m)
}

var xxx_messageInfo_Routing proto.InternalMessageInfo

func (m *Routing) GetService() *wrapperspb.StringValue {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *Routing) GetNamespace() *wrapperspb.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *Routing) GetInbounds() []*Route {
	if m != nil {
		return m.Inbounds
	}
	return nil
}

func (m *Routing) GetOutbounds() []*Route {
	if m != nil {
		return m.Outbounds
	}
	return nil
}

func (m *Routing) GetCtime() *wrapperspb.StringValue {
	if m != nil {
		return m.Ctime
	}
	return nil
}

func (m *Routing) GetMtime() *wrapperspb.StringValue {
	if m != nil {
		return m.Mtime
	}
	return nil
}

func (m *Routing) GetRevision() *wrapperspb.StringValue {
	if m != nil {
		return m.Revision
	}
	return nil
}

func (m *Routing) GetServiceToken() *wrapperspb.StringValue {
	if m != nil {
		return m.ServiceToken
	}
	return nil
}

// deprecated: only for compatible to the old version server
type Route struct {
	// 如果匹配Source规则，按照Destination路由
	// 多个Source之间的关系为或
	Sources      []*Source      `protobuf:"bytes,1,rep,name=sources,proto3" json:"sources,omitempty"`
	Destinations []*Destination `protobuf:"bytes,2,rep,name=destinations,proto3" json:"destinations,omitempty"`
	// extendInfo 用于承载一些额外信息
	// case 1: 升级到 v2 版本时，记录对应到 v2 版本的 id 信息
	ExtendInfo           map[string]string `protobuf:"bytes,3,rep,name=extendInfo,proto3" json:"extendInfo,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_routing_2204fcc6be83eb17, []int{1}
}
func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (dst *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(dst, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetSources() []*Source {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *Route) GetDestinations() []*Destination {
	if m != nil {
		return m.Destinations
	}
	return nil
}

func (m *Route) GetExtendInfo() map[string]string {
	if m != nil {
		return m.ExtendInfo
	}
	return nil
}

// deprecated: only for compatible to the old version server
type Source struct {
	// 主调方服务以及命名空间
	Service   *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Namespace *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// 主调方服务实例标签或者请求标签
	// value支持正则匹配
	Metadata             map[string]*model.MatchString `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *Source) Reset()         { *m = Source{} }
func (m *Source) String() string { return proto.CompactTextString(m) }
func (*Source) ProtoMessage()    {}
func (*Source) Descriptor() ([]byte, []int) {
	return fileDescriptor_routing_2204fcc6be83eb17, []int{2}
}
func (m *Source) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Source.Unmarshal(m, b)
}
func (m *Source) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Source.Marshal(b, m, deterministic)
}
func (dst *Source) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Source.Merge(dst, src)
}
func (m *Source) XXX_Size() int {
	return xxx_messageInfo_Source.Size(m)
}
func (m *Source) XXX_DiscardUnknown() {
	xxx_messageInfo_Source.DiscardUnknown(m)
}

var xxx_messageInfo_Source proto.InternalMessageInfo

func (m *Source) GetService() *wrapperspb.StringValue {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *Source) GetNamespace() *wrapperspb.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *Source) GetMetadata() map[string]*model.MatchString {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// deprecated: only for compatible to the old version server
type Destination struct {
	// 被调方服务以及命名空间
	Service   *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Namespace *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// 被调方服务实例标签
	// value支持正则匹配
	Metadata map[string]*model.MatchString `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// 根据服务名和服务实例metadata筛选符合条件的服务实例子集
	// 服务实例子集可以设置优先级和权重
	// 优先级：整型，范围[0, 9]，最高优先级为0
	// 权重：整型
	// 先按优先级路由，如果存在高优先级，不会使用低优先级
	// 如果存在优先级相同的子集，再按权重分配
	// 优先级和权重可以都不设置/设置一个/设置两个
	// 如果部分设置优先级，部分没有设置，认为没有设置的优先级最低
	// 如果部分设置权重，部分没有设置，认为没有设置的权重为0
	// 如果全部没有设置权重，认为权重相同
	Priority *wrapperspb.UInt32Value `protobuf:"bytes,4,opt,name=priority,proto3" json:"priority,omitempty"`
	Weight   *wrapperspb.UInt32Value `protobuf:"bytes,5,opt,name=weight,proto3" json:"weight,omitempty"`
	// 将请求转发到代理服务
	Transfer *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=transfer,proto3" json:"transfer,omitempty"`
	// 是否对该set执行隔离，隔离后，不会再分配流量
	Isolate              *wrapperspb.BoolValue `protobuf:"bytes,7,opt,name=isolate,proto3" json:"isolate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Destination) Reset()         { *m = Destination{} }
func (m *Destination) String() string { return proto.CompactTextString(m) }
func (*Destination) ProtoMessage()    {}
func (*Destination) Descriptor() ([]byte, []int) {
	return fileDescriptor_routing_2204fcc6be83eb17, []int{3}
}
func (m *Destination) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Destination.Unmarshal(m, b)
}
func (m *Destination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Destination.Marshal(b, m, deterministic)
}
func (dst *Destination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Destination.Merge(dst, src)
}
func (m *Destination) XXX_Size() int {
	return xxx_messageInfo_Destination.Size(m)
}
func (m *Destination) XXX_DiscardUnknown() {
	xxx_messageInfo_Destination.DiscardUnknown(m)
}

var xxx_messageInfo_Destination proto.InternalMessageInfo

func (m *Destination) GetService() *wrapperspb.StringValue {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *Destination) GetNamespace() *wrapperspb.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *Destination) GetMetadata() map[string]*model.MatchString {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Destination) GetPriority() *wrapperspb.UInt32Value {
	if m != nil {
		return m.Priority
	}
	return nil
}

func (m *Destination) GetWeight() *wrapperspb.UInt32Value {
	if m != nil {
		return m.Weight
	}
	return nil
}

func (m *Destination) GetTransfer() *wrapperspb.StringValue {
	if m != nil {
		return m.Transfer
	}
	return nil
}

func (m *Destination) GetIsolate() *wrapperspb.BoolValue {
	if m != nil {
		return m.Isolate
	}
	return nil
}

// configuration root for route
type Router struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// route rule name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// namespace namingspace of routing rules
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Enable this router
	Enable bool `protobuf:"varint,4,opt,name=enable,proto3" json:"enable,omitempty"`
	// Router type
	RoutingPolicy RoutingPolicy `protobuf:"varint,5,opt,name=routing_policy,proto3,enum=v1.RoutingPolicy" json:"routing_policy,omitempty"`
	// Routing configuration for router
	RoutingConfig *anypb.Any `protobuf:"bytes,6,opt,name=routing_config,proto3" json:"routing_config,omitempty"`
	// revision routing version
	Revision string `protobuf:"bytes,7,opt,name=revision,proto3" json:"revision,omitempty"`
	// ctime create time of the rules
	Ctime string `protobuf:"bytes,8,opt,name=ctime,proto3" json:"ctime,omitempty"`
	// mtime modify time of the rules
	Mtime string `protobuf:"bytes,9,opt,name=mtime,proto3" json:"mtime,omitempty"`
	// etime enable time of the rules
	Etime string `protobuf:"bytes,10,opt,name=etime,proto3" json:"etime,omitempty"`
	// priority rules priority
	Priority uint32 `protobuf:"varint,11,opt,name=priority,proto3" json:"priority,omitempty"`
	// description simple description rules
	Description string `protobuf:"bytes,12,opt,name=description,proto3" json:"description,omitempty"`
	// extendInfo 用于承载一些额外信息
	// case 1: 升级到 v2 版本时，记录对应到 v1 版本的 id 信息
	ExtendInfo           map[string]string `protobuf:"bytes,20,rep,name=extendInfo,proto3" json:"extendInfo,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Router) Reset()         { *m = Router{} }
func (m *Router) String() string { return proto.CompactTextString(m) }
func (*Router) ProtoMessage()    {}
func (*Router) Descriptor() ([]byte, []int) {
	return fileDescriptor_routing_2204fcc6be83eb17, []int{4}
}
func (m *Router) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Router.Unmarshal(m, b)
}
func (m *Router) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Router.Marshal(b, m, deterministic)
}
func (dst *Router) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Router.Merge(dst, src)
}
func (m *Router) XXX_Size() int {
	return xxx_messageInfo_Router.Size(m)
}
func (m *Router) XXX_DiscardUnknown() {
	xxx_messageInfo_Router.DiscardUnknown(m)
}

var xxx_messageInfo_Router proto.InternalMessageInfo

func (m *Router) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Router) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Router) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Router) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *Router) GetRoutingPolicy() RoutingPolicy {
	if m != nil {
		return m.RoutingPolicy
	}
	return RoutingPolicy_RulePolicy
}

func (m *Router) GetRoutingConfig() *anypb.Any {
	if m != nil {
		return m.RoutingConfig
	}
	return nil
}

func (m *Router) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func (m *Router) GetCtime() string {
	if m != nil {
		return m.Ctime
	}
	return ""
}

func (m *Router) GetMtime() string {
	if m != nil {
		return m.Mtime
	}
	return ""
}

func (m *Router) GetEtime() string {
	if m != nil {
		return m.Etime
	}
	return ""
}

func (m *Router) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *Router) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Router) GetExtendInfo() map[string]string {
	if m != nil {
		return m.ExtendInfo
	}
	return nil
}

type MetadataFailover struct {
	// failover_range metadata route bottom type
	FailoverRange MetadataFailover_FailoverRange `protobuf:"varint,1,opt,name=failover_range,json=failoverRange,proto3,enum=v1.MetadataFailover_FailoverRange" json:"failover_range,omitempty"`
	// only use to failover_range == OTHER_KEYS
	Labels               map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MetadataFailover) Reset()         { *m = MetadataFailover{} }
func (m *MetadataFailover) String() string { return proto.CompactTextString(m) }
func (*MetadataFailover) ProtoMessage()    {}
func (*MetadataFailover) Descriptor() ([]byte, []int) {
	return fileDescriptor_routing_2204fcc6be83eb17, []int{5}
}
func (m *MetadataFailover) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetadataFailover.Unmarshal(m, b)
}
func (m *MetadataFailover) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetadataFailover.Marshal(b, m, deterministic)
}
func (dst *MetadataFailover) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataFailover.Merge(dst, src)
}
func (m *MetadataFailover) XXX_Size() int {
	return xxx_messageInfo_MetadataFailover.Size(m)
}
func (m *MetadataFailover) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataFailover.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataFailover proto.InternalMessageInfo

func (m *MetadataFailover) GetFailoverRange() MetadataFailover_FailoverRange {
	if m != nil {
		return m.FailoverRange
	}
	return MetadataFailover_ALL
}

func (m *MetadataFailover) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

// MetadataRoutingConfig metadata routing configuration
type MetadataRoutingConfig struct {
	// service
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// namespace
	Namespace string            `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Labels    map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// When metadata not found, it will fall back to the
	Failover             *MetadataFailover `protobuf:"bytes,4,opt,name=failover,proto3" json:"failover,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MetadataRoutingConfig) Reset()         { *m = MetadataRoutingConfig{} }
func (m *MetadataRoutingConfig) String() string { return proto.CompactTextString(m) }
func (*MetadataRoutingConfig) ProtoMessage()    {}
func (*MetadataRoutingConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_routing_2204fcc6be83eb17, []int{6}
}
func (m *MetadataRoutingConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetadataRoutingConfig.Unmarshal(m, b)
}
func (m *MetadataRoutingConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetadataRoutingConfig.Marshal(b, m, deterministic)
}
func (dst *MetadataRoutingConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataRoutingConfig.Merge(dst, src)
}
func (m *MetadataRoutingConfig) XXX_Size() int {
	return xxx_messageInfo_MetadataRoutingConfig.Size(m)
}
func (m *MetadataRoutingConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataRoutingConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataRoutingConfig proto.InternalMessageInfo

func (m *MetadataRoutingConfig) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *MetadataRoutingConfig) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *MetadataRoutingConfig) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *MetadataRoutingConfig) GetFailover() *MetadataFailover {
	if m != nil {
		return m.Failover
	}
	return nil
}

// RuleRoutingConfig routing configuration
type RuleRoutingConfig struct {
	// source source info
	Sources []*SourceService `protobuf:"bytes,1,rep,name=sources,proto3" json:"sources,omitempty"`
	// destination destinations info
	Destinations         []*DestinationGroup `protobuf:"bytes,2,rep,name=destinations,proto3" json:"destinations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *RuleRoutingConfig) Reset()         { *m = RuleRoutingConfig{} }
func (m *RuleRoutingConfig) String() string { return proto.CompactTextString(m) }
func (*RuleRoutingConfig) ProtoMessage()    {}
func (*RuleRoutingConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_routing_2204fcc6be83eb17, []int{7}
}
func (m *RuleRoutingConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RuleRoutingConfig.Unmarshal(m, b)
}
func (m *RuleRoutingConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RuleRoutingConfig.Marshal(b, m, deterministic)
}
func (dst *RuleRoutingConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuleRoutingConfig.Merge(dst, src)
}
func (m *RuleRoutingConfig) XXX_Size() int {
	return xxx_messageInfo_RuleRoutingConfig.Size(m)
}
func (m *RuleRoutingConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RuleRoutingConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RuleRoutingConfig proto.InternalMessageInfo

func (m *RuleRoutingConfig) GetSources() []*SourceService {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *RuleRoutingConfig) GetDestinations() []*DestinationGroup {
	if m != nil {
		return m.Destinations
	}
	return nil
}

type SourceService struct {
	// Main tuning service and namespace
	Service   string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Master Control Service Example Tag or Request Label
	// Value supports regular matching
	Arguments            []*SourceMatch `protobuf:"bytes,3,rep,name=arguments,proto3" json:"arguments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SourceService) Reset()         { *m = SourceService{} }
func (m *SourceService) String() string { return proto.CompactTextString(m) }
func (*SourceService) ProtoMessage()    {}
func (*SourceService) Descriptor() ([]byte, []int) {
	return fileDescriptor_routing_2204fcc6be83eb17, []int{8}
}
func (m *SourceService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourceService.Unmarshal(m, b)
}
func (m *SourceService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourceService.Marshal(b, m, deterministic)
}
func (dst *SourceService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceService.Merge(dst, src)
}
func (m *SourceService) XXX_Size() int {
	return xxx_messageInfo_SourceService.Size(m)
}
func (m *SourceService) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceService.DiscardUnknown(m)
}

var xxx_messageInfo_SourceService proto.InternalMessageInfo

func (m *SourceService) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *SourceService) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *SourceService) GetArguments() []*SourceMatch {
	if m != nil {
		return m.Arguments
	}
	return nil
}

type DestinationGroup struct {
	// Templated service and namespace
	Service   string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Templated service example label
	// Value supports regular matching
	Labels map[string]*model.MatchString `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// According to the service name and service instance Metadata Filter the
	// qualified service instance subset Service instance subset can set priority
	// and weight Priority: integer, range [0, 9], the highest priority is 0
	// Weight: Integer
	// Press priority routing, if there is high priority, low priority will not
	// use If there is a subset of the same priority, then assign by weight
	// Priority and weight can be not set / set up one / set two
	// If the section is set priority, some are not set, it is considered that the
	// priority is not set. If the part is set, some is not set, it is considered
	// that the weight is not set to 0 If you have no weight, you think the weight
	// is the same
	Priority uint32 `protobuf:"varint,4,opt,name=priority,proto3" json:"priority,omitempty"`
	Weight   uint32 `protobuf:"varint,5,opt,name=weight,proto3" json:"weight,omitempty"`
	// Forward requests to proxy service
	Transfer string `protobuf:"bytes,6,opt,name=transfer,proto3" json:"transfer,omitempty"`
	// Whether to isolate the SET, after isolation, no traffic will be allocated
	Isolate bool `protobuf:"varint,7,opt,name=isolate,proto3" json:"isolate,omitempty"`
	// name desition name
	Name                 string   `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DestinationGroup) Reset()         { *m = DestinationGroup{} }
func (m *DestinationGroup) String() string { return proto.CompactTextString(m) }
func (*DestinationGroup) ProtoMessage()    {}
func (*DestinationGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_routing_2204fcc6be83eb17, []int{9}
}
func (m *DestinationGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestinationGroup.Unmarshal(m, b)
}
func (m *DestinationGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestinationGroup.Marshal(b, m, deterministic)
}
func (dst *DestinationGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestinationGroup.Merge(dst, src)
}
func (m *DestinationGroup) XXX_Size() int {
	return xxx_messageInfo_DestinationGroup.Size(m)
}
func (m *DestinationGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_DestinationGroup.DiscardUnknown(m)
}

var xxx_messageInfo_DestinationGroup proto.InternalMessageInfo

func (m *DestinationGroup) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *DestinationGroup) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *DestinationGroup) GetLabels() map[string]*model.MatchString {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *DestinationGroup) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *DestinationGroup) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *DestinationGroup) GetTransfer() string {
	if m != nil {
		return m.Transfer
	}
	return ""
}

func (m *DestinationGroup) GetIsolate() bool {
	if m != nil {
		return m.Isolate
	}
	return false
}

func (m *DestinationGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// SourceMatch
type SourceMatch struct {
	Type SourceMatch_Type `protobuf:"varint,1,opt,name=type,proto3,enum=v1.SourceMatch_Type" json:"type,omitempty"`
	// header key or query key
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// header value or query value
	Value                *model.MatchString `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SourceMatch) Reset()         { *m = SourceMatch{} }
func (m *SourceMatch) String() string { return proto.CompactTextString(m) }
func (*SourceMatch) ProtoMessage()    {}
func (*SourceMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_routing_2204fcc6be83eb17, []int{10}
}
func (m *SourceMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourceMatch.Unmarshal(m, b)
}
func (m *SourceMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourceMatch.Marshal(b, m, deterministic)
}
func (dst *SourceMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceMatch.Merge(dst, src)
}
func (m *SourceMatch) XXX_Size() int {
	return xxx_messageInfo_SourceMatch.Size(m)
}
func (m *SourceMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceMatch.DiscardUnknown(m)
}

var xxx_messageInfo_SourceMatch proto.InternalMessageInfo

func (m *SourceMatch) GetType() SourceMatch_Type {
	if m != nil {
		return m.Type
	}
	return SourceMatch_CUSTOM
}

func (m *SourceMatch) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SourceMatch) GetValue() *model.MatchString {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Routing)(nil), "v1.Routing")
	proto.RegisterType((*Route)(nil), "v1.Route")
	proto.RegisterMapType((map[string]string)(nil), "v1.Route.ExtendInfoEntry")
	proto.RegisterType((*Source)(nil), "v1.Source")
	proto.RegisterMapType((map[string]*model.MatchString)(nil), "v1.Source.MetadataEntry")
	proto.RegisterType((*Destination)(nil), "v1.Destination")
	proto.RegisterMapType((map[string]*model.MatchString)(nil), "v1.Destination.MetadataEntry")
	proto.RegisterType((*Router)(nil), "v1.Router")
	proto.RegisterMapType((map[string]string)(nil), "v1.Router.ExtendInfoEntry")
	proto.RegisterType((*MetadataFailover)(nil), "v1.MetadataFailover")
	proto.RegisterMapType((map[string]string)(nil), "v1.MetadataFailover.LabelsEntry")
	proto.RegisterType((*MetadataRoutingConfig)(nil), "v1.MetadataRoutingConfig")
	proto.RegisterMapType((map[string]string)(nil), "v1.MetadataRoutingConfig.LabelsEntry")
	proto.RegisterType((*RuleRoutingConfig)(nil), "v1.RuleRoutingConfig")
	proto.RegisterType((*SourceService)(nil), "v1.SourceService")
	proto.RegisterType((*DestinationGroup)(nil), "v1.DestinationGroup")
	proto.RegisterMapType((map[string]*model.MatchString)(nil), "v1.DestinationGroup.LabelsEntry")
	proto.RegisterType((*SourceMatch)(nil), "v1.SourceMatch")
	proto.RegisterEnum("v1.RoutingPolicy", RoutingPolicy_name, RoutingPolicy_value)
	proto.RegisterEnum("v1.MetadataFailover_FailoverRange", MetadataFailover_FailoverRange_name, MetadataFailover_FailoverRange_value)
	proto.RegisterEnum("v1.SourceMatch_Type", SourceMatch_Type_name, SourceMatch_Type_value)
}

func init() { proto.RegisterFile("routing.proto", fileDescriptor_routing_2204fcc6be83eb17) }

var fileDescriptor_routing_2204fcc6be83eb17 = []byte{
	// 1182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0xc1, 0x6e, 0xdb, 0x46,
	0x13, 0x0e, 0x29, 0x59, 0x16, 0x47, 0x91, 0xc2, 0x2c, 0xf4, 0x07, 0x8c, 0x90, 0xbf, 0x30, 0x84,
	0x06, 0x35, 0x5a, 0x94, 0x6a, 0x6c, 0xa3, 0xb5, 0x83, 0xe6, 0x60, 0xc7, 0x6a, 0xed, 0xc6, 0x86,
	0x53, 0xca, 0x29, 0x90, 0x5c, 0x04, 0x8a, 0x5a, 0xd1, 0x8b, 0x50, 0xbb, 0x04, 0xb9, 0x52, 0xaa,
	0x9e, 0x7a, 0xec, 0xa5, 0x4f, 0xd3, 0xd7, 0x28, 0xfa, 0x04, 0x7d, 0x82, 0x1e, 0x7b, 0xed, 0xa5,
	0xe0, 0xee, 0x92, 0x22, 0x19, 0xb9, 0x55, 0xd2, 0x16, 0xb9, 0x71, 0x77, 0xbe, 0x99, 0xdd, 0xf9,
	0x76, 0xe6, 0x1b, 0x42, 0x33, 0x62, 0x33, 0x4e, 0xa8, 0x6f, 0x87, 0x11, 0xe3, 0x0c, 0xe9, 0xf3,
	0x07, 0x9d, 0xf7, 0x7c, 0xc6, 0xfc, 0x00, 0xf7, 0xc4, 0xce, 0x68, 0x36, 0xe9, 0xbd, 0x8a, 0xdc,
	0x30, 0xc4, 0x51, 0x2c, 0x31, 0x9d, 0xbb, 0x65, 0xbb, 0x4b, 0x17, 0xca, 0xd4, 0x98, 0xb2, 0x31,
	0x0e, 0xe4, 0xa2, 0xfb, 0x53, 0x05, 0x36, 0x1d, 0x19, 0x1d, 0x7d, 0x0a, 0x9b, 0x31, 0x8e, 0xe6,
	0xc4, 0xc3, 0x96, 0xb6, 0xa5, 0x6d, 0x37, 0x76, 0xee, 0xd9, 0x32, 0x8a, 0x9d, 0x46, 0xb1, 0x07,
	0x3c, 0x22, 0xd4, 0xff, 0xc6, 0x0d, 0x66, 0xd8, 0x49, 0xc1, 0xe8, 0x21, 0x18, 0xd4, 0x9d, 0xe2,
	0x38, 0x74, 0x3d, 0x6c, 0xe9, 0x6b, 0x78, 0x2e, 0xe1, 0xe8, 0x3e, 0xd4, 0x09, 0x1d, 0xb1, 0x19,
	0x1d, 0xc7, 0x56, 0x65, 0xab, 0xb2, 0xdd, 0xd8, 0x31, 0xec, 0xf9, 0x03, 0x3b, 0xb9, 0x12, 0x76,
	0x32, 0x13, 0xfa, 0x00, 0x0c, 0x36, 0xe3, 0x0a, 0x57, 0x2d, 0xe3, 0x96, 0x36, 0xb4, 0x03, 0x1b,
	0x1e, 0x27, 0x53, 0x6c, 0x6d, 0xac, 0x71, 0x0f, 0x09, 0x4d, 0x7c, 0xa6, 0xc2, 0xa7, 0xb6, 0x8e,
	0x8f, 0x80, 0xa2, 0x7d, 0xa8, 0x47, 0x78, 0x4e, 0x62, 0xc2, 0xa8, 0xb5, 0xb9, 0x86, 0x5b, 0x86,
	0x46, 0x47, 0xd0, 0x54, 0xc4, 0x0d, 0x39, 0x7b, 0x89, 0xa9, 0x55, 0x5f, 0xc3, 0xbd, 0xe8, 0xd2,
	0xfd, 0x55, 0x83, 0x0d, 0x91, 0x3a, 0x7a, 0x1f, 0x36, 0x63, 0x36, 0x8b, 0x3c, 0x1c, 0x5b, 0x9a,
	0xa0, 0x05, 0x12, 0x5a, 0x06, 0x62, 0xcb, 0x49, 0x4d, 0x68, 0x17, 0x6e, 0x8e, 0x71, 0xcc, 0x09,
	0x75, 0x39, 0x61, 0x34, 0xb6, 0x74, 0x01, 0xbd, 0x95, 0x40, 0x8f, 0x97, 0xfb, 0x4e, 0x01, 0x84,
	0x0e, 0x00, 0xf0, 0xb7, 0x1c, 0xd3, 0xf1, 0x29, 0x9d, 0x30, 0xf5, 0x38, 0x77, 0x33, 0xd2, 0xed,
	0x7e, 0x66, 0xeb, 0x53, 0x1e, 0x2d, 0x9c, 0x1c, 0xb8, 0xf3, 0x08, 0x6e, 0x95, 0xcc, 0xc8, 0x84,
	0xca, 0x4b, 0xbc, 0x10, 0x85, 0x65, 0x38, 0xc9, 0x27, 0x6a, 0xc3, 0xc6, 0x3c, 0x49, 0x4e, 0x94,
	0x8c, 0xe1, 0xc8, 0xc5, 0x43, 0x7d, 0x5f, 0xeb, 0x7e, 0xaf, 0x43, 0x4d, 0xa6, 0xf0, 0x4e, 0x6a,
	0x72, 0x0f, 0xea, 0x53, 0xcc, 0xdd, 0xb1, 0xcb, 0x5d, 0x95, 0xb6, 0xb5, 0x24, 0xd5, 0x3e, 0x57,
	0x26, 0x99, 0x75, 0x86, 0xec, 0x9c, 0x41, 0xb3, 0x60, 0x5a, 0x91, 0xf1, 0xfd, 0x7c, 0xc6, 0x8a,
	0xff, 0x73, 0x97, 0x7b, 0x57, 0xf2, 0x22, 0x79, 0x0a, 0x7e, 0xab, 0x40, 0x23, 0xf7, 0x34, 0xef,
	0x84, 0x87, 0x83, 0xd7, 0x78, 0xf8, 0x7f, 0xa9, 0x62, 0xae, 0x23, 0x23, 0x69, 0x8f, 0x30, 0x22,
	0x2c, 0x22, 0x7c, 0x61, 0x55, 0xaf, 0x39, 0xf5, 0xd9, 0x29, 0xe5, 0xbb, 0x3b, 0xaa, 0x3d, 0x52,
	0x34, 0xda, 0x83, 0xda, 0x2b, 0x4c, 0xfc, 0x2b, 0x7e, 0x6d, 0x07, 0xe7, 0xfd, 0x14, 0x36, 0x39,
	0x8f, 0x47, 0x2e, 0x8d, 0x27, 0x38, 0x5a, 0xab, 0x8b, 0x33, 0x34, 0xda, 0x83, 0x4d, 0x12, 0xb3,
	0xc0, 0xe5, 0x58, 0xf5, 0x71, 0xe7, 0x35, 0xc7, 0x23, 0xc6, 0x02, 0x45, 0xab, 0x82, 0xfe, 0xcb,
	0x8f, 0xfd, 0x7b, 0x05, 0x6a, 0xa2, 0xa9, 0x22, 0xd4, 0x02, 0x9d, 0x8c, 0x55, 0x18, 0x9d, 0x8c,
	0x11, 0x82, 0x6a, 0xf2, 0x20, 0xaa, 0x47, 0xc4, 0x37, 0xba, 0x97, 0x7f, 0xd3, 0x8a, 0x30, 0xe4,
	0x5e, 0xed, 0x0e, 0xd4, 0x30, 0x75, 0x47, 0x01, 0x16, 0xc4, 0xd7, 0x1d, 0xb5, 0x42, 0x07, 0xd0,
	0x52, 0x63, 0x64, 0x18, 0xb2, 0x80, 0x78, 0x0b, 0x41, 0x70, 0x6b, 0xe7, 0x76, 0xda, 0xd2, 0x84,
	0xfa, 0x4f, 0x85, 0xc1, 0x29, 0x01, 0xd1, 0xe7, 0x4b, 0x57, 0x8f, 0xd1, 0x09, 0xf1, 0x15, 0xc7,
	0xed, 0xd7, 0xa8, 0x3a, 0xa4, 0x39, 0x6f, 0x89, 0x45, 0x9d, 0x92, 0x54, 0x1a, 0x39, 0x31, 0x6c,
	0xa7, 0x72, 0x5d, 0x97, 0x1a, 0x20, 0x05, 0xb9, 0x9d, 0x0a, 0xb2, 0x21, 0x77, 0xa7, 0xe9, 0x2e,
	0x16, 0xbb, 0x20, 0x77, 0xc5, 0x22, 0x89, 0x9e, 0x55, 0x5a, 0x63, 0x4b, 0xdb, 0x6e, 0xe6, 0x6a,
	0x69, 0x0b, 0x1a, 0x63, 0x1c, 0x7b, 0x11, 0x09, 0x93, 0x62, 0xb5, 0x6e, 0x0a, 0xbf, 0xfc, 0x16,
	0x7a, 0x58, 0xd0, 0xb8, 0xb6, 0x28, 0xf2, 0x4e, 0xa6, 0x71, 0xd1, 0x7f, 0x29, 0x72, 0x3f, 0xe8,
	0x60, 0xa6, 0x35, 0xf4, 0x85, 0x4b, 0x02, 0x36, 0xc7, 0x11, 0x3a, 0x85, 0xd6, 0x44, 0x7d, 0x0f,
	0x23, 0x97, 0xfa, 0xb2, 0xdb, 0x5b, 0x3b, 0x5d, 0x51, 0x3d, 0x25, 0xb4, 0x9d, 0x7e, 0x38, 0x09,
	0xd2, 0x69, 0x4e, 0xf2, 0x4b, 0xb4, 0x0f, 0xb5, 0xc0, 0x1d, 0xe1, 0x20, 0x55, 0xfb, 0xad, 0x95,
	0x21, 0xce, 0x04, 0x44, 0x26, 0xa7, 0xf0, 0x9d, 0x03, 0x68, 0xe4, 0xb6, 0xdf, 0x28, 0xa9, 0x3d,
	0x68, 0x16, 0x2e, 0x85, 0x36, 0xa1, 0x72, 0x78, 0x76, 0x66, 0xde, 0x40, 0x00, 0xb5, 0x8b, 0xcb,
	0x93, 0xbe, 0x33, 0x30, 0x35, 0xd4, 0x02, 0x10, 0xdf, 0xc3, 0x27, 0xfd, 0xe7, 0x03, 0x53, 0xef,
	0xfe, 0xa1, 0xc1, 0xff, 0xd2, 0x9b, 0xa9, 0x4a, 0x7c, 0x2c, 0x6b, 0xc7, 0x2a, 0xca, 0x9e, 0xb1,
	0x14, 0xb6, 0x7b, 0x65, 0x61, 0x2b, 0x34, 0xc1, 0xa3, 0x2c, 0x79, 0x29, 0x5c, 0xf7, 0xf3, 0xc9,
	0x17, 0x8e, 0x58, 0xc5, 0x00, 0xfa, 0x04, 0xea, 0x29, 0x99, 0x4a, 0xbe, 0xda, 0xab, 0xd8, 0x73,
	0x32, 0xd4, 0x3f, 0xe1, 0xec, 0x3b, 0xb8, 0xed, 0xcc, 0x02, 0x5c, 0x4c, 0xfc, 0xa3, 0xf2, 0x5c,
	0xbf, 0xbd, 0x1c, 0x41, 0x03, 0x49, 0xc1, 0x72, 0xbc, 0xef, 0xaf, 0x1c, 0xef, 0xed, 0x92, 0x58,
	0x7f, 0x19, 0xb1, 0x59, 0x58, 0x9c, 0xf1, 0xdd, 0x39, 0x34, 0x0b, 0x31, 0xdf, 0x9a, 0xf0, 0x8f,
	0xc1, 0x70, 0x23, 0x7f, 0x36, 0xc5, 0x94, 0xa7, 0x9c, 0xdf, 0x5a, 0xde, 0x58, 0xe8, 0x9e, 0xb3,
	0x44, 0x74, 0x7f, 0xd1, 0xc1, 0x2c, 0x5f, 0xed, 0xad, 0xcf, 0xde, 0x2f, 0x3d, 0xf6, 0xd6, 0xaa,
	0xc4, 0x57, 0xbe, 0x73, 0xa7, 0x34, 0xa6, 0xf2, 0xe2, 0x71, 0xa7, 0x30, 0x88, 0x9a, 0xd9, 0xa8,
	0xe9, 0x94, 0x46, 0x8d, 0x91, 0x1b, 0x26, 0x56, 0x71, 0x98, 0xd4, 0xb3, 0x81, 0x91, 0xe9, 0x78,
	0x7d, 0xa9, 0xe3, 0x9d, 0xaf, 0xfe, 0xae, 0x66, 0xde, 0x60, 0x84, 0xfc, 0xac, 0x41, 0x23, 0xc7,
	0x35, 0xda, 0x86, 0x2a, 0x5f, 0x84, 0xa9, 0x7c, 0xb4, 0x4b, 0x4f, 0x61, 0x5f, 0x2e, 0x42, 0xec,
	0x08, 0x44, 0x7a, 0xac, 0xbe, 0xe2, 0xd8, 0xca, 0x5f, 0x1d, 0xdb, 0x7d, 0x01, 0xd5, 0x24, 0x4c,
	0xd2, 0xd9, 0x8f, 0x9f, 0x0d, 0x2e, 0x2f, 0xce, 0x65, 0x97, 0x9f, 0xf7, 0x2f, 0x4f, 0x2e, 0x8e,
	0x4d, 0x2d, 0xf9, 0x3e, 0xe9, 0x1f, 0x1e, 0xf7, 0x1d, 0x53, 0x47, 0x06, 0x6c, 0x7c, 0xfd, 0xac,
	0xef, 0x3c, 0x37, 0x2b, 0xa8, 0x09, 0xc6, 0xe3, 0xc3, 0xb3, 0xb3, 0xbe, 0x33, 0x3c, 0x7d, 0x6a,
	0x56, 0x51, 0x1d, 0xaa, 0x4f, 0x0f, 0x2f, 0x4f, 0xcc, 0x0d, 0x11, 0xe7, 0xe2, 0xe2, 0xc9, 0x69,
	0xdf, 0xac, 0x7d, 0xb8, 0x0b, 0xcd, 0xc2, 0x48, 0x4a, 0x24, 0x23, 0x69, 0x12, 0xb9, 0x32, 0x6f,
	0x20, 0x04, 0xad, 0xb4, 0x1b, 0xd5, 0x9e, 0x76, 0xf4, 0xa3, 0x06, 0x9f, 0x79, 0x6c, 0x6a, 0x73,
	0x4c, 0x3d, 0x4c, 0xb9, 0x1d, 0xb2, 0xc0, 0x8d, 0x48, 0x6c, 0xc7, 0x21, 0xf6, 0xc8, 0x84, 0x78,
	0xf2, 0x97, 0xc5, 0x0d, 0x49, 0x92, 0x10, 0x8f, 0xdc, 0xc9, 0x84, 0x78, 0xf6, 0xd4, 0xa5, 0xae,
	0x8f, 0x8f, 0x6e, 0xa6, 0xc7, 0x25, 0x93, 0xec, 0xc5, 0xb1, 0x4f, 0xf8, 0xd5, 0x6c, 0x64, 0x7b,
	0x6c, 0xda, 0x53, 0x51, 0xa6, 0x38, 0xbe, 0xea, 0x15, 0x22, 0xf5, 0x64, 0xfb, 0xf5, 0x7c, 0xd6,
	0x73, 0x43, 0xd2, 0x9b, 0x3f, 0xe8, 0xa9, 0x98, 0x43, 0x19, 0x73, 0x54, 0x13, 0x63, 0x71, 0xf7,
	0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xab, 0x15, 0x43, 0x97, 0xbf, 0x0d, 0x00, 0x00,
}