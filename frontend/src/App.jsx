import React from 'react'
import { BrowserRouter, Routes, Route } from 'react-router-dom'
import { ThemeProvider, createTheme } from '@mui/material/styles'
import CssBaseline from '@mui/material/CssBaseline'
import { Toaster } from 'react-hot-toast'

import SignIn from './pages/SignIn.jsx'
import SignUp from './pages/SignUp.jsx'
import Chat from './pages/Chat.jsx'

const theme = createTheme({
  palette: {
    primary: {
      main: '#2C3333',
    },
    secondary: {
      main: '#395B64',
    },
    error: {
      main: '#E7F6F2',
    },
    warning: {
      main: '#E7F6F2',
    },
    info: {
      main: '#A5C9CA',
    },
    success: {
      main: '#A5C9CA',
    },
  },
})

const App = () => {
  return (
    <BrowserRouter>
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <Toaster
          toastOptions={{
            position: 'top-right',
          }}
        />
        <Routes>
          <Route path="/" element={<SignIn />} />
          <Route path="/signup" element={<SignUp />} />
          <Route path="/chat" element={<Chat />} />
        </Routes>
      </ThemeProvider>
    </BrowserRouter>
  )
}
export default App
