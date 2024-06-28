/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "reserve";

export interface Denom {
  base: string;
  display: string;
  init_time: number;
  /** Positive Interest rate on denom debt pool */
  debt_interest_rate: string;
  /** Total shares of the denom debt pool */
  debt_shares: string;
  /** Negative Interest rate on Denoms Bonded */
  bond_interest_rate: string;
  /** Total shares of the denom bond pool */
  bond_shares: string;
  /** Unbonding blocks */
  bond_blocks: string;
}

const baseDenom: object = {
  base: "",
  display: "",
  init_time: 0,
  debt_interest_rate: "",
  debt_shares: "",
  bond_interest_rate: "",
  bond_shares: "",
  bond_blocks: "",
};

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
    if (message.debt_interest_rate !== "") {
      writer.uint32(42).string(message.debt_interest_rate);
    }
    if (message.debt_shares !== "") {
      writer.uint32(50).string(message.debt_shares);
    }
    if (message.bond_interest_rate !== "") {
      writer.uint32(58).string(message.bond_interest_rate);
    }
    if (message.bond_shares !== "") {
      writer.uint32(66).string(message.bond_shares);
    }
    if (message.bond_blocks !== "") {
      writer.uint32(74).string(message.bond_blocks);
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
        case 5:
          message.debt_interest_rate = reader.string();
          break;
        case 6:
          message.debt_shares = reader.string();
          break;
        case 7:
          message.bond_interest_rate = reader.string();
          break;
        case 8:
          message.bond_shares = reader.string();
          break;
        case 9:
          message.bond_blocks = reader.string();
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
    if (
      object.debt_interest_rate !== undefined &&
      object.debt_interest_rate !== null
    ) {
      message.debt_interest_rate = String(object.debt_interest_rate);
    } else {
      message.debt_interest_rate = "";
    }
    if (object.debt_shares !== undefined && object.debt_shares !== null) {
      message.debt_shares = String(object.debt_shares);
    } else {
      message.debt_shares = "";
    }
    if (
      object.bond_interest_rate !== undefined &&
      object.bond_interest_rate !== null
    ) {
      message.bond_interest_rate = String(object.bond_interest_rate);
    } else {
      message.bond_interest_rate = "";
    }
    if (object.bond_shares !== undefined && object.bond_shares !== null) {
      message.bond_shares = String(object.bond_shares);
    } else {
      message.bond_shares = "";
    }
    if (object.bond_blocks !== undefined && object.bond_blocks !== null) {
      message.bond_blocks = String(object.bond_blocks);
    } else {
      message.bond_blocks = "";
    }
    return message;
  },

  toJSON(message: Denom): unknown {
    const obj: any = {};
    message.base !== undefined && (obj.base = message.base);
    message.display !== undefined && (obj.display = message.display);
    message.init_time !== undefined && (obj.init_time = message.init_time);
    message.debt_interest_rate !== undefined &&
      (obj.debt_interest_rate = message.debt_interest_rate);
    message.debt_shares !== undefined &&
      (obj.debt_shares = message.debt_shares);
    message.bond_interest_rate !== undefined &&
      (obj.bond_interest_rate = message.bond_interest_rate);
    message.bond_shares !== undefined &&
      (obj.bond_shares = message.bond_shares);
    message.bond_blocks !== undefined &&
      (obj.bond_blocks = message.bond_blocks);
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
    if (
      object.debt_interest_rate !== undefined &&
      object.debt_interest_rate !== null
    ) {
      message.debt_interest_rate = object.debt_interest_rate;
    } else {
      message.debt_interest_rate = "";
    }
    if (object.debt_shares !== undefined && object.debt_shares !== null) {
      message.debt_shares = object.debt_shares;
    } else {
      message.debt_shares = "";
    }
    if (
      object.bond_interest_rate !== undefined &&
      object.bond_interest_rate !== null
    ) {
      message.bond_interest_rate = object.bond_interest_rate;
    } else {
      message.bond_interest_rate = "";
    }
    if (object.bond_shares !== undefined && object.bond_shares !== null) {
      message.bond_shares = object.bond_shares;
    } else {
      message.bond_shares = "";
    }
    if (object.bond_blocks !== undefined && object.bond_blocks !== null) {
      message.bond_blocks = object.bond_blocks;
    } else {
      message.bond_blocks = "";
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
