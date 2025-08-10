import React, { useState, useEffect, useMemo } from 'react'
import { motion, AnimatePresence } from 'framer-motion'
import { useListRooms, useRoomInfo } from '../../../modules/room'

export const RoomList = ({ onRoomSelect, selectedRoom }) => {
  const { data } = useListRooms()
  const [isOpen, setIsOpen] = useState(false)
  const [infoRoom, setInfoRoom] = useState(null)
  const { handleRoomInfo } = useRoomInfo()

  const handleRoomInfoClick = async (e, roomId) => {
    e.stopPropagation()
    const room = await handleRoomInfo(roomId)
    setInfoRoom(room)
    setIsOpen(true)
  }

  const rooms = useMemo(() => (data && data.rooms ? data.rooms : []), [data])

  useEffect(() => {
    if (rooms && rooms.length > 0 && !selectedRoom) {
      onRoomSelect(rooms[0])
    }
  }, [rooms, selectedRoom, onRoomSelect])

  if (!rooms || rooms.length === 0) {
    return (
      <div style={{ padding: '1rem', textAlign: 'center' }}>
        <p style={{ color: 'rgb(107, 114, 128)', fontSize: '0.875rem' }}>まだルームがありません</p>
        <p style={{ color: 'rgb(75, 85, 99)', fontSize: '0.75rem', marginTop: '0.25rem' }}>+ボタンから作成してください</p>
      </div>
    )
  }

  return (
    <>
      <div style={{ display: 'flex', flexDirection: 'column', gap: '0.25rem' }}>
        {rooms.map((room) => (
          <motion.button
            key={room.id}
            whileHover={{ x: 2 }}
            whileTap={{ scale: 0.98 }}
            onClick={() => onRoomSelect(room)}
            style={{
              width: '100%',
              textAlign: 'left',
              padding: '0.5rem 0.75rem',
              borderRadius: '0.5rem',
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'space-between',
              transition: 'all 0.2s',
              backgroundColor: selectedRoom?.id === room.id ? 'rgb(55, 65, 81)' : 'transparent',
              color: selectedRoom?.id === room.id ? 'white' : 'rgb(156, 163, 175)',
              border: 'none',
              cursor: 'pointer'
            }}
            onMouseEnter={(e) => {
              if (selectedRoom?.id !== room.id) {
                e.currentTarget.style.backgroundColor = 'rgba(55, 65, 81, 0.5)'
                e.currentTarget.style.color = 'white'
              }
            }}
            onMouseLeave={(e) => {
              if (selectedRoom?.id !== room.id) {
                e.currentTarget.style.backgroundColor = 'transparent'
                e.currentTarget.style.color = 'rgb(156, 163, 175)'
              }
            }}
          >
            <div style={{ display: 'flex', alignItems: 'center', gap: '0.5rem', minWidth: 0 }}>
              <svg style={{ height: '1rem', width: '1rem', flexShrink: 0 }} fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M7 20l4-16m2 16l4-16M6 9h14M4 15h14" />
              </svg>
              <span style={{ 
                overflow: 'hidden',
                textOverflow: 'ellipsis',
                whiteSpace: 'nowrap',
                fontSize: '0.875rem',
                fontWeight: '500'
              }}>{room.name}</span>
            </div>
            <motion.button
              whileHover={{ scale: 1.1 }}
              onClick={(e) => handleRoomInfoClick(e, room.id)}
              style={{
                opacity: 0,
                transition: 'opacity 0.2s',
                background: 'none',
                border: 'none',
                padding: 0,
                cursor: 'pointer'
              }}
              onMouseEnter={(e) => {
                e.currentTarget.style.opacity = 1
                e.currentTarget.parentElement.querySelector('svg:last-child').style.color = 'white'
              }}
              onMouseLeave={(e) => {
                e.currentTarget.style.opacity = 0
                e.currentTarget.parentElement.querySelector('svg:last-child').style.color = 'rgb(156, 163, 175)'
              }}
            >
              <svg style={{ height: '1rem', width: '1rem', color: 'rgb(156, 163, 175)' }} fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </motion.button>
          </motion.button>
        ))}
      </div>

      {/* モーダル */}
      <AnimatePresence>
        {isOpen && (
          <>
            <motion.div
              initial={{ opacity: 0 }}
              animate={{ opacity: 1 }}
              exit={{ opacity: 0 }}
              onClick={() => setIsOpen(false)}
              style={{
                position: 'fixed',
                inset: 0,
                backgroundColor: 'rgba(0, 0, 0, 0.6)',
                backdropFilter: 'blur(4px)',
                zIndex: 50
              }}
            />
            <motion.div
              initial={{ opacity: 0, scale: 0.9 }}
              animate={{ opacity: 1, scale: 1 }}
              exit={{ opacity: 0, scale: 0.9 }}
              style={{
                position: 'fixed',
                top: '50%',
                left: '50%',
                transform: 'translate(-50%, -50%)',
                zIndex: 50,
                width: '100%',
                maxWidth: '28rem'
              }}
            >
              <div className="glass-morphism" style={{
                borderRadius: '1rem',
                padding: '1.5rem',
                boxShadow: '0 20px 25px -5px rgba(0, 0, 0, 0.1)'
              }}>
                <div style={{
                  display: 'flex',
                  alignItems: 'center',
                  justifyContent: 'space-between',
                  marginBottom: '1rem'
                }}>
                  <h3 style={{
                    fontSize: '1.25rem',
                    fontWeight: 'bold',
                    color: 'white',
                    display: 'flex',
                    alignItems: 'center',
                    gap: '0.5rem'
                  }}>
                    <svg style={{ height: '1.25rem', width: '1.25rem' }} fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
                    </svg>
                    ルーム詳細
                  </h3>
                  <button
                    onClick={() => setIsOpen(false)}
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
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
                    </svg>
                  </button>
                </div>

                <div style={{ display: 'flex', flexDirection: 'column', gap: '1rem' }}>
                  <div>
                    <p style={{ color: 'rgb(156, 163, 175)', fontSize: '0.875rem', marginBottom: '0.25rem' }}>ルーム名</p>
                    <p style={{ color: 'white', fontWeight: '500' }}>{infoRoom?.name}</p>
                  </div>

                  <div>
                    <p style={{ color: 'rgb(156, 163, 175)', fontSize: '0.875rem', marginBottom: '0.25rem' }}>説明</p>
                    <p style={{ color: 'white' }}>
                      {infoRoom?.description || '説明はありません'}
                    </p>
                  </div>

                  <div>
                    <p style={{
                      color: 'rgb(156, 163, 175)',
                      fontSize: '0.875rem',
                      marginBottom: '0.25rem',
                      display: 'flex',
                      alignItems: 'center',
                      gap: '0.25rem'
                    }}>
                      <svg style={{ height: '1rem', width: '1rem' }} fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                      </svg>
                      作成日時
                    </p>
                    <p style={{ color: 'white' }}>
                      {infoRoom?.createdAt ? new Date(infoRoom.createdAt).toLocaleString('ja-JP') : '-'}
                    </p>
                  </div>
                </div>

                <button
                  onClick={() => setIsOpen(false)}
                  style={{
                    width: '100%',
                    marginTop: '1.5rem',
                    padding: '0.5rem 1rem',
                    background: 'linear-gradient(to right, rgb(147, 51, 234), rgb(236, 72, 153))',
                    color: 'white',
                    fontWeight: '500',
                    borderRadius: '0.5rem',
                    border: 'none',
                    cursor: 'pointer',
                    boxShadow: '0 10px 15px -3px rgba(0, 0, 0, 0.1)',
                    transform: 'scale(1)',
                    transition: 'all 0.2s'
                  }}
                  onMouseEnter={(e) => e.currentTarget.style.transform = 'scale(1.02)'}
                  onMouseLeave={(e) => e.currentTarget.style.transform = 'scale(1)'}
                >
                  閉じる
                </button>
              </div>
            </motion.div>
          </>
        )}
      </AnimatePresence>
    </>
  )
}