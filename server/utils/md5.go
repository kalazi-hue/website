package utils

import (
	"crypto/md5"
	"encoding/hex"
	"time"
	"math/rand"
)

func MD5V(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}

func GetRandInt() int32 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int31()%10000
}
