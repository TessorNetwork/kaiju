package cli

import (
	"github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	libkaiju "github.com/tessornetwork/kaiju/lib/kaiju"
	"github.com/tessornetwork/kaiju/lib/legacydid"
	"github.com/tessornetwork/kaiju/x/iid/types"
	"github.com/spf13/cobra"
)

func NewCreateIidDocumentFormLegacyDidCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create-iid-from-legacy-did [did]",
		Short:   "create decentralized did (did) document from legacy did",
		Example: "creates a did document for users",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			kaijuDid, err := legacydid.UnmarshalKaijuDid(args[0])
			if err != nil {
				return err
			}

			// chaincmd.Flags().GetString(flags.FlagChainID)
			// if err != nil {
			// 	return err
			// }
			// did
			did := types.DID(kaijuDid.Did)
			// verification
			// signer := clientCtx.GetFromAddress()
			// pubkey

			pubKey := kaijuDid.VerifyKey

			clientCtx = clientCtx.WithFromAddress(kaijuDid.Address())

			// understand the vmType

			auth := types.NewVerification(
				types.NewVerificationMethod(
					kaijuDid.Did,
					did,
					types.NewPublicKeyMultibase(base58.Decode(pubKey), types.DIDVMethodTypeEd25519VerificationKey2018),
				),
				[]string{types.Authentication},
				nil,
			)
			// create the message
			msg := types.NewMsgCreateIidDocument(
				did.String(),
				types.Verifications{auth},
				types.Services{},
				types.AccordedRights{},
				types.LinkedResources{},
				types.LinkedEntities{},
				kaijuDid.Address().String(),
				types.Contexts{},
			)
			// validate
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			// execute
			return libkaiju.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), kaijuDid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
