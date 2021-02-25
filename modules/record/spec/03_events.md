<!--
order: 3
-->

# Events

The record module emits the following events:

## Handlers

### MsgCreateRecord

| Type          | Attribute Key | Attribute Value  |
| :------------ | :------------ | :--------------- |
| create_record | creator       | {creatorAddress} |
| create_record | record_id     | {recordId}       |
| message       | module        | record           |
| message       | sender        | {creatorAddress} |
