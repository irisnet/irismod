package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

// Rest variable names
// nolint
const (
	RestPoolName = "pool-name"
)

// RegisterHandlers defines routes that get registered by the main application
func RegisterHandlers(cliCtx client.Context, r *mux.Router) {
	registerTxRoutes(cliCtx, r)
}

type CreateFarmPoolReq struct {
	BaseReq        rest.BaseReq `json:"base_req"` // base req
	Name           string       `json:"name"`
	Description    string       `json:"description,omitempty"`
	LpTokenDenom   string       `json:"lp_token_denom"`
	StartHeight    uint64       `json:"start_height"`
	RewardPerBlock sdk.Coins    `json:"reward_per_block"`
	TotalReward    sdk.Coins    `json:"total_reward"`
	Destructible   bool         `json:"destructible,omitempty"`
	Creator        string       `json:"creator"`
}

type DestroyFarmPoolReq struct {
	BaseReq rest.BaseReq `json:"base_req"` // base req
	Creator string       `json:"creator"`
}

type AppendRewardReq struct {
	BaseReq  rest.BaseReq `json:"base_req"` // base req
	PoolName string       `json:"pool_name"`
	Reward   sdk.Coins    `json:"reward"`
	Creator  string       `json:"creator"`
}

type StakeReq struct {
	BaseReq rest.BaseReq `json:"base_req"` // base req
	Amount  sdk.Coin     `json:"reward"`
	Sender  string       `json:"creator"`
}

type UnstakeReq struct {
	BaseReq rest.BaseReq `json:"base_req"` // base req
	Amount  sdk.Coin     `json:"reward"`
	Sender  string       `json:"creator"`
}

type HarvestReq struct {
	BaseReq rest.BaseReq `json:"base_req"` // base req
	Sender  string       `json:"creator"`
}
