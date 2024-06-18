/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "reserve";

export interface MsgCreateVault {
  creator: string;
  collateral: string;
  name: string;
}

export interface MsgCreateVaultResponse {
  uid: number;
}

export interface MsgDeposit {
  creator: string;
  uid: string;
  coin: string;
}

export interface MsgDepositResponse {
  uid: string;
  coin: string;
}

export interface MsgWithdraw {
  creator: string;
  uid: string;
  coin: string;
}

export interface MsgWithdrawResponse {}

export interface MsgLiquidate {
  creator: string;
  uid: string;
}

export interface MsgLiquidateResponse {}

const baseMsgCreateVault: object = { creator: "", collateral: "", name: "" };

export const MsgCreateVault = {
  encode(message: MsgCreateVault, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.collateral !== "") {
      writer.uint32(18).string(message.collateral);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateVault {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateVault } as MsgCreateVault;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.collateral = reader.string();
          break;
        case 3:
          message.name = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateVault {
    const message = { ...baseMsgCreateVault } as MsgCreateVault;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.collateral !== undefined && object.collateral !== null) {
      message.collateral = String(object.collateral);
    } else {
      message.collateral = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    return message;
  },

  toJSON(message: MsgCreateVault): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.collateral !== undefined && (obj.collateral = message.collateral);
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateVault>): MsgCreateVault {
    const message = { ...baseMsgCreateVault } as MsgCreateVault;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.collateral !== undefined && object.collateral !== null) {
      message.collateral = object.collateral;
    } else {
      message.collateral = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    return message;
  },
};

const baseMsgCreateVaultResponse: object = { uid: 0 };

export const MsgCreateVaultResponse = {
  encode(
    message: MsgCreateVaultResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.uid !== 0) {
      writer.uint32(8).uint64(message.uid);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateVaultResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateVaultResponse } as MsgCreateVaultResponse;
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

  fromJSON(object: any): MsgCreateVaultResponse {
    const message = { ...baseMsgCreateVaultResponse } as MsgCreateVaultResponse;
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = Number(object.uid);
    } else {
      message.uid = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateVaultResponse): unknown {
    const obj: any = {};
    message.uid !== undefined && (obj.uid = message.uid);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateVaultResponse>
  ): MsgCreateVaultResponse {
    const message = { ...baseMsgCreateVaultResponse } as MsgCreateVaultResponse;
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = object.uid;
    } else {
      message.uid = 0;
    }
    return message;
  },
};

const baseMsgDeposit: object = { creator: "", uid: "", coin: "" };

export const MsgDeposit = {
  encode(message: MsgDeposit, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.uid !== "") {
      writer.uint32(18).string(message.uid);
    }
    if (message.coin !== "") {
      writer.uint32(26).string(message.coin);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeposit {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeposit } as MsgDeposit;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.uid = reader.string();
          break;
        case 3:
          message.coin = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeposit {
    const message = { ...baseMsgDeposit } as MsgDeposit;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = String(object.uid);
    } else {
      message.uid = "";
    }
    if (object.coin !== undefined && object.coin !== null) {
      message.coin = String(object.coin);
    } else {
      message.coin = "";
    }
    return message;
  },

  toJSON(message: MsgDeposit): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.uid !== undefined && (obj.uid = message.uid);
    message.coin !== undefined && (obj.coin = message.coin);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeposit>): MsgDeposit {
    const message = { ...baseMsgDeposit } as MsgDeposit;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = object.uid;
    } else {
      message.uid = "";
    }
    if (object.coin !== undefined && object.coin !== null) {
      message.coin = object.coin;
    } else {
      message.coin = "";
    }
    return message;
  },
};

const baseMsgDepositResponse: object = { uid: "", coin: "" };

export const MsgDepositResponse = {
  encode(
    message: MsgDepositResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.uid !== "") {
      writer.uint32(10).string(message.uid);
    }
    if (message.coin !== "") {
      writer.uint32(18).string(message.coin);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDepositResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDepositResponse } as MsgDepositResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.uid = reader.string();
          break;
        case 2:
          message.coin = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDepositResponse {
    const message = { ...baseMsgDepositResponse } as MsgDepositResponse;
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = String(object.uid);
    } else {
      message.uid = "";
    }
    if (object.coin !== undefined && object.coin !== null) {
      message.coin = String(object.coin);
    } else {
      message.coin = "";
    }
    return message;
  },

  toJSON(message: MsgDepositResponse): unknown {
    const obj: any = {};
    message.uid !== undefined && (obj.uid = message.uid);
    message.coin !== undefined && (obj.coin = message.coin);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDepositResponse>): MsgDepositResponse {
    const message = { ...baseMsgDepositResponse } as MsgDepositResponse;
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = object.uid;
    } else {
      message.uid = "";
    }
    if (object.coin !== undefined && object.coin !== null) {
      message.coin = object.coin;
    } else {
      message.coin = "";
    }
    return message;
  },
};

const baseMsgWithdraw: object = { creator: "", uid: "", coin: "" };

export const MsgWithdraw = {
  encode(message: MsgWithdraw, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.uid !== "") {
      writer.uint32(18).string(message.uid);
    }
    if (message.coin !== "") {
      writer.uint32(26).string(message.coin);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgWithdraw {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgWithdraw } as MsgWithdraw;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.uid = reader.string();
          break;
        case 3:
          message.coin = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgWithdraw {
    const message = { ...baseMsgWithdraw } as MsgWithdraw;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = String(object.uid);
    } else {
      message.uid = "";
    }
    if (object.coin !== undefined && object.coin !== null) {
      message.coin = String(object.coin);
    } else {
      message.coin = "";
    }
    return message;
  },

  toJSON(message: MsgWithdraw): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.uid !== undefined && (obj.uid = message.uid);
    message.coin !== undefined && (obj.coin = message.coin);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgWithdraw>): MsgWithdraw {
    const message = { ...baseMsgWithdraw } as MsgWithdraw;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = object.uid;
    } else {
      message.uid = "";
    }
    if (object.coin !== undefined && object.coin !== null) {
      message.coin = object.coin;
    } else {
      message.coin = "";
    }
    return message;
  },
};

const baseMsgWithdrawResponse: object = {};

export const MsgWithdrawResponse = {
  encode(_: MsgWithdrawResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgWithdrawResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgWithdrawResponse } as MsgWithdrawResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgWithdrawResponse {
    const message = { ...baseMsgWithdrawResponse } as MsgWithdrawResponse;
    return message;
  },

  toJSON(_: MsgWithdrawResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgWithdrawResponse>): MsgWithdrawResponse {
    const message = { ...baseMsgWithdrawResponse } as MsgWithdrawResponse;
    return message;
  },
};

const baseMsgLiquidate: object = { creator: "", uid: "" };

export const MsgLiquidate = {
  encode(message: MsgLiquidate, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.uid !== "") {
      writer.uint32(18).string(message.uid);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgLiquidate {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgLiquidate } as MsgLiquidate;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.uid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgLiquidate {
    const message = { ...baseMsgLiquidate } as MsgLiquidate;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = String(object.uid);
    } else {
      message.uid = "";
    }
    return message;
  },

  toJSON(message: MsgLiquidate): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.uid !== undefined && (obj.uid = message.uid);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgLiquidate>): MsgLiquidate {
    const message = { ...baseMsgLiquidate } as MsgLiquidate;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = object.uid;
    } else {
      message.uid = "";
    }
    return message;
  },
};

const baseMsgLiquidateResponse: object = {};

export const MsgLiquidateResponse = {
  encode(_: MsgLiquidateResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgLiquidateResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgLiquidateResponse } as MsgLiquidateResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgLiquidateResponse {
    const message = { ...baseMsgLiquidateResponse } as MsgLiquidateResponse;
    return message;
  },

  toJSON(_: MsgLiquidateResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgLiquidateResponse>): MsgLiquidateResponse {
    const message = { ...baseMsgLiquidateResponse } as MsgLiquidateResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateVault(request: MsgCreateVault): Promise<MsgCreateVaultResponse>;
  Deposit(request: MsgDeposit): Promise<MsgDepositResponse>;
  Withdraw(request: MsgWithdraw): Promise<MsgWithdrawResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  Liquidate(request: MsgLiquidate): Promise<MsgLiquidateResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateVault(request: MsgCreateVault): Promise<MsgCreateVaultResponse> {
    const data = MsgCreateVault.encode(request).finish();
    const promise = this.rpc.request("reserve.Msg", "CreateVault", data);
    return promise.then((data) =>
      MsgCreateVaultResponse.decode(new Reader(data))
    );
  }

  Deposit(request: MsgDeposit): Promise<MsgDepositResponse> {
    const data = MsgDeposit.encode(request).finish();
    const promise = this.rpc.request("reserve.Msg", "Deposit", data);
    return promise.then((data) => MsgDepositResponse.decode(new Reader(data)));
  }

  Withdraw(request: MsgWithdraw): Promise<MsgWithdrawResponse> {
    const data = MsgWithdraw.encode(request).finish();
    const promise = this.rpc.request("reserve.Msg", "Withdraw", data);
    return promise.then((data) => MsgWithdrawResponse.decode(new Reader(data)));
  }

  Liquidate(request: MsgLiquidate): Promise<MsgLiquidateResponse> {
    const data = MsgLiquidate.encode(request).finish();
    const promise = this.rpc.request("reserve.Msg", "Liquidate", data);
    return promise.then((data) =>
      MsgLiquidateResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
