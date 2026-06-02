import { Bot, MessageSquare, FolderOpen, Cpu } from 'lucide-react'
import { type AppState } from '../App'

interface Props {
  state: AppState
  updateState: (u: Partial<AppState>) => void
}

export default function WelcomeView({ state }: Props) {
  return (
    <div className="flex flex-col items-center justify-center h-full text-center px-8">
      <div className="w-16 h-16 rounded-2xl bg-indigo-600 flex items-center justify-center mb-6">
        <Bot size={32} className="text-white" />
      </div>
      <h1 className="text-2xl font-semibold text-white mb-2">tiny-ai</h1>
      <p className="text-gray-400 text-sm max-w-sm mb-10">
        Local AI assistant powered by Ollama. Small models, smart pipelines.
      </p>

      <div className="grid grid-cols-3 gap-4 max-w-lg w-full">
        {[
          { icon: <MessageSquare size={18} />, title: 'Conversations', desc: 'Multi-turn chat with full history' },
          { icon: <FolderOpen size={18} />, title: 'Context', desc: 'Add local folders to include in prompts' },
          { icon: <Cpu size={18} />, title: 'Local models', desc: `Using ${state.selectedModel}` },
        ].map(card => (
          <div key={card.title} className="bg-gray-800 rounded-xl p-4 border border-gray-700 text-left">
            <div className="text-indigo-400 mb-2">{card.icon}</div>
            <div className="text-sm font-medium text-gray-200 mb-1">{card.title}</div>
            <div className="text-xs text-gray-500">{card.desc}</div>
          </div>
        ))}
      </div>

      {state.projects.length === 0 && (
        <p className="mt-8 text-xs text-gray-500">
          Create a project in the sidebar to start chatting →
        </p>
      )}
    </div>
  )
}
