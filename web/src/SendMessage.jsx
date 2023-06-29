import * as React from 'react';
import { useState} from 'react';
import {useNavigate} from 'react-router-dom'

import IconButton from '@mui/material/IconButton';
import Box from '@mui/material/Box';
import SendIcon from '@mui/icons-material/Send';
import TextField from '@mui/material/TextField';

import { fetcher } from './fetcher'

export const SendMessage = ({selectedRoomId,setSelectedRoomId,setAllReload}) => {
  const [chatInput, setChatInput] = useState('');
  const rawUserData = sessionStorage.getItem("userData");
  const user = rawUserData ? JSON.parse(rawUserData) : null;
  const navigate = useNavigate();

  const handleChatInputChange = (event) => {
    setChatInput(event.target.value);
  };

  const handleChatSubmit = async (event) => {
    try{
      event.preventDefault();
      const data = new FormData(event.currentTarget);
      const text = data.get("text")
      console.log(text,user.token,user.userId)
      const response = await fetcher(`http://localhost:1323/rooms/${selectedRoomId}/messages`, {
        method: "POST",
        headers: {
          'Content-Type': 'application/json',
          Authorization: 'Bearer ' + user.token,
        
        },
        data: {
          userId:user.userId,
          text:text
        }
      });
      console.log("success to send message",response)
      setChatInput('');
      setSelectedRoomId(selectedRoomId);
      setAllReload(true);
    }catch(error){
      console.log("failure to send message",error)
      if(error?.requst?.status === 401){
        navigate("/")
      }
    }
  };

  return (
    <Box 
    sx={{ mt: 'auto', p: 2, backgroundColor: 'white' }} 
    component="form"
    onSubmit={handleChatSubmit}>
      <Box sx={{ display: 'flex', alignItems: 'center' }}>
        <TextField
          id="text"
          name="text"
          label="チャットメッセージ"
          value={chatInput}
          onChange={handleChatInputChange}
          fullWidth
        />
        <IconButton color="primary" type="submit">
          <SendIcon />
        </IconButton>
      </Box>
    </Box>
  );
}
