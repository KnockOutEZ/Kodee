import './App.css'
import { MemoryRouter as Router, Routes, Route } from 'react-router-dom';
import Home from './views/Home';
import Todo from './views/Todo';
function App() {
// const Todo = require('./views/Todo');
    return (
        <>
      <Router>
        <Routes>
          <Route path="/" element={<Home/>} />
          <Route path="/todo" element={<Todo/>} />
        </Routes>
      </Router>
    </>
    )
}

export default App
