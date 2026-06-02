import { useEffect, useRef, useState } from 'react'
import { Send, Loader2 } from 'lucide-react'
import { type AppState } from '../App'
import { sendMessage, updateChatTitle } from '../wailsbridge'

interface Props {
  state: AppState
  updateState: (u: Partial<AppState>) => void
}

export default function ChatView({ state, updateState }: Props) {
  const [input, setInput] = useState('')
  const bottomRef = useRef<HTMLDivElement>(null)
  const textareaRef = useRef<HTMLTextAreaElement>(null)
  const isStreaming = state.streamingChatID === state.selectedChat?.ID

  useEffect(() => {
    bottomRef.current?.scrollIntoView({ behavior: 'smooth' })
  }, [state.messages, state.streamingContent])

  async function handleSend() {
    if (!input.trim() || !state.selectedChat || isStreaming) return
    const chatID = state.selectedChat.ID
    const content = input.trim()
    setInput('')

    const userMsg = { role: 'user', content }
    updateState({
      messages: [...state.messages, userMsg],
      streamingChatID: chatID,
      streamingContent: '',
    })

    // Auto-title on first message
    if (state.messages.length === 0 && state.selectedChat.Title === 'New Chat') {
      const title = content.slice(0, 40) + (content.length > 40 ? '…' : '')
      await updateChatTitle(chatID, title)
      const chats = state.chats.map(c => c.ID === chatID ? { ...c, Title: title } : c)
      updateState({ chats, selectedChat: { ...state.selectedChat, Title: title } })
    }

    await sendMessage(chatID, state.selectedModel, content)
  }

  function handleKeyDown(e: React.KeyboardEvent) {
    if (e.key === 'Enter' && !e.shiftKey) {
      e.preventDefault()
      handleSend()
    }
  }

  function autoResize(e: React.ChangeEvent<HTMLTextAreaElement>) {
    setInput(e.target.value)
    const el = e.target
    el.style.height = 'auto'
    el.style.height = Math.min(el.scrollHeight, 200) + 'px'
  }

  const allMessages = [...state.messages]
  const streamingMsg = isStreaming && state.streamingContent
    ? { role: 'assistant', content: state.streamingContent }
    : null

  return (
    <div className="flex flex-col h-full">
      {/* Header */}
      <div className="flex items-center px-6 py-3 border-b border-gray-700 bg-gray-850">
        <h1 className="text-sm font-medium text-gray-200 truncate">
          {state.selectedChat?.Title || 'Chat'}
        </h1>
        <span className="ml-auto text-xs text-gray-500">{state.selectedModel}</span>
      </div>

      {/* Messages */}
      <div className="flex-1 overflow-y-auto px-4 py-6 space-y-6">
        {allMessages.map((msg, i) => (
          <MessageBubble key={i} role={msg.role} content={msg.content} />
        ))}
        {streamingMsg && (
          <MessageBubble role="assistant" content={streamingMsg.content} streaming />
        )}
        {isStreaming && !state.streamingContent && (
          <div className="flex items-start gap-3">
            <div className="w-7 h-7 rounded-full bg-indigo-600 flex items-center justify-center flex-shrink-0 mt-0.5">
              <Loader2 size={14} className="animate-spin text-white" />
            </div>
          </div>
        )}
        <div ref={bottomRef} />
      </div>

      {/* Input */}
      <div className="px-4 py-4 border-t border-gray-700">
        <div className="flex gap-3 items-end bg-gray-700 rounded-xl px-4 py-3 border border-gray-600 focus-within:border-indigo-500 transition-colors">
          <textarea
            ref={textareaRef}
            rows={1}
            value={input}
            onChange={autoResize}
            onKeyDown={handleKeyDown}
            placeholder="Message…"
            disabled={isStreaming}
            className="flex-1 bg-transparent resize-none text-gray-100 text-sm placeholder-gray-500 focus:outline-none disabled:opacity-50"
            style={{ minHeight: '24px', maxHeight: '200px' }}
          />
          <button
            onClick={handleSend}
            disabled={!input.trim() || isStreaming}
            className="flex-shrink-0 w-8 h-8 rounded-lg bg-indigo-600 hover:bg-indigo-500 disabled:bg-gray-600 disabled:cursor-not-allowed flex items-center justify-center transition-colors"
          >
            {isStreaming ? (
              <Loader2 size={14} className="animate-spin text-white" />
            ) : (
              <Send size={14} className="text-white" />
            )}
          </button>
        </div>
        <p className="text-xs text-gray-600 text-center mt-2">Enter to send · Shift+Enter for newline</p>
      </div>
    </div>
  )
}

function MessageBubble({ role, content, streaming }: { role: string; content: string; streaming?: boolean }) {
  const isUser = role === 'user'
  return (
    <div className={`flex items-start gap-3 ${isUser ? 'flex-row-reverse' : ''}`}>
      <div className={`w-7 h-7 rounded-full flex items-center justify-center flex-shrink-0 mt-0.5 text-xs font-bold ${
        isUser ? 'bg-gray-600 text-gray-200' : 'bg-indigo-600 text-white'
      }`}>
        {isUser ? 'U' : 'AI'}
      </div>
      <div className={`max-w-[75%] rounded-2xl px-4 py-2.5 text-sm leading-relaxed ${
        isUser
          ? 'bg-indigo-600 text-white rounded-tr-sm'
          : 'bg-gray-700 text-gray-100 rounded-tl-sm'
      }`}>
        <pre className="whitespace-pre-wrap font-sans">{content}{streaming ? '▋' : ''}</pre>
      </div>
    </div>
  )
}
