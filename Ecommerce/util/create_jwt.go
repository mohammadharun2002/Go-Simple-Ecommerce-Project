package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub         int    `json:"sub"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func CreateJwt(secret string, data Payload) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	//convert to byte arrays
	byteArrHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	byteArrData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	//convert to Base64
	headerBase64 := Base64Encode(byteArrHeader)
	payloadBase64 := Base64Encode(byteArrData)

	message := headerBase64 + "." + payloadBase64

	byteArrSecret := []byte(secret)
	byteArrMessage := []byte(message)

	h := hmac.New(sha256.New, byteArrSecret)
	h.Write(byteArrMessage)

	signature := h.Sum(nil)
	signatureBase64 := Base64Encode(signature)

	jwt := headerBase64 + "." + payloadBase64 + "." + signatureBase64

	return jwt, nil
}

func Base64Encode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
