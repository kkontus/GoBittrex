package http

import (
	"net/http"
	"time"
	"fmt"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"io"
	. "GoBittrex/config"
)

func createClient() *http.Client {
	return &http.Client{}
}

func NewRequest(method string, url string, body io.Reader, authRequired AuthType) (response *http.Response, err error) {
	client := createClient()

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	if authRequired == BITTREX {
		nonce := time.Now().UnixNano()
		q := req.URL.Query()
		q.Set("apikey", API_KEY)
		q.Set("nonce", fmt.Sprintf("%d", nonce))
		req.URL.RawQuery = q.Encode()
		mac := hmac.New(sha512.New, []byte(API_SECRET))
		_, err = mac.Write([]byte(req.URL.String()))
		sig := hex.EncodeToString(mac.Sum(nil))
		req.Header.Add("apisign", sig)
	} else if authRequired == NONE {

	}

	return client.Do(req)
}
