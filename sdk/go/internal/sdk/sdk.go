package sdk

import (
	"context"

	sdk "github.com/jaimehgb/5minano-go-sdk"
)

type fiveminano struct {
}

func New() sdk.SDK {
	return &fiveminano{}
}

func (s *fiveminano) NewAddress(context.Context) (*sdk.Address, error) {
	return nil, nil
}

func (s *fiveminano) SendTo(context.Context, *sdk.Address, *sdk.Amount) (*sdk.Block, error) {
	return nil, nil
}

func (s *fiveminano) SendFromTo(context.Context, *sdk.Address, *sdk.Address, *sdk.Amount) (*sdk.Block, error) {
	return nil, nil
}

func (s *fiveminano) WalletBalance(context.Context) (*sdk.Balance, error) {
	return nil, nil
}

func (s *fiveminano) ChangeRepresentative(context.Context, *sdk.Address, *sdk.Address) error {
	return nil
}

func (s *fiveminano) PendingBlocks(context.Context) ([]*sdk.Block, error) {
	return nil, nil
}

func (s *fiveminano) UnpocketedBlocks(context.Context) ([]*sdk.Block, error) {
	return nil, nil
}
