package routes

func SelectGeneralRoute(cmd string, args interface{}) bool {
	status := false
	switch cmd {
	case "runServer":
		RunServer()
		status = true
	default:
		status = false
	}
	return status
}
