package nsfw

import (
	"net/http"
	"github.com/danielsamuels/sscaas/sscaas"
	"fmt"
	"errors"
)

type Plugin struct {
	Writer http.ResponseWriter
	Request *http.Request
}

func (p Plugin) Run(http.ResponseWriter, *http.Request) (*sscaas.PluginResponse, error) {
	text := ""
	username := ""

	if p.Request.Method == "POST" {
		text = p.Request.Form.Get("text")
		username = p.Request.Form.Get("user_name")
	} else {
		text = p.Request.URL.Query().Get("text")
		username = p.Request.URL.Query().Get("user_name")
	}

	if len(text) == 0 {
		http.Error(p.Writer, "No NSFW text supplied.", http.StatusBadRequest)
		return &sscaas.PluginResponse{}, errors.New("No NSFW text supplied.")
	}

	return &sscaas.PluginResponse{
		Username: "NSFW Bot",
		Emoji: ":nsfw:",
		Text: fmt.Sprintf("NSFW content from %s - %s", username, text),
		UnfurlLinks: false,
		UnfurlMedia: false,
	}, nil

}