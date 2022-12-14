# ADR 3525: SFT Module

## Changelog

- 2022-10-19: Initial Draft

## Status

DRAFT

## Abstract

This ADR defines the `x/sft` module which implements the **core** functionality of SFT. It also specifies the **rules** for utilizing the `x/sft` module as the base for building SFT applications.

## Context

Semi-fungible Token, also known as SFT, has emerged as NFT compatible token while allowing more flexibility than NFT. It assumes that tokens with the same **slot** (an abstraction of the quantifiable characteristic of a group of tokens) can transfer **value** (the quantitative value of that characteristic) between themselves.

SFT is still in its early days. A typical application scenario for SFT is as a financial instrument. [Solv.Finance](), the [ERC-3525](https://eips.ethereum.org/EIPS/eip-3525) proposer, uses SFT in their Voucher product. With the use of SFT, [Voucher](https://whitepaper.sftlabs.io/SFT%20Whitepaper.pdf) can be divided into hundreds of tokens with different value but subjected to the same set of rules for the underlying asset. 

## Decision

We created an `x/sft` module, which contains the following functionality:

- Unified definition of `Class`, `Slot`, and `SFT`
- Store `Slot` and `Value` of SFT.
- Expose the `Keeper` interface for composing modules to transfer value and create slots.
- Query `Slot`, `Value`, and related info.

The module is intended to be a base module for SFT application logic and must not be used standalone. As SFT is compatible with NFT, it is possible to make use of `x/nft` module facilities in the implementation. Considering this flexibility, we make applications choose their preferred ways to make full implementation, but they **must** adhere to the following **rules**:

- Functionality that is compatible with NFT **should** reuse methods of the `x/nft` module.
- Functionality that is dedicated to SFT **must** call relevant methods of the `x/sft` module.

### Types

We propose three main types:

- `Class` describes SFT class, which can be thought of as a smart contract address.
- `Slot` is an abstraction of SFT fungible characteristic. SFTs with the same slot can transfer value to each other.
- `SFT` is Semi-fungible Token that is compatible with NFT. Each SFT has one slot and a transferable value.

#### Class

SFT Class is comparable to an ERC-3525 smart contract, which creates and manages a collection of SFTs. Its definition is the same as NFT Class's in `x/nft`.

```protobuf
message Class {  
  string id = 1;
  string name = 2;  
  string symbol = 3;  
  string description = 4;  
  string uri = 5;  
  string uri_hash = 6;  
  google.protobuf.Any Data = 10;  
}
```

- `id`  is a **required** alphanumeric identifier of the SFT class.
- `name`  is an **optional** descriptive name of the SFT class.
- `symbol` is an **optional** symbol for the SFT class.
- `description` is an **optional** detailed description of the SFT class.
- `uri` is an **optional** URI for the class metadata stored off-chain.
- `uri_hash` is an **optional** hash of the document pointed to by `uri`.
- `data` is an **optional** app-specific metadata of the class.

#### Slot

Value transfer currently is **class relevant**, which means SFTs with the same slot but belonging to different classes are not value transferable.

```protobuf
message Slot {  
  string id = 1;  
  string name = 2;  
  string description = 3;  
  string uri = 4;  
  string uri_hash = 5;  
  google.protobuf.Any data = 10;  
}
```

- `id` is a **required** and module scope unique identifier of a slot.
- `name` is an **optional** descriptive name of a slot.
- `description` is an **optional** detailed description.
- `uri` is an **optional** URI for the **slot metadata** stored off-chain.
- `uri_hash` is an **optional** hash of the document pointed to by `uri`.
- `data` is an **optional** but very recommended app-specific metadata of a slot for calculating its `id`.

`slot_id` is a module-scoped and we maintain a global current slot `sequence_id`. Each time a new slot is created,  `sequence_id` increases by one and is assigned to `slot_id`.

#### SFT

SFT is defined as a triad consisting of `base`, `slot`, and `value`:

- `base` is a **required** BaseSFT of the SFT.
- `slot_id` is a **required** identifier of an existing slot.
- `value` is a **required** transferable value of the SFT.

```protobuf
message SFT {  
  BaseSFT base = 1;  
  string slot_id =  2;  
  uint64 value = 3;
}
```

BaseSFT has the same definition as NFT in `x/nft` and can be viewed as the Non-fungible part of an SFT:

- `class_id` is the **required** identifier of the SFT class to which the SFT belongs;
- `id` is a **required** identifier of the SFT which is unique in the class scope. It follows the same generation rule as NFT.
- `uri` is a **required** URI for the SFT metadata stored off-chain.
- `uri_hash` is an **optional** hash of the document pointed by `uri`.
- `data` is an **optional** app-specific data of the SFT.

```protobuf
message BaseSFT {  
  string class_id = 1;  
  string id = 2;  
  string uri = 3;  
  string uri_hash = 4;  
  google.protobuf.Any data = 10;  
}
```

### `Keeper` Interface

#### `x/sft`

The `Keeper` interface is expected to be used by the SFT application to make full SFT implementation.

```go
type Keeper interface {  
	TransferValue(ctx sdk.Context, classId string, srcId string, toId string, val uint64)

	SaveSlot(ctx sdk.Context, slot Slot)

	GetSlot(ctx sdk.Context, slotId string) Slot
	GetSlots(ctx sdk.Context) []Slot
	GetSFTsOfSlot(ctx sdk.Context, classId string, slotId string) []SFT
	
	GetValue(ctx sdk.Context, classId string, sftId string) uint64
	GetSlotSFT(ctx sdk.Context, classId string, sftId string) Slot
}
```

#### `app/sft`

We recommend the SFT application implement the following interface. It has almost the same interface definition as `x/nft`, except for considering `slot` and `value` in implementation.

```go
type Keeper interface {
	SaveClass(ctx sdk.Context, class Class)
	UpdateClass(ctx sdk.Context, class Class)

	Mint(ctx sdk.Context, sft SFT, receiver sdk.AccAddress)
	Burn(ctx sdk.Context, classId string, sftId string)
	Update(ctx sdk.Context, sft SFT)
	Transfer(ctx sdk.Context, sft SFT)

	GetClass(ctx sdk.Context, classId string) Class
	GetClasses(ctx sdk.Context) []Class
	
	GetSFT(ctx sdk.Context, classId string, sftId string) SFT
	GetSFTsOfClassByOwner(ctx sdk.Context, classId string, owner sdk.AccAddress) []SFT
	GetSFTsOfClass(ctx sdk.Context, classId string) []SFT
	
	GetOwner(ctx sdk.Context, classId string, sftId string) sdk.AccAddress
	GetBalance(ctx sdk.Context, classId string, owner sdk.AccAddress) uint64
	GetTotalSupply(ctx sdk.Context, classId string) uint64
}
```

#### Example

The application can use `x/nft` to implement functionality such as `Mint`.

```go
// The following code is only used as a demo
func NewKeeper(cdc codec.Codec,  
   storeKey storetypes.StoreKey,  
   ak       AccountKeeper,  
   bk       BankKeeper,  
) Keeper {  
   return Keeper{  
      storeKey: storeKey,
      cdc:      cdc,  
      sk:       sftkeeper.NewKeeper(storeKey, cdc),
      nk:       nftkeeper.NewKeeper(storeKey, cdc, ak, bk),  
   }  
}

func(k Keeper) Mint(ctx sdk.Context, sft SFT, recipient sdk.AccAddress) {
	// baseSFT converted to NFT type.
	baseSFT := nft.NFT(sft.BaseSFT)
	if err := k.nk.Mint(ctx sdk.Context, baseSFT, recipient); err != nil {
		return err
	}
	k.sk.setSlot(ctx, sft.BaseSFT.ClassId, sft.BaseSFT.Id, sft.Slot)
	k.sk.setValue(ctx, sft.BaseSFT.ClassId, sft.BaseSFT.Id, sft.Value)
}
```

Here is another way to implement it, which needs extra work and adds more redundancy to the code repo.

```go
// The following code is only used as a demo
func NewKeeper(cdc codec.Codec,  
   storeKey storetypes.StoreKey,  
   ak       AccountKeeper,  
   bk       BankKeeper,  
) Keeper {  
   return Keeper{  
      storeKey: storeKey,
      cdc:      cdc,  
      ak:       ak,
      bk:       bk,
      sk:       sftkeeper.NewKeeper(storeKey, cdc),
   }  
}

func(k Keeper) Mint(ctx sdk.Context, sft SFT, recipient sdk.AccAddress) {
	// application implemented method
	k.setBaseSFT(ctx sdk.Context, sft.BaseSFT)
	// application implemented method 
	k.setOwner(ctx sdk.Context, sft.BaseSFT.ClassId, sft.BaseSFT.Id, recipient)
	k.sk.setSlot(ctx, sft.BaseSFT.ClassId, sft.BaseSFT.Id, sft.Slot)
	k.sk.setValue(ctx, sft.BaseSFT.ClassId, sft.BaseSFT.Id, sft.Value)
}
```

### `Msg` Service

#### `x/sft`

```protobuf
// Msg defines the sft Msg service.
service Msg {
  // TransferValueSFT defines a method to transfer value between sfts that are of the same slot.
  rpc TransferValueSFT(MsgTransferValueSFT) returns (MsgTransferValueSFTResponse);

  // CreateSlot defines a method to create a slot.
  rpc CreateSlot(MsgCreateSlot) returns (MsgCreateSlotResponse);

  // EditSlot defines a method to edit a slot.
  rpc EditSlot(MsgEditSlot) returns (MsgEditSlotResponse);
}
```

#### `app/sft`

Applications can have their specific message services, which look as follows:

```protobuf
Service Msg {
  // TransferValueSFT defines a method to transfer value between sfts that are of the same slot.
  rpc TransferValueSFT(MsgTransferValueSFT) returns (MsgTransferValueSFTResponse);

  // CreateSlot defines a method to create a slot.
  rpc CreateSlot(MsgCreateSlot) returns (MsgCreateSlotResponse);

  // EditSlot defines a method to edit a slot.
  rpc EditSlot(MsgEditSlot) returns (MsgEditSlotResponse);
  // IssueClass defines a method to issue a class.
  rpc IssueClass(MsgIssueClass) returns (MsgIssueClassResponse);

  // EditClass defines a method to edit a class.
  rpc EditClass(MsgEditClass) returns (MsgEditClassResponse);

  // TransferClass defines a method to transfer ownership of a sft class from one account to another.
  rpc TransferClass(MsgTransferClass) returns (MsgTransferClassResponse);

  // MintSFT defines a method to mint an sft under a class.
  rpc MintSFT(MsgMintSFT) returns (MsgMintSFTResponse);

  // EditSFT defines a method to edit mutable part of an sft infos.
  rpc EditSFT(MsgEditBaseSFT) returns (MsgEditBaseSFTResponse);

  // TransferSFT defines a method to transfer an sft from one account to another.
  rpc TransferSFT(MsgTransferSFT) returns (MsgTransferSFTResponse);

  // BurnSFT defines a method to burn an sft under a class.
  rpc BurnSFT(MsgBurnSFT) returns (MsgBurnSFTResponse);
}

// MsgIssueDenom represents a message to issue a class.  
message MsgIssueClass {  
  string id = 1;  
  string name = 2;  
  string symbol = 3;  
  string description = 4;  
  string uri = 5;  
  string uri_hash = 6;  
  string schema = 7;  
  bool mint_restricted = 8;  
  bool update_restricted = 9;  
  string data = 10;  
  string sender = 11;  
}  
  
// MsgIssueClassResponse defines the Msg/IssueClass response type.  
message MsgIssueClassResponse {}  
  
// MsgEditClass represents a message to edit a class  
message MsgEditClass {  
  string id = 1;  
  string name = 2;  
  string symbol = 3;  
  string description = 4;  
  string uri = 5;  
  string uri_hash = 6;  
  string schema = 7;  
  bool mint_restricted = 8;  
  bool update_restricted = 9;  
  string data = 10;  
  string sender = 11;  
}  
  
// MsgEditClassResponse defines the Msg/EditClass response type.  
message MsgEditClassResponse {}  
  
// MsgTransferClass represents a message to transfer a class  
message MsgTransferClass {  
  string id = 1;  
  string sender = 2;  
  string recipient = 3;  
}  
  
// MsgTransferClassResponse defines the Msg/TransferClass response type.  
message MsgTransferClassResponse {}  
  
// MsgMintSFT represents a message to mint an sft.  
message MsgMintSFT {  
  string id = 1;  
  string class_id = 2;  
  string slot_id = 3;  
  string value = 4;  
  string name = 5;  
  string uri = 6;  
  string uri_hash = 7;  
  string data = 8;  
  string sender = 9;  
  string recipient = 10;  
}  
  
// MsgMintSFTResponse defines the Msg/MintSFT response type.  
message MsgMintSFTResponse {}  
  
// MsgEditSFT represents a message to edit the base part of an sft.  
message MsgEditBaseSFT {  
  string id = 1;  
  string class_id = 2;  
  string name = 3;  
  string uri = 4;  
  string uri_hash = 5;  
  string data = 6;  
  string sender = 7;  
}  
  
// MsgEditBaseSFTResponse defines the Msg/EditBaseSFT response type.  
message MsgEditBaseSFTResponse {}  
  
// MsgTransferSFT represents a message to transfer an sft.  
message MsgTransferSFT {  
  string id = 1;  
  string class_id = 2;  
  string name = 3;  
  string uri = 4;  
  string uri_hash = 5;  
  string data = 6;  
  string sender = 7;  
  string recipient = 8;  
}  
  
// MsgTransferSFTResponse defines the Msg/TransferSFT response type.  
message MsgTransferSFTResponse {}  
  
// MsgBurnSFT represents a message to burn an sft.  
message MsgBurnSFT {  
  string id = 1;  
  string class_id = 2;  
  string sender = 3;  
}  
  
// MsgBurnSFTResponse defines the Msg/BurnSFT response type.  
message MsgBurnSFTResponse {}  
  
// MsgTransferValueSFT represents a message to transfer value between sfts.  
message MsgTransferValueSFT {  
  string class_id = 1;  
  string slot_id = 2;  
  string src_id = 3;  
  string dest_id = 4;  
  uint64 value = 5;  
  string slot_uri = 6;  
  string slot_uri_hash = 7;  
  string sender = 8;  
  string recipient = 9;  
}  
  
// MsgTransferValueSFTResponse defines the Msg/TransferValueSFT response type.  
message MsgTransferValueSFTResponse {  
  string id = 1;  
  string recipient = 2;  
  bool minted = 3;  
}
  
// MsgCreateSlot  
message MsgCreateSlot {  
  string id = 1;  
  string name = 2;  
  string description = 3;  
  string uri = 4;  
  string uri_hash = 5;  
  string sender = 6;  
}  
  
message MsgCreateSlotResponse {}  
  
// MsgEditSlot  
message MsgEditSlot {  
  string id = 1;  
  string name = 2;  
  string description = 3;  
  string uri = 4;  
  string uri_hash = 5;  
  string sender = 6;  
}  
  
message MsgEditSlotResponse {} 
```

### `Query` Service

#### `x/sft`

```protobuf
service Query {
  // Slot queries an Slot by slot id.
  rpc Slot(QuerySlotRequest) returns (QuerySlotResponse) {
    option (google.api.http).get = "cosmos/x/sft/slots/{slot_id}";
  }

  // Slots queries all Slots
  rpc Slots(QuerySlotsRequest) returns (QuerySlotsResponse) {
    option (google.api.http).get = "cosmos/x/sft/slots";
  }

  // SFTsOfSlot queries all SFTs of a class with the same slot.
  rpc SFTsOfSlot(QuerySFTsOfSlotRequest) returns (QuerySFTsOfSlotResponse) {
    option (google.api.http).get = "cosmos/x/sft/slots/{slot_id}/{class_id}";
  }
}
```

#### `app/sft`

Applications can have their specific query services, which look as follows:

```protobuf
service Query {  
  // Class queries an SFT Class by class id.  
  rpc Class(QueryClassRequest) returns (QueryClassResponse) {  
    option (google.api.http).get = "irismod/sft/classes/{class_id}";  
  }  
  
  // Classes queries all SFT classes.  
  rpc Classes(QueryClassesRequest) returns (QueryClassesResponse) {  
    option (google.api.http).get = "irismod/sft/classes";  
  }  
  
  // Supply queries the number of SFTs of a given class by class id.  
  rpc Supply(QuerySupplyRequest) returns (QuerySupplyResponse) {  
    option (google.api.http).get = "irismod/sft/supply/{class_id}";  
  }  
  
  // Collection queries all SFTs of a given class by class id.  
  rpc Collection(QueryCollectionRequest) returns (QueryCollectionResponse) {  
    option (google.api.http).get = "irismod/sft/sfts/{class_id}";  
  }  
  
  // SFT queries an SFT by class id and sft id.  
  rpc SFT(QuerySFTRequest) returns (QuerySFTResponse) {  
    option (google.api.http).get = "irismod/sft/sfts/{class_id}/{sft_id}";  
  }  
  
  // Owner queries the owner of an SFT by class id and sft id.  
  rpc Owner(QueryOwnerRequest) returns (QueryOwnerResponse) {  
    option (google.api.http).get = "irismod/sft/owner/{class_id}/{sft_id}";  
  }  
  
  // SFTsOfOwner queries all SFTs of an owner by class_id  
  rpc SFTsOfOwner(QuerySFTsOfOwnerRequest) returns (QuerySFTsOfOwnerResponse) {  
    option (google.api.http).get = "irismod/sft/sfts";  
  }  
  
  // Balance queries the number of SFTs of an owner by owner and class id.  
  rpc Balance(QueryBalanceRequest) returns (QueryBalanceResponse) {  
    option (google.api.http).get = "irismod/sft/balance/{owner}/{class_id}";  
  }  
  
  // Slot queries an Slot by slot id.  
  rpc Slot(QuerySlotRequest) returns (QuerySlotResponse) {  
    option (google.api.http).get = "irismod/sft/slots/{slot_id}";  
  }  
  
  // Slots queries all Slots  
  rpc Slots(QuerySlotsRequest) returns (QuerySlotsResponse) {  
    option (google.api.http).get = "irismod/sft/slots";  
  }  
  
  // SFTsOfSlot queries all SFTs of a class with the same slot.  
  rpc SFTsOfSlot(QuerySFTsOfSlotRequest) returns (QuerySFTsOfSlotResponse) {  
    option (google.api.http).get = "irismod/sft/slots/{slot_id}/{class_id}";  
  }  
}  
  
// QueryClassRequest is the request type for the Query/Class RPC method  
message QueryClassRequest {  
  string class_id = 1;  
}  
  
// QueryClassResponse is the response type for the Query/Class RPC method  
message QueryClassResponse {  
  irismod.sft.Class class = 1;  
}  
  
// QueryClassesRequest is the request type for the Query/Classes RPC method  
message QueryClassesRequest {  
  cosmos.base.query.v1beta1.PageRequest pagination = 1;  
}  
  
// QueryClassesResponse is the response type for the Query/Classes RPC method  
message QueryClassesResponse {  
  repeated irismod.sft.Class classes = 1;  
  cosmos.base.query.v1beta1.PageResponse pagination = 2;  
}  
  
// QuerySupplyRequest is the request type for the Query/Supply RPC method  
message QuerySupplyRequest {  
  string class_id = 1;  
}  
  
// QuerySupplyResponse is the response type for the Query/Supply RPC method  
message QuerySupplyResponse {  
  uint64 amount = 1;  
}  
  
// QueryCollectionRequest is the request type for the Query/Collection RPC method  
message QueryCollectionRequest {  
  string class_id = 1;  
  cosmos.base.query.v1beta1.PageRequest pagination = 1;  
}  
  
// QueryCollectionResponse is the response type for the Query/Collection RPC method  
message QueryCollectionResponse {  
  repeated irismod.sft.SFT sfts = 1;  
  cosmos.base.query.v1beta1.PageResponse pagination = 2;  
}  
  
// QuerySFTRequest is the request type for the Query/SFT RPC method  
message QuerySFTRequest {  
  string class_id = 1;  
  string sft_id = 2;  
}  
  
// QuerySFTResponse is the response type for the Query/SFT RPC method  
message QuerySFTResponse {  
  irismod.sft.SFT sft = 1;  
}  
  
// QueryOwnerRequest is the request type for the Query/Owner RPC method  
message QueryOwnerRequest {  
  string class_id = 1;  
  string sft_id = 2;  
}  
  
// QueryOwnerResponse is the response type for the Query/Owner RPC method  
message QueryOwnerResponse {  
  string owner = 1;  
}  
  
// QuerySFTsOfOwnerRequest is the request type for the Query/SFTsOfOwner RPC method  
message QuerySFTsOfOwnerRequest {  
  string class_id = 1;  
  string owner = 2;  
  cosmos.base.query.v1beta1.PageRequest pagination = 1;  
}  
  
// QuerySFTsOfOwnerResponse is the response type for the Query/SFTsOfOwner RPC method  
message QuerySFTsOfOwnerResponse {  
  repeated irismod.sft.SFT sfts = 1;  
  cosmos.base.query.v1beta1.PageResponse pagination = 2;  
}  
  
// QueryBalanceRequest is the request type for the Query/Balance RPC method  
message QueryBalanceRequest {  
  string class_id = 1;  
  string owner = 2;  
}  
  
// QueryBalanceResponse is the response type for the Query/Balance RPC method  
message QueryBalanceResponse {  
  uint64 amount = 1;  
}  
  
// QuerySlotRequest is the request type for the Query/Slot RPC method  
message QuerySlotRequest {  
  string slot_id = 1;  
}  
  
// QuerySlotResponse is the response type for the Query/Slot RPC method  
message QuerySlotResponse {  
  irismod.sft.Slot slot = 1;  
}  
  
// QuerySlotsRequest is the request type for the Query/Slots RPC method  
message QuerySlotsRequest {  
  cosmos.base.query.v1beta1.PageRequest pagination = 1;  
}  
  
// QuerySlotsResponse is the response type for the Query/Slots RPC method  
message QuerySlotsResponse {  
  repeated irismod.sft.Slot slots = 1;  
  cosmos.base.query.v1beta1.PageResponse pagination = 2;  
}  
  
// QuerySFTsOfSlotRequest is the request type for the Query/SFTsOfSlot RPC method  
message QuerySFTsOfSlotRequest {  
  string slot_id = 1;  
  string class_id = 2;  
  cosmos.base.query.v1beta1.PageRequest pagination = 3;  
}  
  
// QuerySFTsOfSlotResponse is the response type for the Query/SFTsOfSlot RPC method  
message QuerySFTsOfSlotResponse {  
  repeated irismod.sft.Slot slots = 1;  
  cosmos.base.query.v1beta1.PageResponse pagination = 2;  
}
```

## Consequences

### Backwards Compatibility

No Cosmos-SDK backward incompatibilities.

### Positive

- Flexibility of buiding SFT app modules on the Cosmos Ecosystem.
- Ability to reuse the `x/nft` module.

### Negative

### Neutral

## Further Discussions

The interoperability needs further discussion.

## References

- [ERC-3525](https://eips.ethereum.org/EIPS/eip-3525)
- [ADR 043](https://github.com/cosmos/cosmos-sdk/blob/main/docs/architecture/adr-043-nft-module.md)
