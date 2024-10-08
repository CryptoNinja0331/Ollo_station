package keeper_test

import (
	utils "github.com/ollo-station/ollo/x/ollo/types"
)

func (s *KeeperTestSuite) TestCreateMarket() {
	creatorAddr := s.FundedAccount(1, enoughCoins)
	_, err := s.keeper.CreateMarket(s.Ctx, creatorAddr, "nonexistent1", "nonexistent2")
	s.Require().EqualError(err, "base denom nonexistent1 has no supply: invalid request")
	_, err = s.keeper.CreateMarket(s.Ctx, creatorAddr, "uollo", "nonexistent2")
	s.Require().EqualError(err, "quote denom nonexistent2 has no supply: invalid request")

	s.CreateMarket("uollo", "uusd")

	_, err = s.keeper.CreateMarket(s.Ctx, creatorAddr, "uollo", "uusd")
	s.Require().EqualError(err, "market already exists: 1: invalid request")

	emptyAddr := utils.TestAddress(2)
	_, err = s.keeper.CreateMarket(s.Ctx, emptyAddr, "uatom", "uusd")
	s.Require().EqualError(err, "insufficient market creation fee: 0stake is smaller than 1000000stake: insufficient funds")
}
