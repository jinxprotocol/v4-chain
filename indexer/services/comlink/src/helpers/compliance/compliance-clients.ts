import {
  ComplianceClient,
  getComplianceClient,
} from '@jinxprotocol-indexer/compliance';
import { ComplianceProvider } from '@jinxprotocol-indexer/postgres';

export interface ClientAndProvider {
  client: ComplianceClient;
  provider: ComplianceProvider;
}

export const complianceProvider: ClientAndProvider = {
  client: getComplianceClient(),
  provider: ComplianceProvider.ELLIPTIC,
};
