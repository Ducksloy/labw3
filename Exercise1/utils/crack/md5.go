package crack

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5Hex returns the lowercase hex-encoded MD5 checksum of s.
func MD5Hex(s string) string {
    sum := md5.Sum([]byte(s))
    return hex.EncodeToString(sum[:])
}
