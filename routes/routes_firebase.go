package routes

import (
	gbtFcmDb "GoBittrex/databases"
	gbtFcm "GoBittrex/fcm"
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
	gbtFcmDb.RealtimeDB()
}
