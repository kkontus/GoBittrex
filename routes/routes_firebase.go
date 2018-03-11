package routes

import (
	"fmt"
	gbtFcmDb "GoBittrex/databases"
	gbtFcm "GoBittrex/fcm"
	gbtError "GoBittrex/error"
)

func SelectFirebaseRoute(cmd string, args interface{}) bool {
	status := false
	switch cmd {
	case "sendPush":
		sendPush()
		status = true
	case "startRealtimeDatabase":
		startRealtimeDatabase()
		status = true
	default:
		status = false
	}
	return status
}

func sendPush() {
	gbtFcm.SendPush(false) // refactor this to send message instead hardcoded data
}

func startRealtimeDatabase() {
	//gbtFcmDb.RealtimeDb()

	client, err := gbtFcmDb.RealtimeDbClient()
	if err != nil {
		gbtError.ShowError(err)
	} else {
		fmt.Println("Connection successful.")
	}

	//gbtFcmDb.SetUsers(client)
	//gbtFcmDb.SetCurrencies(client)
	gbtFcmDb.GetCurrencies(client)
	gbtFcmDb.GetUsers(client)

}
