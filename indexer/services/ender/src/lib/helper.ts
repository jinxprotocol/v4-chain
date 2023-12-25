import { logger, ParseMessageError } from '@dydxprotocol-indexer/base';
import {
  OrderSide,
  PerpetualMarketFromDatabase,
  PositionSide,
  protocolTranslations,
} from '@dydxprotocol-indexer/postgres';
import {
  IndexerTendermintEvent,
  IndexerTendermintEvent_BlockEvent,
  Timestamp,
  OrderFillEventV1,
  MarketEventV1,
  SubaccountUpdateEventV1,
  TransferEventV1,
  IndexerOrder,
  StatefulOrderEventV1,
  FundingEventV1,
  AssetCreateEventV1,
  PerpetualMarketCreateEventV1,
  LiquidityTierUpsertEventV1,
  UpdatePerpetualEventV1,
  UpdateClobPairEventV1,
  SubaccountMessage,
  DeleveragingEventV1,
  TradingRewardsEventV1,
} from '@dydxprotocol-indexer/v4-protos';
import Big from 'big.js';
import _ from 'lodash';
import { DateTime } from 'luxon';

import {
  MILLIS_IN_NANOS,
  SECONDS_IN_MILLIS,
} from '../constants';
import {
  AnnotatedSubaccountMessage,
  JinxIndexerSubtypes,
  EventProtoWithTypeAndVersion,
} from './types';

export function indexerTendermintEventToTransactionIndex(
  event: IndexerTendermintEvent,
): number {
  if (event.transactionIndex !== undefined) {
    return event.transactionIndex;
  } else if (event.blockEvent !== undefined) {
    switch (event.blockEvent) {
      case IndexerTendermintEvent_BlockEvent.BLOCK_EVENT_BEGIN_BLOCK:
        return -2;
      case IndexerTendermintEvent_BlockEvent.BLOCK_EVENT_END_BLOCK:
        return -1;
      default:
        throw new ParseMessageError(`Received V4 event with invalid block event type: ${event.blockEvent}`);
    }
  }

  throw new ParseMessageError(
    'Either transactionIndex or blockEvent must be defined in IndexerTendermintEvent',
  );
}

export function convertToSubaccountMessage(
  annotatedMessage: AnnotatedSubaccountMessage,
): SubaccountMessage {
  const subaccountMessage: SubaccountMessage = _.omit(
    annotatedMessage,
    ['orderId', 'isFill', 'subaccountMessageContents'],
  );
  return subaccountMessage;
}

export function protoTimestampToDate(
  protoTime: Timestamp,
): Date {
  const timeInMillis: number = Number(protoTime.seconds) * SECONDS_IN_MILLIS +
    Math.floor(protoTime.nanos / MILLIS_IN_NANOS);

  return new Date(timeInMillis);
}

export function dateToDateTime(
  protoTime: Date,
): DateTime {
  return DateTime.fromJSDate(protoTime);
}

/**
 * Determines the event subtype and parses the IndexerTendermintEvent
 * to the correct EventProto and returns it all in an object.
 * @param blockEventIndex - the index of the event in the block.
 * @param event - the event.
 * @returns
 */
export function indexerTendermintEventToEventProtoWithType(
  blockEventIndex: number,
  event: IndexerTendermintEvent,
): EventProtoWithTypeAndVersion | undefined {
  const eventDataBinary: Uint8Array = event.dataBytes;
  // set the default version to 1
  const version: number = event.version === 0 ? 1 : event.version;
  switch (event.subtype) {
    case (JinxIndexerSubtypes.ORDER_FILL.toString()): {
      return {
        type: JinxIndexerSubtypes.ORDER_FILL,
        eventProto: OrderFillEventV1.decode(eventDataBinary),
        indexerTendermintEvent: event,
        version,
        blockEventIndex,
      };
    }
    case (JinxIndexerSubtypes.SUBACCOUNT_UPDATE.toString()): {
      return {
        type: JinxIndexerSubtypes.SUBACCOUNT_UPDATE,
        eventProto: SubaccountUpdateEventV1.decode(eventDataBinary),
        indexerTendermintEvent: event,
        version,
        blockEventIndex,
      };
    }
    case (JinxIndexerSubtypes.TRANSFER.toString()): {
      return {
        type: JinxIndexerSubtypes.TRANSFER,
        eventProto: TransferEventV1.decode(eventDataBinary),
        indexerTendermintEvent: event,
        version,
        blockEventIndex,
      };
    }
    case (JinxIndexerSubtypes.MARKET.toString()): {
      return {
        type: JinxIndexerSubtypes.MARKET,
        eventProto: MarketEventV1.decode(eventDataBinary),
        indexerTendermintEvent: event,
        version,
        blockEventIndex,
      };
    }
    case (JinxIndexerSubtypes.STATEFUL_ORDER.toString()): {
      return {
        type: JinxIndexerSubtypes.STATEFUL_ORDER,
        eventProto: StatefulOrderEventV1.decode(eventDataBinary),
        indexerTendermintEvent: event,
        version,
        blockEventIndex,
      };
    }
    case (JinxIndexerSubtypes.FUNDING.toString()): {
      return {
        type: JinxIndexerSubtypes.FUNDING,
        eventProto: FundingEventV1.decode(eventDataBinary),
        indexerTendermintEvent: event,
        version,
        blockEventIndex,
      };
    }
    case (JinxIndexerSubtypes.ASSET.toString()): {
      return {
        type: JinxIndexerSubtypes.ASSET,
        eventProto: AssetCreateEventV1.decode(eventDataBinary),
        indexerTendermintEvent: event,
        version,
        blockEventIndex,
      };
    }
    case (JinxIndexerSubtypes.PERPETUAL_MARKET.toString()): {
      return {
        type: JinxIndexerSubtypes.PERPETUAL_MARKET,
        eventProto: PerpetualMarketCreateEventV1.decode(eventDataBinary),
        indexerTendermintEvent: event,
        version,
        blockEventIndex,
      };
    }
    case (JinxIndexerSubtypes.LIQUIDITY_TIER.toString()): {
      return {
        type: JinxIndexerSubtypes.LIQUIDITY_TIER,
        eventProto: LiquidityTierUpsertEventV1.decode(eventDataBinary),
        indexerTendermintEvent: event,
        version,
        blockEventIndex,
      };
    }
    case (JinxIndexerSubtypes.UPDATE_PERPETUAL.toString()): {
      return {
        type: JinxIndexerSubtypes.UPDATE_PERPETUAL,
        eventProto: UpdatePerpetualEventV1.decode(eventDataBinary),
        indexerTendermintEvent: event,
        version,
        blockEventIndex,
      };
    }
    case (JinxIndexerSubtypes.UPDATE_CLOB_PAIR.toString()): {
      return {
        type: JinxIndexerSubtypes.UPDATE_CLOB_PAIR,
        eventProto: UpdateClobPairEventV1.decode(eventDataBinary),
        indexerTendermintEvent: event,
        version,
        blockEventIndex,
      };
    }
    case (JinxIndexerSubtypes.DELEVERAGING.toString()): {
      return {
        type: JinxIndexerSubtypes.DELEVERAGING,
        eventProto: DeleveragingEventV1.decode(eventDataBinary),
        indexerTendermintEvent: event,
        version,
        blockEventIndex,
      };
    }
    case (JinxIndexerSubtypes.TRADING_REWARD.toString()): {
      return {
        type: JinxIndexerSubtypes.TRADING_REWARD,
        eventProto: TradingRewardsEventV1.decode(eventDataBinary),
        indexerTendermintEvent: event,
        version,
        blockEventIndex,
      };
    }
    default: {
      const message: string = `Unable to parse event subtype: ${event.subtype}`;
      logger.error({
        at: 'helpers#indexerTendermintEventToEventWithType',
        message,
      });
      return undefined;
    }
  }
}

/**
 * Returns the size of an order in human readable form.
 * @param order
 * @param perpetualMarket
 */
export function getSize(
  order: IndexerOrder,
  perpetualMarket: PerpetualMarketFromDatabase,
): string {
  return protocolTranslations.quantumsToHumanFixedString(
    order.quantums.toString(10),
    perpetualMarket.atomicResolution,
  );
}

/**
 * Returns the price of an order in human readable form.
 *
 * @param order
 * @param perpetualMarket
 */
export function getPrice(
  order: IndexerOrder,
  perpetualMarket: PerpetualMarketFromDatabase,
): string {
  return protocolTranslations.subticksToPrice(
    order.subticks.toString(10),
    perpetualMarket,
  );
}

/**
 * Returns the trigger price of an order in human readable form.
 *
 * @param order
 * @param perpetualMarket
 * @returns
 */
export function getTriggerPrice(
  order: IndexerOrder,
  perpetualMarket: PerpetualMarketFromDatabase,
): string {
  return protocolTranslations.subticksToPrice(
    order.conditionalOrderTriggerSubticks.toString(10),
    perpetualMarket,
  );
}

/**
 * Returns the weighted average between two prices
 * @param firstPrice
 * @param firstWeight
 * @param secondPrice
 * @param secondWeight
 * @returns
 */
export function getWeightedAverage(
  firstPrice: string,
  firstWeight: string,
  secondPrice: string,
  secondWeight: string,
): string {
  return Big(firstPrice).times(firstWeight).plus(
    Big(secondPrice).times(secondWeight),
  )
    .div(Big(firstWeight).plus(secondWeight))
    .toFixed();
}

/**
 * Returns true if perpetualPostionSide is LONG and orderSide is BUY or
 * if perpetualPostionSide is SHORT and orderSide is SELL
 * @param perpetualPositionSide
 * @param orderSide
 */
export function perpetualPositionAndOrderSideMatching(
  perpetualPositionSide: PositionSide,
  orderSide: OrderSide,
): boolean {
  return (perpetualPositionSide === PositionSide.LONG && orderSide === OrderSide.BUY) ||
    (perpetualPositionSide === PositionSide.SHORT && orderSide === OrderSide.SELL);
}
