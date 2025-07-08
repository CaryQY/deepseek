package main

import (
	"context"
	"fmt"
	"github.com/CaryQY/deepseek"
	"github.com/CaryQY/deepseek/config"
	"github.com/CaryQY/deepseek/request"
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

	inputMessage := "1 + 1 = ?" // set your input message
	chatReq := &request.ChatCompletionsRequest{
		Model:  chatModel,
		Stream: false,
		Messages: []*request.Message{
			{
				Role:    "user",
				Content: inputMessage,
			},
		},
	}
	fmt.Printf("input message => %s\n", chatReq.Messages[0].Content)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// call deepseek api
	cr, err := cli.CallChatCompletionsChat(ctx, chatReq)
	if err != nil {
		fmt.Println("error => ", err)
		return
	}
	fmt.Printf("output message => %s\n", cr.Choices[0].Message.Content)
}
