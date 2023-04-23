package types

// ERC721Data is the struct that holds the data of an ERC721 token
type ERC721Data struct {
	Name   string
	Symbol string
}

// ERC721StringResponse defines the string value from the call response
type ERC721StringResponse struct {
	Value string
}

// NewERC721Data creates a new ERC20Data instance
func NewERC721Data(name, symbol string) ERC721Data {
	return ERC721Data{
		Name:   name,
		Symbol: symbol,
	}
}
