#!/bin/bash

# Create directories
mkdir -p proto/amino
mkdir -p proto/cosmos_proto
mkdir -p proto/gogoproto
mkdir -p proto/google/api
mkdir -p proto/google/protobuf
mkdir -p proto/cometbft/abci/v1
mkdir -p proto/cometbft/crypto/v1
mkdir -p proto/cometbft/types/v1

# Download required proto files
curl -o proto/amino/amino.proto https://raw.githubusercontent.com/cosmos/cosmos-sdk/main/proto/amino/amino.proto
curl -o proto/cosmos_proto/cosmos.proto https://raw.githubusercontent.com/cosmos/cosmos-proto/main/proto/cosmos_proto/cosmos.proto
curl -o proto/gogoproto/gogo.proto https://raw.githubusercontent.com/cosmos/gogoproto/main/gogoproto/gogo.proto
curl -o proto/google/api/annotations.proto https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto
curl -o proto/google/api/http.proto https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto
curl -o proto/google/protobuf/descriptor.proto https://raw.githubusercontent.com/protocolbuffers/protobuf/main/src/google/protobuf/descriptor.proto
curl -o proto/cometbft/abci/v1/types.proto https://raw.githubusercontent.com/cometbft/cometbft/main/proto/cometbft/abci/v1/types.proto
curl -o proto/cometbft/crypto/v1/keys.proto https://raw.githubusercontent.com/cometbft/cometbft/main/proto/cometbft/crypto/v1/keys.proto
curl -o proto/cometbft/crypto/v1/proof.proto https://raw.githubusercontent.com/cometbft/cometbft/main/proto/cometbft/crypto/v1/proof.proto
curl -o proto/cometbft/types/v1/types.proto https://raw.githubusercontent.com/cometbft/cometbft/main/proto/cometbft/types/v1/types.proto
curl -o proto/cometbft/types/v1/validator.proto https://raw.githubusercontent.com/cometbft/cometbft/main/proto/cometbft/types/v1/validator.proto
curl -o proto/cometbft/types/v1/params.proto https://raw.githubusercontent.com/cometbft/cometbft/main/proto/cometbft/types/v1/params.proto
curl -o proto/cometbft/types/v1/block.proto https://raw.githubusercontent.com/cometbft/cometbft/main/proto/cometbft/types/v1/block.proto
curl -o proto/cometbft/types/v1/evidence.proto https://raw.githubusercontent.com/cometbft/cometbft/main/proto/cometbft/types/v1/evidence.proto

# P2P protos
mkdir -p proto/cometbft/p2p/v1
curl -o proto/cometbft/p2p/v1/types.proto https://raw.githubusercontent.com/cometbft/cometbft/main/proto/cometbft/p2p/v1/types.proto
curl -o proto/cometbft/p2p/v1/conn.proto https://raw.githubusercontent.com/cometbft/cometbft/main/proto/cometbft/p2p/v1/conn.proto
curl -o proto/cometbft/p2p/v1/pex.proto https://raw.githubusercontent.com/cometbft/cometbft/main/proto/cometbft/p2p/v1/pex.proto
