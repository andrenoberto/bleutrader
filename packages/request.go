package packages

import (
	"net/http"
	"io/ioutil"
	"io"
)

func RequestHandler(method string, uri string, body io.Reader, signature string) []byte {
	req, err := http.NewRequest(method, uri, body)
	ErrorHandler(err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("apisign", signature)
	resp, err := http.DefaultClient.Do(req)
	ErrorHandler(err)
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	ErrorHandler(err)
	return []byte(string(responseBody))
}
