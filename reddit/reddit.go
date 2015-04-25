package reddit

import (
	"errors"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"net/http"
)


func Run(w http.ResponseWriter, r *http.Request) (string, string, string, error) {
    subreddit := r.URL.Query().Get("text")

    if len(subreddit) == 0 {
        http.Error(w, "No subreddit supplied.", http.StatusBadRequest)
        return "", "", "", errors.New("No subreddit supplied.")
    }

    url := fmt.Sprintf("http://www.reddit.com/r/%s/about.json", r.URL.Query().Get("text"))

    client := &http.Client{}

    req, err := http.NewRequest("GET", url, nil)
    req.Header.Add("User-Agent", "Slack slash command")
    resp, err := client.Do(req)

    if err != nil || resp.StatusCode != 200 {
        if resp.StatusCode == 404 {
            http.Error(w, "That subreddit does not exist.", http.StatusNotFound)
            return "", "", "", errors.New("That subreddit does not exist.")
        } else {
            return "", "", "", errors.New("There was an error with the request.")
        }
    }

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        fmt.Println("There was an error parsing the response.")
        return "", "", "", err
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
        r.URL.Query().Get("user_name"),
        data["display_name"],
        data["title"],
        subreddit,
        nsfw,
    )

    return "Reddit Bot", ":reddit:", returnString, nil
}
