// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateUserInput struct {
	Name     *string  `json:"name,omitempty"`
	LastName *string  `json:"lastName,omitempty"`
	Gender   *float64 `json:"gender,omitempty"`
	Email    *string  `json:"email,omitempty"`
	Password *string  `json:"password,omitempty"`
}

type CreateWalletInput struct {
	Name *string `json:"name,omitempty"`
}

type DeployContractInput struct {
	Name       string  `json:"name"`
	Symbol     string  `json:"symbol"`
	Supply     float64 `json:"supply"`
	RPCLink    string  `json:"rpcLink"`
	ChainID    float64 `json:"chainId"`
	PrivateKey string  `json:"privateKey"`
}

type TransferCustomTokenInput struct {
	FromAddress     string  `json:"fromAddress"`
	ContractAddress string  `json:"contractAddress"`
	RPCLink         string  `json:"rpcLink"`
	ChainID         float64 `json:"chainId"`
	ToAddress       string  `json:"toAddress"`
	Amount          float64 `json:"amount"`
}

type TransferToken struct {
	RPCLink     string `json:"rpcLink"`
	Amount      string `json:"amount"`
	FromPrivate string `json:"fromPrivate"`
	ToPublic    string `json:"toPublic"`
}

type User struct {
	Name     string  `json:"name"`
	LastName string  `json:"lastName"`
	Gender   float64 `json:"gender"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
}

type Wallet struct {
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"privateKey"`
	Address    string `json:"address"`
}
