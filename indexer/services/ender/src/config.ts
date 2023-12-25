/**
 * Environment variables required by Ender.
 */

import {
  parseSchema,
  baseConfigSchema,
  parseBoolean,
} from '@jinxprotocol-indexer/base';
import {
  kafkaConfigSchema,
} from '@jinxprotocol-indexer/kafka';
import {
  postgresConfigSchema,
} from '@jinxprotocol-indexer/postgres';
import { redisConfigSchema } from '@jinxprotocol-indexer/redis';

export const configSchema = {
  ...baseConfigSchema,
  ...postgresConfigSchema,
  ...redisConfigSchema,
  ...kafkaConfigSchema,
  SEND_WEBSOCKET_MESSAGES: parseBoolean({
    default: true,
  }),
};

export default parseSchema(configSchema);
