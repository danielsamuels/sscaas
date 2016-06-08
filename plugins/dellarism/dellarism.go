package dellarism

import (
	"fmt"
	"github.com/danielsamuels/sscaas/sscaas"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Plugin struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

func shuffle(arr []string) {
	rand.Seed(int64(time.Now().Nanosecond()))

	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func random(min, max int) int {
	rand.Seed(int64(time.Now().Nanosecond()))
	return rand.Intn(max-min) + min
}

func (p Plugin) Run(http.ResponseWriter, *http.Request) (*sscaas.PluginResponse, error) {
	userName := ""

	if p.Request.Method == "POST" {
		userName = p.Request.Form.Get("user_name")
	} else {
		userName = p.Request.URL.Query().Get("user_name")
	}

	generatedString := ""

	firstWord := []string{"Change", "Make", "Build", "Need", "Put"}
	secondWord := []string{"black", "social", "keyline", "content", "copy", "red", "scrolling", "footer", "header", "button", "navigation", "grid", "font", "spacing", "typekit", "website"}
	thirdWord := []string{"to", "from", "into", "into the", "under the", "on top of"}
	fourthWord := []string{"blue", "carousel", "section", "map", "paragraph", "ajax", "footer", "responsive", "packery", "python", "social"}

	shuffle(firstWord)
	shuffle(secondWord)
	shuffle(thirdWord)
	shuffle(fourthWord)

	generatedString = fmt.Sprintf("%v %v %v %v", firstWord[0], secondWord[0], thirdWord[0], fourthWord[0])

	if random(1, 5) == 4 {
		generatedString = strings.ToUpper(generatedString)
	}

	return &sscaas.PluginResponse{
		Username: "Dellarism",
		Emoji:    ":dellar:",
		Text: fmt.Sprintf(
			"%v, your Dellarism is: %v",
			userName,
			generatedString,
		),
		UnfurlLinks: true,
		UnfurlMedia: true,
	}, nil
}
