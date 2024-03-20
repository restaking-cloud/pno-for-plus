# Peferred Node Operator (PNO) Module for MEV Plus

The PNO Liquid Delegation Module for MEV Plus enables liquid delegations in two forms: self-delegation and delegation to another PNO (Pools Node Operator). This module allows node operators to receive liquid delegations from other K2 users if they have PNO status. The module also allows node operators to self-delegate their own validators to the K2 contract pool.

## Tutorial

To follow a tutorial for the installation of the PNO Module, please visit the [PNO Module Tutorial]().

## Prerequisites

- [MEV Plus with PNO]() module plugged in.
Alternatively, you can install the PNO module in your existing MEV Plus software by following the [installation instructions]().

- Eth1 Wallet: You should have an Ethereum wallet with sufficient funds to pay for gas fees. This wallet is used to pay for gas fees when interacting with the K2 and Node Operator contracts.

- MEV Plus connected to node by Builder API or your beacon node URL (without the need for a builder proposals and builder API enabled). This would use the registration messages from the node or epoch changes respectively as a cue to redelegate the validators own by the node operator to the K2 contract pool.

## Supported Networks

- Holesky

## Configuration

You can configure the PNO Module by setting the following flags when running MEV Plus:

### Required Flags

- `pno.eth1-private-key`: The private key of your node operator representative wallet. This key is essential for signing transactions and to configure your lien contract recipient address ad perform redelegation of your ETH capacity in the protocol to yourself or other node operators of your preference.

NOTE: This module also supports the use of multiple representative wallets, to multiple K2 ETH balances and validators to liquid delegate to any other preferred node operator or to yourself if there is capacity to do so.
, in which case provide a comma separated list of private keys. [eg. `pno.eth1-private-key 1234567890abcdef1234567890abcdef12345f,1234567890abcdef1234567890abcdef12345f`
].

- `pno.beacon-node-url`: The URL of the beacon node. This URL is required for syncing with the Ethereum Consensus Layer, and run the module in the background to listen for epoch changes to redelegate the validators owned by the node operator to the K2 contract pool.

- `pno.execution-node-url`: The URL of the execution node to connect to for on-chain execution.

### Optional Flags

- `pno.max-gas-price`: The maximum gas price to be used for on-chain transactions. This flag is optional and defaults to using the current network gas price for execution. If set, any transaction in which the gas price exceeds the maximum gas price, will not be executed.

- `pno.listen-address`: The address on which the module will listen for incoming requests. This flag is optional and defaults to `localhost:9000` if not specified. The API specifications can be found [here](#api).

- `pno.auto-redelegate`: This flag is used to enable or disable the automatic redelegation of validators owned by the node operator to one's self in the K2 contract pool as a node operator of preference. This feature is optional and is only enabled if this flag is set. Auto redelegation would occur when the module is connected to the beacon node and listens for epoch changes, or recieves a validator registration message from your node through the Builder API.

- `pno.k2-lending-contract-address`: The address of the K2 lending contract you wish to provide to override the default contract address for a supported network, or to provide a contract address for an unsupported network.

- `pno.k2-node-operator-contract-address`: The address of the K2 node operator contract you wish to provide to override the default contract address for a supported network, or to provide a contract address for an unsupported network.

## How It Works


## Benefits for Node Runners



## Installation

The PNO module is a plugin for MEV Plus. To install the module, import the module into your MEV Plus `moduleList.go` file and instantiate the `NewPNOService()` and `NewCommand()` from [service.go](service.go) functions in the respective service and command lists variables within the [moduleList.go](https://github.com/pon-network/mev-plus/blob/main/moduleList/moduleList.go) file.

Build your modified MEV Plus binary and run it with the required flags.

For more information on building MEV Plus, check out
[How to build MEV Plus](https://github.com/pon-network/mev-plus#building-mev-plus)

## Usage

The module functionality can be enabled/disabled by running mev-plus with the pno flags available.

```bash

# Enable PNO Module
mevPlus ............ -pno.eth1-private 1234567890abcdef1234567890abcdef12345f -pno.beacon-node-url http://localhost:5052 -pno.execution-node-url http://localhost:8545
```

## License
[MIT](LICENSE.md)
