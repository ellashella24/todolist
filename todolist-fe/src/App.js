import React from 'react';
import './App.css';
import Create from './components/create/create';
import Read from './components/read/read';
import { BrowserRouter as Router, Route, Routes  } from 'react-router-dom'

function App() {
  return (
    <div className='main'>
      <h3>Todolist</h3>
        <Router>
          <Routes>
            <Route path="/create" element={<Create/>} />
            <Route path="/" element={<Read/>} />
          </Routes>
        </Router>
    </div>
  );
}

export default App;
