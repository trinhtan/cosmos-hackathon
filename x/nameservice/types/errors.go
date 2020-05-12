package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrNameDoesNotExist = sdkerrors.Register(ModuleName, 1, "name does not exist")

	ErrProductDoesNotExist  = sdkerrors.Register(ModuleName, 2, "product does not exist")
	ErrProductAlreadyExists = sdkerrors.Register(ModuleName, 3, "product already exists")

	ErrSellDoesNotExist  = sdkerrors.Register(ModuleName, 4, "sell does not exist")
	ErrSellAlreadyExists = sdkerrors.Register(ModuleName, 5, "sell already exists")

	ErrReservationDoesNotExist = sdkerrors.Register(ModuleName, 6, "reservation does not exist")
	ErrReservationAlreadyExists = sdkerrors.Register(ModuleName, 8, "reservation already exists")

	ErrNothingChanges = sdkerrors.Register(ModuleName, 7, "nothing changes")
)
