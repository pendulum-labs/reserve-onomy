package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DenomKeyPrefix is the prefix to retrieve all Drop
	DenomKeyPrefix = "Denom/"
	// DenomsKeyPrefix is the prefix to retrieve all Denoms of Owner
	DenomsKeyPrefix = "Denom/Owner/"
)

// DropKey returns the store key to retrieve a Drop from the index fields
func DenomKey(
	base string,
) []byte {
	var key []byte

	baseBytes := []byte(base)
	key = append(key, baseBytes...)
	key = append(key, []byte("/")...)

	return key
}
