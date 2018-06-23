package api

import "bytes"
import "encoding/json"
import "io/ioutil"
import "net/http"

type QuoteRequest struct {
	Action   string   `json:"action"`
	Projects []string `json:"projects"`
}

type QuoteResponse struct {
	Projects []QuoteProject `json:"projects"`
}

type QuoteProject struct {
	Licensor    LicensorInformation `json:"licensor"`
	ProjectID   string              `json:"projectID"`
	Description string              `json:"description"`
	Repository  string              `json:"homepage"`
	Pricing     Pricing             `json:"pricing"`
	Retracted   bool                `json:"retracted"`
}

type LicensorInformation struct {
	Name         string
	Jurisdiction string
	PublicKey    string
}

type Pricing struct {
	Private int
}

func Quote(projectIDs []string) (QuoteResponse, error) {
	bodyData := QuoteRequest{
		Action:   "quote",
		Projects: projectIDs,
	}
	body, err := json.Marshal(bodyData)
	if err != nil {
		return QuoteResponse{}, err
	}
	response, err := http.Post("https://licensezero.com/api/v0", "application/json", bytes.NewBuffer(body))
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return QuoteResponse{}, err
	}
	var parsed QuoteResponse
	err = json.Unmarshal(responseBody, &parsed)
	if err != nil {
		return QuoteResponse{}, err
	}
	return parsed, nil
}