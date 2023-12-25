import { baseConfigSchema, parseSchema } from '@jinxprotocol-indexer/base';
import { kafkaConfigSchema } from '@jinxprotocol-indexer/kafka';
import { postgresConfigSchema } from '@jinxprotocol-indexer/postgres';

export const configSchema = {
  ...baseConfigSchema,
  ...postgresConfigSchema,
  ...kafkaConfigSchema,
};

export default parseSchema(configSchema);
