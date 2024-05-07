# node_health_check
Simple healtch checker for nodes.

## Config
`NHC_SIMULATE_HEALTH` can be used to test expected health outcome

`NHC_MAX_BLOCK_LAG` the max block lag before the node is considered unhealthy

`NHC_SELF_URL` the URL to the node under test itself

`NHC_PEER_URLS` comma-separated list of node peers to check block heights against

## Build
`CGO_ENABLED=0 go build -ldflags="-extldflags=-static" -o nhc`
