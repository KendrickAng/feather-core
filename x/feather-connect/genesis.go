package feather_connect

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/terra-money/feather-core/x/feather-connect/types"
)

// ValidateGenesis
func ValidateGenesis(data *types.GenesisState) error {
	if data.Params.AllianceBondHeight < 1 {
		panic(fmt.Errorf("AllianceBondHeight cannot be less than 1 on genesis state"))
	}

	if err := sdk.ValidateDenom(data.Params.BaseDenom); err != nil {
		panic(fmt.Errorf("invalid denom on genesis state: %s", err))
	}

	alliance := data.Params.Alliance

	if len(alliance.Title) == 0 {
		panic(fmt.Errorf("title is empty on genesis state"))
	}

	if len(alliance.Description) == 0 {
		panic(fmt.Errorf("description is empty on genesis state"))
	}

	if err := sdk.ValidateDenom(alliance.Denom); err != nil {
		panic(fmt.Errorf("invalid denom on genesis state: %s", err))
	}

	if alliance.RewardWeight.IsNil() || alliance.RewardWeight.IsNegative() {
		panic(fmt.Errorf("rewardWeight cannot be negative nor nil on genesis state"))
	}

	if alliance.TakeRate.IsNil() || alliance.TakeRate.IsNegative() {
		panic(fmt.Errorf("takeRate cannot be negative nor nil on genesis state"))
	}

	if alliance.RewardChangeRate.IsNil() || alliance.RewardChangeRate.IsNegative() {
		panic(fmt.Errorf("rewardChangeRate cannot be negative nor nil on genesis state"))
	}

	if alliance.RewardChangeInterval < 0 {
		panic(fmt.Errorf("rewardChangeInterval cannot be negative nor nil on genesis state"))
	}

	if alliance.RewardWeightRange.Min.IsNegative() || alliance.RewardWeightRange.Max.IsNegative() {
		panic(fmt.Errorf("rewardWeightRange Min or Max cannot be negative on genesis state"))
	}

	return nil
}

func DefaultGenesisState() *types.GenesisState {
	return &types.GenesisState{
		Params: types.DefaultParams(),
	}
}
