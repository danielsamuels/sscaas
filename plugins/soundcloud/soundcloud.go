package soundcloud

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

type soundcloudTrack struct {
	Permalink_url string
	Title         string
}

func (p Plugin) Run(http.ResponseWriter, *http.Request) (*sscaas.PluginResponse, error) {
	text := ""

	if p.Request.Method == "POST" {
		text = p.Request.Form.Get("text")
	} else {
		text = p.Request.URL.Query().Get("text")
	}

	clientID := "a45552608d49ea7babda4dde4a1e82e4"
	url := fmt.Sprintf("https://api.soundcloud.com/tracks/?q=%v&client_id=%v", url.QueryEscape(text), clientID)
	resp, err := http.Get(url)

	if err != nil || resp.StatusCode != 200 {
		return &sscaas.PluginResponse{}, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return &sscaas.PluginResponse{}, errors.New(resp.Status)
	}

	// var baseData map[string]interface{}
	baseData := make([]soundcloudTrack, 0)
	json.Unmarshal(body, &baseData)

	if len(baseData) == 0 {
		return &sscaas.PluginResponse{}, errors.New(fmt.Sprintf("Unable to find any results for %v.", text))
	}

	return &sscaas.PluginResponse{
		Username: "SoundCloud",
		Emoji:    ":speaker:",
		Text:     fmt.Sprintf("<%v|%v>", baseData[0].Permalink_url, baseData[0].Title),
		UnfurlLinks: true,
		UnfurlMedia: true,
	}, nil
}
