#!/bin/bash
set -e

# Define variables
HOME_DIR="$HOME/.taskbounty"
RPC_PORT=26657
P2P_PORT=26656
API_PORT=1317
GRPC_PORT=9090

echo "Starting local Cosmos SDK blockchain with taskbounty module..."

# Start the node with specific ports to avoid conflicts
./daod start \
  --home $HOME_DIR \
  --rpc.laddr tcp://0.0.0.0:$RPC_PORT \
  --p2p.laddr tcp://0.0.0.0:$P2P_PORT \
  --api.address tcp://0.0.0.0:$API_PORT \
  --grpc.address 0.0.0.0:$GRPC_PORT \
  --log_level info
