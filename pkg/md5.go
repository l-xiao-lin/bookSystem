package pkg

import (
	"crypto/md5"
	"encoding/hex"
)

const secret = "cisco46589"

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString([]byte(s))
}
