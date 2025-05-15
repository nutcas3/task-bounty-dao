package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewMsgCreateTask creates a new MsgCreateTask instance
func NewMsgCreateTask(creator, title, description string, bounty sdk.Coins) *MsgCreateTask {
	return &MsgCreateTask{
		Creator:     creator,
		Title:       title,
		Description: description,
		Bounty:      bounty, // Bounty field is sdk.Coins in the struct
	}
}

// NewMsgClaimTask creates a new MsgClaimTask instance
func NewMsgClaimTask(claimer, taskID string) *MsgClaimTask {
	return &MsgClaimTask{
		TaskID:  taskID,
		Claimer: claimer,
	}
}

// NewMsgSubmitProof creates a new MsgSubmitProof instance
func NewMsgSubmitProof(worker, taskID, proof string) *MsgSubmitProof {
	return &MsgSubmitProof{
		TaskID: taskID,
		Worker: worker,
		Proof:  proof,
	}
}

// NewMsgApproveTask creates a new MsgApproveTask instance
func NewMsgApproveTask(approver, taskID string) *MsgApproveTask {
	return &MsgApproveTask{
		TaskID:   taskID,
		Approver: approver,
	}
}

// NewMsgCreateWallet creates a new MsgCreateWallet instance
func NewMsgCreateWallet(owner string) *MsgCreateWallet {
	return &MsgCreateWallet{
		Owner: owner,
	}
}
