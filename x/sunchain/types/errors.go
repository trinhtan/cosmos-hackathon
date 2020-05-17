package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidBasicMsg        = sdkerrors.Register(ModuleName, 1, "InvalidBasicMsg")
	ErrBadDataValue           = sdkerrors.Register(ModuleName, 2, "BadDataValue")
	ErrUnauthorizedPermission = sdkerrors.Register(ModuleName, 3, "UnauthorizedPermission")
	ErrItemDuplication        = sdkerrors.Register(ModuleName, 4, "ItemDuplication")
	ErrItemNotFound           = sdkerrors.Register(ModuleName, 5, "ItemNotFound")
	ErrInvalidState           = sdkerrors.Register(ModuleName, 6, "InvalidState")
	ErrBadWasmExecution       = sdkerrors.Register(ModuleName, 7, "BadWasmExecution")
	ErrOnlyOneDenomAllowed    = sdkerrors.Register(ModuleName, 8, "OnlyOneDenomAllowed")
	ErrInvalidDenom           = sdkerrors.Register(ModuleName, 9, "InvalidDenom")
	ErrUnknownClientID        = sdkerrors.Register(ModuleName, 10, "UnknownClientID")

	ErrProductDoesNotExist  = sdkerrors.Register(ModuleName, 11, "product does not exist")
	ErrProductAlreadyExists = sdkerrors.Register(ModuleName, 12, "product already exists")

	ErrSellDoesNotExist  = sdkerrors.Register(ModuleName, 13, "sell does not exist")
	ErrSellAlreadyExists = sdkerrors.Register(ModuleName, 14, "sell already exists")

	ErrReservationDoesNotExist  = sdkerrors.Register(ModuleName, 15, "reservation does not exist")
	ErrReservationAlreadyExists = sdkerrors.Register(ModuleName, 16, "reservation already exists")
	ErrReservationNotDecided    = sdkerrors.Register(ModuleName, 17, "reservation not decided")
)
