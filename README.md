# node_health_check
Simple healtch checker for nodes.

## Config
`HC_SIMULATE_HEALTH` can be used to test expected health outcome

`HC_MAX_BLOCK_LAG` the max block lag before the node is considered unhealthy

`HC_SELF_URL` the URL to the node under test itself

`HC_PEER_URLS` comma-separated list of node peers to check block heights against

## Build
`CGO_ENABLED=0 go build -ldflags="-extldflags=-static" -o nhc`
