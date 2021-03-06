package api

import "bytes"
import "encoding/json"
import "errors"
import "licensezero.com/cli/data"
import "io/ioutil"
import "net/http"
import "strconv"

type retractRequest struct {
	Action      string `json:"action"`
	DeveloperID string `json:"developerID"`
	Token       string `json:"token"`
	OfferID     string `json:"offerID"`
}

type retractResponse struct {
	Error interface{} `json:"error"`
}

// Retract sends retract API requests.
func Retract(developer *data.Developer, offerID string) error {
	bodyData := retractRequest{
		Action:      "retract",
		DeveloperID: developer.DeveloperID,
		Token:       developer.Token,
		OfferID:     offerID,
	}
	body, err := json.Marshal(bodyData)
	if err != nil {
		return err
	}
	response, err := http.Post("https://licensezero.com/api/v0", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return errors.New("error sending request")
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return errors.New("Server responded " + strconv.Itoa(response.StatusCode))
	}
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	var parsed retractResponse
	err = json.Unmarshal(responseBody, &parsed)
	if err != nil {
		return err
	}
	if message, ok := parsed.Error.(string); ok {
		return errors.New(message)
	}
	return nil
}
