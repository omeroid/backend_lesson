import * as React from 'react';
import { useState} from 'react';

import IconButton from '@mui/material/IconButton';
import Box from '@mui/material/Box';
import SendIcon from '@mui/icons-material/Send';
import TextField from '@mui/material/TextField';

import { useCreateMessage } from '../../../modules/message';

export const MessageForm = ({ roomId }) => {
  const { isMutating, handleCreateMessage } = useCreateMessage(roomId)

  const [chatInput, setChatInput] = useState('')

  const handleChatSubmit = async (e) => {
    e.preventDefault();
    await handleCreateMessage(chatInput)
    setChatInput('');
  };

  return (
    <Box
    sx={{ mt: 'auto', p: 2, backgroundColor: 'white' }}
    component="form"
    onSubmit={handleChatSubmit}>
     {roomId && (
        <Box sx={{ display: 'flex', alignItems: 'center' }}>
          <TextField
            id="text"
            name="text"
            label="チャットメッセージ"
            value={chatInput}
            onChange={(e) => setChatInput(e.target.value)}
            fullWidth
          />
          <IconButton color='primary' type='submit' disabled={isMutating}>
            <SendIcon />
          </IconButton>
        </Box>
      )}
    </Box>
  );
}
