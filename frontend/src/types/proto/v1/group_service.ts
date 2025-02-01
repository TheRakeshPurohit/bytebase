// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.3.0
//   protoc               unknown
// source: v1/group_service.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import Long from "long";
import { Empty } from "../google/protobuf/empty";
import { FieldMask } from "../google/protobuf/field_mask";
import { Timestamp } from "../google/protobuf/timestamp";

export const protobufPackage = "bytebase.v1";

export interface GetGroupRequest {
  /**
   * The name of the group to retrieve.
   * Format: groups/{email}
   */
  name: string;
}

export interface ListGroupsRequest {
  /**
   * Not used.
   * The maximum number of groups to return. The service may return fewer than
   * this value.
   * If unspecified, at most 10 groups will be returned.
   * The maximum value is 1000; values above 1000 will be coerced to 1000.
   */
  pageSize: number;
  /**
   * Not used.
   * A page token, received from a previous `ListGroups` call.
   * Provide this to retrieve the subsequent page.
   *
   * When paginating, all other parameters provided to `ListGroups` must match
   * the call that provided the page token.
   */
  pageToken: string;
}

export interface ListGroupsResponse {
  /** The groups from the specified request. */
  groups: Group[];
  /**
   * A token, which can be sent as `page_token` to retrieve the next page.
   * If this field is omitted, there are no subsequent pages.
   */
  nextPageToken: string;
}

export interface CreateGroupRequest {
  /** The group to create. */
  group:
    | Group
    | undefined;
  /**
   * The email to use for the group, which will become the final component
   * of the group's resource name.
   */
  groupEmail: string;
}

export interface UpdateGroupRequest {
  /**
   * The group to update.
   *
   * The group's `name` field is used to identify the group to update.
   * Format: groups/{email}
   */
  group:
    | Group
    | undefined;
  /** The list of fields to update. */
  updateMask:
    | string[]
    | undefined;
  /**
   * If set to true, and the group is not found, a new group will be created.
   * In this situation, `update_mask` is ignored.
   */
  allowMissing: boolean;
}

export interface DeleteGroupRequest {
  /**
   * The name of the group to delete.
   * Format: groups/{email}
   */
  name: string;
}

export interface GroupMember {
  /**
   * Member is the principal who belong to this group.
   *
   * Format: users/hello@world.com
   */
  member: string;
  role: GroupMember_Role;
}

export enum GroupMember_Role {
  ROLE_UNSPECIFIED = "ROLE_UNSPECIFIED",
  OWNER = "OWNER",
  MEMBER = "MEMBER",
  UNRECOGNIZED = "UNRECOGNIZED",
}

export function groupMember_RoleFromJSON(object: any): GroupMember_Role {
  switch (object) {
    case 0:
    case "ROLE_UNSPECIFIED":
      return GroupMember_Role.ROLE_UNSPECIFIED;
    case 1:
    case "OWNER":
      return GroupMember_Role.OWNER;
    case 2:
    case "MEMBER":
      return GroupMember_Role.MEMBER;
    case -1:
    case "UNRECOGNIZED":
    default:
      return GroupMember_Role.UNRECOGNIZED;
  }
}

export function groupMember_RoleToJSON(object: GroupMember_Role): string {
  switch (object) {
    case GroupMember_Role.ROLE_UNSPECIFIED:
      return "ROLE_UNSPECIFIED";
    case GroupMember_Role.OWNER:
      return "OWNER";
    case GroupMember_Role.MEMBER:
      return "MEMBER";
    case GroupMember_Role.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export function groupMember_RoleToNumber(object: GroupMember_Role): number {
  switch (object) {
    case GroupMember_Role.ROLE_UNSPECIFIED:
      return 0;
    case GroupMember_Role.OWNER:
      return 1;
    case GroupMember_Role.MEMBER:
      return 2;
    case GroupMember_Role.UNRECOGNIZED:
    default:
      return -1;
  }
}

export interface Group {
  /**
   * The name of the group to retrieve.
   * Format: groups/{group}, group is an email.
   */
  name: string;
  title: string;
  description: string;
  /**
   * The name for the creator.
   * Format: users/hello@world.com
   */
  creator: string;
  members: GroupMember[];
  /** The timestamp when the group was created. */
  createTime:
    | Timestamp
    | undefined;
  /** source means where the group comes from. For now we support Entra ID SCIM sync, so the source could be Entra ID. */
  source: string;
}

function createBaseGetGroupRequest(): GetGroupRequest {
  return { name: "" };
}

export const GetGroupRequest: MessageFns<GetGroupRequest> = {
  encode(message: GetGroupRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): GetGroupRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetGroupRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetGroupRequest {
    return { name: isSet(object.name) ? globalThis.String(object.name) : "" };
  },

  toJSON(message: GetGroupRequest): unknown {
    const obj: any = {};
    if (message.name !== "") {
      obj.name = message.name;
    }
    return obj;
  },

  create(base?: DeepPartial<GetGroupRequest>): GetGroupRequest {
    return GetGroupRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<GetGroupRequest>): GetGroupRequest {
    const message = createBaseGetGroupRequest();
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseListGroupsRequest(): ListGroupsRequest {
  return { pageSize: 0, pageToken: "" };
}

export const ListGroupsRequest: MessageFns<ListGroupsRequest> = {
  encode(message: ListGroupsRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.pageSize !== 0) {
      writer.uint32(8).int32(message.pageSize);
    }
    if (message.pageToken !== "") {
      writer.uint32(18).string(message.pageToken);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): ListGroupsRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListGroupsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 8) {
            break;
          }

          message.pageSize = reader.int32();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.pageToken = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListGroupsRequest {
    return {
      pageSize: isSet(object.pageSize) ? globalThis.Number(object.pageSize) : 0,
      pageToken: isSet(object.pageToken) ? globalThis.String(object.pageToken) : "",
    };
  },

  toJSON(message: ListGroupsRequest): unknown {
    const obj: any = {};
    if (message.pageSize !== 0) {
      obj.pageSize = Math.round(message.pageSize);
    }
    if (message.pageToken !== "") {
      obj.pageToken = message.pageToken;
    }
    return obj;
  },

  create(base?: DeepPartial<ListGroupsRequest>): ListGroupsRequest {
    return ListGroupsRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<ListGroupsRequest>): ListGroupsRequest {
    const message = createBaseListGroupsRequest();
    message.pageSize = object.pageSize ?? 0;
    message.pageToken = object.pageToken ?? "";
    return message;
  },
};

function createBaseListGroupsResponse(): ListGroupsResponse {
  return { groups: [], nextPageToken: "" };
}

export const ListGroupsResponse: MessageFns<ListGroupsResponse> = {
  encode(message: ListGroupsResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    for (const v of message.groups) {
      Group.encode(v!, writer.uint32(10).fork()).join();
    }
    if (message.nextPageToken !== "") {
      writer.uint32(18).string(message.nextPageToken);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): ListGroupsResponse {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListGroupsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.groups.push(Group.decode(reader, reader.uint32()));
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.nextPageToken = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListGroupsResponse {
    return {
      groups: globalThis.Array.isArray(object?.groups) ? object.groups.map((e: any) => Group.fromJSON(e)) : [],
      nextPageToken: isSet(object.nextPageToken) ? globalThis.String(object.nextPageToken) : "",
    };
  },

  toJSON(message: ListGroupsResponse): unknown {
    const obj: any = {};
    if (message.groups?.length) {
      obj.groups = message.groups.map((e) => Group.toJSON(e));
    }
    if (message.nextPageToken !== "") {
      obj.nextPageToken = message.nextPageToken;
    }
    return obj;
  },

  create(base?: DeepPartial<ListGroupsResponse>): ListGroupsResponse {
    return ListGroupsResponse.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<ListGroupsResponse>): ListGroupsResponse {
    const message = createBaseListGroupsResponse();
    message.groups = object.groups?.map((e) => Group.fromPartial(e)) || [];
    message.nextPageToken = object.nextPageToken ?? "";
    return message;
  },
};

function createBaseCreateGroupRequest(): CreateGroupRequest {
  return { group: undefined, groupEmail: "" };
}

export const CreateGroupRequest: MessageFns<CreateGroupRequest> = {
  encode(message: CreateGroupRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.group !== undefined) {
      Group.encode(message.group, writer.uint32(10).fork()).join();
    }
    if (message.groupEmail !== "") {
      writer.uint32(18).string(message.groupEmail);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): CreateGroupRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateGroupRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.group = Group.decode(reader, reader.uint32());
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.groupEmail = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateGroupRequest {
    return {
      group: isSet(object.group) ? Group.fromJSON(object.group) : undefined,
      groupEmail: isSet(object.groupEmail) ? globalThis.String(object.groupEmail) : "",
    };
  },

  toJSON(message: CreateGroupRequest): unknown {
    const obj: any = {};
    if (message.group !== undefined) {
      obj.group = Group.toJSON(message.group);
    }
    if (message.groupEmail !== "") {
      obj.groupEmail = message.groupEmail;
    }
    return obj;
  },

  create(base?: DeepPartial<CreateGroupRequest>): CreateGroupRequest {
    return CreateGroupRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<CreateGroupRequest>): CreateGroupRequest {
    const message = createBaseCreateGroupRequest();
    message.group = (object.group !== undefined && object.group !== null) ? Group.fromPartial(object.group) : undefined;
    message.groupEmail = object.groupEmail ?? "";
    return message;
  },
};

function createBaseUpdateGroupRequest(): UpdateGroupRequest {
  return { group: undefined, updateMask: undefined, allowMissing: false };
}

export const UpdateGroupRequest: MessageFns<UpdateGroupRequest> = {
  encode(message: UpdateGroupRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.group !== undefined) {
      Group.encode(message.group, writer.uint32(10).fork()).join();
    }
    if (message.updateMask !== undefined) {
      FieldMask.encode(FieldMask.wrap(message.updateMask), writer.uint32(18).fork()).join();
    }
    if (message.allowMissing !== false) {
      writer.uint32(24).bool(message.allowMissing);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): UpdateGroupRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateGroupRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.group = Group.decode(reader, reader.uint32());
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.updateMask = FieldMask.unwrap(FieldMask.decode(reader, reader.uint32()));
          continue;
        }
        case 3: {
          if (tag !== 24) {
            break;
          }

          message.allowMissing = reader.bool();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateGroupRequest {
    return {
      group: isSet(object.group) ? Group.fromJSON(object.group) : undefined,
      updateMask: isSet(object.updateMask) ? FieldMask.unwrap(FieldMask.fromJSON(object.updateMask)) : undefined,
      allowMissing: isSet(object.allowMissing) ? globalThis.Boolean(object.allowMissing) : false,
    };
  },

  toJSON(message: UpdateGroupRequest): unknown {
    const obj: any = {};
    if (message.group !== undefined) {
      obj.group = Group.toJSON(message.group);
    }
    if (message.updateMask !== undefined) {
      obj.updateMask = FieldMask.toJSON(FieldMask.wrap(message.updateMask));
    }
    if (message.allowMissing !== false) {
      obj.allowMissing = message.allowMissing;
    }
    return obj;
  },

  create(base?: DeepPartial<UpdateGroupRequest>): UpdateGroupRequest {
    return UpdateGroupRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<UpdateGroupRequest>): UpdateGroupRequest {
    const message = createBaseUpdateGroupRequest();
    message.group = (object.group !== undefined && object.group !== null) ? Group.fromPartial(object.group) : undefined;
    message.updateMask = object.updateMask ?? undefined;
    message.allowMissing = object.allowMissing ?? false;
    return message;
  },
};

function createBaseDeleteGroupRequest(): DeleteGroupRequest {
  return { name: "" };
}

export const DeleteGroupRequest: MessageFns<DeleteGroupRequest> = {
  encode(message: DeleteGroupRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): DeleteGroupRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteGroupRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteGroupRequest {
    return { name: isSet(object.name) ? globalThis.String(object.name) : "" };
  },

  toJSON(message: DeleteGroupRequest): unknown {
    const obj: any = {};
    if (message.name !== "") {
      obj.name = message.name;
    }
    return obj;
  },

  create(base?: DeepPartial<DeleteGroupRequest>): DeleteGroupRequest {
    return DeleteGroupRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<DeleteGroupRequest>): DeleteGroupRequest {
    const message = createBaseDeleteGroupRequest();
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseGroupMember(): GroupMember {
  return { member: "", role: GroupMember_Role.ROLE_UNSPECIFIED };
}

export const GroupMember: MessageFns<GroupMember> = {
  encode(message: GroupMember, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.member !== "") {
      writer.uint32(10).string(message.member);
    }
    if (message.role !== GroupMember_Role.ROLE_UNSPECIFIED) {
      writer.uint32(16).int32(groupMember_RoleToNumber(message.role));
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): GroupMember {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGroupMember();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.member = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 16) {
            break;
          }

          message.role = groupMember_RoleFromJSON(reader.int32());
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GroupMember {
    return {
      member: isSet(object.member) ? globalThis.String(object.member) : "",
      role: isSet(object.role) ? groupMember_RoleFromJSON(object.role) : GroupMember_Role.ROLE_UNSPECIFIED,
    };
  },

  toJSON(message: GroupMember): unknown {
    const obj: any = {};
    if (message.member !== "") {
      obj.member = message.member;
    }
    if (message.role !== GroupMember_Role.ROLE_UNSPECIFIED) {
      obj.role = groupMember_RoleToJSON(message.role);
    }
    return obj;
  },

  create(base?: DeepPartial<GroupMember>): GroupMember {
    return GroupMember.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<GroupMember>): GroupMember {
    const message = createBaseGroupMember();
    message.member = object.member ?? "";
    message.role = object.role ?? GroupMember_Role.ROLE_UNSPECIFIED;
    return message;
  },
};

function createBaseGroup(): Group {
  return { name: "", title: "", description: "", creator: "", members: [], createTime: undefined, source: "" };
}

export const Group: MessageFns<Group> = {
  encode(message: Group, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.title !== "") {
      writer.uint32(18).string(message.title);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    if (message.creator !== "") {
      writer.uint32(34).string(message.creator);
    }
    for (const v of message.members) {
      GroupMember.encode(v!, writer.uint32(42).fork()).join();
    }
    if (message.createTime !== undefined) {
      Timestamp.encode(message.createTime, writer.uint32(50).fork()).join();
    }
    if (message.source !== "") {
      writer.uint32(58).string(message.source);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): Group {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGroup();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.title = reader.string();
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.description = reader.string();
          continue;
        }
        case 4: {
          if (tag !== 34) {
            break;
          }

          message.creator = reader.string();
          continue;
        }
        case 5: {
          if (tag !== 42) {
            break;
          }

          message.members.push(GroupMember.decode(reader, reader.uint32()));
          continue;
        }
        case 6: {
          if (tag !== 50) {
            break;
          }

          message.createTime = Timestamp.decode(reader, reader.uint32());
          continue;
        }
        case 7: {
          if (tag !== 58) {
            break;
          }

          message.source = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Group {
    return {
      name: isSet(object.name) ? globalThis.String(object.name) : "",
      title: isSet(object.title) ? globalThis.String(object.title) : "",
      description: isSet(object.description) ? globalThis.String(object.description) : "",
      creator: isSet(object.creator) ? globalThis.String(object.creator) : "",
      members: globalThis.Array.isArray(object?.members) ? object.members.map((e: any) => GroupMember.fromJSON(e)) : [],
      createTime: isSet(object.createTime) ? fromJsonTimestamp(object.createTime) : undefined,
      source: isSet(object.source) ? globalThis.String(object.source) : "",
    };
  },

  toJSON(message: Group): unknown {
    const obj: any = {};
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.title !== "") {
      obj.title = message.title;
    }
    if (message.description !== "") {
      obj.description = message.description;
    }
    if (message.creator !== "") {
      obj.creator = message.creator;
    }
    if (message.members?.length) {
      obj.members = message.members.map((e) => GroupMember.toJSON(e));
    }
    if (message.createTime !== undefined) {
      obj.createTime = fromTimestamp(message.createTime).toISOString();
    }
    if (message.source !== "") {
      obj.source = message.source;
    }
    return obj;
  },

  create(base?: DeepPartial<Group>): Group {
    return Group.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<Group>): Group {
    const message = createBaseGroup();
    message.name = object.name ?? "";
    message.title = object.title ?? "";
    message.description = object.description ?? "";
    message.creator = object.creator ?? "";
    message.members = object.members?.map((e) => GroupMember.fromPartial(e)) || [];
    message.createTime = (object.createTime !== undefined && object.createTime !== null)
      ? Timestamp.fromPartial(object.createTime)
      : undefined;
    message.source = object.source ?? "";
    return message;
  },
};

export type GroupServiceDefinition = typeof GroupServiceDefinition;
export const GroupServiceDefinition = {
  name: "GroupService",
  fullName: "bytebase.v1.GroupService",
  methods: {
    getGroup: {
      name: "GetGroup",
      requestType: GetGroupRequest,
      requestStream: false,
      responseType: Group,
      responseStream: false,
      options: {
        _unknownFields: {
          8410: [new Uint8Array([4, 110, 97, 109, 101])],
          800010: [new Uint8Array([13, 98, 98, 46, 103, 114, 111, 117, 112, 115, 46, 103, 101, 116])],
          800016: [new Uint8Array([1])],
          578365826: [
            new Uint8Array([
              21,
              18,
              19,
              47,
              118,
              49,
              47,
              123,
              110,
              97,
              109,
              101,
              61,
              103,
              114,
              111,
              117,
              112,
              115,
              47,
              42,
              125,
            ]),
          ],
        },
      },
    },
    listGroups: {
      name: "ListGroups",
      requestType: ListGroupsRequest,
      requestStream: false,
      responseType: ListGroupsResponse,
      responseStream: false,
      options: {
        _unknownFields: {
          8410: [new Uint8Array([0])],
          800010: [new Uint8Array([14, 98, 98, 46, 103, 114, 111, 117, 112, 115, 46, 108, 105, 115, 116])],
          800016: [new Uint8Array([1])],
          578365826: [new Uint8Array([12, 18, 10, 47, 118, 49, 47, 103, 114, 111, 117, 112, 115])],
        },
      },
    },
    createGroup: {
      name: "CreateGroup",
      requestType: CreateGroupRequest,
      requestStream: false,
      responseType: Group,
      responseStream: false,
      options: {
        _unknownFields: {
          8410: [new Uint8Array([5, 103, 114, 111, 117, 112])],
          800010: [new Uint8Array([16, 98, 98, 46, 103, 114, 111, 117, 112, 115, 46, 99, 114, 101, 97, 116, 101])],
          800016: [new Uint8Array([1])],
          800024: [new Uint8Array([1])],
          578365826: [
            new Uint8Array([19, 58, 5, 103, 114, 111, 117, 112, 34, 10, 47, 118, 49, 47, 103, 114, 111, 117, 112, 115]),
          ],
        },
      },
    },
    /**
     * UpdateGroup updates the group.
     * Users with "bb.groups.update" permission on the workspace or the group owner can access this method.
     */
    updateGroup: {
      name: "UpdateGroup",
      requestType: UpdateGroupRequest,
      requestStream: false,
      responseType: Group,
      responseStream: false,
      options: {
        _unknownFields: {
          8410: [new Uint8Array([17, 103, 114, 111, 117, 112, 44, 117, 112, 100, 97, 116, 101, 95, 109, 97, 115, 107])],
          800010: [new Uint8Array([16, 98, 98, 46, 103, 114, 111, 117, 112, 115, 46, 117, 112, 100, 97, 116, 101])],
          800016: [new Uint8Array([2])],
          800024: [new Uint8Array([1])],
          578365826: [
            new Uint8Array([
              34,
              58,
              5,
              103,
              114,
              111,
              117,
              112,
              50,
              25,
              47,
              118,
              49,
              47,
              123,
              103,
              114,
              111,
              117,
              112,
              46,
              110,
              97,
              109,
              101,
              61,
              103,
              114,
              111,
              117,
              112,
              115,
              47,
              42,
              125,
            ]),
          ],
        },
      },
    },
    deleteGroup: {
      name: "DeleteGroup",
      requestType: DeleteGroupRequest,
      requestStream: false,
      responseType: Empty,
      responseStream: false,
      options: {
        _unknownFields: {
          8410: [new Uint8Array([4, 110, 97, 109, 101])],
          800010: [new Uint8Array([16, 98, 98, 46, 103, 114, 111, 117, 112, 115, 46, 100, 101, 108, 101, 116, 101])],
          800016: [new Uint8Array([2])],
          800024: [new Uint8Array([1])],
          578365826: [
            new Uint8Array([
              21,
              42,
              19,
              47,
              118,
              49,
              47,
              123,
              110,
              97,
              109,
              101,
              61,
              103,
              114,
              111,
              117,
              112,
              115,
              47,
              42,
              125,
            ]),
          ],
        },
      },
    },
  },
} as const;

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Long ? string | number | Long : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function toTimestamp(date: Date): Timestamp {
  const seconds = numberToLong(Math.trunc(date.getTime() / 1_000));
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = (t.seconds.toNumber() || 0) * 1_000;
  millis += (t.nanos || 0) / 1_000_000;
  return new globalThis.Date(millis);
}

function fromJsonTimestamp(o: any): Timestamp {
  if (o instanceof globalThis.Date) {
    return toTimestamp(o);
  } else if (typeof o === "string") {
    return toTimestamp(new globalThis.Date(o));
  } else {
    return Timestamp.fromJSON(o);
  }
}

function numberToLong(number: number) {
  return Long.fromNumber(number);
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

export interface MessageFns<T> {
  encode(message: T, writer?: BinaryWriter): BinaryWriter;
  decode(input: BinaryReader | Uint8Array, length?: number): T;
  fromJSON(object: any): T;
  toJSON(message: T): unknown;
  create(base?: DeepPartial<T>): T;
  fromPartial(object: DeepPartial<T>): T;
}
