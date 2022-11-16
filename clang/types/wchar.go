//go:build !windows
// +build !windows

package types

import (
	"go/types"
)

var WcharT = types.Typ[types.Int32]
