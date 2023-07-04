import React from 'react';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';

import { MessageList } from '../../molecules/MessageList';
import { MessageForm } from '../../molecules/MessageForm';


export const MessagePanel = ({roomId})  => {

  return (
    <Box
      component='main'
      sx={{
        backgroundColor: (theme) =>
          theme.palette.mode === 'light'
            ? theme.palette.grey[100]
            : theme.palette.grey[900],
        flexGrow: 1,
        height: '100vh',
        overflow: 'auto',
        margin: 'auto',
        display: 'flex',
        flexDirection: 'column',
      }}
    >
      <Toolbar />
      <MessageList
        roomId={roomId}
      />
      <MessageForm
        roomId={roomId}
      />
    </Box>
  );
}
