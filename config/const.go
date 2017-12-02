package config

const (
	API_BASE              = "https://bittrex.com/api/" // Bittrex API endpoint
	API_VERSION           = "v1.1"
	API_VERSION_V2        = "v2.0"
	WS_BASE               = "socket.bittrex.com" // Bittrex WS API endpoint
	WS_HUB                = "CoreHub"            // SignalR main hub
	API_PATH              = API_BASE + API_VERSION
	API_PATH_V2           = API_BASE + API_VERSION_V2
	FCM_URL               = "https://fcm.googleapis.com/v1/projects/%s/messages:send"
	FCM_PROJECT           = "gobittrex"
	FCM_CREDS             = "service-account.json"
	FCM_URL_LEGACY        = "https://fcm.googleapis.com/fcm/send"     // when using legacy HTTP POST
	FCM_SERVER_KEY_LEGACY = "xxxxx" // when using legacy HTTP POST
	API_KEY               = "xxxxx"
	API_SECRET            = "xxxxx"
)
