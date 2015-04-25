package main

import (
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
