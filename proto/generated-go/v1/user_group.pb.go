// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: v1/user_group.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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

type UserGroupMember_Role int32

const (
	UserGroupMember_ROLE_UNSPECIFIED UserGroupMember_Role = 0
	UserGroupMember_OWNER            UserGroupMember_Role = 1
	UserGroupMember_MEMBER           UserGroupMember_Role = 2
)

// Enum value maps for UserGroupMember_Role.
var (
	UserGroupMember_Role_name = map[int32]string{
		0: "ROLE_UNSPECIFIED",
		1: "OWNER",
		2: "MEMBER",
	}
	UserGroupMember_Role_value = map[string]int32{
		"ROLE_UNSPECIFIED": 0,
		"OWNER":            1,
		"MEMBER":           2,
	}
)

func (x UserGroupMember_Role) Enum() *UserGroupMember_Role {
	p := new(UserGroupMember_Role)
	*p = x
	return p
}

func (x UserGroupMember_Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserGroupMember_Role) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_user_group_proto_enumTypes[0].Descriptor()
}

func (UserGroupMember_Role) Type() protoreflect.EnumType {
	return &file_v1_user_group_proto_enumTypes[0]
}

func (x UserGroupMember_Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserGroupMember_Role.Descriptor instead.
func (UserGroupMember_Role) EnumDescriptor() ([]byte, []int) {
	return file_v1_user_group_proto_rawDescGZIP(), []int{6, 0}
}

type GetUserGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the group to retrieve.
	// Format: groups/{email}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetUserGroupRequest) Reset() {
	*x = GetUserGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_user_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserGroupRequest) ProtoMessage() {}

func (x *GetUserGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserGroupRequest.ProtoReflect.Descriptor instead.
func (*GetUserGroupRequest) Descriptor() ([]byte, []int) {
	return file_v1_user_group_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListUserGroupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum number of groups to return. The service may return fewer than
	// this value.
	// If unspecified, at most 50 groups will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A page token, received from a previous `ListUsers` call.
	// Provide this to retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to `ListUsers` must match
	// the call that provided the page token.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListUserGroupsRequest) Reset() {
	*x = ListUserGroupsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_user_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserGroupsRequest) ProtoMessage() {}

func (x *ListUserGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserGroupsRequest.ProtoReflect.Descriptor instead.
func (*ListUserGroupsRequest) Descriptor() ([]byte, []int) {
	return file_v1_user_group_proto_rawDescGZIP(), []int{1}
}

func (x *ListUserGroupsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListUserGroupsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListUserGroupsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The groups from the specified request.
	Groups []*UserGroup `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
	// A token, which can be sent as `page_token` to retrieve the next page.
	// If this field is omitted, there are no subsequent pages.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListUserGroupsResponse) Reset() {
	*x = ListUserGroupsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_user_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserGroupsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserGroupsResponse) ProtoMessage() {}

func (x *ListUserGroupsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserGroupsResponse.ProtoReflect.Descriptor instead.
func (*ListUserGroupsResponse) Descriptor() ([]byte, []int) {
	return file_v1_user_group_proto_rawDescGZIP(), []int{2}
}

func (x *ListUserGroupsResponse) GetGroups() []*UserGroup {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *ListUserGroupsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type CreateUserGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The group to create.
	Group *UserGroup `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *CreateUserGroupRequest) Reset() {
	*x = CreateUserGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_user_group_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserGroupRequest) ProtoMessage() {}

func (x *CreateUserGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_group_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateUserGroupRequest) Descriptor() ([]byte, []int) {
	return file_v1_user_group_proto_rawDescGZIP(), []int{3}
}

func (x *CreateUserGroupRequest) GetGroup() *UserGroup {
	if x != nil {
		return x.Group
	}
	return nil
}

type UpdateUserGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The group to update.
	//
	// The group's `name` field is used to identify the group to update.
	// Format: groups/{email}
	Group *UserGroup `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	// The list of fields to update.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateUserGroupRequest) Reset() {
	*x = UpdateUserGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_user_group_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserGroupRequest) ProtoMessage() {}

func (x *UpdateUserGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_group_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserGroupRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserGroupRequest) Descriptor() ([]byte, []int) {
	return file_v1_user_group_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateUserGroupRequest) GetGroup() *UserGroup {
	if x != nil {
		return x.Group
	}
	return nil
}

func (x *UpdateUserGroupRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type DeleteUserGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the group to delete.
	// Format: groups/{email}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteUserGroupRequest) Reset() {
	*x = DeleteUserGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_user_group_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserGroupRequest) ProtoMessage() {}

func (x *DeleteUserGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_group_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserGroupRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserGroupRequest) Descriptor() ([]byte, []int) {
	return file_v1_user_group_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteUserGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UserGroupMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Member is the principal who belong to this user group.
	//
	// Format: users/hello@world.com
	Member string               `protobuf:"bytes,1,opt,name=member,proto3" json:"member,omitempty"`
	Role   UserGroupMember_Role `protobuf:"varint,2,opt,name=role,proto3,enum=bytebase.v1.UserGroupMember_Role" json:"role,omitempty"`
}

func (x *UserGroupMember) Reset() {
	*x = UserGroupMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_user_group_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserGroupMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGroupMember) ProtoMessage() {}

func (x *UserGroupMember) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_group_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGroupMember.ProtoReflect.Descriptor instead.
func (*UserGroupMember) Descriptor() ([]byte, []int) {
	return file_v1_user_group_proto_rawDescGZIP(), []int{6}
}

func (x *UserGroupMember) GetMember() string {
	if x != nil {
		return x.Member
	}
	return ""
}

func (x *UserGroupMember) GetRole() UserGroupMember_Role {
	if x != nil {
		return x.Role
	}
	return UserGroupMember_ROLE_UNSPECIFIED
}

type UserGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the group to retrieve.
	// Format: groups/{email}
	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// The name for the creator.
	// Format: users/hello@world.com
	Creator string             `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
	Members []*UserGroupMember `protobuf:"bytes,5,rep,name=members,proto3" json:"members,omitempty"`
	// The timestamp when the group was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

func (x *UserGroup) Reset() {
	*x = UserGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_user_group_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGroup) ProtoMessage() {}

func (x *UserGroup) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_group_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGroup.ProtoReflect.Descriptor instead.
func (*UserGroup) Descriptor() ([]byte, []int) {
	return file_v1_user_group_proto_rawDescGZIP(), []int{7}
}

func (x *UserGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserGroup) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UserGroup) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UserGroup) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *UserGroup) GetMembers() []*UserGroupMember {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *UserGroup) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

var File_v1_user_group_proto protoreflect.FileDescriptor

var file_v1_user_group_proto_rawDesc = []byte{
	0x0a, 0x13, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d,
	0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x53, 0x0a, 0x15, 0x4c,
	0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x70, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x79, 0x74,
	0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65,
	0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x4c, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x05,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x79,
	0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x22, 0x89, 0x01, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x05, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x79, 0x74,
	0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x32, 0x0a, 0x16,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x95, 0x01, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x62, 0x79, 0x74,
	0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x22, 0x33, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x52,
	0x4f, 0x4c, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x02, 0x22, 0xf8, 0x01, 0x0a, 0x09, 0x55, 0x73, 0x65,
	0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x18, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x36, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x79, 0x74, 0x65,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x12, 0x41, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x32, 0xe8, 0x04, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6c, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x20, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x62, 0x79, 0x74,
	0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x22, 0x22, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x15, 0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x70, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x22, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x62,
	0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x15, 0xda, 0x41, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x76,
	0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x71, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x23, 0x2e, 0x62, 0x79,
	0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x21, 0xda, 0x41, 0x05, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22,
	0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x8c, 0x01, 0x0a, 0x0f,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x23, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x3c, 0xda, 0x41,
	0x11, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61,
	0x73, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x32,
	0x19, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x6e, 0x61, 0x6d, 0x65,
	0x3d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x72, 0x0a, 0x0f, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x23, 0x2e,
	0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x22, 0xda, 0x41, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x2a, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x7b,
	0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x2a, 0x7d, 0x42, 0x11,
	0x5a, 0x0f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_user_group_proto_rawDescOnce sync.Once
	file_v1_user_group_proto_rawDescData = file_v1_user_group_proto_rawDesc
)

func file_v1_user_group_proto_rawDescGZIP() []byte {
	file_v1_user_group_proto_rawDescOnce.Do(func() {
		file_v1_user_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_user_group_proto_rawDescData)
	})
	return file_v1_user_group_proto_rawDescData
}

var file_v1_user_group_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_user_group_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_v1_user_group_proto_goTypes = []any{
	(UserGroupMember_Role)(0),      // 0: bytebase.v1.UserGroupMember.Role
	(*GetUserGroupRequest)(nil),    // 1: bytebase.v1.GetUserGroupRequest
	(*ListUserGroupsRequest)(nil),  // 2: bytebase.v1.ListUserGroupsRequest
	(*ListUserGroupsResponse)(nil), // 3: bytebase.v1.ListUserGroupsResponse
	(*CreateUserGroupRequest)(nil), // 4: bytebase.v1.CreateUserGroupRequest
	(*UpdateUserGroupRequest)(nil), // 5: bytebase.v1.UpdateUserGroupRequest
	(*DeleteUserGroupRequest)(nil), // 6: bytebase.v1.DeleteUserGroupRequest
	(*UserGroupMember)(nil),        // 7: bytebase.v1.UserGroupMember
	(*UserGroup)(nil),              // 8: bytebase.v1.UserGroup
	(*fieldmaskpb.FieldMask)(nil),  // 9: google.protobuf.FieldMask
	(*timestamppb.Timestamp)(nil),  // 10: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),          // 11: google.protobuf.Empty
}
var file_v1_user_group_proto_depIdxs = []int32{
	8,  // 0: bytebase.v1.ListUserGroupsResponse.groups:type_name -> bytebase.v1.UserGroup
	8,  // 1: bytebase.v1.CreateUserGroupRequest.group:type_name -> bytebase.v1.UserGroup
	8,  // 2: bytebase.v1.UpdateUserGroupRequest.group:type_name -> bytebase.v1.UserGroup
	9,  // 3: bytebase.v1.UpdateUserGroupRequest.update_mask:type_name -> google.protobuf.FieldMask
	0,  // 4: bytebase.v1.UserGroupMember.role:type_name -> bytebase.v1.UserGroupMember.Role
	7,  // 5: bytebase.v1.UserGroup.members:type_name -> bytebase.v1.UserGroupMember
	10, // 6: bytebase.v1.UserGroup.create_time:type_name -> google.protobuf.Timestamp
	1,  // 7: bytebase.v1.UserGroupService.GetUserGroup:input_type -> bytebase.v1.GetUserGroupRequest
	2,  // 8: bytebase.v1.UserGroupService.ListUserGroups:input_type -> bytebase.v1.ListUserGroupsRequest
	4,  // 9: bytebase.v1.UserGroupService.CreateUserGroup:input_type -> bytebase.v1.CreateUserGroupRequest
	5,  // 10: bytebase.v1.UserGroupService.UpdateUserGroup:input_type -> bytebase.v1.UpdateUserGroupRequest
	6,  // 11: bytebase.v1.UserGroupService.DeleteUserGroup:input_type -> bytebase.v1.DeleteUserGroupRequest
	8,  // 12: bytebase.v1.UserGroupService.GetUserGroup:output_type -> bytebase.v1.UserGroup
	3,  // 13: bytebase.v1.UserGroupService.ListUserGroups:output_type -> bytebase.v1.ListUserGroupsResponse
	8,  // 14: bytebase.v1.UserGroupService.CreateUserGroup:output_type -> bytebase.v1.UserGroup
	8,  // 15: bytebase.v1.UserGroupService.UpdateUserGroup:output_type -> bytebase.v1.UserGroup
	11, // 16: bytebase.v1.UserGroupService.DeleteUserGroup:output_type -> google.protobuf.Empty
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_v1_user_group_proto_init() }
func file_v1_user_group_proto_init() {
	if File_v1_user_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_user_group_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserGroupRequest); i {
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
		file_v1_user_group_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListUserGroupsRequest); i {
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
		file_v1_user_group_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ListUserGroupsResponse); i {
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
		file_v1_user_group_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CreateUserGroupRequest); i {
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
		file_v1_user_group_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateUserGroupRequest); i {
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
		file_v1_user_group_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteUserGroupRequest); i {
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
		file_v1_user_group_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UserGroupMember); i {
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
		file_v1_user_group_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*UserGroup); i {
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
			RawDescriptor: file_v1_user_group_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_user_group_proto_goTypes,
		DependencyIndexes: file_v1_user_group_proto_depIdxs,
		EnumInfos:         file_v1_user_group_proto_enumTypes,
		MessageInfos:      file_v1_user_group_proto_msgTypes,
	}.Build()
	File_v1_user_group_proto = out.File
	file_v1_user_group_proto_rawDesc = nil
	file_v1_user_group_proto_goTypes = nil
	file_v1_user_group_proto_depIdxs = nil
}
