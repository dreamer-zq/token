package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// Register concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgIssueToken{}, "irismod/token/MsgIssueToken", nil)
	cdc.RegisterConcrete(MsgEditToken{}, "irismod/token/MsgEditToken", nil)
	cdc.RegisterConcrete(MsgMintToken{}, "irismod/token/MsgMintToken", nil)
	cdc.RegisterConcrete(MsgTransferTokenOwner{}, "irismod/token/MsgTransferTokenOwner", nil)

	cdc.RegisterConcrete(Token{}, "irismod/token/Token", nil)

	cdc.RegisterConcrete(&Params{}, "irismod/token/Params", nil)
}

var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	codec.RegisterCrypto(ModuleCdc)
	RegisterCodec(ModuleCdc)
	ModuleCdc.Seal()
}
