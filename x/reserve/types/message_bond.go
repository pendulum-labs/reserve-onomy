package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBond = "create_vault"

var _ sdk.Msg = &MsgBond{}

func NewMsgBond(creator string, coin string) *MsgBond {
	return &MsgBond{
		Creator: creator,
		Coin:    coin,
	}
}

func (msg *MsgBond) Route() string {
	return RouterKey
}

func (msg *MsgBond) Type() string {
	return TypeMsgBond
}

func (msg *MsgBond) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBond) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBond) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	coin, _ := sdk.ParseCoinNormalized(msg.Coin)
	if !coin.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount is not a valid Coin object")
	}

	return nil
}
