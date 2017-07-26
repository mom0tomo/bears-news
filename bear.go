package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// IncomingURL is Slack incoming Url.
var IncomingURL = ""

// Slack is Message Struct.
type Slack struct {
	Text    string `json:"text"`
	Channel string `json:"channel"`
}

func main() {
	params, _ := json.Marshal(Slack{
		"ʕ•ᴥ•ʔ＜Hello",
		"#bears-news",
	})

	resp, _ := http.PostForm(
		IncomingURL,
		url.Values{"payload": {string(params)}},
	)

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Println(string(body))
}
