package keeper

import (
	"encoding/binary"
	"fmt"
	"strconv"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// "github.com/tendermint/tendermint/libs/log"

	"github.com/nutcase/dao-golang/x/taskbounty/types"
)


type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey storetypes.StoreKey
	bank     types.BankKeeper
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	bank types.BankKeeper,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		bank:     bank,
	}
}

// GetTaskCount get the total number of tasks
func (k Keeper) GetTaskCount(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.TaskCountKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

// SetTaskCount set the total number of tasks
func (k Keeper) SetTaskCount(ctx sdk.Context, count uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(types.TaskCountKey, bz)
}

// CreateTask creates a new task
func (k Keeper) CreateTask(ctx sdk.Context, task types.Task) (string, error) {
	count := k.GetTaskCount(ctx)
	task.ID = strconv.FormatUint(count+1, 10)
	task.Status = types.TaskStatus_OPEN

	// Ensure creator has enough funds for bounty
	if !k.bank.HasBalance(ctx, sdk.AccAddress(task.Creator), task.Bounty[0]) {
		return "", fmt.Errorf("insufficient funds for bounty")
	}

	// Lock the bounty funds
	if err := k.bank.SendCoinsFromAccountToModule(ctx, sdk.AccAddress(task.Creator), types.ModuleName, task.Bounty); err != nil {
		return "", err
	}

	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&task)
	store.Set(types.TaskKey(task.ID), bz)
	k.SetTaskCount(ctx, count+1)

	// Emit event
	ctx.EventManager().EmitEvent(
		types.NewEventCreateTask(
			task.ID,
			task.Creator,
			task.Title,
			task.Bounty.String(),
		),
	)

	return task.ID, nil
}

// GetTask returns a task by its id
func (k Keeper) GetTask(ctx sdk.Context, id string) (types.Task, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.TaskKey(id))
	if bz == nil {
		return types.Task{}, fmt.Errorf("task not found: %s", id)
	}

	var task types.Task
	k.cdc.MustUnmarshal(bz, &task)
	return task, nil
}

// ClaimTask allows a user to claim a task
func (k Keeper) ClaimTask(ctx sdk.Context, id string, claimer string) error {
	task, err := k.GetTask(ctx, id)
	if err != nil {
		return err
	}

	// Check if task is already claimed
	if task.Claimer != "" {
		return fmt.Errorf("task already claimed")
	}

	// Update task
	task.Claimer = claimer
	task.Status = types.TaskStatus_IN_PROGRESS

	// Save updated task
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&task)
	store.Set(types.TaskKey(id), bz)

	// Emit event
	ctx.EventManager().EmitEvent(
		types.NewEventClaimTask(
			task.ID,
			task.Creator,
			claimer,
		),
	)

	return nil
}

// SubmitTaskProof allows a claimer to submit proof for a task
func (k Keeper) SubmitTaskProof(ctx sdk.Context, id string, proof string) error {
	task, err := k.GetTask(ctx, id)
	if err != nil {
		return err
	}

	// Update task
	task.Proof = proof
	task.Status = types.TaskStatus_PENDING_APPROVAL

	// Save updated task
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&task)
	store.Set(types.TaskKey(id), bz)

	// Emit event
	ctx.EventManager().EmitEvent(
		types.NewEventSubmitTaskProof(
			task.ID,
			task.Creator,
			task.Claimer,
			proof,
		),
	)

	return nil
}

// ApproveTask allows the creator to approve a task and release the bounty
func (k Keeper) ApproveTask(ctx sdk.Context, id string) error {
	task, err := k.GetTask(ctx, id)
	if err != nil {
		return err
	}

	// Transfer bounty to claimer
	if err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(task.Claimer), task.Bounty); err != nil {
		return err
	}

	// Update task
	task.Status = types.TaskStatus_COMPLETED

	// Save updated task
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&task)
	store.Set(types.TaskKey(id), bz)

	// Emit event
	ctx.EventManager().EmitEvent(
		types.NewEventApproveTask(
			task.ID,
			task.Creator,
			task.Claimer,
		),
	)

	return nil
}

// GetAllTasks returns all tasks
func (k Keeper) GetAllTasks(ctx sdk.Context) []types.Task {
	var tasks []types.Task
	store := ctx.KVStore(k.storeKey)
	prefixedStore := prefix.NewStore(store, []byte(types.TaskKeyPrefix))
	iterator := prefixedStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var task types.Task
		k.cdc.MustUnmarshal(iterator.Value(), &task)
		tasks = append(tasks, task)
	}
	return tasks
}
