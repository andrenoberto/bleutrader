package main

import (
	"crypto/hmac"
	"fmt"
	"time"
	"strconv"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/hex"
	"crypto/sha512"
)

// API
const apiKey = "1a01622b0dd6ad7d4c3f03b86169f168"
const apiSecret = "659add3d23c11efd40fc9dd3d028b0a093b81925"
const rootUrl = "https://bleutrade.com/api/v2/"

// Routes
const balanceUrl = "account/getbalances"

func main() {
	signature, uri := getAPISign(balanceUrl)
	req, err := http.NewRequest("GET", uri, nil)
	errorHandler(err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("apisign", signature)
	resp, err := http.DefaultClient.Do(req)
	errorHandler(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	errorHandler(err)
	fmt.Println(string(body))
}

func getAPISign(s string) (string, string) {
	var (
		nonce = strconv.FormatInt(time.Now().UnixNano(), 10)
		uri   = rootUrl + s + "?apikey=" + apiKey + "&nonce=" + string(nonce)
	)
	signature := createHash(uri)
	return signature, uri
}

func createHash(s string) string {
	secret := []byte(apiSecret)
	message := []byte(s)
	mac := hmac.New(sha512.New, secret)
	mac.Write(message)
	hash := mac.Sum(nil)
	encodedHash := hex.EncodeToString(hash)
	return encodedHash
}

func errorHandler(err error) {
	if err != nil {
		log.Fatal("Fatal error thrown:", err)
	}
}
