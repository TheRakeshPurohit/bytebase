// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: store/audit_log.proto

package store

import (
	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

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
}

func (x *AuditLog) Reset() {
	*x = AuditLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_audit_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuditLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditLog) ProtoMessage() {}

func (x *AuditLog) ProtoReflect() protoreflect.Message {
	mi := &file_store_audit_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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

var File_store_audit_log_proto protoreflect.FileDescriptor

var file_store_audit_log_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x85, 0x03, 0x0a, 0x08, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x3d, 0x0a,
	0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x21, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x2e, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69,
	0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x78,
	0x0a, 0x08, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45,
	0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x45, 0x42, 0x55, 0x47,
	0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06,
	0x4e, 0x4f, 0x54, 0x49, 0x43, 0x45, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x52, 0x4e,
	0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x05,
	0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x49, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x06, 0x12, 0x09,
	0x0a, 0x05, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x10, 0x07, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x4d, 0x45,
	0x52, 0x47, 0x45, 0x4e, 0x43, 0x59, 0x10, 0x08, 0x42, 0x14, 0x5a, 0x12, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_audit_log_proto_rawDescOnce sync.Once
	file_store_audit_log_proto_rawDescData = file_store_audit_log_proto_rawDesc
)

func file_store_audit_log_proto_rawDescGZIP() []byte {
	file_store_audit_log_proto_rawDescOnce.Do(func() {
		file_store_audit_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_audit_log_proto_rawDescData)
	})
	return file_store_audit_log_proto_rawDescData
}

var file_store_audit_log_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_store_audit_log_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_store_audit_log_proto_goTypes = []any{
	(AuditLog_Severity)(0), // 0: bytebase.store.AuditLog.Severity
	(*AuditLog)(nil),       // 1: bytebase.store.AuditLog
	(*status.Status)(nil),  // 2: google.rpc.Status
}
var file_store_audit_log_proto_depIdxs = []int32{
	0, // 0: bytebase.store.AuditLog.severity:type_name -> bytebase.store.AuditLog.Severity
	2, // 1: bytebase.store.AuditLog.status:type_name -> google.rpc.Status
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_store_audit_log_proto_init() }
func file_store_audit_log_proto_init() {
	if File_store_audit_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_store_audit_log_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AuditLog); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_store_audit_log_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_audit_log_proto_goTypes,
		DependencyIndexes: file_store_audit_log_proto_depIdxs,
		EnumInfos:         file_store_audit_log_proto_enumTypes,
		MessageInfos:      file_store_audit_log_proto_msgTypes,
	}.Build()
	File_store_audit_log_proto = out.File
	file_store_audit_log_proto_rawDesc = nil
	file_store_audit_log_proto_goTypes = nil
	file_store_audit_log_proto_depIdxs = nil
}
