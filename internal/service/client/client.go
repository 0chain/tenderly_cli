package client

import (
	"context"

	"github.com/ybbus/jsonrpc/v3"

	"github.com/pkg/errors"
)

var (
	ErrTenderlyRequest = errors.New("err happened during external Tenderly request")
)

// SetBalance sets initial balance for the given ethereum address with the given value.
// Given value should be in HEX based interpretation.
func SetBalance(client jsonrpc.RPCClient, wallet, value string) error {
	resp, err := client.Call(
		context.Background(), "tenderly_setBalance", []string{wallet}, value)
	if err != nil {
		return errors.Wrap(err, ErrTenderlyRequest.Error())
	}

	if resp.Error != nil {
		return errors.New(resp.Error.Error())
	}

	return nil
}

// AddBalance adds balance for the given ethereum address with the given value.
// Given value should be in HEX based interpretation.
func AddBalance(client jsonrpc.RPCClient, wallet, value string) error {
	resp, err := client.Call(
		context.Background(), "tenderly_addBalance", []string{wallet}, value)
	if err != nil {
		return errors.Wrap(err, ErrTenderlyRequest.Error())
	}

	if resp.Error != nil {
		return errors.New(resp.Error.Error())
	}

	return nil
}

// SetErc20Balance sets initial balance for the given ethereum address with the given value.
// Given value should be in HEX based interpretation.
func SetErc20Balance(client jsonrpc.RPCClient, token, wallet, value string) error {
	resp, err := client.Call(
		context.Background(), "tenderly_setErc20Balance", token, wallet, value)
	if err != nil {
		return errors.Wrap(err, ErrTenderlyRequest.Error())
	}

	if resp.Error != nil {
		return errors.Wrap(errors.New(resp.Error.Error()), ErrTenderlyRequest.Error())
	}

	return nil
}

// AddErc20Balance adds initial balance for the given ethereum address with the given value.
// Given value should be in HEX based interpretation.
func AddErc20Balance(client jsonrpc.RPCClient, token, wallet, value string) error {
	resp, err := client.Call(
		context.Background(), "tenderly_addErc20Balance", token, wallet, value)
	if err != nil {
		return errors.Wrap(err, ErrTenderlyRequest.Error())
	}

	if resp.Error != nil {
		return errors.Wrap(errors.New(resp.Error.Error()), ErrTenderlyRequest.Error())
	}

	return nil
}
