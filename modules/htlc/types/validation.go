package types

import (
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// SecretLength is the length for the secret in hex string
	SecretLength = 64
	// HashLockLength is the length for the hash lock in hex string
	HTLCIDLength = 64
	// MaxLengthForAddressOnOtherChain is the maximum length for the address on other chains
	MaxLengthForAddressOnOtherChain = 128
	// MinTimeLock is the minimum time span for HTLC in blocks
	MinTimeLock = 50
	// MaxTimeLock is the maximum time span for HTLC in blocks
	MaxTimeLock = 25480
)

// ValidateReceiverOnOtherChain verifies if the receiver on the other chain is legal
func ValidateReceiverOnOtherChain(receiverOnOtherChain string) error {
	if len(receiverOnOtherChain) > MaxLengthForAddressOnOtherChain {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "length of the receiver on other chain must be between [0,%d]", MaxLengthForAddressOnOtherChain)
	}
	return nil
}

// ValidateSenderOnOtherChain verifies if the receiver on the other chain is legal
func ValidateSenderOnOtherChain(senderOnOtherChain string) error {
	if len(senderOnOtherChain) > MaxLengthForAddressOnOtherChain {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "length of the sender on other chain must be between [0,%d]", MaxLengthForAddressOnOtherChain)
	}
	return nil
}

// ValidateAmount verifies whether the given amount is legal
func ValidateAmount(amount sdk.Coins) error {
	if !(amount.IsValid() && amount.IsAllPositive()) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "the transferred amount must be valid")
	}
	return nil
}

// ValidateID verifies whether the given ID lock is legal
func ValidateID(id string) error {
	if len(id) != HTLCIDLength {
		return sdkerrors.Wrapf(ErrInvalidID, "length of the htlc id must be %d", HTLCIDLength)
	}

	if _, err := hex.DecodeString(id); err != nil {
		return sdkerrors.Wrapf(ErrInvalidID, "id must be a hex encoded string")
	}
	return nil
}

// ValidateTimeLock verifies whether the given time lock is legal
func ValidateTimeLock(timeLock uint64) error {
	if timeLock < MinTimeLock || timeLock > MaxTimeLock {
		return sdkerrors.Wrapf(ErrInvalidTimeLock, "the time lock must be between [%d,%d]", MinTimeLock, MaxTimeLock)
	}
	return nil
}

// ValidateSecret verifies whether the given secret is legal
func ValidateSecret(secret string) error {
	if len(secret) != SecretLength {
		return sdkerrors.Wrapf(ErrInvalidSecret, "length of the secret must be %d", SecretLength)
	}

	if _, err := hex.DecodeString(secret); err != nil {
		return sdkerrors.Wrapf(ErrInvalidSecret, "secret must be a hex encoded string")
	}
	return nil
}
