// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: v1/audit_log_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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
	return file_v1_audit_log_service_proto_enumTypes[0].Descriptor()
}

func (AuditLog_Severity) Type() protoreflect.EnumType {
	return &file_v1_audit_log_service_proto_enumTypes[0]
}

func (x AuditLog_Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuditLog_Severity.Descriptor instead.
func (AuditLog_Severity) EnumDescriptor() ([]byte, []int) {
	return file_v1_audit_log_service_proto_rawDescGZIP(), []int{4, 0}
}

type SearchAuditLogsRequest struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	Parent string                 `protobuf:"bytes,5,opt,name=parent,proto3" json:"parent,omitempty"`
	// The filter of the log. It should be a valid CEL expression.
	// The syntax and semantics of CEL are documented at https://github.com/google/cel-spec
	//
	// Supported filter:
	// - method: the API name, can be found in the docs, should start with "/bytebase.v1." prefix. For example "/bytebase.v1.UserService/CreateUser". Support "==" operator.
	// - severity: support "==" operator, check Severity enum in AuditLog message for values.
	// - user: the actor, should in "users/{email}" format, support "==" operator.
	// - create_time: support ">=" and "<=" operator.
	//
	// For example:
	//   - filter = "method == '/bytebase.v1.SQLService/Query'"
	//   - filter = "method == '/bytebase.v1.SQLService/Query' && severity == 'ERROR'"
	//   - filter = "method == '/bytebase.v1.SQLService/Query' && severity == 'ERROR' && user == 'users/bb@bytebase.com'"
	//   - filter = "method == '/bytebase.v1.SQLService/Query' && severity == 'ERROR' && create_time <= '2021-01-01T00:00:00Z' && create_time >= '2020-01-01T00:00:00Z'"
	Filter string `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	// The order by of the log.
	// Only support order by create_time.
	// For example:
	//   - order_by = "create_time asc"
	//   - order_by = "create_time desc"
	OrderBy string `protobuf:"bytes,2,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	// The maximum number of logs to return.
	// The service may return fewer than this value.
	// If unspecified, at most 10 log entries will be returned.
	// The maximum value is 5000; values above 5000 will be coerced to 5000.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A page token, received from a previous `SearchLogs` call.
	// Provide this to retrieve the subsequent page.
	PageToken     string `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchAuditLogsRequest) Reset() {
	*x = SearchAuditLogsRequest{}
	mi := &file_v1_audit_log_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchAuditLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchAuditLogsRequest) ProtoMessage() {}

func (x *SearchAuditLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_audit_log_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchAuditLogsRequest.ProtoReflect.Descriptor instead.
func (*SearchAuditLogsRequest) Descriptor() ([]byte, []int) {
	return file_v1_audit_log_service_proto_rawDescGZIP(), []int{0}
}

func (x *SearchAuditLogsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *SearchAuditLogsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *SearchAuditLogsRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *SearchAuditLogsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SearchAuditLogsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type SearchAuditLogsResponse struct {
	state     protoimpl.MessageState `protogen:"open.v1"`
	AuditLogs []*AuditLog            `protobuf:"bytes,1,rep,name=audit_logs,json=auditLogs,proto3" json:"audit_logs,omitempty"`
	// A token to retrieve next page of log entities.
	// Pass this value in the page_token field in the subsequent call
	// to retrieve the next page of log entities.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchAuditLogsResponse) Reset() {
	*x = SearchAuditLogsResponse{}
	mi := &file_v1_audit_log_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchAuditLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchAuditLogsResponse) ProtoMessage() {}

func (x *SearchAuditLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_audit_log_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchAuditLogsResponse.ProtoReflect.Descriptor instead.
func (*SearchAuditLogsResponse) Descriptor() ([]byte, []int) {
	return file_v1_audit_log_service_proto_rawDescGZIP(), []int{1}
}

func (x *SearchAuditLogsResponse) GetAuditLogs() []*AuditLog {
	if x != nil {
		return x.AuditLogs
	}
	return nil
}

func (x *SearchAuditLogsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type ExportAuditLogsRequest struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	Parent string                 `protobuf:"bytes,4,opt,name=parent,proto3" json:"parent,omitempty"`
	// The filter of the log. It should be a valid CEL expression.
	// Check the filter field in the SearchAuditLogsRequest message.
	Filter string `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	// The order by of the log.
	// Only support order by create_time.
	// For example:
	//   - order_by = "create_time asc"
	//   - order_by = "create_time desc"
	OrderBy string `protobuf:"bytes,2,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	// The export format.
	Format ExportFormat `protobuf:"varint,3,opt,name=format,proto3,enum=bytebase.v1.ExportFormat" json:"format,omitempty"`
	// The maximum number of logs to return.
	// The service may return fewer than this value.
	// If unspecified, at most 10 log entries will be returned.
	// The maximum value is 5000; values above 5000 will be coerced to 5000.
	PageSize int32 `protobuf:"varint,5,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A page token, received from a previous `ExportAuditLogs` call.
	// Provide this to retrieve the subsequent page.
	PageToken     string `protobuf:"bytes,6,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExportAuditLogsRequest) Reset() {
	*x = ExportAuditLogsRequest{}
	mi := &file_v1_audit_log_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExportAuditLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportAuditLogsRequest) ProtoMessage() {}

func (x *ExportAuditLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_audit_log_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportAuditLogsRequest.ProtoReflect.Descriptor instead.
func (*ExportAuditLogsRequest) Descriptor() ([]byte, []int) {
	return file_v1_audit_log_service_proto_rawDescGZIP(), []int{2}
}

func (x *ExportAuditLogsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ExportAuditLogsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *ExportAuditLogsRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *ExportAuditLogsRequest) GetFormat() ExportFormat {
	if x != nil {
		return x.Format
	}
	return ExportFormat_FORMAT_UNSPECIFIED
}

func (x *ExportAuditLogsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ExportAuditLogsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ExportAuditLogsResponse struct {
	state   protoimpl.MessageState `protogen:"open.v1"`
	Content []byte                 `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	// A token to retrieve next page of log entities.
	// Pass this value in the page_token field in the subsequent call
	// to retrieve the next page of log entities.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExportAuditLogsResponse) Reset() {
	*x = ExportAuditLogsResponse{}
	mi := &file_v1_audit_log_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExportAuditLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportAuditLogsResponse) ProtoMessage() {}

func (x *ExportAuditLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_audit_log_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportAuditLogsResponse.ProtoReflect.Descriptor instead.
func (*ExportAuditLogsResponse) Descriptor() ([]byte, []int) {
	return file_v1_audit_log_service_proto_rawDescGZIP(), []int{3}
}

func (x *ExportAuditLogsResponse) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *ExportAuditLogsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type AuditLog struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the log.
	// Formats:
	// - projects/{project}/auditLogs/{uid}
	// - workspaces/{workspace}/auditLogs/{uid}
	Name       string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Format: users/d@d.com
	User string `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	// e.g. `/bytebase.v1.SQLService/Query`, `bb.project.repository.push`
	Method   string            `protobuf:"bytes,4,opt,name=method,proto3" json:"method,omitempty"`
	Severity AuditLog_Severity `protobuf:"varint,5,opt,name=severity,proto3,enum=bytebase.v1.AuditLog_Severity" json:"severity,omitempty"`
	// The associated resource.
	Resource string `protobuf:"bytes,6,opt,name=resource,proto3" json:"resource,omitempty"`
	// JSON-encoded request.
	Request string `protobuf:"bytes,7,opt,name=request,proto3" json:"request,omitempty"`
	// JSON-encoded response.
	// Some fields are omitted because they are too large or contain sensitive information.
	Response string         `protobuf:"bytes,8,opt,name=response,proto3" json:"response,omitempty"`
	Status   *status.Status `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	// The latency of the RPC.
	Latency *durationpb.Duration `protobuf:"bytes,10,opt,name=latency,proto3" json:"latency,omitempty"`
	// service-specific data about the request, response, and other activities.
	ServiceData *anypb.Any `protobuf:"bytes,11,opt,name=service_data,json=serviceData,proto3" json:"service_data,omitempty"`
	// Metadata about the operation.
	RequestMetadata *RequestMetadata `protobuf:"bytes,12,opt,name=request_metadata,json=requestMetadata,proto3" json:"request_metadata,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *AuditLog) Reset() {
	*x = AuditLog{}
	mi := &file_v1_audit_log_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuditLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditLog) ProtoMessage() {}

func (x *AuditLog) ProtoReflect() protoreflect.Message {
	mi := &file_v1_audit_log_service_proto_msgTypes[4]
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
	return file_v1_audit_log_service_proto_rawDescGZIP(), []int{4}
}

func (x *AuditLog) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AuditLog) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *AuditLog) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *AuditLog) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *AuditLog) GetSeverity() AuditLog_Severity {
	if x != nil {
		return x.Severity
	}
	return AuditLog_DEFAULT
}

func (x *AuditLog) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
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

type AuditData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PolicyDelta   *PolicyDelta           `protobuf:"bytes,1,opt,name=policy_delta,json=policyDelta,proto3" json:"policy_delta,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuditData) Reset() {
	*x = AuditData{}
	mi := &file_v1_audit_log_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuditData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditData) ProtoMessage() {}

func (x *AuditData) ProtoReflect() protoreflect.Message {
	mi := &file_v1_audit_log_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditData.ProtoReflect.Descriptor instead.
func (*AuditData) Descriptor() ([]byte, []int) {
	return file_v1_audit_log_service_proto_rawDescGZIP(), []int{5}
}

func (x *AuditData) GetPolicyDelta() *PolicyDelta {
	if x != nil {
		return x.PolicyDelta
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
	mi := &file_v1_audit_log_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMetadata) ProtoMessage() {}

func (x *RequestMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_v1_audit_log_service_proto_msgTypes[6]
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
	return file_v1_audit_log_service_proto_rawDescGZIP(), []int{6}
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

var File_v1_audit_log_service_proto protoreflect.FileDescriptor

const file_v1_audit_log_service_proto_rawDesc = "" +
	"\n" +
	"\x1av1/audit_log_service.proto\x12\vbytebase.v1\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\x1a\x19google/protobuf/any.proto\x1a\x1egoogle/protobuf/duration.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x17google/rpc/status.proto\x1a\x13v1/annotation.proto\x1a\x0fv1/common.proto\x1a\x13v1/iam_policy.proto\"\xbe\x01\n" +
	"\x16SearchAuditLogsRequest\x125\n" +
	"\x06parent\x18\x05 \x01(\tB\x1d\xe0A\x02\xfaA\x17\x12\x15bytebase.com/AuditLogR\x06parent\x12\x16\n" +
	"\x06filter\x18\x01 \x01(\tR\x06filter\x12\x19\n" +
	"\border_by\x18\x02 \x01(\tR\aorderBy\x12\x1b\n" +
	"\tpage_size\x18\x03 \x01(\x05R\bpageSize\x12\x1d\n" +
	"\n" +
	"page_token\x18\x04 \x01(\tR\tpageToken\"w\n" +
	"\x17SearchAuditLogsResponse\x124\n" +
	"\n" +
	"audit_logs\x18\x01 \x03(\v2\x15.bytebase.v1.AuditLogR\tauditLogs\x12&\n" +
	"\x0fnext_page_token\x18\x02 \x01(\tR\rnextPageToken\"\xf1\x01\n" +
	"\x16ExportAuditLogsRequest\x125\n" +
	"\x06parent\x18\x04 \x01(\tB\x1d\xe0A\x02\xfaA\x17\x12\x15bytebase.com/AuditLogR\x06parent\x12\x16\n" +
	"\x06filter\x18\x01 \x01(\tR\x06filter\x12\x19\n" +
	"\border_by\x18\x02 \x01(\tR\aorderBy\x121\n" +
	"\x06format\x18\x03 \x01(\x0e2\x19.bytebase.v1.ExportFormatR\x06format\x12\x1b\n" +
	"\tpage_size\x18\x05 \x01(\x05R\bpageSize\x12\x1d\n" +
	"\n" +
	"page_token\x18\x06 \x01(\tR\tpageToken\"[\n" +
	"\x17ExportAuditLogsResponse\x12\x18\n" +
	"\acontent\x18\x01 \x01(\fR\acontent\x12&\n" +
	"\x0fnext_page_token\x18\x02 \x01(\tR\rnextPageToken\"\xfc\x04\n" +
	"\bAuditLog\x12\x17\n" +
	"\x04name\x18\x01 \x01(\tB\x03\xe0A\x03R\x04name\x12@\n" +
	"\vcreate_time\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampB\x03\xe0A\x03R\n" +
	"createTime\x12\x12\n" +
	"\x04user\x18\x03 \x01(\tR\x04user\x12\x16\n" +
	"\x06method\x18\x04 \x01(\tR\x06method\x12:\n" +
	"\bseverity\x18\x05 \x01(\x0e2\x1e.bytebase.v1.AuditLog.SeverityR\bseverity\x12\x1a\n" +
	"\bresource\x18\x06 \x01(\tR\bresource\x12\x18\n" +
	"\arequest\x18\a \x01(\tR\arequest\x12\x1a\n" +
	"\bresponse\x18\b \x01(\tR\bresponse\x12*\n" +
	"\x06status\x18\t \x01(\v2\x12.google.rpc.StatusR\x06status\x123\n" +
	"\alatency\x18\n" +
	" \x01(\v2\x19.google.protobuf.DurationR\alatency\x127\n" +
	"\fservice_data\x18\v \x01(\v2\x14.google.protobuf.AnyR\vserviceData\x12G\n" +
	"\x10request_metadata\x18\f \x01(\v2\x1c.bytebase.v1.RequestMetadataR\x0frequestMetadata\"x\n" +
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
	"\tEMERGENCY\x10\b\"H\n" +
	"\tAuditData\x12;\n" +
	"\fpolicy_delta\x18\x01 \x01(\v2\x18.bytebase.v1.PolicyDeltaR\vpolicyDelta\"k\n" +
	"\x0fRequestMetadata\x12\x1b\n" +
	"\tcaller_ip\x18\x01 \x01(\tR\bcallerIp\x12;\n" +
	"\x1acaller_supplied_user_agent\x18\x02 \x01(\tR\x17callerSuppliedUserAgent2\xa5\x03\n" +
	"\x0fAuditLogService\x12\xc7\x01\n" +
	"\x0fSearchAuditLogs\x12#.bytebase.v1.SearchAuditLogsRequest\x1a$.bytebase.v1.SearchAuditLogsResponse\"i\x8a\xea0\x13bb.auditLogs.search\x90\xea0\x01\x82\xd3\xe4\x93\x02H:\x01*Z\x19:\x01*\"\x14/v1/auditLogs:search\"(/v1/{parent=projects/*}/auditLogs:search\x12\xc7\x01\n" +
	"\x0fExportAuditLogs\x12#.bytebase.v1.ExportAuditLogsRequest\x1a$.bytebase.v1.ExportAuditLogsResponse\"i\x8a\xea0\x13bb.auditLogs.export\x90\xea0\x01\x82\xd3\xe4\x93\x02H:\x01*Z\x19:\x01*\"\x14/v1/auditLogs:export\"(/v1/{parent=projects/*}/auditLogs:exportB6Z4github.com/bytebase/bytebase/backend/generated-go/v1b\x06proto3"

var (
	file_v1_audit_log_service_proto_rawDescOnce sync.Once
	file_v1_audit_log_service_proto_rawDescData []byte
)

func file_v1_audit_log_service_proto_rawDescGZIP() []byte {
	file_v1_audit_log_service_proto_rawDescOnce.Do(func() {
		file_v1_audit_log_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_v1_audit_log_service_proto_rawDesc), len(file_v1_audit_log_service_proto_rawDesc)))
	})
	return file_v1_audit_log_service_proto_rawDescData
}

var file_v1_audit_log_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_audit_log_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_v1_audit_log_service_proto_goTypes = []any{
	(AuditLog_Severity)(0),          // 0: bytebase.v1.AuditLog.Severity
	(*SearchAuditLogsRequest)(nil),  // 1: bytebase.v1.SearchAuditLogsRequest
	(*SearchAuditLogsResponse)(nil), // 2: bytebase.v1.SearchAuditLogsResponse
	(*ExportAuditLogsRequest)(nil),  // 3: bytebase.v1.ExportAuditLogsRequest
	(*ExportAuditLogsResponse)(nil), // 4: bytebase.v1.ExportAuditLogsResponse
	(*AuditLog)(nil),                // 5: bytebase.v1.AuditLog
	(*AuditData)(nil),               // 6: bytebase.v1.AuditData
	(*RequestMetadata)(nil),         // 7: bytebase.v1.RequestMetadata
	(ExportFormat)(0),               // 8: bytebase.v1.ExportFormat
	(*timestamppb.Timestamp)(nil),   // 9: google.protobuf.Timestamp
	(*status.Status)(nil),           // 10: google.rpc.Status
	(*durationpb.Duration)(nil),     // 11: google.protobuf.Duration
	(*anypb.Any)(nil),               // 12: google.protobuf.Any
	(*PolicyDelta)(nil),             // 13: bytebase.v1.PolicyDelta
}
var file_v1_audit_log_service_proto_depIdxs = []int32{
	5,  // 0: bytebase.v1.SearchAuditLogsResponse.audit_logs:type_name -> bytebase.v1.AuditLog
	8,  // 1: bytebase.v1.ExportAuditLogsRequest.format:type_name -> bytebase.v1.ExportFormat
	9,  // 2: bytebase.v1.AuditLog.create_time:type_name -> google.protobuf.Timestamp
	0,  // 3: bytebase.v1.AuditLog.severity:type_name -> bytebase.v1.AuditLog.Severity
	10, // 4: bytebase.v1.AuditLog.status:type_name -> google.rpc.Status
	11, // 5: bytebase.v1.AuditLog.latency:type_name -> google.protobuf.Duration
	12, // 6: bytebase.v1.AuditLog.service_data:type_name -> google.protobuf.Any
	7,  // 7: bytebase.v1.AuditLog.request_metadata:type_name -> bytebase.v1.RequestMetadata
	13, // 8: bytebase.v1.AuditData.policy_delta:type_name -> bytebase.v1.PolicyDelta
	1,  // 9: bytebase.v1.AuditLogService.SearchAuditLogs:input_type -> bytebase.v1.SearchAuditLogsRequest
	3,  // 10: bytebase.v1.AuditLogService.ExportAuditLogs:input_type -> bytebase.v1.ExportAuditLogsRequest
	2,  // 11: bytebase.v1.AuditLogService.SearchAuditLogs:output_type -> bytebase.v1.SearchAuditLogsResponse
	4,  // 12: bytebase.v1.AuditLogService.ExportAuditLogs:output_type -> bytebase.v1.ExportAuditLogsResponse
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_v1_audit_log_service_proto_init() }
func file_v1_audit_log_service_proto_init() {
	if File_v1_audit_log_service_proto != nil {
		return
	}
	file_v1_annotation_proto_init()
	file_v1_common_proto_init()
	file_v1_iam_policy_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_v1_audit_log_service_proto_rawDesc), len(file_v1_audit_log_service_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_audit_log_service_proto_goTypes,
		DependencyIndexes: file_v1_audit_log_service_proto_depIdxs,
		EnumInfos:         file_v1_audit_log_service_proto_enumTypes,
		MessageInfos:      file_v1_audit_log_service_proto_msgTypes,
	}.Build()
	File_v1_audit_log_service_proto = out.File
	file_v1_audit_log_service_proto_goTypes = nil
	file_v1_audit_log_service_proto_depIdxs = nil
}
