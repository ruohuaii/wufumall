package hashutil

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

/**
* Created by : GoLand
* User: self-denial
* Date: 2023/4/24
* Time: 11:41
**/

//HMACWithHex 使用HMAC加密字符串,返回十六进制字符串
func HMACWithHex(key []byte, salt []byte) (string, error) {
	sum, err := HMAC(key, salt)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(sum), nil
}

//HMACWithBase64 使用HMAC加密字符串,返回base64格式字符串
func HMACWithBase64(key []byte, salt []byte) (string, error) {
	sum, err := HMAC(key, salt)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sum), nil
}

//HMAC 使用HMAC加密,返回已加密的字节片
func HMAC(key []byte, salt []byte) ([]byte, error) {
	h := hmac.New(sha256.New, key)
	_, err := h.Write(key)
	if err != nil {
		return nil, err
	}
	return h.Sum(salt), nil
}
