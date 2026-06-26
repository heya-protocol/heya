package tokenfactory

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"heya/x/tokenfactory/types"
)

type msgServer struct {
	keeper Keeper
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (s msgServer) CreateDenom(goCtx context.Context, msg *types.MsgCreateDenom) (*types.MsgCreateDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	denom := types.NewDenom(msg.Sender, msg.Subdenom)
	if _, exists := s.keeper.GetDenomAdmin(ctx, denom); exists {
		return nil, types.ErrDenomExists
	}
	s.keeper.SetDenomAdmin(ctx, denom, msg.Sender)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.ModuleName,
		sdk.NewAttribute("creator", msg.Sender),
		sdk.NewAttribute("denom", denom),
		sdk.NewAttribute("subdenom", msg.Subdenom),
	))
	return &types.MsgCreateDenomResponse{Denom: denom}, nil
}

func (s msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	coin, err := sdk.ParseCoinNormalized(msg.Amount)
	if err != nil {
		return nil, err
	}
	admin, exists := s.keeper.GetDenomAdmin(ctx, coin.Denom)
	if !exists {
		return nil, types.ErrDenomNotFound
	}
	if admin != msg.Sender {
		return nil, types.ErrUnauthorized
	}
	if err := s.keeper.bankKeeper.MintCoins(sdk.WrapSDKContext(ctx), types.ModuleName, sdk.NewCoins(coin)); err != nil {
		return nil, err
	}
	recipient := msg.MintTo
	if recipient == "" {
		recipient = msg.Sender
	}
	recipientAddr, err := sdk.AccAddressFromBech32(recipient)
	if err != nil {
		return nil, err
	}
	if err := s.keeper.bankKeeper.SendCoinsFromModuleToAccount(sdk.WrapSDKContext(ctx), types.ModuleName, recipientAddr, sdk.NewCoins(coin)); err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.ModuleName,
		sdk.NewAttribute("minter", msg.Sender),
		sdk.NewAttribute("recipient", recipient),
		sdk.NewAttribute("denom", coin.Denom),
		sdk.NewAttribute("amount", coin.Amount.String()),
	))
	return &types.MsgMintResponse{}, nil
}

func (s msgServer) Burn(goCtx context.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	coin, err := sdk.ParseCoinNormalized(msg.Amount)
	if err != nil {
		return nil, err
	}
	admin, exists := s.keeper.GetDenomAdmin(ctx, coin.Denom)
	if !exists {
		return nil, types.ErrDenomNotFound
	}
	if admin != msg.Sender {
		return nil, types.ErrUnauthorized
	}
	sender, _ := sdk.AccAddressFromBech32(msg.Sender)
	if err := s.keeper.bankKeeper.SendCoinsFromAccountToModule(sdk.WrapSDKContext(ctx), sender, types.ModuleName, sdk.NewCoins(coin)); err != nil {
		return nil, err
	}
	if err := s.keeper.bankKeeper.BurnCoins(sdk.WrapSDKContext(ctx), types.ModuleName, sdk.NewCoins(coin)); err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.ModuleName,
		sdk.NewAttribute("burner", msg.Sender),
		sdk.NewAttribute("denom", coin.Denom),
		sdk.NewAttribute("amount", coin.Amount.String()),
	))
	return &types.MsgBurnResponse{}, nil
}

func (s msgServer) ChangeAdmin(goCtx context.Context, msg *types.MsgChangeAdmin) (*types.MsgChangeAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	admin, exists := s.keeper.GetDenomAdmin(ctx, msg.Denom)
	if !exists {
		return nil, types.ErrDenomNotFound
	}
	if admin != msg.Sender {
		return nil, types.ErrUnauthorized
	}
	s.keeper.SetDenomAdmin(ctx, msg.Denom, msg.NewAdmin)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.ModuleName,
		sdk.NewAttribute("old_admin", msg.Sender),
		sdk.NewAttribute("new_admin", msg.NewAdmin),
		sdk.NewAttribute("denom", msg.Denom),
	))
	return &types.MsgChangeAdminResponse{}, nil
}

func (s msgServer) ForceTransfer(goCtx context.Context, msg *types.MsgForceTransfer) (*types.MsgForceTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	coin, err := sdk.ParseCoinNormalized(msg.Amount)
	if err != nil {
		return nil, err
	}
	admin, exists := s.keeper.GetDenomAdmin(ctx, coin.Denom)
	if !exists {
		return nil, types.ErrDenomNotFound
	}
	if admin != msg.Sender {
		return nil, types.ErrUnauthorized
	}
	destAddr, err := sdk.AccAddressFromBech32(msg.DestAddr)
	if err != nil {
		return nil, err
	}
	if err := s.keeper.bankKeeper.MintCoins(sdk.WrapSDKContext(ctx), types.ModuleName, sdk.NewCoins(coin)); err != nil {
		return nil, err
	}
	if err := s.keeper.bankKeeper.SendCoinsFromModuleToAccount(sdk.WrapSDKContext(ctx), types.ModuleName, destAddr, sdk.NewCoins(coin)); err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.ModuleName,
		sdk.NewAttribute("force_transfer_to", msg.DestAddr),
		sdk.NewAttribute("denom", coin.Denom),
		sdk.NewAttribute("amount", coin.Amount.String()),
	))
	return &types.MsgForceTransferResponse{}, nil
}
