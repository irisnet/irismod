package types

import "fmt"

// CreateClass generates a string the module name plus the address to avoid conflicts with names staring with a number
func CreateClass(address string) string {
	return fmt.Sprintf("%s/%s", ModuleName, address)
}
