package reddit

import (
    "github.com/danielsamuels/sscaas/sscaas"
    "errors"
    "encoding/json"
    "io/ioutil"
    "fmt"
    "net/http"
)

type Reddit struct {
    Writer http.ResponseWriter
    Request *http.Request
}

func (r Reddit) Run(http.ResponseWriter, *http.Request) (*sscaas.PluginResponse, error) {
    subreddit := r.Request.URL.Query().Get("text")

    if len(subreddit) == 0 {
        http.Error(r.Writer, "No subreddit supplied.", http.StatusBadRequest)
        return &sscaas.PluginResponse{}, errors.New("No subreddit supplied.")
    }

    url := fmt.Sprintf("http://www.reddit.com/r/%s/about.json", r.Request.URL.Query().Get("text"))

    client := &http.Client{}

    req, err := http.NewRequest("GET", url, nil)
    req.Header.Add("User-Agent", "Slack slash command")
    resp, err := client.Do(req)

    if err != nil || resp.StatusCode != 200 {
        if resp.StatusCode == 404 {
            http.Error(r.Writer, "That subreddit does not exist.", http.StatusNotFound)
            return &sscaas.PluginResponse{}, errors.New("That subreddit does not exist.")
        } else {
            return &sscaas.PluginResponse{}, errors.New("There was an error with the request.")
        }
    }

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        fmt.Println("There was an error parsing the response.")
        return &sscaas.PluginResponse{}, err
    }

    var base_data map[string]interface{}
    json.Unmarshal([]byte(body), &base_data)
    data_key := base_data["data"]

    data := data_key.(map[string]interface{})
    nsfw := ""

    if data["over18"] == true {
        nsfw = "(NSFW)"
    }

    returnString := fmt.Sprintf(
        "%v - %v (%v): http://www.reddit.com/r/%v %v",
        r.Request.URL.Query().Get("user_name"),
        data["display_name"],
        data["title"],
        subreddit,
        nsfw,
    )

    return &sscaas.PluginResponse{
        Username: "Reddit Bot",
        Emoji: ":reddit:",
        Text: returnString,
    }, nil
}
