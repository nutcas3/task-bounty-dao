package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/cosmos/gogoproto/proto"
	// "time"
)

// TaskStatus represents the current state of a task
type TaskStatus int32

const (
	TaskStatus_OPEN             TaskStatus = 0
	TaskStatus_IN_PROGRESS      TaskStatus = 1
	TaskStatus_PENDING_APPROVAL TaskStatus = 2
	TaskStatus_COMPLETED        TaskStatus = 3
)

// String returns the string representation of a TaskStatus
func (ts TaskStatus) String() string {
	switch ts {
	case TaskStatus_OPEN:
		return "OPEN"
	case TaskStatus_IN_PROGRESS:
		return "IN_PROGRESS"
	case TaskStatus_PENDING_APPROVAL:
		return "PENDING_APPROVAL"
	case TaskStatus_COMPLETED:
		return "COMPLETED"
	default:
		return "UNKNOWN"
	}
}

// Task represents a task in the bounty system
type Task struct {
	ID          string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Creator     string     `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator"`
	Title       string     `protobuf:"bytes,3,opt,name=title,proto3" json:"title"`
	Description string     `protobuf:"bytes,4,opt,name=description,proto3" json:"description"`
	Bounty      sdk.Coins  `protobuf:"bytes,5,rep,name=bounty,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"bounty"`
	Status      TaskStatus `protobuf:"varint,6,opt,name=status,proto3,enum=taskbounty.TaskStatus" json:"status"`
	Claimer     string     `protobuf:"bytes,7,opt,name=claimer,proto3" json:"claimer,omitempty"`
	Proof       string     `protobuf:"bytes,8,opt,name=proof,proto3" json:"proof,omitempty"`
	Claimant    string     `protobuf:"bytes,9,opt,name=claimer,proto3" json:"claimant,omitempty"`

}

// Required methods for proto.Message interface
func (t *Task) Reset()         { *t = Task{} }
func (t *Task) String() string { return proto.CompactTextString(t) }
func (t *Task) ProtoMessage()  {}

// NewTask creates a new Task instance
func NewTask(id, creator, title, description string, bounty sdk.Coins) Task {
	return Task{
		ID:          id,
		Creator:     creator,
		Title:       title,
		Description: description,
		Bounty:      bounty,
		Status:      TaskStatus_OPEN,
	}
}
