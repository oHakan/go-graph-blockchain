# oHakan
## GoLang Blockchain Smart Contract Deployer

## Features

- You can deploy your contract with repository.
- Compile your .abi and .bin file with abigen under the "Contract" folder.
- Update your schema according to your contract constructor parameters.



This application is schema-based. For this reason, you must update only .graphqls schemas. After, you can run this command;
```sh
go run github.com/99designs/gqlgen generate
```
It's will create resolver and model according to your schema.

>If you want deploy your contract, you should use "abigen" for generate .go file generated from your solidity contract .abi and .bin file.

You can run this program. Functions required RPC Provider Link and ChainID. It's important. With this parameters, contract will deploy to your selected blockchain network. 

For an example RPC Link and ChainID for Avalanche Testnet.

> "https://api.avax-test.network/ext/bc/C/rpc"
> 43113

You can create your wallet with this program. 
```
mutation {createWallet(
input: {name: "test"}
) {
  publicKey
  privateKey
  address
}}
```
You can transfer your main chain network with this query.
```
mutation {
  transferToken(
    input: {
    rpcLink: "https://api.avax-test.network/ext/bc/C/rpc", 
    fromPrivate: "YOUR_PRIVATE_KEY", 
    toPublic: "0x87578Bf088AC0198d3bfab5b54Ce70360aEd2457", 
    amount: "0.01"
    }
  )
}
```

You can deploy contract with this query. 
```
mutation {
  deployContract(
  input: {
    rpcLink: "https://api.avax-test.network/ext/bc/C/rpc",
    name: "oHakan Test Token",
    symbol: "OTT",
    supply: 1000,
    privateKey: "YOUR_PRIVATE_KEY",
    chainId: 43113
      }
  )
}
```

With default contract (ERC-20), create your own token on network. 

You can send your token with this query to another wallet;
```
mutation {
  transferCustomToken(
  input: {
    rpcLink: "https://api.avax-test.network/ext/bc/C/rpc",
    amount: 10,
    chainId: 43113,
    toAddress: "0x339aF7605c00dC892D71141d6Fa6C87c49a4fFbd",
    contractAddress: "0x404e58f19F311d7C4539872534fA6B71b597401f",
    fromAddress: "YOUR_PRIVATE_KEY"
      }
  )
}
```

Regards, oHakan.