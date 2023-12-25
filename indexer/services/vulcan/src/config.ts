/**
 * Environment variables required by Vulcan.
 */

import {
  parseNumber,
  parseSchema,
  baseConfigSchema,
  parseBoolean,
} from '@jinxprotocol-indexer/base';
import {
  kafkaConfigSchema,
} from '@jinxprotocol-indexer/kafka';
import {
  redisConfigSchema,
} from '@jinxprotocol-indexer/redis';

export const configSchema = {
  ...baseConfigSchema,
  ...kafkaConfigSchema,
  ...redisConfigSchema,

  FLUSH_KAFKA_MESSAGES_INTERVAL_MS: parseNumber({
    default: 10,
  }),
  MAX_WEBSOCKET_MESSAGES_TO_QUEUE_PER_TOPIC: parseNumber({
    default: 20,
  }),
  // Set this flag to false during fast sync.
  SEND_WEBSOCKET_MESSAGES: parseBoolean({
    default: true,
  }),
};

export default parseSchema(configSchema);
