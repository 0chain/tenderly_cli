# TenderlyCLI

This repository contains the implementation of Tenderly CLI, used to simplify interaction with Tenderly API, intended to expose basic comamnds.

## Setup

It's required to execute the following command to build the node:

```bash
make build
```

The built file is located at **./build** directory.

## Examples

The following command represents request to Tenderly API, which provokes balance to be set to be equal to the given value:
```shell
tenderly_cli set-balance --host="https://rpc.tenderly.co/fork/34689cb9-e797-4786-8a1e-9512e3183222" --wallet="0xfda881a5f23c62b75f76390Ee8aBFA441F681A0e" --amount=10
```

The following command represents request to Tenderly API, which provokes balance to be increased by the given value:
```shell
tenderly_cli add-balance --host="https://rpc.tenderly.co/fork/34689cb9-e797-4786-8a1e-9512e3183222" --wallet="0xfda881a5f23c62b75f76390Ee8aBFA441F681A0e" --amount=10
```

The following command represents request to Tenderly API, which provokes balance of the selected ERC20 smart contract to be set to be equal to the given value:
```shell
tenderly_cli set-erc20-balance --host="https://rpc.tenderly.co/fork/34689cb9-e797-4786-8a1e-9512e3183222" --wallet="0xfda881a5f23c62b75f76390Ee8aBFA441F681A0e" --token="0xb9EF770B6A5e12E45983C5D80545258aA38F3B78" --amount=10
```

The following command represents request to Tenderly API, which provokes balance of the selected ERC20 smart contract to be increased by the given value:
```shell
tenderly_cli add-erc20-balance --host="https://rpc.tenderly.co/fork/34689cb9-e797-4786-8a1e-9512e3183222" --wallet="0xfda881a5f23c62b75f76390Ee8aBFA441F681A0e" --token="0xb9EF770B6A5e12E45983C5D80545258aA38F3B78" --amount=10
```

> Remember that the value expected to be in **ETH**