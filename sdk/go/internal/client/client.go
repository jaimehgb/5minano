package client

import (
	"context"
	"net/url"

	sdk "github.com/jaimehgb/5minano-go-sdk"
)

type client struct {
}

func NewClient() sdk.Client {
	return &client{}
}

func (c *client) RegisterWallet(ctx context.Context, w *sdk.Wallet) error {
	return nil
}

func (c *client) SubscribeAddress(ctx context.Context, a *sdk.Address, url *url.URL) error {
	return nil
}

func (c *client) WalletBalance(ctx context.Context, w *sdk.Wallet) (*sdk.Balance, error) {
	return nil, nil
}

func (c *client) AddressBalance(ctx context.Context, a *sdk.Address) (*sdk.Balance, error) {
	return nil, nil
}

func (c *client) PendingWebhooks(ctx context.Context, w *sdk.Wallet) ([]*sdk.Webhook, error) {
	return nil, nil
}

func (c *client) FailedWebhooks(ctx context.Context, w *sdk.Wallet) ([]*sdk.Webhook, error) {
	return nil, nil
}

func (c *client) PendingReceives(ctx context.Context, a []*sdk.Address) ([]*sdk.Block, error) {
	return nil, nil
}

func (c *client) Send(ctx context.Context, b *sdk.Block) error {
	return nil
}

func (c *client) PendingPoW(ctx context.Context, w *sdk.Wallet) ([]*sdk.Block, error) {
	return nil, nil
}
