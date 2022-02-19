# Messages

## MsgIssueDenom

This message defines a type of mt, there can be multiple mt of the same type

| **Field**        | **Type** | **Description**                                                                                                                        |
|:-----------------|:---------|:---------------------------------------------------------------------------------------------------------------------------------------|
| Id               | `string` | The denomination ID of the MT, necessary as multiple denominations are able to be represented on each chain.                          |
| Name             | `string` | The denomination name of the MT, necessary as multiple denominations are able to be represented on each chain.                        |
| Sender           | `string` | The account address of the user sending the MT. By default it is __not__ required that the sender is also the owner of the MT.       |
| Schema           | `string` | MT specifications defined under this category                                                                                         |
| Symbol           | `string` | The abbreviated name of a specific MT type                                                                                            |
| MintRestricted   | `bool`   | MintRestricted is true means that only Denom owners can issue MTs under this category, false means anyone can                         |                                                                        |
| UpdateRestricted | `bool`   | UpdateRestricted is true means that no one in this category can update the MT, false means that only the owner of this MT can update |                                                                             |
| Description      | `string` | Description is a brief description of mt classification. Optional                                                                     |                                                                             |
| URI              | `string` | The uri_hash is a hash of the document pointed by uri. Optional                                                                        |                                                                             |
| UriHash          | `string` | UriHash is true means that no one in this category can update the MT, false means that only the owner of this MT can update          |                                                                             |
| Data             | `string` | data is the app specific metadata of the MT class. Optional                                                                           |                                                                             |

```go
type MsgIssueDenom struct {
    Id               string
    Name             string
    Schema           string
    Sender           string
    Symbol           string
    MintRestricted   bool
    UpdateRestricted bool
    Description      string
    Uri              string
    UriHash          string
    Data             string
}
```

## MsgTransferMT

This is the most commonly expected MsgType to be supported across chains. While each application specific blockchain will have very different adoption of the `MsgMintMT`, `MsgBurnMT` and `MsgEditMT` it should be expected that most chains support the ability to transfer ownership of the MT asset. The exception to this would be non-transferable MTs that might be attached to reputation or some asset which should not be transferable. It still makes sense for this to be represented as an MT because there are common queriers which will remain relevant to the MT type even if non-transferable. This Message will fail if the MT does not exist. By default it will not fail if the transfer is executed by someone beside the owner. **It is highly recommended that a custom handler is made to restrict use of this Message type to prevent unintended use.**

| **Field** | **Type** | **Description**                                                                                                                  |
|:----------|:---------|:---------------------------------------------------------------------------------------------------------------------------------|
| ID        | `string` | The unique ID of the MT being transferred.                                                                                      |
| DenomId   | `string` | The unique ID of the denomination, necessary as multiple denominations are able to be represented on each chain.                 |
| Name      | `string` | The name of the MT being transferred.                                                                                           |
| URI       | `string` | The URI pointing to a JSON object that contains subsequent tokenData information off-chain                                       |
| UriHash   | `string` | The uri_hash is a hash of the document pointed by uri. Optional                                                                  |                                                                             |
| Data      | `string` | The data of the MT                                                                                                              |
| Sender    | `string` | The account address of the user sending the MT. By default it is __not__ required that the sender is also the owner of the MT. |
| Recipient | `string` | The account address who will receive the MT as a result of the transfer transaction.                                            |

```go
// MsgTransferMT defines an SDK message for transferring an MT to recipient.
type MsgTransferMT struct {
    Id        string
    DenomId   string
    Name      string
    URI       string
    UriHash   string
    Data      string
    Sender    string
    Recipient string
}
```

## MsgEditMT

This message type allows the `TokenURI` to be updated. By default anyone can execute this Message type. **It is highly recommended that a custom handler is made to restrict use of this Message type to prevent unintended use.**

| **Field** | **Type** | **Description**                                                                                                  |
|:----------|:---------|:-----------------------------------------------------------------------------------------------------------------|
| Id        | `string` | The unique ID of the MT being edited.                                                                           |
| DenomId   | `string` | The unique ID of the denomination, necessary as multiple denominations are able to be represented on each chain. |
| Name      | `string` | The name of the MT being edited.                                                                                |
| URI       | `string` | The URI pointing to a JSON object that contains subsequent tokenData information off-chain                       |
| UriHash   | `string` | The uri_hash is a hash of the document pointed by uri. Optional                                                  |                                                                             |
| Data      | `string` | The data of the MT                                                                                              |
| Sender    | `string` | The creator of the message                                                                                       |

```go
// MsgEditMT defines an SDK message for editing a mt.
type MsgEditMT struct {
    Id      string
    DenomId string
    Name    string
    URI     string
    UriHash string
    Data    string
    Sender  string
}
```

## MsgMintMT

This message type is used for minting new tokens. If a new `MT` is minted under a new `Denom`, a new `Collection` will also be created, otherwise the `MT` is added to the existing `Collection`. If a new `MT` is minted by a new account, a new `Owner` is created, otherwise the `MT` `ID` is added to the existing `Owner`'s `IDCollection`. By default anyone can execute this Message type. **It is highly recommended that a custom handler is made to restrict use of this Message type to prevent unintended use.**

| **Field** | **Type** | **Description**                                                                            |
|:----------|:---------|:-------------------------------------------------------------------------------------------|
| Id        | `string` | The unique ID of the MT being minted                                                      |
| DenomId   | `string` | The unique ID of the denomination.                                                         |
| Name      | `string` | The name of the MT being minted.                                                          |
| URI       | `string` | The URI pointing to a JSON object that contains subsequent tokenData information off-chain |
| UriHash   | `string` | The uri_hash is a hash of the document pointed by uri. Optional                            |                                                                             |
| Data      | `string` | The data of the MT.                                                                       |
| Sender    | `string` | The sender of the Message                                                                  |
| Recipient | `string` | The recipiet of the new MT                                                                |

```go
// MsgMintMT defines an SDK message for creating a new MT.
type MsgMintMT struct {
    Id        string
    DenomId   string
    Name      string
    URI       string
	UriHash   string
    Data      string
    Sender    string
    Recipient string
}
```

### MsgBurnMT

This message type is used for burning tokens which destroys and deletes them. By default anyone can execute this Message type. **It is highly recommended that a custom handler is made to restrict use of this Message type to prevent unintended use.**

| **Field** | **Type** | **Description**                                    |
|:----------|:---------|:---------------------------------------------------|
| Id        | `string` | The ID of the Token.                               |
| DenomId   | `string` | The Denom ID of the Token.                         |
| Sender    | `string` | The account address of the user burning the token. |

```go
// MsgBurnMT defines an SDK message for burning a MT.
type MsgBurnMT struct {
    Id      string
    DenomId string
    Sender  string
}
```

## MsgTransferDenom
This message is used by the owner of the MT classification to transfer the ownership of the MT classification to others

| **Field** | **Type** | **Description**                                                                         |
|:----------|:---------|:----------------------------------------------------------------------------------------|
| ID        | `string` | The unique ID of the Denom being transferred.                                           | 
| Sender    | `string` | The account address of the user sending the Denom.                                      |
| Recipient | `string` | The account address who will receive the Denom as a result of the transfer transaction. |

```go
// MsgTransferDenom defines an SDK message for transferring an Denom to recipient.
type MsgTransferDenom struct {
    Id        string
    Sender    string
    Recipient string
}
```