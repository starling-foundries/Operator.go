# The Awesome Golang Json-RPC starter Kit for Zilliqa Hybrid DAPPs

A hybrid dapp is an application with some critical logic executed from a distributed context (for example on a zilliqa smart contract) and some off-chain. This off-chain logic could be a gasless transaction server, a stamping signature authority, etc. We want to maintain the greatest level of interoperability with the existing Zilliqa tools so we will use a JSON-RPC 2.0 server.

## Twirp

Twirp allows us to use protobuf `.proto` files and still operate over HTTP 1.1 and 2.0 without proxy. This makes it a "better gRPC" for our purposes as it does not immediately mandate an Envoy proxy and is more serverless-friendly. Twirp's automatic generation makes it easy to add or modify functionality around the core purpose of this server by simply adding or changing the definition files and regenerating the stub codes. 

## Gasless transactions

Many smart contract platforms require all calls to their chaincodes to be paid for in one core denomination of gas. For Ethereum it is only ETH (and not ERC-20 tokens, for example). This creates a KYC and full-wallet requirement as new users must get ETH or ZIL from an exchange even to interact solely with a token contract. Now, with this bouncer/operator pattern the transition fee can be paid in the contract's *native token!* This is made possible by bundling and signing a transaction and then passing it to a semi-trusted off-chain server that signs and encapsulates that transaction - paying gas and perhaps taking a token-denominated fee in return. 

## Smart Contract

This pattern is best demonstrated securely in the ERC-865 reference. There is a corresponding Zilliqa reference used here, ZRC-4. 

## Demonstration

You might like to test the live server with this online json-rpc testing tool [Guru JSON-RPC Tester](https://gurujsonrpc.appspot.com/). 

### Local installation

Make sure you have twirp, protoc, etc installed.
```
// Install Twirp protoc generator
go get github.com/twitchtv/twirp/protoc-gen-twirp

// Install Protoc - The protobuf compiler
https://github.com/google/protobuf/releases/tag/v3.6.0

// Mac users can install the above with HomeBrew
brew update && brew install protobuf

// Install Golang Protobuf support
go get -u github.com/golang/protobuf/protoc-gen-go
```

## License: 
 Copyright Cameron Sajedi, 2019 All rights reserved
