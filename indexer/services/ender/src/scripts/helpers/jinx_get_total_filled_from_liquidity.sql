CREATE OR REPLACE FUNCTION jinx_get_total_filled(fill_liquidity text, event_data jsonb) RETURNS numeric AS $$
/**
  Returns the order total filled amount given the liquidity side.

  (Note that no text should exist before the function declaration to ensure that exception line numbers are correct.)
*/
BEGIN
    IF fill_liquidity = 'TAKER' THEN
        RETURN jinx_from_jsonlib_long(event_data->'totalFilledTaker');
    ELSE
        RETURN jinx_from_jsonlib_long(event_data->'totalFilledMaker');
    END IF;
END;
$$ LANGUAGE plpgsql IMMUTABLE PARALLEL SAFE;
