import _default from 'antd/lib/time-picker'
import React, { useState, useEffect } from 'react'

const UpdateStatus = ({ id, getShowStatus }) => {
  const [newStatus, setNewStatus] = useState('')
  const handleCancel = (event) => {
    event.preventDefault()
    getShowStatus(false)
  }
  return (
    <div className='modalForm'>
      <h3>Update Status: {id}</h3>
      <div>
        <label>Status: </label>
        <select>
          <option>Ongoing to Origin</option>
          <option>At Origin</option>
          <option>Ongoing to Destination</option>
          <option>At Destination</option>
          <option>Completed</option>
        </select>
      </div>
      <button className='modalButton'>Update</button>
      <button className='modalButton' onClick={handleCancel}>
        Cancel
      </button>
    </div>
  )
}

export default UpdateStatus
