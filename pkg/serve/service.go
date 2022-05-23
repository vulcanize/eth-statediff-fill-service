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
	"sync"

	ethnode "github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/rpc"
	log "github.com/sirupsen/logrus"
)

// Server is the top level interface for exposing remote RPC API
type Server interface {
	ethnode.Lifecycle
	APIs() []rpc.API
	Protocols() []p2p.Protocol
	Serve(wg *sync.WaitGroup)
}

// Service is the underlying struct for the service
type Service struct {
	wg *sync.WaitGroup
	// rpc client for forwarding to geth
	client   *rpc.Client
	quitChan chan struct{}
}

// NewServer creates a new Server using an underlying Service struct
func NewServer(settings *Config) (Server, error) {
	sap := new(Service)
	sap.quitChan = make(chan struct{})
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

// Serve is the listening loop
func (sap *Service) Serve(wg *sync.WaitGroup) {
	sap.wg = wg
	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			select {
			case <-sap.quitChan:
				log.Info("quitting the server process")
				return
			}
		}
	}()
	log.Info("server process successfully spun up")
}

// Start is used to begin the service
// This is mostly just to satisfy the node.Service interface
func (sap *Service) Start() error {
	log.Info("starting server")
	wg := new(sync.WaitGroup)
	sap.Serve(wg)
	return nil
}

// Stop is used to close down the service
// This is mostly just to satisfy the node.Service interface
func (sap *Service) Stop() error {
	log.Infof("stopping server")
	close(sap.quitChan)
	return nil
}
