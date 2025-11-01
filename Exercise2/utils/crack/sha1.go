package crack

import (
	"crypto/sha1"
	"encoding/hex"
)

// SHA1Hex returns the lowercase hex-encoded SHA1 checksum of s.
func SHA1Hex(s string) string {
    sum := sha1.Sum([]byte(s))
    return hex.EncodeToString(sum[:])
}
