# ADR 4907: Rental NFT

## Changelog

* 2022-10-19: Initial Draft

## Status

DRAFT 

## Abstract

This ADR is an extension to `x/nft`. It proposes an additional role (user) that can be granted to an address, and when that role is automatically revoked (expired). User roles represent permissions to "use" NFTs, but not the ability to transfer or set up users. This standard is fully compliant with `EIP-4907`

## Context

At present, `x/nft` only implements the basic interface of `EIP-721`. In some cases, many developers may need certain functions. For example, the current NFT has high artistic value, and of course the price will be high. At this time, it is meaningful for the entire ecology to have a module with the ability to lease NFT.

## Decision

We extend based on `irismod/modules/nft`, which extends the following functions:

- Store rental information for NFTs
- Expose the `Keeper` interface for writing modules for renting NFTs.
- Expose the external `Message` interface for users to rent out the right to use the NFT they own.
- Query NFT rental information.

### Rental

Rental is an extension of NFT, mainly to allow NFT to support rental.

- rentalInfo: --  The structure in which the Rental information is kept.

### Types

#### RentalInfo

```
message RentalInfo {
    string renter = 1;
    uint64 expires = 2;
}
```

- `renter`：is the renter address of `RentalInfo`; required
- `expires`：is the time of the rental of `RentalInfo`; required

### Storage

A total of one style needs to be stored for royalties:

1. `{class_id}/{nft_id} ---->RentalInfo (bytes) `：Store the information that an NFT is leased

### `Keeper` Interface

```
type Keeper interface {
	UserOf(ctx sdk.Context, classId, nftId string) (renter string)
	UserExpires(ctx sdk.Context, classId, nftId string) (expire uint64)

	SetUser(ctx sdk.Context, classId, nftId string, renter string, expire uint64)

	// determines whether the NFT is being rented
	HaveUser(ctx sdk.Context, classId, nftId string) bool

	// Delete rental information after expiration
	// should be call in EndBlock
	DeleteUser(ctx sdk.Context, classId, nftId string)
}
```



### `Msg` Service

```
service Msg {
    rpc SetUser(MsgSetUser) returns (MsgSetUserResponse) {}
}

message MsgSetUser {
    string class_id = 1;
    string nft_id = 2;
    string renter = 3;
    uint64 expire = 4;

    string sender = 5;
}

message MsgSetUserResponse {}

```

The implementation outline of the server is as follows:

```
type msgServer struct {
	k Keeper
}

func (m msgServer) SetUser(goCtx context.Context, msg *types.MsgSetUser) (*types.MsgSetUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check current ownership
	assertEqual(msg.Sender, m.k.GetOwnerFromClass(msg.Module, msg.ClassId))

	// check whether the nft is not already rented
	assertEqual(false, m.k.HaveUser(ctx, msg.classId, msg.NftId))

	m.k.SetUser(ctx, msg.classId, msg.NftId, msg.Renter, uint64(msg.Expire))
	return &types.MsgSetUserResponse{}, nil
}

```

Query method:

```
service Query {
    rpc UserOf(MsgUserOfRequest) returns (MsgUserOfResponse);
    rpc UserExpires(MsgUserExpiresRequest) returns (MsgUserExpiresResponse);
    rpc HaveUser(MsgHaveUserRequest) returns (MsgHaveUserResponse);
}

message MsgUserOfRequest {
    string class_id = 1;
    string nft_id = 2;
}

message MsgUserOfResponse {
    string renter = 1;
}

message MsgUserExpiresRequest {
    string class_id = 1;
    string nft_id = 2;
}

message MsgUserExpiresResponse {
    uint64 expires = 1;
}

message MsgHaveUserRequest {
    string class_id = 1;
    string nft_id = 2;
}

message MsgHaveUserResponse {
    bool had_renter = 1;
}

```



## Consequences

### Backwards Compatibility

No backwards incompatibility.

### Positive

- NFT Rental Standard on Cosmos Hub.

### Negative

None

### Neutral

None

## Further Discussions





## References

