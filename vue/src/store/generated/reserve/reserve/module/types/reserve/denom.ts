/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "reserve";

export interface Denom {
  denom_base: string;
  denom_display: string;
  bond_base: string;
  bond_display: string;
  init_time: number;
  /** Positive Interest rate on denom debt pool */
  debt_interest_rate: number;
  /** Total shares of the denom debt pool */
  debt_shares: string;
  /** Total amount of denom in the denom debt pool */
  debt_denoms: string;
  /** Negative Interest rate on Denoms Bonded */
  bond_interest_rate: number;
  /** Total shares of the denom bond pool */
  bond_shares: string;
  /** Total amount of denom in the denom bond pool */
  bond_denoms: string;
  /** Unbonding blocks */
  bond_blocks: string;
}

const baseDenom: object = {
  denom_base: "",
  denom_display: "",
  bond_base: "",
  bond_display: "",
  init_time: 0,
  debt_interest_rate: 0,
  debt_shares: "",
  debt_denoms: "",
  bond_interest_rate: 0,
  bond_shares: "",
  bond_denoms: "",
  bond_blocks: "",
};

export const Denom = {
  encode(message: Denom, writer: Writer = Writer.create()): Writer {
    if (message.denom_base !== "") {
      writer.uint32(10).string(message.denom_base);
    }
    if (message.denom_display !== "") {
      writer.uint32(18).string(message.denom_display);
    }
    if (message.bond_base !== "") {
      writer.uint32(26).string(message.bond_base);
    }
    if (message.bond_display !== "") {
      writer.uint32(34).string(message.bond_display);
    }
    if (message.init_time !== 0) {
      writer.uint32(40).int64(message.init_time);
    }
    if (message.debt_interest_rate !== 0) {
      writer.uint32(48).uint64(message.debt_interest_rate);
    }
    if (message.debt_shares !== "") {
      writer.uint32(58).string(message.debt_shares);
    }
    if (message.debt_denoms !== "") {
      writer.uint32(66).string(message.debt_denoms);
    }
    if (message.bond_interest_rate !== 0) {
      writer.uint32(72).uint64(message.bond_interest_rate);
    }
    if (message.bond_shares !== "") {
      writer.uint32(82).string(message.bond_shares);
    }
    if (message.bond_denoms !== "") {
      writer.uint32(90).string(message.bond_denoms);
    }
    if (message.bond_blocks !== "") {
      writer.uint32(98).string(message.bond_blocks);
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
          message.denom_base = reader.string();
          break;
        case 2:
          message.denom_display = reader.string();
          break;
        case 3:
          message.bond_base = reader.string();
          break;
        case 4:
          message.bond_display = reader.string();
          break;
        case 5:
          message.init_time = longToNumber(reader.int64() as Long);
          break;
        case 6:
          message.debt_interest_rate = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.debt_shares = reader.string();
          break;
        case 8:
          message.debt_denoms = reader.string();
          break;
        case 9:
          message.bond_interest_rate = longToNumber(reader.uint64() as Long);
          break;
        case 10:
          message.bond_shares = reader.string();
          break;
        case 11:
          message.bond_denoms = reader.string();
          break;
        case 12:
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
    if (object.denom_base !== undefined && object.denom_base !== null) {
      message.denom_base = String(object.denom_base);
    } else {
      message.denom_base = "";
    }
    if (object.denom_display !== undefined && object.denom_display !== null) {
      message.denom_display = String(object.denom_display);
    } else {
      message.denom_display = "";
    }
    if (object.bond_base !== undefined && object.bond_base !== null) {
      message.bond_base = String(object.bond_base);
    } else {
      message.bond_base = "";
    }
    if (object.bond_display !== undefined && object.bond_display !== null) {
      message.bond_display = String(object.bond_display);
    } else {
      message.bond_display = "";
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
      message.debt_interest_rate = Number(object.debt_interest_rate);
    } else {
      message.debt_interest_rate = 0;
    }
    if (object.debt_shares !== undefined && object.debt_shares !== null) {
      message.debt_shares = String(object.debt_shares);
    } else {
      message.debt_shares = "";
    }
    if (object.debt_denoms !== undefined && object.debt_denoms !== null) {
      message.debt_denoms = String(object.debt_denoms);
    } else {
      message.debt_denoms = "";
    }
    if (
      object.bond_interest_rate !== undefined &&
      object.bond_interest_rate !== null
    ) {
      message.bond_interest_rate = Number(object.bond_interest_rate);
    } else {
      message.bond_interest_rate = 0;
    }
    if (object.bond_shares !== undefined && object.bond_shares !== null) {
      message.bond_shares = String(object.bond_shares);
    } else {
      message.bond_shares = "";
    }
    if (object.bond_denoms !== undefined && object.bond_denoms !== null) {
      message.bond_denoms = String(object.bond_denoms);
    } else {
      message.bond_denoms = "";
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
    message.denom_base !== undefined && (obj.denom_base = message.denom_base);
    message.denom_display !== undefined &&
      (obj.denom_display = message.denom_display);
    message.bond_base !== undefined && (obj.bond_base = message.bond_base);
    message.bond_display !== undefined &&
      (obj.bond_display = message.bond_display);
    message.init_time !== undefined && (obj.init_time = message.init_time);
    message.debt_interest_rate !== undefined &&
      (obj.debt_interest_rate = message.debt_interest_rate);
    message.debt_shares !== undefined &&
      (obj.debt_shares = message.debt_shares);
    message.debt_denoms !== undefined &&
      (obj.debt_denoms = message.debt_denoms);
    message.bond_interest_rate !== undefined &&
      (obj.bond_interest_rate = message.bond_interest_rate);
    message.bond_shares !== undefined &&
      (obj.bond_shares = message.bond_shares);
    message.bond_denoms !== undefined &&
      (obj.bond_denoms = message.bond_denoms);
    message.bond_blocks !== undefined &&
      (obj.bond_blocks = message.bond_blocks);
    return obj;
  },

  fromPartial(object: DeepPartial<Denom>): Denom {
    const message = { ...baseDenom } as Denom;
    if (object.denom_base !== undefined && object.denom_base !== null) {
      message.denom_base = object.denom_base;
    } else {
      message.denom_base = "";
    }
    if (object.denom_display !== undefined && object.denom_display !== null) {
      message.denom_display = object.denom_display;
    } else {
      message.denom_display = "";
    }
    if (object.bond_base !== undefined && object.bond_base !== null) {
      message.bond_base = object.bond_base;
    } else {
      message.bond_base = "";
    }
    if (object.bond_display !== undefined && object.bond_display !== null) {
      message.bond_display = object.bond_display;
    } else {
      message.bond_display = "";
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
      message.debt_interest_rate = 0;
    }
    if (object.debt_shares !== undefined && object.debt_shares !== null) {
      message.debt_shares = object.debt_shares;
    } else {
      message.debt_shares = "";
    }
    if (object.debt_denoms !== undefined && object.debt_denoms !== null) {
      message.debt_denoms = object.debt_denoms;
    } else {
      message.debt_denoms = "";
    }
    if (
      object.bond_interest_rate !== undefined &&
      object.bond_interest_rate !== null
    ) {
      message.bond_interest_rate = object.bond_interest_rate;
    } else {
      message.bond_interest_rate = 0;
    }
    if (object.bond_shares !== undefined && object.bond_shares !== null) {
      message.bond_shares = object.bond_shares;
    } else {
      message.bond_shares = "";
    }
    if (object.bond_denoms !== undefined && object.bond_denoms !== null) {
      message.bond_denoms = object.bond_denoms;
    } else {
      message.bond_denoms = "";
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
