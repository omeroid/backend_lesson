import { useState } from 'react';
import useSWR from 'swr'

import { MessageList as ChatMessageList } from "react-chat-elements"
import { Button, Card } from 'react-bootstrap';
import { fetcher } from './fetcher'

export const MessageList = () => {
  // const { data: rooms, error, isLoading } = useSWR('/rooms', fetcher);
  const [messages, setMessages] = useState([
    {
      position: "left",
      type: "text",
      title: "Kursat",
      text: "Give me a message list example!",
      removeButton:true,
    },
    {
      position: "right",
      type: "text",
      title: "Emre",
      text: "That's all.",
    }
  ]);

  return (
    <div>
      <ChatMessageList
        className='message-list'
        lockable={true}
        toBottomHeight={'100%'}
        dataSource={messages}
      />
    </div>
  );
};