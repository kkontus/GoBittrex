package config

const (
	API_BASE                  = "https://bittrex.com/api/" // Bittrex API endpoint
	API_VERSION               = "v1.1"
	API_VERSION_V2            = "v2.0"
	WS_BASE                   = "socket.bittrex.com" // Bittrex WS API endpoint
	WS_HUB                    = "CoreHub"            // SignalR main hub
	API_PATH                  = API_BASE + API_VERSION
	API_PATH_V2               = API_BASE + API_VERSION_V2
	FCM_URL                   = "https://fcm.googleapis.com/v1/projects/%s/messages:send"
	FCM_PROJECT               = "xxxxx"
	FCM_CREDS                 = "service-account.json"
	FCM_REALTIME_CREDS        = "https://" + FCM_PROJECT + ".firebaseio.com/restricted.json"
	FCM_URL_LEGACY            = "https://fcm.googleapis.com/fcm/send"     // when using legacy HTTP POST
	FCM_SERVER_KEY_LEGACY     = "xxxxx" // when using legacy HTTP POST
	API_KEY                   = "xxxxx"
	API_SECRET                = "xxxxx"
	API_BASE_COINMARKETCAP    = "https://api.coinmarketcap.com/" // Coinmarketcap API endpoint
	API_VERSION_COINMARKETCAP = "v1"
	API_PATH_COINMARKETCAP    = API_BASE_COINMARKETCAP + API_VERSION_COINMARKETCAP
	API_BASE_COINMARKETCAL    = "https://www.coinmarketcal.com/" // Coinmarketcal API endpoint
	API_VERSION_COINMARKETCAL = "api"
	API_PATH_COINMARKETCAL    = API_BASE_COINMARKETCAL + API_VERSION_COINMARKETCAL
	MYSQL_HOST                = "127.0.0.1" // or localhost
	MYSQL_PORT                = "3306"
	MYSQL_USERNAME            = "xxxxx"
	MYSQL_PASSWORD            = "xxxxx"
	MYSQL_DBNAME              = "crypto"
	CRYPTO_API_PORT           = "8080"
	DEFAULT_CURRENCY          = "USD"
)

var RoutesBittrex = []string{"getMarkets", "getCurrencies", "getTicks", "getOpenOrders", "getOrderBook", "startTrailing"}
var RoutesCmc = []string{"getCoinInfo", "getCoinsInfo"}
var RoutesCmcal = []string{"getCmcalCoins", "getCmcalCategories", "getCmcalEvents"}
var RoutesFirebase = []string{"sendPush", "startRealtimeDatabase"}
var RoutesGeneral = []string{"runServer"}
