package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/mtfuller/tiny-ai-poc/internal/color"
	"github.com/mtfuller/tiny-ai-poc/internal/logger"
	"github.com/mtfuller/tiny-ai-poc/pkg/ollama"
	"github.com/spf13/cobra"
)

var (
	chatModel  string
	ollamaURL  string
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Start an interactive chat session with a local model",
	Long: `Start an interactive chat session powered by a small local model via Ollama.

The chat command manages conversation history and uses structured prompting to
get reliable responses from compact models.`,
	Example: `  tiny-ai chat
  tiny-ai chat --model llama3.2:1b
  tiny-ai chat --model llama3.2:1b --ollama-url http://localhost:11434`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := ollama.NewClient(ollamaURL)

		color.Info("Model: %s  (Ollama: %s)", chatModel, ollamaURL)
		color.Info("Type 'exit' or press Ctrl+C to quit.\n")

		var history []ollama.Message
		scanner := bufio.NewScanner(os.Stdin)

		for {
			fmt.Print(color.Bold("You: "))
			if !scanner.Scan() {
				break
			}
			input := strings.TrimSpace(scanner.Text())
			if input == "" {
				continue
			}
			if input == "exit" || input == "quit" {
				break
			}

			history = append(history, ollama.Message{Role: "user", Content: input})

			logger.Debug("Sending %d messages to model %s", len(history), chatModel)

			resp, err := client.Chat(context.Background(), ollama.ChatRequest{
				Model:    chatModel,
				Messages: history,
			})
			if err != nil {
				return fmt.Errorf("chat request failed: %w", err)
			}

			history = append(history, resp.Message)
			fmt.Printf("%s %s\n\n", color.Bold("Assistant:"), resp.Message.Content)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)
	chatCmd.Flags().StringVarP(&chatModel, "model", "m", "llama3.2:latest", "Ollama model to use")
	chatCmd.Flags().StringVar(&ollamaURL, "ollama-url", "http://localhost:11434", "Ollama server URL")
}
