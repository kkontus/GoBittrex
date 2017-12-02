package fcm

func FcmBody() string {
	//return topicBodyExample()
	return tokenBodyExample()
}

func FcmBodyLegacy() string {
	return legacyApiExample()
}

func topicBodyExample() string {
	json := `{
					"message": {
						"topic": "subscriber-updates",
						"notification": {
							"body": "This week's edition is now available.",
							"title": "NewsMagazine.com"
						},
						"data": {
							"volume": "3.21.15",
							"contents": "http://www.news-magazine.com/world-week/21659772"
						},
						"android": {
							"priority": "normal"
						},
						"apns": {
							"headers": {
								"apns-priority": "5"
							}
						},
						"webpush": {
							"headers": {
								"Urgency": "high"
							}
						}
					}
				}`

	return json
}

func tokenBodyExample() string {
	json := `{
				"message": {
					"token": "fvztYIIwJ7Y:APA91bEhzhnbv5B8OCZQejOUDbjaW-87CZUUJv92yNeVIDGmpNLpSJ8JRhg5s1ESVaOGDCI0dGjGJdyWmhNspOyD_OkQPZbkETjV_Bv0MgR1fIAR766AUBd7hzmtmCDoA9bv2nXB17_d",
					"notification": {
						"body": "This week's edition is now available.",
						"title": "NewsMagazine.com"
					},
					"data": {
						"volume": "3.21.15",
						"contents": "http://www.news-magazine.com/world-week/21659772"
					},
					"android": {
						"priority": "normal"
					},
					"apns": {
						"headers": {
							"apns-priority": "5"
						}
					},
					"webpush": {
						"headers": {
							"Urgency": "high"
						}
					}
				}
			}`

	return json
}

func legacyApiExample() string {
	json := `{
				"registration_ids":["fvztYIIwJ7Y:APA91bEhzhnbv5B8OCZQejOUDbjaW-87CZUUJv92yNeVIDGmpNLpSJ8JRhg5s1ESVaOGDCI0dGjGJdyWmhNspOyD_OkQPZbkETjV_Bv0MgR1fIAR766AUBd7hzmtmCDoA9bv2nXB17_d"],
				"data": {
                	"Coin" : "BTC",
                	"Price": "10000",
           		}
			}`

	return json
}
