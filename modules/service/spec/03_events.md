<!--
order: 3
-->

# Events

The service module emits the following events:

## Handlers

### MsgDefineService

| Type              | Attribute Key          | Attribute Value |
| ----------------- | ---------------------- | --------------- |
| create_definition | service_name           | service         |
| create_definition | author                 | {senderAddress} |
| message           | module                 | service         |
| message           | sender                 | {senderAddress} |

### MsgBindService

| Type    | Attribute Key | Attribute Value |
| ------- | ------------- | --------------- |
| create_binding | service_name        | service         |
| create_binding | provider        | {senderAddress} |
| create_binding | owner        | {senderAddress} |
| message | module        | service         |
| message | sender        | {senderAddress} |

### MsgUpdateServiceBinding

| Type    | Attribute Key | Attribute Value |
| ------- | ------------- | --------------- |
| update_binding | service_name        | service         |
| update_binding | provider        | {senderAddress} |
| update_binding | owner        | {senderAddress} |
| message | module        | service         |
| message | sender        | {senderAddress} |

### MsgDisableServiceBinding

| Type    | Attribute Key | Attribute Value |
| ------- | ------------- | --------------- |
| enable_binding | service_name        | service         |
| enable_binding | provider        | {senderAddress} |
| enable_binding | owner        | {senderAddress} |
| message | module        | service         |
| message | sender        | {senderAddress} |

### MsgEnableServiceBinding

| Type    | Attribute Key | Attribute Value |
| ------- | ------------- | --------------- |
| enable_binding | service_name        | service         |
| enable_binding | provider        | {senderAddress} |
| enable_binding | owner        | {senderAddress} |
| message | module        | service         |
| message | sender        | {senderAddress} |

### MsgRefundServiceDeposit

| Type    | Attribute Key | Attribute Value |
| ------- | ------------- | --------------- |
| message | module        | service         |
| message | sender        | {senderAddress} |

### MsgSetWithdrawAddress

| Type    | Attribute Key | Attribute Value |
| ------- | ------------- | --------------- |
| set_withdraw_address | withdraw_address        | {senderAddress} |
| set_withdraw_address | owner        | {senderAddress} |
| message | module        | service         |
| message | sender        | {senderAddress} |

### MsgCallService

| Type           | Attribute Key      | Attribute Value    |
| -------------- | ------------------ | ------------------ |
| create_context | request_context_id | {requestContextID} |
| message        | module             | service            |
| message        | sender             | {senderAddress}    |

### MsgRespondService

| Type            | Attribute Key | Attribute Value |
| --------------- | ------------- | --------------- |
| respond_service | request_id    | {requestID}     |
| message         | module        | service         |
| message         | sender        | {senderAddress} |

### MsgUpdateRequestContext

| Type    | Attribute Key | Attribute Value |
| ------- | ------------- | --------------- |
| message | module        | service         |
| message | sender        | {senderAddress} |

### MsgPauseRequestContext

| Type    | Attribute Key | Attribute Value |
| ------- | ------------- | --------------- |
| message | module        | service         |
| message | sender        | {senderAddress} |

### MsgStartRequestContext

| Type    | Attribute Key | Attribute Value |
| ------- | ------------- | --------------- |
| message | module        | service         |
| message | sender        | {senderAddress} |

### MsgKillRequestContext

| Type    | Attribute Key | Attribute Value |
| ------- | ------------- | --------------- |
| message | module        | service         |
| message | sender        | {senderAddress} |

### MsgWithdrawEarnedFees

| Type    | Attribute Key | Attribute Value |
| ------- | ------------- | --------------- |
| message | module        | service         |
| message | sender        | {senderAddress} |
