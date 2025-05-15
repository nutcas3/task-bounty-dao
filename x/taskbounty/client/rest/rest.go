package rest

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"
)

// RegisterRoutes registers taskbounty-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	registerQueryRoutes(clientCtx, r)
	registerTxRoutes(clientCtx, r)
}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// Query endpoints
	r.HandleFunc("/taskbounty/tasks", listTasksHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/taskbounty/tasks/{id}", getTaskHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/taskbounty/wallets/{address}/balance", getBalanceHandler(clientCtx)).Methods("GET")
}

func registerTxRoutes(clientCtx client.Context, r *mux.Router) {
	// Transaction endpoints
	r.HandleFunc("/taskbounty/tasks", createTaskHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/taskbounty/tasks/{id}/claim", claimTaskHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/taskbounty/tasks/{id}/proof", submitProofHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/taskbounty/tasks/{id}/approve", approveTaskHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/taskbounty/wallets", createWalletHandler(clientCtx)).Methods("POST")
}
