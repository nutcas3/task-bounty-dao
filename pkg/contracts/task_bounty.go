package contracts

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

// TaskStatus represents the current state of a task
type TaskStatus int

const (
	TaskStatusOpen TaskStatus = iota
	TaskStatusClaimed
	TaskStatusCompleted
	TaskStatusApproved
)

// Task represents a bounty task
type Task struct {
	ID          uint64     `json:"id"`
	Creator     string     `json:"creator"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Bounty      sdk.Coins `json:"bounty"`
	Status      TaskStatus `json:"status"`
	Claimant    string     `json:"claimant,omitempty"`
	ProofURI    string     `json:"proof_uri,omitempty"`
	CreatedAt   int64      `json:"created_at"`
	UpdatedAt   int64      `json:"updated_at"`
}

// TaskBountyContract manages the task bounty system
type TaskBountyContract struct {
	tasks    map[uint64]*Task
	nextID   uint64
	keeper   bankkeeper.Keeper
	ctx      sdk.Context
}

// NewTaskBountyContract creates a new instance of TaskBountyContract
func NewTaskBountyContract(keeper bankkeeper.Keeper, ctx sdk.Context) *TaskBountyContract {
	return &TaskBountyContract{
		tasks:    make(map[uint64]*Task),
		nextID:   1,
		keeper:   keeper,
		ctx:      ctx,
	}
}

// CreateTask creates a new task with a bounty
func (c *TaskBountyContract) CreateTask(creator sdk.AccAddress, title, description string, bounty sdk.Coins) (*Task, error) {
	// Verify creator has enough balance
	balance := c.keeper.GetBalance(c.ctx, creator, bounty[0].Denom)
	if balance.IsLT(bounty[0]) {
		return nil, fmt.Errorf("insufficient balance: got %v, need %v", balance, bounty[0])
	}

	// Lock bounty tokens
	err := c.keeper.SendCoinsFromAccountToModule(c.ctx, creator, "taskbounty", bounty)
	if err != nil {
		return nil, fmt.Errorf("failed to lock bounty: %w", err)
	}

	task := &Task{
		ID:          c.nextID,
		Creator:     creator.String(),
		Title:       title,
		Description: description,
		Bounty:      bounty,
		Status:      TaskStatusOpen,
		CreatedAt:   c.ctx.BlockTime().Unix(),
		UpdatedAt:   c.ctx.BlockTime().Unix(),
	}

	c.tasks[c.nextID] = task
	c.nextID++

	return task, nil
}

// ClaimTask allows a user to claim an open task
func (c *TaskBountyContract) ClaimTask(taskID uint64, claimant sdk.AccAddress) error {
	task, exists := c.tasks[taskID]
	if !exists {
		return fmt.Errorf("task not found: %d", taskID)
	}

	if task.Status != TaskStatusOpen {
		return fmt.Errorf("task is not open: %d", taskID)
	}

	task.Status = TaskStatusClaimed
	task.Claimant = claimant.String()
	task.UpdatedAt = c.ctx.BlockTime().Unix()

	return nil
}

// SubmitProof allows a claimant to submit proof of work
func (c *TaskBountyContract) SubmitProof(taskID uint64, claimant sdk.AccAddress, proofURI string) error {
	task, exists := c.tasks[taskID]
	if !exists {
		return fmt.Errorf("task not found: %d", taskID)
	}

	if task.Status != TaskStatusClaimed {
		return fmt.Errorf("task is not claimed: %d", taskID)
	}

	if task.Claimant != claimant.String() {
		return fmt.Errorf("only claimant can submit proof")
	}

	task.Status = TaskStatusCompleted
	task.ProofURI = proofURI
	task.UpdatedAt = c.ctx.BlockTime().Unix()

	return nil
}

// ApproveTask allows the creator to approve the task and release the bounty
func (c *TaskBountyContract) ApproveTask(taskID uint64, creator sdk.AccAddress) error {
	task, exists := c.tasks[taskID]
	if !exists {
		return fmt.Errorf("task not found: %d", taskID)
	}

	if task.Status != TaskStatusCompleted {
		return fmt.Errorf("task is not completed: %d", taskID)
	}

	if task.Creator != creator.String() {
		return fmt.Errorf("only creator can approve task")
	}

	// Convert claimant address string back to AccAddress
	claimant, err := sdk.AccAddressFromBech32(task.Claimant)
	if err != nil {
		return fmt.Errorf("invalid claimant address: %w", err)
	}

	// Release bounty to claimant
	err = c.keeper.SendCoinsFromModuleToAccount(c.ctx, "taskbounty", claimant, task.Bounty)
	if err != nil {
		return fmt.Errorf("failed to release bounty: %w", err)
	}

	task.Status = TaskStatusApproved
	task.UpdatedAt = c.ctx.BlockTime().Unix()

	return nil
}

// GetTask returns a task by ID
func (c *TaskBountyContract) GetTask(taskID uint64) (*Task, error) {
	task, exists := c.tasks[taskID]
	if !exists {
		return nil, fmt.Errorf("task not found: %d", taskID)
	}
	return task, nil
}

// ListTasks returns all tasks
func (c *TaskBountyContract) ListTasks() []*Task {
	tasks := make([]*Task, 0, len(c.tasks))
	for _, task := range c.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

// ListTasksByStatus returns tasks filtered by status
func (c *TaskBountyContract) ListTasksByStatus(status TaskStatus) []*Task {
	tasks := make([]*Task, 0)
	for _, task := range c.tasks {
		if task.Status == status {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

// ListTasksByCreator returns tasks created by a specific address
func (c *TaskBountyContract) ListTasksByCreator(creator string) []*Task {
	tasks := make([]*Task, 0)
	for _, task := range c.tasks {
		if task.Creator == creator {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

// String returns a string representation of TaskStatus
func (s TaskStatus) String() string {
	switch s {
	case TaskStatusOpen:
		return "OPEN"
	case TaskStatusClaimed:
		return "CLAIMED"
	case TaskStatusCompleted:
		return "COMPLETED"
	case TaskStatusApproved:
		return "APPROVED"
	default:
		return "UNKNOWN"
	}
}

// MarshalJSON implements the json.Marshaler interface
func (s TaskStatus) MarshalJSON() ([]byte, error) {
	return fmt.Appendf(nil, "\"%s\"", s.String()), nil
}
