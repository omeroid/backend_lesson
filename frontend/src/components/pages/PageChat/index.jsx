import React, { useState } from 'react'
import { styled } from '@mui/material/styles'
import Box from '@mui/material/Box'
import MuiAppBar from '@mui/material/AppBar'
import Toolbar from '@mui/material/Toolbar'
import Typography from '@mui/material/Typography'

import { RoomPanel } from '../../organisms/RoomPanel'
import { MessagePanel } from '../../organisms/MessagePanel'
import { UserIcon } from '../../molecules/UserIcon'

const drawerWidth = 500

const AppBar = styled(MuiAppBar, {
  shouldForwardProp: (prop) => prop !== 'open',
})(({ theme, open }) => ({
  zIndex: theme.zIndex.drawer + 1,
  transition: theme.transitions.create(['width', 'margin'], {
    easing: theme.transitions.easing.sharp,
    duration: theme.transitions.duration.leavingScreen,
  }),
  ...(open && {
    marginLeft: drawerWidth,
    width: `calc(100% - ${drawerWidth}px)`,
    transition: theme.transitions.create(['width', 'margin'], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  }),
}))

export const PageChat = () => {
  const [roomId, setRoomId] = useState(null)

  return (
    <Box sx={{ display: 'flex' }}>
      <AppBar position="absolute" open={true}>
        <Toolbar
          sx={{
            pr: '24px', // keep right padding when drawer closed
          }}
        >
          <Typography
            component="h1"
            variant="h6"
            color="inherit"
            noWrap
            sx={{ flexGrow: 1 }}
          >
            omeroidChatApp
          </Typography>
          <UserIcon />
        </Toolbar>
      </AppBar>
      <RoomPanel roomId={roomId} setRoomId={setRoomId} />
      <MessagePanel roomId={roomId} />
    </Box>
  )
}
