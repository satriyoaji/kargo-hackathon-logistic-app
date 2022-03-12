import './App.scss';
import 'antd/dist/antd.css';
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import Login from './pages/login';
import DriverPage from './pages/transporter/driverPage';
import AddDriverPage from './pages/transporter/addDriver';

function App() {
  return (
    <BrowserRouter>
      <div className="app">
        <Routes>
          <Route path="/login" element={<Login/>}></Route>
          {/* YOUR CODE HERE */}

          <Route
            exact
            path="/"
            element={<Navigate to="/login" />}
          />
          <Route path="/login/driver" element={<DriverPage/>}></Route>
          <Route path="/login/driver/add" element={<AddDriverPage/>}></Route>
        </Routes>
      </div>
    </BrowserRouter>
  );
}

export default App;