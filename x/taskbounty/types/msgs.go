package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Message types for the taskbounty module
const (
	TypeMsgCreateTask    = "create_task"
	TypeMsgClaimTask     = "claim_task"
	TypeMsgSubmitProof   = "submit_proof"
	TypeMsgApproveTask   = "approve_task"
	TypeMsgCreateWallet  = "create_wallet"
)

// MsgCreateTask defines a message to create a new task
type MsgCreateTask struct {
	Creator     string      `json:"creator"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Bounty      sdk.Coins  `json:"bounty"`
}

// MsgClaimTask defines a message to claim an existing task
type MsgClaimTask struct {
	TaskID  string `json:"task_id"`
	Claimer string `json:"claimer"`
}

// MsgSubmitProof defines a message to submit proof for a claimed task
type MsgSubmitProof struct {
	TaskID  string `json:"task_id"`
	Worker  string `json:"worker"`
	Proof   string `json:"proof"`
}

// MsgApproveTask defines a message to approve a completed task
type MsgApproveTask struct {
	TaskID    string `json:"task_id"`
	Approver  string `json:"approver"`
}

// MsgCreateWallet defines a message to create a new wallet
type MsgCreateWallet struct {
	Owner string `json:"owner"`
}

// Implement sdk.Msg interface
func (msg MsgCreateTask) Route() string { return ModuleName }
func (msg MsgCreateTask) Type() string  { return TypeMsgCreateTask }
func (msg MsgCreateTask) ValidateBasic() error { return nil }
func (msg MsgCreateTask) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg MsgClaimTask) Route() string { return ModuleName }
func (msg MsgClaimTask) Type() string  { return TypeMsgClaimTask }
func (msg MsgClaimTask) ValidateBasic() error { return nil }
func (msg MsgClaimTask) GetSigners() []sdk.AccAddress {
	claimer, err := sdk.AccAddressFromBech32(msg.Claimer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{claimer}
}

func (msg MsgSubmitProof) Route() string { return ModuleName }
func (msg MsgSubmitProof) Type() string  { return TypeMsgSubmitProof }
func (msg MsgSubmitProof) ValidateBasic() error { return nil }
func (msg MsgSubmitProof) GetSigners() []sdk.AccAddress {
	worker, err := sdk.AccAddressFromBech32(msg.Worker)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{worker}
}

func (msg MsgApproveTask) Route() string { return ModuleName }
func (msg MsgApproveTask) Type() string  { return TypeMsgApproveTask }
func (msg MsgApproveTask) ValidateBasic() error { return nil }
func (msg MsgApproveTask) GetSigners() []sdk.AccAddress {
	approver, err := sdk.AccAddressFromBech32(msg.Approver)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{approver}
}

func (msg MsgCreateWallet) Route() string { return ModuleName }
func (msg MsgCreateWallet) Type() string  { return TypeMsgCreateWallet }
func (msg MsgCreateWallet) ValidateBasic() error { return nil }
func (msg MsgCreateWallet) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}
