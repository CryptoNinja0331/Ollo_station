package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ollo-station/ollo/testutil"
	utils "github.com/ollo-station/ollo/x/ollo/types"
	"github.com/ollo-station/ollo/x/exchange/keeper"
	"github.com/ollo-station/ollo/x/exchange/types"
)

var enoughCoins = utils.ParseCoins(
	"10000_000000000000000000ucre,10000_000000000000000000uatom,10000_000000000000000000uusd,10000_000000000000000000stake")

type KeeperTestSuite struct {
	testutil.TestSuite
	keeper  keeper.Keeper
	querier keeper.Querier
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (s *KeeperTestSuite) SetupTest() {
	s.TestSuite.SetupTest()
	s.keeper = s.App.ExchangeKeeper
	s.querier = keeper.Querier{Keeper: s.keeper}
	s.FundAccount(utils.TestAddress(0), utils.ParseCoins("1ucre,1uusd,1uatom")) // make positive supplies
}

func (s *KeeperTestSuite) TestSetOrderSources() {
	// Same source name
	s.Require().PanicsWithValue("duplicate order source name: a", func() {
		k := keeper.Keeper{}
		k.SetOrderSources(types.NewMockOrderSource("a"), types.NewMockOrderSource("a"))
	})
	k := keeper.Keeper{}
	k.SetOrderSources(types.NewMockOrderSource("a"), types.NewMockOrderSource("b"))
	// Already set
	s.Require().PanicsWithValue("cannot set order sources twice", func() {
		s.keeper.SetOrderSources(types.NewMockOrderSource("b"), types.NewMockOrderSource("c"))
	})
}

func (s *KeeperTestSuite) createLiquidity(
	marketId uint64, ordererAddr sdk.AccAddress, centerPrice, totalQty sdk.Dec) {
	tick := types.TickAtPrice(centerPrice)
	interval := types.PriceIntervalAtTick(tick + 10*10)
	for i := 0; i < 10; i++ {
		sellPrice := centerPrice.Add(interval.MulInt64(int64(i+1) * 10))
		buyPrice := centerPrice.Sub(interval.MulInt64(int64(i+1) * 10))

		qty := totalQty.QuoInt64(200).Add(totalQty.QuoInt64(100).MulInt64(int64(i)))
		s.PlaceLimitOrder(marketId, ordererAddr, false, sellPrice, qty, time.Hour)
		s.PlaceLimitOrder(marketId, ordererAddr, true, buyPrice, qty, time.Hour)
	}
}

func (s *KeeperTestSuite) createLiquidity2(
	marketId uint64, ordererAddr sdk.AccAddress, centerPrice, maxOrderPriceRatio, qtyPerTick sdk.Dec) {
	minPrice, maxPrice := types.OrderPriceLimit(centerPrice, maxOrderPriceRatio)
	for i := 1; ; i++ {
		buyPrice := types.PriceAtTick(types.TickAtPrice(centerPrice) - 100*int32(i))
		if buyPrice.LT(minPrice) {
			break
		}
		s.PlaceLimitOrder(
			marketId, ordererAddr, true, buyPrice, qtyPerTick, time.Hour)
	}
	for i := 1; ; i++ {
		sellPrice := types.PriceAtTick(types.TickAtPrice(centerPrice) + 100*int32(i))
		if sellPrice.GT(maxPrice) {
			break
		}
		s.PlaceLimitOrder(
			marketId, ordererAddr, false, sellPrice, qtyPerTick, time.Hour)
	}
}
