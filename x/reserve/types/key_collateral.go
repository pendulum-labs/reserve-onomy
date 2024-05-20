package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// CollateralKeyPrefix is the prefix to retrieve all Drop
	CollateralKeyPrefix = "Collateral/"
	// CollateralsKeyPrefix is the prefix to retrieve all Collaterals of Owner
	CollateralsKeyPrefix = "Collateral/Owner/"
)

// DropKey returns the store key to retrieve a Drop from the index fields
func CollateralKey(
	base string,
) []byte {
	var key []byte

	baseBytes := []byte(base)
	key = append(key, baseBytes...)
	key = append(key, []byte("/")...)

	return key
}
