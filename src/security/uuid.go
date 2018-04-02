package security

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

func UniqueId() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}

	return GetMD5Hash(base64.URLEncoding.EncodeToString(b))
}
