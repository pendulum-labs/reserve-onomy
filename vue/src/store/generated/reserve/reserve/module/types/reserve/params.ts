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
  /** reserve burn coin */
  burn_coin: string;
  /**
   * Vault liquidator receives reward for succesfully identifying and triggering
   * vault liquidations.  The reward is defined as a percentage.
   * liquidator_reward is (parameter / 10000), 9999 representing 99.99%
   */
  liquidator_reward: string;
}

const baseParams: object = {
  burn_rate: "",
  burn_coin: "",
  liquidator_reward: "",
};

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    if (message.burn_rate !== "") {
      writer.uint32(10).string(message.burn_rate);
    }
    if (message.burn_coin !== "") {
      writer.uint32(18).string(message.burn_coin);
    }
    if (message.liquidator_reward !== "") {
      writer.uint32(26).string(message.liquidator_reward);
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
        case 1:
          message.burn_rate = reader.string();
          break;
        case 2:
          message.burn_coin = reader.string();
          break;
        case 3:
          message.liquidator_reward = reader.string();
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
    if (object.burn_coin !== undefined && object.burn_coin !== null) {
      message.burn_coin = String(object.burn_coin);
    } else {
      message.burn_coin = "";
    }
    if (
      object.liquidator_reward !== undefined &&
      object.liquidator_reward !== null
    ) {
      message.liquidator_reward = String(object.liquidator_reward);
    } else {
      message.liquidator_reward = "";
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.burn_rate !== undefined && (obj.burn_rate = message.burn_rate);
    message.burn_coin !== undefined && (obj.burn_coin = message.burn_coin);
    message.liquidator_reward !== undefined &&
      (obj.liquidator_reward = message.liquidator_reward);
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    if (object.burn_rate !== undefined && object.burn_rate !== null) {
      message.burn_rate = object.burn_rate;
    } else {
      message.burn_rate = "";
    }
    if (object.burn_coin !== undefined && object.burn_coin !== null) {
      message.burn_coin = object.burn_coin;
    } else {
      message.burn_coin = "";
    }
    if (
      object.liquidator_reward !== undefined &&
      object.liquidator_reward !== null
    ) {
      message.liquidator_reward = object.liquidator_reward;
    } else {
      message.liquidator_reward = "";
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
