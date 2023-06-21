import React from 'react';
import { BrowserRouter,Routes, Route } from 'react-router-dom';

import SignIn from './SignIn.jsx';
import SignUp from './SignUp.jsx';
import Dashboard from './Chat.jsx';
const Index = () =>{
  return (
    <div>
      <ul>
        <li><a href='/signin'>signin</a></li>
        <li><a href='/signup'>signup</a></li>
        <li><a href='/chat'>chat</a></li>
      </ul>
    </div>
  )
}

const App = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path='/' element={<Index />} />
        <Route path='/signin' element={<SignIn />} />
        <Route path='/signup' element={<SignUp />} />
        <Route path='/chat' element={<Dashboard />} />
      </Routes>
    </BrowserRouter>
  );
};
export default App;