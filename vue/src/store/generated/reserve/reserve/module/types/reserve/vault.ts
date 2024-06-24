/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Coin } from "../cosmos/base/v1beta1/coin";

export const protobufPackage = "reserve";

export interface Vault {
  uid: number;
  owner: string;
  status: string;
  collateral: Coin | undefined;
  denom: Coin | undefined;
}

export interface VaultMap {
  uid: number;
}

const baseVault: object = { uid: 0, owner: "", status: "" };

export const Vault = {
  encode(message: Vault, writer: Writer = Writer.create()): Writer {
    if (message.uid !== 0) {
      writer.uint32(8).uint64(message.uid);
    }
    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }
    if (message.status !== "") {
      writer.uint32(34).string(message.status);
    }
    if (message.collateral !== undefined) {
      Coin.encode(message.collateral, writer.uint32(42).fork()).ldelim();
    }
    if (message.denom !== undefined) {
      Coin.encode(message.denom, writer.uint32(50).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Vault {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseVault } as Vault;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.uid = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.owner = reader.string();
          break;
        case 4:
          message.status = reader.string();
          break;
        case 5:
          message.collateral = Coin.decode(reader, reader.uint32());
          break;
        case 6:
          message.denom = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Vault {
    const message = { ...baseVault } as Vault;
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = Number(object.uid);
    } else {
      message.uid = 0;
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = String(object.status);
    } else {
      message.status = "";
    }
    if (object.collateral !== undefined && object.collateral !== null) {
      message.collateral = Coin.fromJSON(object.collateral);
    } else {
      message.collateral = undefined;
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = Coin.fromJSON(object.denom);
    } else {
      message.denom = undefined;
    }
    return message;
  },

  toJSON(message: Vault): unknown {
    const obj: any = {};
    message.uid !== undefined && (obj.uid = message.uid);
    message.owner !== undefined && (obj.owner = message.owner);
    message.status !== undefined && (obj.status = message.status);
    message.collateral !== undefined &&
      (obj.collateral = message.collateral
        ? Coin.toJSON(message.collateral)
        : undefined);
    message.denom !== undefined &&
      (obj.denom = message.denom ? Coin.toJSON(message.denom) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<Vault>): Vault {
    const message = { ...baseVault } as Vault;
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = object.uid;
    } else {
      message.uid = 0;
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = "";
    }
    if (object.collateral !== undefined && object.collateral !== null) {
      message.collateral = Coin.fromPartial(object.collateral);
    } else {
      message.collateral = undefined;
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = Coin.fromPartial(object.denom);
    } else {
      message.denom = undefined;
    }
    return message;
  },
};

const baseVaultMap: object = { uid: 0 };

export const VaultMap = {
  encode(message: VaultMap, writer: Writer = Writer.create()): Writer {
    if (message.uid !== 0) {
      writer.uint32(8).uint64(message.uid);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): VaultMap {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseVaultMap } as VaultMap;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.uid = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VaultMap {
    const message = { ...baseVaultMap } as VaultMap;
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = Number(object.uid);
    } else {
      message.uid = 0;
    }
    return message;
  },

  toJSON(message: VaultMap): unknown {
    const obj: any = {};
    message.uid !== undefined && (obj.uid = message.uid);
    return obj;
  },

  fromPartial(object: DeepPartial<VaultMap>): VaultMap {
    const message = { ...baseVaultMap } as VaultMap;
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = object.uid;
    } else {
      message.uid = 0;
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
