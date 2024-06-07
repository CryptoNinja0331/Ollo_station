package keeper_test

import (
	"testing"

	testkeeper "github.com/ollo-station/ollo/testutil/keeper"
	"github.com/ollo-station/ollo/x/ons/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.OnsKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
