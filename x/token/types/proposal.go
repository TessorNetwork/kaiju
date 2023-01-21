package types

import (
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

var (
	ProposalTypeSetTokenContractCodes                  = "SetTokenContractCodes"
	_                                 govtypes.Content = &SetTokenContractCodes{}
)

func NewSetTokenContract(cw20Code, cw721Code, kaiju1155Code uint64) SetTokenContractCodes {
	return SetTokenContractCodes{
		Cw20ContractCode:    cw20Code,
		Cw721ContractCode:   cw721Code,
		Kaiju1155ContractCode: kaiju1155Code,
	}
}

func (p *SetTokenContractCodes) GetDescription() string {
	return "update token contract codes"
}

func (p *SetTokenContractCodes) GetTitle() string {
	return "set token contract codes"
}

func (sup *SetTokenContractCodes) ProposalRoute() string { return RouterKey }
func (sup *SetTokenContractCodes) ProposalType() string  { return ProposalTypeSetTokenContractCodes }

func (sup *SetTokenContractCodes) ValidateBasic() error { return nil }

func init() {
	govtypes.RegisterProposalType(ProposalTypeSetTokenContractCodes)
	govtypes.RegisterProposalTypeCodec(&SetTokenContractCodes{}, "token.kaiju.token.v1beta1.SetTokenContractCodes")
}
