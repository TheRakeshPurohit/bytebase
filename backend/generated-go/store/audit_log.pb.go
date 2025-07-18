// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: store/audit_log.proto

package store

import (
	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuditLog_Severity int32

const (
	AuditLog_DEFAULT   AuditLog_Severity = 0
	AuditLog_DEBUG     AuditLog_Severity = 1
	AuditLog_INFO      AuditLog_Severity = 2
	AuditLog_NOTICE    AuditLog_Severity = 3
	AuditLog_WARNING   AuditLog_Severity = 4
	AuditLog_ERROR     AuditLog_Severity = 5
	AuditLog_CRITICAL  AuditLog_Severity = 6
	AuditLog_ALERT     AuditLog_Severity = 7
	AuditLog_EMERGENCY AuditLog_Severity = 8
)

// Enum value maps for AuditLog_Severity.
var (
	AuditLog_Severity_name = map[int32]string{
		0: "DEFAULT",
		1: "DEBUG",
		2: "INFO",
		3: "NOTICE",
		4: "WARNING",
		5: "ERROR",
		6: "CRITICAL",
		7: "ALERT",
		8: "EMERGENCY",
	}
	AuditLog_Severity_value = map[string]int32{
		"DEFAULT":   0,
		"DEBUG":     1,
		"INFO":      2,
		"NOTICE":    3,
		"WARNING":   4,
		"ERROR":     5,
		"CRITICAL":  6,
		"ALERT":     7,
		"EMERGENCY": 8,
	}
)

func (x AuditLog_Severity) Enum() *AuditLog_Severity {
	p := new(AuditLog_Severity)
	*p = x
	return p
}

func (x AuditLog_Severity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuditLog_Severity) Descriptor() protoreflect.EnumDescriptor {
	return file_store_audit_log_proto_enumTypes[0].Descriptor()
}

func (AuditLog_Severity) Type() protoreflect.EnumType {
	return &file_store_audit_log_proto_enumTypes[0]
}

func (x AuditLog_Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuditLog_Severity.Descriptor instead.
func (AuditLog_Severity) EnumDescriptor() ([]byte, []int) {
	return file_store_audit_log_proto_rawDescGZIP(), []int{0, 0}
}

type AuditLog struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The project or workspace the audit log belongs to.
	// Formats:
	// - projects/{project}
	// - workspaces/{workspace}
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// e.g. /bytebase.v1.SQLService/Query
	Method string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	// resource name
	// projects/{project}
	Resource string `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
	// Format: users/{userUID}.
	User     string            `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	Severity AuditLog_Severity `protobuf:"varint,5,opt,name=severity,proto3,enum=bytebase.store.AuditLog_Severity" json:"severity,omitempty"`
	// Marshalled request.
	Request string `protobuf:"bytes,6,opt,name=request,proto3" json:"request,omitempty"`
	// Marshalled response.
	// Some fields are omitted because they are too large or contain sensitive information.
	Response string         `protobuf:"bytes,7,opt,name=response,proto3" json:"response,omitempty"`
	Status   *status.Status `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	// The latency of the RPC.
	Latency *durationpb.Duration `protobuf:"bytes,9,opt,name=latency,proto3" json:"latency,omitempty"`
	// service-specific data about the request, response, and other activities.
	ServiceData *anypb.Any `protobuf:"bytes,10,opt,name=service_data,json=serviceData,proto3" json:"service_data,omitempty"`
	// Metadata about the operation.
	RequestMetadata *RequestMetadata `protobuf:"bytes,11,opt,name=request_metadata,json=requestMetadata,proto3" json:"request_metadata,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *AuditLog) Reset() {
	*x = AuditLog{}
	mi := &file_store_audit_log_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuditLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditLog) ProtoMessage() {}

func (x *AuditLog) ProtoReflect() protoreflect.Message {
	mi := &file_store_audit_log_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditLog.ProtoReflect.Descriptor instead.
func (*AuditLog) Descriptor() ([]byte, []int) {
	return file_store_audit_log_proto_rawDescGZIP(), []int{0}
}

func (x *AuditLog) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *AuditLog) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *AuditLog) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *AuditLog) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *AuditLog) GetSeverity() AuditLog_Severity {
	if x != nil {
		return x.Severity
	}
	return AuditLog_DEFAULT
}

func (x *AuditLog) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

func (x *AuditLog) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *AuditLog) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *AuditLog) GetLatency() *durationpb.Duration {
	if x != nil {
		return x.Latency
	}
	return nil
}

func (x *AuditLog) GetServiceData() *anypb.Any {
	if x != nil {
		return x.ServiceData
	}
	return nil
}

func (x *AuditLog) GetRequestMetadata() *RequestMetadata {
	if x != nil {
		return x.RequestMetadata
	}
	return nil
}

// Metadata about the request.
type RequestMetadata struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The IP address of the caller.
	CallerIp string `protobuf:"bytes,1,opt,name=caller_ip,json=callerIp,proto3" json:"caller_ip,omitempty"`
	// The user agent of the caller.
	// This information is not authenticated and should be treated accordingly.
	CallerSuppliedUserAgent string `protobuf:"bytes,2,opt,name=caller_supplied_user_agent,json=callerSuppliedUserAgent,proto3" json:"caller_supplied_user_agent,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *RequestMetadata) Reset() {
	*x = RequestMetadata{}
	mi := &file_store_audit_log_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMetadata) ProtoMessage() {}

func (x *RequestMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_store_audit_log_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMetadata.ProtoReflect.Descriptor instead.
func (*RequestMetadata) Descriptor() ([]byte, []int) {
	return file_store_audit_log_proto_rawDescGZIP(), []int{1}
}

func (x *RequestMetadata) GetCallerIp() string {
	if x != nil {
		return x.CallerIp
	}
	return ""
}

func (x *RequestMetadata) GetCallerSuppliedUserAgent() string {
	if x != nil {
		return x.CallerSuppliedUserAgent
	}
	return ""
}

var File_store_audit_log_proto protoreflect.FileDescriptor

const file_store_audit_log_proto_rawDesc = "" +
	"\n" +
	"\x15store/audit_log.proto\x12\x0ebytebase.store\x1a\x19google/protobuf/any.proto\x1a\x1egoogle/protobuf/duration.proto\x1a\x17google/rpc/status.proto\"\xbf\x04\n" +
	"\bAuditLog\x12\x16\n" +
	"\x06parent\x18\x01 \x01(\tR\x06parent\x12\x16\n" +
	"\x06method\x18\x02 \x01(\tR\x06method\x12\x1a\n" +
	"\bresource\x18\x03 \x01(\tR\bresource\x12\x12\n" +
	"\x04user\x18\x04 \x01(\tR\x04user\x12=\n" +
	"\bseverity\x18\x05 \x01(\x0e2!.bytebase.store.AuditLog.SeverityR\bseverity\x12\x18\n" +
	"\arequest\x18\x06 \x01(\tR\arequest\x12\x1a\n" +
	"\bresponse\x18\a \x01(\tR\bresponse\x12*\n" +
	"\x06status\x18\b \x01(\v2\x12.google.rpc.StatusR\x06status\x123\n" +
	"\alatency\x18\t \x01(\v2\x19.google.protobuf.DurationR\alatency\x127\n" +
	"\fservice_data\x18\n" +
	" \x01(\v2\x14.google.protobuf.AnyR\vserviceData\x12J\n" +
	"\x10request_metadata\x18\v \x01(\v2\x1f.bytebase.store.RequestMetadataR\x0frequestMetadata\"x\n" +
	"\bSeverity\x12\v\n" +
	"\aDEFAULT\x10\x00\x12\t\n" +
	"\x05DEBUG\x10\x01\x12\b\n" +
	"\x04INFO\x10\x02\x12\n" +
	"\n" +
	"\x06NOTICE\x10\x03\x12\v\n" +
	"\aWARNING\x10\x04\x12\t\n" +
	"\x05ERROR\x10\x05\x12\f\n" +
	"\bCRITICAL\x10\x06\x12\t\n" +
	"\x05ALERT\x10\a\x12\r\n" +
	"\tEMERGENCY\x10\b\"k\n" +
	"\x0fRequestMetadata\x12\x1b\n" +
	"\tcaller_ip\x18\x01 \x01(\tR\bcallerIp\x12;\n" +
	"\x1acaller_supplied_user_agent\x18\x02 \x01(\tR\x17callerSuppliedUserAgentB\x14Z\x12generated-go/storeb\x06proto3"

var (
	file_store_audit_log_proto_rawDescOnce sync.Once
	file_store_audit_log_proto_rawDescData []byte
)

func file_store_audit_log_proto_rawDescGZIP() []byte {
	file_store_audit_log_proto_rawDescOnce.Do(func() {
		file_store_audit_log_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_store_audit_log_proto_rawDesc), len(file_store_audit_log_proto_rawDesc)))
	})
	return file_store_audit_log_proto_rawDescData
}

var file_store_audit_log_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_store_audit_log_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_store_audit_log_proto_goTypes = []any{
	(AuditLog_Severity)(0),      // 0: bytebase.store.AuditLog.Severity
	(*AuditLog)(nil),            // 1: bytebase.store.AuditLog
	(*RequestMetadata)(nil),     // 2: bytebase.store.RequestMetadata
	(*status.Status)(nil),       // 3: google.rpc.Status
	(*durationpb.Duration)(nil), // 4: google.protobuf.Duration
	(*anypb.Any)(nil),           // 5: google.protobuf.Any
}
var file_store_audit_log_proto_depIdxs = []int32{
	0, // 0: bytebase.store.AuditLog.severity:type_name -> bytebase.store.AuditLog.Severity
	3, // 1: bytebase.store.AuditLog.status:type_name -> google.rpc.Status
	4, // 2: bytebase.store.AuditLog.latency:type_name -> google.protobuf.Duration
	5, // 3: bytebase.store.AuditLog.service_data:type_name -> google.protobuf.Any
	2, // 4: bytebase.store.AuditLog.request_metadata:type_name -> bytebase.store.RequestMetadata
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_store_audit_log_proto_init() }
func file_store_audit_log_proto_init() {
	if File_store_audit_log_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_store_audit_log_proto_rawDesc), len(file_store_audit_log_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_audit_log_proto_goTypes,
		DependencyIndexes: file_store_audit_log_proto_depIdxs,
		EnumInfos:         file_store_audit_log_proto_enumTypes,
		MessageInfos:      file_store_audit_log_proto_msgTypes,
	}.Build()
	File_store_audit_log_proto = out.File
	file_store_audit_log_proto_goTypes = nil
	file_store_audit_log_proto_depIdxs = nil
}
