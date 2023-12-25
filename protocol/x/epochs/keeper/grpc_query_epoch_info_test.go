package keeper_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	keepertest "github.com/jinxprotocol/v4-chain/protocol/testutil/keeper"
	"github.com/jinxprotocol/v4-chain/protocol/testutil/nullify"
	"github.com/jinxprotocol/v4-chain/protocol/x/epochs/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestEpochInfoQuerySingle(t *testing.T) {
	ctx, keeper, _ := keepertest.EpochsKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNEpochInfo(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetEpochInfoRequest
		response *types.QueryEpochInfoResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetEpochInfoRequest{
				Name: msgs[0].Name,
			},
			response: &types.QueryEpochInfoResponse{EpochInfo: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetEpochInfoRequest{
				Name: msgs[1].Name,
			},
			response: &types.QueryEpochInfoResponse{EpochInfo: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetEpochInfoRequest{
				Name: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.EpochInfo(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response), //nolint:staticcheck
					nullify.Fill(response),    //nolint:staticcheck
				)
			}
		})
	}
}

func TestEpochInfoQueryPaginated(t *testing.T) {
	ctx, keeper, _ := keepertest.EpochsKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNEpochInfo(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllEpochInfoRequest {
		return &types.QueryAllEpochInfoRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.EpochInfoAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.EpochInfo), step)
			require.Subset(t,
				nullify.Fill(msgs),           //nolint:staticcheck
				nullify.Fill(resp.EpochInfo), //nolint:staticcheck
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.EpochInfoAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.EpochInfo), step)
			require.Subset(t,
				nullify.Fill(msgs),           //nolint:staticcheck
				nullify.Fill(resp.EpochInfo), //nolint:staticcheck
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.EpochInfoAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),           //nolint:staticcheck
			nullify.Fill(resp.EpochInfo), //nolint:staticcheck
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.EpochInfoAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
