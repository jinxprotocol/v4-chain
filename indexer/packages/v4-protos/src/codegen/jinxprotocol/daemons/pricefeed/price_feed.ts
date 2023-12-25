import { Timestamp } from "../../../google/protobuf/timestamp";
import * as _m0 from "protobufjs/minimal";
import { DeepPartial, toTimestamp, Long, fromTimestamp } from "../../../helpers";
/** UpdateMarketPriceRequest is a request message updating market prices. */

export interface UpdateMarketPricesRequest {
  marketPriceUpdates: MarketPriceUpdate[];
}
/** UpdateMarketPriceRequest is a request message updating market prices. */

export interface UpdateMarketPricesRequestSDKType {
  market_price_updates: MarketPriceUpdateSDKType[];
}
/** UpdateMarketPricesResponse is a response message for updating market prices. */

export interface UpdateMarketPricesResponse {}
/** UpdateMarketPricesResponse is a response message for updating market prices. */

export interface UpdateMarketPricesResponseSDKType {}
/** ExchangePrice represents a specific exchange's market price */

export interface ExchangePrice {
  exchangeId: string;
  price: Long;
  lastUpdateTime?: Date;
}
/** ExchangePrice represents a specific exchange's market price */

export interface ExchangePriceSDKType {
  exchange_id: string;
  price: Long;
  last_update_time?: Date;
}
/** MarketPriceUpdate represents an update to a single market */

export interface MarketPriceUpdate {
  marketId: number;
  exchangePrices: ExchangePrice[];
}
/** MarketPriceUpdate represents an update to a single market */

export interface MarketPriceUpdateSDKType {
  market_id: number;
  exchange_prices: ExchangePriceSDKType[];
}

function createBaseUpdateMarketPricesRequest(): UpdateMarketPricesRequest {
  return {
    marketPriceUpdates: []
  };
}

export const UpdateMarketPricesRequest = {
  encode(message: UpdateMarketPricesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.marketPriceUpdates) {
      MarketPriceUpdate.encode(v!, writer.uint32(10).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateMarketPricesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateMarketPricesRequest();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.marketPriceUpdates.push(MarketPriceUpdate.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<UpdateMarketPricesRequest>): UpdateMarketPricesRequest {
    const message = createBaseUpdateMarketPricesRequest();
    message.marketPriceUpdates = object.marketPriceUpdates?.map(e => MarketPriceUpdate.fromPartial(e)) || [];
    return message;
  }

};

function createBaseUpdateMarketPricesResponse(): UpdateMarketPricesResponse {
  return {};
}

export const UpdateMarketPricesResponse = {
  encode(_: UpdateMarketPricesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateMarketPricesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateMarketPricesResponse();

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

  fromPartial(_: DeepPartial<UpdateMarketPricesResponse>): UpdateMarketPricesResponse {
    const message = createBaseUpdateMarketPricesResponse();
    return message;
  }

};

function createBaseExchangePrice(): ExchangePrice {
  return {
    exchangeId: "",
    price: Long.UZERO,
    lastUpdateTime: undefined
  };
}

export const ExchangePrice = {
  encode(message: ExchangePrice, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.exchangeId !== "") {
      writer.uint32(10).string(message.exchangeId);
    }

    if (!message.price.isZero()) {
      writer.uint32(16).uint64(message.price);
    }

    if (message.lastUpdateTime !== undefined) {
      Timestamp.encode(toTimestamp(message.lastUpdateTime), writer.uint32(26).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ExchangePrice {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseExchangePrice();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.exchangeId = reader.string();
          break;

        case 2:
          message.price = (reader.uint64() as Long);
          break;

        case 3:
          message.lastUpdateTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<ExchangePrice>): ExchangePrice {
    const message = createBaseExchangePrice();
    message.exchangeId = object.exchangeId ?? "";
    message.price = object.price !== undefined && object.price !== null ? Long.fromValue(object.price) : Long.UZERO;
    message.lastUpdateTime = object.lastUpdateTime ?? undefined;
    return message;
  }

};

function createBaseMarketPriceUpdate(): MarketPriceUpdate {
  return {
    marketId: 0,
    exchangePrices: []
  };
}

export const MarketPriceUpdate = {
  encode(message: MarketPriceUpdate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.marketId !== 0) {
      writer.uint32(8).uint32(message.marketId);
    }

    for (const v of message.exchangePrices) {
      ExchangePrice.encode(v!, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MarketPriceUpdate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMarketPriceUpdate();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.marketId = reader.uint32();
          break;

        case 2:
          message.exchangePrices.push(ExchangePrice.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MarketPriceUpdate>): MarketPriceUpdate {
    const message = createBaseMarketPriceUpdate();
    message.marketId = object.marketId ?? 0;
    message.exchangePrices = object.exchangePrices?.map(e => ExchangePrice.fromPartial(e)) || [];
    return message;
  }

};