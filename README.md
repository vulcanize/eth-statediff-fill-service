# eth-statediff-fill-service
Service to fill statediff gaps for watched addresses

## Setup

Build the binary:

```bash
make build
```

### Local Setup

* Create `config.toml` in [environments](./environments/) from [example.toml](./environments/example.toml).

* Set the necessary params in `config.toml` file
  * The `database` fields are for connecting to a Postgres database.
  * The `server` fields set the paths for exposing the eth-statediff-fill-service endpoints.
  * The `ethereum` can be used to configure the remote eth node for making write_stateDiff calls.

## Usage

### `serve`

To run the service:

`eth-statediff-fill-service serve --config=<config path>`

Example:

```bash
./eth-statediff-fill-service serve --config environments/config.toml
```

Available RPC method: `vdb_watchAddress()`

e.g. `curl -X POST -H 'Content-Type: application/json' --data '{"jsonrpc":"2.0","method":"vdb_watchAddress","params":["add",[{"Address":"0x825a6eec09e44Cb0fa19b84353ad0f7858d7F61a","CreatedAt":22}]],"id":1}' "$HOST":"$PORT"`

## Tests

To run tests follow [readme](./test/README.md).
