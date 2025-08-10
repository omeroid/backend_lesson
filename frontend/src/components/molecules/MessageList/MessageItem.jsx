import React from 'react'
import { motion } from 'framer-motion'

export const MessageItem = ({ message, onDelete }) => {
  const isMyMessage = message.position === 'right'
  
  return (
    <motion.div
      initial={{ opacity: 0, y: 10 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.2 }}
      style={{
        display: 'flex',
        justifyContent: isMyMessage ? 'flex-end' : 'flex-start',
        marginBottom: '1rem',
        paddingLeft: '1rem',
        paddingRight: '1rem',
        position: 'relative'
      }}
    >
      <div style={{
        display: 'flex',
        gap: '0.75rem',
        maxWidth: '70%',
        flexDirection: isMyMessage ? 'row-reverse' : 'row'
      }}>
        {/* アバター */}
        <motion.div 
          whileHover={{ scale: 1.1 }}
          style={{ flexShrink: 0 }}
        >
          <div style={{
            width: '2.5rem',
            height: '2.5rem',
            borderRadius: '50%',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            color: isMyMessage ? 'black' : 'white',
            fontWeight: 'bold',
            fontSize: '0.875rem',
            background: isMyMessage ? '#ffffff' : '#333333',
            boxShadow: '0 4px 12px rgba(0, 0, 0, 0.3)'
          }}>
            {message.title?.charAt(0).toUpperCase() || 'U'}
          </div>
        </motion.div>

        {/* メッセージ本体 */}
        <div style={{
          display: 'flex',
          flexDirection: 'column',
          alignItems: isMyMessage ? 'flex-end' : 'flex-start',
          gap: '0.25rem'
        }}>
          <div style={{
            display: 'flex',
            alignItems: 'center',
            gap: '0.5rem',
            paddingLeft: isMyMessage ? 0 : '0.5rem',
            paddingRight: isMyMessage ? '0.5rem' : 0
          }}>
            <span style={{
              fontSize: '0.75rem',
              color: 'rgb(209, 213, 219)',
              fontWeight: '600'
            }}>
              {message.title}
            </span>
            <span style={{
              fontSize: '0.625rem',
              color: 'rgb(107, 114, 128)'
            }}>
              {new Date().toLocaleTimeString('ja-JP', { hour: '2-digit', minute: '2-digit' })}
            </span>
          </div>
          
          <motion.div
            whileHover={{ scale: 1.01 }}
            style={{
              position: 'relative',
              background: isMyMessage 
                ? 'rgba(255, 255, 255, 0.1)' 
                : 'rgba(0, 0, 0, 0.3)',
              border: isMyMessage
                ? '1px solid rgba(255, 255, 255, 0.2)'
                : '1px solid rgba(255, 255, 255, 0.1)',
              borderRadius: isMyMessage 
                ? '1.25rem 1.25rem 0.25rem 1.25rem'
                : '1.25rem 1.25rem 1.25rem 0.25rem',
              padding: '0.75rem 1rem',
              backdropFilter: 'blur(10px)',
              boxShadow: '0 4px 12px rgba(0, 0, 0, 0.2)',
              minWidth: '4rem'
            }}
            onMouseEnter={(e) => {
              if (message.removeButton) {
                const deleteBtn = e.currentTarget.querySelector('button')
                if (deleteBtn) deleteBtn.style.opacity = '1'
              }
            }}
            onMouseLeave={(e) => {
              if (message.removeButton) {
                const deleteBtn = e.currentTarget.querySelector('button')
                if (deleteBtn) deleteBtn.style.opacity = '0'
              }
            }}
          >
            <p style={{
              color: 'white',
              fontSize: '0.875rem',
              lineHeight: '1.5',
              margin: 0,
              wordBreak: 'break-word'
            }}>
              {message.text}
            </p>
            
            {message.removeButton && (
              <motion.button
                initial={{ opacity: 0 }}
                whileHover={{ scale: 1.2, rotate: 10 }}
                whileTap={{ scale: 0.9 }}
                onClick={() => onDelete(message)}
                style={{
                  position: 'absolute',
                  top: '-0.5rem',
                  right: '-0.5rem',
                  opacity: 0,
                  transition: 'opacity 0.2s',
                  padding: '0.375rem',
                  background: 'linear-gradient(135deg, rgb(239, 68, 68), rgb(220, 38, 38))',
                  borderRadius: '50%',
                  color: 'white',
                  border: 'none',
                  cursor: 'pointer',
                  boxShadow: '0 2px 8px rgba(239, 68, 68, 0.4)',
                  display: 'flex',
                  alignItems: 'center',
                  justifyContent: 'center'
                }}
              >
                <svg style={{ width: '0.75rem', height: '0.75rem' }} fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={3} d="M6 18L18 6M6 6l12 12" />
                </svg>
              </motion.button>
            )}
          </motion.div>
        </div>
      </div>
    </motion.div>
  )
}