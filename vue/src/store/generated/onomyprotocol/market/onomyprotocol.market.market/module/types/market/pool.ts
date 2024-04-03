/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "pendulum-labs.market.market";

export interface Pool {
  pair: string;
  denom1: string;
  denom2: string;
  leader: string;
  drops: string;
}

const basePool: object = {
  pair: "",
  denom1: "",
  denom2: "",
  leader: "",
  drops: "",
};

export const Pool = {
  encode(message: Pool, writer: Writer = Writer.create()): Writer {
    if (message.pair !== "") {
      writer.uint32(10).string(message.pair);
    }
    if (message.denom1 !== "") {
      writer.uint32(18).string(message.denom1);
    }
    if (message.denom2 !== "") {
      writer.uint32(26).string(message.denom2);
    }
    if (message.leader !== "") {
      writer.uint32(34).string(message.leader);
    }
    if (message.drops !== "") {
      writer.uint32(42).string(message.drops);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Pool {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basePool } as Pool;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pair = reader.string();
          break;
        case 2:
          message.denom1 = reader.string();
          break;
        case 3:
          message.denom2 = reader.string();
          break;
        case 4:
          message.leader = reader.string();
          break;
        case 5:
          message.drops = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Pool {
    const message = { ...basePool } as Pool;
    if (object.pair !== undefined && object.pair !== null) {
      message.pair = String(object.pair);
    } else {
      message.pair = "";
    }
    if (object.denom1 !== undefined && object.denom1 !== null) {
      message.denom1 = String(object.denom1);
    } else {
      message.denom1 = "";
    }
    if (object.denom2 !== undefined && object.denom2 !== null) {
      message.denom2 = String(object.denom2);
    } else {
      message.denom2 = "";
    }
    if (object.leader !== undefined && object.leader !== null) {
      message.leader = String(object.leader);
    } else {
      message.leader = "";
    }
    if (object.drops !== undefined && object.drops !== null) {
      message.drops = String(object.drops);
    } else {
      message.drops = "";
    }
    return message;
  },

  toJSON(message: Pool): unknown {
    const obj: any = {};
    message.pair !== undefined && (obj.pair = message.pair);
    message.denom1 !== undefined && (obj.denom1 = message.denom1);
    message.denom2 !== undefined && (obj.denom2 = message.denom2);
    message.leader !== undefined && (obj.leader = message.leader);
    message.drops !== undefined && (obj.drops = message.drops);
    return obj;
  },

  fromPartial(object: DeepPartial<Pool>): Pool {
    const message = { ...basePool } as Pool;
    if (object.pair !== undefined && object.pair !== null) {
      message.pair = object.pair;
    } else {
      message.pair = "";
    }
    if (object.denom1 !== undefined && object.denom1 !== null) {
      message.denom1 = object.denom1;
    } else {
      message.denom1 = "";
    }
    if (object.denom2 !== undefined && object.denom2 !== null) {
      message.denom2 = object.denom2;
    } else {
      message.denom2 = "";
    }
    if (object.leader !== undefined && object.leader !== null) {
      message.leader = object.leader;
    } else {
      message.leader = "";
    }
    if (object.drops !== undefined && object.drops !== null) {
      message.drops = object.drops;
    } else {
      message.drops = "";
    }
    return message;
  },
};

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
