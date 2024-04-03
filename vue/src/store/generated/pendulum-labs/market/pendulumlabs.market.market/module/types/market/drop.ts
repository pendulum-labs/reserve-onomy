/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "pendulumlabs.market.market";

export interface Drop {
  uid: number;
  owner: string;
  pair: string;
  drops: string;
  product: string;
  active: boolean;
}

export interface Drops {
  uids: number[];
  sum: string;
}

export interface DropPairs {
  pairs: string[];
}

const baseDrop: object = {
  uid: 0,
  owner: "",
  pair: "",
  drops: "",
  product: "",
  active: false,
};

export const Drop = {
  encode(message: Drop, writer: Writer = Writer.create()): Writer {
    if (message.uid !== 0) {
      writer.uint32(8).uint64(message.uid);
    }
    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }
    if (message.pair !== "") {
      writer.uint32(26).string(message.pair);
    }
    if (message.drops !== "") {
      writer.uint32(34).string(message.drops);
    }
    if (message.product !== "") {
      writer.uint32(42).string(message.product);
    }
    if (message.active === true) {
      writer.uint32(48).bool(message.active);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Drop {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDrop } as Drop;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.uid = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.owner = reader.string();
          break;
        case 3:
          message.pair = reader.string();
          break;
        case 4:
          message.drops = reader.string();
          break;
        case 5:
          message.product = reader.string();
          break;
        case 6:
          message.active = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Drop {
    const message = { ...baseDrop } as Drop;
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
    if (object.pair !== undefined && object.pair !== null) {
      message.pair = String(object.pair);
    } else {
      message.pair = "";
    }
    if (object.drops !== undefined && object.drops !== null) {
      message.drops = String(object.drops);
    } else {
      message.drops = "";
    }
    if (object.product !== undefined && object.product !== null) {
      message.product = String(object.product);
    } else {
      message.product = "";
    }
    if (object.active !== undefined && object.active !== null) {
      message.active = Boolean(object.active);
    } else {
      message.active = false;
    }
    return message;
  },

  toJSON(message: Drop): unknown {
    const obj: any = {};
    message.uid !== undefined && (obj.uid = message.uid);
    message.owner !== undefined && (obj.owner = message.owner);
    message.pair !== undefined && (obj.pair = message.pair);
    message.drops !== undefined && (obj.drops = message.drops);
    message.product !== undefined && (obj.product = message.product);
    message.active !== undefined && (obj.active = message.active);
    return obj;
  },

  fromPartial(object: DeepPartial<Drop>): Drop {
    const message = { ...baseDrop } as Drop;
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
    if (object.pair !== undefined && object.pair !== null) {
      message.pair = object.pair;
    } else {
      message.pair = "";
    }
    if (object.drops !== undefined && object.drops !== null) {
      message.drops = object.drops;
    } else {
      message.drops = "";
    }
    if (object.product !== undefined && object.product !== null) {
      message.product = object.product;
    } else {
      message.product = "";
    }
    if (object.active !== undefined && object.active !== null) {
      message.active = object.active;
    } else {
      message.active = false;
    }
    return message;
  },
};

const baseDrops: object = { uids: 0, sum: "" };

export const Drops = {
  encode(message: Drops, writer: Writer = Writer.create()): Writer {
    writer.uint32(10).fork();
    for (const v of message.uids) {
      writer.uint64(v);
    }
    writer.ldelim();
    if (message.sum !== "") {
      writer.uint32(18).string(message.sum);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Drops {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDrops } as Drops;
    message.uids = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.uids.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.uids.push(longToNumber(reader.uint64() as Long));
          }
          break;
        case 2:
          message.sum = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Drops {
    const message = { ...baseDrops } as Drops;
    message.uids = [];
    if (object.uids !== undefined && object.uids !== null) {
      for (const e of object.uids) {
        message.uids.push(Number(e));
      }
    }
    if (object.sum !== undefined && object.sum !== null) {
      message.sum = String(object.sum);
    } else {
      message.sum = "";
    }
    return message;
  },

  toJSON(message: Drops): unknown {
    const obj: any = {};
    if (message.uids) {
      obj.uids = message.uids.map((e) => e);
    } else {
      obj.uids = [];
    }
    message.sum !== undefined && (obj.sum = message.sum);
    return obj;
  },

  fromPartial(object: DeepPartial<Drops>): Drops {
    const message = { ...baseDrops } as Drops;
    message.uids = [];
    if (object.uids !== undefined && object.uids !== null) {
      for (const e of object.uids) {
        message.uids.push(e);
      }
    }
    if (object.sum !== undefined && object.sum !== null) {
      message.sum = object.sum;
    } else {
      message.sum = "";
    }
    return message;
  },
};

const baseDropPairs: object = { pairs: "" };

export const DropPairs = {
  encode(message: DropPairs, writer: Writer = Writer.create()): Writer {
    for (const v of message.pairs) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): DropPairs {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDropPairs } as DropPairs;
    message.pairs = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pairs.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DropPairs {
    const message = { ...baseDropPairs } as DropPairs;
    message.pairs = [];
    if (object.pairs !== undefined && object.pairs !== null) {
      for (const e of object.pairs) {
        message.pairs.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: DropPairs): unknown {
    const obj: any = {};
    if (message.pairs) {
      obj.pairs = message.pairs.map((e) => e);
    } else {
      obj.pairs = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<DropPairs>): DropPairs {
    const message = { ...baseDropPairs } as DropPairs;
    message.pairs = [];
    if (object.pairs !== undefined && object.pairs !== null) {
      for (const e of object.pairs) {
        message.pairs.push(e);
      }
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
