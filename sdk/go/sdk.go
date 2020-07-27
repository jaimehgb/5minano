package _5minano

import "context"

type SDK interface {
	// NewAddress will create a new address and subscribe it to 5minano listener.
	NewAddress(context.Context) (*Address, error)

	// SendTo will send the specified amount to the designed address.
	// The origin address will be chose automatically based on the already pocketed balance.
	SendTo(context.Context, *Address, *Amount) (*Block, error)

	// SendFromTo will send the specified amount from a specific address to the destination chosen.
	SendFromTo(context.Context, *Address, *Address, *Amount) (*Block, error)

	// WalletBalance will return the wallet balance.
	WalletBalance(context.Context) (*Balance, error)

	// ChangeRepresentative will change the representative of the specified address.
	ChangeRepresentative(context.Context, *Address, *Address) error

	// PendingBlocks will return all blocks which are waiting for PoW for this wallet.
	PendingBlocks(context.Context) ([]*Block, error)

	// UnpocketedBlocks will return all unpocketed blocks for this wallet.
	UnpocketedBlocks(context.Context) ([]*Block, error)
}
