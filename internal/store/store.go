package store

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/mtfuller/tiny-ai-poc/pkg/ollama"
)

type Project struct {
	ID             string          `json:"id"`
	Name           string          `json:"name"`
	CreatedAt      time.Time       `json:"createdAt"`
	ContextFolders []ContextFolder `json:"contextFolders"`
}

type ContextFolder struct {
	ID   string `json:"id"`
	Path string `json:"path"`
}

type Chat struct {
	ID        string    `json:"id"`
	ProjectID string    `json:"projectId"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
}

type data struct {
	Projects []*Project `json:"projects"`
	Chats    []*Chat    `json:"chats"`
}

type Store struct {
	mu      sync.Mutex
	dataDir string
	d       data
	msgs    map[string][]ollama.Message // chatID -> messages
}

func New(dataDir string) *Store {
	return &Store{
		dataDir: dataDir,
		msgs:    make(map[string][]ollama.Message),
	}
}

func (s *Store) Load() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := os.MkdirAll(s.dataDir, 0755); err != nil {
		return err
	}

	path := filepath.Join(s.dataDir, "data.json")
	b, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &s.d)
}

func (s *Store) save() error {
	b, err := json.MarshalIndent(s.d, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(s.dataDir, "data.json"), b, 0644)
}

func (s *Store) loadMessages(chatID string) ([]ollama.Message, error) {
	if msgs, ok := s.msgs[chatID]; ok {
		return msgs, nil
	}
	path := filepath.Join(s.dataDir, "chats", chatID+".json")
	b, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var msgs []ollama.Message
	if err := json.Unmarshal(b, &msgs); err != nil {
		return nil, err
	}
	s.msgs[chatID] = msgs
	return msgs, nil
}

func (s *Store) saveMessages(chatID string, msgs []ollama.Message) error {
	dir := filepath.Join(s.dataDir, "chats")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	b, err := json.MarshalIndent(msgs, "", "  ")
	if err != nil {
		return err
	}
	s.msgs[chatID] = msgs
	return os.WriteFile(filepath.Join(dir, chatID+".json"), b, 0644)
}

func newID() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func (s *Store) CreateProject(name string) (Project, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	p := &Project{
		ID:             newID(),
		Name:           name,
		CreatedAt:      time.Now(),
		ContextFolders: []ContextFolder{},
	}
	s.d.Projects = append(s.d.Projects, p)
	return *p, s.save()
}

func (s *Store) ListProjects() []Project {
	s.mu.Lock()
	defer s.mu.Unlock()
	result := make([]Project, len(s.d.Projects))
	for i, p := range s.d.Projects {
		result[i] = *p
	}
	return result
}

func (s *Store) DeleteProject(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, p := range s.d.Projects {
		if p.ID == id {
			s.d.Projects = append(s.d.Projects[:i], s.d.Projects[i+1:]...)
			return s.save()
		}
	}
	return fmt.Errorf("project %s not found", id)
}

func (s *Store) AddContextFolder(projectID, path string) (ContextFolder, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	cf := ContextFolder{ID: newID(), Path: path}
	for _, p := range s.d.Projects {
		if p.ID == projectID {
			p.ContextFolders = append(p.ContextFolders, cf)
			return cf, s.save()
		}
	}
	return ContextFolder{}, fmt.Errorf("project %s not found", projectID)
}

func (s *Store) RemoveContextFolder(projectID, folderID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, p := range s.d.Projects {
		if p.ID == projectID {
			for i, cf := range p.ContextFolders {
				if cf.ID == folderID {
					p.ContextFolders = append(p.ContextFolders[:i], p.ContextFolders[i+1:]...)
					return s.save()
				}
			}
		}
	}
	return fmt.Errorf("folder %s not found", folderID)
}

func (s *Store) ListContextFolders(projectID string) []ContextFolder {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, p := range s.d.Projects {
		if p.ID == projectID {
			result := make([]ContextFolder, len(p.ContextFolders))
			copy(result, p.ContextFolders)
			return result
		}
	}
	return []ContextFolder{}
}

func (s *Store) CreateChat(projectID, title string) (Chat, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	c := &Chat{
		ID:        newID(),
		ProjectID: projectID,
		Title:     title,
		CreatedAt: time.Now(),
	}
	s.d.Chats = append(s.d.Chats, c)
	return *c, s.save()
}

func (s *Store) ListChats(projectID string) []Chat {
	s.mu.Lock()
	defer s.mu.Unlock()
	var result []Chat
	for _, c := range s.d.Chats {
		if c.ProjectID == projectID {
			result = append(result, *c)
		}
	}
	if result == nil {
		result = []Chat{}
	}
	return result
}

func (s *Store) DeleteChat(chatID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, c := range s.d.Chats {
		if c.ID == chatID {
			s.d.Chats = append(s.d.Chats[:i], s.d.Chats[i+1:]...)
			delete(s.msgs, chatID)
			os.Remove(filepath.Join(s.dataDir, "chats", chatID+".json"))
			return s.save()
		}
	}
	return fmt.Errorf("chat %s not found", chatID)
}

func (s *Store) GetMessages(chatID string) []ollama.Message {
	s.mu.Lock()
	defer s.mu.Unlock()
	msgs, _ := s.loadMessages(chatID)
	if msgs == nil {
		return []ollama.Message{}
	}
	return msgs
}

func (s *Store) AppendMessage(chatID string, msg ollama.Message) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	msgs, err := s.loadMessages(chatID)
	if err != nil {
		return err
	}
	msgs = append(msgs, msg)
	return s.saveMessages(chatID, msgs)
}

func (s *Store) UpdateChatTitle(chatID, title string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, c := range s.d.Chats {
		if c.ID == chatID {
			c.Title = title
			return s.save()
		}
	}
	return fmt.Errorf("chat %s not found", chatID)
}
