import React from 'react'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
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

const router = createBrowserRouter([
  {
    path: '/',
    element: <SignIn />,
  },
  {
    path: '/signup',
    element: <SignUp />,
  },
  {
    path: '/chat',
    element: <Chat />,
  },
])

const App = () => {
  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Toaster
        toastOptions={{
          position: 'top-right',
        }}
      />
      <RouterProvider router={router} />
    </ThemeProvider>
  )
}
export default App