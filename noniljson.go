package noniljson

import (
	"go/types"
)

// isNullableType checks if the given type is considered nullable for JSON purposes.
func isNullableType(typ types.Type) bool {
	switch typ.Underlying().(type) {
	case *types.Pointer, *types.Slice, *types.Map, *types.Interface:
		return true
	default:
		return false
	}
}
