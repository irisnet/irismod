package cli

import (
	flag "github.com/spf13/pflag"

	randomtypes "github.com/irisnet/irismod/modules/random/types"
)

const (
	FlagReqID         = "request-id"
	FlagBlockInterval = "block-interval"
	FlagOracle        = "oracle"
	FlagServiceFeeCap = "service-fee-cap"
	FlagQueueHeight   = "queue-height"
)

var (
	FsRequestRand = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	FsRequestRand.Uint64(FlagBlockInterval, randomtypes.DefaultBlockInterval, "the block interval")
	FsRequestRand.Bool(FlagOracle, false, "woth oracle method")
	FsRequestRand.String(FlagServiceFeeCap, "", "maximal fee to pay for a service request")
}
