package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// TypeMsgCreatePool is the type for MsgCreatePool
	TypeMsgCreatePool = "create_pool"

	// TypeMsgDestroyPool is the type for MsgDestroyPool
	TypeMsgDestroyPool = "destroy_pool"

	// TypeMsgMsgAppendReward is the type for MsgAppendReward
	TypeMsgAppendReward = "append_reward"

	// TypeMsgStake is the type for MsgStake
	TypeMsgStake = "stake"

	// TypeMsgUnstake is the type for MsgUnstake
	TypeMsgUnstake = "unstake"

	// TypeMsgHarvest is the type for MsgHarvest
	TypeMsgHarvest = "harvest"
)

var (
	_ sdk.Msg = &MsgCreatePool{}
	_ sdk.Msg = &MsgDestroyPool{}
	_ sdk.Msg = &MsgAppendReward{}
	_ sdk.Msg = &MsgStake{}
	_ sdk.Msg = &MsgUnstake{}
	_ sdk.Msg = &MsgHarvest{}
)

// Route implements Msg
func (msg MsgCreatePool) Route() string { return RouterKey }

// Type implements Msg
func (msg MsgCreatePool) Type() string { return TypeMsgCreatePool }

// ValidateBasic implements Msg
func (msg MsgCreatePool) ValidateBasic() error {
	if err := ValidatePoolName(msg.Name); err != nil {
		return err
	}

	if err := ValidateDescription(msg.Description); err != nil {
		return err
	}

	if err := ValidateLpTokenDenom(msg.LpTokenDenom); err != nil {
		return err
	}

	if err := ValidateCoins(msg.RewardPerBlock...); err != nil {
		return err
	}

	if err := ValidateAddress(msg.Creator); err != nil {
		return err
	}

	if err := ValidateCoins(msg.TotalReward...); err != nil {
		return err
	}
	return ValidateReward(msg.RewardPerBlock, msg.TotalReward)
}

// GetSignBytes implements Msg
func (msg MsgCreatePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgCreatePool) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// -----------------------------------------------------------------------------
// Route implements Msg
func (msg MsgDestroyPool) Route() string { return RouterKey }

// Type implements Msg
func (msg MsgDestroyPool) Type() string { return TypeMsgDestroyPool }

// ValidateBasic implements Msg
func (msg MsgDestroyPool) ValidateBasic() error {
	if err := ValidateAddress(msg.Creator); err != nil {
		return err
	}
	return ValidatePoolName(msg.PoolName)
}

// GetSignBytes implements Msg
func (msg MsgDestroyPool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgDestroyPool) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// -----------------------------------------------------------------------------
// Route implements Msg
func (msg MsgAppendReward) Route() string { return RouterKey }

// Type implements Msg
func (msg MsgAppendReward) Type() string { return TypeMsgAppendReward }

// ValidateBasic implements Msg
func (msg MsgAppendReward) ValidateBasic() error {
	if err := ValidateAddress(msg.Creator); err != nil {
		return err
	}

	if err := ValidateCoins(msg.Amount...); err != nil {
		return err
	}
	return ValidatePoolName(msg.PoolName)
}

// GetSignBytes implements Msg
func (msg MsgAppendReward) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgAppendReward) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// -----------------------------------------------------------------------------
// Route implements Msg
func (msg MsgStake) Route() string { return RouterKey }

// Type implements Msg
func (msg MsgStake) Type() string { return TypeMsgStake }

// ValidateBasic implements Msg
func (msg MsgStake) ValidateBasic() error {
	if err := ValidateAddress(msg.Sender); err != nil {
		return err
	}

	if err := ValidateCoins(msg.Amount); err != nil {
		return err
	}
	return ValidatePoolName(msg.PoolName)
}

// GetSignBytes implements Msg
func (msg MsgStake) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgStake) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// -----------------------------------------------------------------------------
// Route implements Msg
func (msg MsgUnstake) Route() string { return RouterKey }

// Type implements Msg
func (msg MsgUnstake) Type() string { return TypeMsgUnstake }

// ValidateBasic implements Msg
func (msg MsgUnstake) ValidateBasic() error {
	if err := ValidateAddress(msg.Sender); err != nil {
		return err
	}

	if err := ValidateCoins(msg.Amount); err != nil {
		return err
	}
	return ValidatePoolName(msg.PoolName)
}

// GetSignBytes implements Msg
func (msg MsgUnstake) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgUnstake) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// -----------------------------------------------------------------------------
// Route implements Msg
func (msg MsgHarvest) Route() string { return RouterKey }

// Type implements Msg
func (msg MsgHarvest) Type() string { return TypeMsgHarvest }

// ValidateBasic implements Msg
func (msg MsgHarvest) ValidateBasic() error {
	if err := ValidateAddress(msg.Sender); err != nil {
		return err
	}

	return ValidatePoolName(msg.PoolName)
}

// GetSignBytes implements Msg
func (msg MsgHarvest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgHarvest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}
