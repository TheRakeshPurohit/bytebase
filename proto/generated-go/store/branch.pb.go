// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: store/branch.proto

package store

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BranchSnapshot struct {
	state          protoimpl.MessageState  `protogen:"open.v1"`
	Metadata       *DatabaseSchemaMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	DatabaseConfig *BranchDatabaseConfig   `protobuf:"bytes,2,opt,name=database_config,json=databaseConfig,proto3" json:"database_config,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *BranchSnapshot) Reset() {
	*x = BranchSnapshot{}
	mi := &file_store_branch_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BranchSnapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BranchSnapshot) ProtoMessage() {}

func (x *BranchSnapshot) ProtoReflect() protoreflect.Message {
	mi := &file_store_branch_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BranchSnapshot.ProtoReflect.Descriptor instead.
func (*BranchSnapshot) Descriptor() ([]byte, []int) {
	return file_store_branch_proto_rawDescGZIP(), []int{0}
}

func (x *BranchSnapshot) GetMetadata() *DatabaseSchemaMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *BranchSnapshot) GetDatabaseConfig() *BranchDatabaseConfig {
	if x != nil {
		return x.DatabaseConfig
	}
	return nil
}

type BranchConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of source database.
	// Optional.
	// Example: instances/instance-id/databases/database-name.
	SourceDatabase string `protobuf:"bytes,1,opt,name=source_database,json=sourceDatabase,proto3" json:"source_database,omitempty"`
	// The name of the source branch.
	// Optional.
	// Example: projects/project-id/branches/branch-id.
	SourceBranch  string `protobuf:"bytes,2,opt,name=source_branch,json=sourceBranch,proto3" json:"source_branch,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BranchConfig) Reset() {
	*x = BranchConfig{}
	mi := &file_store_branch_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BranchConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BranchConfig) ProtoMessage() {}

func (x *BranchConfig) ProtoReflect() protoreflect.Message {
	mi := &file_store_branch_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BranchConfig.ProtoReflect.Descriptor instead.
func (*BranchConfig) Descriptor() ([]byte, []int) {
	return file_store_branch_proto_rawDescGZIP(), []int{1}
}

func (x *BranchConfig) GetSourceDatabase() string {
	if x != nil {
		return x.SourceDatabase
	}
	return ""
}

func (x *BranchConfig) GetSourceBranch() string {
	if x != nil {
		return x.SourceBranch
	}
	return ""
}

type BranchDatabaseConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Name  string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The schema_configs is the list of configs for schemas in a database.
	SchemaConfigs []*BranchSchemaConfig `protobuf:"bytes,2,rep,name=schema_configs,json=schemaConfigs,proto3" json:"schema_configs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BranchDatabaseConfig) Reset() {
	*x = BranchDatabaseConfig{}
	mi := &file_store_branch_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BranchDatabaseConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BranchDatabaseConfig) ProtoMessage() {}

func (x *BranchDatabaseConfig) ProtoReflect() protoreflect.Message {
	mi := &file_store_branch_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BranchDatabaseConfig.ProtoReflect.Descriptor instead.
func (*BranchDatabaseConfig) Descriptor() ([]byte, []int) {
	return file_store_branch_proto_rawDescGZIP(), []int{2}
}

func (x *BranchDatabaseConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BranchDatabaseConfig) GetSchemaConfigs() []*BranchSchemaConfig {
	if x != nil {
		return x.SchemaConfigs
	}
	return nil
}

type BranchSchemaConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name is the schema name.
	// It is an empty string for databases without such concept such as MySQL.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The table_configs is the list of configs for tables in a schema.
	TableConfigs     []*BranchTableConfig     `protobuf:"bytes,2,rep,name=table_configs,json=tableConfigs,proto3" json:"table_configs,omitempty"`
	FunctionConfigs  []*BranchFunctionConfig  `protobuf:"bytes,3,rep,name=function_configs,json=functionConfigs,proto3" json:"function_configs,omitempty"`
	ProcedureConfigs []*BranchProcedureConfig `protobuf:"bytes,4,rep,name=procedure_configs,json=procedureConfigs,proto3" json:"procedure_configs,omitempty"`
	ViewConfigs      []*BranchViewConfig      `protobuf:"bytes,5,rep,name=view_configs,json=viewConfigs,proto3" json:"view_configs,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *BranchSchemaConfig) Reset() {
	*x = BranchSchemaConfig{}
	mi := &file_store_branch_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BranchSchemaConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BranchSchemaConfig) ProtoMessage() {}

func (x *BranchSchemaConfig) ProtoReflect() protoreflect.Message {
	mi := &file_store_branch_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BranchSchemaConfig.ProtoReflect.Descriptor instead.
func (*BranchSchemaConfig) Descriptor() ([]byte, []int) {
	return file_store_branch_proto_rawDescGZIP(), []int{3}
}

func (x *BranchSchemaConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BranchSchemaConfig) GetTableConfigs() []*BranchTableConfig {
	if x != nil {
		return x.TableConfigs
	}
	return nil
}

func (x *BranchSchemaConfig) GetFunctionConfigs() []*BranchFunctionConfig {
	if x != nil {
		return x.FunctionConfigs
	}
	return nil
}

func (x *BranchSchemaConfig) GetProcedureConfigs() []*BranchProcedureConfig {
	if x != nil {
		return x.ProcedureConfigs
	}
	return nil
}

func (x *BranchSchemaConfig) GetViewConfigs() []*BranchViewConfig {
	if x != nil {
		return x.ViewConfigs
	}
	return nil
}

type BranchTableConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name is the name of a table.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The column_configs is the ordered list of configs for columns in a table.
	ColumnConfigs    []*BranchColumnConfig `protobuf:"bytes,2,rep,name=column_configs,json=columnConfigs,proto3" json:"column_configs,omitempty"`
	ClassificationId string                `protobuf:"bytes,3,opt,name=classification_id,json=classificationId,proto3" json:"classification_id,omitempty"`
	// The last updater of the table in branch.
	// Format: users/{userUID}.
	Updater string `protobuf:"bytes,4,opt,name=updater,proto3" json:"updater,omitempty"`
	// The last change come from branch.
	// Format: projcets/{project}/branches/{branch}
	SourceBranch string `protobuf:"bytes,6,opt,name=source_branch,json=sourceBranch,proto3" json:"source_branch,omitempty"`
	// The timestamp when the table is updated in branch.
	UpdateTime    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BranchTableConfig) Reset() {
	*x = BranchTableConfig{}
	mi := &file_store_branch_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BranchTableConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BranchTableConfig) ProtoMessage() {}

func (x *BranchTableConfig) ProtoReflect() protoreflect.Message {
	mi := &file_store_branch_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BranchTableConfig.ProtoReflect.Descriptor instead.
func (*BranchTableConfig) Descriptor() ([]byte, []int) {
	return file_store_branch_proto_rawDescGZIP(), []int{4}
}

func (x *BranchTableConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BranchTableConfig) GetColumnConfigs() []*BranchColumnConfig {
	if x != nil {
		return x.ColumnConfigs
	}
	return nil
}

func (x *BranchTableConfig) GetClassificationId() string {
	if x != nil {
		return x.ClassificationId
	}
	return ""
}

func (x *BranchTableConfig) GetUpdater() string {
	if x != nil {
		return x.Updater
	}
	return ""
}

func (x *BranchTableConfig) GetSourceBranch() string {
	if x != nil {
		return x.SourceBranch
	}
	return ""
}

func (x *BranchTableConfig) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

type BranchFunctionConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name is the name of a function.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The last updater of the function in branch.
	// Format: users/{userUID}.
	Updater string `protobuf:"bytes,2,opt,name=updater,proto3" json:"updater,omitempty"`
	// The last change come from branch.
	// Format: projcets/{project}/branches/{branch}
	SourceBranch string `protobuf:"bytes,4,opt,name=source_branch,json=sourceBranch,proto3" json:"source_branch,omitempty"`
	// The timestamp when the function is updated in branch.
	UpdateTime    *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BranchFunctionConfig) Reset() {
	*x = BranchFunctionConfig{}
	mi := &file_store_branch_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BranchFunctionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BranchFunctionConfig) ProtoMessage() {}

func (x *BranchFunctionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_store_branch_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BranchFunctionConfig.ProtoReflect.Descriptor instead.
func (*BranchFunctionConfig) Descriptor() ([]byte, []int) {
	return file_store_branch_proto_rawDescGZIP(), []int{5}
}

func (x *BranchFunctionConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BranchFunctionConfig) GetUpdater() string {
	if x != nil {
		return x.Updater
	}
	return ""
}

func (x *BranchFunctionConfig) GetSourceBranch() string {
	if x != nil {
		return x.SourceBranch
	}
	return ""
}

func (x *BranchFunctionConfig) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

type BranchProcedureConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name is the name of a procedure.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The last updater of the procedure in branch.
	// Format: users/{userUID}.
	Updater string `protobuf:"bytes,2,opt,name=updater,proto3" json:"updater,omitempty"`
	// The last change come from branch.
	// Format: projcets/{project}/branches/{branch}
	SourceBranch string `protobuf:"bytes,4,opt,name=source_branch,json=sourceBranch,proto3" json:"source_branch,omitempty"`
	// The timestamp when the procedure is updated in branch.
	UpdateTime    *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BranchProcedureConfig) Reset() {
	*x = BranchProcedureConfig{}
	mi := &file_store_branch_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BranchProcedureConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BranchProcedureConfig) ProtoMessage() {}

func (x *BranchProcedureConfig) ProtoReflect() protoreflect.Message {
	mi := &file_store_branch_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BranchProcedureConfig.ProtoReflect.Descriptor instead.
func (*BranchProcedureConfig) Descriptor() ([]byte, []int) {
	return file_store_branch_proto_rawDescGZIP(), []int{6}
}

func (x *BranchProcedureConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BranchProcedureConfig) GetUpdater() string {
	if x != nil {
		return x.Updater
	}
	return ""
}

func (x *BranchProcedureConfig) GetSourceBranch() string {
	if x != nil {
		return x.SourceBranch
	}
	return ""
}

func (x *BranchProcedureConfig) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

type BranchViewConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name is the name of a view.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The last updater of the view in branch.
	// Format: users/{userUID}.
	Updater string `protobuf:"bytes,2,opt,name=updater,proto3" json:"updater,omitempty"`
	// The last change come from branch.
	// Format: projcets/{project}/branches/{branch}
	SourceBranch string `protobuf:"bytes,4,opt,name=source_branch,json=sourceBranch,proto3" json:"source_branch,omitempty"`
	// The timestamp when the view is updated in branch.
	UpdateTime    *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BranchViewConfig) Reset() {
	*x = BranchViewConfig{}
	mi := &file_store_branch_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BranchViewConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BranchViewConfig) ProtoMessage() {}

func (x *BranchViewConfig) ProtoReflect() protoreflect.Message {
	mi := &file_store_branch_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BranchViewConfig.ProtoReflect.Descriptor instead.
func (*BranchViewConfig) Descriptor() ([]byte, []int) {
	return file_store_branch_proto_rawDescGZIP(), []int{7}
}

func (x *BranchViewConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BranchViewConfig) GetUpdater() string {
	if x != nil {
		return x.Updater
	}
	return ""
}

func (x *BranchViewConfig) GetSourceBranch() string {
	if x != nil {
		return x.SourceBranch
	}
	return ""
}

func (x *BranchViewConfig) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

type BranchColumnConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name is the name of a column.
	Name           string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SemanticTypeId string `protobuf:"bytes,2,opt,name=semantic_type_id,json=semanticTypeId,proto3" json:"semantic_type_id,omitempty"`
	// The user labels for a column.
	Labels                    map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ClassificationId          string            `protobuf:"bytes,4,opt,name=classification_id,json=classificationId,proto3" json:"classification_id,omitempty"`
	MaskingLevel              MaskingLevel      `protobuf:"varint,5,opt,name=masking_level,json=maskingLevel,proto3,enum=bytebase.store.MaskingLevel" json:"masking_level,omitempty"`
	FullMaskingAlgorithmId    string            `protobuf:"bytes,6,opt,name=full_masking_algorithm_id,json=fullMaskingAlgorithmId,proto3" json:"full_masking_algorithm_id,omitempty"`
	PartialMaskingAlgorithmId string            `protobuf:"bytes,7,opt,name=partial_masking_algorithm_id,json=partialMaskingAlgorithmId,proto3" json:"partial_masking_algorithm_id,omitempty"`
	unknownFields             protoimpl.UnknownFields
	sizeCache                 protoimpl.SizeCache
}

func (x *BranchColumnConfig) Reset() {
	*x = BranchColumnConfig{}
	mi := &file_store_branch_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BranchColumnConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BranchColumnConfig) ProtoMessage() {}

func (x *BranchColumnConfig) ProtoReflect() protoreflect.Message {
	mi := &file_store_branch_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BranchColumnConfig.ProtoReflect.Descriptor instead.
func (*BranchColumnConfig) Descriptor() ([]byte, []int) {
	return file_store_branch_proto_rawDescGZIP(), []int{8}
}

func (x *BranchColumnConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BranchColumnConfig) GetSemanticTypeId() string {
	if x != nil {
		return x.SemanticTypeId
	}
	return ""
}

func (x *BranchColumnConfig) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *BranchColumnConfig) GetClassificationId() string {
	if x != nil {
		return x.ClassificationId
	}
	return ""
}

func (x *BranchColumnConfig) GetMaskingLevel() MaskingLevel {
	if x != nil {
		return x.MaskingLevel
	}
	return MaskingLevel_MASKING_LEVEL_UNSPECIFIED
}

func (x *BranchColumnConfig) GetFullMaskingAlgorithmId() string {
	if x != nil {
		return x.FullMaskingAlgorithmId
	}
	return ""
}

func (x *BranchColumnConfig) GetPartialMaskingAlgorithmId() string {
	if x != nil {
		return x.PartialMaskingAlgorithmId
	}
	return ""
}

var File_store_branch_proto protoreflect.FileDescriptor

var file_store_branch_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa3, 0x01, 0x0a, 0x0e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x53, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x12, 0x42, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x4d, 0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x5c, 0x0a, 0x0c, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x22, 0x7b, 0x0a, 0x14, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x49, 0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0d, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x4a, 0x04, 0x08, 0x03, 0x10,
	0x04, 0x22, 0xda, 0x02, 0x0a, 0x12, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x0d,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0c, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x73, 0x12, 0x4f, 0x0a, 0x10, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x0f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x52, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75,
	0x72, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75,
	0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x43, 0x0a, 0x0c, 0x76, 0x69, 0x65,
	0x77, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x56, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x0b, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x22, 0xad,
	0x02, 0x0a, 0x11, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x49, 0x0a, 0x0e, 0x63, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x0d, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x1e, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72,
	0x12, 0x29, 0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0c, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x41, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41,
	0x01, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xb8,
	0x01, 0x0a, 0x14, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x07, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41,
	0x01, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x0d, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x41, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xb9, 0x01, 0x0a, 0x15, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x07,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04,
	0xe2, 0x41, 0x01, 0x03, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x12, 0x41, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xb4, 0x01, 0x0a, 0x10, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x56, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e,
	0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x12, 0x29,
	0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0c, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x41, 0x0a, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03,
	0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xc1, 0x03, 0x0a,
	0x12, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x65, 0x6d, 0x61, 0x6e,
	0x74, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x49,
	0x64, 0x12, 0x46, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x0d, 0x6d, 0x61, 0x73, 0x6b, 0x69, 0x6e,
	0x67, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e,
	0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4d,
	0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x0c, 0x6d, 0x61, 0x73,
	0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x39, 0x0a, 0x19, 0x66, 0x75, 0x6c,
	0x6c, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69,
	0x74, 0x68, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x66, 0x75,
	0x6c, 0x6c, 0x4d, 0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74,
	0x68, 0x6d, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x1c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x19, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x61, 0x6c, 0x4d, 0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69,
	0x74, 0x68, 0x6d, 0x49, 0x64, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x14, 0x5a, 0x12, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f,
	0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_branch_proto_rawDescOnce sync.Once
	file_store_branch_proto_rawDescData = file_store_branch_proto_rawDesc
)

func file_store_branch_proto_rawDescGZIP() []byte {
	file_store_branch_proto_rawDescOnce.Do(func() {
		file_store_branch_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_branch_proto_rawDescData)
	})
	return file_store_branch_proto_rawDescData
}

var file_store_branch_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_store_branch_proto_goTypes = []any{
	(*BranchSnapshot)(nil),         // 0: bytebase.store.BranchSnapshot
	(*BranchConfig)(nil),           // 1: bytebase.store.BranchConfig
	(*BranchDatabaseConfig)(nil),   // 2: bytebase.store.BranchDatabaseConfig
	(*BranchSchemaConfig)(nil),     // 3: bytebase.store.BranchSchemaConfig
	(*BranchTableConfig)(nil),      // 4: bytebase.store.BranchTableConfig
	(*BranchFunctionConfig)(nil),   // 5: bytebase.store.BranchFunctionConfig
	(*BranchProcedureConfig)(nil),  // 6: bytebase.store.BranchProcedureConfig
	(*BranchViewConfig)(nil),       // 7: bytebase.store.BranchViewConfig
	(*BranchColumnConfig)(nil),     // 8: bytebase.store.BranchColumnConfig
	nil,                            // 9: bytebase.store.BranchColumnConfig.LabelsEntry
	(*DatabaseSchemaMetadata)(nil), // 10: bytebase.store.DatabaseSchemaMetadata
	(*timestamppb.Timestamp)(nil),  // 11: google.protobuf.Timestamp
	(MaskingLevel)(0),              // 12: bytebase.store.MaskingLevel
}
var file_store_branch_proto_depIdxs = []int32{
	10, // 0: bytebase.store.BranchSnapshot.metadata:type_name -> bytebase.store.DatabaseSchemaMetadata
	2,  // 1: bytebase.store.BranchSnapshot.database_config:type_name -> bytebase.store.BranchDatabaseConfig
	3,  // 2: bytebase.store.BranchDatabaseConfig.schema_configs:type_name -> bytebase.store.BranchSchemaConfig
	4,  // 3: bytebase.store.BranchSchemaConfig.table_configs:type_name -> bytebase.store.BranchTableConfig
	5,  // 4: bytebase.store.BranchSchemaConfig.function_configs:type_name -> bytebase.store.BranchFunctionConfig
	6,  // 5: bytebase.store.BranchSchemaConfig.procedure_configs:type_name -> bytebase.store.BranchProcedureConfig
	7,  // 6: bytebase.store.BranchSchemaConfig.view_configs:type_name -> bytebase.store.BranchViewConfig
	8,  // 7: bytebase.store.BranchTableConfig.column_configs:type_name -> bytebase.store.BranchColumnConfig
	11, // 8: bytebase.store.BranchTableConfig.update_time:type_name -> google.protobuf.Timestamp
	11, // 9: bytebase.store.BranchFunctionConfig.update_time:type_name -> google.protobuf.Timestamp
	11, // 10: bytebase.store.BranchProcedureConfig.update_time:type_name -> google.protobuf.Timestamp
	11, // 11: bytebase.store.BranchViewConfig.update_time:type_name -> google.protobuf.Timestamp
	9,  // 12: bytebase.store.BranchColumnConfig.labels:type_name -> bytebase.store.BranchColumnConfig.LabelsEntry
	12, // 13: bytebase.store.BranchColumnConfig.masking_level:type_name -> bytebase.store.MaskingLevel
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_store_branch_proto_init() }
func file_store_branch_proto_init() {
	if File_store_branch_proto != nil {
		return
	}
	file_store_common_proto_init()
	file_store_database_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_store_branch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_branch_proto_goTypes,
		DependencyIndexes: file_store_branch_proto_depIdxs,
		MessageInfos:      file_store_branch_proto_msgTypes,
	}.Build()
	File_store_branch_proto = out.File
	file_store_branch_proto_rawDesc = nil
	file_store_branch_proto_goTypes = nil
	file_store_branch_proto_depIdxs = nil
}
