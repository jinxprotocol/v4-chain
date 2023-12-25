import { Params, ParamsSDKType } from "./params";
import { StatsMetadata, StatsMetadataSDKType, GlobalStats, GlobalStatsSDKType, UserStats, UserStatsSDKType } from "./stats";
import * as _m0 from "protobufjs/minimal";
import { DeepPartial } from "../../helpers";
/** QueryParamsRequest is a request type for the Params RPC method. */

export interface QueryParamsRequest {}
/** QueryParamsRequest is a request type for the Params RPC method. */

export interface QueryParamsRequestSDKType {}
/** QueryParamsResponse is a response type for the Params RPC method. */

export interface QueryParamsResponse {
  params?: Params;
}
/** QueryParamsResponse is a response type for the Params RPC method. */

export interface QueryParamsResponseSDKType {
  params?: ParamsSDKType;
}
/** QueryStatsMetadataRequest is a request type for the StatsMetadata RPC method. */

export interface QueryStatsMetadataRequest {}
/** QueryStatsMetadataRequest is a request type for the StatsMetadata RPC method. */

export interface QueryStatsMetadataRequestSDKType {}
/**
 * QueryStatsMetadataResponse is a response type for the StatsMetadata RPC
 * method.
 */

export interface QueryStatsMetadataResponse {
  /**
   * QueryStatsMetadataResponse is a response type for the StatsMetadata RPC
   * method.
   */
  metadata?: StatsMetadata;
}
/**
 * QueryStatsMetadataResponse is a response type for the StatsMetadata RPC
 * method.
 */

export interface QueryStatsMetadataResponseSDKType {
  /**
   * QueryStatsMetadataResponse is a response type for the StatsMetadata RPC
   * method.
   */
  metadata?: StatsMetadataSDKType;
}
/** QueryGlobalStatsRequest is a request type for the GlobalStats RPC method. */

export interface QueryGlobalStatsRequest {}
/** QueryGlobalStatsRequest is a request type for the GlobalStats RPC method. */

export interface QueryGlobalStatsRequestSDKType {}
/** QueryGlobalStatsResponse is a response type for the GlobalStats RPC method. */

export interface QueryGlobalStatsResponse {
  /** QueryGlobalStatsResponse is a response type for the GlobalStats RPC method. */
  stats?: GlobalStats;
}
/** QueryGlobalStatsResponse is a response type for the GlobalStats RPC method. */

export interface QueryGlobalStatsResponseSDKType {
  /** QueryGlobalStatsResponse is a response type for the GlobalStats RPC method. */
  stats?: GlobalStatsSDKType;
}
/** QueryUserStatsRequest is a request type for the UserStats RPC method. */

export interface QueryUserStatsRequest {
  /** QueryUserStatsRequest is a request type for the UserStats RPC method. */
  user: string;
}
/** QueryUserStatsRequest is a request type for the UserStats RPC method. */

export interface QueryUserStatsRequestSDKType {
  /** QueryUserStatsRequest is a request type for the UserStats RPC method. */
  user: string;
}
/** QueryUserStatsResponse is a request type for the UserStats RPC method. */

export interface QueryUserStatsResponse {
  /** QueryUserStatsResponse is a request type for the UserStats RPC method. */
  stats?: UserStats;
}
/** QueryUserStatsResponse is a request type for the UserStats RPC method. */

export interface QueryUserStatsResponseSDKType {
  /** QueryUserStatsResponse is a request type for the UserStats RPC method. */
  stats?: UserStatsSDKType;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();

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

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  }

};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return {
    params: undefined
  };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();

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

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = object.params !== undefined && object.params !== null ? Params.fromPartial(object.params) : undefined;
    return message;
  }

};

function createBaseQueryStatsMetadataRequest(): QueryStatsMetadataRequest {
  return {};
}

export const QueryStatsMetadataRequest = {
  encode(_: QueryStatsMetadataRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryStatsMetadataRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryStatsMetadataRequest();

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

  fromPartial(_: DeepPartial<QueryStatsMetadataRequest>): QueryStatsMetadataRequest {
    const message = createBaseQueryStatsMetadataRequest();
    return message;
  }

};

function createBaseQueryStatsMetadataResponse(): QueryStatsMetadataResponse {
  return {
    metadata: undefined
  };
}

export const QueryStatsMetadataResponse = {
  encode(message: QueryStatsMetadataResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.metadata !== undefined) {
      StatsMetadata.encode(message.metadata, writer.uint32(10).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryStatsMetadataResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryStatsMetadataResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.metadata = StatsMetadata.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryStatsMetadataResponse>): QueryStatsMetadataResponse {
    const message = createBaseQueryStatsMetadataResponse();
    message.metadata = object.metadata !== undefined && object.metadata !== null ? StatsMetadata.fromPartial(object.metadata) : undefined;
    return message;
  }

};

function createBaseQueryGlobalStatsRequest(): QueryGlobalStatsRequest {
  return {};
}

export const QueryGlobalStatsRequest = {
  encode(_: QueryGlobalStatsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGlobalStatsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGlobalStatsRequest();

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

  fromPartial(_: DeepPartial<QueryGlobalStatsRequest>): QueryGlobalStatsRequest {
    const message = createBaseQueryGlobalStatsRequest();
    return message;
  }

};

function createBaseQueryGlobalStatsResponse(): QueryGlobalStatsResponse {
  return {
    stats: undefined
  };
}

export const QueryGlobalStatsResponse = {
  encode(message: QueryGlobalStatsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.stats !== undefined) {
      GlobalStats.encode(message.stats, writer.uint32(10).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGlobalStatsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGlobalStatsResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.stats = GlobalStats.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryGlobalStatsResponse>): QueryGlobalStatsResponse {
    const message = createBaseQueryGlobalStatsResponse();
    message.stats = object.stats !== undefined && object.stats !== null ? GlobalStats.fromPartial(object.stats) : undefined;
    return message;
  }

};

function createBaseQueryUserStatsRequest(): QueryUserStatsRequest {
  return {
    user: ""
  };
}

export const QueryUserStatsRequest = {
  encode(message: QueryUserStatsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.user !== "") {
      writer.uint32(10).string(message.user);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryUserStatsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryUserStatsRequest();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.user = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryUserStatsRequest>): QueryUserStatsRequest {
    const message = createBaseQueryUserStatsRequest();
    message.user = object.user ?? "";
    return message;
  }

};

function createBaseQueryUserStatsResponse(): QueryUserStatsResponse {
  return {
    stats: undefined
  };
}

export const QueryUserStatsResponse = {
  encode(message: QueryUserStatsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.stats !== undefined) {
      UserStats.encode(message.stats, writer.uint32(10).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryUserStatsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryUserStatsResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.stats = UserStats.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryUserStatsResponse>): QueryUserStatsResponse {
    const message = createBaseQueryUserStatsResponse();
    message.stats = object.stats !== undefined && object.stats !== null ? UserStats.fromPartial(object.stats) : undefined;
    return message;
  }

};