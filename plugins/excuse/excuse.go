package excuse

import (
    "errors"
    "fmt"
    "net/http"
    "github.com/danielsamuels/sscaas/sscaas"
    "github.com/PuerkitoBio/goquery"
)

type Plugin struct {
    Writer http.ResponseWriter
    Request *http.Request
}

func (p Plugin) Run(http.ResponseWriter, *http.Request) (*sscaas.PluginResponse, error) {
    url := fmt.Sprintf("http://www.programmerexcuses.com/")
    resp, err := goquery.NewDocument(url)

    if err != nil {
        return &sscaas.PluginResponse{}, errors.New(err.Error())
    }

    return &sscaas.PluginResponse{
        Username: "Developer Excuse",
        Emoji: ":no_entry_sign:",
        Text: fmt.Sprintf(
            "%v says: %v",
            p.Request.URL.Query().Get("user_name"),
            resp.Find("center a").Text(),
        ),
    }, nil
}
