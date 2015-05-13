package main

import (
	"encoding/json"
	"fmt"
	"github.com/danielsamuels/sscaas/sscaas"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/danielsamuels/sscaas/plugins/define"
	"github.com/danielsamuels/sscaas/plugins/dellarism"
	"github.com/danielsamuels/sscaas/plugins/excuse"
	"github.com/danielsamuels/sscaas/plugins/reddit"
	"github.com/danielsamuels/sscaas/plugins/soundcloud"
	"github.com/danielsamuels/sscaas/plugins/uptime"
	"github.com/danielsamuels/sscaas/plugins/urbandictionary"
)

func logRequest(w http.ResponseWriter, r *http.Request, contentLength string, statusCode int) {
	fmt.Printf(
		"%v - %v - [%v] \"%v %v %v\" %v -\n",
		r.RemoteAddr,
		contentLength,
		time.Now().Format("2/Jan/2006 15:04:05"),
		r.Method,
		r.URL.String(),
		r.Proto,
		statusCode,
	)
}

type responsePayload struct {
	Channel     string `json:"channel"`
	Username    string `json:"username"`
	IconEmoji   string `json:"icon_emoji"`
	Text        string `json:"text"`
	UnfurlMedia bool   `json:"unfurl_media"`
	UnfurlLinks bool   `json:"unfurl_links"`
}

func main() {
	port := 8080

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		statusCode := 200
		contentLength := ""

		if len(r.URL.Query().Get("channel_id")) == 0 {
			errorText := "No channel supplied."
			http.Error(w, errorText, http.StatusBadRequest)
			statusCode = 400
			contentLength = strconv.Itoa(len(errorText))
		} else if len(r.URL.Query().Get("callback")) == 0 {
			errorText := "No callback supplied."
			http.Error(w, errorText, http.StatusBadRequest)
			statusCode = 400
			contentLength = strconv.Itoa(len(errorText))
		} else {
			parts := strings.Split(r.URL.Path[1:], "/")
			key := parts[0]

			var plugin sscaas.Plugin

			// Plugin definitions..
			// TODO: Make some sort of plugin registration system?
			switch key {
			case "reddit":
				plugin = reddit.Plugin{w, r}
			case "urbandictionary":
				plugin = urbandictionary.Plugin{w, r}
			case "dellarism":
				plugin = dellarism.Plugin{w, r}
			case "define":
				plugin = define.Plugin{w, r}
			case "excuse":
				plugin = excuse.Plugin{w, r}
			case "soundcloud":
				plugin = soundcloud.Plugin{w, r}
			case "uptime":
				plugin = uptime.Plugin{w, r}
			}

			if plugin != nil {
				res, err := sscaas.Plugin(plugin).Run(w, r)

				if err == nil {
					// Create the JSON payload.
					responsePayload := &responsePayload{
						Channel:     r.URL.Query().Get("channel_id"),
						Username:    res.Username,
						IconEmoji:   res.Emoji,
						Text:        res.Text,
						UnfurlMedia: true,
						UnfurlLinks: true,
					}

					responsePayloadJSON, _ := json.Marshal(responsePayload)
					stringJSON := string(responsePayloadJSON[:])

					// Make the request to the Slack API.
					http.PostForm(r.URL.Query().Get("callback"), url.Values{"payload": {stringJSON}})
				} else {
					http.Error(w, err.Error(), 200)
				}
			} else {
				errorText := "Sorry, it was not possible to load that plugin."
				http.Error(w, errorText, http.StatusNotFound)
				statusCode = 404
				contentLength = strconv.Itoa(len(errorText))
			}
		}

		logRequest(w, r, contentLength, statusCode)
	})

	fmt.Printf("Running server on port %d..\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
