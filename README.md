# Tokenized Task Bounty System

A blockchain-powered decentralized task reward platform built with Go and Cosmos SDK, enabling users to create, claim, and complete tasks with token bounties.

## Core Features

### Task Management
- Create tasks with token bounties
- List tasks (open/closed)
- Claim tasks
- Submit proof of work
- Approve and release bounties

### Smart Contract Features
- Token management
- Task state management
- Bounty release logic
- On-chain metadata storage

### User Interaction
- REST API for task operations
- CLI commands for blockchain interaction
- Wallet integration

## Project Structure

```
.
├── app/                    # Main application
├── cmd/                    # Application entry points
│   └── daod/               # Daemon application
├── x/                      # Cosmos SDK modules
│   └── taskbounty/         # Core module
│       ├── client/         # CLI and REST handlers
│       ├── keeper/         # Business logic
│       └── types/          # Data structures and messages
├── proto/                  # Protocol buffer definitions
│   └── amino/              # Amino codec definitions
├── config/                 # Configuration files
│   └── network.toml        # Network configuration
└── internal/               # Private application code
    ├── types/              # Custom types and interfaces
    └── handlers/           # Request handlers
```

## Implementation Phases

### 1. Smart Contract (Cosmos SDK Module)
- State management (Task data structure)
- Message types (CreateTask, ClaimTask, SubmitProof, ApproveTask)
- Keeper logic implementation
- Module registration and endpoints

### 2. Go API/CLI
- Generated client code
- Custom API logic
- CLI commands
- Blockchain integration

### 3. User Interaction and Data Storage
- Task creation/management
- Task listing and querying
- Proof submission
- Bounty release
- On-chain metadata storage

### 4. Stretch Goals
- Role-based access control
- SERV wallet integration
- Know Freedom testnet integration
- Comprehensive tests and documentation

## Integration Requirements

### Know Freedom Testnet
- RPC Endpoint: 
- gRPC Endpoint: 
- Chain ID: To be determined
- Native Token: To be determined

### Key Integration Points
1. Network Configuration
2. Token Denomination Usage
3. Wallet Compatibility
4. Transaction Handling

## Prerequisites

- Go 1.20 or later
- Cosmos SDK
- Access to Know Freedom testnet

## Getting Started

1. Clone the repository
2. Install dependencies: `go mod tidy`
3. Configure network settings in `config/network.toml`
4. Run the application: `go run cmd/daod/main.go`

## API Documentation

The API provides endpoints for:
- Task creation and management
- Task listing and querying
- Proof submission
- Bounty release

Detailed API documentation will be available after initial implementation.

## Testing

Run tests with: `go test ./...`

See `testing.md` for more detailed testing procedures and scenarios.

## License

MIT
