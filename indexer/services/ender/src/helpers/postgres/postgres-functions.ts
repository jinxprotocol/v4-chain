import { readFileSync } from 'fs';
import path from 'path';

import { logger } from '@dydxprotocol-indexer/base';
import { dbHelpers, storeHelpers } from '@dydxprotocol-indexer/postgres';

export type PostgresFunction = {
  // The name of the script
  readonly name: string;
  // The contents of the script
  readonly script: string;
};

/**
 * Loads a named script from the specified path.
 *
 * @param name The name of the script.
 * @param scriptPath The path to the script.
 * @returns The created script object
 */
function newScript(name: string, scriptPath: string): PostgresFunction {
  const script: string = readFileSync(path.resolve(__dirname, scriptPath)).toString();
  return {
    name,
    script,
  };
}

const HANDLER_SCRIPTS: string[] = [
  'jinx_asset_create_handler.sql',
  'jinx_block_processor_ordered_handlers.sql',
  'jinx_block_processor_unordered_handlers.sql',
  'jinx_deleveraging_handler.sql',
  'jinx_funding_handler.sql',
  'jinx_liquidity_tier_handler.sql',
  'jinx_market_create_handler.sql',
  'jinx_market_modify_handler.sql',
  'jinx_market_price_update_handler.sql',
  'jinx_perpetual_market_handler.sql',
  'jinx_stateful_order_handler.sql',
  'jinx_subaccount_update_handler.sql',
  'jinx_trading_rewards_handler.sql',
  'jinx_transfer_handler.sql',
  'jinx_update_clob_pair_handler.sql',
  'jinx_update_perpetual_handler.sql',
];

const DB_SETUP_SCRIPTS: string[] = [
  'create_extension_pg_stat_statements.sql',
  'create_extension_uuid_ossp.sql',
];

const HELPER_SCRIPTS: string[] = [
  'jinx_clob_pair_status_to_market_status.sql',
  'jinx_create_initial_rows_for_tendermint_block.sql',
  'jinx_create_tendermint_event.sql',
  'jinx_create_transaction.sql',
  'jinx_event_id_from_parts.sql',
  'jinx_from_jsonlib_long.sql',
  'jinx_from_protocol_order_side.sql',
  'jinx_from_protocol_time_in_force.sql',
  'jinx_from_serializable_int.sql',
  'jinx_get_fee_from_liquidity.sql',
  'jinx_get_order_status.sql',
  'jinx_get_perpetual_market_for_clob_pair.sql',
  'jinx_get_total_filled_from_liquidity.sql',
  'jinx_get_weighted_average.sql',
  'jinx_liquidation_fill_handler_per_order.sql',
  'jinx_order_fill_handler_per_order.sql',
  'jinx_perpetual_position_and_order_side_matching.sql',
  'jinx_process_trading_reward_event.sql',
  'jinx_protocol_condition_type_to_order_type.sql',
  'jinx_tendermint_event_to_transaction_index.sql',
  'jinx_trim_scale.sql',
  'jinx_update_perpetual_position_aggregate_fields.sql',
  'jinx_uuid.sql',
  'jinx_uuid_from_asset_position_parts.sql',
  'jinx_uuid_from_fill_event_parts.sql',
  'jinx_uuid_from_funding_index_update_parts.sql',
  'jinx_uuid_from_oracle_price_parts.sql',
  'jinx_uuid_from_order_id.sql',
  'jinx_uuid_from_order_id_parts.sql',
  'jinx_uuid_from_perpetual_position_parts.sql',
  'jinx_uuid_from_subaccount_id.sql',
  'jinx_uuid_from_subaccount_id_parts.sql',
  'jinx_uuid_from_trading_rewards_parts.sql',
  'jinx_uuid_from_transaction_parts.sql',
  'jinx_uuid_from_transfer_parts.sql',
];

const MAIN_SCRIPTS: string[] = [
  'jinx_block_processor.sql',
];

const SCRIPTS: string[] = [
  ...HANDLER_SCRIPTS.map((script: string) => `handlers/${script}`),
  ...HELPER_SCRIPTS.map((script: string) => `helpers/${script}`),
  ...DB_SETUP_SCRIPTS.map((script: string) => `setup/${script}`),
  ...MAIN_SCRIPTS,
];

export async function createPostgresFunctions(): Promise<void> {
  await Promise.all([
    dbHelpers.createModelToJsonFunctions(),
    ...SCRIPTS.map((script: string) => storeHelpers.rawQuery(newScript(script, `../../scripts/${script}`).script, {})
      .catch((error: Error) => {
        logger.error({
          at: 'postgres-functions#createPostgresFunctions',
          message: `Failed to create or replace function contained in ${script}`,
          error,
        });
        throw error;
      }),
    ),
  ]);
}
