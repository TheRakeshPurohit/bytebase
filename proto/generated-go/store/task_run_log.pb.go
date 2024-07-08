// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: store/task_run_log.proto

package store

import (
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

type TaskRunLog_Type int32

const (
	TaskRunLog_TYPE_UNSPECIFIED       TaskRunLog_Type = 0
	TaskRunLog_SCHEMA_DUMP_START      TaskRunLog_Type = 1
	TaskRunLog_SCHEMA_DUMP_END        TaskRunLog_Type = 2
	TaskRunLog_COMMAND_EXECUTE        TaskRunLog_Type = 3
	TaskRunLog_COMMAND_RESPONSE       TaskRunLog_Type = 4
	TaskRunLog_DATABASE_SYNC_START    TaskRunLog_Type = 5
	TaskRunLog_DATABASE_SYNC_END      TaskRunLog_Type = 6
	TaskRunLog_TASK_RUN_STATUS_UPDATE TaskRunLog_Type = 7
)

// Enum value maps for TaskRunLog_Type.
var (
	TaskRunLog_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "SCHEMA_DUMP_START",
		2: "SCHEMA_DUMP_END",
		3: "COMMAND_EXECUTE",
		4: "COMMAND_RESPONSE",
		5: "DATABASE_SYNC_START",
		6: "DATABASE_SYNC_END",
		7: "TASK_RUN_STATUS_UPDATE",
	}
	TaskRunLog_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED":       0,
		"SCHEMA_DUMP_START":      1,
		"SCHEMA_DUMP_END":        2,
		"COMMAND_EXECUTE":        3,
		"COMMAND_RESPONSE":       4,
		"DATABASE_SYNC_START":    5,
		"DATABASE_SYNC_END":      6,
		"TASK_RUN_STATUS_UPDATE": 7,
	}
)

func (x TaskRunLog_Type) Enum() *TaskRunLog_Type {
	p := new(TaskRunLog_Type)
	*p = x
	return p
}

func (x TaskRunLog_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskRunLog_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_store_task_run_log_proto_enumTypes[0].Descriptor()
}

func (TaskRunLog_Type) Type() protoreflect.EnumType {
	return &file_store_task_run_log_proto_enumTypes[0]
}

func (x TaskRunLog_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskRunLog_Type.Descriptor instead.
func (TaskRunLog_Type) EnumDescriptor() ([]byte, []int) {
	return file_store_task_run_log_proto_rawDescGZIP(), []int{0, 0}
}

type TaskRunLog_TaskRunStatusUpdate_Status int32

const (
	TaskRunLog_TaskRunStatusUpdate_STATUS_UNSPECIFIED TaskRunLog_TaskRunStatusUpdate_Status = 0
	// the task run is ready to be executed by the scheduler
	TaskRunLog_TaskRunStatusUpdate_RUNNING_WAITING TaskRunLog_TaskRunStatusUpdate_Status = 1
	// the task run is being executed by the scheduler
	TaskRunLog_TaskRunStatusUpdate_RUNNING_RUNNING TaskRunLog_TaskRunStatusUpdate_Status = 2
)

// Enum value maps for TaskRunLog_TaskRunStatusUpdate_Status.
var (
	TaskRunLog_TaskRunStatusUpdate_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "RUNNING_WAITING",
		2: "RUNNING_RUNNING",
	}
	TaskRunLog_TaskRunStatusUpdate_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"RUNNING_WAITING":    1,
		"RUNNING_RUNNING":    2,
	}
)

func (x TaskRunLog_TaskRunStatusUpdate_Status) Enum() *TaskRunLog_TaskRunStatusUpdate_Status {
	p := new(TaskRunLog_TaskRunStatusUpdate_Status)
	*p = x
	return p
}

func (x TaskRunLog_TaskRunStatusUpdate_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskRunLog_TaskRunStatusUpdate_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_store_task_run_log_proto_enumTypes[1].Descriptor()
}

func (TaskRunLog_TaskRunStatusUpdate_Status) Type() protoreflect.EnumType {
	return &file_store_task_run_log_proto_enumTypes[1]
}

func (x TaskRunLog_TaskRunStatusUpdate_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskRunLog_TaskRunStatusUpdate_Status.Descriptor instead.
func (TaskRunLog_TaskRunStatusUpdate_Status) EnumDescriptor() ([]byte, []int) {
	return file_store_task_run_log_proto_rawDescGZIP(), []int{0, 6, 0}
}

type TaskRunLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type                TaskRunLog_Type                 `protobuf:"varint,1,opt,name=type,proto3,enum=bytebase.store.TaskRunLog_Type" json:"type,omitempty"`
	SchemaDumpStart     *TaskRunLog_SchemaDumpStart     `protobuf:"bytes,2,opt,name=schema_dump_start,json=schemaDumpStart,proto3" json:"schema_dump_start,omitempty"`
	SchemaDumpEnd       *TaskRunLog_SchemaDumpEnd       `protobuf:"bytes,3,opt,name=schema_dump_end,json=schemaDumpEnd,proto3" json:"schema_dump_end,omitempty"`
	CommandExecute      *TaskRunLog_CommandExecute      `protobuf:"bytes,4,opt,name=command_execute,json=commandExecute,proto3" json:"command_execute,omitempty"`
	CommandResponse     *TaskRunLog_CommandResponse     `protobuf:"bytes,5,opt,name=command_response,json=commandResponse,proto3" json:"command_response,omitempty"`
	DatabaseSyncStart   *TaskRunLog_DatabaseSyncStart   `protobuf:"bytes,6,opt,name=database_sync_start,json=databaseSyncStart,proto3" json:"database_sync_start,omitempty"`
	DatabaseSyncEnd     *TaskRunLog_DatabaseSyncEnd     `protobuf:"bytes,7,opt,name=database_sync_end,json=databaseSyncEnd,proto3" json:"database_sync_end,omitempty"`
	TaskRunStatusUpdate *TaskRunLog_TaskRunStatusUpdate `protobuf:"bytes,8,opt,name=task_run_status_update,json=taskRunStatusUpdate,proto3" json:"task_run_status_update,omitempty"`
}

func (x *TaskRunLog) Reset() {
	*x = TaskRunLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_task_run_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRunLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRunLog) ProtoMessage() {}

func (x *TaskRunLog) ProtoReflect() protoreflect.Message {
	mi := &file_store_task_run_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRunLog.ProtoReflect.Descriptor instead.
func (*TaskRunLog) Descriptor() ([]byte, []int) {
	return file_store_task_run_log_proto_rawDescGZIP(), []int{0}
}

func (x *TaskRunLog) GetType() TaskRunLog_Type {
	if x != nil {
		return x.Type
	}
	return TaskRunLog_TYPE_UNSPECIFIED
}

func (x *TaskRunLog) GetSchemaDumpStart() *TaskRunLog_SchemaDumpStart {
	if x != nil {
		return x.SchemaDumpStart
	}
	return nil
}

func (x *TaskRunLog) GetSchemaDumpEnd() *TaskRunLog_SchemaDumpEnd {
	if x != nil {
		return x.SchemaDumpEnd
	}
	return nil
}

func (x *TaskRunLog) GetCommandExecute() *TaskRunLog_CommandExecute {
	if x != nil {
		return x.CommandExecute
	}
	return nil
}

func (x *TaskRunLog) GetCommandResponse() *TaskRunLog_CommandResponse {
	if x != nil {
		return x.CommandResponse
	}
	return nil
}

func (x *TaskRunLog) GetDatabaseSyncStart() *TaskRunLog_DatabaseSyncStart {
	if x != nil {
		return x.DatabaseSyncStart
	}
	return nil
}

func (x *TaskRunLog) GetDatabaseSyncEnd() *TaskRunLog_DatabaseSyncEnd {
	if x != nil {
		return x.DatabaseSyncEnd
	}
	return nil
}

func (x *TaskRunLog) GetTaskRunStatusUpdate() *TaskRunLog_TaskRunStatusUpdate {
	if x != nil {
		return x.TaskRunStatusUpdate
	}
	return nil
}

type TaskRunLog_SchemaDumpStart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TaskRunLog_SchemaDumpStart) Reset() {
	*x = TaskRunLog_SchemaDumpStart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_task_run_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRunLog_SchemaDumpStart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRunLog_SchemaDumpStart) ProtoMessage() {}

func (x *TaskRunLog_SchemaDumpStart) ProtoReflect() protoreflect.Message {
	mi := &file_store_task_run_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRunLog_SchemaDumpStart.ProtoReflect.Descriptor instead.
func (*TaskRunLog_SchemaDumpStart) Descriptor() ([]byte, []int) {
	return file_store_task_run_log_proto_rawDescGZIP(), []int{0, 0}
}

type TaskRunLog_SchemaDumpEnd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *TaskRunLog_SchemaDumpEnd) Reset() {
	*x = TaskRunLog_SchemaDumpEnd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_task_run_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRunLog_SchemaDumpEnd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRunLog_SchemaDumpEnd) ProtoMessage() {}

func (x *TaskRunLog_SchemaDumpEnd) ProtoReflect() protoreflect.Message {
	mi := &file_store_task_run_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRunLog_SchemaDumpEnd.ProtoReflect.Descriptor instead.
func (*TaskRunLog_SchemaDumpEnd) Descriptor() ([]byte, []int) {
	return file_store_task_run_log_proto_rawDescGZIP(), []int{0, 1}
}

func (x *TaskRunLog_SchemaDumpEnd) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type TaskRunLog_CommandExecute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The indexes of the executed commands.
	CommandIndexes []int32 `protobuf:"varint,1,rep,packed,name=command_indexes,json=commandIndexes,proto3" json:"command_indexes,omitempty"`
}

func (x *TaskRunLog_CommandExecute) Reset() {
	*x = TaskRunLog_CommandExecute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_task_run_log_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRunLog_CommandExecute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRunLog_CommandExecute) ProtoMessage() {}

func (x *TaskRunLog_CommandExecute) ProtoReflect() protoreflect.Message {
	mi := &file_store_task_run_log_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRunLog_CommandExecute.ProtoReflect.Descriptor instead.
func (*TaskRunLog_CommandExecute) Descriptor() ([]byte, []int) {
	return file_store_task_run_log_proto_rawDescGZIP(), []int{0, 2}
}

func (x *TaskRunLog_CommandExecute) GetCommandIndexes() []int32 {
	if x != nil {
		return x.CommandIndexes
	}
	return nil
}

type TaskRunLog_CommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The indexes of the executed commands.
	CommandIndexes []int32 `protobuf:"varint,1,rep,packed,name=command_indexes,json=commandIndexes,proto3" json:"command_indexes,omitempty"`
	Error          string  `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	AffectedRows   int32   `protobuf:"varint,3,opt,name=affected_rows,json=affectedRows,proto3" json:"affected_rows,omitempty"`
	// `all_affected_rows` is the affected rows of each command.
	// `all_affected_rows` may be unavailable if the database driver doesn't support it. Caller should fallback to `affected_rows` in that case.
	AllAffectedRows []int32 `protobuf:"varint,4,rep,packed,name=all_affected_rows,json=allAffectedRows,proto3" json:"all_affected_rows,omitempty"`
}

func (x *TaskRunLog_CommandResponse) Reset() {
	*x = TaskRunLog_CommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_task_run_log_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRunLog_CommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRunLog_CommandResponse) ProtoMessage() {}

func (x *TaskRunLog_CommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_store_task_run_log_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRunLog_CommandResponse.ProtoReflect.Descriptor instead.
func (*TaskRunLog_CommandResponse) Descriptor() ([]byte, []int) {
	return file_store_task_run_log_proto_rawDescGZIP(), []int{0, 3}
}

func (x *TaskRunLog_CommandResponse) GetCommandIndexes() []int32 {
	if x != nil {
		return x.CommandIndexes
	}
	return nil
}

func (x *TaskRunLog_CommandResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *TaskRunLog_CommandResponse) GetAffectedRows() int32 {
	if x != nil {
		return x.AffectedRows
	}
	return 0
}

func (x *TaskRunLog_CommandResponse) GetAllAffectedRows() []int32 {
	if x != nil {
		return x.AllAffectedRows
	}
	return nil
}

type TaskRunLog_DatabaseSyncStart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TaskRunLog_DatabaseSyncStart) Reset() {
	*x = TaskRunLog_DatabaseSyncStart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_task_run_log_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRunLog_DatabaseSyncStart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRunLog_DatabaseSyncStart) ProtoMessage() {}

func (x *TaskRunLog_DatabaseSyncStart) ProtoReflect() protoreflect.Message {
	mi := &file_store_task_run_log_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRunLog_DatabaseSyncStart.ProtoReflect.Descriptor instead.
func (*TaskRunLog_DatabaseSyncStart) Descriptor() ([]byte, []int) {
	return file_store_task_run_log_proto_rawDescGZIP(), []int{0, 4}
}

type TaskRunLog_DatabaseSyncEnd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *TaskRunLog_DatabaseSyncEnd) Reset() {
	*x = TaskRunLog_DatabaseSyncEnd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_task_run_log_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRunLog_DatabaseSyncEnd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRunLog_DatabaseSyncEnd) ProtoMessage() {}

func (x *TaskRunLog_DatabaseSyncEnd) ProtoReflect() protoreflect.Message {
	mi := &file_store_task_run_log_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRunLog_DatabaseSyncEnd.ProtoReflect.Descriptor instead.
func (*TaskRunLog_DatabaseSyncEnd) Descriptor() ([]byte, []int) {
	return file_store_task_run_log_proto_rawDescGZIP(), []int{0, 5}
}

func (x *TaskRunLog_DatabaseSyncEnd) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type TaskRunLog_TaskRunStatusUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status TaskRunLog_TaskRunStatusUpdate_Status `protobuf:"varint,1,opt,name=status,proto3,enum=bytebase.store.TaskRunLog_TaskRunStatusUpdate_Status" json:"status,omitempty"`
}

func (x *TaskRunLog_TaskRunStatusUpdate) Reset() {
	*x = TaskRunLog_TaskRunStatusUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_task_run_log_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRunLog_TaskRunStatusUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRunLog_TaskRunStatusUpdate) ProtoMessage() {}

func (x *TaskRunLog_TaskRunStatusUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_store_task_run_log_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRunLog_TaskRunStatusUpdate.ProtoReflect.Descriptor instead.
func (*TaskRunLog_TaskRunStatusUpdate) Descriptor() ([]byte, []int) {
	return file_store_task_run_log_proto_rawDescGZIP(), []int{0, 6}
}

func (x *TaskRunLog_TaskRunStatusUpdate) GetStatus() TaskRunLog_TaskRunStatusUpdate_Status {
	if x != nil {
		return x.Status
	}
	return TaskRunLog_TaskRunStatusUpdate_STATUS_UNSPECIFIED
}

var File_store_task_run_log_proto protoreflect.FileDescriptor

var file_store_task_run_log_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x72, 0x75, 0x6e,
	0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x79, 0x74, 0x65,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x22, 0xfd, 0x0a, 0x0a, 0x0a, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x4c, 0x6f, 0x67, 0x12, 0x33, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e,
	0x4c, 0x6f, 0x67, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x56,
	0x0a, 0x11, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x64, 0x75, 0x6d, 0x70, 0x5f, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x62, 0x79, 0x74, 0x65,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x75, 0x6e, 0x4c, 0x6f, 0x67, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x44, 0x75, 0x6d, 0x70,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x0f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x44, 0x75, 0x6d,
	0x70, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x50, 0x0a, 0x0f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x5f, 0x64, 0x75, 0x6d, 0x70, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x4c, 0x6f, 0x67, 0x2e, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x44, 0x75, 0x6d, 0x70, 0x45, 0x6e, 0x64, 0x52, 0x0d, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x44, 0x75, 0x6d, 0x70, 0x45, 0x6e, 0x64, 0x12, 0x52, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x4c, 0x6f, 0x67, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x0e, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x55, 0x0a, 0x10,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x4c,
	0x6f, 0x67, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x13, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x73, 0x79, 0x6e, 0x63, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2c, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x4c, 0x6f, 0x67, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x11,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x56, 0x0a, 0x11, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x79,
	0x6e, 0x63, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x62,
	0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x75, 0x6e, 0x4c, 0x6f, 0x67, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x53, 0x79, 0x6e, 0x63, 0x45, 0x6e, 0x64, 0x52, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x53, 0x79, 0x6e, 0x63, 0x45, 0x6e, 0x64, 0x12, 0x63, 0x0a, 0x16, 0x74, 0x61, 0x73,
	0x6b, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x62, 0x79, 0x74, 0x65,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x75, 0x6e, 0x4c, 0x6f, 0x67, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x13, 0x74, 0x61, 0x73, 0x6b, 0x52,
	0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x11,
	0x0a, 0x0f, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x44, 0x75, 0x6d, 0x70, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x1a, 0x25, 0x0a, 0x0d, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x44, 0x75, 0x6d, 0x70, 0x45,
	0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x39, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x05, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x65, 0x73, 0x1a, 0xa1, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x66, 0x66, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x5f, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x61,
	0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x52, 0x6f, 0x77, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x61,
	0x6c, 0x6c, 0x5f, 0x61, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x77, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0f, 0x61, 0x6c, 0x6c, 0x41, 0x66, 0x66, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x52, 0x6f, 0x77, 0x73, 0x1a, 0x13, 0x0a, 0x11, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x72, 0x74, 0x1a, 0x27, 0x0a, 0x0f,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x53, 0x79, 0x6e, 0x63, 0x45, 0x6e, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0xb0, 0x01, 0x0a, 0x13, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x75,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x4d, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e,
	0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x4c, 0x6f, 0x67, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x75,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x4a, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13,
	0x0a, 0x0f, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e,
	0x47, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x52,
	0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x22, 0xbf, 0x01, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x43, 0x48, 0x45, 0x4d,
	0x41, 0x5f, 0x44, 0x55, 0x4d, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x01, 0x12, 0x13,
	0x0a, 0x0f, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x5f, 0x44, 0x55, 0x4d, 0x50, 0x5f, 0x45, 0x4e,
	0x44, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x45,
	0x58, 0x45, 0x43, 0x55, 0x54, 0x45, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x4f, 0x4d, 0x4d,
	0x41, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x04, 0x12, 0x17,
	0x0a, 0x13, 0x44, 0x41, 0x54, 0x41, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x5f,
	0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x05, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x41, 0x54, 0x41, 0x42,
	0x41, 0x53, 0x45, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x45, 0x4e, 0x44, 0x10, 0x06, 0x12, 0x1a,
	0x0a, 0x16, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x52, 0x55, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x07, 0x42, 0x14, 0x5a, 0x12, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_task_run_log_proto_rawDescOnce sync.Once
	file_store_task_run_log_proto_rawDescData = file_store_task_run_log_proto_rawDesc
)

func file_store_task_run_log_proto_rawDescGZIP() []byte {
	file_store_task_run_log_proto_rawDescOnce.Do(func() {
		file_store_task_run_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_task_run_log_proto_rawDescData)
	})
	return file_store_task_run_log_proto_rawDescData
}

var file_store_task_run_log_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_store_task_run_log_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_store_task_run_log_proto_goTypes = []any{
	(TaskRunLog_Type)(0),                       // 0: bytebase.store.TaskRunLog.Type
	(TaskRunLog_TaskRunStatusUpdate_Status)(0), // 1: bytebase.store.TaskRunLog.TaskRunStatusUpdate.Status
	(*TaskRunLog)(nil),                         // 2: bytebase.store.TaskRunLog
	(*TaskRunLog_SchemaDumpStart)(nil),         // 3: bytebase.store.TaskRunLog.SchemaDumpStart
	(*TaskRunLog_SchemaDumpEnd)(nil),           // 4: bytebase.store.TaskRunLog.SchemaDumpEnd
	(*TaskRunLog_CommandExecute)(nil),          // 5: bytebase.store.TaskRunLog.CommandExecute
	(*TaskRunLog_CommandResponse)(nil),         // 6: bytebase.store.TaskRunLog.CommandResponse
	(*TaskRunLog_DatabaseSyncStart)(nil),       // 7: bytebase.store.TaskRunLog.DatabaseSyncStart
	(*TaskRunLog_DatabaseSyncEnd)(nil),         // 8: bytebase.store.TaskRunLog.DatabaseSyncEnd
	(*TaskRunLog_TaskRunStatusUpdate)(nil),     // 9: bytebase.store.TaskRunLog.TaskRunStatusUpdate
}
var file_store_task_run_log_proto_depIdxs = []int32{
	0, // 0: bytebase.store.TaskRunLog.type:type_name -> bytebase.store.TaskRunLog.Type
	3, // 1: bytebase.store.TaskRunLog.schema_dump_start:type_name -> bytebase.store.TaskRunLog.SchemaDumpStart
	4, // 2: bytebase.store.TaskRunLog.schema_dump_end:type_name -> bytebase.store.TaskRunLog.SchemaDumpEnd
	5, // 3: bytebase.store.TaskRunLog.command_execute:type_name -> bytebase.store.TaskRunLog.CommandExecute
	6, // 4: bytebase.store.TaskRunLog.command_response:type_name -> bytebase.store.TaskRunLog.CommandResponse
	7, // 5: bytebase.store.TaskRunLog.database_sync_start:type_name -> bytebase.store.TaskRunLog.DatabaseSyncStart
	8, // 6: bytebase.store.TaskRunLog.database_sync_end:type_name -> bytebase.store.TaskRunLog.DatabaseSyncEnd
	9, // 7: bytebase.store.TaskRunLog.task_run_status_update:type_name -> bytebase.store.TaskRunLog.TaskRunStatusUpdate
	1, // 8: bytebase.store.TaskRunLog.TaskRunStatusUpdate.status:type_name -> bytebase.store.TaskRunLog.TaskRunStatusUpdate.Status
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_store_task_run_log_proto_init() }
func file_store_task_run_log_proto_init() {
	if File_store_task_run_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_store_task_run_log_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TaskRunLog); i {
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
		file_store_task_run_log_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TaskRunLog_SchemaDumpStart); i {
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
		file_store_task_run_log_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*TaskRunLog_SchemaDumpEnd); i {
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
		file_store_task_run_log_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*TaskRunLog_CommandExecute); i {
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
		file_store_task_run_log_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*TaskRunLog_CommandResponse); i {
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
		file_store_task_run_log_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*TaskRunLog_DatabaseSyncStart); i {
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
		file_store_task_run_log_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*TaskRunLog_DatabaseSyncEnd); i {
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
		file_store_task_run_log_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*TaskRunLog_TaskRunStatusUpdate); i {
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
			RawDescriptor: file_store_task_run_log_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_task_run_log_proto_goTypes,
		DependencyIndexes: file_store_task_run_log_proto_depIdxs,
		EnumInfos:         file_store_task_run_log_proto_enumTypes,
		MessageInfos:      file_store_task_run_log_proto_msgTypes,
	}.Build()
	File_store_task_run_log_proto = out.File
	file_store_task_run_log_proto_rawDesc = nil
	file_store_task_run_log_proto_goTypes = nil
	file_store_task_run_log_proto_depIdxs = nil
}
