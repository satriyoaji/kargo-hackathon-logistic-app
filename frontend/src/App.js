import './App.scss';
import 'antd/dist/antd.css';
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import Login from './pages/login';
import DriverPage from './pages/transporter/driverPage';
import AddDriverPage from './pages/transporter/addDriver';
import Shipper from './pages/shipper/index'
import TruckPage from './pages/transporter/truckPage';
import './App.scss'
import 'antd/dist/antd.css'

function App() {
  return (
    <BrowserRouter>
      <div className='app'>
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
          <Route path='/login' element={<Login />}></Route>
          <Route exact path='/' element={<Navigate to='/login' />} />
          <Route path='/shipper' element={<Shipper />}></Route>
          <Route path="/transporter/trucks" element={<TruckPage/>}></Route>
        </Routes>
      </div>
    </BrowserRouter>
  )
}

export default App
