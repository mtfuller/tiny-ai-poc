import { useState, useEffect, useCallback } from 'react'
import Sidebar from './components/Sidebar'
import ChatView from './components/ChatView'
import WelcomeView from './components/WelcomeView'
import { type Project, type Chat, type Message, type ModelInfo, listProjects, listModels, onChatToken, onChatError } from './wailsbridge'

export interface AppState {
  projects: Project[]
  selectedProject: Project | null
  chats: Chat[]
  selectedChat: Chat | null
  messages: Message[]
  models: ModelInfo[]
  selectedModel: string
  streamingChatID: string | null
  streamingContent: string
  error: string | null
}

export default function App() {
  const [state, setState] = useState<AppState>({
    projects: [],
    selectedProject: null,
    chats: [],
    selectedChat: null,
    messages: [],
    models: [],
    selectedModel: 'llama3.2:latest',
    streamingChatID: null,
    streamingContent: '',
    error: null,
  })

  useEffect(() => {
    listProjects().then(projects => setState(s => ({ ...s, projects })))
    listModels().then(models => {
      if (models.length > 0) {
        setState(s => ({ ...s, models, selectedModel: models[0].Name }))
      }
    }).catch(() => {})

    onChatToken((chatID, token, done) => {
      setState(s => {
        if (s.streamingChatID !== chatID) return s
        if (done) {
          const assistantMsg: Message = { role: 'assistant', content: s.streamingContent + token }
          return {
            ...s,
            messages: [...s.messages, assistantMsg],
            streamingChatID: null,
            streamingContent: '',
          }
        }
        return { ...s, streamingContent: s.streamingContent + token }
      })
    })

    onChatError((chatID, error) => {
      setState(s => {
        if (s.streamingChatID !== chatID) return s
        return { ...s, streamingChatID: null, streamingContent: '', error }
      })
    })
  }, [])

  const updateState = useCallback((updates: Partial<AppState>) => {
    setState(s => ({ ...s, ...updates }))
  }, [])

  return (
    <div className="flex h-screen bg-gray-900 text-gray-100 overflow-hidden">
      <Sidebar state={state} updateState={updateState} />
      <main className="flex-1 flex flex-col min-w-0">
        {state.selectedChat ? (
          <ChatView state={state} updateState={updateState} />
        ) : (
          <WelcomeView state={state} updateState={updateState} />
        )}
      </main>
    </div>
  )
}
