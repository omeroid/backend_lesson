import { useState } from 'react'
import useSWR from 'swr'

import { fetcher, mutater } from './fetcher'
import { useUser } from './user'

export const useListRooms = () => {
  return useSWR('/rooms', fetcher)
}
// export const useCreateMessage = (roomId) => {
//   const { user } = useUser()
//   const [isMutating, setIsMutating] = useState(false)
//   const { mutate } = useListMessages(roomId)

//   const handleCreateMessage = async (text) => {
//     setIsMutating(true)
//     const data = {
//       userId: user.userId,
//       text: text
//     }
//     await mutater(`/rooms/${roomId}/messages`, 'POST', data)
//     mutate()
//     setIsMutating(false)
//   }


//   return {
//     handleCreateMessage,
//     isMutating
//   }

// }

