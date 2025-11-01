package crack

import (
	"crypto/sha512"
	"encoding/hex"
)

// SHA512Hex returns the lowercase hex-encoded SHA-512 checksum of s.
func SHA512Hex(s string) string {
    sum := sha512.Sum512([]byte(s))
    return hex.EncodeToString(sum[:])
}
