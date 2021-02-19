<!--
order: 2
-->

# Messages

## MsgCreateFeed

xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

```go
type MsgCreateFeed struct {
    FeedName          string
    LatestHistory     uint64
    Description       string
    Creator           string
    ServiceName       string
    Providers         []string
    Input             string
    Timeout           int64
    ServiceFeeCap     sdk.Coins
    RepeatedFrequency uint64
    AggregateFunc     string
    ValueJsonPath     string
    ResponseThreshold uint32
}
```

xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

## MsgStartFeed

xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

```go
type MsgStartFeed struct {
    FeedName string
    Creator  string
}
```

xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

## MsgPauseFeed

xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

```go
type MsgPauseFeed struct {
    FeedName string
    Creator  string
}
```

xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

## MsgEditFeed

xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

```go
type MsgEditFeed struct {
    FeedName          string
    Description       string
    LatestHistory     uint64
    Providers         []string
    Timeout           int64
    ServiceFeeCap     sdk.Coins
    RepeatedFrequency uint64
    ResponseThreshold uint32
    Creator           string
}
```

xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
