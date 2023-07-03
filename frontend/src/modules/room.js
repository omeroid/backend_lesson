import { useState } from 'react'
import useSWR from 'swr'

import { fetcher, mutater } from './fetcher'

export const useListRooms = () => {
  return useSWR('/rooms', fetcher)
}

export const useRoomInfo = () => {
  const handleRoomInfo = async (roomId) => {
    const room = await mutater(`/rooms/${roomId}`, 'GET', null)
    const datetime = new Date(room?.createdAt);
    const year = datetime.getFullYear();
    const month = datetime.getMonth() + 1;
    const day = datetime.getDate();
    room.createdAt = `${year}年${month}月${day}日`
    return room
  }


  return {
    handleRoomInfo,
  }
}

export const useCreateRoom = () => {
  const [isMutating, setIsMutating] = useState(false)
  const { mutate } = useListRooms()
  const handleCreateRoom = async (name,description) => {
    setIsMutating(true)
    const data = {
      name: name,
      description: description,
    }
    await mutater('/rooms','POST',data)
    mutate()
    setIsMutating(false)
  }

  
  return {
    handleCreateRoom,
    isMutating,
  }
}
