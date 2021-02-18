<!--
order: 2
-->

# Messages

## MsgSwapOrder

xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

```go
type MsgSwapOrder struct {
    Input      Input
    Output     Output
    Deadline   int64
    IsBuyOrder bool
}
```

```go
type Input struct {
    Address string
    Coin    types.Coin
}
```

```go
type Output struct {
    Address string
    Coin    types.Coin
}

```

xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

## MsgAddLiquidity

xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

```go
type MsgAddLiquidity struct {
    MaxToken         types.Coin
    ExactStandardAmt sdk.Int
    MinLiquidity     sdk.Int
    Deadline         int64
    Sender           string
}
```

xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

## MsgRemoveLiquidity

xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

```go
type MsgRemoveLiquidity struct {
    WithdrawLiquidity types.Coin
    MinToken          sdk.Int
    MinStandardAmt    sdk.Int
    Deadline          int64
    Sender            string
}
```

xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
