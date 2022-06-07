# Test Instructions

## Setup

- Clone [stack-orchestrator](https://github.com/vulcanize/stack-orchestrator), [go-ethereum](https://github.com/vulcanize/go-ethereum) and [ipld-eth-server](https://github.com/vulcanize/ipld-eth-server) repositories.

- Checkout [v3 release](https://github.com/vulcanize/go-ethereum/releases/tag/v1.10.18-statediff-3.2.2) in go-ethereum repo.
  ```bash
  # In go-ethereum repo.
  git checkout v1.10.18-statediff-3.2.2
  ```

- Checkout [v3 release](https://github.com/vulcanize/ipld-eth-server/releases/tag/v3.2.1) in ipld-eth-server repo.
  ```bash
  # In ipld-eth-server repo.
  git checkout v3.2.1
  ```

- Checkout working commit in stack-orchestrator repo.
  ```bash
  # In stack-orchestrator repo.
  git checkout e10025903ee8f2e7379cb89724f8f5b9f4dccc7b
  ```

## Run

- Run unit tests:

  ```bash
  # In eth-statediff-fill-service root directory.
  ./scripts/run_unit_test.sh
  ```

- Run integration tests:
  - Create config file in stack-orchestrator repo.
    ```bash
    cd stack-orchestrator/helper-scripts
    ./create-config.sh
    ```

  - Update (Replace existing content) the generated config file `config.sh` in stack-orchestrator repo:
    ```bash
    #!/bin/bash

    # Path to go-ethereum repo.
    vulcanize_go_ethereum=~/vulcanize/go-ethereum/

    # Path to ipld-eth-server repo.
    vulcanize_ipld_eth_server=~/vulcanize/ipld-eth-server/

    # Path to docker for test contract.
    vulcanize_test_contract=~/vulcanize/ipld-eth-server/test/contract

    # Path to eth-statediff-fill-service repo.
    vulcanize_eth_statediff_fill_service=~/vulcanize/eth-statediff-fill-service/

    db_write=true
    eth_forward_eth_calls=false
    eth_proxy_on_error=false
    eth_http_path="go-ethereum:8545"
    watched_address_gap_filler_interval=5
    ```

  - Run stack-orchestrator:
    ```bash
    # In stack-orchestrator root directory.
    cd helper-scripts

    ./wrapper.sh \
    -e docker \
    -d ../docker/latest/docker-compose-db.yml \
    -d ../docker/local/docker-compose-go-ethereum.yml \
    -d ../docker/local/docker-compose-ipld-eth-server.yml \
    -d ../docker/local/docker-compose-contract.yml \
    -d ../docker/local/docker-compose-eth-statediff-fill-service.yml \
    -v remove \
    -p ../config.sh
    ```

  - Run test:
    ```bash
    # In ipld-eth-server root directory.
    ./scripts/run_integration_test.sh
    ```
