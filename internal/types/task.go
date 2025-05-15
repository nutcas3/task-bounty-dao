package types

import (
	"time"
)

// TaskStatus represents the current state of a task
type TaskStatus string

const (
	TaskStatusOpen     TaskStatus = "open"
	TaskStatusClaimed  TaskStatus = "claimed"
	TaskStatusComplete TaskStatus = "complete"
)

// Task represents a bounty task in the system
type Task struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Bounty      uint64     `json:"bounty"`      // Token amount
	Status      TaskStatus `json:"status"`
	CreatorID   string     `json:"creator_id"`  // Creator's wallet address
	ClaimerID   string     `json:"claimer_id"`  // Claimer's wallet address
	ProofOfWork string     `json:"proof"`       // Link or hash of the proof
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// CreateTaskRequest represents the request to create a new task
type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Bounty      uint64 `json:"bounty"`
}

// ClaimTaskRequest represents the request to claim a task
type ClaimTaskRequest struct {
	TaskID    string `json:"task_id"`
	ClaimerID string `json:"claimer_id"`
}

// SubmitProofRequest represents the request to submit proof of work
type SubmitProofRequest struct {
	TaskID      string `json:"task_id"`
	ProofOfWork string `json:"proof"`
}

// ApproveTaskRequest represents the request to approve and release bounty
type ApproveTaskRequest struct {
	TaskID string `json:"task_id"`
}
