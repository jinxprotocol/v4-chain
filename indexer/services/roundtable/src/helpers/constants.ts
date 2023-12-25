import { testConstants } from '@jinxprotocol-indexer/postgres';
import { PnlTickForSubaccounts } from '@jinxprotocol-indexer/redis';

export const defaultPnlTickForSubaccounts: PnlTickForSubaccounts = {
  [testConstants.defaultSubaccountId]: testConstants.defaultPnlTick,
  [testConstants.defaultSubaccountId2]: {
    ...testConstants.defaultPnlTick,
    subaccountId: testConstants.defaultSubaccountId2,
    equity: '9000',
  },
};
