// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.3.0
//   protoc               unknown
// source: store/sheet.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import Long from "long";
import { Engine, engineFromJSON, engineToJSON, engineToNumber } from "./common";

export const protobufPackage = "bytebase.store";

export interface SheetPayload {
  /** The SQL dialect. */
  engine: Engine;
  /** The start and end position of each command in the sheet statement. */
  commands: SheetCommand[];
}

export interface SheetCommand {
  start: number;
  end: number;
}

function createBaseSheetPayload(): SheetPayload {
  return { engine: Engine.ENGINE_UNSPECIFIED, commands: [] };
}

export const SheetPayload: MessageFns<SheetPayload> = {
  encode(message: SheetPayload, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.engine !== Engine.ENGINE_UNSPECIFIED) {
      writer.uint32(24).int32(engineToNumber(message.engine));
    }
    for (const v of message.commands) {
      SheetCommand.encode(v!, writer.uint32(34).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): SheetPayload {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSheetPayload();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 3: {
          if (tag !== 24) {
            break;
          }

          message.engine = engineFromJSON(reader.int32());
          continue;
        }
        case 4: {
          if (tag !== 34) {
            break;
          }

          message.commands.push(SheetCommand.decode(reader, reader.uint32()));
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

  fromJSON(object: any): SheetPayload {
    return {
      engine: isSet(object.engine) ? engineFromJSON(object.engine) : Engine.ENGINE_UNSPECIFIED,
      commands: globalThis.Array.isArray(object?.commands)
        ? object.commands.map((e: any) => SheetCommand.fromJSON(e))
        : [],
    };
  },

  toJSON(message: SheetPayload): unknown {
    const obj: any = {};
    if (message.engine !== Engine.ENGINE_UNSPECIFIED) {
      obj.engine = engineToJSON(message.engine);
    }
    if (message.commands?.length) {
      obj.commands = message.commands.map((e) => SheetCommand.toJSON(e));
    }
    return obj;
  },

  create(base?: DeepPartial<SheetPayload>): SheetPayload {
    return SheetPayload.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<SheetPayload>): SheetPayload {
    const message = createBaseSheetPayload();
    message.engine = object.engine ?? Engine.ENGINE_UNSPECIFIED;
    message.commands = object.commands?.map((e) => SheetCommand.fromPartial(e)) || [];
    return message;
  },
};

function createBaseSheetCommand(): SheetCommand {
  return { start: 0, end: 0 };
}

export const SheetCommand: MessageFns<SheetCommand> = {
  encode(message: SheetCommand, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.start !== 0) {
      writer.uint32(8).int32(message.start);
    }
    if (message.end !== 0) {
      writer.uint32(16).int32(message.end);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): SheetCommand {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSheetCommand();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 8) {
            break;
          }

          message.start = reader.int32();
          continue;
        }
        case 2: {
          if (tag !== 16) {
            break;
          }

          message.end = reader.int32();
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

  fromJSON(object: any): SheetCommand {
    return {
      start: isSet(object.start) ? globalThis.Number(object.start) : 0,
      end: isSet(object.end) ? globalThis.Number(object.end) : 0,
    };
  },

  toJSON(message: SheetCommand): unknown {
    const obj: any = {};
    if (message.start !== 0) {
      obj.start = Math.round(message.start);
    }
    if (message.end !== 0) {
      obj.end = Math.round(message.end);
    }
    return obj;
  },

  create(base?: DeepPartial<SheetCommand>): SheetCommand {
    return SheetCommand.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<SheetCommand>): SheetCommand {
    const message = createBaseSheetCommand();
    message.start = object.start ?? 0;
    message.end = object.end ?? 0;
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
