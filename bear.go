package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

// Slack is Message Struct.
type Slack struct {
	Text    string `json:"text"`
	Channel string `json:"channel"`
}

func main() {
	// IncomingURL is Slack incoming webhook URL.
	IncomingURL := os.Getenv("WEBHOOK_URL")

	params, err := json.Marshal(Slack{
		"ʕ ◔ϖ◔ʔ < Hello",
		os.Getenv("SLACK_CHANNEL"),
	})
	if err != nil {
		fmt.Println(err)
	}

	resp, err := http.PostForm(
		IncomingURL,
		url.Values{"payload": {string(params)}},
	)
	if err != nil {
		log.Fatal("HTTPリクエスト時にエラーが発生しました")
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal("リクエストボディの読み込み時にエラーが発生しました")
	}

	fmt.Println(string(body))
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("環境変数を読み込めませんでした")
	}
}
