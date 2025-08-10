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
  const scrollToBottomOfList = useCallback(() => {
    scrollRef.current?.scrollIntoView({
      behavior: 'smooth',
      block: 'end',
    })
  }, [])

  useEffect(() => {
    if (!data || !data.messages) {
      setMessages([])
      return
    }
    if (!user) return
    const list = data.messages.map((item) => ({
      id: item.id,
      position: user.userId === item.user.id ? 'right' : 'left',
      type: 'text',
      title: item.user.name,
      text: item.text,
      removeButton: user.userId === item.user.id,
      className: user.userId === item.user.id ? 'my-message' : '',
    }))
    setMessages(list)
  }, [data, user])

  useEffect(() => {
    scrollToBottomOfList()
  }, [messages, scrollToBottomOfList])

  return (
    <div className="flex-1 overflow-y-auto scrollbar-thin bg-gray-800 py-4">
      {messages.length === 0 ? (
        <motion.div
          initial={{ opacity: 0 }}
          animate={{ opacity: 1 }}
          className="flex items-center justify-center h-full"
        >
          <div className="text-center">
            <p className="text-gray-500 text-lg mb-2">まだメッセージがありません</p>
            <p className="text-gray-600 text-sm">最初のメッセージを送信してみましょう！</p>
          </div>
        </motion.div>
      ) : (
        <>
          {messages.map((message, index) => (
            <MessageItem
              key={message.id}
              message={message}
              onDelete={(msg) => handleDeleteMessage(msg.id)}
            />
          ))}
        </>
      )}
      <div ref={scrollRef}></div>
    </div>
  )
}