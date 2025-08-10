import * as React from 'react'
import { useState } from 'react'
import { motion } from 'framer-motion'
import { useCreateMessage } from '../../../modules/message'

export const MessageForm = ({ roomId, roomName }) => {
  const { isMutating, handleCreateMessage } = useCreateMessage(roomId)
  const [chatInput, setChatInput] = useState('')

  const handleChatSubmit = async (e) => {
    e.preventDefault()
    if (!chatInput.trim()) return
    await handleCreateMessage(chatInput)
    setChatInput('')
  }

  if (!roomId) return null

  return (
    <form onSubmit={handleChatSubmit} style={{
      display: 'flex',
      alignItems: 'center',
      gap: '0.75rem'
    }}>
      <motion.button
        type="button"
        whileHover={{ scale: 1.1 }}
        whileTap={{ scale: 0.95 }}
        style={{
          background: 'none',
          border: 'none',
          color: 'rgb(156, 163, 175)',
          cursor: 'pointer',
          padding: 0,
          transition: 'color 0.2s'
        }}
        onMouseEnter={(e) => e.currentTarget.style.color = 'white'}
        onMouseLeave={(e) => e.currentTarget.style.color = 'rgb(156, 163, 175)'}
      >
        <svg style={{ height: '1.5rem', width: '1.5rem' }} fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 9v3m0 0v3m0-3h3m-3 0H9m12 0a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
      </motion.button>

      <div style={{ flex: 1, position: 'relative' }}>
        <input
          type="text"
          value={chatInput}
          onChange={(e) => setChatInput(e.target.value)}
          placeholder={`#${roomName || 'general'} にメッセージを送信`}
          style={{
            width: '100%',
            padding: '0.75rem 5rem 0.75rem 1rem',
            backgroundColor: 'rgba(55, 65, 81, 0.8)',
            color: 'white',
            borderRadius: '1.5rem',
            outline: 'none',
            border: '1px solid rgba(75, 85, 99, 0.5)',
            transition: 'all 0.2s',
            fontSize: '0.875rem',
            backdropFilter: 'blur(10px)'
          }}
          disabled={isMutating}
          onFocus={(e) => {
            e.target.style.border = '1px solid rgba(147, 51, 234, 0.5)'
            e.target.style.boxShadow = '0 0 0 3px rgba(147, 51, 234, 0.1)'
          }}
          onBlur={(e) => {
            e.target.style.border = '1px solid rgba(75, 85, 99, 0.5)'
            e.target.style.boxShadow = 'none'
          }}
        />
        
        <div style={{
          position: 'absolute',
          right: '0.5rem',
          top: '50%',
          transform: 'translateY(-50%)',
          display: 'flex',
          alignItems: 'center',
          gap: '0.5rem'
        }}>
          <motion.button
            type="button"
            whileHover={{ scale: 1.1 }}
            style={{
              background: 'none',
              border: 'none',
              color: 'rgb(156, 163, 175)',
              cursor: 'pointer',
              padding: 0,
              transition: 'color 0.2s'
            }}
            onMouseEnter={(e) => e.currentTarget.style.color = 'white'}
            onMouseLeave={(e) => e.currentTarget.style.color = 'rgb(156, 163, 175)'}
          >
            <svg style={{ height: '1.25rem', width: '1.25rem' }} fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15.172 7l-6.586 6.586a2 2 0 102.828 2.828l6.414-6.586a4 4 0 00-5.656-5.656l-6.415 6.585a6 6 0 108.486 8.486L20.5 13" />
            </svg>
          </motion.button>
          <motion.button
            type="button"
            whileHover={{ scale: 1.1 }}
            style={{
              background: 'none',
              border: 'none',
              color: 'rgb(156, 163, 175)',
              cursor: 'pointer',
              padding: 0,
              transition: 'color 0.2s'
            }}
            onMouseEnter={(e) => e.currentTarget.style.color = 'white'}
            onMouseLeave={(e) => e.currentTarget.style.color = 'rgb(156, 163, 175)'}
          >
            <svg style={{ height: '1.25rem', width: '1.25rem' }} fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M14.828 14.828a4 4 0 01-5.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </motion.button>
        </div>
      </div>

      <motion.button
        type="submit"
        disabled={isMutating || !chatInput.trim()}
        whileHover={{ scale: 1.05 }}
        whileTap={{ scale: 0.95 }}
        style={{
          padding: '0.75rem',
          borderRadius: '50%',
          transition: 'all 0.2s',
          background: chatInput.trim() && !isMutating
            ? 'linear-gradient(135deg, rgb(147, 51, 234), rgb(236, 72, 153))'
            : 'rgb(55, 65, 81)',
          color: chatInput.trim() && !isMutating ? 'white' : 'rgb(107, 114, 128)',
          cursor: chatInput.trim() && !isMutating ? 'pointer' : 'not-allowed',
          boxShadow: chatInput.trim() && !isMutating
            ? '0 4px 12px rgba(147, 51, 234, 0.3)'
            : 'none',
          border: 'none',
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center'
        }}
      >
        {isMutating ? (
          <div style={{
            width: '1.25rem',
            height: '1.25rem',
            border: '2px solid white',
            borderTopColor: 'transparent',
            borderRadius: '50%',
            animation: 'spin 1s linear infinite'
          }} />
        ) : (
          <svg style={{ height: '1.25rem', width: '1.25rem' }} fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8" />
          </svg>
        )}
      </motion.button>
    </form>
  )
}