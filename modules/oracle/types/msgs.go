package types

import (
	"regexp"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	MaxLatestHistory    = 100
	MaxNameLen          = 70
	MaxAggregateFuncLen = 10
	MaxValueJsonPath    = 70
	MaxDescriptionLen   = 200

	TypeMsgCreateFeed = "create_feed" // type for MsgCreateFeed
	TypeMsgStartFeed  = "start_feed"  // type for MsgStartFeed
	TypeMsgPauseFeed  = "pause_feed"  // type for MsgPauseFeed
	TypeMsgEditFeed   = "edit_feed"   // type for MsgEditFeed

	DoNotModify = "do-not-modify"
)

var (
	_ sdk.Msg = &MsgCreateFeed{}
	_ sdk.Msg = &MsgStartFeed{}
	_ sdk.Msg = &MsgPauseFeed{}
	_ sdk.Msg = &MsgEditFeed{}

	// the feed/service name only accepts alphanumeric characters, _ and -
	regPlainText = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_-]*$`)
)

// ______________________________________________________________________

// Route implements Msg.
func (msg MsgCreateFeed) Route() string {
	return RouterKey
}

// Type implements Msg.
func (msg MsgCreateFeed) Type() string {
	return TypeMsgCreateFeed
}

// ValidateBasic implements Msg.
func (msg MsgCreateFeed) ValidateBasic() error {
	if err := ValidateFeedName(msg.FeedName); err != nil {
		return err
	}

	if err := ValidateDescription(msg.Description); err != nil {
		return err
	}

	if err := validateServiceName(msg.ServiceName); err != nil {
		return err
	}

	if err := ValidateLatestHistory(msg.LatestHistory); err != nil {
		return err
	}

	if err := validateTimeout(msg.Timeout, msg.RepeatedFrequency); err != nil {
		return err
	}
	if len(msg.Providers) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "providers missing")
	}

	if err := ValidateAggregateFunc(msg.AggregateFunc); err != nil {
		return err
	}

	if err := ValidateValueJSONPath(msg.ValueJsonPath); err != nil {
		return err
	}

	if !msg.ServiceFeeCap.IsValid() {
		return sdkerrors.Wrapf(ErrInvalidServiceFeeCap, msg.ServiceFeeCap.String())
	}

	if err := ValidateCreator(msg.Creator); err != nil {
		return err
	}

	return validateResponseThreshold(msg.ResponseThreshold, len(msg.Providers))
}

// Normalize return a string with spaces removed and lowercase
func (msg *MsgCreateFeed) Normalize() *MsgCreateFeed {
	msg.FeedName = strings.TrimSpace(msg.FeedName)
	msg.ServiceName = strings.TrimSpace(msg.ServiceName)
	msg.Input = strings.TrimSpace(msg.Input)
	msg.AggregateFunc = strings.ToLower(strings.TrimSpace(msg.AggregateFunc))
	msg.ValueJsonPath = strings.TrimSpace(msg.ValueJsonPath)
	return msg
}

// GetSignBytes implements Msg.
func (msg MsgCreateFeed) GetSignBytes() []byte {
	if len(msg.Providers) == 0 {
		msg.Providers = nil
	}
	if msg.ServiceFeeCap.Empty() {
		msg.ServiceFeeCap = nil
	}
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners implements Msg.
func (msg MsgCreateFeed) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// _____________________________________________________________________

// Route implements Msg.
func (msg MsgStartFeed) Route() string {
	return RouterKey
}

// Type implements Msg.
func (msg MsgStartFeed) Type() string {
	return TypeMsgStartFeed
}

// ValidateBasic implements Msg.
func (msg MsgStartFeed) ValidateBasic() error {
	if err := ValidateCreator(msg.Creator); err != nil {
		return err
	}
	return ValidateFeedName(msg.FeedName)
}

// Normalize return a string with spaces removed and lowercase
func (msg *MsgStartFeed) Normalize() *MsgStartFeed {
	msg.FeedName = strings.TrimSpace(msg.FeedName)
	return msg
}

// GetSignBytes implements Msg.
func (msg MsgStartFeed) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners implements Msg.
func (msg MsgStartFeed) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// ______________________________________________________________________

// Route implements Msg.
func (msg MsgPauseFeed) Route() string {
	return RouterKey
}

// Type implements Msg.
func (msg MsgPauseFeed) Type() string {
	return TypeMsgPauseFeed
}

// ValidateBasic implements Msg.
func (msg MsgPauseFeed) ValidateBasic() error {
	if err := ValidateCreator(msg.Creator); err != nil {
		return err
	}
	return ValidateFeedName(msg.FeedName)
}

// Normalize return a string with spaces removed and lowercase
func (msg *MsgPauseFeed) Normalize() *MsgPauseFeed {
	msg.FeedName = strings.TrimSpace(msg.FeedName)
	return msg
}

// GetSignBytes implements Msg.
func (msg MsgPauseFeed) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners implements Msg.
func (msg MsgPauseFeed) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// ______________________________________________________________________

// Route implements Msg.
func (msg MsgEditFeed) Route() string {
	return RouterKey
}

// Type implements Msg.
func (msg MsgEditFeed) Type() string {
	return TypeMsgEditFeed
}

// ValidateBasic implements Msg.
func (msg MsgEditFeed) ValidateBasic() error {
	if err := ValidateFeedName(msg.FeedName); err != nil {
		return err
	}
	if err := ValidateDescription(msg.Description); err != nil {
		return err
	}
	if msg.LatestHistory != 0 {
		if err := ValidateLatestHistory(msg.LatestHistory); err != nil {
			return err
		}
	}
	if msg.ServiceFeeCap != nil && !msg.ServiceFeeCap.IsValid() {
		return sdkerrors.Wrapf(ErrInvalidServiceFeeCap, msg.ServiceFeeCap.String())
	}
	if msg.Timeout != 0 && msg.RepeatedFrequency != 0 {
		if err := validateTimeout(msg.Timeout, msg.RepeatedFrequency); err != nil {
			return err
		}
	}
	if msg.ResponseThreshold != 0 {
		if err := validateResponseThreshold(msg.ResponseThreshold, len(msg.Providers)); err != nil {
			return err
		}
	}
	return ValidateCreator(msg.Creator)
}

// Normalize return a string with spaces removed and lowercase
func (msg *MsgEditFeed) Normalize() *MsgEditFeed {
	msg.FeedName = strings.TrimSpace(msg.FeedName)
	return msg
}

// GetSignBytes implements Msg.
func (msg MsgEditFeed) GetSignBytes() []byte {
	if len(msg.Providers) == 0 {
		msg.Providers = nil
	}
	if msg.ServiceFeeCap.Empty() {
		msg.ServiceFeeCap = nil
	}
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners implements Msg.
func (msg MsgEditFeed) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// ValidateFeedName verify that the feedName is legal
func ValidateFeedName(feedName string) error {
	feedName = strings.TrimSpace(feedName)
	if len(feedName) == 0 || len(feedName) > MaxNameLen {
		return sdkerrors.Wrap(ErrInvalidFeedName, feedName)
	}
	if !regPlainText.MatchString(feedName) {
		return sdkerrors.Wrap(ErrInvalidFeedName, feedName)
	}
	return nil
}

// ValidateDescription verify that the desc is legal
func ValidateDescription(desc string) error {
	desc = strings.TrimSpace(desc)
	if len(desc) > MaxDescriptionLen {
		return sdkerrors.Wrap(ErrInvalidDescription, desc)
	}
	return nil
}

// ValidateAggregateFunc verify that the aggregateFunc is legal
func ValidateAggregateFunc(aggregateFunc string) error {
	aggregateFunc = strings.TrimSpace(aggregateFunc)
	if len(aggregateFunc) == 0 || len(aggregateFunc) > MaxAggregateFuncLen {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "aggregate func must between [1, %d], got: %d", MaxAggregateFuncLen, len(aggregateFunc))
	}
	if _, err := GetAggregateFunc(aggregateFunc); err != nil {
		return err
	}
	return nil
}

// ValidateValueJSONPath verify that the valueJsonPath is legal
func ValidateValueJSONPath(valueJSONPath string) error {
	valueJSONPath = strings.ToLower(strings.TrimSpace(valueJSONPath))
	if len(valueJSONPath) == 0 || len(valueJSONPath) > MaxValueJsonPath {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "the length of valueJson path func must less than %d, got: %d", MaxAggregateFuncLen, len(valueJSONPath))
	}
	return nil
}

// ValidateLatestHistory verify that the latestHistory is legal
func ValidateLatestHistory(latestHistory uint64) error {
	if latestHistory < 1 || latestHistory > MaxLatestHistory {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "latest history is invalid, should be between 1 and %d", MaxLatestHistory)
	}
	return nil
}

// ValidateCreator verify that the creator is legal
func ValidateCreator(creator string) error {
	if _, err := sdk.AccAddressFromBech32(creator); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator")
	}
	return nil
}

func validateServiceName(serviceName string) error {
	serviceName = strings.TrimSpace(serviceName)
	if len(serviceName) == 0 || len(serviceName) > MaxNameLen {
		return sdkerrors.Wrapf(ErrInvalidServiceName, serviceName)
	}
	if !regPlainText.MatchString(serviceName) {
		return sdkerrors.Wrapf(ErrInvalidServiceName, serviceName)
	}
	return nil
}

func validateResponseThreshold(responseThreshold uint32, maxCnt int) error {
	if (maxCnt != 0 && int(responseThreshold) > maxCnt) || responseThreshold < 1 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "response threshold should be between 1 and %d", maxCnt)
	}
	return nil
}

func validateTimeout(timeout int64, frequency uint64) error {
	if frequency < uint64(timeout) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "timeout [%d] should be no more than frequency [%d]", timeout, frequency)
	}
	return nil
}

func IsModified(target string) bool {
	target = strings.TrimSpace(target)
	return target != DoNotModify
}
