package main

import (
	"log"
	"os"

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
			Description: "Only slots after 4pm (exclusive)",
			DoesSlotPass: func(slot mlpapi.Slot) bool {
				return slot.Attributes.Starts.Hour() > 16
			},
		},
		mlpapi.Rule{
			Description: "Only slots before 7pm (exclusive)",
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

	log.Println("Reading config.json for credentials")
	cred := &roborooney.Credentials{
		APIToken:  os.Getenv("API_TOKEN"),
		ChannelID: os.Getenv("CHANNEL_ID"),
		BotID:     os.Getenv("BOT_ID"),
	}

	if cred.BotID == "" {
		log.Println("BotID not set, mentions like @roborooney will not work...")
	}

	robo := roborooney.NewRobo(pitches, rules, cred)
	robo.Connect()
}
