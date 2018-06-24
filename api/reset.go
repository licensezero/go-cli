package api

import "bytes"
import "encoding/json"
import "errors"
import "github.com/licensezero/cli/data"
import "net/http"

type ResetRequest struct {
	Action     string `json:"action"`
	LicensorID string `json:"licensorID"`
	Name       string `json:"name"`
	EMail      string `json:"email"`
}

func Reset(identity *data.Identity, licensor *data.Licensor) error {
	bodyData := ResetRequest{
		Action:     "reset",
		LicensorID: licensor.LicensorID,
		Name:       identity.Name,
		EMail:      identity.EMail,
	}
	body, err := json.Marshal(bodyData)
	if err != nil {
		return errors.New("could not construct reset request")
	}
	response, err := http.Post("https://licensezero.com/api/v0", "application/json", bytes.NewBuffer(body))
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return errors.New("Server responded " + string(response.StatusCode))
	}
	return nil
}