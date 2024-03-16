// nolint:typecheck
package main

import (
	"github.com/0chain/tenderly_cli/internal/logging"
	"github.com/0chain/tenderly_cli/internal/service/command"
)

func init() {
	logging.Init()
}

func main() {
	command.Run()
}
