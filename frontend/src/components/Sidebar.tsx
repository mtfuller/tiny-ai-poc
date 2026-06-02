import { useState } from 'react'
import { Plus, MessageSquare, Folder, Trash2, ChevronDown, ChevronRight, Bot } from 'lucide-react'
import { type AppState } from '../App'
import { type Project, type Chat, listChats, createProject, deleteProject, createChat, deleteChat, getChatMessages, addContextFolder, removeContextFolder, listContextFolders, openFolderDialog } from '../wailsbridge'

interface Props {
  state: AppState
  updateState: (u: Partial<AppState>) => void
}

export default function Sidebar({ state, updateState }: Props) {
  const [newProjectName, setNewProjectName] = useState('')
  const [showNewProject, setShowNewProject] = useState(false)
  const [expandedProjects, setExpandedProjects] = useState<Set<string>>(new Set())

  async function handleCreateProject() {
    if (!newProjectName.trim()) return
    const project = await createProject(newProjectName.trim())
    updateState({ projects: [...state.projects, project] })
    setNewProjectName('')
    setShowNewProject(false)
  }

  async function handleSelectProject(project: Project) {
    const chats = await listChats(project.ID)
    const folders = await listContextFolders(project.ID)
    const projectWithFolders = { ...project, ContextFolders: folders }
    updateState({
      selectedProject: projectWithFolders,
      chats,
      selectedChat: null,
      messages: [],
      streamingChatID: null,
      streamingContent: '',
    })
    setExpandedProjects(s => new Set([...s, project.ID]))
  }

  async function handleDeleteProject(e: React.MouseEvent, id: string) {
    e.stopPropagation()
    await deleteProject(id)
    const projects = state.projects.filter(p => p.ID !== id)
    const updates: Partial<AppState> = { projects }
    if (state.selectedProject?.ID === id) {
      updates.selectedProject = null
      updates.chats = []
      updates.selectedChat = null
      updates.messages = []
    }
    updateState(updates)
  }

  async function handleCreateChat() {
    if (!state.selectedProject) return
    const chat = await createChat(state.selectedProject.ID, 'New Chat')
    updateState({ chats: [...state.chats, chat], selectedChat: chat, messages: [] })
  }

  async function handleSelectChat(chat: Chat) {
    const messages = await getChatMessages(chat.ID)
    updateState({ selectedChat: chat, messages, streamingChatID: null, streamingContent: '' })
  }

  async function handleDeleteChat(e: React.MouseEvent, chatID: string) {
    e.stopPropagation()
    await deleteChat(chatID)
    const chats = state.chats.filter(c => c.ID !== chatID)
    const updates: Partial<AppState> = { chats }
    if (state.selectedChat?.ID === chatID) {
      updates.selectedChat = null
      updates.messages = []
    }
    updateState(updates)
  }

  async function handleAddFolder() {
    if (!state.selectedProject) return
    const path = await openFolderDialog()
    if (!path) return
    const folder = await addContextFolder(state.selectedProject.ID, path)
    const updatedProject = {
      ...state.selectedProject,
      ContextFolders: [...(state.selectedProject.ContextFolders || []), folder],
    }
    const projects = state.projects.map(p => p.ID === updatedProject.ID ? updatedProject : p)
    updateState({ selectedProject: updatedProject, projects })
  }

  async function handleRemoveFolder(folderID: string) {
    if (!state.selectedProject) return
    await removeContextFolder(state.selectedProject.ID, folderID)
    const updatedProject = {
      ...state.selectedProject,
      ContextFolders: state.selectedProject.ContextFolders.filter(f => f.ID !== folderID),
    }
    const projects = state.projects.map(p => p.ID === updatedProject.ID ? updatedProject : p)
    updateState({ selectedProject: updatedProject, projects })
  }

  function toggleProject(id: string) {
    setExpandedProjects(s => {
      const next = new Set(s)
      if (next.has(id)) next.delete(id)
      else next.add(id)
      return next
    })
  }

  return (
    <aside className="w-64 flex flex-col bg-gray-800 border-r border-gray-700 overflow-hidden">
      {/* Header */}
      <div className="flex items-center gap-2 px-4 py-4 border-b border-gray-700">
        <Bot size={20} className="text-indigo-400" />
        <span className="font-semibold text-white">tiny-ai</span>
      </div>

      {/* Model selector */}
      <div className="px-3 py-2 border-b border-gray-700">
        <select
          value={state.selectedModel}
          onChange={e => updateState({ selectedModel: e.target.value })}
          className="w-full bg-gray-700 text-gray-200 text-xs rounded px-2 py-1.5 border border-gray-600 focus:outline-none focus:border-indigo-500"
        >
          {state.models.length === 0 ? (
            <option value={state.selectedModel}>{state.selectedModel}</option>
          ) : (
            state.models.map(m => (
              <option key={m.Name} value={m.Name}>{m.Name}</option>
            ))
          )}
        </select>
      </div>

      {/* Projects */}
      <div className="flex-1 overflow-y-auto py-2">
        <div className="flex items-center justify-between px-3 py-1 mb-1">
          <span className="text-xs font-medium text-gray-400 uppercase tracking-wider">Projects</span>
          <button
            onClick={() => setShowNewProject(v => !v)}
            className="text-gray-400 hover:text-white transition-colors"
            title="New project"
          >
            <Plus size={14} />
          </button>
        </div>

        {showNewProject && (
          <div className="px-3 mb-2">
            <input
              autoFocus
              type="text"
              placeholder="Project name…"
              value={newProjectName}
              onChange={e => setNewProjectName(e.target.value)}
              onKeyDown={e => {
                if (e.key === 'Enter') handleCreateProject()
                if (e.key === 'Escape') setShowNewProject(false)
              }}
              className="w-full bg-gray-700 text-sm text-gray-200 rounded px-2 py-1 border border-gray-600 focus:outline-none focus:border-indigo-500"
            />
          </div>
        )}

        {state.projects.map(project => {
          const expanded = expandedProjects.has(project.ID)
          const isSelected = state.selectedProject?.ID === project.ID
          return (
            <div key={project.ID}>
              <div
                className={`flex items-center gap-1 px-2 py-1.5 cursor-pointer group rounded mx-1 ${
                  isSelected ? 'bg-gray-700' : 'hover:bg-gray-750'
                }`}
                onClick={() => {
                  handleSelectProject(project)
                  toggleProject(project.ID)
                }}
              >
                <button
                  onClick={e => { e.stopPropagation(); toggleProject(project.ID) }}
                  className="text-gray-400 flex-shrink-0"
                >
                  {expanded ? <ChevronDown size={12} /> : <ChevronRight size={12} />}
                </button>
                <span className="flex-1 text-sm text-gray-200 truncate">{project.Name}</span>
                <button
                  onClick={e => handleDeleteProject(e, project.ID)}
                  className="opacity-0 group-hover:opacity-100 text-gray-500 hover:text-red-400 transition-all"
                >
                  <Trash2 size={12} />
                </button>
              </div>

              {expanded && isSelected && (
                <div className="ml-4 mr-1">
                  {/* New chat button */}
                  <button
                    onClick={handleCreateChat}
                    className="flex items-center gap-1.5 w-full px-2 py-1 text-xs text-indigo-400 hover:text-indigo-300 hover:bg-gray-700 rounded transition-colors"
                  >
                    <Plus size={11} />
                    New Chat
                  </button>

                  {/* Chats */}
                  {state.chats.map(chat => (
                    <div
                      key={chat.ID}
                      onClick={() => handleSelectChat(chat)}
                      className={`flex items-center gap-1.5 px-2 py-1 rounded cursor-pointer group ${
                        state.selectedChat?.ID === chat.ID
                          ? 'bg-indigo-600/30 text-indigo-200'
                          : 'text-gray-300 hover:bg-gray-700'
                      }`}
                    >
                      <MessageSquare size={11} className="flex-shrink-0 text-gray-400" />
                      <span className="flex-1 text-xs truncate">{chat.Title}</span>
                      <button
                        onClick={e => handleDeleteChat(e, chat.ID)}
                        className="opacity-0 group-hover:opacity-100 text-gray-500 hover:text-red-400 transition-all"
                      >
                        <Trash2 size={10} />
                      </button>
                    </div>
                  ))}

                  {/* Context folders */}
                  <div className="mt-2 mb-1">
                    <div className="flex items-center justify-between px-2 py-0.5">
                      <span className="text-xs text-gray-500">Context</span>
                      <button
                        onClick={handleAddFolder}
                        className="text-gray-500 hover:text-gray-300 transition-colors"
                        title="Add folder"
                      >
                        <Plus size={11} />
                      </button>
                    </div>
                    {(state.selectedProject?.ContextFolders || []).map(folder => (
                      <div key={folder.ID} className="flex items-center gap-1 px-2 py-0.5 group">
                        <Folder size={10} className="text-gray-500 flex-shrink-0" />
                        <span className="flex-1 text-xs text-gray-400 truncate" title={folder.Path}>
                          {folder.Path.split('/').pop() || folder.Path}
                        </span>
                        <button
                          onClick={() => handleRemoveFolder(folder.ID)}
                          className="opacity-0 group-hover:opacity-100 text-gray-500 hover:text-red-400 transition-all"
                        >
                          <Trash2 size={10} />
                        </button>
                      </div>
                    ))}
                  </div>
                </div>
              )}
            </div>
          )
        })}

        {state.projects.length === 0 && (
          <div className="px-3 py-4 text-center text-xs text-gray-500">
            Create a project to get started
          </div>
        )}
      </div>
    </aside>
  )
}
