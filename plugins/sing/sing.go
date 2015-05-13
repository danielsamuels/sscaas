package sing

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
	text := ""

	if p.Request.Method == "POST" {
		text = p.Request.Form.Get("text")
	} else {
		text = p.Request.URL.Query().Get("text")
	}

	return &sscaas.PluginResponse{
		Username: "Sing Bot",
		Emoji:    ":musical_score:",
		Text: fmt.Sprintf(
			":musical_note: %v :musical_note:",
			text,
		),
	}, nil
}
