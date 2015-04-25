package main

import (
	"net/url"
	"encoding/json"
	"log"
	"strings"
	"strconv"
	"github.com/danielsamuels/sscaas/sscaas"
	"github.com/danielsamuels/sscaas/reddit"
	"time"
	"net/http"
	"fmt"
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
    Channel     string  `json:"channel"`
    Username    string  `json:"username"`
    IconEmoji   string  `json:"icon_emoji"`
    Text        string  `json:"text"`
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
                    plugin = reddit.Reddit{w, r}
            }

            if plugin != nil {
                res, err := sscaas.Plugin(plugin).Run(w, r)

                if err == nil {
                    // Create the JSON payload.
                    responsePayload := &responsePayload{
                        Channel: r.URL.Query().Get("channel_id"),
                        Username: res.Username,
                        IconEmoji: res.Emoji,
                        Text: res.Text,
                    }

                    responsePayloadJSON, _ := json.Marshal(responsePayload)
                    stringJSON := string(responsePayloadJSON[:])

                    // Make the request to the Slack API.
                    http.PostForm(r.URL.Query().Get("callback"), url.Values{"payload": {stringJSON}})
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
