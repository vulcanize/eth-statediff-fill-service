set -e
set -o xtrace

export WATCHED_ADDRESS_GAP_FILLER_INTERVAL=5

# Wait for containers to be up and execute the integration test.
while [ "$(curl -s -o /dev/null -w ''%{http_code}'' localhost:8081)" != "200" ]; do echo "waiting for ipld-eth-server..." && sleep 5; done && \
while [ "$(curl -s -o /dev/null -w ''%{http_code}'' localhost:8545)" != "200" ]; do echo "waiting for geth-statediff..." && sleep 5; done && \
make integrationtest
