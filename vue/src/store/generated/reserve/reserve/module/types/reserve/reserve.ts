/* eslint-disable */
import { Coin } from "../cosmos/base/v1beta1/coin";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "reserve";

export interface Reserve {
  reserves: Coin[];
}

const baseReserve: object = {};

export const Reserve = {
  encode(message: Reserve, writer: Writer = Writer.create()): Writer {
    for (const v of message.reserves) {
      Coin.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Reserve {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseReserve } as Reserve;
    message.reserves = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.reserves.push(Coin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Reserve {
    const message = { ...baseReserve } as Reserve;
    message.reserves = [];
    if (object.reserves !== undefined && object.reserves !== null) {
      for (const e of object.reserves) {
        message.reserves.push(Coin.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: Reserve): unknown {
    const obj: any = {};
    if (message.reserves) {
      obj.reserves = message.reserves.map((e) =>
        e ? Coin.toJSON(e) : undefined
      );
    } else {
      obj.reserves = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Reserve>): Reserve {
    const message = { ...baseReserve } as Reserve;
    message.reserves = [];
    if (object.reserves !== undefined && object.reserves !== null) {
      for (const e of object.reserves) {
        message.reserves.push(Coin.fromPartial(e));
      }
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
