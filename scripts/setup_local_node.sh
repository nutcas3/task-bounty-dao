#!/bin/bash
set -e

# Define variables
CHAIN_ID="taskbounty-local"
KEYRING_BACKEND="test"
MONIKER="taskbounty-node"
HOME_DIR="$HOME/.taskbounty"
KEY_NAME="validator"
STAKE_DENOM="uknow"

# Create directories if they don't exist
mkdir -p $HOME_DIR

echo "Setting up local Cosmos SDK blockchain with taskbounty module..."

# Initialize the chain
echo "Initializing chain..."
./daod init $MONIKER --chain-id $CHAIN_ID --home $HOME_DIR

# Create a key for the validator
echo "Creating validator key..."
./daod keys add $KEY_NAME --keyring-backend $KEYRING_BACKEND --home $HOME_DIR

# Get the validator address
VALIDATOR_ADDRESS=$(./daod keys show $KEY_NAME -a --keyring-backend $KEYRING_BACKEND --home $HOME_DIR)
echo "Validator address: $VALIDATOR_ADDRESS"

# Add genesis account
echo "Adding genesis account..."
./daod add-genesis-account $VALIDATOR_ADDRESS 100000000000$STAKE_DENOM --home $HOME_DIR

# Create genesis transaction
echo "Creating genesis transaction..."
./daod gentx $KEY_NAME 10000000000$STAKE_DENOM --chain-id $CHAIN_ID --keyring-backend $KEYRING_BACKEND --home $HOME_DIR

# Collect genesis transactions
echo "Collecting genesis transactions..."
./daod collect-gentxs --home $HOME_DIR

# Validate genesis file
echo "Validating genesis file..."
./daod validate-genesis --home $HOME_DIR

echo "Local blockchain setup complete!"
echo "To start the node, run: ./daod start --home $HOME_DIR"
