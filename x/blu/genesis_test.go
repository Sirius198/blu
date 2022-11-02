package blu_test

import (
	"testing"

	keepertest "blu/testutil/keeper"
	"blu/testutil/nullify"
	"blu/x/blu"
	"blu/x/blu/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BluKeeper(t)
	blu.InitGenesis(ctx, *k, genesisState)
	got := blu.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
