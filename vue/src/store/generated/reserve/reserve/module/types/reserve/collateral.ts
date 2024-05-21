/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "reserve";

export interface Collateral {
  base: string;
  display: string;
  minimum_deposit: string;
  minting_ratio: string;
  liquidation_ratio: string;
  init_time: number;
}

const baseCollateral: object = {
  base: "",
  display: "",
  minimum_deposit: "",
  minting_ratio: "",
  liquidation_ratio: "",
  init_time: 0,
};

export const Collateral = {
  encode(message: Collateral, writer: Writer = Writer.create()): Writer {
    if (message.base !== "") {
      writer.uint32(10).string(message.base);
    }
    if (message.display !== "") {
      writer.uint32(18).string(message.display);
    }
    if (message.minimum_deposit !== "") {
      writer.uint32(26).string(message.minimum_deposit);
    }
    if (message.minting_ratio !== "") {
      writer.uint32(34).string(message.minting_ratio);
    }
    if (message.liquidation_ratio !== "") {
      writer.uint32(42).string(message.liquidation_ratio);
    }
    if (message.init_time !== 0) {
      writer.uint32(48).int64(message.init_time);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Collateral {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseCollateral } as Collateral;
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
          message.minimum_deposit = reader.string();
          break;
        case 4:
          message.minting_ratio = reader.string();
          break;
        case 5:
          message.liquidation_ratio = reader.string();
          break;
        case 6:
          message.init_time = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Collateral {
    const message = { ...baseCollateral } as Collateral;
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
    if (
      object.minimum_deposit !== undefined &&
      object.minimum_deposit !== null
    ) {
      message.minimum_deposit = String(object.minimum_deposit);
    } else {
      message.minimum_deposit = "";
    }
    if (object.minting_ratio !== undefined && object.minting_ratio !== null) {
      message.minting_ratio = String(object.minting_ratio);
    } else {
      message.minting_ratio = "";
    }
    if (
      object.liquidation_ratio !== undefined &&
      object.liquidation_ratio !== null
    ) {
      message.liquidation_ratio = String(object.liquidation_ratio);
    } else {
      message.liquidation_ratio = "";
    }
    if (object.init_time !== undefined && object.init_time !== null) {
      message.init_time = Number(object.init_time);
    } else {
      message.init_time = 0;
    }
    return message;
  },

  toJSON(message: Collateral): unknown {
    const obj: any = {};
    message.base !== undefined && (obj.base = message.base);
    message.display !== undefined && (obj.display = message.display);
    message.minimum_deposit !== undefined &&
      (obj.minimum_deposit = message.minimum_deposit);
    message.minting_ratio !== undefined &&
      (obj.minting_ratio = message.minting_ratio);
    message.liquidation_ratio !== undefined &&
      (obj.liquidation_ratio = message.liquidation_ratio);
    message.init_time !== undefined && (obj.init_time = message.init_time);
    return obj;
  },

  fromPartial(object: DeepPartial<Collateral>): Collateral {
    const message = { ...baseCollateral } as Collateral;
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
    if (
      object.minimum_deposit !== undefined &&
      object.minimum_deposit !== null
    ) {
      message.minimum_deposit = object.minimum_deposit;
    } else {
      message.minimum_deposit = "";
    }
    if (object.minting_ratio !== undefined && object.minting_ratio !== null) {
      message.minting_ratio = object.minting_ratio;
    } else {
      message.minting_ratio = "";
    }
    if (
      object.liquidation_ratio !== undefined &&
      object.liquidation_ratio !== null
    ) {
      message.liquidation_ratio = object.liquidation_ratio;
    } else {
      message.liquidation_ratio = "";
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
