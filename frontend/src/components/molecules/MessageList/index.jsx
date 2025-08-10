import { useState, useEffect, useRef, useCallback } from 'react'
import { Box } from '@mui/material'
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
    <Box sx={{ overflow: 'auto', height: 'calc(100% - 5rem)', py: 2 }}>
      {messages.map((message) => (
        <MessageItem
          key={message.id}
          message={message}
          onDelete={(msg) => handleDeleteMessage(msg.id)}
        />
      ))}
      <div ref={scrollRef}></div>
    </Box>
  )
}