package ignit_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "ignit/testutil/keeper"
	"ignit/testutil/nullify"
	"ignit/x/ignit"
	"ignit/x/ignit/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IgnitKeeper(t)
	ignit.InitGenesis(ctx, *k, genesisState)
	got := ignit.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
