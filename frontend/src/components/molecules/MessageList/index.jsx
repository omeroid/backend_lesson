import { useState, useEffect, useRef, useCallback } from 'react'
import { motion } from 'framer-motion'
import { MessageItem } from './MessageItem'

import { useListMessages, useDeleteMessage } from '../../../modules/message'
import { useUser } from '../../../modules/user'

export const MessageList = ({ roomId }) => {
  const [messages, setMessages] = useState([])
  const { data } = useListMessages(roomId)
  const { handleDeleteMessage } = useDeleteMessage(roomId)
  const { user } = useUser()

  const scrollRef = useRef(null)
  const containerRef = useRef(null)
  const [userScrolled, setUserScrolled] = useState(false)
  
  const scrollToBottomOfList = useCallback(() => {
    if (!userScrolled && scrollRef.current) {
      scrollRef.current.scrollIntoView({
        behavior: 'smooth',
        block: 'end',
      })
    }
  }, [userScrolled])

  const handleScroll = useCallback(() => {
    if (!containerRef.current) return
    const { scrollTop, scrollHeight, clientHeight } = containerRef.current
    const isAtBottom = scrollHeight - scrollTop - clientHeight < 100
    setUserScrolled(!isAtBottom)
  }, [])

  useEffect(() => {
    if (!data || !data.messages) {
      setMessages([])
      return
    }
    if (!user) return
    const list = data.messages.map((item) => ({
      id: item.id,
      position: 'left',  // 全てのメッセージを左側に表示
      type: 'text',
      title: item.user.name,
      text: item.text,
      removeButton: user.userId === item.user.id,
      isMyMessage: user.userId === item.user.id,  // 自分のメッセージかどうかのフラグを追加
      className: user.userId === item.user.id ? 'my-message' : '',
    }))
    setMessages(list)
  }, [data, user])

  useEffect(() => {
    scrollToBottomOfList()
  }, [messages, scrollToBottomOfList])

  // スクロールバーのスタイルを追加
  useEffect(() => {
    const style = document.createElement('style')
    style.textContent = `
      .custom-scrollbar::-webkit-scrollbar {
        width: 8px;
      }
      .custom-scrollbar::-webkit-scrollbar-track {
        background: #1a1a1a;
      }
      .custom-scrollbar::-webkit-scrollbar-thumb {
        background: #4a4a4a;
        border-radius: 4px;
      }
      .custom-scrollbar::-webkit-scrollbar-thumb:hover {
        background: #5a5a5a;
      }
    `
    document.head.appendChild(style)
    return () => {
      document.head.removeChild(style)
    }
  }, [])

  return (
    <div 
      ref={containerRef}
      onScroll={handleScroll}
      className="custom-scrollbar"
      style={{
        height: '100%',
        overflowY: 'auto',
        overflowX: 'hidden',
        backgroundColor: '#1a1a1a',
        paddingTop: '1rem',
        paddingBottom: '1rem',
        position: 'relative'
      }}
    >
      {messages.length === 0 ? (
        <motion.div
          initial={{ opacity: 0 }}
          animate={{ opacity: 1 }}
          style={{
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            height: '100%',
            minHeight: '400px'
          }}
        >
          <div style={{ textAlign: 'center' }}>
            <p style={{ color: 'rgb(107, 114, 128)', fontSize: '1.125rem', marginBottom: '0.5rem' }}>
              まだメッセージがありません
            </p>
            <p style={{ color: 'rgb(75, 85, 99)', fontSize: '0.875rem' }}>
              最初のメッセージを送信してみましょう！
            </p>
          </div>
        </motion.div>
      ) : (
        <div style={{ minHeight: '100%' }}>
          {messages.map((message) => (
            <MessageItem
              key={message.id}
              message={message}
              onDelete={(msg) => handleDeleteMessage(msg.id)}
            />
          ))}
          <div ref={scrollRef} style={{ height: '1px' }}></div>
        </div>
      )}
      
      {/* スクロールトップボタン */}
      {userScrolled && messages.length > 0 && (
        <motion.button
          initial={{ opacity: 0, scale: 0.8 }}
          animate={{ opacity: 1, scale: 1 }}
          exit={{ opacity: 0, scale: 0.8 }}
          onClick={() => {
            containerRef.current?.scrollTo({ top: containerRef.current.scrollHeight, behavior: 'smooth' })
            setUserScrolled(false)
          }}
          style={{
            position: 'fixed',
            bottom: '5rem',
            right: '2rem',
            backgroundColor: 'rgba(255, 255, 255, 0.9)',
            border: '1px solid rgba(0, 0, 0, 0.1)',
            borderRadius: '2rem',
            padding: '0.5rem 1rem',
            color: '#000000',
            fontSize: '0.875rem',
            cursor: 'pointer',
            display: 'flex',
            alignItems: 'center',
            gap: '0.5rem',
            boxShadow: '0 4px 12px rgba(0, 0, 0, 0.2)',
            transition: 'all 0.2s',
            zIndex: 10
          }}
          onMouseEnter={(e) => {
            e.currentTarget.style.backgroundColor = '#ffffff'
            e.currentTarget.style.transform = 'scale(1.05)'
          }}
          onMouseLeave={(e) => {
            e.currentTarget.style.backgroundColor = 'rgba(255, 255, 255, 0.9)'
            e.currentTarget.style.transform = 'scale(1)'
          }}
        >
          <svg style={{ width: '1rem', height: '1rem' }} fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 14l-7 7m0 0l-7-7m7 7V3" />
          </svg>
          新しいメッセージ
        </motion.button>
      )}
    </div>
  )
}