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

	// client.InitErc20Balance("0x0E85ef63337299eCF06dEBa43a61ce8B5982B38e")

	// client.InitBalance("0xeca881a5f23c62b75f76390Ee8aBFA441F681A0e", 7766279631452241920)

	command.Run()
}

// func main() {
// 	client := NewClient("https://rpc.tenderly.co/fork/835ecb4e-1f60-4329-adc2-b0c741193830")

// 	client.InitErc20Balance("0x0E85ef63337299eCF06dEBa43a61ce8B5982B38e")
// 	client.InitBalance("0x0E85ef63337299eCF06dEBa43a61ce8B5982B38e")
// 	// fmt.Println(client.LatestBlock())
// }
