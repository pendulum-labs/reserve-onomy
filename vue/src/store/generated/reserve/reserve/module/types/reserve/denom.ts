/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "reserve";

export interface Denom {
  base: string;
  display: string;
  init_time: number;
}

const baseDenom: object = { base: "", display: "", init_time: 0 };

export const Denom = {
  encode(message: Denom, writer: Writer = Writer.create()): Writer {
    if (message.base !== "") {
      writer.uint32(10).string(message.base);
    }
    if (message.display !== "") {
      writer.uint32(18).string(message.display);
    }
    if (message.init_time !== 0) {
      writer.uint32(24).int64(message.init_time);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Denom {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDenom } as Denom;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.base = reader.string();
          break;
        case 2:
          message.display = reader.string();
          break;
        case 3:
          message.init_time = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Denom {
    const message = { ...baseDenom } as Denom;
    if (object.base !== undefined && object.base !== null) {
      message.base = String(object.base);
    } else {
      message.base = "";
    }
    if (object.display !== undefined && object.display !== null) {
      message.display = String(object.display);
    } else {
      message.display = "";
    }
    if (object.init_time !== undefined && object.init_time !== null) {
      message.init_time = Number(object.init_time);
    } else {
      message.init_time = 0;
    }
    return message;
  },

  toJSON(message: Denom): unknown {
    const obj: any = {};
    message.base !== undefined && (obj.base = message.base);
    message.display !== undefined && (obj.display = message.display);
    message.init_time !== undefined && (obj.init_time = message.init_time);
    return obj;
  },

  fromPartial(object: DeepPartial<Denom>): Denom {
    const message = { ...baseDenom } as Denom;
    if (object.base !== undefined && object.base !== null) {
      message.base = object.base;
    } else {
      message.base = "";
    }
    if (object.display !== undefined && object.display !== null) {
      message.display = object.display;
    } else {
      message.display = "";
    }
    if (object.init_time !== undefined && object.init_time !== null) {
      message.init_time = object.init_time;
    } else {
      message.init_time = 0;
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
