import * as _m0 from "protobufjs/minimal";
import { DeepPartial } from "../../helpers";
/** Params defines the parameters for x/perpetuals module. */

export interface Params {
  /**
   * Funding rate clamp factor in parts-per-million, used for clamping 8-hour
   * funding rates according to equation: |R| <= funding_rate_clamp_factor *
   * (initial margin - maintenance margin).
   */
  fundingRateClampFactorPpm: number;
  /**
   * Premium vote clamp factor in parts-per-million, used for clamping premium
   * votes according to equation: |V| <= premium_vote_clamp_factor *
   * (initial margin - maintenance margin).
   */

  premiumVoteClampFactorPpm: number;
  /**
   * Minimum number of premium votes per premium sample. If number of premium
   * votes is smaller than this number, pad with zeros up to this number.
   */

  minNumVotesPerSample: number;
}
/** Params defines the parameters for x/perpetuals module. */

export interface ParamsSDKType {
  /**
   * Funding rate clamp factor in parts-per-million, used for clamping 8-hour
   * funding rates according to equation: |R| <= funding_rate_clamp_factor *
   * (initial margin - maintenance margin).
   */
  funding_rate_clamp_factor_ppm: number;
  /**
   * Premium vote clamp factor in parts-per-million, used for clamping premium
   * votes according to equation: |V| <= premium_vote_clamp_factor *
   * (initial margin - maintenance margin).
   */

  premium_vote_clamp_factor_ppm: number;
  /**
   * Minimum number of premium votes per premium sample. If number of premium
   * votes is smaller than this number, pad with zeros up to this number.
   */

  min_num_votes_per_sample: number;
}

function createBaseParams(): Params {
  return {
    fundingRateClampFactorPpm: 0,
    premiumVoteClampFactorPpm: 0,
    minNumVotesPerSample: 0
  };
}

export const Params = {
  encode(message: Params, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.fundingRateClampFactorPpm !== 0) {
      writer.uint32(8).uint32(message.fundingRateClampFactorPpm);
    }

    if (message.premiumVoteClampFactorPpm !== 0) {
      writer.uint32(16).uint32(message.premiumVoteClampFactorPpm);
    }

    if (message.minNumVotesPerSample !== 0) {
      writer.uint32(24).uint32(message.minNumVotesPerSample);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseParams();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.fundingRateClampFactorPpm = reader.uint32();
          break;

        case 2:
          message.premiumVoteClampFactorPpm = reader.uint32();
          break;

        case 3:
          message.minNumVotesPerSample = reader.uint32();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = createBaseParams();
    message.fundingRateClampFactorPpm = object.fundingRateClampFactorPpm ?? 0;
    message.premiumVoteClampFactorPpm = object.premiumVoteClampFactorPpm ?? 0;
    message.minNumVotesPerSample = object.minNumVotesPerSample ?? 0;
    return message;
  }

};