package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/memory"
)

func main() {
	llm, err := ollama.New(ollama.WithModel("gemma2:2b"))
	if err != nil {
		log.Fatal(err)
	}

	chatMemory := memory.NewConversationBuffer()

	scanner := bufio.NewScanner(os.Stdin)
	ctx := context.Background()

	for {
		fmt.Print("Your turn:")
		if !scanner.Scan() {
			break
		}

		userInput := strings.TrimSpace(scanner.Text())
		if userInput == "exit" || userInput == "quit" {
			fmt.Println("Exiting chat...")
			break
		}

		response, er := llm.GenerateContent(ctx, []llms.MessageContent{llms.TextParts(llms.ChatMessageTypeHuman, userInput)})
		if er != nil {
			fmt.Println("Error generating response:", er)
		}

		aiResponse := response.Choices[0].Content
		fmt.Printf("AI: %s\n\n", aiResponse)

		// Store conversation in memory
		chatMemory.ChatHistory.AddUserMessage(ctx, userInput)
		chatMemory.ChatHistory.AddAIMessage(ctx, aiResponse)

	}

}
