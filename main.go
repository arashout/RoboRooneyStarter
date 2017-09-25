package main

import (
	"github.com/arashout/mlpapi"
	"github.com/arashout/roborooney"
)

//go:generate echo '{"apiToken":"","channelId":""}' > config.json

func main() {
	rules := []mlpapi.SlotFilter{
		func(slot mlpapi.Slot) bool {
			// Return only the available slots
			return slot.Attributes.Availabilities > 0
		},
		func(slot mlpapi.Slot) bool {
			// Only show slots after 4pm
			return slot.Attributes.Starts.Hour() > 16
		},
		func(slot mlpapi.Slot) bool {
			// Only show slots before 8pm
			return slot.Attributes.Starts.Hour() < 20
		},
		func(slot mlpapi.Slot) bool {
			// Week starts at Sunday == index 0
			return slot.Attributes.Starts.Weekday() != 6 && slot.Attributes.Starts.Weekday() != 0
		},
	}
	pitches := []mlpapi.Pitch{
		mlpapi.Pitch{
			ID:   "34933",
			Path: "three-corners-adventure-playground/football-5-a-side-34933",
			City: "london",
			Name: "Three Corners",
		},
		mlpapi.Pitch{
			ID:   "32208",
			Path: "finsbury-leisure-centre/football-5-a-side-32208",
			City: "london",
			Name: "Finsbury Leisure Centre",
		},
	}
	robo := roborooney.NewRobo(pitches, rules)
	robo.Connect()
}
