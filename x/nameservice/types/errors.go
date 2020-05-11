package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrNameDoesNotExist        = sdkerrors.Register(ModuleName, 1, "name does not exist")
	ErrProductDoesNotExist     = sdkerrors.Register(ModuleName, 2, "product does not exist")
	ErrSellDoesNotExist        = sdkerrors.Register(ModuleName, 3, "sell does not exist")
	ErrReservationDoesNotExist = sdkerrors.Register(ModuleName, 4, "reservation does not exist")
	ErrProductAlreadyExists    = sdkerrors.Register(ModuleName, 5, "product already exist")
	ErrNothingChanges          = sdkerrors.Register(ModuleName, 6, "Nothing changes")
)
