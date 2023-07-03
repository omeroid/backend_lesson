import * as React from 'react';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import Link from '@mui/material/Link';
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import axios from 'axios'
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';


export default function SignUp() {
  const navigate = useNavigate()
  const [error, setError] = useState(null);

  const handleSubmit = async (event) => {
    const url = 'http://localhost:1323/user/signup'

    event.preventDefault();
    const data = new FormData(event.currentTarget);
    const username = data.get('username');
    const password = data.get('password');
    var response
    try{
        response = await axios.post(url,{
        userName: username,
        password: password,
      })
      navigate("/")
    }catch(e){
      setError('そのユーザ名はすでに使用されています。');
      response = e?.response
    }
    console.log("method:",response?.config?.method,"url:",response?.config?.url)
  };

  return (
    <Container component="main" maxWidth="xs">
      <Box
        sx={{
          marginTop: 8,
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
        }}
      >
      <Box
          component="img"
          sx={{height: 100,width: 100}}
          alt="omeroid icon"
          src="https://assets.st-note.com/production/uploads/images/38911312/profile_5e2d06172918f8d8fae54589aa5e2217.jpg"
        />
        <Typography component="h1" variant="h5">アカウント作成</Typography>
        <Box component="form" onSubmit={handleSubmit} noValidate sx={{ mt: 1 }}>
          <TextField
            margin="normal"
            required
            fullWidth
            id="username"
            label="username"
            name="username"
            autoComplete="email"
            autoFocus
          />
          <TextField
            margin="normal"
            required
            fullWidth
            name="password"
            label="password"
            type="password"
            id="password"
            autoComplete="current-password"
          />
          {error && (
            <Typography variant="body2" color="error" sx={{ mt: 1 }}>
              {error}
            </Typography>
          )}
          <Button
            type="submit"
            fullWidth
            variant="contained"
            sx={{ mt: 3, mb: 2 }}
          >作成</Button>
          <Grid container>
            <Grid item>
              <Link href="/" variant="body2">{"ログイン"}</Link>
            </Grid>
          </Grid>
        </Box>
      </Box>
    </Container>
  );
}