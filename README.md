# Slack Slash Commands as a Service
## Making Slack better since some time early 2015.

https://sscaas.eu/

## Example package

``` go
package example

import (
    "github.com/danielsamuels/sscaas/sscaas"
)

type Plugin struct {
    Writer http.ResponseWriter
    Request *http.Request
}

func (p Plugin) Run(http.ResponseWriter, *http.Request) (*sscaas.PluginResponse, error) {
    return &sscaas.PluginResponse{
        Username: "Example Bot",
        Emoji: ":smile:",
        Text: "Hello world.",
    }, nil
}
```
