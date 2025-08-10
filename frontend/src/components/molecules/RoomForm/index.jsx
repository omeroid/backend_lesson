import * as React from 'react'
import { useState } from 'react'
import { motion } from 'framer-motion'
import { useCreateRoom } from '../../../modules/room'
import toast from 'react-hot-toast'

export const RoomForm = ({ onClose }) => {
  const [nameInput, setNameInput] = useState('')
  const [descriptionInput, setDescriptionInput] = useState('')
  const [loading, setLoading] = useState(false)

  const { handleCreateRoom } = useCreateRoom()

  const handleFormSubmit = async (e) => {
    e.preventDefault()
    if (!nameInput.trim()) {
      toast.error('ルーム名を入力してください')
      return
    }
    
    setLoading(true)
    try {
      await handleCreateRoom(nameInput, descriptionInput)
      toast.success('ルームを作成しました！')
      onClose()
    } catch (error) {
      toast.error('ルームの作成に失敗しました')
    } finally {
      setLoading(false)
    }
  }

  return (
    <motion.div
      initial={{ opacity: 0, y: -10 }}
      animate={{ opacity: 1, y: 0 }}
      exit={{ opacity: 0, y: -10 }}
      style={{
        backgroundColor: 'rgb(31, 41, 55)',
        borderRadius: '0.5rem',
        padding: '1rem'
      }}
    >
      <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between', marginBottom: '1rem' }}>
        <h3 style={{ color: 'white', fontWeight: '600', display: 'flex', alignItems: 'center', gap: '0.5rem' }}>
          <svg style={{ height: '1.25rem', width: '1.25rem' }} fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 4v16m8-8H4" />
          </svg>
          新規ルーム作成
        </h3>
        <button
          onClick={onClose}
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

      <form onSubmit={handleFormSubmit} style={{ display: 'flex', flexDirection: 'column', gap: '1rem' }}>
        <div>
          <label style={{ display: 'block', fontSize: '0.875rem', fontWeight: '500', color: 'rgb(209, 213, 219)', marginBottom: '0.25rem' }}>
            ルーム名 <span style={{ color: 'rgb(248, 113, 113)' }}>*</span>
          </label>
          <div style={{ position: 'relative' }}>
            <div style={{ position: 'absolute', top: '50%', transform: 'translateY(-50%)', left: '0.75rem', pointerEvents: 'none' }}>
              <svg style={{ height: '1rem', width: '1rem', color: 'rgb(107, 114, 128)' }} fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M7 20l4-16m2 16l4-16M6 9h14M4 15h14" />
              </svg>
            </div>
            <input
              type="text"
              value={nameInput}
              onChange={(e) => setNameInput(e.target.value)}
              style={{
                display: 'block',
                width: '100%',
                paddingLeft: '2.25rem',
                paddingRight: '0.75rem',
                paddingTop: '0.5rem',
                paddingBottom: '0.5rem',
                backgroundColor: 'rgb(55, 65, 81)',
                border: '1px solid rgb(75, 85, 99)',
                borderRadius: '0.5rem',
                color: 'white',
                outline: 'none',
                transition: 'all 0.2s'
              }}
              placeholder="general"
              required
            />
          </div>
        </div>

        <div>
          <label style={{ display: 'block', fontSize: '0.875rem', fontWeight: '500', color: 'rgb(209, 213, 219)', marginBottom: '0.25rem' }}>
            説明
          </label>
          <div style={{ position: 'relative' }}>
            <div style={{ position: 'absolute', top: '0.5rem', left: '0.75rem', pointerEvents: 'none' }}>
              <svg style={{ height: '1rem', width: '1rem', color: 'rgb(107, 114, 128)' }} fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
              </svg>
            </div>
            <textarea
              value={descriptionInput}
              onChange={(e) => setDescriptionInput(e.target.value)}
              style={{
                display: 'block',
                width: '100%',
                paddingLeft: '2.25rem',
                paddingRight: '0.75rem',
                paddingTop: '0.5rem',
                paddingBottom: '0.5rem',
                backgroundColor: 'rgb(55, 65, 81)',
                border: '1px solid rgb(75, 85, 99)',
                borderRadius: '0.5rem',
                color: 'white',
                outline: 'none',
                transition: 'all 0.2s',
                resize: 'none'
              }}
              placeholder="このルームの説明を入力..."
              rows="3"
            />
          </div>
        </div>

        <div style={{ display: 'flex', gap: '0.5rem' }}>
          <button
            type="button"
            onClick={onClose}
            style={{
              flex: 1,
              padding: '0.5rem 1rem',
              backgroundColor: 'rgb(55, 65, 81)',
              color: 'rgb(209, 213, 219)',
              borderRadius: '0.5rem',
              border: 'none',
              cursor: 'pointer',
              transition: 'background-color 0.2s'
            }}
            onMouseEnter={(e) => e.currentTarget.style.backgroundColor = 'rgb(75, 85, 99)'}
            onMouseLeave={(e) => e.currentTarget.style.backgroundColor = 'rgb(55, 65, 81)'}
          >
            キャンセル
          </button>
          <button
            type="submit"
            disabled={loading || !nameInput.trim()}
            style={{
              flex: 1,
              padding: '0.5rem 1rem',
              background: (loading || !nameInput.trim()) ? 'rgb(107, 114, 128)' : 'linear-gradient(to right, rgb(147, 51, 234), rgb(236, 72, 153))',
              color: 'white',
              fontWeight: '500',
              borderRadius: '0.5rem',
              border: 'none',
              cursor: (loading || !nameInput.trim()) ? 'not-allowed' : 'pointer',
              opacity: (loading || !nameInput.trim()) ? 0.5 : 1,
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'center',
              gap: '0.5rem',
              transition: 'all 0.2s',
              transform: 'scale(1)'
            }}
            onMouseEnter={(e) => !(loading || !nameInput.trim()) && (e.currentTarget.style.transform = 'scale(1.02)')}
            onMouseLeave={(e) => !(loading || !nameInput.trim()) && (e.currentTarget.style.transform = 'scale(1)')}
          >
            {loading ? (
              <div style={{
                width: '1rem',
                height: '1rem',
                border: '2px solid white',
                borderTopColor: 'transparent',
                borderRadius: '50%',
                animation: 'spin 1s linear infinite'
              }} />
            ) : (
              <>
                <svg style={{ height: '1rem', width: '1rem' }} fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 4v16m8-8H4" />
                </svg>
                作成
              </>
            )}
          </button>
        </div>
      </form>
    </motion.div>
  )
}