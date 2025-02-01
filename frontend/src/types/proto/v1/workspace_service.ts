// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.3.0
//   protoc               unknown
// source: v1/workspace_service.proto

/* eslint-disable */
import { GetIamPolicyRequest, IamPolicy, SetIamPolicyRequest } from "./iam_policy";

export const protobufPackage = "bytebase.v1";

export type WorkspaceServiceDefinition = typeof WorkspaceServiceDefinition;
export const WorkspaceServiceDefinition = {
  name: "WorkspaceService",
  fullName: "bytebase.v1.WorkspaceService",
  methods: {
    getIamPolicy: {
      name: "GetIamPolicy",
      requestType: GetIamPolicyRequest,
      requestStream: false,
      responseType: IamPolicy,
      responseStream: false,
      options: {
        _unknownFields: {
          800010: [new Uint8Array([15, 98, 98, 46, 112, 111, 108, 105, 99, 105, 101, 115, 46, 103, 101, 116])],
          800016: [new Uint8Array([1])],
          578365826: [
            new Uint8Array([
              42,
              18,
              40,
              47,
              118,
              49,
              47,
              123,
              114,
              101,
              115,
              111,
              117,
              114,
              99,
              101,
              61,
              119,
              111,
              114,
              107,
              115,
              112,
              97,
              99,
              101,
              115,
              47,
              42,
              125,
              58,
              103,
              101,
              116,
              73,
              97,
              109,
              80,
              111,
              108,
              105,
              99,
              121,
            ]),
          ],
        },
      },
    },
    setIamPolicy: {
      name: "SetIamPolicy",
      requestType: SetIamPolicyRequest,
      requestStream: false,
      responseType: IamPolicy,
      responseStream: false,
      options: {
        _unknownFields: {
          800010: [
            new Uint8Array([18, 98, 98, 46, 112, 111, 108, 105, 99, 105, 101, 115, 46, 117, 112, 100, 97, 116, 101]),
          ],
          800016: [new Uint8Array([1])],
          800024: [new Uint8Array([1])],
          578365826: [
            new Uint8Array([
              45,
              58,
              1,
              42,
              34,
              40,
              47,
              118,
              49,
              47,
              123,
              114,
              101,
              115,
              111,
              117,
              114,
              99,
              101,
              61,
              119,
              111,
              114,
              107,
              115,
              112,
              97,
              99,
              101,
              115,
              47,
              42,
              125,
              58,
              115,
              101,
              116,
              73,
              97,
              109,
              80,
              111,
              108,
              105,
              99,
              121,
            ]),
          ],
        },
      },
    },
  },
} as const;
