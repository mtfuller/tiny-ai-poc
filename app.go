package main

import (
	"context"
	"os"
	"path/filepath"

	"github.com/mtfuller/tiny-ai-poc/internal/store"
	"github.com/mtfuller/tiny-ai-poc/pkg/ollama"
	wailsruntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx    context.Context
	ollama *ollama.Client
	store  *store.Store
}

func NewApp() *App {
	homeDir, _ := os.UserHomeDir()
	dataDir := filepath.Join(homeDir, ".tiny-ai")
	return &App{
		store: store.New(dataDir),
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.ollama = ollama.NewClient("http://localhost:11434")
	a.store.Load()
}

func (a *App) ListModels() ([]ollama.ModelInfo, error) {
	return a.ollama.ListModels(a.ctx)
}

func (a *App) CreateProject(name string) (store.Project, error) {
	return a.store.CreateProject(name)
}

func (a *App) ListProjects() []store.Project {
	return a.store.ListProjects()
}

func (a *App) DeleteProject(id string) error {
	return a.store.DeleteProject(id)
}

func (a *App) CreateChat(projectID, title string) (store.Chat, error) {
	return a.store.CreateChat(projectID, title)
}

func (a *App) ListChats(projectID string) []store.Chat {
	return a.store.ListChats(projectID)
}

func (a *App) GetChatMessages(chatID string) []ollama.Message {
	return a.store.GetMessages(chatID)
}

func (a *App) DeleteChat(chatID string) error {
	return a.store.DeleteChat(chatID)
}

func (a *App) UpdateChatTitle(chatID, title string) error {
	return a.store.UpdateChatTitle(chatID, title)
}

// SendMessage appends the user message and starts a streaming response.
// Tokens are emitted as "chat:token" events: { chatID, token, done }.
// Errors are emitted as "chat:error" events: { chatID, error }.
func (a *App) SendMessage(chatID, model, content string) error {
	msgs := a.store.GetMessages(chatID)
	userMsg := ollama.Message{Role: "user", Content: content}
	if err := a.store.AppendMessage(chatID, userMsg); err != nil {
		return err
	}
	msgs = append(msgs, userMsg)

	go func() {
		var full string
		err := a.ollama.ChatStream(a.ctx, ollama.ChatRequest{
			Model:    model,
			Messages: msgs,
		}, func(token string, done bool) {
			full += token
			wailsruntime.EventsEmit(a.ctx, "chat:token", map[string]any{
				"chatID": chatID,
				"token":  token,
				"done":   done,
			})
		})
		if err != nil {
			wailsruntime.EventsEmit(a.ctx, "chat:error", map[string]any{
				"chatID": chatID,
				"error":  err.Error(),
			})
			return
		}
		a.store.AppendMessage(chatID, ollama.Message{Role: "assistant", Content: full})
	}()

	return nil
}

func (a *App) OpenFolderDialog() (string, error) {
	return wailsruntime.OpenDirectoryDialog(a.ctx, wailsruntime.OpenDialogOptions{
		Title: "Select Folder for Context",
	})
}

func (a *App) AddContextFolder(projectID, path string) (store.ContextFolder, error) {
	return a.store.AddContextFolder(projectID, path)
}

func (a *App) RemoveContextFolder(projectID, folderID string) error {
	return a.store.RemoveContextFolder(projectID, folderID)
}

func (a *App) ListContextFolders(projectID string) []store.ContextFolder {
	return a.store.ListContextFolders(projectID)
}
