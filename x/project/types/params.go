package types

import (
	"fmt"

	xcotypes "github.com/petrinetwork/xco-blockchain/lib/xco"

	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
	didexported "github.com/petrinetwork/xco-blockchain/lib/legacydid"
)

// Parameter store keys
var (
	KeyXcoDid                       = []byte("XcoDid")
	KeyProjectMinimumInitialFunding = []byte("ProjectMinimumInitialFunding")
	KeyOracleFeePercentage          = []byte("OracleFeePercentage")
	KeyNodeFeePercentage            = []byte("NodeFeePercentage")
)

// ParamTable for project module.
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(projectMinimumInitialFunding sdk.Coins, xcoDid didexported.Did,
	oracleFeePercentage, nodeFeePercentage sdk.Dec) Params {
	return Params{
		XcoDid:                       xcoDid,
		ProjectMinimumInitialFunding: projectMinimumInitialFunding,
		OracleFeePercentage:          oracleFeePercentage,
		NodeFeePercentage:            nodeFeePercentage,
	}

}

// default project module parameters
func DefaultParams() Params {
	defaultXcoDid := didexported.Did("did:xco:U4tSpzzv91HHqWW1YmFkHJ")
	defaultMinInitFunding := sdk.NewCoins(sdk.NewCoin(
		xcotypes.XcoNativeToken, sdk.OneInt()))
	tenPercentFee := sdk.NewDec(10)

	return Params{
		XcoDid:                       defaultXcoDid,         // invalid blank
		ProjectMinimumInitialFunding: defaultMinInitFunding, // 1uxco
		OracleFeePercentage:          tenPercentFee,         // 10.0 (10%)
		NodeFeePercentage:            tenPercentFee,         // 10.0 (10%)
	}
}

func validateXcoDid(i interface{}) error {
	v, ok := i.(didexported.Did)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if len(v) == 0 {
		return fmt.Errorf("xco did cannot be empty")
	}

	return nil
}

func validateProjectMinimumInitialFunding(i interface{}) error {
	v, ok := i.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsAnyNegative() {
		return fmt.Errorf("invalid project minimum initial "+
			"funding should be positive, is %s ", v.String())
	}

	return nil
}

func validateOracleFeePercentage(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.LT(sdk.ZeroDec()) {
		return fmt.Errorf("invalid parameter oracle fee percentage; should be >= 0.0, is %s ", v.String())
	} else if v.GT(sdk.NewDec(100)) {
		return fmt.Errorf("invalid parameter oracle fee percentage; should be <= 100, is %s ", v.String())
	}

	return nil
}

func validateNodeFeePercentage(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.LT(sdk.ZeroDec()) {
		return fmt.Errorf("invalid parameter node fee percentage; should be >= 0.0, is %s ", v.String())
	} else if v.GT(sdk.NewDec(100)) {
		return fmt.Errorf("invalid parameter node fee percentage; should be <= 100, is %s ", v.String())
	}

	return nil
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{KeyXcoDid, &p.XcoDid, validateXcoDid},
		{KeyProjectMinimumInitialFunding, &p.ProjectMinimumInitialFunding, validateProjectMinimumInitialFunding},
		{KeyOracleFeePercentage, &p.OracleFeePercentage, validateOracleFeePercentage},
		{KeyNodeFeePercentage, &p.NodeFeePercentage, validateNodeFeePercentage},
	}
}
