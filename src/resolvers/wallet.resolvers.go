package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"fmt"

	"osman-hakan.com/graphql-blockchain/src/model"
	"osman-hakan.com/graphql-blockchain/src/services"
	"osman-hakan.com/graphql-blockchain/src/utils"
)

// CreateWallet is the resolver for the createWallet field.
func (r *mutationResolver) CreateWallet(ctx context.Context, input *model.CreateWalletInput) (*model.Wallet, error) {
	newWallet, err := services.CreateWallet()

	if err != nil {
		return nil, err
	}

	return newWallet, nil
}

// TransferToken is the resolver for the transferToken field.
func (r *mutationResolver) TransferToken(ctx context.Context, input *model.TransferToken) (string, error) {
	bigInt := utils.StringToBigInt(input.Amount)

	if bigInt == nil {
		return "", ctx.Err()
	}

	tra, err := services.TransferToken(input.RPCLink, input.FromPrivate, input.ToPublic, bigInt)

	fmt.Println(tra)

	if err != nil {
		return "", err
	}

	return "ok", nil
}

// DeployContract is the resolver for the deployContract field.
func (r *mutationResolver) DeployContract(ctx context.Context, input *model.DeployContractInput) (string, error) {
	deployedContractAddress, err := services.DeployContract(input.RPCLink, uint64(input.ChainID), input.PrivateKey, input.Name, input.Symbol, uint64(input.Supply))

	if err != nil {
		return "", err
	}

	hexAddress := deployedContractAddress.Hex()

	return hexAddress, nil
}
