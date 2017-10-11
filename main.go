package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/arashout/mlpapi"
	"github.com/arashout/roborooney"
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

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
		mlpapi.Pitch{
			ID:   "32180",
			Path: "calthorpe-project-sports-facilities/football-5-a-side-32180",
			City: "london",
			Name: "Calthorpe Project Sports Facilities",
		},
	}

	log.Println("Reading credentials from enviroment")
	cred := &roborooney.Credentials{
		APIToken: os.Getenv("API_TOKEN"),
		BotID:    os.Getenv("BOT_ID"),
	}

	if cred.BotID == "" {
		log.Println("BotID not set, mentions like @roborooney will not work...")
	}
	if cred.APIToken == "" {
		log.Fatal("Credentials not set! Ending program.")
	}

	robo := roborooney.NewRobo(pitches, rules, cred)
	robo.Connect()
}
