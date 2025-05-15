package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Event types for the taskbounty module
const (
	EventTypeCreateTask    = "create_task"
	EventTypeClaimTask     = "claim_task"
	EventTypeSubmitProof   = "submit_proof"
	EventTypeApproveTask   = "approve_task"
	EventTypeCancelTask    = "cancel_task"

	// Attribute keys
	AttributeKeyTaskID     = "task_id"
	AttributeKeyCreator    = "creator"
	AttributeKeyClaimant   = "claimant"
	AttributeKeyTitle      = "title"
	AttributeKeyBounty     = "bounty"
	AttributeKeyProofURI   = "proof_uri"
	AttributeKeyStatus     = "status"
	AttributeKeyTimestamp  = "timestamp"
)

// NewEventCreateTask creates a new task creation event
func NewEventCreateTask(taskID string, creator, title, bounty string) sdk.Event {
	return sdk.NewEvent(
		EventTypeCreateTask,
		sdk.NewAttribute(AttributeKeyTaskID, taskID),
		sdk.NewAttribute(AttributeKeyCreator, creator),
		sdk.NewAttribute(AttributeKeyTitle, title),
		sdk.NewAttribute(AttributeKeyBounty, bounty),
	)
}

// NewEventClaimTask creates a new task claim event
func NewEventClaimTask(taskID, creator, claimer string) sdk.Event {
	return sdk.NewEvent(
		EventTypeClaimTask,
		sdk.NewAttribute(AttributeKeyTaskID, taskID),
		sdk.NewAttribute(AttributeKeyCreator, creator),
		sdk.NewAttribute(AttributeKeyClaimant, claimer),
	)
}

// NewEventSubmitTaskProof creates a new proof submission event
func NewEventSubmitTaskProof(taskID, creator, claimer, proof string) sdk.Event {
	return sdk.NewEvent(
		EventTypeSubmitProof,
		sdk.NewAttribute(AttributeKeyTaskID, taskID),
		sdk.NewAttribute(AttributeKeyCreator, creator),
		sdk.NewAttribute(AttributeKeyClaimant, claimer),
		sdk.NewAttribute(AttributeKeyProofURI, proof),
	)
}

// NewEventApproveTask creates a new task approval event
func NewEventApproveTask(taskID, creator, claimer string) sdk.Event {
	return sdk.NewEvent(
		EventTypeApproveTask,
		sdk.NewAttribute(AttributeKeyTaskID, taskID),
		sdk.NewAttribute(AttributeKeyCreator, creator),
		sdk.NewAttribute(AttributeKeyClaimant, claimer),
	)
}

// NewEventCancelTask creates a new task cancellation event
func NewEventCancelTask(taskID, creator string) sdk.Event {
	return sdk.NewEvent(
		EventTypeCancelTask,
		sdk.NewAttribute(AttributeKeyTaskID, taskID),
		sdk.NewAttribute(AttributeKeyCreator, creator),
	)
}
