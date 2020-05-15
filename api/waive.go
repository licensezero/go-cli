package api

import "bytes"
import "encoding/json"
import "errors"
import "licensezero.com/cli/data"
import "io/ioutil"
import "net/http"
import "strconv"

type waiveRequest struct {
	Action       string      `json:"action"`
	DeveloperID  string      `json:"developerID"`
	Token        string      `json:"token"`
	OfferID      string      `json:"offerID"`
	Beneficiary  string      `json:"beneficiary"`
	Jurisdiction string      `json:"jurisdiction"`
	Term         interface{} `json:"term"`
}

type waiveResponse struct {
	Error interface{} `json:"error"`
}

// Waive sends waiver API requests.
func Waive(developer *data.Developer, offerID, beneficiary, jurisdiction string, term interface{}) ([]byte, error) {
	bodyData := waiveRequest{
		Action:       "waiver",
		DeveloperID:  developer.DeveloperID,
		Token:        developer.Token,
		OfferID:      offerID,
		Beneficiary:  beneficiary,
		Jurisdiction: jurisdiction,
		Term:         term,
	}
	body, err := json.Marshal(bodyData)
	if err != nil {
		return nil, errors.New("error serializing request body")
	}
	response, err := http.Post("https://licensezero.com/api/v0", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.New("error sending request")
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return nil, errors.New("Server responded " + strconv.Itoa(response.StatusCode))
	}
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var parsed struct {
		Error interface{} `json:"error"`
	}
	err = json.Unmarshal(responseBody, &parsed)
	if err != nil {
		return nil, err
	}
	if message, ok := parsed.Error.(string); ok {
		return nil, errors.New(message)
	}
	return responseBody, nil
}
