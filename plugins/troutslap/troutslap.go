package troutslap

import (
	"fmt"
	"github.com/danielsamuels/sscaas/sscaas"
	"net/http"
)

type Plugin struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

func (p Plugin) Run(http.ResponseWriter, *http.Request) (*sscaas.PluginResponse, error) {
	userName := ""
	text := ""

	if p.Request.Method == "POST" {
		userName = p.Request.Form.Get("user_name")
		text = p.Request.Form.Get("text")
	} else {
		userName = p.Request.URL.Query().Get("user_name")
		text = p.Request.URL.Query().Get("text")
	}

	return &sscaas.PluginResponse{
		Username: "Troutslap",
		Emoji:    ":fish:",
		Text: fmt.Sprintf(
			"_%v slaps %v around a bit with a large trout_",
			userName,
			text,
		),
		UnfurlLinks: true,
		UnfurlMedia: true,
	}, nil
}
