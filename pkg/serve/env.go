package serve

import (
	"github.com/ethereum/go-ethereum/rpc"
)

// Env variables
const (
	ETH_WS_PATH   = "ETH_WS_PATH"
	ETH_HTTP_PATH = "ETH_HTTP_PATH"

	DATABASE_NAME                 = "DATABASE_NAME"
	DATABASE_HOSTNAME             = "DATABASE_HOSTNAME"
	DATABASE_PORT                 = "DATABASE_PORT"
	DATABASE_USER                 = "DATABASE_USER"
	DATABASE_PASSWORD             = "DATABASE_PASSWORD"
	DATABASE_MAX_IDLE_CONNECTIONS = "DATABASE_MAX_IDLE_CONNECTIONS"
	DATABASE_MAX_OPEN_CONNECTIONS = "DATABASE_MAX_OPEN_CONNECTIONS"
	DATABASE_MAX_CONN_LIFETIME    = "DATABASE_MAX_CONN_LIFETIME"
)

// GetEthClient returns client from path url
func getEthClient(path string) (*rpc.Client, error) {
	rpcClient, err := rpc.Dial(path)
	if err != nil {
		return nil, err
	}

	return rpcClient, nil
}
