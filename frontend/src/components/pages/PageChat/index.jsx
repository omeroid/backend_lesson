import React, { useState } from 'react'
import { motion, AnimatePresence } from 'framer-motion'
import { useNavigate } from 'react-router-dom'
import toast from 'react-hot-toast'

import { MessageList } from '../../molecules/MessageList'
import { MessageForm } from '../../molecules/MessageForm'
import { RoomList } from '../../molecules/RoomList'
import { RoomForm } from '../../molecules/RoomForm'
import { useUser } from '../../../modules/user'

export const PageChat = () => {
  const [selectedRoom, setSelectedRoom] = useState(null)
  const [showCreateRoom, setShowCreateRoom] = useState(false)
  const [mobileSidebarOpen, setMobileSidebarOpen] = useState(false)
  const { user } = useUser()
  const navigate = useNavigate()

  const handleLogout = () => {
    sessionStorage.removeItem('userData')
    toast.success('ログアウトしました')
    navigate('/')
  }

  const handleRoomSelect = (room) => {
    setSelectedRoom(room)
    setShowCreateRoom(false)
    setMobileSidebarOpen(false)
  }

  return (
    <div style={{ height: '100vh', display: 'flex', backgroundColor: 'rgb(17, 24, 39)' }}>
      {/* モバイルメニューボタン */}
      <button
        onClick={() => setMobileSidebarOpen(!mobileSidebarOpen)}
        style={{
          display: window.innerWidth >= 1024 ? 'none' : 'block',
          position: 'fixed',
          top: '1rem',
          left: '1rem',
          zIndex: 50,
          padding: '0.5rem',
          backgroundColor: 'rgb(31, 41, 55)',
          borderRadius: '0.5rem',
          color: 'white',
          border: 'none',
          cursor: 'pointer'
        }}
      >
        <svg style={{ width: '1.5rem', height: '1.5rem' }} fill="none" viewBox="0 0 24 24" stroke="currentColor">
          {mobileSidebarOpen ? (
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
          ) : (
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 6h16M4 12h16M4 18h16" />
          )}
        </svg>
      </button>

      {/* サイドバー - サーバー一覧 (Discord風) */}
      <motion.div 
        initial={{ x: -100 }}
        animate={{ x: 0 }}
        style={{
          display: window.innerWidth >= 1024 ? 'flex' : 'none',
          width: '5rem',
          backgroundColor: 'rgb(3, 7, 18)',
          flexDirection: 'column',
          alignItems: 'center',
          paddingTop: '1rem',
          paddingBottom: '1rem',
          gap: '0.75rem'
        }}
      >
        <motion.div
          whileHover={{ scale: 1.1 }}
          whileTap={{ scale: 0.95 }}
          style={{
            width: '3rem',
            height: '3rem',
            background: 'linear-gradient(to bottom right, rgb(147, 51, 234), rgb(236, 72, 153))',
            borderRadius: '1rem',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            cursor: 'pointer',
            transition: 'border-radius 0.2s'
          }}
          onMouseEnter={(e) => e.currentTarget.style.borderRadius = '0.75rem'}
          onMouseLeave={(e) => e.currentTarget.style.borderRadius = '1rem'}
        >
          <svg style={{ width: '1.75rem', height: '1.75rem', color: 'white' }} fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
          </svg>
        </motion.div>
        
        <div style={{ height: '1px', width: '2rem', backgroundColor: 'rgb(55, 65, 81)' }} />
        
        <motion.button
          whileHover={{ scale: 1.1 }}
          whileTap={{ scale: 0.95 }}
          onClick={() => setShowCreateRoom(!showCreateRoom)}
          style={{
            width: '3rem',
            height: '3rem',
            backgroundColor: 'rgb(31, 41, 55)',
            borderRadius: '1rem',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            cursor: 'pointer',
            transition: 'all 0.2s',
            border: 'none'
          }}
          onMouseEnter={(e) => {
            e.currentTarget.style.backgroundColor = 'rgb(147, 51, 234)'
            e.currentTarget.style.borderRadius = '0.75rem'
          }}
          onMouseLeave={(e) => {
            e.currentTarget.style.backgroundColor = 'rgb(31, 41, 55)'
            e.currentTarget.style.borderRadius = '1rem'
          }}
        >
          <svg style={{ width: '1.5rem', height: '1.5rem', color: 'rgb(156, 163, 175)' }} fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 4v16m8-8H4" />
          </svg>
        </motion.button>

        <div style={{ flex: 1 }} />

        <motion.button
          whileHover={{ scale: 1.1 }}
          whileTap={{ scale: 0.95 }}
          onClick={handleLogout}
          style={{
            width: '3rem',
            height: '3rem',
            backgroundColor: 'rgb(31, 41, 55)',
            borderRadius: '1rem',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            cursor: 'pointer',
            transition: 'all 0.2s',
            border: 'none'
          }}
          onMouseEnter={(e) => {
            e.currentTarget.style.backgroundColor = 'rgb(220, 38, 38)'
            e.currentTarget.style.borderRadius = '0.75rem'
          }}
          onMouseLeave={(e) => {
            e.currentTarget.style.backgroundColor = 'rgb(31, 41, 55)'
            e.currentTarget.style.borderRadius = '1rem'
          }}
        >
          <svg style={{ width: '1.5rem', height: '1.5rem', color: 'rgb(156, 163, 175)' }} fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
          </svg>
        </motion.button>
      </motion.div>

      {/* チャンネルサイドバー */}
      <AnimatePresence>
        {(window.innerWidth >= 1024 || mobileSidebarOpen) && (
          <motion.div
            initial={{ width: 0, opacity: 0 }}
            animate={{ width: mobileSidebarOpen ? '100%' : 240, opacity: 1 }}
            exit={{ width: 0, opacity: 0 }}
            transition={{ duration: 0.2 }}
            style={{
              position: mobileSidebarOpen ? 'fixed' : 'relative',
              inset: mobileSidebarOpen ? 0 : 'auto',
              zIndex: mobileSidebarOpen ? 40 : 'auto',
              display: 'flex',
              backgroundColor: '#2b2d31',
              borderRight: '1px solid rgb(31, 41, 55)',
              flexDirection: 'column',
              overflow: 'hidden'
            }}
          >
            <div style={{
              padding: '1rem',
              borderBottom: '1px solid rgb(31, 41, 55)',
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'space-between'
            }}>
              <h2 style={{
                color: 'white',
                fontWeight: 'bold',
                fontSize: '1.125rem',
                display: 'flex',
                alignItems: 'center',
                gap: '0.5rem'
              }}>
                <svg style={{ width: '1.25rem', height: '1.25rem' }} fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
                </svg>
                チャットルーム
              </h2>
              {mobileSidebarOpen && (
                <button
                  onClick={() => setMobileSidebarOpen(false)}
                  style={{
                    background: 'none',
                    border: 'none',
                    color: 'rgb(156, 163, 175)',
                    cursor: 'pointer',
                    padding: 0
                  }}
                  onMouseEnter={(e) => e.currentTarget.style.color = 'white'}
                  onMouseLeave={(e) => e.currentTarget.style.color = 'rgb(156, 163, 175)'}
                >
                  <svg style={{ width: '1.5rem', height: '1.5rem' }} fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
                  </svg>
                </button>
              )}
            </div>

            <div style={{
              flex: 1,
              overflowY: 'auto',
              padding: '0.5rem'
            }}>
              <RoomList onRoomSelect={handleRoomSelect} selectedRoom={selectedRoom} />
            </div>

            <div style={{
              padding: '1rem',
              borderTop: '1px solid rgb(31, 41, 55)',
              backgroundColor: 'rgba(17, 24, 39, 0.5)'
            }}>
              <div style={{
                display: 'flex',
                alignItems: 'center',
                gap: '0.75rem'
              }}>
                <div style={{
                  width: '2.5rem',
                  height: '2.5rem',
                  background: 'linear-gradient(to bottom right, rgb(34, 197, 94), rgb(59, 130, 246))',
                  borderRadius: '50%',
                  display: 'flex',
                  alignItems: 'center',
                  justifyContent: 'center',
                  color: 'white',
                  fontWeight: 'bold'
                }}>
                  {user?.userName?.charAt(0).toUpperCase() || 'U'}
                </div>
                <div style={{ flex: 1 }}>
                  <p style={{ color: 'white', fontSize: '0.875rem', fontWeight: '500' }}>
                    {user?.userName || 'Guest'}
                  </p>
                  <p style={{ color: 'rgb(156, 163, 175)', fontSize: '0.75rem' }}>
                    オンライン
                  </p>
                </div>
                <div style={{ display: 'flex', gap: '0.5rem' }}>
                  <motion.button
                    whileHover={{ rotate: 180 }}
                    transition={{ duration: 0.3 }}
                    style={{
                      background: 'none',
                      border: 'none',
                      color: 'rgb(156, 163, 175)',
                      cursor: 'pointer',
                      padding: 0
                    }}
                    onMouseEnter={(e) => e.currentTarget.style.color = 'white'}
                    onMouseLeave={(e) => e.currentTarget.style.color = 'rgb(156, 163, 175)'}
                  >
                    <svg style={{ width: '1.25rem', height: '1.25rem' }} fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                    </svg>
                  </motion.button>
                  {mobileSidebarOpen && (
                    <button
                      onClick={handleLogout}
                      style={{
                        background: 'none',
                        border: 'none',
                        color: 'rgb(156, 163, 175)',
                        cursor: 'pointer',
                        padding: 0
                      }}
                      onMouseEnter={(e) => e.currentTarget.style.color = 'rgb(239, 68, 68)'}
                      onMouseLeave={(e) => e.currentTarget.style.color = 'rgb(156, 163, 175)'}
                    >
                      <svg style={{ width: '1.25rem', height: '1.25rem' }} fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
                      </svg>
                    </button>
                  )}
                </div>
              </div>
            </div>
          </motion.div>
        )}
      </AnimatePresence>

      {/* ルーム作成ポップアップ */}
      <AnimatePresence>
        {showCreateRoom && (
          <>
            <motion.div
              initial={{ opacity: 0 }}
              animate={{ opacity: 1 }}
              exit={{ opacity: 0 }}
              onClick={() => setShowCreateRoom(false)}
              style={{
                position: 'fixed',
                inset: 0,
                backgroundColor: 'rgba(0, 0, 0, 0.7)',
                backdropFilter: 'blur(8px)',
                zIndex: 60,
                display: 'flex',
                alignItems: 'center',
                justifyContent: 'center'
              }}
            />
            <motion.div
              initial={{ opacity: 0, scale: 0.9, y: 20 }}
              animate={{ opacity: 1, scale: 1, y: 0 }}
              exit={{ opacity: 0, scale: 0.9, y: 20 }}
              transition={{ type: 'spring', damping: 20, stiffness: 300 }}
              style={{
                position: 'fixed',
                top: '50%',
                left: '50%',
                transform: 'translate(-50%, -50%)',
                zIndex: 70,
                width: '100%',
                maxWidth: '32rem',
                padding: '1rem'
              }}
            >
              <div className="glass-morphism" style={{
                borderRadius: '1.5rem',
                padding: '2rem',
                boxShadow: '0 25px 50px -12px rgba(0, 0, 0, 0.25)',
                border: '1px solid rgba(147, 51, 234, 0.2)',
                background: 'linear-gradient(135deg, rgba(31, 41, 55, 0.9), rgba(17, 24, 39, 0.9))'
              }}>
                <RoomForm onClose={() => setShowCreateRoom(false)} />
              </div>
            </motion.div>
          </>
        )}
      </AnimatePresence>

      {/* メインチャットエリア */}
      <div style={{
        flex: 1,
        display: 'flex',
        flexDirection: 'column',
        backgroundColor: 'rgb(31, 41, 55)'
      }}>
        {selectedRoom ? (
          <>
            {/* チャットヘッダー */}
            <div style={{
              height: '3.5rem',
              backgroundColor: 'rgb(31, 41, 55)',
              borderBottom: '1px solid rgb(55, 65, 81)',
              display: 'flex',
              alignItems: 'center',
              paddingLeft: '1rem',
              paddingRight: '1rem',
              justifyContent: 'space-between'
            }}>
              <div style={{
                display: 'flex',
                alignItems: 'center',
                gap: '0.75rem'
              }}>
                <svg style={{ width: '1.25rem', height: '1.25rem', color: 'rgb(156, 163, 175)' }} fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M7 20l4-16m2 16l4-16M6 9h14M4 15h14" />
                </svg>
                <h3 style={{ color: 'white', fontWeight: '600' }}>
                  {selectedRoom.name}
                </h3>
                <span style={{
                  display: window.innerWidth >= 640 ? 'inline' : 'none',
                  color: 'rgb(156, 163, 175)',
                  fontSize: '0.875rem'
                }}>
                  {selectedRoom.description || 'チャットルーム'}
                </span>
              </div>
              <div style={{
                display: 'flex',
                alignItems: 'center',
                gap: '0.75rem'
              }}>
                <motion.button
                  whileHover={{ scale: 1.1 }}
                  style={{
                    background: 'none',
                    border: 'none',
                    color: 'rgb(156, 163, 175)',
                    cursor: 'pointer',
                    padding: 0
                  }}
                  onMouseEnter={(e) => e.currentTarget.style.color = 'white'}
                  onMouseLeave={(e) => e.currentTarget.style.color = 'rgb(156, 163, 175)'}
                >
                  <svg style={{ width: '1.25rem', height: '1.25rem' }} fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
                  </svg>
                </motion.button>
                <motion.button
                  whileHover={{ scale: 1.1 }}
                  style={{
                    background: 'none',
                    border: 'none',
                    color: 'rgb(156, 163, 175)',
                    cursor: 'pointer',
                    padding: 0
                  }}
                  onMouseEnter={(e) => e.currentTarget.style.color = 'white'}
                  onMouseLeave={(e) => e.currentTarget.style.color = 'rgb(156, 163, 175)'}
                >
                  <svg style={{ width: '1.25rem', height: '1.25rem' }} fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                  </svg>
                </motion.button>
              </div>
            </div>

            {/* メッセージエリア */}
            <div style={{ flex: 1, overflow: 'hidden' }}>
              <MessageList roomId={selectedRoom.id} />
            </div>

            {/* メッセージ入力エリア */}
            <div style={{
              padding: '1rem',
              backgroundColor: 'rgb(31, 41, 55)'
            }}>
              <MessageForm roomId={selectedRoom.id} roomName={selectedRoom.name} />
            </div>
          </>
        ) : (
          <div style={{
            flex: 1,
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            padding: '1rem'
          }}>
            <motion.div
              initial={{ opacity: 0, scale: 0.9 }}
              animate={{ opacity: 1, scale: 1 }}
              transition={{ duration: 0.3 }}
              style={{ textAlign: 'center' }}
            >
              <div style={{
                marginBottom: '1rem',
                marginLeft: 'auto',
                marginRight: 'auto',
                width: '5rem',
                height: '5rem',
                backgroundColor: 'rgb(55, 65, 81)',
                borderRadius: '50%',
                display: 'flex',
                alignItems: 'center',
                justifyContent: 'center'
              }}>
                <svg style={{ width: '2.5rem', height: '2.5rem', color: 'rgb(107, 114, 128)' }} fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
                </svg>
              </div>
              <h3 style={{
                fontSize: '1.25rem',
                fontWeight: '600',
                color: 'rgb(209, 213, 219)',
                marginBottom: '0.5rem'
              }}>
                ルームを選択してください
              </h3>
              <p style={{ color: 'rgb(107, 114, 128)' }}>
                左側のリストからチャットルームを選択するか、
                <br />
                新しいルームを作成してください
              </p>
              <button
                onClick={() => setShowCreateRoom(true)}
                style={{
                  marginTop: '1rem',
                  padding: '0.5rem 1rem',
                  background: 'rgb(147, 51, 234)',
                  color: 'white',
                  borderRadius: '0.5rem',
                  border: 'none',
                  cursor: 'pointer',
                  transition: 'background-color 0.2s'
                }}
                onMouseEnter={(e) => e.currentTarget.style.backgroundColor = 'rgb(126, 34, 206)'}
                onMouseLeave={(e) => e.currentTarget.style.backgroundColor = 'rgb(147, 51, 234)'}
              >
                ルームを作成
              </button>
            </motion.div>
          </div>
        )}
      </div>
    </div>
  )
}