package urbandictionary

import (
    "errors"
    "encoding/json"
    "io/ioutil"
    "net/url"
    "fmt"
    "github.com/danielsamuels/sscaas/sscaas"
    "net/http"
)

type Plugin struct {
    Writer http.ResponseWriter
    Request *http.Request
}

func (p Plugin) Run(http.ResponseWriter, *http.Request) (*sscaas.PluginResponse, error) {
    url := fmt.Sprintf("http://urbanscraper.herokuapp.com/define/%v", url.QueryEscape(p.Request.URL.Query().Get("text")))
    resp, err := http.Get(url)

    body, _ := ioutil.ReadAll(resp.Body)
    var baseData map[string]string
    json.Unmarshal([]byte(body), &baseData)

    if err != nil || resp.StatusCode != 200 {
        return &sscaas.PluginResponse{}, errors.New(baseData["message"])
    }

    if err == nil {
        returnString := fmt.Sprintf(
            "%v: %v - %v",
            p.Request.URL.Query().Get("user_name"),
            p.Request.URL.Query().Get("text"),
            baseData["definition"],
        )

        return &sscaas.PluginResponse{
            Username: "Urban Dictionary Bot",
            Emoji: ":poop:",
            Text: returnString,
        }, nil
    }

    return &sscaas.PluginResponse{}, nil
}
