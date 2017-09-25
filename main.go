package main

import (
	"github.com/arashout/mlpapi"
	"github.com/arashout/roborooney"
)

//go:generate echo '{"apiToken":"","channelId":""}' > config.json

// Rule functions
// Function signature should be func(mlpapi.Slot) bool

func filterAvailable(slot mlpapi.Slot) bool {
	return slot.Attributes.Availabilities > 0
}
func filterAfterTime(slot mlpapi.Slot) bool {
	// Only show slots after 4pm
	return slot.Attributes.Starts.Hour() > 16
}
func filterBeforeTime(slot mlpapi.Slot) bool {
	// Only show slots before 8pm
	return slot.Attributes.Starts.Hour() < 20
}
func filterExcludeWeekends(slot mlpapi.Slot) bool {
	// Week starts at Sunday == index 0
	return slot.Attributes.Starts.Weekday() != 6 && slot.Attributes.Starts.Weekday() != 0
}

func main() {
	rules := make([]func(mlpapi.Slot) bool, 0)
	// Return only available slots
	rules = append(rules, filterAvailable)
	rules = append(rules, filterAfterTime)
	rules = append(rules, filterExcludeWeekends)
	pitches := []mlpapi.Pitch{
		mlpapi.Pitch{
			ID:   "34933",
			Path: "three-corners-adventure-playground/football-5-a-side-34933",
			City: "london",
		},
		mlpapi.Pitch{
			ID:   "32208",
			Path: "finsbury-leisure-centre/football-5-a-side-32208",
			City: "london",
		},
	}
	robo := roborooney.NewRobo(pitches, rules)
	robo.Connect()
}
