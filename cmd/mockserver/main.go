package main

import (
	"encoding/json"
	// "fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

// Task represents a task in the system
type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Reward      string    `json:"reward"`
	Creator     string    `json:"creator"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	Deadline    time.Time `json:"deadline,omitempty"`
	Claimer     string    `json:"claimer,omitempty"`
	CompletedAt time.Time `json:"completed_at,omitempty"`
}

// TasksResponse represents the response for the tasks endpoint
type TasksResponse struct {
	Tasks []Task `json:"tasks"`
}

// TaskResponse represents the response for a single task endpoint
type TaskResponse struct {
	Task Task `json:"task"`
}

// BalanceResponse represents the response for the balance endpoint
type BalanceResponse struct {
	Balance struct {
		Address string `json:"address"`
		Coins   []Coin `json:"coins"`
	} `json:"balance"`
}

// Coin represents a coin in the system
type Coin struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

func main() {
	log.Println("Starting Mock Task Bounty Server...")

	// Initialize router
	router := mux.NewRouter()

	// Register routes
	router.HandleFunc("/taskbounty/tasks", listTasksHandler).Methods("GET")
	router.HandleFunc("/taskbounty/tasks/{id}", getTaskHandler).Methods("GET")
	router.HandleFunc("/taskbounty/wallets/{address}/balance", getBalanceHandler).Methods("GET")

	// Create HTTP server
	srv := &http.Server{
		Addr:    ":1317", // Standard Cosmos SDK REST port
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Mock server listening on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down gracefully...")
}

// listTasksHandler returns a list of mock tasks
func listTasksHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling request for /taskbounty/tasks")

	// Create mock tasks
	tasks := []Task{
		{
			ID:          "task-1",
			Title:       "Implement REST API",
			Description: "Create REST endpoints for the taskbounty module",
			Reward:      "100uknow",
			Creator:     "know1...",
			Status:      "open",
			CreatedAt:   time.Date(2025, 5, 15, 10, 0, 0, 0, time.UTC),
			Deadline:    time.Date(2025, 5, 30, 10, 0, 0, 0, time.UTC),
		},
		{
			ID:          "task-2",
			Title:       "Design UI",
			Description: "Design a user interface for the taskbounty dApp",
			Reward:      "200uknow",
			Creator:     "know1...",
			Status:      "claimed",
			CreatedAt:   time.Date(2025, 5, 14, 15, 30, 0, 0, time.UTC),
			Deadline:    time.Date(2025, 5, 28, 15, 30, 0, 0, time.UTC),
			Claimer:     "know2...",
		},
		{
			ID:          "task-3",
			Title:       "Write tests",
			Description: "Write unit tests for the taskbounty module",
			Reward:      "150uknow",
			Creator:     "know1...",
			Status:      "completed",
			CreatedAt:   time.Date(2025, 5, 13, 9, 15, 0, 0, time.UTC),
			CompletedAt: time.Date(2025, 5, 15, 14, 20, 0, 0, time.UTC),
			Claimer:     "know3...",
		},
	}

	// Create response
	response := TasksResponse{
		Tasks: tasks,
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// getTaskHandler returns a mock task by ID
func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["id"]

	log.Printf("Handling request for /taskbounty/tasks/%s", taskID)

	// Create mock tasks map
	tasks := map[string]Task{
		"task-1": {
			ID:          "task-1",
			Title:       "Implement REST API",
			Description: "Create REST endpoints for the taskbounty module",
			Reward:      "100uknow",
			Creator:     "know1...",
			Status:      "open",
			CreatedAt:   time.Date(2025, 5, 15, 10, 0, 0, 0, time.UTC),
			Deadline:    time.Date(2025, 5, 30, 10, 0, 0, 0, time.UTC),
		},
		"task-2": {
			ID:          "task-2",
			Title:       "Design UI",
			Description: "Design a user interface for the taskbounty dApp",
			Reward:      "200uknow",
			Creator:     "know1...",
			Status:      "claimed",
			CreatedAt:   time.Date(2025, 5, 14, 15, 30, 0, 0, time.UTC),
			Deadline:    time.Date(2025, 5, 28, 15, 30, 0, 0, time.UTC),
			Claimer:     "know2...",
		},
		"task-3": {
			ID:          "task-3",
			Title:       "Write tests",
			Description: "Write unit tests for the taskbounty module",
			Reward:      "150uknow",
			Creator:     "know1...",
			Status:      "completed",
			CreatedAt:   time.Date(2025, 5, 13, 9, 15, 0, 0, time.UTC),
			CompletedAt: time.Date(2025, 5, 15, 14, 20, 0, 0, time.UTC),
			Claimer:     "know3...",
		},
	}

	// Check if task exists
	task, exists := tasks[taskID]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Task not found"})
		return
	}

	// Create response
	response := TaskResponse{
		Task: task,
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// getBalanceHandler returns a mock balance for an address
func getBalanceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]

	log.Printf("Handling request for /taskbounty/wallets/%s/balance", address)

	// Create mock balances map
	balances := map[string]string{
		"know1...": "10000000",
		"know2...": "5000000",
		"know3...": "7500000",
	}

	// Get balance or use default
	amount := "1000000" // Default amount
	if val, exists := balances[address]; exists {
		amount = val
	}

	// Create response
	response := BalanceResponse{}
	response.Balance.Address = address
	response.Balance.Coins = []Coin{
		{
			Denom:  "uknow",
			Amount: amount,
		},
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
