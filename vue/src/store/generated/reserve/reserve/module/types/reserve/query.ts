/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../reserve/params";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { Vault } from "../reserve/vault";

export const protobufPackage = "reserve";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetAllVaultsRequest {
  pagination: PageRequest | undefined;
}

export interface QueryGetAllVaultsResponse {
  vaults: Vault[];
  pagination: PageResponse | undefined;
}

export interface QueryGetAllVaultsByOwnerRequest {
  pagination: PageRequest | undefined;
  address: string;
}

export interface QueryGetAllVaultsByOwnerResponse {
  vaults: Vault[];
  pagination: PageResponse | undefined;
}

export interface QueryGetAllVaultsInDefaultRequest {
  pagination: PageRequest | undefined;
}

export interface QueryGetAllVaultsInDefaultResponse {
  vaults: Vault[];
  pagination: PageResponse | undefined;
}

export interface QueryGetVaultByUidRequest {
  uid: number;
}

export interface QueryGetVaultByUidResponse {
  vault: Vault | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
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

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetAllVaultsRequest: object = {};

export const QueryGetAllVaultsRequest = {
  encode(
    message: QueryGetAllVaultsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetAllVaultsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetAllVaultsRequest,
    } as QueryGetAllVaultsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetAllVaultsRequest {
    const message = {
      ...baseQueryGetAllVaultsRequest,
    } as QueryGetAllVaultsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetAllVaultsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetAllVaultsRequest>
  ): QueryGetAllVaultsRequest {
    const message = {
      ...baseQueryGetAllVaultsRequest,
    } as QueryGetAllVaultsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetAllVaultsResponse: object = {};

export const QueryGetAllVaultsResponse = {
  encode(
    message: QueryGetAllVaultsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.vaults) {
      Vault.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetAllVaultsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetAllVaultsResponse,
    } as QueryGetAllVaultsResponse;
    message.vaults = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vaults.push(Vault.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetAllVaultsResponse {
    const message = {
      ...baseQueryGetAllVaultsResponse,
    } as QueryGetAllVaultsResponse;
    message.vaults = [];
    if (object.vaults !== undefined && object.vaults !== null) {
      for (const e of object.vaults) {
        message.vaults.push(Vault.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetAllVaultsResponse): unknown {
    const obj: any = {};
    if (message.vaults) {
      obj.vaults = message.vaults.map((e) => (e ? Vault.toJSON(e) : undefined));
    } else {
      obj.vaults = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetAllVaultsResponse>
  ): QueryGetAllVaultsResponse {
    const message = {
      ...baseQueryGetAllVaultsResponse,
    } as QueryGetAllVaultsResponse;
    message.vaults = [];
    if (object.vaults !== undefined && object.vaults !== null) {
      for (const e of object.vaults) {
        message.vaults.push(Vault.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetAllVaultsByOwnerRequest: object = { address: "" };

export const QueryGetAllVaultsByOwnerRequest = {
  encode(
    message: QueryGetAllVaultsByOwnerRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetAllVaultsByOwnerRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetAllVaultsByOwnerRequest,
    } as QueryGetAllVaultsByOwnerRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        case 2:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetAllVaultsByOwnerRequest {
    const message = {
      ...baseQueryGetAllVaultsByOwnerRequest,
    } as QueryGetAllVaultsByOwnerRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: QueryGetAllVaultsByOwnerRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetAllVaultsByOwnerRequest>
  ): QueryGetAllVaultsByOwnerRequest {
    const message = {
      ...baseQueryGetAllVaultsByOwnerRequest,
    } as QueryGetAllVaultsByOwnerRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseQueryGetAllVaultsByOwnerResponse: object = {};

export const QueryGetAllVaultsByOwnerResponse = {
  encode(
    message: QueryGetAllVaultsByOwnerResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.vaults) {
      Vault.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetAllVaultsByOwnerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetAllVaultsByOwnerResponse,
    } as QueryGetAllVaultsByOwnerResponse;
    message.vaults = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vaults.push(Vault.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetAllVaultsByOwnerResponse {
    const message = {
      ...baseQueryGetAllVaultsByOwnerResponse,
    } as QueryGetAllVaultsByOwnerResponse;
    message.vaults = [];
    if (object.vaults !== undefined && object.vaults !== null) {
      for (const e of object.vaults) {
        message.vaults.push(Vault.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetAllVaultsByOwnerResponse): unknown {
    const obj: any = {};
    if (message.vaults) {
      obj.vaults = message.vaults.map((e) => (e ? Vault.toJSON(e) : undefined));
    } else {
      obj.vaults = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetAllVaultsByOwnerResponse>
  ): QueryGetAllVaultsByOwnerResponse {
    const message = {
      ...baseQueryGetAllVaultsByOwnerResponse,
    } as QueryGetAllVaultsByOwnerResponse;
    message.vaults = [];
    if (object.vaults !== undefined && object.vaults !== null) {
      for (const e of object.vaults) {
        message.vaults.push(Vault.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetAllVaultsInDefaultRequest: object = {};

export const QueryGetAllVaultsInDefaultRequest = {
  encode(
    message: QueryGetAllVaultsInDefaultRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetAllVaultsInDefaultRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetAllVaultsInDefaultRequest,
    } as QueryGetAllVaultsInDefaultRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetAllVaultsInDefaultRequest {
    const message = {
      ...baseQueryGetAllVaultsInDefaultRequest,
    } as QueryGetAllVaultsInDefaultRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetAllVaultsInDefaultRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetAllVaultsInDefaultRequest>
  ): QueryGetAllVaultsInDefaultRequest {
    const message = {
      ...baseQueryGetAllVaultsInDefaultRequest,
    } as QueryGetAllVaultsInDefaultRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetAllVaultsInDefaultResponse: object = {};

export const QueryGetAllVaultsInDefaultResponse = {
  encode(
    message: QueryGetAllVaultsInDefaultResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.vaults) {
      Vault.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetAllVaultsInDefaultResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetAllVaultsInDefaultResponse,
    } as QueryGetAllVaultsInDefaultResponse;
    message.vaults = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vaults.push(Vault.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetAllVaultsInDefaultResponse {
    const message = {
      ...baseQueryGetAllVaultsInDefaultResponse,
    } as QueryGetAllVaultsInDefaultResponse;
    message.vaults = [];
    if (object.vaults !== undefined && object.vaults !== null) {
      for (const e of object.vaults) {
        message.vaults.push(Vault.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetAllVaultsInDefaultResponse): unknown {
    const obj: any = {};
    if (message.vaults) {
      obj.vaults = message.vaults.map((e) => (e ? Vault.toJSON(e) : undefined));
    } else {
      obj.vaults = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetAllVaultsInDefaultResponse>
  ): QueryGetAllVaultsInDefaultResponse {
    const message = {
      ...baseQueryGetAllVaultsInDefaultResponse,
    } as QueryGetAllVaultsInDefaultResponse;
    message.vaults = [];
    if (object.vaults !== undefined && object.vaults !== null) {
      for (const e of object.vaults) {
        message.vaults.push(Vault.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetVaultByUidRequest: object = { uid: 0 };

export const QueryGetVaultByUidRequest = {
  encode(
    message: QueryGetVaultByUidRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.uid !== 0) {
      writer.uint32(8).uint64(message.uid);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetVaultByUidRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetVaultByUidRequest,
    } as QueryGetVaultByUidRequest;
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

  fromJSON(object: any): QueryGetVaultByUidRequest {
    const message = {
      ...baseQueryGetVaultByUidRequest,
    } as QueryGetVaultByUidRequest;
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = Number(object.uid);
    } else {
      message.uid = 0;
    }
    return message;
  },

  toJSON(message: QueryGetVaultByUidRequest): unknown {
    const obj: any = {};
    message.uid !== undefined && (obj.uid = message.uid);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetVaultByUidRequest>
  ): QueryGetVaultByUidRequest {
    const message = {
      ...baseQueryGetVaultByUidRequest,
    } as QueryGetVaultByUidRequest;
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = object.uid;
    } else {
      message.uid = 0;
    }
    return message;
  },
};

const baseQueryGetVaultByUidResponse: object = {};

export const QueryGetVaultByUidResponse = {
  encode(
    message: QueryGetVaultByUidResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.vault !== undefined) {
      Vault.encode(message.vault, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetVaultByUidResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetVaultByUidResponse,
    } as QueryGetVaultByUidResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vault = Vault.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetVaultByUidResponse {
    const message = {
      ...baseQueryGetVaultByUidResponse,
    } as QueryGetVaultByUidResponse;
    if (object.vault !== undefined && object.vault !== null) {
      message.vault = Vault.fromJSON(object.vault);
    } else {
      message.vault = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetVaultByUidResponse): unknown {
    const obj: any = {};
    message.vault !== undefined &&
      (obj.vault = message.vault ? Vault.toJSON(message.vault) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetVaultByUidResponse>
  ): QueryGetVaultByUidResponse {
    const message = {
      ...baseQueryGetVaultByUidResponse,
    } as QueryGetVaultByUidResponse;
    if (object.vault !== undefined && object.vault !== null) {
      message.vault = Vault.fromPartial(object.vault);
    } else {
      message.vault = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of GetAllVaults items. */
  GetAllVaults(
    request: QueryGetAllVaultsRequest
  ): Promise<QueryGetAllVaultsResponse>;
  /** Queries a list of GetAllVaultsByOwner items. */
  GetAllVaultsByOwner(
    request: QueryGetAllVaultsByOwnerRequest
  ): Promise<QueryGetAllVaultsByOwnerResponse>;
  /** Queries a list of GetAllVaultsInDefault items. */
  GetAllVaultsInDefault(
    request: QueryGetAllVaultsInDefaultRequest
  ): Promise<QueryGetAllVaultsInDefaultResponse>;
  /** Queries a list of GetVaultByUid items. */
  GetVaultByUid(
    request: QueryGetVaultByUidRequest
  ): Promise<QueryGetVaultByUidResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("reserve.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  GetAllVaults(
    request: QueryGetAllVaultsRequest
  ): Promise<QueryGetAllVaultsResponse> {
    const data = QueryGetAllVaultsRequest.encode(request).finish();
    const promise = this.rpc.request("reserve.Query", "GetAllVaults", data);
    return promise.then((data) =>
      QueryGetAllVaultsResponse.decode(new Reader(data))
    );
  }

  GetAllVaultsByOwner(
    request: QueryGetAllVaultsByOwnerRequest
  ): Promise<QueryGetAllVaultsByOwnerResponse> {
    const data = QueryGetAllVaultsByOwnerRequest.encode(request).finish();
    const promise = this.rpc.request(
      "reserve.Query",
      "GetAllVaultsByOwner",
      data
    );
    return promise.then((data) =>
      QueryGetAllVaultsByOwnerResponse.decode(new Reader(data))
    );
  }

  GetAllVaultsInDefault(
    request: QueryGetAllVaultsInDefaultRequest
  ): Promise<QueryGetAllVaultsInDefaultResponse> {
    const data = QueryGetAllVaultsInDefaultRequest.encode(request).finish();
    const promise = this.rpc.request(
      "reserve.Query",
      "GetAllVaultsInDefault",
      data
    );
    return promise.then((data) =>
      QueryGetAllVaultsInDefaultResponse.decode(new Reader(data))
    );
  }

  GetVaultByUid(
    request: QueryGetVaultByUidRequest
  ): Promise<QueryGetVaultByUidResponse> {
    const data = QueryGetVaultByUidRequest.encode(request).finish();
    const promise = this.rpc.request("reserve.Query", "GetVaultByUid", data);
    return promise.then((data) =>
      QueryGetVaultByUidResponse.decode(new Reader(data))
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
