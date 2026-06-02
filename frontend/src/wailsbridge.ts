// TypeScript wrappers for Wails Go bindings.
// These match the method signatures on App in app.go.

declare global {
  interface Window {
    go?: {
      main?: {
        App?: {
          ListModels: () => Promise<Array<{Name: string}>>
          CreateProject: (name: string) => Promise<{ID: string; Name: string; CreatedAt: string; ContextFolders: ContextFolder[]}>
          ListProjects: () => Promise<Project[]>
          DeleteProject: (id: string) => Promise<void>
          CreateChat: (projectID: string, title: string) => Promise<Chat>
          ListChats: (projectID: string) => Promise<Chat[]>
          GetChatMessages: (chatID: string) => Promise<Message[]>
          DeleteChat: (chatID: string) => Promise<void>
          UpdateChatTitle: (chatID: string, title: string) => Promise<void>
          SendMessage: (chatID: string, model: string, content: string) => Promise<void>
          OpenFolderDialog: () => Promise<string>
          AddContextFolder: (projectID: string, path: string) => Promise<ContextFolder>
          RemoveContextFolder: (projectID: string, folderID: string) => Promise<void>
          ListContextFolders: (projectID: string) => Promise<ContextFolder[]>
        }
      }
    }
    runtime?: {
      EventsOn: (event: string, cb: (data: unknown) => void) => void
      EventsOff: (...events: string[]) => void
    }
  }
}

export interface Project {
  ID: string
  Name: string
  CreatedAt: string
  ContextFolders: ContextFolder[]
}

export interface Chat {
  ID: string
  ProjectID: string
  Title: string
  CreatedAt: string
}

export interface Message {
  role: string
  content: string
}

export interface ContextFolder {
  ID: string
  Path: string
}

export interface ModelInfo {
  Name: string
}

function getApp() {
  return window.go?.main?.App
}

export async function listModels(): Promise<ModelInfo[]> {
  const app = getApp()
  if (!app) return [{Name: 'llama3.2:latest'}]
  return app.ListModels()
}

export async function createProject(name: string): Promise<Project> {
  return getApp()!.CreateProject(name)
}

export async function listProjects(): Promise<Project[]> {
  const app = getApp()
  if (!app) return []
  return app.ListProjects()
}

export async function deleteProject(id: string): Promise<void> {
  return getApp()!.DeleteProject(id)
}

export async function createChat(projectID: string, title: string): Promise<Chat> {
  return getApp()!.CreateChat(projectID, title)
}

export async function listChats(projectID: string): Promise<Chat[]> {
  const app = getApp()
  if (!app) return []
  return app.ListChats(projectID)
}

export async function getChatMessages(chatID: string): Promise<Message[]> {
  const app = getApp()
  if (!app) return []
  return app.GetChatMessages(chatID)
}

export async function deleteChat(chatID: string): Promise<void> {
  return getApp()!.DeleteChat(chatID)
}

export async function updateChatTitle(chatID: string, title: string): Promise<void> {
  return getApp()!.UpdateChatTitle(chatID, title)
}

export async function sendMessage(chatID: string, model: string, content: string): Promise<void> {
  return getApp()!.SendMessage(chatID, model, content)
}

export async function openFolderDialog(): Promise<string> {
  return getApp()!.OpenFolderDialog()
}

export async function addContextFolder(projectID: string, path: string): Promise<ContextFolder> {
  return getApp()!.AddContextFolder(projectID, path)
}

export async function removeContextFolder(projectID: string, folderID: string): Promise<void> {
  return getApp()!.RemoveContextFolder(projectID, folderID)
}

export async function listContextFolders(projectID: string): Promise<ContextFolder[]> {
  const app = getApp()
  if (!app) return []
  return app.ListContextFolders(projectID)
}

export function onChatToken(cb: (chatID: string, token: string, done: boolean) => void) {
  window.runtime?.EventsOn('chat:token', (data: unknown) => {
    const d = data as {chatID: string; token: string; done: boolean}
    cb(d.chatID, d.token, d.done)
  })
}

export function onChatError(cb: (chatID: string, error: string) => void) {
  window.runtime?.EventsOn('chat:error', (data: unknown) => {
    const d = data as {chatID: string; error: string}
    cb(d.chatID, d.error)
  })
}

export function offEvents(...events: string[]) {
  window.runtime?.EventsOff(...events)
}
