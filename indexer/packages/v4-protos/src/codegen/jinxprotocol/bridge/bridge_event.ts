import { Coin, CoinSDKType } from "../../cosmos/base/v1beta1/coin";
import * as _m0 from "protobufjs/minimal";
import { Long, DeepPartial } from "../../helpers";
/** BridgeEvent is a recognized event from the Ethereum blockchain. */

export interface BridgeEvent {
  /** The unique id of the Ethereum event log. */
  id: number;
  /** The tokens bridged. */

  coin?: Coin;
  /** The account address or module address to bridge to. */

  address: string;
  /** The Ethereum block height of the event. */

  ethBlockHeight: Long;
}
/** BridgeEvent is a recognized event from the Ethereum blockchain. */

export interface BridgeEventSDKType {
  /** The unique id of the Ethereum event log. */
  id: number;
  /** The tokens bridged. */

  coin?: CoinSDKType;
  /** The account address or module address to bridge to. */

  address: string;
  /** The Ethereum block height of the event. */

  eth_block_height: Long;
}

function createBaseBridgeEvent(): BridgeEvent {
  return {
    id: 0,
    coin: undefined,
    address: "",
    ethBlockHeight: Long.UZERO
  };
}

export const BridgeEvent = {
  encode(message: BridgeEvent, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint32(message.id);
    }

    if (message.coin !== undefined) {
      Coin.encode(message.coin, writer.uint32(18).fork()).ldelim();
    }

    if (message.address !== "") {
      writer.uint32(26).string(message.address);
    }

    if (!message.ethBlockHeight.isZero()) {
      writer.uint32(32).uint64(message.ethBlockHeight);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BridgeEvent {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBridgeEvent();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.id = reader.uint32();
          break;

        case 2:
          message.coin = Coin.decode(reader, reader.uint32());
          break;

        case 3:
          message.address = reader.string();
          break;

        case 4:
          message.ethBlockHeight = (reader.uint64() as Long);
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<BridgeEvent>): BridgeEvent {
    const message = createBaseBridgeEvent();
    message.id = object.id ?? 0;
    message.coin = object.coin !== undefined && object.coin !== null ? Coin.fromPartial(object.coin) : undefined;
    message.address = object.address ?? "";
    message.ethBlockHeight = object.ethBlockHeight !== undefined && object.ethBlockHeight !== null ? Long.fromValue(object.ethBlockHeight) : Long.UZERO;
    return message;
  }

};