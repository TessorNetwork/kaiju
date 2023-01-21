package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/petrinetwork/xco-blockchain/x/entity/client/cli"
	"github.com/petrinetwork/xco-blockchain/x/entity/client/rest"
)

var ProposalHandler = govclient.NewProposalHandler(cli.NewCmdUpdateEntityParamsProposal, rest.ProposalRESTHandler)
