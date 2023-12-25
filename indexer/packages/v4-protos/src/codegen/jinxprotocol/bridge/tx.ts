import { BridgeEvent, BridgeEventSDKType } from "./bridge_event";
import { EventParams, EventParamsSDKType, ProposeParams, ProposeParamsSDKType, SafetyParams, SafetyParamsSDKType } from "./params";
import * as _m0 from "protobufjs/minimal";
import { DeepPartial } from "../../helpers";
/** MsgAcknowledgeBridges is the Msg/AcknowledgeBridges request type. */

export interface MsgAcknowledgeBridges {
  /** The events to acknowledge. */
  events: BridgeEvent[];
}
/** MsgAcknowledgeBridges is the Msg/AcknowledgeBridges request type. */

export interface MsgAcknowledgeBridgesSDKType {
  /** The events to acknowledge. */
  events: BridgeEventSDKType[];
}
/**
 * MsgAcknowledgeBridgesResponse is the Msg/AcknowledgeBridgesResponse response
 * type.
 */

export interface MsgAcknowledgeBridgesResponse {}
/**
 * MsgAcknowledgeBridgesResponse is the Msg/AcknowledgeBridgesResponse response
 * type.
 */

export interface MsgAcknowledgeBridgesResponseSDKType {}
/** MsgCompleteBridge is the Msg/CompleteBridgeResponse request type. */

export interface MsgCompleteBridge {
  authority: string;
  /** The event to complete. */

  event?: BridgeEvent;
}
/** MsgCompleteBridge is the Msg/CompleteBridgeResponse request type. */

export interface MsgCompleteBridgeSDKType {
  authority: string;
  /** The event to complete. */

  event?: BridgeEventSDKType;
}
/** MsgCompleteBridgeResponse is the Msg/CompleteBridgeResponse response type. */

export interface MsgCompleteBridgeResponse {}
/** MsgCompleteBridgeResponse is the Msg/CompleteBridgeResponse response type. */

export interface MsgCompleteBridgeResponseSDKType {}
/** MsgUpdateEventParams is the Msg/UpdateEventParams request type. */

export interface MsgUpdateEventParams {
  authority: string;
  /** The parameters to update. Each field must be set. */

  params?: EventParams;
}
/** MsgUpdateEventParams is the Msg/UpdateEventParams request type. */

export interface MsgUpdateEventParamsSDKType {
  authority: string;
  /** The parameters to update. Each field must be set. */

  params?: EventParamsSDKType;
}
/** MsgUpdateEventParamsResponse is the Msg/UpdateEventParams response type. */

export interface MsgUpdateEventParamsResponse {}
/** MsgUpdateEventParamsResponse is the Msg/UpdateEventParams response type. */

export interface MsgUpdateEventParamsResponseSDKType {}
/** MsgUpdateProposeParams is the Msg/UpdateProposeParams request type. */

export interface MsgUpdateProposeParams {
  authority: string;
  /** The parameters to update. Each field must be set. */

  params?: ProposeParams;
}
/** MsgUpdateProposeParams is the Msg/UpdateProposeParams request type. */

export interface MsgUpdateProposeParamsSDKType {
  authority: string;
  /** The parameters to update. Each field must be set. */

  params?: ProposeParamsSDKType;
}
/** MsgUpdateProposeParamsResponse is the Msg/UpdateProposeParams response type. */

export interface MsgUpdateProposeParamsResponse {}
/** MsgUpdateProposeParamsResponse is the Msg/UpdateProposeParams response type. */

export interface MsgUpdateProposeParamsResponseSDKType {}
/** MsgUpdateSafetyParams is the Msg/UpdateSafetyParams request type. */

export interface MsgUpdateSafetyParams {
  authority: string;
  /** The parameters to update. Each field must be set. */

  params?: SafetyParams;
}
/** MsgUpdateSafetyParams is the Msg/UpdateSafetyParams request type. */

export interface MsgUpdateSafetyParamsSDKType {
  authority: string;
  /** The parameters to update. Each field must be set. */

  params?: SafetyParamsSDKType;
}
/** MsgUpdateSafetyParamsResponse is the Msg/UpdateSafetyParams response type. */

export interface MsgUpdateSafetyParamsResponse {}
/** MsgUpdateSafetyParamsResponse is the Msg/UpdateSafetyParams response type. */

export interface MsgUpdateSafetyParamsResponseSDKType {}

function createBaseMsgAcknowledgeBridges(): MsgAcknowledgeBridges {
  return {
    events: []
  };
}

export const MsgAcknowledgeBridges = {
  encode(message: MsgAcknowledgeBridges, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.events) {
      BridgeEvent.encode(v!, writer.uint32(10).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAcknowledgeBridges {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAcknowledgeBridges();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.events.push(BridgeEvent.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgAcknowledgeBridges>): MsgAcknowledgeBridges {
    const message = createBaseMsgAcknowledgeBridges();
    message.events = object.events?.map(e => BridgeEvent.fromPartial(e)) || [];
    return message;
  }

};

function createBaseMsgAcknowledgeBridgesResponse(): MsgAcknowledgeBridgesResponse {
  return {};
}

export const MsgAcknowledgeBridgesResponse = {
  encode(_: MsgAcknowledgeBridgesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAcknowledgeBridgesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAcknowledgeBridgesResponse();

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

  fromPartial(_: DeepPartial<MsgAcknowledgeBridgesResponse>): MsgAcknowledgeBridgesResponse {
    const message = createBaseMsgAcknowledgeBridgesResponse();
    return message;
  }

};

function createBaseMsgCompleteBridge(): MsgCompleteBridge {
  return {
    authority: "",
    event: undefined
  };
}

export const MsgCompleteBridge = {
  encode(message: MsgCompleteBridge, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.authority !== "") {
      writer.uint32(10).string(message.authority);
    }

    if (message.event !== undefined) {
      BridgeEvent.encode(message.event, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCompleteBridge {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCompleteBridge();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.authority = reader.string();
          break;

        case 2:
          message.event = BridgeEvent.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgCompleteBridge>): MsgCompleteBridge {
    const message = createBaseMsgCompleteBridge();
    message.authority = object.authority ?? "";
    message.event = object.event !== undefined && object.event !== null ? BridgeEvent.fromPartial(object.event) : undefined;
    return message;
  }

};

function createBaseMsgCompleteBridgeResponse(): MsgCompleteBridgeResponse {
  return {};
}

export const MsgCompleteBridgeResponse = {
  encode(_: MsgCompleteBridgeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCompleteBridgeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCompleteBridgeResponse();

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

  fromPartial(_: DeepPartial<MsgCompleteBridgeResponse>): MsgCompleteBridgeResponse {
    const message = createBaseMsgCompleteBridgeResponse();
    return message;
  }

};

function createBaseMsgUpdateEventParams(): MsgUpdateEventParams {
  return {
    authority: "",
    params: undefined
  };
}

export const MsgUpdateEventParams = {
  encode(message: MsgUpdateEventParams, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.authority !== "") {
      writer.uint32(10).string(message.authority);
    }

    if (message.params !== undefined) {
      EventParams.encode(message.params, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateEventParams {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateEventParams();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.authority = reader.string();
          break;

        case 2:
          message.params = EventParams.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgUpdateEventParams>): MsgUpdateEventParams {
    const message = createBaseMsgUpdateEventParams();
    message.authority = object.authority ?? "";
    message.params = object.params !== undefined && object.params !== null ? EventParams.fromPartial(object.params) : undefined;
    return message;
  }

};

function createBaseMsgUpdateEventParamsResponse(): MsgUpdateEventParamsResponse {
  return {};
}

export const MsgUpdateEventParamsResponse = {
  encode(_: MsgUpdateEventParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateEventParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateEventParamsResponse();

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

  fromPartial(_: DeepPartial<MsgUpdateEventParamsResponse>): MsgUpdateEventParamsResponse {
    const message = createBaseMsgUpdateEventParamsResponse();
    return message;
  }

};

function createBaseMsgUpdateProposeParams(): MsgUpdateProposeParams {
  return {
    authority: "",
    params: undefined
  };
}

export const MsgUpdateProposeParams = {
  encode(message: MsgUpdateProposeParams, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.authority !== "") {
      writer.uint32(10).string(message.authority);
    }

    if (message.params !== undefined) {
      ProposeParams.encode(message.params, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateProposeParams {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateProposeParams();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.authority = reader.string();
          break;

        case 2:
          message.params = ProposeParams.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgUpdateProposeParams>): MsgUpdateProposeParams {
    const message = createBaseMsgUpdateProposeParams();
    message.authority = object.authority ?? "";
    message.params = object.params !== undefined && object.params !== null ? ProposeParams.fromPartial(object.params) : undefined;
    return message;
  }

};

function createBaseMsgUpdateProposeParamsResponse(): MsgUpdateProposeParamsResponse {
  return {};
}

export const MsgUpdateProposeParamsResponse = {
  encode(_: MsgUpdateProposeParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateProposeParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateProposeParamsResponse();

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

  fromPartial(_: DeepPartial<MsgUpdateProposeParamsResponse>): MsgUpdateProposeParamsResponse {
    const message = createBaseMsgUpdateProposeParamsResponse();
    return message;
  }

};

function createBaseMsgUpdateSafetyParams(): MsgUpdateSafetyParams {
  return {
    authority: "",
    params: undefined
  };
}

export const MsgUpdateSafetyParams = {
  encode(message: MsgUpdateSafetyParams, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.authority !== "") {
      writer.uint32(10).string(message.authority);
    }

    if (message.params !== undefined) {
      SafetyParams.encode(message.params, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateSafetyParams {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateSafetyParams();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.authority = reader.string();
          break;

        case 2:
          message.params = SafetyParams.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgUpdateSafetyParams>): MsgUpdateSafetyParams {
    const message = createBaseMsgUpdateSafetyParams();
    message.authority = object.authority ?? "";
    message.params = object.params !== undefined && object.params !== null ? SafetyParams.fromPartial(object.params) : undefined;
    return message;
  }

};

function createBaseMsgUpdateSafetyParamsResponse(): MsgUpdateSafetyParamsResponse {
  return {};
}

export const MsgUpdateSafetyParamsResponse = {
  encode(_: MsgUpdateSafetyParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateSafetyParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateSafetyParamsResponse();

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

  fromPartial(_: DeepPartial<MsgUpdateSafetyParamsResponse>): MsgUpdateSafetyParamsResponse {
    const message = createBaseMsgUpdateSafetyParamsResponse();
    return message;
  }

};