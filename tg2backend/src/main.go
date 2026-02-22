package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func sendResponse(url_ string, id float64, text string) {
	chatID := int64(id)
	json := `{"chat_id" : ` + fmt.Sprintf("%v", chatID) + `, "text" : "` + text + `"}`
	request, _ := http.NewRequest(http.MethodPost, url_+"/sendMessage", bytes.NewBuffer([]byte(json)))
	request.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("Something going wrong\n")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("Telegram API error:", resp.Status, string(body))
		return
	}
}

func getChatId(byt map[string]interface{}) float64 {
	temp := (byt["result"].([]interface{}))
	return ((temp[len(temp)-1].(map[string]interface{}))["message"].(map[string]interface{}))["chat"].(map[string]interface{})["id"].(float64)
}

func getUserText(byt map[string]interface{}) string {
	temp := (byt["result"].([]interface{}))
	return ((temp[len(temp)-1].(map[string]interface{}))["message"].(map[string]interface{}))["text"].(string)
}

func handler() {
	token := os.Getenv("API_TOKEN")
	url_ := "https://api.telegram.org/bot" + token

	request := &http.Request{
		Method: http.MethodGet,
		Header: http.Header{
			"content-type": {"application/json"},
		},
	}
	defer fmt.Println("pizda")
	request.URL, _ = url.Parse(url_ + "/getUpdates")
	i := 0
	for {
		fmt.Printf("=============%v=============\n", i)
		response, err := http.DefaultClient.Do(request)
		if err != nil {
			fmt.Printf("Somtehing going wrong\n")
		}

		var dat map[string]interface{}
		respBody, _ := io.ReadAll(response.Body)
		if err := json.Unmarshal(respBody, &dat); err != nil {
			panic(err)
		}

		text := getUserText(dat)
		id := getChatId(dat)
		fmt.Printf("text: %v\n", text)
		fmt.Printf("chat id: %v\n", id)
		sendResponse(url_, id, text)
		i++

		time.Sleep(10 * time.Second)
	}
}

func main() {
	godotenv.Load(".env")
	fmt.Printf("Someting working\n")
	handler()
}
