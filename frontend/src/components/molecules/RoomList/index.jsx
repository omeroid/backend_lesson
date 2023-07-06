import { useState, useEffect, useMemo } from 'react';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemText from '@mui/material/ListItemText';
import IconButton from '@mui/material/IconButton';
import Box from '@mui/material/Box';
import Divider from '@mui/material/Divider';
import InfoIcon from '@mui/icons-material/Info';
import Modal from '@mui/material/Modal';
import Typography from '@mui/material/Typography';

import { useListRooms, useRoomInfo } from '../../../modules/room'

export const RoomList = ({ roomId, setRoomId }) => {
  const { data } = useListRooms()
  const [isOpen, setIsOpen] = useState(false);
  const [infoRoom, setInfoRoom] = useState(null);
  const { handleRoomInfo } = useRoomInfo()

  const handleRoomInfoClick = async (roomId) => {
    const room = await handleRoomInfo(roomId)
    setInfoRoom(room)
    setIsOpen(true)
  }

  const rooms = useMemo(() => data && data.rooms ? data.rooms : [], [data])

  useEffect(() => {
    if(rooms && rooms.length > 0 && roomId === null) {
      setRoomId(rooms[0].id);
    }
  }, [rooms, roomId, setRoomId])

  return (
    <div>
      {rooms &&
        rooms.map((r) => roomId === r.id ?
          (
            <div key={r.id}>
              <Box
                display='flex'
                alignItems='center'
                p={1}
                backgroundColor='primary.light'
                color='white'
                >
                <Box flexGrow={1}>
                  <ListItemButton onClick={() => setRoomId(r.id)}>
                    <ListItemText primary={r.name} style={{ textAlign: 'center' }} />
                  </ListItemButton>
                </Box>
                <IconButton onClick={() => handleRoomInfoClick(r.id)}>
                  <InfoIcon sx={{ color: 'white' }}/>
                </IconButton>
              </Box>
              <Divider />
            </div>
          ) :
          (
            <div key={r.id}>
              <Box display='flex' alignItems='center' p={1}>
                <Box flexGrow={1}>
                  <ListItemButton onClick={() => setRoomId(r.id)}>
                    <ListItemText primary={r.name} style={{ textAlign: 'center' }} />
                  </ListItemButton>
                </Box>
                <IconButton onClick={() => handleRoomInfoClick(r.id)}>
                  <InfoIcon />
                </IconButton>
              </Box>
              <Divider />
            </div>
          )
        )}
        <Modal open={isOpen} onClose={() => setIsOpen(false)} aria-labelledby='modal-title'>
          <Box
            sx={{
              position: 'absolute',
              top: '50%',
              left: '50%',
              transform: 'translate(-50%, -50%)',
              bgcolor: 'white',
              p: 4,
              outline: 'none',
              borderRadius: 8,
              boxShadow: '0px 2px 10px rgba(0, 0, 0, 0.15)',
              maxWidth: 400,
              width: '100%',
            }}
          >
            <Typography variant='h6' id='modal-title' sx={{ marginBottom: 2 }}>
              ルーム詳細
            </Typography>
            <Box sx={{ marginBottom: 2 }}>
              <Typography variant='body1' component='div' sx={{ fontWeight: 'bold', marginBottom: 1 }}>
                Name:
              </Typography>
              <Typography variant='body1' component='div' sx={{ marginBottom: '1rem' }}>
                {infoRoom?.name}
              </Typography>
            </Box>
            <Box sx={{ marginBottom: 2 }}>
              <Typography variant='body1' component='div' sx={{ fontWeight: 'bold', marginBottom: 1 }}>
                Description:
              </Typography>
              <Typography variant='body1' component='div' sx={{ marginBottom: '1rem' }}>
                {infoRoom?.description}
              </Typography>
            <Box sx={{ marginBottom: 2 }}>
              <Typography variant='body1' component='div' sx={{ fontWeight: 'bold', marginBottom: 1 }}>
                CreatedAt:
              </Typography>
              <Typography variant='body1' component='div' sx={{ marginBottom: '1rem' }}>
                {infoRoom?.createdAt}
              </Typography>
            </Box>
          </Box>
        </Box>
      </Modal>
    </div>
  );
};