package urbandictionary

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
	text := ""
	userName := ""

	if p.Request.Method == "POST" {
		text = p.Request.Form.Get("text")
		userName = p.Request.Form.Get("user_name")
	} else {
		text = p.Request.URL.Query().Get("text")
		userName = p.Request.URL.Query().Get("user_name")
	}

	url := fmt.Sprintf("http://urbanscraper.herokuapp.com/define/%v", url.QueryEscape(text))
	resp, err := http.Get(url)

	body, _ := ioutil.ReadAll(resp.Body)
	var baseData map[string]string
	json.Unmarshal([]byte(body), &baseData)

	if err != nil || resp.StatusCode != 200 {
		return &sscaas.PluginResponse{}, errors.New(baseData["message"])
	}

	if err == nil {
		returnString := fmt.Sprintf(
			"%v: %v - %v\n\n_%v_",
			userName,
			text,
			baseData["definition"],
			baseData["example"],
		)

		return &sscaas.PluginResponse{
			Username: "Urban Dictionary",
			Emoji:    ":poop:",
			Text:     returnString,
			UnfurlLinks: true,
			UnfurlMedia: true,
		}, nil
	}

	return &sscaas.PluginResponse{}, nil
}
