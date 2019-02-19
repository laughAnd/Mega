package model

import (
	"crypto/md5"
	"encoding/hex"
)

func GeneratePasswordHash(pwd string) string{
	hasher := md5.New()
	hasher.Write([]byte(pwd))
	pwdhash := hex.EncodeToString(hasher.Sum(nil))
	return pwdhash
}

