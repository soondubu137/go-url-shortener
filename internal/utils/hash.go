package utils

/**
* This file contains all the utility functions related to hashing.
* @Author: Yinfeng Lu
* @Date: 2024-08-20
 */

import (
	"crypto/md5"
	"encoding/hex"
)

// GenerateMD5 generates the MD5 hash of the given string.
func GenerateMD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
