package keeper

import "github.com/irisnet/irismod/modules/nft/types"

func (k Keeper) UserDataToDenomPlugin(data string) *types.DenomPlugin {

	// 1. Unmarshal denomPlugin
	var denomPlugin *types.DenomPlugin
	if err := k.cdc.Unmarshal([]byte(data), denomPlugin); err != nil {
		return nil
	}
	return denomPlugin
}

func (k Keeper) UserDataToTokenPlugin(data string) *types.TokenPlugin {

	// 1. Unmarshal denomPlugin
	var tokenPlugin *types.TokenPlugin
	if err := k.cdc.Unmarshal([]byte(data), tokenPlugin); err != nil {
		return nil
	}
	return tokenPlugin
}
