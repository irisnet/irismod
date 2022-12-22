package cli

import (
	"encoding/json"
	"github.com/irisnet/irismod/modules/nft/types"
	"github.com/spf13/cobra"
)

// composeDenomData composes an json string from denomPlugin and user data
func composeDenomData(cmd *cobra.Command, data string) (string, error) {
	denomPlugin, err := composeDenomPlugin(cmd)
	if err != nil {
		return data, err
	}

	composedData := types.DenomComposedData{
		UserData:    data,
		DenomPlugin: denomPlugin,
	}

	bz, err := json.Marshal(composedData)
	if err != nil {
		return data, err
	}

	return string(bz), nil
}

// composeDenomPlugin compose DenomPlugin from cli flag
func composeDenomPlugin(cmd *cobra.Command) (*types.DenomPlugin, error) {
	rentalPlugin, err := composeRentalPlugin(cmd)
	if err != nil {
		return nil, err
	}

	return &types.DenomPlugin{
		RentalPlugin: rentalPlugin,
	}, nil
}

// composeRentalPlugin compose RentalPlugin from cli flag
func composeRentalPlugin(cmd *cobra.Command) (*types.RentalPlugin, error) {
	rental, err := cmd.Flags().GetBool(FlagRentable)
	if err != nil {
		return nil, err
	}
	return &types.RentalPlugin{
		Enabled: rental,
	}, nil
}
