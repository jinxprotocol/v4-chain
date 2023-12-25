import { EventParams, EventParamsSDKType, ProposeParams, ProposeParamsSDKType, SafetyParams, SafetyParamsSDKType } from "./params";
import { BridgeEventInfo, BridgeEventInfoSDKType } from "./bridge_event_info";
import * as _m0 from "protobufjs/minimal";
import { DeepPartial } from "../../helpers";
/** GenesisState defines the bridge module's genesis state. */

export interface GenesisState {
  /** The parameters of the module. */
  eventParams?: EventParams;
  proposeParams?: ProposeParams;
  safetyParams?: SafetyParams;
  /**
   * Acknowledged event info that stores:
   * - the next event ID to be added to consensus.
   * - Ethereum block height of the most recently acknowledged bridge event.
   */

  acknowledgedEventInfo?: BridgeEventInfo;
}
/** GenesisState defines the bridge module's genesis state. */

export interface GenesisStateSDKType {
  /** The parameters of the module. */
  event_params?: EventParamsSDKType;
  propose_params?: ProposeParamsSDKType;
  safety_params?: SafetyParamsSDKType;
  /**
   * Acknowledged event info that stores:
   * - the next event ID to be added to consensus.
   * - Ethereum block height of the most recently acknowledged bridge event.
   */

  acknowledged_event_info?: BridgeEventInfoSDKType;
}

function createBaseGenesisState(): GenesisState {
  return {
    eventParams: undefined,
    proposeParams: undefined,
    safetyParams: undefined,
    acknowledgedEventInfo: undefined
  };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.eventParams !== undefined) {
      EventParams.encode(message.eventParams, writer.uint32(10).fork()).ldelim();
    }

    if (message.proposeParams !== undefined) {
      ProposeParams.encode(message.proposeParams, writer.uint32(18).fork()).ldelim();
    }

    if (message.safetyParams !== undefined) {
      SafetyParams.encode(message.safetyParams, writer.uint32(26).fork()).ldelim();
    }

    if (message.acknowledgedEventInfo !== undefined) {
      BridgeEventInfo.encode(message.acknowledgedEventInfo, writer.uint32(34).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.eventParams = EventParams.decode(reader, reader.uint32());
          break;

        case 2:
          message.proposeParams = ProposeParams.decode(reader, reader.uint32());
          break;

        case 3:
          message.safetyParams = SafetyParams.decode(reader, reader.uint32());
          break;

        case 4:
          message.acknowledgedEventInfo = BridgeEventInfo.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = createBaseGenesisState();
    message.eventParams = object.eventParams !== undefined && object.eventParams !== null ? EventParams.fromPartial(object.eventParams) : undefined;
    message.proposeParams = object.proposeParams !== undefined && object.proposeParams !== null ? ProposeParams.fromPartial(object.proposeParams) : undefined;
    message.safetyParams = object.safetyParams !== undefined && object.safetyParams !== null ? SafetyParams.fromPartial(object.safetyParams) : undefined;
    message.acknowledgedEventInfo = object.acknowledgedEventInfo !== undefined && object.acknowledgedEventInfo !== null ? BridgeEventInfo.fromPartial(object.acknowledgedEventInfo) : undefined;
    return message;
  }

};