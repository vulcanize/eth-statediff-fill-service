#!/bin/sh

echo "Beginning the eth-statediff-fill-service process"

echo running: ./eth-statediff-fill-service ${VDB_COMMAND} --config=config.toml
./eth-statediff-fill-service ${VDB_COMMAND} --config=config.toml
rv=$?

if [ $rv != 0 ]; then
  echo "eth-statediff-fill-service startup failed"
  exit 1
fi
