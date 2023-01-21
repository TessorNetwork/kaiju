package ante

// import (
// 	"errors"
// 	"fmt"

// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

// 	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
// 	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
// 	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
// 	iidkeeper "github.com/tessornetwork/kaiju/x/iid/keeper"
// )

// type CheckTxForIncompatibleMsgsDecorator struct {
// }

// func NewCheckTxForIncompatibleMsgsDecorator() sdk.AnteDecorator {
// 	return CheckTxForIncompatibleMsgsDecorator{}
// }

// func (dec CheckTxForIncompatibleMsgsDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
// 	// feeTx, ok := tx.(KaijuFeeTx)

// 	//Check if feegranter is set. or if not kaijuFeeTx
// 	// if !ok || feeTx.FeeGranter() != nil {
// 	// 	// return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a IIDTx")
// 	// 	panic("")
// 	// }

// 	// events := sdk.Events{sdk.NewEvent(sdk.EventTypeTx,
// 	// 	sdk.NewAttribute(sdk.AttributeKeyFee, feeTx.GetFee().String()),
// 	// )}
// 	// ctx.EventManager().EmitEvents(events)

// 	return next(ctx, tx, simulate)
// }

// type KaijuFeeHandlerDecorator struct {
// 	iidKeeper         iidkeeper.Keeper
// 	accountKeeper     authante.AccountKeeper
// 	bankKeeper        bankkeeper.Keeper
// 	defaultFeeHandler authante.DeductFeeDecorator
// }

// func NewKaijuFeeHandlerDecorator(iidKeeper iidkeeper.Keeper, accountKeeper authante.AccountKeeper, bankKeeper bankkeeper.Keeper, defaultFeeHandler authante.DeductFeeDecorator) sdk.AnteDecorator {
// 	return KaijuFeeHandlerDecorator{
// 		iidKeeper:         iidKeeper,
// 		accountKeeper:     accountKeeper,
// 		bankKeeper:        bankKeeper,
// 		defaultFeeHandler: defaultFeeHandler,
// 	}
// }

// func (dec KaijuFeeHandlerDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
// 	feeTx, ok := tx.(KaijuFeeTx)

// 	//Check if feegranter is set. or if not kaijuFeeTx
// 	if !ok || feeTx.FeeGranter() != nil {
// 		// return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a IIDTx")
// 		return dec.defaultFeeHandler.AnteHandle(ctx, tx, simulate, next)
// 	}

// 	if addr := dec.accountKeeper.GetModuleAddress(authtypes.FeeCollectorName); addr == nil {
// 		return ctx, fmt.Errorf("fee collector module account (%s) has not been set", authtypes.FeeCollectorName)
// 	}

// 	kaijuFeeMsgs := feeTx.GetFeePayerMsgs()

// 	kaijuMsgCount := len(kaijuFeeMsgs)
// 	if len(feeTx.GetMsgs()) != kaijuMsgCount && kaijuMsgCount == 1 {
// 		return ctx, sdkerrors.Wrapf(errors.New("only one custom fee handler msg allowed per transaction"), "expted 1 and got %d", kaijuMsgCount)
// 	}

// 	feePayer, err := kaijuFeeMsgs[0].FeePayerFromIid(ctx, dec.accountKeeper, dec.iidKeeper)
// 	if err != nil {
// 		return ctx, sdkerrors.Wrap(sdkerrors.ErrUnknownAddress, "fee payer does not exist")
// 	}

// 	// // deduct the fees
// 	if !feeTx.GetFee().IsZero() {
// 		err = authante.DeductFees(dec.bankKeeper, ctx, feePayer.feePayerAccount, feeTx.GetFee())
// 		if err != nil {
// 			return ctx, err
// 		}
// 	}

// 	events := sdk.Events{sdk.NewEvent(sdk.EventTypeTx,
// 		sdk.NewAttribute(sdk.AttributeKeyFee, feeTx.GetFee().String()),
// 	)}
// 	ctx.EventManager().EmitEvents(events)

// 	return next(ctx, tx, simulate)
// }
