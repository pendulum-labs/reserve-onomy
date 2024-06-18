package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var (
	// KeyBurnRate is byte key for BurnRate param.
	KeyBurnRate = []byte("BurnRate") //nolint:gochecknoglobals // cosmos-sdk style
	// KeyBurnCoin is byte key for BurnCoin param.
	KeyBurnCoin = []byte("BurnCoin") //nolint:gochecknoglobals // cosmos-sdk style
	// KeyBurnCoin is byte key for BurnCoin param.
	KeyLiquidatorReward = []byte("LiquidatorReward") //nolint:gochecknoglobals // cosmos-sdk style
)

var (
	// DefaultBurnRate is default value for the DefaultBurnRate param.
	DefaultBurnRate = "1000" //nolint:gomnd,gochecknoglobals // cosmos-sdk style
	// DefaultBurnCoin is default value for the DefaultBurnCoin param.
	DefaultBurnCoin = "stake" //nolint:gomnd,gochecknoglobals // cosmos-sdk style
	// DefaultLiquidatorReward is default value for the LiquidatorReward param.
	DefaultLiquidatorReward = "0010" //nolint:gomnd,gochecknoglobals // cosmos-sdk style

)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	burnRate string,
	burnCoin string,
	liquidatorReward string,
) Params {
	return Params{
		BurnRate:         burnRate,
		BurnCoin:         burnCoin,
		LiquidatorReward: liquidatorReward,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(DefaultBurnRate, DefaultBurnCoin, DefaultLiquidatorReward)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyBurnRate, &p.BurnRate, validateBurnRate),
		paramtypes.NewParamSetPair(KeyBurnCoin, &p.BurnCoin, validateBurnCoin),
		paramtypes.NewParamSetPair(KeyLiquidatorReward, &p.LiquidatorReward, validateLiquidatorReward),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateBurnRate(p.BurnRate); err != nil {
		return err
	}
	if err := validateBurnCoin(p.BurnCoin); err != nil {
		return err
	}
	if err := validateLiquidatorReward(p.LiquidatorReward); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p) //nolint:errcheck // error is not expected here
	return string(out)
}

func validateBurnRate(i interface{}) error {
	value, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	burnRate, ok := sdk.NewIntFromString(value)
	if !ok {
		return fmt.Errorf("invalid string number format: %q", value)
	}
	if burnRate.LTE(sdk.ZeroInt()) {
		return fmt.Errorf("burn rate numerator must be positive and greater than zero: %d", burnRate)
	}
	if burnRate.GTE(sdk.NewInt(10000)) {
		return fmt.Errorf("burn rate numerator must be less than 10000: %d", burnRate)
	}

	return nil
}

func validateBurnCoin(i interface{}) error {
	_, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateLiquidatorReward(i interface{}) error {
	value, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	liquidatorReward, ok := sdk.NewIntFromString(value)
	if !ok {
		return fmt.Errorf("invalid string number format: %q", value)
	}
	if liquidatorReward.LTE(sdk.ZeroInt()) {
		return fmt.Errorf("market fee numerator must be positive and greater than zero: %d", liquidatorReward)
	}
	if liquidatorReward.GTE(sdk.NewInt(10000)) {
		return fmt.Errorf("market fee numerator must be less than 10000: %d", liquidatorReward)
	}

	return nil
}
