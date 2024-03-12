package tools

import (
	"fmt"

	"github.com/ybbus/jsonrpc/v3"
)

// CreateProviderClient creates JSON-RPC client for the selected provider.
func CreateProviderClient(url string) jsonrpc.RPCClient {
	return jsonrpc.NewClient(url)
}

// ConvertIntToHex converts given int value to hex string.
func ConvertIntToHex(value uint64) string {
	// 0x56BC75E2D63100000

	return fmt.Sprintf("%#x", value)
}
