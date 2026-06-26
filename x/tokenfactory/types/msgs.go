package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewDenom(creator, subdenom string) string {
	return DenomPrefix + "/" + creator + "/" + subdenom
}

func (m *MsgCreateDenom) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return ErrInvalidCreator
	}
	if len(m.Subdenom) == 0 {
		return ErrInvalidSubdenom
	}
	return nil
}

func (m *MsgCreateDenom) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

func NewMsgCreateDenom(sender, subdenom string) *MsgCreateDenom {
	return &MsgCreateDenom{Sender: sender, Subdenom: subdenom}
}

func (m *MsgMint) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return ErrInvalidCreator
	}
	coin, err := sdk.ParseCoinNormalized(m.Amount)
	if err != nil {
		return err
	}
	if !coin.IsPositive() {
		return ErrInvalidDenom
	}
	if m.MintTo != "" {
		if _, err := sdk.AccAddressFromBech32(m.MintTo); err != nil {
			return ErrInvalidCreator
		}
	}
	return nil
}

func (m *MsgMint) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

func NewMsgMint(sender, amount, mintTo string) *MsgMint {
	return &MsgMint{Sender: sender, Amount: amount, MintTo: mintTo}
}

func (m *MsgBurn) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return ErrInvalidCreator
	}
	coin, err := sdk.ParseCoinNormalized(m.Amount)
	if err != nil {
		return err
	}
	if !coin.IsPositive() {
		return ErrInvalidDenom
	}
	return nil
}

func (m *MsgBurn) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

func NewMsgBurn(sender, amount string) *MsgBurn {
	return &MsgBurn{Sender: sender, Amount: amount}
}

func (m *MsgChangeAdmin) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return ErrInvalidCreator
	}
	if _, err := sdk.AccAddressFromBech32(m.NewAdmin); err != nil {
		return ErrInvalidCreator
	}
	if len(m.Denom) == 0 {
		return ErrInvalidDenom
	}
	return nil
}

func (m *MsgChangeAdmin) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

func NewMsgChangeAdmin(sender, denom, newAdmin string) *MsgChangeAdmin {
	return &MsgChangeAdmin{Sender: sender, Denom: denom, NewAdmin: newAdmin}
}

func (m *MsgForceTransfer) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return ErrInvalidCreator
	}
	if _, err := sdk.AccAddressFromBech32(m.DestAddr); err != nil {
		return ErrInvalidCreator
	}
	coin, err := sdk.ParseCoinNormalized(m.Amount)
	if err != nil {
		return err
	}
	if !coin.IsPositive() {
		return ErrInvalidDenom
	}
	return nil
}

func (m *MsgForceTransfer) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

func NewMsgForceTransfer(sender, amount, destAddr string) *MsgForceTransfer {
	return &MsgForceTransfer{Sender: sender, Amount: amount, DestAddr: destAddr}
}
