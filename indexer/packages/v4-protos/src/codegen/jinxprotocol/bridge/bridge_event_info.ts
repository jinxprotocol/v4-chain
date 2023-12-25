import * as _m0 from "protobufjs/minimal";
import { Long, DeepPartial } from "../../helpers";
/**
 * BridgeEventInfo stores information about the most recently processed bridge
 * event.
 */

export interface BridgeEventInfo {
  /**
   * The next event id (the last processed id plus one) of the logs from the
   * Ethereum contract.
   */
  nextId: number;
  /** The Ethereum block height of the most recently processed bridge event. */

  ethBlockHeight: Long;
}
/**
 * BridgeEventInfo stores information about the most recently processed bridge
 * event.
 */

export interface BridgeEventInfoSDKType {
  /**
   * The next event id (the last processed id plus one) of the logs from the
   * Ethereum contract.
   */
  next_id: number;
  /** The Ethereum block height of the most recently processed bridge event. */

  eth_block_height: Long;
}

function createBaseBridgeEventInfo(): BridgeEventInfo {
  return {
    nextId: 0,
    ethBlockHeight: Long.UZERO
  };
}

export const BridgeEventInfo = {
  encode(message: BridgeEventInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.nextId !== 0) {
      writer.uint32(8).uint32(message.nextId);
    }

    if (!message.ethBlockHeight.isZero()) {
      writer.uint32(16).uint64(message.ethBlockHeight);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BridgeEventInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBridgeEventInfo();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.nextId = reader.uint32();
          break;

        case 2:
          message.ethBlockHeight = (reader.uint64() as Long);
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<BridgeEventInfo>): BridgeEventInfo {
    const message = createBaseBridgeEventInfo();
    message.nextId = object.nextId ?? 0;
    message.ethBlockHeight = object.ethBlockHeight !== undefined && object.ethBlockHeight !== null ? Long.fromValue(object.ethBlockHeight) : Long.UZERO;
    return message;
  }

};