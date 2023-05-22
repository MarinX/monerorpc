package monerorpc

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strconv"
)

// NewPaymentID64 generates a 64 bit payment ID
func NewPaymentID64() string {
	buf := make([]byte, 8)
	_, _ = rand.Read(buf)
	return hex.EncodeToString(buf)
}

// NewPaymentID256 generates a 256 bit payment ID (hex encoded).
func NewPaymentID256() string {
	buf := make([]byte, 32)
	_, _ = rand.Read(buf)
	return hex.EncodeToString(buf)
}

// XMRToDecimal converts a raw atomic XMR balance to a more
// human readable format.
func XMRToDecimal(xmr uint64) string {
	str0 := fmt.Sprintf("%013d", xmr)
	l := len(str0)
	return str0[:l-12] + "." + str0[l-12:]
}

// XMRToFloat64 converts raw atomic XMR to a float64
func XMRToFloat64(xmr uint64) float64 {
	return float64(xmr) / 1e12
}

// Float64ToXMR converts a float64 to a raw atomic XMR
func Float64ToXMR(xmr float64) uint64 {
	return uint64(xmr * 1e12)
}

// StringToXMR converts a string to a raw atomic XMR
func StringToXMR(xmr string) (uint64, error) {
	f, err := strconv.ParseFloat(xmr, 64)
	if err != nil {
		return 0, err
	}
	return uint64(f * 1e12), nil
}
