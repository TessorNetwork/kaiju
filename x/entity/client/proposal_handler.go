package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/tessornetwork/kaiju/x/entity/client/cli"
	"github.com/tessornetwork/kaiju/x/entity/client/rest"
)

var ProposalHandler = govclient.NewProposalHandler(cli.NewCmdUpdateEntityParamsProposal, rest.ProposalRESTHandler)
