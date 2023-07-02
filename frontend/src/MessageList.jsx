import { useState, useEffect } from 'react';
import axios from 'axios';
import {useNavigate} from 'react-router-dom'

import { MessageList as ChatMessageList } from "react-chat-elements"

import { fetcher } from './fetcher'


export const MessageList = ({ selectedRoomId, setSelectedRoomId,allReload,setAllReload }) => {
  const [messages, setMessages] = useState([]);

  const rawUserData = sessionStorage.getItem("userData");
  const user = rawUserData ? JSON.parse(rawUserData) : null;

  const navigate = useNavigate();
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
      setAllReload(true)
    }catch(e){
      console.log("failure to delete room",e)
      if(e?.requst?.status === 401){
        navigate("/")
      }
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
          let rawMessages = [];
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
        } catch (e) {
          console.error('Error occurred while fetching messages:', e);
          if(e?.requst?.status === 401){
            navigate("/")
          }
        }
      }
    };
    if(allReload){
      setAllReload(false)
    }
    fetchData();
  }, [selectedRoomId, allReload, navigate, setAllReload, user.token, user.userId]);

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