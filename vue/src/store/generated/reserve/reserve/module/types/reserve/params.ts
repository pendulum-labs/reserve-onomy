/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "reserve";

/** Params defines the parameters for the module. */
export interface Params {
  /**
   * Upon vault liquidation event collateral is traded for market module
   * burn coin. burn ratio is the ratio of market burn coin burned vs stored
   * in community pool
   * burn rate is (parameter / 10000), 9999 representing as 99.99%
   */
  burn_rate: string;
}

const baseParams: object = { burn_rate: "" };

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    if (message.burn_rate !== "") {
      writer.uint32(34).string(message.burn_rate);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseParams } as Params;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 4:
          message.burn_rate = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params {
    const message = { ...baseParams } as Params;
    if (object.burn_rate !== undefined && object.burn_rate !== null) {
      message.burn_rate = String(object.burn_rate);
    } else {
      message.burn_rate = "";
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.burn_rate !== undefined && (obj.burn_rate = message.burn_rate);
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    if (object.burn_rate !== undefined && object.burn_rate !== null) {
      message.burn_rate = object.burn_rate;
    } else {
      message.burn_rate = "";
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
