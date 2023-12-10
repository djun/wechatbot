package gtp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"wechatbot/config"
)

type Role int

const (
	User      Role = iota // User = 0
	System                // System = 1
	Assistant             // Assistant = 2
)

func (r Role) String() string {
	return []string{"user", "system", "assistant"}[r]
}

const BASEURL = "https://api.openai.com/v1/chat/"

// ChatGPTResponseBody 响应体
type ChatGPTResponseBody struct {
	ID      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created int                    `json:"created"`
	Model   string                 `json:"model"`
	Choices []ChoiceItem           `json:"choices"`
	Usage   map[string]interface{} `json:"usage"`
}

type ChoiceItem struct {
	Index        int         `json:"index"`
	Message      MessageItem `json:"message"`
	FinishReason string      `json:"finish_reason"`
}

type MessageItem struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatGPTRequestBody 请求体
type ChatGPTRequestBody struct {
	Model    string        `json:"model"`
	Messages []MessageItem `json:"messages"`
}

// Completions gtp文本模型回复
// curl https://api.openai.com/v1/chat/completions
// -H "Content-Type: application/json"
// -H "Authorization: Bearer your chatGPT key"
// -d '{"model": "gpt-3.5-turbo", "messages: [{"role": "user", "content": "your problem"}]"}'
func Completions(msg string) (string, error) {
	gptModel := config.LoadConfig().GptModel
	requestBody := ChatGPTRequestBody{
		Model: gptModel,
		Messages: []MessageItem{
			{
				Role:    Role(0).String(),
				Content: msg,
			},
		},
	}
	requestData, err := json.Marshal(requestBody)

	if err != nil {
		return "", err
	}
	log.Printf("request gtp json string : %v", string(requestData))
	req, err := http.NewRequest("POST", BASEURL+"completions", bytes.NewBuffer(requestData))
	if err != nil {
		return "", err
	}

	apiKey := config.LoadConfig().ApiKey
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	gptResponseBody := &ChatGPTResponseBody{}
	log.Println(string(body))
	err = json.Unmarshal(body, gptResponseBody)
	if err != nil {
		return "", err
	}
	reply := "暂无回复"
	choicesLength := len(gptResponseBody.Choices)
	if choicesLength > 0 {
		reply = gptResponseBody.Choices[choicesLength-1].Message.Content
	}
	log.Printf("gpt response text: %s \n", reply)
	return reply, nil
}
