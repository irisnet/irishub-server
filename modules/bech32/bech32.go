package bech32

import (
	"github.com/btcsuite/btcutil/bech32"
	"encoding/hex"
	"fmt"
)

// Bech32 prefixes
const (
	Bech32PrefixAccAddr = "cosmosaccaddr"
	Bech32PrefixAccPub  = "cosmosaccpub"
	Bech32PrefixValAddr = "cosmosvaladdr"
	Bech32PrefixValPub  = "cosmosvalpub"
)

// convertAndEncode converts from a base64 encoded byte string to base32 encoded byte string and then to bech32
func convertAndEncode(hrp string, data []byte) (string, error) {
	converted, err := bech32.ConvertBits(data, 8, 5, true)
	if err != nil {
		return "", fmt.Errorf("encoding bech32 failed")
	}
	return bech32.Encode(hrp, converted)
}

// decodeString returns the bytes represented by the hexadecimal string s.
//
// decodeString expects that src contain only hexadecimal
// characters and that src should have an even length.
// If the input is malformed, decodeString returns a string
// containing the bytes decoded before the error.
func decodeString(s string) ([]byte, error) {
	src := []byte(s)
	// We can use the source slice itself as the destination
	// because the decode loop increments by one and then the 'seen' byte is not used anymore.
	n, err := hex.Decode(src, src)
	return src[:n], err
}

// convert hex to bech32
func ConvertHexToBech32(addr string) (string, error) {
	addrBytes, err := decodeString(addr)
	if err != nil {
		return "", fmt.Errorf("decode string to bytes failed")
	}

	return convertAndEncode(Bech32PrefixAccAddr, addrBytes)
}
