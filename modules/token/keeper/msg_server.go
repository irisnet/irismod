package keeper

import (
	"context"
	"encoding/hex"
	"math/big"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"

	"github.com/irisnet/irismod/modules/token/types"
	v1 "github.com/irisnet/irismod/modules/token/types/v1"
)

type msgServer struct {
	k Keeper
}

var _ v1.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the token MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) v1.MsgServer {
	return &msgServer{k: keeper}
}

func (m msgServer) IssueToken(
	goCtx context.Context,
	msg *v1.MsgIssueToken,
) (*v1.MsgIssueTokenResponse, error) {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	if m.k.blockedAddrs[msg.Owner] {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "%s is a module account", msg.Owner)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// handle fee for token
	if err := m.k.DeductIssueTokenFee(ctx, owner, msg.Symbol); err != nil {
		return nil, err
	}

	if err := m.k.IssueToken(
		ctx, msg.Symbol, msg.Name, msg.MinUnit, msg.Scale,
		msg.InitialSupply, msg.MaxSupply, msg.Mintable, owner,
	); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeIssueToken,
			sdk.NewAttribute(types.AttributeKeySymbol, msg.Symbol),
			sdk.NewAttribute(types.AttributeKeyCreator, msg.Owner),
		),
	})

	return &v1.MsgIssueTokenResponse{}, nil
}

func (m msgServer) EditToken(
	goCtx context.Context,
	msg *v1.MsgEditToken,
) (*v1.MsgEditTokenResponse, error) {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.k.EditToken(
		ctx, msg.Symbol, msg.Name,
		msg.MaxSupply, msg.Mintable, owner,
	); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeEditToken,
			sdk.NewAttribute(types.AttributeKeySymbol, msg.Symbol),
			sdk.NewAttribute(types.AttributeKeyOwner, msg.Owner),
		),
	})

	return &v1.MsgEditTokenResponse{}, nil
}

func (m msgServer) MintToken(
	goCtx context.Context,
	msg *v1.MsgMintToken,
) (*v1.MsgMintTokenResponse, error) {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	var recipient sdk.AccAddress

	if len(msg.To) != 0 {
		recipient, err = sdk.AccAddressFromBech32(msg.To)
		if err != nil {
			return nil, err
		}
	} else {
		recipient = owner
	}

	if m.k.blockedAddrs[recipient.String()] {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "%s is a module account", recipient)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	symbol, err := m.k.getSymbolByMinUnit(ctx, msg.Coin.Denom)
	if err != nil {
		return nil, err
	}

	if err := m.k.DeductMintTokenFee(ctx, owner, symbol); err != nil {
		return nil, err
	}

	if err := m.k.MintToken(ctx, msg.Coin, recipient, owner); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeMintToken,
			sdk.NewAttribute(types.AttributeKeyAmount, msg.Coin.String()),
			sdk.NewAttribute(types.AttributeKeyRecipient, recipient.String()),
		),
	})

	return &v1.MsgMintTokenResponse{}, nil
}

func (m msgServer) BurnToken(
	goCtx context.Context,
	msg *v1.MsgBurnToken,
) (*v1.MsgBurnTokenResponse, error) {
	owner, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.k.BurnToken(ctx, msg.Coin, owner); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeBurnToken,
			sdk.NewAttribute(types.AttributeKeyAmount, msg.Coin.String()),
			sdk.NewAttribute(types.AttributeKeySender, msg.Sender),
		),
	})

	return &v1.MsgBurnTokenResponse{}, nil
}

func (m msgServer) TransferTokenOwner(
	goCtx context.Context,
	msg *v1.MsgTransferTokenOwner,
) (*v1.MsgTransferTokenOwnerResponse, error) {
	srcOwner, err := sdk.AccAddressFromBech32(msg.SrcOwner)
	if err != nil {
		return nil, err
	}

	dstOwner, err := sdk.AccAddressFromBech32(msg.DstOwner)
	if err != nil {
		return nil, err
	}

	if m.k.blockedAddrs[msg.DstOwner] {
		return nil, errorsmod.Wrapf(
			sdkerrors.ErrUnauthorized,
			"%s is a module account",
			msg.DstOwner,
		)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.k.TransferTokenOwner(ctx, msg.Symbol, srcOwner, dstOwner); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeTransferTokenOwner,
			sdk.NewAttribute(types.AttributeKeySymbol, msg.Symbol),
			sdk.NewAttribute(types.AttributeKeyOwner, msg.SrcOwner),
			sdk.NewAttribute(types.AttributeKeyDstOwner, msg.DstOwner),
		),
	})

	return &v1.MsgTransferTokenOwnerResponse{}, nil
}

func (m msgServer) SwapFeeToken(
	goCtx context.Context,
	msg *v1.MsgSwapFeeToken,
) (*v1.MsgSwapFeeTokenResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	var recipient sdk.AccAddress
	if len(msg.Recipient) > 0 {
		recipient, err = sdk.AccAddressFromBech32(msg.Recipient)
		if err != nil {
			return nil, err
		}

		if m.k.blockedAddrs[msg.Recipient] {
			return nil, errorsmod.Wrapf(
				sdkerrors.ErrUnauthorized,
				"%s is a module account",
				recipient,
			)
		}
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	feePaid, feeGot, err := m.k.SwapFeeToken(ctx, msg.FeePaid, sender, recipient)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSwapFeeToken,
			sdk.NewAttribute(types.AttributeKeySender, msg.Sender),
			sdk.NewAttribute(types.AttributeKeyRecipient, msg.Recipient),
			sdk.NewAttribute(types.AttributeKeyFeePaid, feePaid.String()),
			sdk.NewAttribute(types.AttributeKeyFeeGot, feeGot.String()),
		),
	})

	return &v1.MsgSwapFeeTokenResponse{
		FeeGot: feeGot,
	}, nil
}

func (m msgServer) UpdateParams(
	goCtx context.Context,
	msg *v1.MsgUpdateParams,
) (*v1.MsgUpdateParamsResponse, error) {
	if m.k.authority != msg.Authority {
		return nil, errorsmod.Wrapf(
			sdkerrors.ErrUnauthorized,
			"invalid authority; expected %s, got %s",
			m.k.authority,
			msg.Authority,
		)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.k.SetParams(ctx, msg.Params); err != nil {
		return nil, err
	}
	return &v1.MsgUpdateParamsResponse{}, nil
}

func (m msgServer) SwapToERC20(goCtx context.Context, msg *v1.MsgSwapToERC20) (*v1.MsgSwapToERC20Response, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	receiver := common.BytesToAddress(sender.Bytes())
	if len(msg.Receiver) > 0 {
		bz, err := hex.DecodeString(msg.Receiver)
		if err != nil {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "expecting a hex address of 0x, got %s", msg.Receiver)
		}
		receiver = common.BytesToAddress(bz)
	}

	if err := m.k.SwapToERC20(ctx, sender, receiver, msg.Amount); err != nil {
		return nil, err
	}
	return &v1.MsgSwapToERC20Response{}, nil
}

func (m msgServer) SwapFromERC20(goCtx context.Context, msg *v1.MsgSwapFromERC20) (*v1.MsgSwapFromERC20Response, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sender, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return nil, err
	}

	receiver := sender
	if len(msg.Receiver) > 0 {
		receiver, err = sdk.AccAddressFromBech32(msg.Receiver)
		if err != nil {
			return nil, err
		}
	}

	bz, err := hex.DecodeString(msg.Receiver)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "expecting a hex address of 0x, got %s", msg.Receiver)
	}
	contract := common.BytesToAddress(bz)

	if err := m.k.SwapFromERC20(ctx, common.BytesToAddress(sender.Bytes()), receiver, contract, big.NewInt(msg.Amount)); err != nil {
		return nil, err
	}
	return &v1.MsgSwapFromERC20Response{}, nil
}

func (m msgServer) DeployERC20(goCtx context.Context, msg *v1.MsgDeployERC20) (*v1.MsgDeployERC20Response, error) {
	var (
		ctx     = sdk.UnwrapSDKContext(goCtx)
		name    = msg.Name
		symbol  = msg.Symbol
		scale   = msg.Scale
		minUnit = msg.MinUnit
		token   v1.Token
		err     error
	)

	if !m.k.HasMinUint(ctx, msg.MinUnit) {
		if !m.k.ics20Keeper.HasTrace(ctx, msg.MinUnit) {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "token: %s not exist", msg.MinUnit)
		}
		token = v1.Token{
			Symbol:   symbol,
			Name:     name,
			Scale:    scale,
			MinUnit:  msg.MinUnit,
			Mintable: true,
			Owner:    string(m.k.accountKeeper.GetModuleAddress(types.ModuleName)),
		}
	} else {
		token, err = m.k.getTokenByMinUnit(ctx, msg.MinUnit)
		if err != nil {
			return nil, err
		}
		if len(token.Contract) > 0 {
			return nil, errorsmod.Wrapf(types.ErrERC20AlreadyExists, "token: %s already deployed erc20 contract: %s", token.Symbol, token.Contract)
		}
		name = token.Name
		symbol = token.Symbol
		scale = token.Scale
		minUnit = token.MinUnit
	}

	contractAddr, err := m.k.DeployERC20(ctx, name, symbol, minUnit, int8(scale))
	if err != nil {
		return nil, err
	}
	token.Contract = contractAddr.String()
	m.k.upsertToken(ctx, token)
	return &v1.MsgDeployERC20Response{}, nil
}
