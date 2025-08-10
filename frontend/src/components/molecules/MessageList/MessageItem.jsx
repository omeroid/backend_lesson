import React from 'react'
import { Box, Typography, IconButton, Paper } from '@mui/material'
import DeleteIcon from '@mui/icons-material/Delete'

export const MessageItem = ({ message, onDelete }) => {
  const isMyMessage = message.position === 'right'
  
  return (
    <Box
      sx={{
        display: 'flex',
        justifyContent: isMyMessage ? 'flex-end' : 'flex-start',
        mb: 2,
        px: 2,
      }}
    >
      <Paper
        elevation={1}
        sx={{
          maxWidth: '70%',
          px: 2,
          py: 1,
          backgroundColor: isMyMessage ? '#e3f2fd' : '#f5f5f5',
          borderRadius: 2,
        }}
      >
        <Typography variant="caption" color="text.secondary" display="block">
          {message.title}
        </Typography>
        <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
          <Typography variant="body1">{message.text}</Typography>
          {message.removeButton && (
            <IconButton
              size="small"
              onClick={() => onDelete(message)}
              sx={{ ml: 1 }}
            >
              <DeleteIcon fontSize="small" />
            </IconButton>
          )}
        </Box>
      </Paper>
    </Box>
  )
}