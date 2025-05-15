package types

import (
	"fmt"

	proto "github.com/gogo/protobuf/proto"
)

// Proto file implementation
var _ proto.Message = (*MsgCreateTask)(nil)
var _ proto.Message = (*MsgClaimTask)(nil)
var _ proto.Message = (*MsgSubmitProof)(nil)
var _ proto.Message = (*MsgApproveTask)(nil)
var _ proto.Message = (*MsgCreateWallet)(nil)

// Reset implements proto.Message
func (m *MsgCreateTask) Reset() { *m = MsgCreateTask{} }
func (m *MsgClaimTask) Reset()  { *m = MsgClaimTask{} }
func (m *MsgSubmitProof) Reset() { *m = MsgSubmitProof{} }
func (m *MsgApproveTask) Reset() { *m = MsgApproveTask{} }
func (m *MsgCreateWallet) Reset() { *m = MsgCreateWallet{} }

// String implements proto.Message
func (m *MsgCreateTask) String() string {
	return fmt.Sprintf("CreateTask{Creator: %s, Title: %s}", m.Creator, m.Title)
}
func (m *MsgClaimTask) String() string {
	return fmt.Sprintf("ClaimTask{Claimer: %s, TaskID: %s}", m.Claimer, m.TaskID)
}
func (m *MsgSubmitProof) String() string {
	return fmt.Sprintf("SubmitProof{Worker: %s, TaskID: %s}", m.Worker, m.TaskID)
}
func (m *MsgApproveTask) String() string {
	return fmt.Sprintf("ApproveTask{Approver: %s, TaskID: %s}", m.Approver, m.TaskID)
}
func (m *MsgCreateWallet) String() string {
	return fmt.Sprintf("CreateWallet{Owner: %s}", m.Owner)
}

// ProtoMessage implements proto.Message
func (m *MsgCreateTask) ProtoMessage()    {}
func (m *MsgClaimTask) ProtoMessage()     {}
func (m *MsgSubmitProof) ProtoMessage()   {}
func (m *MsgApproveTask) ProtoMessage()   {}
func (m *MsgCreateWallet) ProtoMessage()  {}
