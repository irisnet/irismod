package types

// ERC721Data is the struct that holds the data of an ERC721 token
type ERC721Data struct {
	Name   string
	Symbol string
}

// NewERC721Data creates a new ERC20Data instance
func NewERC721Data(name, symbol string, decimals uint8) ERC721Data {
	return ERC721Data{
		Name:   name,
		Symbol: symbol,
	}
}
