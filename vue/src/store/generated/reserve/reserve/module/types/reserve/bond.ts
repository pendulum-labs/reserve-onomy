/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "reserve.reserve";

export interface Bond {
  uid: string;
  owner: string;
  denom: string;
  bondShares: string;
}

const baseBond: object = { uid: "", owner: "", denom: "", bondShares: "" };

export const Bond = {
  encode(message: Bond, writer: Writer = Writer.create()): Writer {
    if (message.uid !== "") {
      writer.uint32(10).string(message.uid);
    }
    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }
    if (message.denom !== "") {
      writer.uint32(26).string(message.denom);
    }
    if (message.bondShares !== "") {
      writer.uint32(34).string(message.bondShares);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Bond {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseBond } as Bond;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.uid = reader.string();
          break;
        case 2:
          message.owner = reader.string();
          break;
        case 3:
          message.denom = reader.string();
          break;
        case 4:
          message.bondShares = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Bond {
    const message = { ...baseBond } as Bond;
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = String(object.uid);
    } else {
      message.uid = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom);
    } else {
      message.denom = "";
    }
    if (object.bondShares !== undefined && object.bondShares !== null) {
      message.bondShares = String(object.bondShares);
    } else {
      message.bondShares = "";
    }
    return message;
  },

  toJSON(message: Bond): unknown {
    const obj: any = {};
    message.uid !== undefined && (obj.uid = message.uid);
    message.owner !== undefined && (obj.owner = message.owner);
    message.denom !== undefined && (obj.denom = message.denom);
    message.bondShares !== undefined && (obj.bondShares = message.bondShares);
    return obj;
  },

  fromPartial(object: DeepPartial<Bond>): Bond {
    const message = { ...baseBond } as Bond;
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = object.uid;
    } else {
      message.uid = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom;
    } else {
      message.denom = "";
    }
    if (object.bondShares !== undefined && object.bondShares !== null) {
      message.bondShares = object.bondShares;
    } else {
      message.bondShares = "";
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
