package cosmos

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/strangelove-ventures/interchaintest/v8/ibc"
)

// AccAddressFromBech32 creates an AccAddress from a Bech32 string.
// https://github.com/cosmos/cosmos-sdk/blob/v0.50.2/types/address.go#L193-L212
func (c *CosmosChain) AccAddressFromBech32(address string) (addr sdk.AccAddress, err error) {
	// Use the common implementation
	converter := ibc.NewCommonAddressBech32Converter(c.Config())
	return converter.AccAddressFromBech32(address)
}

func (c *CosmosChain) AccAddressToBech32(addr sdk.AccAddress) (string, error) {
	// Use the common implementation
	converter := ibc.NewCommonAddressBech32Converter(c.Config())
	return converter.AccAddressToBech32(addr)
}
