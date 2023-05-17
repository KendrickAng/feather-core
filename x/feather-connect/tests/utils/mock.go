package test_utils

import (
	"time"

	sdkmath "cosmossdk.io/math"
	trasnfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	ibcchanneltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	commitmenttypes "github.com/cosmos/ibc-go/v7/modules/core/23-commitment/types"
	ibctm "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"
	ibctesting "github.com/cosmos/ibc-go/v7/testing"
	alliancetypes "github.com/terra-money/alliance/x/alliance/types"
	"github.com/terra-money/feather-core/x/feather-connect/types"
)

var (
	MockedParams = types.GenesisState{
		Params: types.Params{
			HaltIfNoChannel:    true,
			BaseDenom:          "uluna",
			BaseChainId:        "phoenix-1",
			AllianceBondHeight: 2,
			Alliance: alliancetypes.MsgCreateAllianceProposal{
				Title:                "test",
				Description:          "test",
				RewardWeight:         sdkmath.LegacyNewDec(1),
				TakeRate:             sdkmath.LegacyNewDec(1),
				RewardChangeRate:     sdkmath.LegacyNewDec(1),
				RewardChangeInterval: 100,
				RewardWeightRange: alliancetypes.RewardWeightRange{
					Min: sdkmath.LegacyNewDec(1),
					Max: sdkmath.LegacyNewDec(1),
				},
			},
		},
	}

	MockedChannel = ibcchanneltypes.Channel{
		State: ibcchanneltypes.OPEN,
		ConnectionHops: []string{
			"connection-0",
		},
	}

	MockedDenomTrace = trasnfertypes.DenomTrace{
		BaseDenom: "uluna",
		Path:      "transfer/channel-0",
	}

	MockedClientState = ibctm.NewClientState(
		"phoenix-1",
		ibctm.DefaultTrustLevel,
		time.Second,
		time.Second,
		time.Second,
		clienttypes.ZeroHeight(),
		commitmenttypes.GetSDKSpecs(),
		ibctesting.UpgradePath,
	)
)
