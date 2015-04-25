package sscaas

import (
    "net/http"
)

type Plugin interface {
    Run(http.ResponseWriter, *http.Request) (*PluginResponse, error)
}

type PluginResponse struct {
    Username    string
    Emoji       string
    Text        string
}
