import { useState } from 'react'
import useSWR from 'swr'

import { fetcher, mutater } from './fetcher'
import { useUser } from './user'

export const useListMessages = (roomId) => {
  return useSWR(roomId ? `/rooms/${roomId}/messages` : null, fetcher)
}

export const useCreateMessage = (roomId) => {
  const { user } = useUser()
  const [isMutating, setIsMutating] = useState(false)
  const { mutate } = useListMessages(roomId)

  const handleCreateMessage = async (text) => {
    setIsMutating(true)
    const data = {
      userId: user.userId,
      text: text,
    }
    await mutater(`/rooms/${roomId}/messages`, 'POST', data)
    mutate()
    setIsMutating(false)
  }

  return {
    handleCreateMessage,
    isMutating,
  }
}

export const useDeleteMessage = (roomId) => {
  const [isMutating, setIsMutating] = useState(false)
  const { mutate } = useListMessages(roomId)

  const handleDeleteMessage = async (messageId) => {
    setIsMutating(true)
    await mutater(`/rooms/${roomId}/messages/${messageId}`, 'DELETE', null)
    mutate()
    setIsMutating(false)
  }

  return {
    handleDeleteMessage,
    isMutating,
  }
}
