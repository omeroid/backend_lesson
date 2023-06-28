import * as React from 'react';
import useSWR from 'swr'
import { useState,useEffect } from 'react';

import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import Button from '@mui/material/Button';
import IconButton from '@mui/material/IconButton';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import AddCircleIcon from '@mui/icons-material/AddCircle';
import SendIcon from '@mui/icons-material/Send';
import TextField from '@mui/material/TextField';

import { fetcher } from './fetcher'

export const SendMessage = ({selectedRoomId,setSelectedRoomId}) => {
  const [chatInput, setChatInput] = useState(''); const rawUserData = sessionStorage.getItem("userData");
  const user = rawUserData ? JSON.parse(rawUserData) : null;

  const handleChatInputChange = (event) => {
    setChatInput(event.target.value);
  };

  const handleChatSubmit = async (event) => {
    console.log("in submit")
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
    }
    );
    setChatInput('');
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
