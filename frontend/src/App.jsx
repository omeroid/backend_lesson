import React from 'react';
import { BrowserRouter,Routes, Route } from 'react-router-dom';

import SignIn from './SignIn.jsx';
import SignUp from './SignUp.jsx';
import Dashboard from './Chat.jsx';

const App = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path='/' element={<SignIn />} />
        <Route path='/signup' element={<SignUp />} />
        <Route path='/chat' element={<Dashboard />} />
      </Routes>
    </BrowserRouter>
  );
};
export default App;