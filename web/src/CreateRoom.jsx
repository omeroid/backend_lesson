import * as React from 'react';
import axios from 'axios';
import { useState } from 'react';
import Button from '@mui/material/Button';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import AddCircleIcon from '@mui/icons-material/AddCircle';
import Modal from '@mui/material/Modal';
import TextField from '@mui/material/TextField';

export const CreateRoom = () => {
  const [isOpen, setIsOpen] = useState(false);

  const handleSubmit = async(event) => {
    event.preventDefault();
    const data = new FormData(event.currentTarget);
    const rawUserData = sessionStorage.getItem("userData");
    const user = rawUserData ? JSON.parse(rawUserData):null;
    const name = data.get('name')
    const description = data.get('description')
    try{
      const response = await axios('http://localhost:1323/rooms',{
        method:"POST",
        data:{
          name: name,
          description: description,
        },
        headers:{
          'Content-Type': 'application/json',
          Authorization: 'Bearer '+user.token,
        }
      })
      // TODO エラーの場合はエラーをフロントに表示
      console.log("success to create room",response)
    }catch(e){
      console.log("failure to create room",e)
    }
    setIsOpen(false);
  };

  return (
    <div>
      <Button
        variant="contained"
        size="large"
        startIcon={<AddCircleIcon />}
        sx={{
          color: 'white',
          width: '100%',
          borderRadius: 0,
        }}
        onClick={() => setIsOpen(true)}
      >
        <Box display="flex" alignItems="center" width="100%">
          <Typography variant="body1" sx={{ marginLeft: 8 }}>
            新規ルーム作成
          </Typography>
        </Box>
      </Button>

      <Modal
        open={isOpen}
        onClose={() => setIsOpen(false)}
        aria-labelledby="modal-title"
      >
        <Box component="form"
          onSubmit={handleSubmit}
          noValidate
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
            id = "name"
            label="name"
            name="name"
            autoComplete = "name"
            fullWidth
            sx={{ marginBottom: 2 }}
          />
          <TextField
            required
            id = "description"
            label="description"
            name="description"
            autoComplete = "description"
            fullWidth
            sx={{ marginBottom: 2 }}
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            sx={{ mt: 3, mb: 2 }}
          >作成</Button>
        </Box>
      </Modal>
    </div>
  );
};
