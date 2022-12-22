package keeper

import (
	"encoding/json"
	"github.com/irisnet/irismod/modules/nft/types"
)

// decomposeDenomData decomposes "{"data":"","denomPlugin":""} into data string and denomPlugin struct
func decomposeDenomData(data string) (string, *types.DenomPlugin) {
	var dcd types.DenomComposedData

	// if failed to unmarshal, return DenomPlugin with nil field
	if err := json.Unmarshal([]byte(data), &dcd); err != nil {
		return data, &types.DenomPlugin{RentalPlugin: nil}
	}

	return dcd.UserData, dcd.DenomPlugin
}
