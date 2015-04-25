package main

import (
	// "net/url"
	// "encoding/json"
    "github.com/danielsamuels/sscaas/sscaas"
	"github.com/danielsamuels/sscaas/reddit"
	// "strconv"
	// "strings"
	"time"
	"net/http"
	"fmt"
	// "log"
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
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        redditObj := reddit.Reddit{
            w,
            r,
        }

        plugin := sscaas.Plugin(redditObj)
        fmt.Println("redditObj details are: ", redditObj)
        fmt.Println("plugin is: ", plugin)
    })
    http.ListenAndServe(":8080", nil)
}

/*
func main2() {
    port := 8080

    // Maintain a map of all available 'plugins'. They must all have the same signature.
    plugins := map[string]func(http.ResponseWriter, *http.Request) (*sscaas.PluginResponse, error) {
        "reddit": reddit.Reddit,
    }

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

            // Check if the path matches a plugin. If it does, execute it.
            if plugin, ok := plugins[key]; ok {
                res, err := plugin {w, r}

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
*/
