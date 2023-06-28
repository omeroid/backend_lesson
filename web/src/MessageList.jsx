import { useState, useEffect } from 'react';
import useSWR from 'swr'
import axios from 'axios';

import { MessageList as ChatMessageList } from "react-chat-elements"
import { Button, Card } from 'react-bootstrap';
import { fetcher } from './fetcher'


export const MessageList = ({ selectedRoomId, setSelectedRoomId }) => {
  const [messages, setMessages] = useState([]);

  const rawUserData = sessionStorage.getItem("userData");
  const user = rawUserData ? JSON.parse(rawUserData) : null;

  const handleRemoveMessage = async (message) => {
    console.log(message)
    try{
      const response = await axios(`http://localhost:1323/rooms/${selectedRoomId}/messages/${message.id}`,{
        method:"GET",
        headers:{
          'Content-Type': 'application/json',
          Authorization: 'Bearer '+user.token,
        }
      })
      console.log("success to delete room",response)
      setSelectedRoomId(selectedRoomId)
    }catch(e){
      console.log("failure to delete room",e)
    }
  }

  useEffect(() => {
    const fetchData = async () => {
      if (selectedRoomId !== null) {
        try {
          const response = await fetcher(`http://localhost:1323/rooms/${selectedRoomId}/messages`, {
            method: "GET",
            headers: {
              'Content-Type': 'application/json',
              Authorization: 'Bearer ' + user.token,
            }
          });
          let rawMessages = [
            {
              position: "left",
              type: "text",
              title: "Kursat",
              text: "Give me a message list example!",
              removeButton: false,
            },
            {
              position: "right",
              type: "text",
              title: "Emre",
              text: "That's all.",
              removeButton: false,
            }
          ];
          if (response && Array.isArray(response.messages)) {
            rawMessages = response.messages.map(item => ({
              id: item.id,
              position: user.userId === item.user.id ? "left" : "right",
              type: "text",
              title: item.user.name,
              text: item.text,
              removeButton: user.userId === item.user.id,
            }));
          }
          setMessages(rawMessages);
        } catch (error) {
          console.error('Error occurred while fetching messages:', error);
        }
      }
    };
    fetchData();
  }, [selectedRoomId]);

  return (
    <div>
      <ChatMessageList
        onRemoveMessageClick={handleRemoveMessage}
        className='message-list'
        lockable={true}
        toBottomHeight={'100%'}
        dataSource={messages}
      />
    </div>
  );
};