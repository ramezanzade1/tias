package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "ignit/testutil/keeper"
	"ignit/x/ignit/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.IgnitKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
