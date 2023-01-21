package cli

import (
	"encoding/json"

	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govcli "github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/tessornetwork/kaiju/x/token/types"
	"github.com/spf13/cobra"
)

func NewTxCmd() *cobra.Command {
	tokenTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "token transaction sub commands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	tokenTxCmd.AddCommand(
		NewCmdCreateToken(),
	)

	return tokenTxCmd
}

// NewCmdSubmitUpgradeProposal implements a command handler for submitting a software upgrade proposal transaction.
func NewCmdUpdateTokenParamsProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-token-params [nft_contract_code] [nft_minter_address] [flags]",
		Args:  cobra.ExactArgs(3),
		Short: "Submit a proposal to update token params",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			cw20CodeId, err := strconv.ParseUint(args[0], 0, 64)
			if err != nil {
				return err
			}

			cw721CodeId, err := strconv.ParseUint(args[1], 0, 64)
			if err != nil {
				return err
			}

			kaiju1155CodeId, err := strconv.ParseUint(args[2], 0, 64)
			if err != nil {
				return err
			}

			content := types.NewSetTokenContract(cw20CodeId, cw721CodeId, kaiju1155CodeId)

			from := clientCtx.GetFromAddress()

			depositStr, err := cmd.Flags().GetString(govcli.FlagDeposit)
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			msg, err := govtypes.NewMsgSubmitProposal(&content, deposit, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(govcli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(govcli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(govcli.FlagDeposit, "", "deposit of proposal")

	return cmd
}

func NewCmdCreateToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-token [token-iid]",
		Short: "Create a new TokenDoc",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			var msg types.MsgMint
			err := json.Unmarshal([]byte(args[0]), &msg)
			if err != nil {
				return err
			}

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg.MinterAddress = clientCtx.GetFromAddress().String()

			err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
			if err != nil {
				return err
			}

			return nil
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// func NewCmdUpdateProjectStatus() *cobra.Command {
// cmd := &cobra.Command{
// 	Use:   "update-project-status [sender-did] [status] [kaiju-did]",
// 	Short: "Update the status of a project signed by the kaijuDid of the project",
// 	Args:  cobra.ExactArgs(3),
// 	RunE: func(cmd *cobra.Command, args []string) error {
// 		senderDid := args[0]
// 		status := args[1]
// 		kaijuDid, err := didtypes.UnmarshalKaijuDid(args[2])
// 		if err != nil {
// 			return err
// 		}

// 		projectStatus := types.ProjectStatus(status)
// 		if projectStatus != types.CreatedProject &&
// 			projectStatus != types.PendingStatus &&
// 			projectStatus != types.FundedStatus &&
// 			projectStatus != types.StartedStatus &&
// 			projectStatus != types.StoppedStatus &&
// 			projectStatus != types.PaidoutStatus {
// 			return errors.New("The status must be one of 'CREATED', " +
// 				"'PENDING', 'FUNDED', 'STARTED', 'STOPPED' or 'PAIDOUT'")
// 		}

// 		updateProjectStatusDoc := types.NewUpdateProjectStatusDoc(
// 			projectStatus, "")

// 		clientCtx, err := client.GetClientTxContext(cmd)
// 		if err != nil {
// 			return err
// 		}
// 		clientCtx = clientCtx.WithFromAddress(kaijuDid.Address())

// 		msg := types.NewMsgUpdateProjectStatus(senderDid, updateProjectStatusDoc, kaijuDid.Did)
// 		err = msg.ValidateBasic()
// 		if err != nil {
// 			return err
// 		}

// 		return kaijutypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), kaijuDid, msg)
// 	},
// }

// 	flags.AddTxFlagsToCmd(cmd)
// 	return cmd
// }

// func NewCmdCreateAgent() *cobra.Command {
// cmd := &cobra.Command{
// 	Use: "create-agent [tx-hash] [sender-did] [agent-did] " +
// 		"[role] [project-did]",
// 	Short: "Create a new agent on a project signed by the kaijuDid of the project",
// 	Args:  cobra.ExactArgs(5),
// 	RunE: func(cmd *cobra.Command, args []string) error {
// 		txHash := args[0]
// 		senderDid := args[1]
// 		agentDid := args[2]
// 		role := args[3]
// 		if role != "SA" && role != "EA" && role != "IA" {
// 			return errors.New("The role must be one of 'SA', 'EA' or 'IA'")
// 		}

// 		createAgentDoc := types.NewCreateAgentDoc(agentDid, role)

// 		kaijuDid, err := didtypes.UnmarshalKaijuDid(args[4])
// 		if err != nil {
// 			return err
// 		}

// 		clientCtx, err := client.GetClientTxContext(cmd)
// 		if err != nil {
// 			return err
// 		}
// 		clientCtx = clientCtx.WithFromAddress(kaijuDid.Address())

// 		msg := types.NewMsgCreateAgent(txHash, senderDid, createAgentDoc, kaijuDid.Did)
// 		err = msg.ValidateBasic()
// 		if err != nil {
// 			return err
// 		}

// 		return kaijutypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), kaijuDid, msg)
// 	},
// }

// 	flags.AddTxFlagsToCmd(cmd)
// 	return cmd
// }

// func NewCmdUpdateAgent() *cobra.Command {
// cmd := &cobra.Command{
// 	Use: "update-agent [tx-hash] [sender-did] [agent-did] " +
// 		"[status] [kaiju-did]",
// 	Short: "Update the status of an agent on a project signed by the kaijuDid of the project",
// 	Args:  cobra.ExactArgs(6),
// 	RunE: func(cmd *cobra.Command, args []string) error {
// 		txHash := args[0]
// 		senderDid := args[1]
// 		agentDid := args[2]
// 		agentRole := args[4]
// 		agentStatus := types.AgentStatus(args[3])
// 		if agentStatus != types.PendingAgent && agentStatus != types.ApprovedAgent && agentStatus != types.RevokedAgent {
// 			return errors.New("The status must be one of '0' (Pending), '1' (Approved) or '2' (Revoked)")
// 		}

// 		updateAgentDoc := types.NewUpdateAgentDoc(
// 			agentDid, agentStatus, agentRole)

// 		kaijuDid, err := didtypes.UnmarshalKaijuDid(args[5])
// 		if err != nil {
// 			return err
// 		}

// 		clientCtx, err := client.GetClientTxContext(cmd)
// 		if err != nil {
// 			return err
// 		}
// 		clientCtx = clientCtx.WithFromAddress(kaijuDid.Address())

// 		msg := types.NewMsgUpdateAgent(txHash, senderDid, updateAgentDoc, kaijuDid.Did)
// 		err = msg.ValidateBasic()
// 		if err != nil {
// 			return err
// 		}

// 		return kaijutypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), kaijuDid, msg)
// 	},
// }

// 	flags.AddTxFlagsToCmd(cmd)
// 	return cmd
// }

// func NewCmdCreateClaim() *cobra.Command {
// cmd := &cobra.Command{
// 	Use:   "create-claim [tx-hash] [sender-did] [claim-id] [claim-template-id] [kaiju-did]",
// 	Short: "Create a new claim on a project signed by the kaijuDid of the project",
// 	Args:  cobra.ExactArgs(5),
// 	RunE: func(cmd *cobra.Command, args []string) error {
// 		txHash := args[0]
// 		senderDid := args[1]
// 		claimId := args[2]
// 		claimTemplateId := args[3]
// 		createClaimDoc := types.NewCreateClaimDoc(claimId, claimTemplateId)

// 		kaijuDid, err := didtypes.UnmarshalKaijuDid(args[4])
// 		if err != nil {
// 			return err
// 		}

// 		clientCtx, err := client.GetClientTxContext(cmd)
// 		if err != nil {
// 			return err
// 		}
// 		clientCtx = clientCtx.WithFromAddress(kaijuDid.Address())

// 		msg := types.NewMsgCreateClaim(txHash, senderDid, createClaimDoc, kaijuDid.Did)
// 		err = msg.ValidateBasic()
// 		if err != nil {
// 			return err
// 		}

// 		return kaijutypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), kaijuDid, msg)
// 	},
// }

// 	flags.AddTxFlagsToCmd(cmd)
// 	return cmd
// }

// func NewCmdCreateEvaluation() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use: "create-evaluation [tx-hash] [sender-did] [claim-id] " +
// 			"[status] [kaiju-did]",
// 		Short: "Create a new claim evaluation on a project signed by the kaijuDid of the project",
// 		Args:  cobra.ExactArgs(5),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			txHash := args[0]
// 			senderDid := args[1]
// 			claimId := args[2]
// 			claimStatus := types.ClaimStatus(args[3])
// 			if claimStatus != types.PendingClaim && claimStatus != types.ApprovedClaim && claimStatus != types.RejectedClaim {
// 				return errors.New("The status must be one of '0' (Pending), '1' (Approved) or '2' (Rejected)")
// 			}

// 			createEvaluationDoc := types.NewCreateEvaluationDoc(
// 				claimId, claimStatus)

// 			kaijuDid, err := didtypes.UnmarshalKaijuDid(args[4])
// 			if err != nil {
// 				return err
// 			}

// 			clientCtx, err := client.GetClientTxContext(cmd)
// 			if err != nil {
// 				return err
// 			}
// 			clientCtx = clientCtx.WithFromAddress(kaijuDid.Address())

// 			msg := types.NewMsgCreateEvaluation(txHash, senderDid, createEvaluationDoc, kaijuDid.Did)
// 			err = msg.ValidateBasic()
// 			if err != nil {
// 				return err
// 			}

// 			return kaijutypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), kaijuDid, msg)
// 		},
// 	}

// 	flags.AddTxFlagsToCmd(cmd)
// 	return cmd
// }

// func NewCmdWithdrawFunds() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "withdraw-funds [sender-did] [iid]",
// 		Short: "Withdraw funds.",
// 		Args:  cobra.ExactArgs(2),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			kaijuDid, err := didtypes.UnmarshalKaijuDid(args[0])
// 			if err != nil {
// 				return err
// 			}

// 			var data types.WithdrawFundsDoc
// 			err = json.Unmarshal([]byte(args[1]), &data)
// 			if err != nil {
// 				return err
// 			}

// 			clientCtx, err := client.GetClientTxContext(cmd)
// 			if err != nil {
// 				return err
// 			}
// 			clientCtx = clientCtx.WithFromAddress(kaijuDid.Address())

// 			msg := types.NewMsgWithdrawFunds(kaijuDid.Did, data)
// 			err = msg.ValidateBasic()
// 			if err != nil {
// 				return err
// 			}

// 			return kaijutypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), kaijuDid, msg)
// 		},
// 	}

// 	flags.AddTxFlagsToCmd(cmd)
// 	return cmd
// }

// func NewCmdUpdateProjectDoc() *cobra.Command {
// cmd := &cobra.Command{
// 	Use:   "update-project-doc [sender-did] [project-iid-json] [kaiju-did]",
// 	Short: "Update a project's iid signed by the kaijuDid of the project",
// 	Args:  cobra.ExactArgs(3),
// 	RunE: func(cmd *cobra.Command, args []string) error {
// 		senderDid := args[0]
// 		projectDataStr := args[1]
// 		kaijuDid, err := didtypes.UnmarshalKaijuDid(args[2])
// 		if err != nil {
// 			return err
// 		}

// 		clientCtx, err := client.GetClientTxContext(cmd)
// 		if err != nil {
// 			return err
// 		}
// 		clientCtx = clientCtx.WithFromAddress(kaijuDid.Address())

// 		msg := types.NewMsgUpdateProjectDoc(senderDid, json.RawMessage(projectDataStr), kaijuDid.Did)
// 		err = msg.ValidateBasic()
// 		if err != nil {
// 			return err
// 		}

// 		return kaijutypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), kaijuDid, msg)
// 	},
// }

// 	flags.AddTxFlagsToCmd(cmd)
// 	return cmd
// }
