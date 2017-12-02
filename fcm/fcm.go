package fcm

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"bytes"
	"io"
	"golang.org/x/oauth2/google"
	"golang.org/x/net/context"
	gbtConfig "GoBittrex/config"
	gbtUtil "GoBittrex/util"
	gbtError "GoBittrex/error"
)

func NewRequestFcm(method string, url string, body io.Reader) (response *http.Response, err error) {
	cred, err := ioutil.ReadFile(gbtConfig.FCM_CREDS)
	if err != nil {
		return nil, err
	}

	conf, err := google.JWTConfigFromJSON(cred, "https://www.googleapis.com/auth/firebase.messaging")
	if err != nil {
		return nil, err
	}

	token, err := conf.TokenSource(context.Background()).Token()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	authorization := fmt.Sprintf("%s %s", token.TokenType, token.AccessToken)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authorization)

	client := conf.Client(context.Background())
	return client.Do(req)
}

func NewRequestFcmLegacy(method string, url string, body io.Reader) (response *http.Response, err error) {
	req, err := http.NewRequest(method, url, body)

	authorization := fmt.Sprintf("key=%s", gbtConfig.FCM_SERVER_KEY_LEGACY)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authorization)

	client := &http.Client{}
	return client.Do(req)
}

func SendPush(useLegacyApi bool) {
	if useLegacyApi {
		URL := gbtConfig.FCM_URL_LEGACY
		body := []byte(FcmBodyLegacy())

		resp, err := NewRequestFcmLegacy(http.MethodPost, URL, bytes.NewBuffer(body))
		handleResponse(err, resp)
	} else {
		URL := fmt.Sprintf(gbtConfig.FCM_URL, gbtConfig.FCM_PROJECT)
		body := []byte(FcmBody())

		resp, err := NewRequestFcm(http.MethodPost, URL, bytes.NewBuffer(body))
		handleResponse(err, resp)
	}
}
func handleResponse(err error, resp *http.Response) {
	if err != nil {
		gbtError.ShowError(err)
	} else {
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)

		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			gbtError.ShowError(err)
		}

		fmt.Printf("%s\n", gbtUtil.JsonPrettyPrint(string(contents)))
	}
}
