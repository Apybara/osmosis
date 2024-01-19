package apybara_indexer

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/types"
	"gorm.io/gorm"
	"strconv"
)

type RewardCalculatorService struct {
	Database *gorm.DB
}

type BlockerAmount struct {
	BeforeBeginBlocker types.Dec
	AfterBeginBlocker  types.Dec
	Denom              string
}

func (r RewardCalculatorService) RewardDeltaForBlockers(ctx types.Context, blocker BlockerAmount, database *gorm.DB) error {

	afterBeginBlocker, err := strconv.ParseFloat(blocker.AfterBeginBlocker.String(), 64)
	if err != nil {
		fmt.Println("error parsing float for afterBeginBlocker", err)
		return err
	}

	beforeBeginBlocker, err := strconv.ParseFloat(blocker.BeforeBeginBlocker.String(), 64)
	if err != nil {
		fmt.Println("error parsing float for beforeBeginBlocker", err)
		return err
	}

	rewardDelta := afterBeginBlocker - beforeBeginBlocker

	//
	var rewardDataDelta RewardDataDelta
	rewardDataDelta.AfterBeginBlockerAmount = blocker.AfterBeginBlocker.String()
	rewardDataDelta.BeforeBeginBlockerAmount = blocker.BeforeBeginBlocker.String()
	rewardDataDelta.Denom = blocker.Denom
	rewardDataDelta.Timestamp = ctx.BlockTime().Unix()
	rewardDataDelta.BlockHeight = ctx.BlockHeight()
	rewardDataDelta.Delta = fmt.Sprintf("%.18f", rewardDelta)
	database.Create(&rewardDataDelta)
	//}
	return nil
}
