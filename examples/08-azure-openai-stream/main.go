package main

import (
	"context"
	"fmt"
	"github.com/CaryQY/deepseek"
	"github.com/CaryQY/deepseek/config"
	"github.com/CaryQY/deepseek/request"
	"io"
	"os"
)

func main() {
	cli, _ := deepseek.NewClientWithConfig(config.Config{
		FullURL:                  os.Getenv("AZUREFULLURL"),
		ApiKey:                   os.Getenv("AZUREKEY"),
		TimeoutSeconds:           120,
		DisableRequestValidation: false,
	})

	chatModel := os.Getenv("AZUREMODEL")

	inputMessage := "rust vs golang" // set your input message
	chatReq := &request.ChatCompletionsRequest{
		Model:  chatModel,
		Stream: true,
		Messages: []*request.Message{
			{
				Role:    "user",
				Content: inputMessage,
			},
		},
		StreamOptions: &request.StreamOptions{IncludeUsage: true},
	}
	fmt.Printf("input message => %s\n", chatReq.Messages[0].Content)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// call deepseek api
	sr, err := cli.StreamChatCompletionsChat(ctx, chatReq)
	if err != nil {
		fmt.Println("error => ", err)
		return
	}

	fmt.Print("output message => ")
	for {
		chatResp, err := sr.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		if len(chatResp.Choices) > 0 {
			fmt.Print(chatResp.Choices[0].Delta.Content)
		} else if chatResp.Usage != nil {
			fmt.Printf("\nPromptTokens: %d, CompletionTokens: %d, TotalTokens: %d \n",
				chatResp.Usage.PromptTokens,
				chatResp.Usage.CompletionTokens,
				chatResp.Usage.TotalTokens)
		}
	}
}
