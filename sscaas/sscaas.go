package sscaas

import (
    "net/http"
)

// Plugin format for all plugins.
type Plugin interface {
    Run(http.ResponseWriter, *http.Request) (*PluginResponse, error)
}

// PluginResponse is structure to be returned from plugins.
type PluginResponse struct {
    Username    string
    Emoji       string
    Text        string
}
