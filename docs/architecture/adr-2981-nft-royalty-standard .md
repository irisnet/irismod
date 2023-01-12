# ADR 2981: NFT Royalty Standard 

## Changelog

* 2022-10-19: Initial Draft

## Status

DRAFT 

## Abstract

This ADR describes the ability to signal the royalty amount to the NFT creator or rights holder each time an NFT is sold or resold. This is for NFT marketplaces that want to support ongoing funding for artists and other NFT creators. Royalty payments must be voluntary, because transfer mechanisms like `transfer()` include `NFT` transfers between wallets, and executing them does not always mean that a transaction has occurred. Markets and individuals implement this standard by using `royalInfo()` to retrieve royalty payment information that specifies how much to pay to which address for a given sale price. The exact mechanism for paying and notifying recipients will be defined in a future EIP. This ERC should be considered a minimal, fuel-efficient building block for further innovation in NFT royalties. This standard is fully compliant with `EIP-2981`.

## Context

At present, `x/nft` is based on `ERC-721`, which only contains basic functions. `irismod/moudles/nft` is an NFT application based on `x/nft`. However, in some cases, many developers may want certain features, such as using a common royalty interface, and users and artists also want to see a place to query or set royalties. And royalties are not only applied to `ERC-721`, if we have another token in the future, we can also use the current standard, so a common royalty protocol interface is beneficial to the whole ecology.

## Decision

We decided to implement a standardized method of retrieving royalty payment information for non-fungible tokens (NFTs) or other tokens to enable universal support for royalty payments for all NFT market and ecosystem participants.

We decided to extend the `irismod/moudles/nft` module to add the following features:

- Store the royalty information of the token
- Stores royalty information for the series the token is in (`denom` in `irismod/moudles/nft`)
- Expose the `Keeper` interface, which is convenient for the application room to refer to this module, which is used to set royalties, delete royalties and obtain royalties information;
- Expose the external `Message` interface for users to set the royalty information of the created token or series.
- Query the royalty information of the token/series

### Royalty

Royalty is a general method of obtaining royalty information, so no other factor should be included

- royaltyInfo: --  The structure in which the royalty information is kept.

### Types

#### RoyaltyInfo

We define a generic model `RoyaltyInfo` as follows.

```
message RoyaltyInfo {
    string address            = 1;
    string royalty_fraction   = 2;
};
```

- `address`：is the recipient address of `RoyaltyInfo`; required
- `royalty_fraction`：is the royalty numerator for `RoyaltyInfo` (royalty rate = numerator/denominator); required

### Storage

The royalty module needs to store two styles in total：`{class_id}/{key_name}`

e g：

1. `{class_id}/default --->RoyaltyInfo(Bytes)`：Stores default royalties set by class (global within class)
2. `{class_id}/{token_id} --->RoyaltyInfo(Bytes)`：Store royalty information of an NFT under a class

### `Keeper` Interface

```
type Keeper interface {

	// FeeDenominator returns the denominator of the fee
	FeeDenominator(ctx sdk.Context) (feeNumerator *big.Int)

	// RoyaltyInfo returns the royalty information of a token under a class
	RoyaltyInfo(ctx sdk.Context, classId string, tokenId string, salePrice *big.Int) (address string, royaltyAmount *big.Int)
	// DefaultRoyaltyInfo returns the default royalty information of a class
	DefaultRoyaltyInfo(ctx sdk.Context, classId string) (address string, feeNumerator *big.Int)
	// TokenRoyaltyInfo returns the royalty information of a token under a class
	TokenRoyaltyInfo(ctx sdk.Context, classId string, tokenId string) (address string, feeNumerator *big.Int)

	// SetDefaultRoyaltyInfo sets the default royalty information of a class
	SetDefaultRoyalty(ctx sdk.Context, classId string, receiver string, feeNumerator *big.Int) error
	// SetTokenRoyaltyInfo sets the royalty information of a token under a class
	SetTokenRoyalty(ctx sdk.Context, classId string, tokenId string, receiver string, feeNumerator *big.Int) error
	// DeleteTokenRoyaltyInfo deletes the royalty information of a token under a class
	ResetTokenRoyalty(ctx sdk.Context, classId string, tokenId string) error
	// DeleteDefaultRoyaltyInfo deletes the default royalty information of a class
	DeleteDefaultRoyalty(ctx sdk.Context, classId string) error
}
```

The approximate logic of royalty information calculation is as follows：

```
// RoyaltyInfo  returns the royalty information of a token under a class
func (k Keeper) RoyaltyInfo(ctx sdk.Context, classId string, tokenId string, salePrice *big.Int) (address string, royaltyAmount *big.Int) {
	address, feeNumerator := k.TokenRoyaltyInfo(ctx, classId, tokenId)

	if address == "" {
		address, feeNumerator = k.DefaultRoyaltyInfo(ctx, classId)
	}
	royaltyAmount = salePrice.Mul(salePrice, feeNumerator).Quo(salePrice, k.FeeDenominator(ctx))
	return
}
```



### `Msg` Service

```
service Msg {
    rpc SetDefaultRoyalty(MsgSetDefaultRoyalty) returns (MsgSetDefaultRoyaltyResponse);
    rpc SetTokenRoyalty(MsgSetTokenRoyalty) returns (MsgSetTokenRoyaltyResponse);
    rpc ResetTokenRoyalty(MsgResetTokenRoyalty) returns (MsgResetTokenRoyaltyResponse);
    rpc DeleteDefaultRoyalty(MsgDeleteDefaultRoyalty) returns (MsgDeleteDefaultRoyaltyResponse);
}
  
message MsgSetDefaultRoyalty {
    string class_id = 1;
    string receiver = 2;
    string fee_numerator = 3;
    
    string sender = 4;
}
message MsgSetDefaultRoyaltyResponse {}

message MsgSetTokenRoyalty {
    string class_id = 1;
    string token_id= 2;
    string receiver = 3;
    string fee_numerator = 4;
    
    string sender = 5;
}
message MsgSetTokenRoyaltyResponse {}

message MsgResetTokenRoyalty {
    string class_id = 1;
    string token_id= 2;
    
    string sender = 3;
}
message MsgResetTokenRoyaltyResponse {}

message MsgDeleteDefaultRoyalty {
    string class_id = 1;
    
    string sender = 2;
}
message MsgDeleteDefaultRoyaltyResponse {}



```

The implementation outline of the server is as follows：

```
type msgServer struct {
	k Keeper
}

func (m msgServer) SetDefaultRoyalty(goCtx context.Context, msg *types.MsgSetDefaultRoyalty) (*types.MsgSetDefaultRoyaltyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// check current ownership
	assertEqual(msg.Sender, m.k.GetOwnerFromClass(msg.ClassId))

	// set default royalty info for a classId
	m.k.SetDefaultRoyalty(ctx, msg.ClassId)
	return &types.MsgSetDefaultRoyaltyResponse, nil
}

func (m msgServer) SetTokenRoyalty(goCtx context.Context, msg *types.MsgSetTokenRoyalty) (*types.MsgSetTokenRoyaltyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// check current ownership
	assertEqual(msg.Sender, m.k.GetOwnerFromToken(msg.ClassId, msg.TokenId))

	// set royalty info for a token
	m.k.SetTokenRoyalty(ctx, msg.ClassId, msg.TokenId)
	return &types.MsgSetTokenRoyaltyResponse, nil
}

func (m msgServer) ResetTokenRoyalty(goCtx context.Context, msg *types.MsgResetTokenRoyalty) (*types.MsgResetTokenRoyaltyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// check current ownership
	assertEqual(msg.Sender, m.k.GetOwnerFromToken(msg.ClassId, msg.TokenId))

	// reset royalty info for a token
	m.k.ResetTokenRoyalty(ctx, msg.ClassId, msg.TokenId)
	return &types.MsgResetTokenRoyaltyResponse, nil
}

func (m msgServer) DeleteDefaultRoyalty(goCtx context.Context, msg *types.MsgDeleteDefaultRoyalty) (*types.MsgDeleteDefaultRoyaltyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// check current ownership
	assertEqual(msg.Sender, m.k.GetOwnerFromClass(msg.ClassId))

	// delete default royalty info for a class
	m.k.DeleteDefaultRoyalty(ctx, msg.ClassId, msg.TokenId)
	return &types.MsgDeleteDefaultRoyaltyResponse, nil
}

```



query method：

```

service Query {
    rpc FeeDenominator(MsgFeeDenominatorRequest) returns (MsgFeeDenominatorResponse);
    rpc RoyaltyInfo(MsgRoyaltyInfoRequest) returns (MsgRoyaltyInfoResponse);
    rpc DefaultRoyaltyInfo(MsgDefaultRoyaltyInfoRequest) returns (MsgDefaultRoyaltyInfoResponse);
    rpc TokenRoyaltyInfo(MsgTokenRoyaltyInfoRequest) returns (MsgTokenRoyaltyInfoResponse);
}

message MsgFeeDenominatorRequest {
}
message MsgFeeDenominatorResponse {
    uint64 royalty_fraction = 1;
}

message MsgRoyaltyInfoRequest {
    string class = 1;
    string token_id = 2;
    uint64 sale_price = 3;
}
message MsgRoyaltyInfoResponse {
    string reciver = 1;
    uint64 royalty_amount = 2;
}

message MsgDefaultRoyaltyInfoRequest {
    string class_id = 1;
}
message MsgDefaultRoyaltyInfoResponse {
    string reciver = 1;
    uint64 royalty_fraction = 2;
}

message MsgTokenRoyaltyInfoRequest {
    string class_id = 1;
    string token_id = 2;
}
message MsgTokenRoyaltyInfoResponse {
    string reciver = 1;
    uint64 royalty_fraction = 2;
}



```



## Consequences

### Backwards Compatibility

No backwards incompatibility.

### Positive

- NFT Royalty Standard on Cosmos Hub.

### Negative

None

### Neutral

None

## Further Discussions





## References

