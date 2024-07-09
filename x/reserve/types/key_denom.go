package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DenomKeyPrefix is the prefix to retrieve all Drop
	DenomKeyPrefix = "Denom/"
	// DenomsKeyPrefix is the prefix to retrieve all Denoms of Owner
	DenomsKeyPrefix = "Denom/Owner/"
	// DenomsKeyPrefix is the prefix to retrieve all Denoms of Owner
	BondedKeyPrefix = "Bonded/"
)

// DenomKey returns the store key to retrieve a Denom from the index fields
func DenomKey(
	base string,
) []byte {
	var key []byte

	baseBytes := []byte(base)
	key = append(key, baseBytes...)
	key = append(key, []byte("/")...)

	return key
}

// BondedKey returns the store key to retrieve a Bonded Denom Name from the index fields
func BondedKey(
	base string,
) []byte {
	var key []byte

	baseBytes := []byte(base)
	key = append(key, baseBytes...)
	key = append(key, []byte("/")...)

	return key
}
