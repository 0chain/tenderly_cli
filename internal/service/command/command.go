package command

import (
	"github.com/0chain/tenderly_cli/internal/config"
	"github.com/0chain/tenderly_cli/internal/service/command/handler"
)

// Run starts command root handler.
func Run() error {
	handler.SetupSetBalanceCommand()
	handler.SetupAddBalanceCommand()
	handler.SetupSetErc20BalanceCommand()
	handler.SetupAddErc20BalanceCommand()

	return config.GetCommandRoot().Execute()
}
