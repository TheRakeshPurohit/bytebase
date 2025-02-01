// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.3.0
//   protoc               unknown
// source: store/anomaly.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import Long from "long";

export const protobufPackage = "bytebase.store";

export interface AnomalyConnectionPayload {
  /** Connection failure detail */
  detail: string;
}

export interface AnomalyDatabaseSchemaDriftPayload {
  /** The schema version corresponds to the expected schema */
  version: string;
  /** The expected latest schema stored in the migration history table */
  expect: string;
  /** The actual schema dumped from the database */
  actual: string;
}

function createBaseAnomalyConnectionPayload(): AnomalyConnectionPayload {
  return { detail: "" };
}

export const AnomalyConnectionPayload: MessageFns<AnomalyConnectionPayload> = {
  encode(message: AnomalyConnectionPayload, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.detail !== "") {
      writer.uint32(10).string(message.detail);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): AnomalyConnectionPayload {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAnomalyConnectionPayload();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.detail = reader.string();
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

  fromJSON(object: any): AnomalyConnectionPayload {
    return { detail: isSet(object.detail) ? globalThis.String(object.detail) : "" };
  },

  toJSON(message: AnomalyConnectionPayload): unknown {
    const obj: any = {};
    if (message.detail !== "") {
      obj.detail = message.detail;
    }
    return obj;
  },

  create(base?: DeepPartial<AnomalyConnectionPayload>): AnomalyConnectionPayload {
    return AnomalyConnectionPayload.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<AnomalyConnectionPayload>): AnomalyConnectionPayload {
    const message = createBaseAnomalyConnectionPayload();
    message.detail = object.detail ?? "";
    return message;
  },
};

function createBaseAnomalyDatabaseSchemaDriftPayload(): AnomalyDatabaseSchemaDriftPayload {
  return { version: "", expect: "", actual: "" };
}

export const AnomalyDatabaseSchemaDriftPayload: MessageFns<AnomalyDatabaseSchemaDriftPayload> = {
  encode(message: AnomalyDatabaseSchemaDriftPayload, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.version !== "") {
      writer.uint32(10).string(message.version);
    }
    if (message.expect !== "") {
      writer.uint32(18).string(message.expect);
    }
    if (message.actual !== "") {
      writer.uint32(26).string(message.actual);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): AnomalyDatabaseSchemaDriftPayload {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAnomalyDatabaseSchemaDriftPayload();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.version = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.expect = reader.string();
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.actual = reader.string();
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

  fromJSON(object: any): AnomalyDatabaseSchemaDriftPayload {
    return {
      version: isSet(object.version) ? globalThis.String(object.version) : "",
      expect: isSet(object.expect) ? globalThis.String(object.expect) : "",
      actual: isSet(object.actual) ? globalThis.String(object.actual) : "",
    };
  },

  toJSON(message: AnomalyDatabaseSchemaDriftPayload): unknown {
    const obj: any = {};
    if (message.version !== "") {
      obj.version = message.version;
    }
    if (message.expect !== "") {
      obj.expect = message.expect;
    }
    if (message.actual !== "") {
      obj.actual = message.actual;
    }
    return obj;
  },

  create(base?: DeepPartial<AnomalyDatabaseSchemaDriftPayload>): AnomalyDatabaseSchemaDriftPayload {
    return AnomalyDatabaseSchemaDriftPayload.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<AnomalyDatabaseSchemaDriftPayload>): AnomalyDatabaseSchemaDriftPayload {
    const message = createBaseAnomalyDatabaseSchemaDriftPayload();
    message.version = object.version ?? "";
    message.expect = object.expect ?? "";
    message.actual = object.actual ?? "";
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Long ? string | number | Long : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

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
