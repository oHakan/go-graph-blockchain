package services

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"osman-hakan.com/graphql-blockchain/src/contract"
	"osman-hakan.com/graphql-blockchain/src/model"
)

func CreateWallet() (*model.Wallet, error) {
	var walletModel model.Wallet

	privateKey, err := crypto.GenerateKey()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	privateKeyByte := crypto.FromECDSA(privateKey)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		log.Fatal("error public key ecdsa")
		return nil, err
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	walletModel.PrivateKey = hexutil.Encode(privateKeyByte)
	walletModel.PublicKey = hexutil.Encode(publicKeyBytes)
	walletModel.Address = address

	return &walletModel, nil
}

func GetClient(rpcLink string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(rpcLink)

	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	return client, nil
}

func TransferToken(rpcLink string, fromPrivateKey string, toPublicKey string, transferAmount *big.Int) (*types.Transaction, error) {
	client, err := GetClient(rpcLink)

	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(fromPrivateKey)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	toPublicKeyAddressType := common.HexToAddress(toPublicKey)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)

	if err != nil {
		return nil, err
	}

	value := transferAmount
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())

	var data []byte

	if err != nil {
		return nil, err
	}

	tx := types.NewTransaction(nonce, toPublicKeyAddressType, value, gasLimit, gasPrice, data)
	chainID, err := client.NetworkID(context.Background())

	if err != nil {
		return nil, err
	}

	ecdsaPrivateKey, err := crypto.HexToECDSA(fromPrivateKey)

	if err != nil {
		return nil, err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), ecdsaPrivateKey)

	if err != nil {
		return nil, err
	}

	err = client.SendTransaction(context.Background(), signedTx)

	if err != nil {
		return nil, err
	}

	return signedTx, err
}

func DeployContract(rpcLink string, chainId uint64, fromPrivateKey string, tokenName string, tokenSymbol string, tokenSupply uint64) (*common.Address, error) {
	client, err := GetClient(rpcLink)

	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(fromPrivateKey)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(chainId)))

	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(15000000)
	auth.GasPrice = gasPrice

	address, tx, instance, err := contract.DeployContract(auth, client, tokenName, tokenSymbol, tokenSupply)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(address, tx, instance)

	return &address, nil
}

func TransferCustomToken(rpcLink string, chainId uint64, fromPrivateKey string, contractAddress string, toAddress string, amount uint64) (*string, error) {
	client, err := GetClient(rpcLink)

	if err != nil {
		return nil, err
	}

	tokenAddress := common.HexToAddress(contractAddress)
	myToken, err := contract.NewContract(tokenAddress, client)

	if err != nil {
		log.Fatal(err)
	}

	toAddressTyped := common.HexToAddress(toAddress)

	privateKey, err := crypto.HexToECDSA(fromPrivateKey)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(chainId)))

	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(15000000)
	auth.GasPrice = gasPrice

	tx, err := myToken.SendToken(auth, toAddressTyped, amount)

	if err != nil {
		return nil, err
	}

	hash := tx.Hash().Hex()

	return &hash, nil
}
