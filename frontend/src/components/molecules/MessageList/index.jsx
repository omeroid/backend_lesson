import { useState, useEffect } from 'react';
import { MessageList as ChatMessageList } from "react-chat-elements"

import { useListMessages, useDeleteMessage } from '../../../modules/message'
import { useUser } from '../../../modules/user'


export const MessageList = ({ roomId }) => {
  const [messages, setMessages] = useState([]);
  const { data } = useListMessages(roomId)
  const { handleDeleteMessage } = useDeleteMessage(roomId)
  console.log("MessageList:",roomId)
  const { user } = useUser()


  useEffect(() => {
    if (!data || !data.messages) {
      setMessages([])
      return
    }
    if(!user) return
    const list = data.messages.map(item => ({
      id: item.id,
      position: user.userId === item.user.id ? "left" : "right",
      type: "text",

      title: item.user.name,
      text: item.text,
      removeButton: user.userId === item.user.id,
    }));
    setMessages(list)
  }, [data, user]);

  return (
    <div>
      <ChatMessageList
        onRemoveMessageClick={(message) => handleDeleteMessage(message.id)}
        className='message-list'
        lockable={true}
        toBottomHeight={'100%'}
        dataSource={messages}
      />
    </div>
  );
};