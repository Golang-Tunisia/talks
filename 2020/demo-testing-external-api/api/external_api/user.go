package external_api

import (
	"Meetup/vars/api_vars"
	"encoding/json"
	"errors"
	"github.com/flannel-dev-lab/cyclops/requester"
	"net/http"
)

type ExternalApi struct {
	Url string
}

func New(url string) *ExternalApi {
	return &ExternalApi{Url: url}
}

func (ea *ExternalApi) CreateUser(user api_vars.User) (id string, err error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	response, err := requester.Post(ea.Url+"/user", headers, nil, user)
	if err != nil {
		return id, err
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		var userId map[string]string

		if err = json.NewDecoder(response.Body).Decode(&userId); err != nil {
			return id, errors.New("json decode err on success: " + err.Error())
		}

		return userId["id"], nil
	}

	var errorResponse map[string]string

	if err = json.NewDecoder(response.Body).Decode(&errorResponse); err != nil {
		return id, errors.New("json decode err on failure: " + err.Error())
	}

	return id, errors.New(errorResponse["error"])
}
