package main

import (
	"github.com/arashout/mlpapi"
	"github.com/arashout/roborooney"
)

//go:generate echo '{"apiToken":"","channelId":""}' > config.json

var arrayPitches = []*mlpapi.Pitch{
	&mlpapi.Pitch{
		VenueID:   "34933",
		VenuePath: "three-corners-adventure-playground/football-5-a-side-34933",
		City:      "london",
	},
}

func main() {
	robo := roborooney.NewRobo(arrayPitches)
	robo.Connect()
}
