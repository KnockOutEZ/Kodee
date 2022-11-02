import './App.css'
import { MemoryRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import Auth from './layouts/auth';
import Dashboard from './layouts/dashboard';
function App() {
    return (
      <Routes>
      <Route path="/dashboard/*" element={<Dashboard />} />
      <Route path="/auth/*" element={<Auth />} />
      <Route path="*" element={<Navigate to="/dashboard/home" replace />} />
    </Routes>
    )
}

export default App
