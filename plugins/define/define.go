package define

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/danielsamuels/sscaas/sscaas"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Plugin struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

func (p Plugin) Run(http.ResponseWriter, *http.Request) (*sscaas.PluginResponse, error) {
	url := fmt.Sprintf(
		"http://api.wordnik.com/v4/word.json/%v/definitions?limit=1&includeRelated=false&useCanonical=false&includeTags=true&api_key=f8ab3913c02c28a5b8a4c086d3b072d3b92e4551e13f52d0f",
		url.QueryEscape(p.Request.URL.Query().Get("text")),
	)
	resp, err := http.Get(url)

	body, _ := ioutil.ReadAll(resp.Body)

	stringJSON := string(body[:])
	stringJSON = stringJSON[1:]
	stringJSON = stringJSON[:len(stringJSON)-1]

	var baseData map[string]string
	json.Unmarshal([]byte(stringJSON), &baseData)

	if stringJSON == "" {
		return &sscaas.PluginResponse{}, errors.New("Word not found.")
	}

	if err == nil {
		returnString := fmt.Sprintf(
			"%v: %v - %v",
			p.Request.URL.Query().Get("user_name"),
			baseData["word"],
			baseData["text"],
		)

		return &sscaas.PluginResponse{
			Username: "Dictionary Bot",
			Emoji:    ":book:",
			Text:     returnString,
		}, nil
	}

	return &sscaas.PluginResponse{
		Username: "Dictionary Bot",
		Emoji:    ":book:",
		Text:     "",
	}, nil
}
