package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	types "github.com/nutcase/dao-golang/x/taskbounty/types"
	taskbountyquery "github.com/nutcase/dao-golang/x/taskbounty/types/taskbounty/v1"
	taskbountytypes "github.com/nutcase/dao-golang/x/taskbounty/types/taskbounty/v1"
	"github.com/spf13/cobra"
)

func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdListTasks(),
		CmdShowTask(),
		CmdTasksByStatus(),
		CmdTasksByCreator(),
		CmdTasksByClaimant(),
	)

	return cmd
}

func CmdListTasks() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-tasks",
		Short: "List all tasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := taskbountyquery.NewQueryClient(clientCtx)

			params := &taskbountyquery.QueryTasksRequest{}

			res, err := queryClient.Tasks(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdShowTask() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-task [id]",
		Short: "Shows a task by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := taskbountyquery.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &taskbountyquery.QueryTaskRequest{
				Id: strconv.FormatUint(id, 10),
			}

			res, err := queryClient.Task(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdTasksByStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tasks-by-status [status]",
		Short: "Query tasks by status (OPEN, CLAIMED, PROOF_SUBMITTED, COMPLETED)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := taskbountyquery.NewQueryClient(clientCtx)

			// Convert string status to TaskStatus enum
			var status int32
			switch args[0] {
			case "OPEN":
				status = 1 // TaskStatus_TASK_STATUS_OPEN
			case "CLAIMED":
				status = 2 // TaskStatus_TASK_STATUS_CLAIMED
			case "PROOF_SUBMITTED":
				status = 3 // TaskStatus_TASK_STATUS_PROOF_SUBMITTED
			case "COMPLETED":
				status = 4 // TaskStatus_TASK_STATUS_COMPLETED
			default:
				return fmt.Errorf("invalid status: %s. Must be one of: OPEN, CLAIMED, PROOF_SUBMITTED, COMPLETED", args[0])
			}

			params := &taskbountyquery.QueryTasksByStatusRequest{
				Status: taskbountytypes.TaskStatus(status),
			}

			res, err := queryClient.TasksByStatus(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdTasksByCreator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tasks-by-creator [creator]",
		Short: "Query tasks by creator address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := taskbountyquery.NewQueryClient(clientCtx)

			params := &taskbountyquery.QueryTasksByCreatorRequest{
				Creator: args[0],
			}

			res, err := queryClient.TasksByCreator(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdTasksByClaimant() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tasks-by-claimant [claimant]",
		Short: "Query tasks by claimant address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := taskbountyquery.NewQueryClient(clientCtx)

			params := &taskbountyquery.QueryTasksByClaimantRequest{
				Claimant: args[0],
			}

			res, err := queryClient.TasksByClaimant(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
