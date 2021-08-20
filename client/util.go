package client

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

// Md5 generate md5 for string
func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return hex.EncodeToString(h.Sum(nil))
}
