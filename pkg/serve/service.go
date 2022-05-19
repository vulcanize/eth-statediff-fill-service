// VulcanizeDB
// Copyright © 2022 Vulcanize

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package serve

import (
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/rpc"
)

const (
	PayloadChanBufferSize = 2000
)

// Server is the top level interface for streaming, converting to IPLDs, publishing,
// and indexing all chain data; screening this data; and serving it up to subscribed clients
// This service is compatible with the Ethereum service interface (node.Service)
type Server interface {
	// Start() and Stop()
	APIs() []rpc.API
	Protocols() []p2p.Protocol
}

// Service is the underlying struct for the watcher
type Service struct {
	// rpc client for forwarding cache misses
	client *rpc.Client
}

// NewServer creates a new Server using an underlying Service struct
func NewServer(settings *Config) (Server, error) {
	sap := new(Service)
	sap.client = settings.Client
	return sap, nil
}

// Protocols exports the services p2p protocols, this service has none
func (sap *Service) Protocols() []p2p.Protocol {
	return []p2p.Protocol{}
}

// APIs returns the RPC descriptors the watcher service offers
func (sap *Service) APIs() []rpc.API {
	apis := []rpc.API{
		{
			Namespace: APIName,
			Version:   APIVersion,
			Service:   NewPublicServerAPI(sap.client),
			Public:    true,
		},
	}
	return apis
}
