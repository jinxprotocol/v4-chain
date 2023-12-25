CREATE OR REPLACE FUNCTION jinx_uuid_from_subaccount_id_parts(address text, subaccount_number text) RETURNS uuid AS $$
/**
  Returns a UUID using the parts of an IndexerSubaccountId (https://github.com/jinxprotocol/v4-chain/blob/9ed26bd/proto/jinxprotocol/indexer/protocol/v1/subaccount.proto#L15).

  (Note that no text should exist before the function declaration to ensure that exception line numbers are correct.)
*/
BEGIN
    RETURN jinx_uuid(concat(address, '-', subaccount_number));
END;
$$ LANGUAGE plpgsql IMMUTABLE PARALLEL SAFE;
