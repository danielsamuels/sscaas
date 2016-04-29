package excuse

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/danielsamuels/sscaas/sscaas"
	"net/http"
)

type Plugin struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

func (p Plugin) Run(http.ResponseWriter, *http.Request) (*sscaas.PluginResponse, error) {
	userName := ""

	if p.Request.Method == "POST" {
		userName = p.Request.Form.Get("user_name")
	} else {
		userName = p.Request.URL.Query().Get("user_name")
	}

	url := fmt.Sprintf("http://www.programmerexcuses.com/")
	resp, err := goquery.NewDocument(url)

	if err != nil {
		return &sscaas.PluginResponse{}, errors.New(err.Error())
	}

	return &sscaas.PluginResponse{
		Username: "Developer Excuse",
		Emoji:    ":no_entry_sign:",
		Text: fmt.Sprintf(
			"%v says: %v",
			userName,
			resp.Find("center a").Text(),
		),
		UnfurlLinks: true,
		UnfurlMedia: true,
	}, nil
}
