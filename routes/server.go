package routes

import (
	gbtConfig "GoBittrex/config"
	gbtCryptoDb "GoBittrex/databases"
	gbtEntity "GoBittrex/entity/cmc"
	gbtError "GoBittrex/error"
	"database/sql"
	"fmt"
	"net/http"
)

func RunServer() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/coin-info", handlerExample)
	http.Handle("/coin-info-gbt", helloHandlerExample())

	http.ListenAndServe(":"+gbtConfig.CRYPTO_API_PORT, nil)
}

func simpleSelectHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//if err != nil {
		//	gbtError.ShowError(err)
		//}

		//defer db.Close()

		var err = db.Ping()
		if err != nil {
			gbtError.ShowError(err)
		}

		// prepare statement for reading data
		rows, err := db.Query("SELECT id, name FROM names")
		if err != nil {
			gbtError.ShowError(err)
		}

		defer rows.Close()

		var id int
		var name string

		for rows.Next() {
			rows.Scan(&id, &name)
			fmt.Printf("%d : %s \n", id, name)
			fmt.Fprintf(w, "%d : %s \n", id, name)
		}
	})
}

func helloHandlerExample() http.Handler {
	db, err := gbtCryptoDb.Connect(gbtConfig.MYSQL_HOST, gbtConfig.MYSQL_PORT, gbtConfig.MYSQL_USERNAME, gbtConfig.MYSQL_PASSWORD, gbtConfig.MYSQL_DBNAME)
	if err != nil {
		gbtError.ShowError(err)
	}

	return simpleSelectHandler(db)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World from GoBittrex server!")
}

func handlerExample(w http.ResponseWriter, r *http.Request) {
	params, exists := r.URL.Query()["coin"]
	if !exists || len(params) < 1 {
		fmt.Println("Url Param 'coin' is missing")
		return
	}
	// Query()["coin"] will return an array of items, we only want the single item.
	coin := params[0]

	fmt.Println("Url Param 'coin' is: " + string(coin))

	resp, err := GetCoinInfoAPI(coin)
	if err != nil {
		fmt.Fprintf(w, "%v", "Error: "+err.Error())
	} else {
		switch v := resp.(type) {
		case []gbtEntity.CmcCoinInfo:
			fmt.Fprintf(w, v[0].Name+" "+v[0].MaxSupply)
		default:
			fmt.Fprintf(w, "%v", "Unsupported value in GetCmcInfo")
		}
	}
}
