package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/memory"
)

func main() {
	llm, err := ollama.New(ollama.WithModel("gemma2:2b"))
	if err != nil {
		log.Fatal(err)
	}

	chatMemory := memory.NewConversationBuffer()

	chain := chains.NewConversation(llm, chatMemory)

	scanner := bufio.NewScanner(os.Stdin)
	ctx := context.Background()

	for {
		fmt.Print("Your turn: ")
		if !scanner.Scan() {
			break
		}

		userInput := strings.TrimSpace(scanner.Text())
		if userInput == "exit" || userInput == "quit" {
			fmt.Println("Exiting chat...")
			break
		}

		response, er := chains.Run(ctx, chain, userInput, chains.WithTemperature(0.8))
		if er != nil {
			fmt.Println("Error generating response:", er)
		}

		fmt.Printf("AI: %s\n\n", response)

	}

}
