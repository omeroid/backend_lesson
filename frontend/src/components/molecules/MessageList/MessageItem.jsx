import React from 'react'
import { motion } from 'framer-motion'

export const MessageItem = ({ message, onDelete }) => {
  const isMyMessage = message.isMyMessage || false
  
  return (
    <motion.div
      initial={{ opacity: 0, x: -20 }}
      animate={{ opacity: 1, x: 0 }}
      transition={{ duration: 0.2 }}
      style={{
        display: 'flex',
        justifyContent: 'flex-start',
        marginBottom: '1rem',
        paddingLeft: '1rem',
        paddingRight: '1rem'
      }}
    >
      <div style={{
        display: 'flex',
        gap: '0.75rem',
        maxWidth: '70%'
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
            background: isMyMessage 
              ? 'linear-gradient(135deg, #ffffff, #f0f0f0)' 
              : 'linear-gradient(135deg, #4a4a4a, #2a2a2a)',
            boxShadow: isMyMessage 
              ? '0 4px 12px rgba(255, 255, 255, 0.3), inset 0 1px 3px rgba(0, 0, 0, 0.1)' 
              : '0 4px 12px rgba(0, 0, 0, 0.3)',
            border: isMyMessage 
              ? '2px solid #ffffff' 
              : '2px solid transparent'
          }}>
            {message.title?.charAt(0).toUpperCase() || 'U'}
          </div>
        </motion.div>

        {/* メッセージ本体 */}
        <div style={{
          display: 'flex',
          flexDirection: 'column',
          gap: '0.25rem',
          flex: 1
        }}>
          <div style={{
            display: 'flex',
            alignItems: 'center',
            gap: '0.5rem'
          }}>
            <span style={{
              fontSize: '0.875rem',
              color: isMyMessage ? '#ffffff' : 'rgb(209, 213, 219)',
              fontWeight: '600',
              display: 'flex',
              alignItems: 'center',
              gap: '0.5rem'
            }}>
              {message.title}
              {isMyMessage && (
                <span style={{
                  padding: '0.125rem 0.5rem',
                  fontSize: '0.625rem',
                  color: 'black',
                  fontWeight: 'bold',
                  background: 'linear-gradient(135deg, #ffffff, #e0e0e0)',
                  borderRadius: '0.75rem',
                  textTransform: 'uppercase',
                  letterSpacing: '0.05em',
                  boxShadow: '0 2px 4px rgba(255, 255, 255, 0.2)'
                }}>
                  You
                </span>
              )}
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
                ? 'linear-gradient(135deg, rgba(255, 255, 255, 0.15), rgba(255, 255, 255, 0.08))' 
                : 'linear-gradient(135deg, rgba(0, 0, 0, 0.4), rgba(0, 0, 0, 0.2))',
              border: isMyMessage
                ? '2px solid rgba(255, 255, 255, 0.3)'
                : '1px solid rgba(255, 255, 255, 0.05)',
              borderLeft: isMyMessage
                ? '4px solid #ffffff'
                : '4px solid #4a4a4a',
              borderRadius: '0.75rem',
              padding: '0.75rem 1rem',
              backdropFilter: 'blur(10px)',
              boxShadow: isMyMessage
                ? '0 4px 16px rgba(255, 255, 255, 0.15), inset 0 1px 3px rgba(255, 255, 255, 0.1)'
                : '0 4px 12px rgba(0, 0, 0, 0.2)',
              minWidth: '4rem',
              maxWidth: '100%',
              transition: 'all 0.2s'
            }}
            onMouseEnter={(e) => {
              if (isMyMessage) {
                e.currentTarget.style.borderLeftWidth = '6px'
                e.currentTarget.style.transform = 'translateX(2px)'
              }
              if (message.removeButton) {
                const deleteBtn = e.currentTarget.querySelector('button')
                if (deleteBtn) deleteBtn.style.opacity = '1'
              }
            }}
            onMouseLeave={(e) => {
              if (isMyMessage) {
                e.currentTarget.style.borderLeftWidth = '4px'
                e.currentTarget.style.transform = 'translateX(0)'
              }
              if (message.removeButton) {
                const deleteBtn = e.currentTarget.querySelector('button')
                if (deleteBtn) deleteBtn.style.opacity = '0'
              }
            }}
          >
            <p style={{
              color: isMyMessage ? '#ffffff' : 'rgba(255, 255, 255, 0.9)',
              fontSize: '0.875rem',
              lineHeight: '1.5',
              margin: 0,
              wordBreak: 'break-word',
              fontWeight: isMyMessage ? '500' : '400'
            }}>
              {message.text}
            </p>
            
            {/* 自分のメッセージの背景装飾 */}
            {isMyMessage && (
              <div style={{
                position: 'absolute',
                top: '-2px',
                right: '-2px',
                width: '8px',
                height: '8px',
                background: '#ffffff',
                borderRadius: '50%',
                boxShadow: '0 0 10px rgba(255, 255, 255, 0.5)'
              }} />
            )}
            
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