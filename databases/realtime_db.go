package databases

import (
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"
	"firebase.google.com/go"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"firebase.google.com/go/db"
	"firebase.google.com/go/auth"
	"log"
	gbtFcm "GoBittrex/fcm"
	gbtConfig "GoBittrex/config"
	gbtError "GoBittrex/error"
	gbtUtil "GoBittrex/util"
	gbtEntity "GoBittrex/entity/frd"
)

func RealtimeDbClient() (*db.Client, error) {
	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: gbtConfig.FCM_DATABASE_URL,
	}
	// Fetch the service account key JSON file contents
	opt := option.WithCredentialsFile(gbtConfig.FCM_CREDS)

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		gbtError.ShowError(err)
	}

	client, err := app.Database(ctx)

	return client, err
}

func RealtimeAuthClient() (*auth.Client, error) {
	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: gbtConfig.FCM_DATABASE_URL,
	}
	// Fetch the service account key JSON file contents
	opt := option.WithCredentialsFile(gbtConfig.FCM_CREDS)

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		gbtError.ShowError(err)
	}

	client, err := app.Auth(ctx)

	return client, err
}

func GetUsers(client *db.Client) {
	// get users
	ctx := context.Background()
	refUsers := client.NewRef("users")
	var users map[string]gbtEntity.UserProfile
	err := refUsers.Get(ctx, &users)
	if err != nil {
		gbtError.ShowError(err)
	}
	for k, v := range users {
		fmt.Println("k: ", k, "v: ", v.UserProfile.Email)
	}
}

func GetCurrencies(client *db.Client) error {
	// get currencies
	ctx := context.Background()
	refCurr := client.NewRef("currencies")
	var currencies map[string]gbtEntity.Currency
	err := refCurr.Get(ctx, &currencies)
	if err != nil {
		gbtError.ShowError(err)
	}
	for k, v := range currencies {
		fmt.Println("k: ", k, "v: ", v.Title)
	}
	return err
}

func SetUsers(client *db.Client) {
	ctx := context.Background()
	refUsers := client.NewRef("users")

	u := gbtEntity.User{
		Email:     "kristijan.kontus@gmail.com",
		Firstname: "Kristijan",
		Lastname:  "Kontus",
		Username:  "kkontus"}

	up := gbtEntity.UserProfile{UserProfile: u}

	fmt.Println(u)
	fmt.Println(up)

	clientAuth, err := RealtimeAuthClient()

	ur, err := clientAuth.GetUserByEmail(ctx, u.Email)
	if err != nil {
		log.Fatalf("Error getting user by email %s: %v\n", u.Email, err)
	}
	log.Printf("Successfully fetched user data: %v\n", ur)

	if err := refUsers.Child(ur.UID).Set(ctx, up); err != nil {
		log.Fatalln("Error setting value:", err)
	}

}

func SetCurrencies(client *db.Client) {
	ctx := context.Background()
	refCurrencies := client.NewRef("currencies")

	c := gbtEntity.Currency{
		Id: "metal",
		//Title:    "Metal",
		Symbol:      "MTL",
		Name:        "Metal",
		Description: "Metal is PPoP cryptocurrency....."}

	fmt.Println(c)

	if err := refCurrencies.Child("dgr3jJDHAGtvdfv43arg").Set(ctx, c); err != nil {
		log.Fatalln("Error setting value VVVV:", err)
	}

}

func RealtimeDb() {
	URL := fmt.Sprintf("%s", gbtConfig.FCM_REALTIME_CREDS)
	resp, err := gbtFcm.NewRequestFcm(http.MethodGet, URL, nil)

	handleResponseRealtimeDB(err, resp)

	saveCurrencies()
	fetchCurrencies()
}

func handleResponseRealtimeDB(err error, resp *http.Response) {
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

func saveCurrencies() {
	URL := "https://crypto-curr-exch.firebaseio.com/currencies.json"
	body := []byte(postExample())

	resp, err := gbtFcm.NewRequestFcm(http.MethodPost, URL, bytes.NewBuffer(body))
	handleResponseRealtimeDB(err, resp)
}

func fetchCurrencies() {
	URL := "https://crypto-curr-exch.firebaseio.com/currencies.json"
	resp, err := gbtFcm.NewRequestFcm(http.MethodGet, URL, nil)
	handleResponseRealtimeDB(err, resp)
}

func postExample() string {
	json := `{
  				"description": "Not so good coin",
				"id": "ripple",
 				"title": "Ripple"
			}`

	return json
}
