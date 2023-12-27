package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// StoredTempTradeKeyPrefix is the prefix to retrieve all StoredTempTrade
	StoredTempTradeKeyPrefix = "StoredTempTrade/value/"
)

// StoredTempTradeKey returns the store key to retrieve a StoredTempTrade from the index fields
func StoredTempTradeKey(
	tradeIndex string,
) []byte {
	var key []byte

	tradeIndexBytes := []byte(tradeIndex)
	key = append(key, tradeIndexBytes...)
	key = append(key, []byte("/")...)

	return key
}
