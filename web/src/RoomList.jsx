
import useSWR from 'swr'
import { useState } from 'react';

import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import HomeIcon from '@mui/icons-material/Home';

import { fetcher } from './fetcher'

export const RoomList = () => {
  // const { data: rooms, error, isLoading } = useSWR('/rooms', fetcher)
  const [ rooms, setRooms ] = useState([
    {
      id:"1",
      name:"ゼルダプレイルーム",
      description:"ゼルダのティアキンについて語る部屋",
    },
    {
      id:"2",
      name:"omeroidに就職したい人のための部屋",
      description:"入るといいことあるかも。。",
    },
])

  return (
    <>
      {rooms && rooms.map(r => (
        <ListItemButton>
          <ListItemIcon>
            <HomeIcon />
          </ListItemIcon>
          <ListItemText primary={r.name} />
        </ListItemButton>
      ))}
    </>
  );
}