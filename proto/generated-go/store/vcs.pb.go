// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: store/vcs.proto

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

type VCSConnector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The title or display name of the VCS connector.
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// Full path from the corresponding VCS provider.
	// For GitLab, this is the project full path. e.g. group1/project-1
	FullPath string `protobuf:"bytes,2,opt,name=full_path,json=fullPath,proto3" json:"full_path,omitempty"`
	// Web url from the corresponding VCS provider.
	// For GitLab, this is the project web url. e.g. https://gitlab.example.com/group1/project-1
	WebUrl string `protobuf:"bytes,3,opt,name=web_url,json=webUrl,proto3" json:"web_url,omitempty"`
	// Branch to listen to.
	Branch string `protobuf:"bytes,4,opt,name=branch,proto3" json:"branch,omitempty"`
	// Base working directory we are interested.
	BaseDirectory string `protobuf:"bytes,5,opt,name=base_directory,json=baseDirectory,proto3" json:"base_directory,omitempty"`
	// Repository id from the corresponding VCS provider.
	// For GitLab, this is the project id. e.g. 123
	ExternalId string `protobuf:"bytes,6,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	// Push webhook id from the corresponding VCS provider.
	// For GitLab, this is the project webhook id. e.g. 123
	ExternalWebhookId string `protobuf:"bytes,7,opt,name=external_webhook_id,json=externalWebhookId,proto3" json:"external_webhook_id,omitempty"`
	// For GitLab, webhook request contains this in the 'X-Gitlab-Token" header and we compare it with the one stored in db to validate it sends to the expected endpoint.
	WebhookSecretToken string `protobuf:"bytes,8,opt,name=webhook_secret_token,json=webhookSecretToken,proto3" json:"webhook_secret_token,omitempty"`
	// Apply changes to the database group. Optional, if not set, will apply changes to all databases in the project.
	// Format: projects/{project}/databaseGroups/{databaseGroup}
	DatabaseGroup string `protobuf:"bytes,9,opt,name=database_group,json=databaseGroup,proto3" json:"database_group,omitempty"`
}

func (x *VCSConnector) Reset() {
	*x = VCSConnector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_vcs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VCSConnector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VCSConnector) ProtoMessage() {}

func (x *VCSConnector) ProtoReflect() protoreflect.Message {
	mi := &file_store_vcs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VCSConnector.ProtoReflect.Descriptor instead.
func (*VCSConnector) Descriptor() ([]byte, []int) {
	return file_store_vcs_proto_rawDescGZIP(), []int{0}
}

func (x *VCSConnector) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *VCSConnector) GetFullPath() string {
	if x != nil {
		return x.FullPath
	}
	return ""
}

func (x *VCSConnector) GetWebUrl() string {
	if x != nil {
		return x.WebUrl
	}
	return ""
}

func (x *VCSConnector) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *VCSConnector) GetBaseDirectory() string {
	if x != nil {
		return x.BaseDirectory
	}
	return ""
}

func (x *VCSConnector) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *VCSConnector) GetExternalWebhookId() string {
	if x != nil {
		return x.ExternalWebhookId
	}
	return ""
}

func (x *VCSConnector) GetWebhookSecretToken() string {
	if x != nil {
		return x.WebhookSecretToken
	}
	return ""
}

func (x *VCSConnector) GetDatabaseGroup() string {
	if x != nil {
		return x.DatabaseGroup
	}
	return ""
}

var File_store_vcs_proto protoreflect.FileDescriptor

var file_store_vcs_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x22, 0xc3, 0x02, 0x0a, 0x0c, 0x56, 0x43, 0x53, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c,
	0x6c, 0x50, 0x61, 0x74, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x65, 0x62, 0x55, 0x72, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x62, 0x61, 0x73, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1f, 0x0a,
	0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x2e,
	0x0a, 0x13, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x77, 0x65, 0x62, 0x68, 0x6f,
	0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x30,
	0x0a, 0x14, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x77, 0x65,
	0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x25, 0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x14, 0x5a, 0x12, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_vcs_proto_rawDescOnce sync.Once
	file_store_vcs_proto_rawDescData = file_store_vcs_proto_rawDesc
)

func file_store_vcs_proto_rawDescGZIP() []byte {
	file_store_vcs_proto_rawDescOnce.Do(func() {
		file_store_vcs_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_vcs_proto_rawDescData)
	})
	return file_store_vcs_proto_rawDescData
}

var file_store_vcs_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_store_vcs_proto_goTypes = []any{
	(*VCSConnector)(nil), // 0: bytebase.store.VCSConnector
}
var file_store_vcs_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_store_vcs_proto_init() }
func file_store_vcs_proto_init() {
	if File_store_vcs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_store_vcs_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*VCSConnector); i {
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
			RawDescriptor: file_store_vcs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_vcs_proto_goTypes,
		DependencyIndexes: file_store_vcs_proto_depIdxs,
		MessageInfos:      file_store_vcs_proto_msgTypes,
	}.Build()
	File_store_vcs_proto = out.File
	file_store_vcs_proto_rawDesc = nil
	file_store_vcs_proto_goTypes = nil
	file_store_vcs_proto_depIdxs = nil
}
