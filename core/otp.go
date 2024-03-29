package core

import (
	"net/http"
	"net/url"
	"strings"
)

func GenerateRandomNumber() string {
	// rand.Seed(time.Now().UnixNano())
	// d1 := rand.Intn(10)
	// d2 := rand.Intn(10)
	// d3 := rand.Intn(10)
	// d4 := rand.Intn(10)
	// numStr := fmt.Sprintf("%v%v%v%v", d1, d2, d3, d4)
	numStr := "1234"
	return numStr
}

func SendOTP(num string) bool {
	sID := ""
	authToken := ""
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + sID + "/Messages.json"

	msgData := url.Values{}
	msgData.Set("To", num)
	msgData.Set("From", "+18456689486")
	msgData.Set("Body", GenerateRandomNumber())
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(sID, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	_, err := client.Do(req)
	if err != nil {
		return false
	}
	return true
}
