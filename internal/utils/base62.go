package utils

import "strings"

/**
* This file contains all the utility functions related to base62 encoding and decoding.
* Note that the base62 reference cypher key should be loaded from the config file.
* @Author: Yinfeng Lu
* @Date: 2024-08-20
 */

// EncodeBase62 encodes an integer to a base62 string.
// The key is the reference cypher key for base62 encoding.
func EncodeBase62(id int64, key string) string {
	if id == 0 {
		return string(key[0])
	}
	sb := strings.Builder{}
	for id > 0 {
		sb.WriteByte(key[id%62])
		id /= 62
	}
	result := sb.String()
	return reverse(result)
}

// DecodeBase62 decodes a base62 string to an integer.
// The key is the reference cypher key for base62 encoding.
func DecodeBase62(s string, key string) int64 {
	result := int64(0)
	for _, c := range s {
		result = result*62 + int64(strings.Index(key, string(c)))
	}
	return result
}

func reverse(s string) string {
	r := []byte(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
