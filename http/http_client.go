package http

import (
	"net/http"
	"time"
	"fmt"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"io"
	gbtConfig "GoBittrex/config"
)

func createClient() *http.Client {
	return &http.Client{}
}

func NewRequest(method string, url string, body io.Reader, authRequired bool) (response *http.Response, err error) {
	client := createClient()

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	if authRequired {
		nonce := time.Now().UnixNano()
		q := req.URL.Query()
		q.Set("apikey", gbtConfig.API_KEY)
		q.Set("nonce", fmt.Sprintf("%d", nonce))
		req.URL.RawQuery = q.Encode()
		mac := hmac.New(sha512.New, []byte(gbtConfig.API_SECRET))
		_, err = mac.Write([]byte(req.URL.String()))
		sig := hex.EncodeToString(mac.Sum(nil))
		req.Header.Add("apisign", sig)
	}

	return client.Do(req)
}
