import { KafkaTopics } from '@dydxprotocol-indexer/kafka';
import {
  Liquidity,
  PerpetualPositionColumns,
  PerpetualPositionFromDatabase,
  SubaccountMessageContents,
} from '@dydxprotocol-indexer/postgres';
import {
  StatefulOrderEventV1,
  IndexerTendermintEvent,
  CandleMessage,
  LiquidationOrderV1,
  MarketCreateEventV1,
  MarketEventV1,
  MarketMessage,
  MarketModifyEventV1,
  MarketPriceUpdateEventV1,
  IndexerOrder,
  OrderFillEventV1,
  SubaccountMessage,
  SubaccountUpdateEventV1,
  TradeMessage,
  TransferEventV1,
  OffChainUpdateV1,
  FundingEventV1_Type,
  FundingEventV1,
  FundingUpdateV1,
  AssetCreateEventV1,
  PerpetualMarketCreateEventV1,
  LiquidityTierUpsertEventV1,
  UpdatePerpetualEventV1,
  UpdateClobPairEventV1,
  DeleveragingEventV1,
  TradingRewardsEventV1,
} from '@dydxprotocol-indexer/v4-protos';
import Long from 'long';

// Type sourced from protocol:
// https://github.com/jinxprotocol/v4-chain/blob/main/protocol/indexer/events/constants.go
export enum JinxIndexerSubtypes {
  ORDER_FILL = 'order_fill',
  SUBACCOUNT_UPDATE = 'subaccount_update',
  TRANSFER = 'transfer',
  MARKET = 'market',
  STATEFUL_ORDER = 'stateful_order',
  FUNDING = 'funding_values',
  ASSET = 'asset',
  PERPETUAL_MARKET = 'perpetual_market',
  LIQUIDITY_TIER = 'liquidity_tier',
  UPDATE_PERPETUAL = 'update_perpetual',
  UPDATE_CLOB_PAIR = 'update_clob_pair',
  DELEVERAGING = 'deleveraging',
  TRADING_REWARD = 'trading_reward',
}

// Generic interface used for creating the Handler objects
// eslint-disable-next-line @typescript-eslint/no-explicit-any
export type EventMessage = any;

export type EventProtoWithTypeAndVersion = {
  type: string,
  eventProto: EventMessage,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} & ({
  type: JinxIndexerSubtypes.ORDER_FILL,
  eventProto: OrderFillEventV1,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} | {
  type: JinxIndexerSubtypes.SUBACCOUNT_UPDATE,
  eventProto: SubaccountUpdateEventV1,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} | {
  type: JinxIndexerSubtypes.TRANSFER,
  eventProto: TransferEventV1,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} | {
  type: JinxIndexerSubtypes.MARKET,
  eventProto: MarketEventV1,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} | {
  type: JinxIndexerSubtypes.STATEFUL_ORDER,
  eventProto: StatefulOrderEventV1,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} | {
  type: JinxIndexerSubtypes.FUNDING,
  eventProto: FundingEventV1,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} | {
  type: JinxIndexerSubtypes.ASSET,
  eventProto: AssetCreateEventV1,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} | {
  type: JinxIndexerSubtypes.PERPETUAL_MARKET,
  eventProto: PerpetualMarketCreateEventV1,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} | {
  type: JinxIndexerSubtypes.LIQUIDITY_TIER,
  eventProto: LiquidityTierUpsertEventV1,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} | {
  type: JinxIndexerSubtypes.UPDATE_PERPETUAL,
  eventProto: UpdatePerpetualEventV1,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} | {
  type: JinxIndexerSubtypes.UPDATE_CLOB_PAIR,
  eventProto: UpdateClobPairEventV1,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} | {
  type: JinxIndexerSubtypes.DELEVERAGING,
  eventProto: DeleveragingEventV1,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} | {
  type: JinxIndexerSubtypes.TRADING_REWARD,
  eventProto: TradingRewardsEventV1,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
});

// Events grouped into events block events and events for each transactionIndex
export interface GroupedEvents {
  transactionEvents: EventProtoWithTypeAndVersion[][],
  blockEvents: EventProtoWithTypeAndVersion[],
}

export type MarketPriceUpdateEventMessage = {
  marketId: number,
  priceUpdate: MarketPriceUpdateEventV1,
};

export type MarketCreateEventMessage = {
  marketId: number,
  marketCreate: MarketCreateEventV1,
};

export type MarketModifyEventMessage = {
  marketId: number,
  marketModify: MarketModifyEventV1,
};

export type OrderFillEventWithOrder = {
  makerOrder: IndexerOrder,
  order: IndexerOrder,
  fillAmount: Long,
  totalFilledMaker: Long,
  totalFilledTaker: Long,
  makerFee: Long,
  takerFee: Long,
};

export type OrderFillEventWithLiquidation = {
  makerOrder: IndexerOrder,
  liquidationOrder: LiquidationOrderV1,
  fillAmount: Long,
  totalFilledMaker: Long,
  totalFilledTaker: Long,
  makerFee: Long,
  takerFee: Long,
};

export type FundingEventMessage = {
  type: FundingEventV1_Type.TYPE_FUNDING_RATE_AND_INDEX | FundingEventV1_Type.TYPE_PREMIUM_SAMPLE,
  updates: FundingUpdateV1[],
};

export type SumFields = PerpetualPositionColumns.sumOpen | PerpetualPositionColumns.sumClose;
export type PriceFields = PerpetualPositionColumns.entryPrice |
PerpetualPositionColumns.exitPrice;

export type OrderFillEventWithLiquidity = {
  event: OrderFillEventV1,
  liquidity: Liquidity,
};

export interface PositionWithPnl extends PerpetualPositionFromDatabase {
  realizedPnl?: string,
  unrealizedPnl?: string,
}

export interface SingleTradeMessage extends TradeMessage {
  transactionIndex: number,
  eventIndex: number,
}

export interface AnnotatedSubaccountMessage extends SubaccountMessage {
  orderId?: string,
  isFill?: boolean,
  subaccountMessageContents?: SubaccountMessageContents,
}

export interface VulcanMessage {
  key: Buffer,
  value: OffChainUpdateV1,
}

export type ConsolidatedKafkaEvent = {
  topic: KafkaTopics.TO_WEBSOCKETS_SUBACCOUNTS,
  message: AnnotatedSubaccountMessage,
} | {
  topic: KafkaTopics.TO_WEBSOCKETS_TRADES,
  message: SingleTradeMessage,
} | {
  topic: KafkaTopics.TO_WEBSOCKETS_MARKETS,
  message: MarketMessage,
} | {
  topic: KafkaTopics.TO_WEBSOCKETS_CANDLES,
  message: CandleMessage,
} | {
  topic: KafkaTopics.TO_VULCAN,
  message: VulcanMessage,
};

export enum TransferEventType {
  DEPOSIT = 'deposit',
  WITHDRAWAL = 'withdrawal',
  TRANSFER = 'transfer',
}
