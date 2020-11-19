package app

import (
  "github.com/tendermint/tendermint/libs/log"
  "github.com/cosmos/cosmos-sdk/x/auth"
  "github.com/cosmos/cosmos-sdk/codec"
  "github.com/cosmos/cosmos-sdk/x/auth"
  "github.com/cosmos/cosmos-sdk/x/bank"
  "github.com/cosmos/cosmos-sdk/x/genutil"
  "github.com/cosmos/cosmos-sdk/x/params"
  "github.com/cosmos/cosmos-sdk/x/staking"
  "github.com/cosmos/cosmos-sdk/x/supply"

  bam "github.com/cosmos/cosmos-sdk/baseapp"
  dbm "github.com/tendermint/tendermint/libs/db"
)

const (
    appName = "nameservice"
)

var (
	ModuleBasics    = module.NewBasicManager(
		genutil.AppModuleBasic{},
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
		params.AppModuleBasic{},
		supply.AppModuleBasic{},
		nameservice.AppModuleBasic{},
		// this line is used by starport scaffolding # 2
	)
)


type nameServiceApp struct {
    *bam.BaseApp
}


func MakeCodec() *codec.Codec {
	var cdc = codec.New()

	ModuleBasics.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)

	return cdc.Seal()
}

func NewNameServiceApp(logger log.Logger, db dbm.DB) *nameServiceApp {

    // First define the top level codec that will be shared by the different modules. Note: Codec will be explained later
    cdc := MakeCodec()

    // BaseApp handles interactions with Tendermint through the ABCI protocol
    bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc))

    var app = &nameServiceApp{
        BaseApp: bApp,
        cdc:     cdc,
    }

    return app
}
