package types

import (
	"fmt"
	"time"

	"gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Service params default values
var (
	DefaultMaxRequestTimeout    = int64(100)
	DefaultMinDepositMultiple   = int64(200)
	DefaultMinDeposit           = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(6000))) // 6000stake
	DefaultServiceFeeTax        = sdk.NewDecWithPrec(1, 1)                                          // 10%
	DefaultSlashFraction        = sdk.NewDecWithPrec(1, 3)                                          // 0.1%
	DefaultComplaintRetrospect  = 15 * 24 * time.Hour                                               // 15 days
	DefaultArbitrationTimeLimit = 5 * 24 * time.Hour                                                // 5 days
	DefaultTxSizeLimit          = uint64(4000)
	DefaultBaseDenom            = sdk.DefaultBondDenom
)

// Keys for parameter access
// nolint
var (
	KeyMaxRequestTimeout    = []byte("MaxRequestTimeout")
	KeyMinDepositMultiple   = []byte("MinDepositMultiple")
	KeyMinDeposit           = []byte("MinDeposit")
	KeyServiceFeeTax        = []byte("ServiceFeeTax")
	KeySlashFraction        = []byte("SlashFraction")
	KeyComplaintRetrospect  = []byte("ComplaintRetrospect")
	KeyArbitrationTimeLimit = []byte("ArbitrationTimeLimit")
	KeyTxSizeLimit          = []byte("TxSizeLimit")
	KeyBaseDenom            = []byte("BaseDenom")
)

var _ paramstypes.ParamSet = (*Params)(nil)

// NewParams creates a new Params instance
func NewParams(
	maxRequestTimeout,
	minDepositMultiple int64,
	minDeposit sdk.Coins,
	serviceFeeTax,
	slashFraction sdk.Dec,
	complaintRetrospect,
	arbitrationTimeLimit time.Duration,
	txSizeLimit uint64,
	baseDenom string,
) Params {
	return Params{
		MaxRequestTimeout:    maxRequestTimeout,
		MinDepositMultiple:   minDepositMultiple,
		MinDeposit:           minDeposit,
		ServiceFeeTax:        serviceFeeTax,
		SlashFraction:        slashFraction,
		ComplaintRetrospect:  complaintRetrospect,
		ArbitrationTimeLimit: arbitrationTimeLimit,
		TxSizeLimit:          txSizeLimit,
		BaseDenom:            baseDenom,
	}
}

// ParamSetPairs implements paramstypes.ParamSet
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(KeyMaxRequestTimeout, &p.MaxRequestTimeout, validateMaxRequestTimeout),
		paramstypes.NewParamSetPair(KeyMinDepositMultiple, &p.MinDepositMultiple, validateMinDepositMultiple),
		paramstypes.NewParamSetPair(KeyMinDeposit, &p.MinDeposit, validateMinDeposit),
		paramstypes.NewParamSetPair(KeyServiceFeeTax, &p.ServiceFeeTax, validateServiceFeeTax),
		paramstypes.NewParamSetPair(KeySlashFraction, &p.SlashFraction, validateSlashFraction),
		paramstypes.NewParamSetPair(KeyComplaintRetrospect, &p.ComplaintRetrospect, validateComplaintRetrospect),
		paramstypes.NewParamSetPair(KeyArbitrationTimeLimit, &p.ArbitrationTimeLimit, validateArbitrationTimeLimit),
		paramstypes.NewParamSetPair(KeyTxSizeLimit, &p.TxSizeLimit, validateTxSizeLimit),
		paramstypes.NewParamSetPair(KeyBaseDenom, &p.BaseDenom, validateTxBaseDenom),
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return NewParams(
		DefaultMaxRequestTimeout,
		DefaultMinDepositMultiple,
		DefaultMinDeposit,
		DefaultServiceFeeTax,
		DefaultSlashFraction,
		DefaultComplaintRetrospect,
		DefaultArbitrationTimeLimit,
		DefaultTxSizeLimit,
		DefaultBaseDenom,
	)
}

// String implements stringer
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// Validate validates a set of params
func (p Params) Validate() error {
	if err := validateMaxRequestTimeout(p.MaxRequestTimeout); err != nil {
		return err
	}
	if err := validateMinDepositMultiple(p.MinDepositMultiple); err != nil {
		return err
	}
	if err := validateMinDeposit(p.MinDeposit); err != nil {
		return err
	}
	if err := validateSlashFraction(p.SlashFraction); err != nil {
		return err
	}
	if err := validateServiceFeeTax(p.ServiceFeeTax); err != nil {
		return err
	}
	if err := validateComplaintRetrospect(p.ComplaintRetrospect); err != nil {
		return err
	}
	if err := validateArbitrationTimeLimit(p.ArbitrationTimeLimit); err != nil {
		return err
	}
	if err := sdk.ValidateDenom(p.BaseDenom); err != nil {
		return err
	}

	return validateTxSizeLimit(p.TxSizeLimit)
}

func validateMaxRequestTimeout(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("maximum request timeout must be positive: %d", v)
	}

	return nil
}

func validateMinDepositMultiple(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("minimum deposit multiple must be positive: %d", v)
	}

	return nil
}

func validateMinDeposit(i interface{}) error {
	v, ok := i.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !v.IsValid() {
		return fmt.Errorf("invalid minimum deposit: %s", v)
	}

	return nil
}

func validateSlashFraction(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.LT(sdk.ZeroDec()) || v.GT(sdk.OneDec()) {
		return fmt.Errorf("slashing fraction must be between [0, 1]: %s", v)
	}

	return nil
}

func validateServiceFeeTax(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.LT(sdk.ZeroDec()) || v.GTE(sdk.OneDec()) {
		return fmt.Errorf("service fee tax must be between [0, 1): %s", v)
	}

	return nil
}

func validateComplaintRetrospect(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("complaint retrospect must be positive: %d", v)
	}

	return nil
}

func validateArbitrationTimeLimit(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("arbitration time limit must be positive: %d", v)
	}

	return nil
}

func validateTxSizeLimit(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("tx size limit must be positive: %d", v)
	}

	return nil
}

func validateTxBaseDenom(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return sdk.ValidateDenom(v)
}
