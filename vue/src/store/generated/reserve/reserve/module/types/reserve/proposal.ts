/* eslint-disable */
import { Metadata } from "../cosmos/bank/v1beta1/bank";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "reserve";

/** CreateDenomProposal details a create-denom proposal. */
export interface CreateDenomProposal {
  sender: string;
  title: string;
  description: string;
  metadata: Metadata | undefined;
  exchange_rate: string[];
  collateral_deposit: string;
}

/** RegisterCollateralProposal details a register-collateral proposal. */
export interface RegisterCollateralProposal {
  sender: string;
  title: string;
  description: string;
  metadata: Metadata | undefined;
  minimum_deposit: string;
  minting_ratio: string;
  liquidation_ratio: string;
}

const baseCreateDenomProposal: object = {
  sender: "",
  title: "",
  description: "",
  exchange_rate: "",
  collateral_deposit: "",
};

export const CreateDenomProposal = {
  encode(
    message: CreateDenomProposal,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.sender !== "") {
      writer.uint32(10).string(message.sender);
    }
    if (message.title !== "") {
      writer.uint32(18).string(message.title);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    if (message.metadata !== undefined) {
      Metadata.encode(message.metadata, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.exchange_rate) {
      writer.uint32(42).string(v!);
    }
    if (message.collateral_deposit !== "") {
      writer.uint32(50).string(message.collateral_deposit);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): CreateDenomProposal {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseCreateDenomProposal } as CreateDenomProposal;
    message.exchange_rate = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sender = reader.string();
          break;
        case 2:
          message.title = reader.string();
          break;
        case 3:
          message.description = reader.string();
          break;
        case 4:
          message.metadata = Metadata.decode(reader, reader.uint32());
          break;
        case 5:
          message.exchange_rate.push(reader.string());
          break;
        case 6:
          message.collateral_deposit = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateDenomProposal {
    const message = { ...baseCreateDenomProposal } as CreateDenomProposal;
    message.exchange_rate = [];
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = String(object.sender);
    } else {
      message.sender = "";
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = String(object.title);
    } else {
      message.title = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      message.metadata = Metadata.fromJSON(object.metadata);
    } else {
      message.metadata = undefined;
    }
    if (object.exchange_rate !== undefined && object.exchange_rate !== null) {
      for (const e of object.exchange_rate) {
        message.exchange_rate.push(String(e));
      }
    }
    if (
      object.collateral_deposit !== undefined &&
      object.collateral_deposit !== null
    ) {
      message.collateral_deposit = String(object.collateral_deposit);
    } else {
      message.collateral_deposit = "";
    }
    return message;
  },

  toJSON(message: CreateDenomProposal): unknown {
    const obj: any = {};
    message.sender !== undefined && (obj.sender = message.sender);
    message.title !== undefined && (obj.title = message.title);
    message.description !== undefined &&
      (obj.description = message.description);
    message.metadata !== undefined &&
      (obj.metadata = message.metadata
        ? Metadata.toJSON(message.metadata)
        : undefined);
    if (message.exchange_rate) {
      obj.exchange_rate = message.exchange_rate.map((e) => e);
    } else {
      obj.exchange_rate = [];
    }
    message.collateral_deposit !== undefined &&
      (obj.collateral_deposit = message.collateral_deposit);
    return obj;
  },

  fromPartial(object: DeepPartial<CreateDenomProposal>): CreateDenomProposal {
    const message = { ...baseCreateDenomProposal } as CreateDenomProposal;
    message.exchange_rate = [];
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = object.sender;
    } else {
      message.sender = "";
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = object.title;
    } else {
      message.title = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      message.metadata = Metadata.fromPartial(object.metadata);
    } else {
      message.metadata = undefined;
    }
    if (object.exchange_rate !== undefined && object.exchange_rate !== null) {
      for (const e of object.exchange_rate) {
        message.exchange_rate.push(e);
      }
    }
    if (
      object.collateral_deposit !== undefined &&
      object.collateral_deposit !== null
    ) {
      message.collateral_deposit = object.collateral_deposit;
    } else {
      message.collateral_deposit = "";
    }
    return message;
  },
};

const baseRegisterCollateralProposal: object = {
  sender: "",
  title: "",
  description: "",
  minimum_deposit: "",
  minting_ratio: "",
  liquidation_ratio: "",
};

export const RegisterCollateralProposal = {
  encode(
    message: RegisterCollateralProposal,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.sender !== "") {
      writer.uint32(10).string(message.sender);
    }
    if (message.title !== "") {
      writer.uint32(18).string(message.title);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    if (message.metadata !== undefined) {
      Metadata.encode(message.metadata, writer.uint32(34).fork()).ldelim();
    }
    if (message.minimum_deposit !== "") {
      writer.uint32(42).string(message.minimum_deposit);
    }
    if (message.minting_ratio !== "") {
      writer.uint32(50).string(message.minting_ratio);
    }
    if (message.liquidation_ratio !== "") {
      writer.uint32(58).string(message.liquidation_ratio);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): RegisterCollateralProposal {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseRegisterCollateralProposal,
    } as RegisterCollateralProposal;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sender = reader.string();
          break;
        case 2:
          message.title = reader.string();
          break;
        case 3:
          message.description = reader.string();
          break;
        case 4:
          message.metadata = Metadata.decode(reader, reader.uint32());
          break;
        case 5:
          message.minimum_deposit = reader.string();
          break;
        case 6:
          message.minting_ratio = reader.string();
          break;
        case 7:
          message.liquidation_ratio = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RegisterCollateralProposal {
    const message = {
      ...baseRegisterCollateralProposal,
    } as RegisterCollateralProposal;
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = String(object.sender);
    } else {
      message.sender = "";
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = String(object.title);
    } else {
      message.title = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      message.metadata = Metadata.fromJSON(object.metadata);
    } else {
      message.metadata = undefined;
    }
    if (
      object.minimum_deposit !== undefined &&
      object.minimum_deposit !== null
    ) {
      message.minimum_deposit = String(object.minimum_deposit);
    } else {
      message.minimum_deposit = "";
    }
    if (object.minting_ratio !== undefined && object.minting_ratio !== null) {
      message.minting_ratio = String(object.minting_ratio);
    } else {
      message.minting_ratio = "";
    }
    if (
      object.liquidation_ratio !== undefined &&
      object.liquidation_ratio !== null
    ) {
      message.liquidation_ratio = String(object.liquidation_ratio);
    } else {
      message.liquidation_ratio = "";
    }
    return message;
  },

  toJSON(message: RegisterCollateralProposal): unknown {
    const obj: any = {};
    message.sender !== undefined && (obj.sender = message.sender);
    message.title !== undefined && (obj.title = message.title);
    message.description !== undefined &&
      (obj.description = message.description);
    message.metadata !== undefined &&
      (obj.metadata = message.metadata
        ? Metadata.toJSON(message.metadata)
        : undefined);
    message.minimum_deposit !== undefined &&
      (obj.minimum_deposit = message.minimum_deposit);
    message.minting_ratio !== undefined &&
      (obj.minting_ratio = message.minting_ratio);
    message.liquidation_ratio !== undefined &&
      (obj.liquidation_ratio = message.liquidation_ratio);
    return obj;
  },

  fromPartial(
    object: DeepPartial<RegisterCollateralProposal>
  ): RegisterCollateralProposal {
    const message = {
      ...baseRegisterCollateralProposal,
    } as RegisterCollateralProposal;
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = object.sender;
    } else {
      message.sender = "";
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = object.title;
    } else {
      message.title = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      message.metadata = Metadata.fromPartial(object.metadata);
    } else {
      message.metadata = undefined;
    }
    if (
      object.minimum_deposit !== undefined &&
      object.minimum_deposit !== null
    ) {
      message.minimum_deposit = object.minimum_deposit;
    } else {
      message.minimum_deposit = "";
    }
    if (object.minting_ratio !== undefined && object.minting_ratio !== null) {
      message.minting_ratio = object.minting_ratio;
    } else {
      message.minting_ratio = "";
    }
    if (
      object.liquidation_ratio !== undefined &&
      object.liquidation_ratio !== null
    ) {
      message.liquidation_ratio = object.liquidation_ratio;
    } else {
      message.liquidation_ratio = "";
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
