import * as _m0 from "protobufjs/minimal";
import { DeepPartial } from "../../helpers";
/** EpochInfo stores metadata of an epoch timer. */

export interface EpochInfo {
  /** name is the unique identifier. */
  name: string;
  /**
   * next_tick indicates when the next epoch starts (in Unix Epoch seconds),
   * if `EpochInfo` has been initialized.
   * If `EpochInfo` is not initialized yet, `next_tick` indicates the earliest
   * initialization time (see `is_initialized` below).
   */

  nextTick: number;
  /** duration of the epoch in seconds. */

  duration: number;
  /**
   * current epoch is the number of the current epoch.
   * 0 if `next_tick` has never been reached, positive otherwise.
   */

  currentEpoch: number;
  /**
   * current_epoch_start_block indicates the block height when the current
   * epoch started. 0 if `current_epoch` is 0.
   */

  currentEpochStartBlock: number;
  /**
   * is_initialized indicates whether the `EpochInfo` has been initialized
   * and started ticking.
   * An `EpochInfo` is initialized when all below conditions are true:
   * - Not yet initialized
   * - `BlockHeight` >= 2
   * - `BlockTime` >= `next_tick`
   */

  isInitialized: boolean;
  /**
   * fast_forward_next_tick specifies whether during initialization, `next_tick`
   * should be fast-forwarded to be greater than the current block time.
   * If `false`, the original `next_tick` value is
   * unchanged during initialization.
   * If `true`, `next_tick` will be set to the smallest value `x` greater than
   * the current block time such that `(x - next_tick) % duration = 0`.
   */

  fastForwardNextTick: boolean;
}
/** EpochInfo stores metadata of an epoch timer. */

export interface EpochInfoSDKType {
  /** name is the unique identifier. */
  name: string;
  /**
   * next_tick indicates when the next epoch starts (in Unix Epoch seconds),
   * if `EpochInfo` has been initialized.
   * If `EpochInfo` is not initialized yet, `next_tick` indicates the earliest
   * initialization time (see `is_initialized` below).
   */

  next_tick: number;
  /** duration of the epoch in seconds. */

  duration: number;
  /**
   * current epoch is the number of the current epoch.
   * 0 if `next_tick` has never been reached, positive otherwise.
   */

  current_epoch: number;
  /**
   * current_epoch_start_block indicates the block height when the current
   * epoch started. 0 if `current_epoch` is 0.
   */

  current_epoch_start_block: number;
  /**
   * is_initialized indicates whether the `EpochInfo` has been initialized
   * and started ticking.
   * An `EpochInfo` is initialized when all below conditions are true:
   * - Not yet initialized
   * - `BlockHeight` >= 2
   * - `BlockTime` >= `next_tick`
   */

  is_initialized: boolean;
  /**
   * fast_forward_next_tick specifies whether during initialization, `next_tick`
   * should be fast-forwarded to be greater than the current block time.
   * If `false`, the original `next_tick` value is
   * unchanged during initialization.
   * If `true`, `next_tick` will be set to the smallest value `x` greater than
   * the current block time such that `(x - next_tick) % duration = 0`.
   */

  fast_forward_next_tick: boolean;
}

function createBaseEpochInfo(): EpochInfo {
  return {
    name: "",
    nextTick: 0,
    duration: 0,
    currentEpoch: 0,
    currentEpochStartBlock: 0,
    isInitialized: false,
    fastForwardNextTick: false
  };
}

export const EpochInfo = {
  encode(message: EpochInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }

    if (message.nextTick !== 0) {
      writer.uint32(16).uint32(message.nextTick);
    }

    if (message.duration !== 0) {
      writer.uint32(24).uint32(message.duration);
    }

    if (message.currentEpoch !== 0) {
      writer.uint32(32).uint32(message.currentEpoch);
    }

    if (message.currentEpochStartBlock !== 0) {
      writer.uint32(40).uint32(message.currentEpochStartBlock);
    }

    if (message.isInitialized === true) {
      writer.uint32(48).bool(message.isInitialized);
    }

    if (message.fastForwardNextTick === true) {
      writer.uint32(56).bool(message.fastForwardNextTick);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EpochInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEpochInfo();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;

        case 2:
          message.nextTick = reader.uint32();
          break;

        case 3:
          message.duration = reader.uint32();
          break;

        case 4:
          message.currentEpoch = reader.uint32();
          break;

        case 5:
          message.currentEpochStartBlock = reader.uint32();
          break;

        case 6:
          message.isInitialized = reader.bool();
          break;

        case 7:
          message.fastForwardNextTick = reader.bool();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<EpochInfo>): EpochInfo {
    const message = createBaseEpochInfo();
    message.name = object.name ?? "";
    message.nextTick = object.nextTick ?? 0;
    message.duration = object.duration ?? 0;
    message.currentEpoch = object.currentEpoch ?? 0;
    message.currentEpochStartBlock = object.currentEpochStartBlock ?? 0;
    message.isInitialized = object.isInitialized ?? false;
    message.fastForwardNextTick = object.fastForwardNextTick ?? false;
    return message;
  }

};