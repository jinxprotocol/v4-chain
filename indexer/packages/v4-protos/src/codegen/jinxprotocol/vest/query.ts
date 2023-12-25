import { VestEntry, VestEntrySDKType } from "./vest_entry";
import * as _m0 from "protobufjs/minimal";
import { DeepPartial } from "../../helpers";
/** QueryVestEntryRequest is a request type for the VestEntry RPC method. */

export interface QueryVestEntryRequest {
  /** QueryVestEntryRequest is a request type for the VestEntry RPC method. */
  vesterAccount: string;
}
/** QueryVestEntryRequest is a request type for the VestEntry RPC method. */

export interface QueryVestEntryRequestSDKType {
  /** QueryVestEntryRequest is a request type for the VestEntry RPC method. */
  vester_account: string;
}
/** QueryVestEntryResponse is a response type for the VestEntry RPC method. */

export interface QueryVestEntryResponse {
  entry?: VestEntry;
}
/** QueryVestEntryResponse is a response type for the VestEntry RPC method. */

export interface QueryVestEntryResponseSDKType {
  entry?: VestEntrySDKType;
}

function createBaseQueryVestEntryRequest(): QueryVestEntryRequest {
  return {
    vesterAccount: ""
  };
}

export const QueryVestEntryRequest = {
  encode(message: QueryVestEntryRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.vesterAccount !== "") {
      writer.uint32(10).string(message.vesterAccount);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryVestEntryRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryVestEntryRequest();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.vesterAccount = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryVestEntryRequest>): QueryVestEntryRequest {
    const message = createBaseQueryVestEntryRequest();
    message.vesterAccount = object.vesterAccount ?? "";
    return message;
  }

};

function createBaseQueryVestEntryResponse(): QueryVestEntryResponse {
  return {
    entry: undefined
  };
}

export const QueryVestEntryResponse = {
  encode(message: QueryVestEntryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.entry !== undefined) {
      VestEntry.encode(message.entry, writer.uint32(10).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryVestEntryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryVestEntryResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.entry = VestEntry.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryVestEntryResponse>): QueryVestEntryResponse {
    const message = createBaseQueryVestEntryResponse();
    message.entry = object.entry !== undefined && object.entry !== null ? VestEntry.fromPartial(object.entry) : undefined;
    return message;
  }

};