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
	return &sscaas.PluginResponse{
		Username: "Troutslap Bot",
		Emoji:    ":fish:",
		Text: fmt.Sprintf(
			"_%v slaps %v around a bit with a large trout_",
			p.Request.URL.Query().Get("user_name"),
			p.Request.URL.Query().Get("text"),
		),
	}, nil
}
