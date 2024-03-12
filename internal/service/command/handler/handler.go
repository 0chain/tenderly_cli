package handler

import (
	"github.com/0chain/tenderly_cli/internal/config"
	"github.com/0chain/tenderly_cli/internal/logging"
	"github.com/0chain/tenderly_cli/internal/service/client"
	"github.com/0chain/tenderly_cli/internal/service/command/common"
	"github.com/0chain/tenderly_cli/internal/tools"
)

// SetupSetBalanceCommand initializes set balance command handler.
func SetupSetBalanceCommand() {
	config.GetCommandRoot().AddCommand(
		common.CreateCommand(
			"set-balance",
			"sets Ethereum wallet balance with the given value.",
			"sets Ethereum wallet balance with the given value.",
			func(args ...*common.Arg) {
				host := common.GetHost(args)
				wallet := common.GetWallet(args)
				amount := common.GetAmount(args)

				err := client.SetBalance(
					tools.CreateProviderClient(host),
					wallet,
					tools.ConvertIntToHex(amount))
				if err != nil {
					logging.Logger.Fatal(err.Error())
				}

				logging.Logger.Info("Balance has been set!")
			},
			common.CreateWalletOption(),
			common.CreateAmountOption()))
}

// SetupAddBalanceCommand initializes add balance command handler.
func SetupAddBalanceCommand() {
	config.GetCommandRoot().AddCommand(
		common.CreateCommand(
			"add-balance",
			"adds Ethereum wallet balance with the given value.",
			"adds Ethereum wallet balance with the given value.",
			func(args ...*common.Arg) {
				host := common.GetHost(args)
				wallet := common.GetWallet(args)
				amount := common.GetAmount(args)

				err := client.AddBalance(
					tools.CreateProviderClient(host),
					wallet,
					tools.ConvertIntToHex(amount))
				if err != nil {
					logging.Logger.Fatal(err.Error())
				}

				logging.Logger.Info("Balance has been set!")
			},
			common.CreateWalletOption(),
			common.CreateAmountOption()))
}

// SetupSetErc20BalanceCommand initializes set erc20 balance command handler.
func SetupSetErc20BalanceCommand() {
	config.GetCommandRoot().AddCommand(
		common.CreateCommand(
			"set-erc20-balance",
			"sets Ethereum ERC20 wallet balance for the selected token with the given value.",
			"sets Ethereum ERC20 wallet balance for the selected token with the given value.",
			func(args ...*common.Arg) {
				host := common.GetHost(args)
				wallet := common.GetWallet(args)
				token := common.GetToken(args)
				amount := common.GetAmount(args)

				err := client.SetErc20Balance(
					tools.CreateProviderClient(host),
					token,
					wallet,
					tools.ConvertIntToHex(amount))
				if err != nil {
					logging.Logger.Fatal(err.Error())
				}

				logging.Logger.Info("Balance has been set!")
			},
			common.CreateWalletOption(),
			common.CreateTokenOption(),
			common.CreateAmountOption()))
}

// SetupAddErc20BalanceCommand initializes add erc20 balance command handler.
func SetupAddErc20BalanceCommand() {
	config.GetCommandRoot().AddCommand(
		common.CreateCommand(
			"add-erc20-balance",
			"adds Ethereum ERC20 wallet balance for the selected token with the given value.",
			"adds Ethereum ERC20 wallet balance for the selected token with the given value.",
			func(args ...*common.Arg) {
				host := common.GetHost(args)
				wallet := common.GetWallet(args)
				token := common.GetToken(args)
				amount := common.GetAmount(args)

				err := client.AddErc20Balance(
					tools.CreateProviderClient(host),
					token,
					wallet,
					tools.ConvertIntToHex(amount))
				if err != nil {
					logging.Logger.Fatal(err.Error())
				}

				logging.Logger.Info("Balance has been set!")
			},
			common.CreateWalletOption(),
			common.CreateTokenOption(),
			common.CreateAmountOption()))
}
