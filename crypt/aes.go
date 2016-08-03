package tokenaes

import (
	"crypto/aes"
	"strings"
	"bytes"
)

var key []byte = []byte{
	1, 2, 3, 4, 5, 6, 7, 8,
	9, 0, 1, 2, 3, 4, 5, 6,
	7, 8, 9, 0, 1, 2, 3, 4,
	5, 6, 7, 8, 9, 0, 1, 2,
}

func Encrypt(src string) (dst []byte) {

	cleartext := make([]byte, aes.BlockSize)
	ciphertext := make([]byte, aes.BlockSize)

	cip, _ := aes.NewCipher(key)
	tmpReader := strings.NewReader(src)

	for _, e := tmpReader.Read(cleartext); e == nil; _, e = tmpReader.Read(cleartext) {
		cip.Encrypt(ciphertext, cleartext)
		dst = append(dst, ciphertext...)
		cleartext = make([]byte, aes.BlockSize)
		ciphertext = make([]byte, aes.BlockSize)
	}
	return
}

func Decrypt(src []byte) (dst string) {

	cleartext := make([]byte, aes.BlockSize)
	ciphertext := make([]byte, aes.BlockSize)

	cip, _ := aes.NewCipher(key)
	tmpReader := bytes.NewReader(src)
	var rst []byte
	for _, e := tmpReader.Read(ciphertext); e == nil; _, e = tmpReader.Read(ciphertext) {
		cip.Decrypt(cleartext, ciphertext)
		rst = append(rst, cleartext...)
		cleartext = make([]byte, aes.BlockSize)
		ciphertext = make([]byte, aes.BlockSize)
	}
	
	return string(rst)
}

