package packages

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"Bleu/environment"
	"strconv"
	"time"
)

func createHash(s string) string {
	secret := []byte(environment.ApiSecret)
	message := []byte(s)
	mac := hmac.New(sha512.New, secret)
	mac.Write(message)
	hash := mac.Sum(nil)
	encodedHash := hex.EncodeToString(hash)
	return encodedHash
}

func GetAPISign(s ...string) (string, string) {
	var (
		nonce = strconv.FormatInt(time.Now().UnixNano(), 10)
		uri   = environment.RootUrl + s[0] + "?apikey=" + environment.ApiKey
	)
	uri += "&nonce=" + string(nonce)
	if len(s) > 1 && len(s[1]) > 0 {
		uri += s[1]
	}
	signature := createHash(uri)
	return signature, uri
}
