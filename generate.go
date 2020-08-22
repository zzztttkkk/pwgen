package main

import (
	"crypto/sha512"
	"fmt"
)

var lowers = "abcdefghijklmnopqrstuvwxyz"
var uppers = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var digits = "0123456789"
var symbols = "`~!@#$%^&*()_+-=[]{}:;,./<>?|"

func generate(hostname string, username string, secret string, length int) string {
	hash := sha512.New()
	hash.Write([]byte(fmt.Sprintf("%s:%s@%s", username, secret, hostname)))

	if length <= 0 {
		length = hash.Size()
	}
	if length > hash.Size() {
		length = hash.Size()
	}

	var buf []byte
	for _, b := range hash.Sum(nil)[:length] {
		v := int(b)
		if v >= 192 {
			buf = append(buf, lowers[v%len(lowers)])
		} else if v >= 128 {
			buf = append(buf, uppers[v%len(uppers)])
		} else if v >= 64 {
			buf = append(buf, digits[v%len(digits)])
		} else {
			buf = append(buf, symbols[v%len(symbols)])
		}
	}
	return string(buf)
}
