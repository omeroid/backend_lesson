

import { useState } from 'react';
import useSWR from 'swr'

import { MessageList as ChatMessageList } from "react-chat-elements"

import { fetcher } from './fetcher'

export const MessageList = () => {
  // const { data: rooms, error, isLoading } = useSWR('/rooms', fetcher)
  const [ messages, setMessages ] = useState([
    {
      position:"left",
      type:"text",
      title:"Kursat",
      text:"Give me a message list example !",
      avatar: "https://avatars.githubusercontent.com/u/80540635?v=4"
    },
    {
      position:"right",
      type:"text",
      title:"Emre",
      text:"That's all.",
      avatar: "https://avatars.githubusercontent.com/u/41473129?v=4"
    }
])

  return (
    <ChatMessageList
      className='message-list'
      lockable={true}
      toBottomHeight={'100%'}
      dataSource={messages}
    />
  );
}