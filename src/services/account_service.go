package services

import (
	"context"
	"github.com/coinbase/rosetta-sdk-go/server"
	"github.com/coinbase/rosetta-sdk-go/types"
)

// AccountAPIServicer implements the server.AccountAPIServicer interface.
type AccountAPIServicer struct {
	network *types.NetworkIdentifier
}

// NewAccountAPIService creates a new instance of a AccountAPIServicer.
func NewAccountAPIService(network *types.NetworkIdentifier) server.AccountAPIServicer {
	return &AccountAPIServicer{
		network: network,
	}
}

// AccountBalance implements the /account/balance endpoint.
// AccountBalance - Get an Account Balance
func (s *AccountAPIServicer) AccountBalance(
	ctx context.Context,
	request *types.AccountBalanceRequest,
) (*types.AccountBalanceResponse, *types.Error) {
	balance, hash, height := GetAccountBalance(request.AccountIdentifier.Address)

	amount := []types.Amount{
		{
			Value: balance,
			Currency: &types.Currency{
				Symbol:   "ICX",
				Decimals: 18,
				Metadata: nil,
			},
			Metadata: nil,
		},
	}

	var ptrAmount []*types.Amount
	ptrAmount = append(ptrAmount, &amount[0])

	return &types.AccountBalanceResponse{
		BlockIdentifier: &types.BlockIdentifier{
			Hash:  hash,
			Index: height,
		},
		Balances: ptrAmount,
	}, nil
}
