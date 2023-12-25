/**
 * Environment variables required by Bazooka.
 */

import {
  parseSchema,
  baseConfigSchema,
  parseInteger,
  parseBoolean,
} from '@jinxprotocol-indexer/base';
import {
  kafkaConfigSchema,
} from '@jinxprotocol-indexer/kafka';
import {
  postgresConfigSchema,
} from '@jinxprotocol-indexer/postgres';
import {
  redisConfigSchema,
} from '@jinxprotocol-indexer/redis';

export const configSchema = {
  ...baseConfigSchema,
  ...postgresConfigSchema,
  ...kafkaConfigSchema,
  ...redisConfigSchema,
  // Parameters for exponential backoff when retrying clearing Kafka.
  // Max wait time between retries of 3*(2**2) = 12 seconds.
  CLEAR_KAFKA_TOPIC_RETRY_MS: parseInteger({ default: 3000 }),
  CLEAR_KAFKA_TOPIC_MAX_RETRIES: parseInteger({ default: 3 }),
  PREVENT_BREAKING_CHANGES_WITHOUT_FORCE: parseBoolean({
    default: true,
  }),
};

export default parseSchema(configSchema);
