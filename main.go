package main

import (
	"github.com/arashout/mlpapi"
	"github.com/arashout/roborooney"
)

func main() {
	rules := []mlpapi.Rule{
		mlpapi.Rule{
			Description: "Only available slots",
			DoesSlotPass: func(slot mlpapi.Slot) bool {
				return slot.Attributes.Availabilities > 0
			},
		},
		mlpapi.Rule{
			Description: "Only slots after 4pm",
			DoesSlotPass: func(slot mlpapi.Slot) bool {
				return slot.Attributes.Starts.Hour() > 16
			},
		},
		mlpapi.Rule{
			Description: "Only slots before 7pm",
			DoesSlotPass: func(slot mlpapi.Slot) bool {
				return slot.Attributes.Starts.Hour() < 19
			},
		},
		mlpapi.Rule{
			Description: "Only slots not on a weekend",
			DoesSlotPass: func(slot mlpapi.Slot) bool {
				// Week starts at Sunday == index 0
				return slot.Attributes.Starts.Weekday() != 6 && slot.Attributes.Starts.Weekday() != 0
			},
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
