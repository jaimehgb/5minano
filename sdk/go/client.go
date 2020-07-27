package _5minano

import (
	"context"
	"net/url"
)

// Client defines the methods we can use to reach 5minano api.
type Client interface {
	// RegisterWallet will register a public seed to generate public keys
	// without having to access wallet the private keys.
	RegisterWallet(context.Context, *Wallet) error

	// SubscribeAddress will register an address to 5minano address listener.
	// All transactions received by that address will send a notification to the set URL.
	SubscribeAddress(context.Context, *Address, *url.URL) error

	// WalletBalance will retrieve the balance of the given wallet.
	WalletBalance(context.Context, *Wallet) (*Balance, error)

	// AddressBalance will retrieve the balance of the given address.
	AddressBalance(context.Context, *Address) (*Balance, error)

	// PendingWebhooks will retrieve all webhooks pending to be received, for the given wallet.
	PendingWebhooks(context.Context, *Wallet) ([]*Webhook, error)

	// FailedWebhooks will retrieve all webhooks which failed to be delivered.
	FailedWebhooks(context.Context, *Wallet) ([]*Webhook, error)

	// PendingReceives will retrieve all unpocketed blocks for the given addresses.
	PendingReceives(context.Context, []*Address) ([]*Block, error)

	// Send will submit a block to be broadcasted to the network.
	Send(context.Context, *Block) error

	// PendingPoW will retrieve all blocks which are waiting for PoW for a given Wallet.
	PendingPoW(context.Context, *Wallet) ([]*Block, error)
}
