package cli

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	kaijutypes "github.com/tessornetwork/kaiju/lib/kaiju"
	didtypes "github.com/tessornetwork/kaiju/lib/legacydid"
	iidtypes "github.com/tessornetwork/kaiju/x/iid/types"
	"github.com/tessornetwork/kaiju/x/project/types"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewTxCmd() *cobra.Command {
	projectTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "project transaction sub commands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	projectTxCmd.AddCommand(
		NewCmdCreateProject(),
		NewCmdCreateAgent(),
		NewCmdUpdateProjectStatus(),
		NewCmdUpdateAgent(),
		NewCmdCreateClaim(),
		NewCmdCreateEvaluation(),
		NewCmdWithdrawFunds(),
		NewCmdUpdateProjectDoc(),
	)

	return projectTxCmd
}

func NewCmdCreateProject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-project [sender-did] [project-iid-json] [kaiju-did]",
		Short: "Create a new ProjectDoc signed by the kaijuDid of the project",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			senderDid := args[0]
			projectDataStr := args[1]
			kaijuDidStr := args[2]

			kaijuDid, err := didtypes.UnmarshalKaijuDid(kaijuDidStr)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithFromAddress(kaijuDid.Address())

			msg := types.NewMsgCreateProject(
				iidtypes.DIDFragment(senderDid), json.RawMessage(projectDataStr), kaijuDid.Did, kaijuDid.VerifyKey, kaijuDid.Address().String())
			msg.ProjectAddress = kaijuDid.Address().String()
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			res, err := kaijutypes.SignAndBroadcastTxFromStdSignMsg(clientCtx, msg, kaijuDid, cmd.Flags())
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdUpdateProjectStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-project-status [sender-did] [status] [kaiju-did]",
		Short: "Update the status of a project signed by the kaijuDid of the project",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			senderDid := args[0]
			status := args[1]
			kaijuDid, err := didtypes.UnmarshalKaijuDid(args[2])
			if err != nil {
				return err
			}

			projectStatus := types.ProjectStatus(status)
			//if projectStatus != types.CreatedProject &&
			//	projectStatus != types.PendingStatus &&
			//	projectStatus != types.FundedStatus &&
			//	projectStatus != types.StartedStatus &&
			//	projectStatus != types.StoppedStatus &&
			//	projectStatus != types.PaidoutStatus {
			//	return errors.New("The status must be one of 'CREATED', " +
			//		"'PENDING', 'FUNDED', 'STARTED', 'STOPPED' or 'PAIDOUT'")
			//}

			updateProjectStatusDoc := types.NewUpdateProjectStatusDoc(
				projectStatus, "")

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithFromAddress(kaijuDid.Address())

			msg := types.NewMsgUpdateProjectStatus(iidtypes.DIDFragment(senderDid), updateProjectStatusDoc, kaijuDid.Did, kaijuDid.Address().String())
			msg.ProjectAddress = kaijuDid.Address().String()
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return kaijutypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), kaijuDid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdCreateAgent() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-agent [tx-hash] [sender-did] [agent-did] " +
			"[role] [project-did]",
		Short: "Create a new agent on a project signed by the kaijuDid of the project",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			txHash := args[0]
			senderDid := args[1]
			agentDid := args[2]
			role := args[3]
			if role != "SA" && role != "EA" && role != "IA" {
				return errors.New("The role must be one of 'SA', 'EA' or 'IA'")
			}

			createAgentDoc := types.NewCreateAgentDoc(agentDid, role)

			kaijuDid, err := didtypes.UnmarshalKaijuDid(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithFromAddress(kaijuDid.Address())

			msg := types.NewMsgCreateAgent(txHash, iidtypes.DIDFragment(senderDid), createAgentDoc, kaijuDid.Did, kaijuDid.Address().String())
			msg.ProjectAddress = kaijuDid.Address().String()
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return kaijutypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), kaijuDid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdUpdateAgent() *cobra.Command {
	cmd := &cobra.Command{
		Use: "update-agent [tx-hash] [sender-did] [agent-did] " +
			"[status] [kaiju-did]",
		Short: "Update the status of an agent on a project signed by the kaijuDid of the project",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			txHash := args[0]
			senderDid := args[1]
			agentDid := args[2]
			agentRole := args[4]
			agentStatus := types.AgentStatus(args[3])
			if agentStatus != types.PendingAgent && agentStatus != types.ApprovedAgent && agentStatus != types.RevokedAgent {
				return errors.New("The status must be one of '0' (Pending), '1' (Approved) or '2' (Revoked)")
			}

			updateAgentDoc := types.NewUpdateAgentDoc(
				agentDid, agentStatus, agentRole)

			kaijuDid, err := didtypes.UnmarshalKaijuDid(args[5])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithFromAddress(kaijuDid.Address())

			msg := types.NewMsgUpdateAgent(txHash, iidtypes.DIDFragment(senderDid), updateAgentDoc, kaijuDid.Did, kaijuDid.Address().String())
			msg.ProjectAddress = kaijuDid.Address().String()
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return kaijutypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), kaijuDid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdCreateClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-claim [tx-hash] [sender-did] [claim-id] [claim-template-id] [kaiju-did]",
		Short: "Create a new claim on a project signed by the kaijuDid of the project",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			txHash := args[0]
			senderDid := args[1]
			claimId := args[2]
			claimTemplateId := args[3]
			createClaimDoc := types.NewCreateClaimDoc(claimId, claimTemplateId)

			kaijuDid, err := didtypes.UnmarshalKaijuDid(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithFromAddress(kaijuDid.Address())

			msg := types.NewMsgCreateClaim(txHash, iidtypes.DIDFragment(senderDid), createClaimDoc, kaijuDid.Did, kaijuDid.Address().String())
			msg.ProjectAddress = kaijuDid.Address().String()
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return kaijutypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), kaijuDid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdCreateEvaluation() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-evaluation [tx-hash] [sender-did] [claim-id] " +
			"[status] [kaiju-did]",
		Short: "Create a new claim evaluation on a project signed by the kaijuDid of the project",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			txHash := args[0]
			senderDid := args[1]
			claimId := args[2]
			claimStatus := types.ClaimStatus(args[3])
			if claimStatus != types.PendingClaim && claimStatus != types.ApprovedClaim && claimStatus != types.RejectedClaim {
				return errors.New("The status must be one of '0' (Pending), '1' (Approved) or '2' (Rejected)")
			}

			createEvaluationDoc := types.NewCreateEvaluationDoc(
				claimId, claimStatus)

			kaijuDid, err := didtypes.UnmarshalKaijuDid(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithFromAddress(kaijuDid.Address())

			msg := types.NewMsgCreateEvaluation(txHash, iidtypes.DIDFragment(senderDid), createEvaluationDoc, kaijuDid.Did, kaijuDid.Address().String())
			msg.ProjectAddress = kaijuDid.Address().String()
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return kaijutypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), kaijuDid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdWithdrawFunds() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw-funds [sender-did] [iid]",
		Short: "Withdraw funds.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			kaijuDid, err := didtypes.UnmarshalKaijuDid(args[0])
			if err != nil {
				return err
			}

			var data types.WithdrawFundsDoc
			err = json.Unmarshal([]byte(args[1]), &data)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithFromAddress(kaijuDid.Address())

			msg := types.NewMsgWithdrawFunds(iidtypes.DIDFragment(kaijuDid.Did), data, kaijuDid.Address().String())
			msg.SenderAddress = kaijuDid.Address().String()
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return kaijutypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), kaijuDid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdUpdateProjectDoc() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-project-doc [sender-did] [project-iid-json] [kaiju-did]",
		Short: "Update a project's iid signed by the kaijuDid of the project",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			senderDid := args[0]
			projectDataStr := args[1]
			kaijuDid, err := didtypes.UnmarshalKaijuDid(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithFromAddress(kaijuDid.Address())

			msg := types.NewMsgUpdateProjectDoc(iidtypes.DIDFragment(senderDid), json.RawMessage(projectDataStr), kaijuDid.Did, kaijuDid.Address().String())
			msg.ProjectAddress = kaijuDid.Address().String()
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return kaijutypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), kaijuDid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
