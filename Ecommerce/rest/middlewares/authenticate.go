package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"net/http"
	"strings"

	"ecommerse/util"
)

func (m *Middlewares) AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//parse jwt
		//parse header and payload or claims
		//hmacSHA256 algorithm -> hash(header, payload, secret key)
		//parse signature part from the jwt
		//if the signature and the hash is same => forward to create products
		//otherwise 401 status code with Unauthorised

		header := r.Header.Get("Authorization")

		if header == "" {
			http.Error(w, "Unthorised", http.StatusUnauthorized)
			return
		}
		headerArr := strings.Split(header, " ")
		if len(headerArr) != 2 {
			http.Error(w, "Unthorised", http.StatusUnauthorized)
			return
		}
		accessToken := headerArr[1]

		tokenParts := strings.Split(accessToken, ".")
		if len(tokenParts) != 3 {
			http.Error(w, "Unthorised", http.StatusUnauthorized)
			return
		}
		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		signature := tokenParts[2]

		message := jwtHeader + "." + jwtPayload

		byteArrSecret := []byte(m.cnf.JwtSecretKey)
		byteArrMesage := []byte(message)

		h := hmac.New(sha256.New, byteArrSecret)
		h.Write(byteArrMesage)

		hash := h.Sum(nil)
		newSignature := util.Base64Encode(hash)

		if newSignature != signature {
			http.Error(w, "Unthorised Tui Hacker", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
