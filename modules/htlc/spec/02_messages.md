<!--
order: 2
-->

# Messages

## MsgCreateHTLC

xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

```go
type MsgCreateHTLC struct {
    Sender               string
    To                   string
    ReceiverOnOtherChain string
    Amount               sdk.Coins
    HashLock             string
    Timestamp            uint64
    TimeLock             uint64
}
```

xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

## MsgClaimHTLC

xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

```go
type MsgClaimHTLC struct {
    Sender   string
    HashLock string
    Secret   string
}
```

xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

## MsgRefundHTLC

xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

```go
type MsgRefundHTLC struct {
    Sender   string
    HashLock string
}
```

xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
