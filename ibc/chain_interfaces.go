package ibc

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AddressBech32Converter defines methods for converting between bech32 addresses and AccAddress.
// This is common functionality needed by Cosmos-based chains.
type AddressBech32Converter interface {
	// AccAddressFromBech32 creates an AccAddress from a Bech32 string.
	AccAddressFromBech32(address string) (sdk.AccAddress, error)

	// AccAddressToBech32 converts an AccAddress to a Bech32 string.
	AccAddressToBech32(addr sdk.AccAddress) (string, error)
}

// User is an interface implemented by wallet types to provide additional chain-specific
// functionality beyond the basic Wallet interface.
type User interface {
	// FormattedAddressWithPrefix returns the address formatted with a specific prefix.
	FormattedAddressWithPrefix(prefix string) string
}

// Execer is an interface for objects that can execute commands, usually through a Docker container.
// This matches the interface described in issue #256.
type Execer interface {
	// Exec runs an arbitrary command and returns stdout, stderr, and any error.
	Exec(ctx context.Context, cmd []string, env []string) (stdout, stderr []byte, err error)
}
