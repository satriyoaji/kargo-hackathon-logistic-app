import './App.scss';
import 'antd/dist/antd.css';
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import Login from './pages/login';
import Shipper from './pages/shipper/index'
import Transporter from './pages/transporter/index'
import TruckPage from './pages/transporter/truckPage';

function App() {
  return (
    <BrowserRouter>
      <div className="app">
        <Routes>
          <Route path="/login" element={<Login/>}></Route>
          <Route path='/login' element={<Login />}></Route>
          <Route exact path='/' element={<Navigate to='/login' />} />
          <Route path='/transporter' element={<Transporter />}></Route>
          <Route path='/shipper' element={<Shipper />}></Route>
          <Route path="/transporter/trucks" element={<TruckPage/>}></Route>
        </Routes>
      </div>
    </BrowserRouter>
  );
}

export default App;