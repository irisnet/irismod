package contracts

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"

	"github.com/irisnet/irismod/modules/erc721-converter/types"

	"github.com/ethereum/go-ethereum/common"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
)

var (
	//go:embed compiled_contracts/ERC721PresetMinterPauserAutoId.json
	ERC721PresetMinterPauserAutoIdJSON []byte //nolint: golint

	// ERC721PresetMinterPauserAutoIdContract is the compiled erc721 contract
	ERC721PresetMinterPauserAutoIdContract evmtypes.CompiledContract

	// ERC721PresetMinterPauserAutoIdAddress is the erc721 module address
	ERC721PresetMinterPauserAutoIdAddress common.Address
)

func init() {
	ERC721PresetMinterPauserAutoIdAddress = types.ModuleAddress

	err := json.Unmarshal(ERC721PresetMinterPauserAutoIdJSON, &ERC721PresetMinterPauserAutoIdContract)
	if err != nil {
		panic(err)
	}
}
