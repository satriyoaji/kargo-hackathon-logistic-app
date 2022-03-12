import './App.scss';
import 'antd/dist/antd.css';
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import Login from './pages/login';
import Shipper from './pages/shipper/index'
import TruckPage from './pages/transporter/truckPage';
import './App.scss'
import 'antd/dist/antd.css'

function App() {
  return (
    <BrowserRouter>
      <div className='app'>
        <Routes>
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
