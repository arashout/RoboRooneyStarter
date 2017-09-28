package mlpapi

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	apiEndpoint = "https://api-v2.mylocalpitch.com"
	baseURL     = "https://www.mylocalpitch.com"
)

func New() *MLPClient {
	return &MLPClient{
		httpClient: &http.Client{},
	}
}

func (mlpClient *MLPClient) Close() {
	// Nothing here yet
}

func (mlpClient *MLPClient) GetPitchSlots(pitch Pitch, starts time.Time, ends time.Time) []Slot {
	u, err := url.Parse(apiEndpoint + "/pitches/" + pitch.ID + "/slots")
	if err != nil {
		log.Fatal(err)
	}
	u.Scheme = "https"

	const layout = "2006-01-02"
	q := u.Query()
	q.Set("filter[starts]", starts.Format(layout))
	q.Set("filter[ends]", ends.Format(layout))
	u.RawQuery = q.Encode()

	// Add request headers
	req, err := http.NewRequest("GET", u.String(), nil)
	req.Host = apiEndpoint
	req.Header.Set("Origin", baseURL)
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "en-US,en;q=0.8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.91 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Referer", baseURL+"/"+pitch.City+"/venue/"+pitch.Path)
	req.Header.Set("Connection", "keep-alive")

	response, err := mlpClient.httpClient.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	mlpResponse := MLPResponse{}
	GetJSON(response, &mlpResponse)

	// Return the array of Slots
	return mlpResponse.Data
}

func (mlpClient *MLPClient) FilterSlotsByRules(slots []Slot, rules []Rule) []Slot {
	if len(rules) == 0 {
		return slots
	}

	filteredSlots := make([]Slot, 0)
	for _, slot := range slots {
		if checkAllRulesForSlot(rules, slot) {
			filteredSlots = append(filteredSlots, slot)
		}
	}
	return filteredSlots
}

// GetSlotCheckoutLink return the url for checking out a booking, given a Pitch and Slot
func GetSlotCheckoutLink(pitch Pitch, slot Slot) string {
	return fmt.Sprintf("%s/%s/venue/%s/checkout/%s", baseURL, pitch.City, pitch.Path, slot.ID)
}

func checkAllRulesForSlot(rules []Rule, slot Slot) bool {
	for _, rule := range rules {
		if !rule.DoesSlotPass(slot) {
			return false
		}
	}
	return true
}
