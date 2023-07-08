import { useState, useEffect, useRef, useCallback } from 'react'
import { MessageList as ChatMessageList } from 'react-chat-elements'

import { useListMessages, useDeleteMessage } from '../../../modules/message'
import { useUser } from '../../../modules/user'

export const MessageList = ({ roomId }) => {
  const [messages, setMessages] = useState([])
  const { data } = useListMessages(roomId)
  const { handleDeleteMessage } = useDeleteMessage(roomId)
  const { user } = useUser()

  const scrollRef = useRef(null)
  const scrollToBottomOfList = useCallback(() => {
    console.log('scrollToBottomOfList', scrollRef)
    scrollRef.current.scrollIntoView({
      behavior: 'smooth',
      block: 'end',
    })
  }, [scrollRef])

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
  }, [data, user, scrollToBottomOfList])

  useEffect(() => {
    scrollToBottomOfList()
  }, [messages, scrollToBottomOfList])

  return (
    <div style={{ overflow: 'scroll', height: 'calc(100% - 5rem)' }}>
      <ChatMessageList
        onRemoveMessageClick={(message) => handleDeleteMessage(message.id)}
        className="message-list"
        lockable={true}
        toBottomHeight={'100%'}
        dataSource={messages}
      />
      <div ref={scrollRef}></div>
    </div>
  )
}
