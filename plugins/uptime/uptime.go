/*
r = requests.get('http://www.downforeveryoneorjustme.com/{}'.format(request.args['text']))

if "It\'s just you." in r.text:
    return "{} is up.".format(request.args['text'])
else:
    return "{} is down.".format(request.args['text'])
*/
package uptime

import (
	"errors"
	"fmt"
	"github.com/danielsamuels/sscaas/sscaas"
	"io/ioutil"
	"net/http"
	"strings"
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

	url := fmt.Sprintf("http://www.downforeveryoneorjustme.com/%v", text)
	resp, err := http.Get(url)

	if err != nil {
		return &sscaas.PluginResponse{}, errors.New(err.Error())
	}

	body, _ := ioutil.ReadAll(resp.Body)
	stringBody := string(body[:])

	if strings.Contains(stringBody, "It's just you") {
		return &sscaas.PluginResponse{}, errors.New(fmt.Sprintf("%v is up", text))
	}
	return &sscaas.PluginResponse{}, errors.New(fmt.Sprintf("%v is down", text))
}
