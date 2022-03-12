import { Space, Typography } from 'antd'

// Routing imports
import { Link } from 'react-router-dom'

import './style.scss'

// This is mostly cover what ticket #1 is all about
function App() {
  return (
    <div className='login'>
      <div className='login-header'>
        <div className='login-header-wrapper'>
          <Space direction='vertical'>
            <Typography.Title code>Kargo TMS</Typography.Title>
            <h2>Welcome to Kargo!</h2>
            <Link to='/trasnsporter'>
              <button>Transporter</button>
            </Link>
            <Link to='/shipper'>
              <button>Shipper</button>
            </Link>
          </Space>
        </div>
      </div>
    </div>
  )
}

export default App
