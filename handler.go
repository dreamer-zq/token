package asset

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github/irismod/asset/internal/types"
)

// handle all "asset" type messages.
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case MsgIssueToken:
			return handleIssueToken(ctx, k, msg)
		case MsgEditToken:
			return handleMsgEditToken(ctx, k, msg)
		case MsgMintToken:
			return handleMsgMintToken(ctx, k, msg)
		case MsgTransferTokenOwner:
			return handleMsgTransferTokenOwner(ctx, k, msg)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized nft message type: %T", msg)
		}
	}
}

// handleIssueToken handles MsgIssueToken
func handleIssueToken(ctx sdk.Context, k Keeper, msg MsgIssueToken) (*sdk.Result, error) {
	// handle fee for token
	if err := k.DeductIssueTokenFee(ctx, msg.Owner, msg.Symbol); err != nil {
		return nil, err
	}

	if err := k.IssueToken(ctx, msg); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeIssueToken,
			sdk.NewAttribute(types.AttributeValueSymbol, msg.Symbol),
			sdk.NewAttribute(types.AttributeValueOwner, msg.Owner.String()),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

// handleMsgEditToken handles MsgEditToken
func handleMsgEditToken(ctx sdk.Context, k Keeper, msg MsgEditToken) (*sdk.Result, error) {
	if err := k.EditToken(ctx, msg); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeEditToken,
			sdk.NewAttribute(types.AttributeValueSymbol, msg.Symbol),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

// handleMsgTransferTokenOwner handles MsgTransferTokenOwner
func handleMsgTransferTokenOwner(ctx sdk.Context, k Keeper, msg MsgTransferTokenOwner) (*sdk.Result, error) {
	if err := k.TransferTokenOwner(ctx, msg); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeTransferTokenOwner,
			sdk.NewAttribute(types.AttributeValueSymbol, msg.Symbol),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

// handleMsgMintToken handles MsgMintToken
func handleMsgMintToken(ctx sdk.Context, k Keeper, msg MsgMintToken) (*sdk.Result, error) {
	if err := k.DeductMintTokenFee(ctx, msg.Owner, msg.Symbol); err != nil {
		return nil, err
	}

	if err := k.MintToken(ctx, msg); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMintToken,
			sdk.NewAttribute(types.AttributeValueSymbol, msg.Symbol),
			sdk.NewAttribute(types.AttributeValueAmount, strconv.FormatUint(msg.Amount, 10)),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
