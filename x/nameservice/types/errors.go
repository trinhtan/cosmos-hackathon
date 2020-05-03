package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrNameDoesNotExist = sdkerrors.Register(ModuleName, 1, "name does not exist")
	ProductDoesNotExist = sdkerrors.Register(ModuleName, 2, "product does not exist")
)
