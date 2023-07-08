import * as React from 'react'
import { useState } from 'react'
import Button from '@mui/material/Button'
import Box from '@mui/material/Box'
import Typography from '@mui/material/Typography'
import AddCircleIcon from '@mui/icons-material/AddCircle'
import Modal from '@mui/material/Modal'
import TextField from '@mui/material/TextField'
import { useCreateRoom } from '../../../modules/room'

export const RoomForm = () => {
  const [isOpen, setIsOpen] = useState(false)
  const [nameInput, setNameInput] = useState('')
  const [descriptionInput, setDescriptionInput] = useState('')

  const { handleCreateRoom } = useCreateRoom()

  const handleFormSubmit = async () => {
    await handleCreateRoom(nameInput, descriptionInput)
  }

  return (
    <div>
      <Button
        variant="contained"
        size="large"
        endIcon={<AddCircleIcon />}
        sx={{
          color: 'white',
          width: '100%',
          borderRadius: 0,
          backgroundColor: 'secondary.main',
          margin: 'auto',
        }}
        onClick={() => setIsOpen(true)}
      >
        <Box display="flex" alignItems="center" width="100%">
          <Typography variant="body1" sx={{ width: '100%' }}>
            新規ルーム作成
          </Typography>
        </Box>
      </Button>

      <Modal
        open={isOpen}
        onClose={() => setIsOpen(false)}
        aria-labelledby="modal-title"
      >
        <Box
          component="form"
          onSubmit={handleFormSubmit}
          sx={{
            position: 'absolute',
            top: '50%',
            left: '50%',
            transform: 'translate(-50%, -50%)',
            bgcolor: 'white',
            p: 4,
            outline: 'none',
            borderRadius: 8,
          }}
        >
          <Typography variant="h6" id="modal-title" sx={{ marginBottom: 2 }}>
            ルーム作成
          </Typography>
          <TextField
            required
            id="name"
            label="name"
            name="name"
            value={nameInput}
            onChange={(e) => setNameInput(e.target.value)}
            autoComplete="name"
            fullWidth
            sx={{ marginBottom: 2 }}
          />
          <TextField
            required
            id="description"
            label="description"
            name="description"
            value={descriptionInput}
            onChange={(e) => setDescriptionInput(e.target.value)}
            autoComplete="description"
            fullWidth
            sx={{ marginBottom: 2 }}
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            sx={{ mt: 3, mb: 2 }}
          >
            作成
          </Button>
        </Box>
      </Modal>
    </div>
  )
}
